//go:generate go run go.uber.org/mock/mockgen -typed -package internal -source client.go -destination mockclient_test.go --self_package github.com/pulumi/pulumi-docker/provider/v4/internal
package internal

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/distribution/reference"
	buildx "github.com/docker/buildx/build"
	"github.com/docker/buildx/builder"
	controllerapi "github.com/docker/buildx/controller/pb"
	"github.com/docker/buildx/store"
	"github.com/docker/buildx/store/storeutil"
	"github.com/docker/buildx/util/dockerutil"
	"github.com/docker/buildx/util/platformutil"
	"github.com/docker/buildx/util/progress"
	"github.com/docker/cli/cli/command"
	"github.com/docker/cli/cli/config"
	cfgtypes "github.com/docker/cli/cli/config/types"
	"github.com/docker/cli/cli/flags"
	manifesttypes "github.com/docker/cli/cli/manifest/types"
	"github.com/docker/docker/api/types"
	registrytypes "github.com/docker/docker/api/types/registry"
	"github.com/moby/buildkit/client"
	"github.com/moby/buildkit/session"
	"github.com/moby/buildkit/session/auth/authprovider"
	"github.com/moby/buildkit/util/progress/progressui"
	cp "github.com/otiai10/copy"

	provider "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi/sdk/v3/go/common/diag"

	"github.com/pulumi/pulumi-docker/provider/v4/internal/properties"
)

// Client handles all our Docker API calls.
type Client interface {
	Auth(ctx context.Context, creds properties.RegistryAuth) error
	Build(ctx provider.Context, b Build) (map[string]*client.SolveResponse, error)
	Close(ctx context.Context) error
	BuildKitEnabled() (bool, error)
	Inspect(ctx context.Context, id string) ([]manifesttypes.ImageManifest, error)
	Delete(ctx context.Context, id string) ([]types.ImageDeleteResponseItem, error)
}

type Build interface {
	BuildOptions() controllerapi.BuildOptions
	Targets() []string
}

type docker struct {
	mu sync.Mutex

	cli *command.DockerCli
	txn *store.Txn
	dir string
}

var _ Client = (*docker)(nil)

func newDockerClient() (*docker, error) {
	cli, err := command.NewDockerCli(
		command.WithCombinedStreams(os.Stdout),
	)
	if err != nil {
		return nil, err
	}

	// We create a temporary directory for our config to not disturb the host's
	// existing settings.
	dir, err := os.MkdirTemp("", "pulumi-docker-")
	if err != nil {
		return nil, err
	}
	// Attempt to copy the host's existing config, if it exists, over to our
	// temporary config directory. This ensures we preserve things like
	// credential helpers, builders, etc.
	if _, serr := os.Stat(config.Dir()); serr == nil {
		_ = cp.Copy(config.Dir(), dir)
	}

	opts := &flags.ClientOptions{
		ConfigDir: dir,
		// TODO(github.com/pulumi/pulumi-docker/issues/946): Support TLS options
	}
	err = cli.Initialize(opts)
	if err != nil {
		return nil, err
	}

	txn, _, err := storeutil.GetStore(cli)
	if err != nil {
		return nil, err
	}

	return &docker{cli: cli, txn: txn, dir: dir}, err
}

func (d *docker) Auth(ctx context.Context, creds properties.RegistryAuth) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	cfg := d.cli.ConfigFile()

	// Special handling for legacy DockerHub domains. The OCI-compliant
	// registry is registry-1.docker.io but this is stored in config under the
	// legacy name.
	// https://github.com/docker/cli/issues/3793#issuecomment-1269051403
	if creds.Address == "docker.io" {
		creds.Address = "https://index.docker.io/v1/"
	}

	auth := cfgtypes.AuthConfig{
		ServerAddress: creds.Address,
		Username:      creds.Username,
		Password:      creds.Password,
	}

	// Workaround for https://github.com/docker/docker-credential-helpers/issues/37.
	if existing, err := cfg.GetAuthConfig(creds.Address); err == nil {
		// Confirm the auth is still valid. Otherwise we'll set it to the
		// provided config.
		if existing.Username == creds.Username {
			_, err = d.cli.Client().RegistryLogin(ctx, registrytypes.AuthConfig{
				Auth:          existing.Auth,
				Email:         existing.Email,
				IdentityToken: existing.IdentityToken,
				Password:      existing.Password,
				RegistryToken: existing.RegistryToken,
				ServerAddress: creds.Address, // ServerAddress is sometimes empty?
				Username:      existing.Username,
			})
			if err == nil {
				return nil // Creds still work, nothing to do.
			}
		}
	}

	err := cfg.GetCredentialsStore(creds.Address).Store(auth)
	if err != nil {
		return fmt.Errorf("%q: %w", creds.Address, err)
	}
	return nil
}

// Build performs a buildkit build.
func (d *docker) Build(
	pctx provider.Context,
	build Build,
) (map[string]*client.SolveResponse, error) {
	// Use a separate context for the build. We don't want to kill our request's
	// context if the build fails.
	bctx := context.Background()
	opts := build.BuildOptions()

	// Create a pipe to forward buildx output to our HostClient.
	r, w, err := os.Pipe()
	if err != nil {
		return nil, fmt.Errorf("creating pipe: %w", err)
	}
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Fprintf(os.Stderr, "Panic recovered: %s", err)
			}
		}()
		s := bufio.NewScanner(r)
		for s.Scan() {
			pctx.LogStatus(diag.Info, s.Text())
		}
	}()

	b, nodes, err := d.builder(opts)
	if err != nil {
		return nil, err
	}
	printer, err := progress.NewPrinter(bctx, w,
		progressui.PlainMode,
		progress.WithDesc(
			fmt.Sprintf("building with %q instance using %s driver", b.Name, b.Driver),
			fmt.Sprintf("%s:%s", b.Driver, b.Name),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("creating printer: %w", err)
	}

	cacheFrom := []client.CacheOptionsEntry{}
	for _, c := range opts.CacheFrom {
		cacheFrom = append(cacheFrom, client.CacheOptionsEntry{
			Type:  c.Type,
			Attrs: c.Attrs,
		})
	}
	cacheTo := []client.CacheOptionsEntry{}
	for _, c := range opts.CacheTo {
		cacheTo = append(cacheTo, client.CacheOptionsEntry{
			Type:  c.Type,
			Attrs: c.Attrs,
		})
	}
	exports := []client.ExportEntry{}
	for _, e := range opts.Exports {
		exports = append(exports, client.ExportEntry{
			Type:      e.Type,
			Attrs:     e.Attrs,
			OutputDir: e.Destination,
		})
	}
	platforms, _ := platformutil.Parse(opts.Platforms)
	platforms = platformutil.Dedupe(platforms)

	namedContexts := map[string]buildx.NamedContext{}
	for k, v := range opts.NamedContexts {
		ref, err := reference.ParseNormalizedNamed(k)
		if err != nil {
			return nil, err
		}
		name := strings.TrimSuffix(reference.FamiliarString(ref), ":latest")
		namedContexts[name] = buildx.NamedContext{Path: v}
	}

	payload := map[string]buildx.Options{}
	for _, target := range build.Targets() {
		targetName := target
		if target == "" {
			targetName = "default"
		}
		payload[targetName] = buildx.Options{
			Inputs: buildx.Inputs{
				ContextPath:    opts.ContextPath,
				DockerfilePath: opts.DockerfileName,
				NamedContexts:  namedContexts,
			},
			BuildArgs: opts.BuildArgs,
			CacheFrom: cacheFrom,
			CacheTo:   cacheTo,
			Exports:   exports,
			NoCache:   opts.NoCache,
			Platforms: platforms,
			Pull:      opts.Pull,
			Tags:      opts.Tags,
			Target:    target,

			Session: []session.Attachable{
				authprovider.NewDockerAuthProvider(d.cli.ConfigFile(), nil),
			},
		}
	}

	// Perform the build.
	results, err := buildx.Build(
		bctx,
		nodes,
		payload,
		dockerutil.NewClient(d.cli),
		d.dir,
		printer,
	)
	if err != nil {
		return nil, err
	}

	for _, w := range printer.Warnings() {
		b := &bytes.Buffer{}
		fmt.Fprintf(b, "%s", w.Short)
		for _, d := range w.Detail {
			fmt.Fprintf(b, "\n%s", d)
		}
		pctx.Log(diag.Warning, b.String())
	}

	return results, err
}

// BuildKitEnabled returns true if the client supports buildkit.
func (d *docker) BuildKitEnabled() (bool, error) {
	return d.cli.BuildKitEnabled()
}

// Close cleans up temporary configs.
func (d *docker) Close(_ context.Context) error {
	return os.RemoveAll(d.dir)
}

// Inspect inspects an image.
func (d *docker) Inspect(ctx context.Context, id string) ([]manifesttypes.ImageManifest, error) {
	ref, err := normalizeReference(id)
	if err != nil {
		return []manifesttypes.ImageManifest{}, err
	}

	rc := d.cli.RegistryClient(d.cli.DockerEndpoint().SkipTLSVerify)
	manifests, err := rc.GetManifestList(ctx, ref)

	if err != nil && strings.Contains(err.Error(), "unsupported manifest format") {
		manifest, err := rc.GetManifest(ctx, ref)
		manifests = append(manifests, manifest)
		return manifests, err
	}

	return manifests, err
}

// Delete deletes an image with the given ID.
func (d *docker) Delete(ctx context.Context, id string) ([]types.ImageDeleteResponseItem, error) {
	return d.cli.Client().ImageRemove(ctx, id, types.ImageRemoveOptions{
		Force: true, // Needed in case the image has multiple tags.
	})
}

func normalizeReference(ref string) (reference.Named, error) {
	namedRef, err := reference.ParseNormalizedNamed(ref)
	if err != nil {
		return nil, err
	}
	if _, isDigested := namedRef.(reference.Canonical); !isDigested {
		return reference.TagNameOnly(namedRef), nil
	}
	return namedRef, nil
}

// builder ensures a builder is available and running. This is guarded by a
// mutex to ensure other resources don't attempt to use the builder until it's
// ready.
func (d *docker) builder(
	opts controllerapi.BuildOptions,
) (*builder.Builder, []builder.Node, error) {
	d.mu.Lock()
	defer d.mu.Unlock()

	contextPathHash := opts.ContextPath
	if absContextPath, err := filepath.Abs(contextPathHash); err == nil {
		contextPathHash = absContextPath
	}
	b, err := builder.New(d.cli,
		builder.WithName(opts.Builder),
		builder.WithContextPathHash(contextPathHash),
		builder.WithStore(d.txn),
	)
	if err != nil {
		return nil, nil, err
	}

	nodes, err := b.LoadNodes(context.Background())
	if err != nil {
		return nil, nil, err
	}

	return b, nodes, nil
}
