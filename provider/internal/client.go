//go:generate go run go.uber.org/mock/mockgen -typed -package internal -source client.go -destination mockclient_test.go --self_package github.com/pulumi/pulumi-docker/provider/v4/internal
package internal

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/distribution/reference"
	buildx "github.com/docker/buildx/build"
	"github.com/docker/buildx/builder"
	controllerapi "github.com/docker/buildx/controller/pb"
	"github.com/docker/buildx/util/dockerutil"
	"github.com/docker/buildx/util/platformutil"
	"github.com/docker/buildx/util/progress"
	"github.com/docker/cli/cli/command"
	cfgtypes "github.com/docker/cli/cli/config/types"
	"github.com/docker/cli/cli/flags"
	manifesttypes "github.com/docker/cli/cli/manifest/types"
	"github.com/docker/docker/api/types"
	"github.com/moby/buildkit/client"
	"github.com/moby/buildkit/session"
	"github.com/moby/buildkit/util/progress/progressui"

	provider "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi/sdk/v3/go/common/diag"
)

// Client handles all our Docker API calls.
type Client interface {
	Auth(ctx context.Context, name string, creds RegistryAuth) error
	Build(ctx provider.Context, name string, b Build) (map[string]*client.SolveResponse, error)
	BuildKitEnabled() (bool, error)
	Inspect(ctx context.Context, id string) ([]manifesttypes.ImageManifest, error)
	Delete(ctx context.Context, id string) ([]types.ImageDeleteResponseItem, error)
}

type Build interface {
	BuildOptions() controllerapi.BuildOptions
	Targets() []string
	Inline() string
	Secrets() session.Attachable
}

type docker struct {
	mu sync.Mutex

	cli   *command.DockerCli
	auths map[string]map[string]cfgtypes.AuthConfig
}

var _ Client = (*docker)(nil)

var _baseAuth = ""

func newDockerClient() (*docker, error) {
	cli, err := command.NewDockerCli(
		command.WithCombinedStreams(os.Stdout),
	)
	if err != nil {
		return nil, err
	}

	opts := &flags.ClientOptions{
		// TODO(github.com/pulumi/pulumi-docker/issues/946): Support TLS options
	}
	err = cli.Initialize(opts)
	if err != nil {
		return nil, err
	}

	d := &docker{cli: cli, auths: map[string]map[string]cfgtypes.AuthConfig{}}

	// Load existing credentials into memory.
	creds, err := cli.ConfigFile().GetAllCredentials()
	if err != nil {
		return nil, err
	}
	for _, cred := range creds {
		err := d.Auth(context.Background(), _baseAuth, RegistryAuth{
			Address:  cred.ServerAddress,
			Username: cred.Username,
			Password: cred.Password,
		})
		if err != nil {
			return nil, err
		}
	}

	return d, nil
}

func (d *docker) Auth(_ context.Context, name string, creds RegistryAuth) error {
	d.mu.Lock()
	defer d.mu.Unlock()


	if _, ok := d.auths[name]; !ok {
		d.auths[name] = map[string]cfgtypes.AuthConfig{}
	}

	// Special handling for legacy DockerHub domains. The OCI-compliant
	// registry is registry-1.docker.io but this is stored in config under the
	// legacy name.
	// https://github.com/docker/cli/issues/3793#issuecomment-1269051403

	u, err := url.Parse(creds.Address)
	if err != nil {
		return err
	}
	key := u.Host

	if creds.Address == "registry-1.docker.io" || creds.Address == "https://index.docker.io/v1/" {
		key = "https://index.docker.io/v1/"
	}

	auth := cfgtypes.AuthConfig{
		ServerAddress: creds.Address,
		Username:      creds.Username,
		Password:      creds.Password,
	}

	if _, ok := d.auths[name][key]; ok {
		return nil // Already saved these creds. Nothing to do.
	}

	d.auths[name][key] = auth

	return nil
}

// Build performs a buildkit build. Returns a map of target names (or one name,
// "default", if no targets were specified) to SolveResponses, which capture
// the build's digest and tags (if any).
func (d *docker) Build(
	pctx provider.Context,
	name string,
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

	auths := map[string]*cfgtypes.AuthConfig{}
	for host, cfg := range d.auths[_baseAuth] {
		cfg := cfg
		auths[host] = &cfg
	}
	for host, cfg := range d.auths[name] {
		cfg := cfg
		auths[host] = &cfg
	}

	payload := map[string]buildx.Options{}
	for _, target := range build.Targets() {
		targetName := target
		if target == "" {
			targetName = "default"
		}
		payload[targetName] = buildx.Options{
			Inputs: buildx.Inputs{
				ContextPath:      opts.ContextPath,
				DockerfilePath:   opts.DockerfileName,
				DockerfileInline: build.Inline(),
				NamedContexts:    namedContexts,
				InStream:         strings.NewReader(""),
			},
			BuildArgs: opts.BuildArgs,
			CacheFrom: cacheFrom,
			CacheTo:   cacheTo,
			Exports:   exports,
			NoCache:   opts.NoCache,
			Labels:    opts.Labels,
			Platforms: platforms,
			Pull:      opts.Pull,
			Tags:      opts.Tags,
			Target:    target,

			Session: []session.Attachable{
				build.Secrets(),
				newInMemoryAuthProvider(auths, printer.Write, nil),
			},
		}
	}

	// Perform the build.
	results, err := buildx.Build(
		bctx,
		nodes,
		payload,
		dockerutil.NewClient(d.cli),
		filepath.Dir(d.cli.ConfigFile().Filename),
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
