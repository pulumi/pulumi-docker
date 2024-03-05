//go:generate go run go.uber.org/mock/mockgen -typed -package mock -source client.go -destination mock/client.go
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
	"github.com/docker/buildx/build"
	"github.com/docker/buildx/builder"
	controllerapi "github.com/docker/buildx/controller/pb"
	"github.com/docker/buildx/util/dockerutil"
	"github.com/docker/buildx/util/platformutil"
	"github.com/docker/buildx/util/progress"
	"github.com/docker/cli/cli/command"
	"github.com/docker/cli/cli/config/configfile"
	"github.com/docker/cli/cli/config/credentials"
	cfgtypes "github.com/docker/cli/cli/config/types"
	"github.com/docker/cli/cli/flags"
	manifesttypes "github.com/docker/cli/cli/manifest/types"
	registryclient "github.com/docker/cli/cli/registry/client"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/registry"
	registryconst "github.com/docker/docker/registry"
	"github.com/moby/buildkit/client"
	"github.com/moby/buildkit/session"
	"github.com/moby/buildkit/session/auth/authprovider"
	"github.com/moby/buildkit/util/progress/progressui"
	"github.com/sirupsen/logrus"

	provider "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi/sdk/v3/go/common/diag"

	"github.com/pulumi/pulumi-docker/provider/v4/internal/properties"
)

// Client handles all our Docker API calls.
type Client interface {
	Auth(ctx context.Context, name string, creds properties.RegistryAuth) error
	Build(ctx provider.Context, name string, opts controllerapi.BuildOptions) (*client.SolveResponse, error)
	BuildKitEnabled() (bool, error)
	Inspect(ctx context.Context, name string, id string) ([]manifesttypes.ImageManifest, error)
	Delete(ctx context.Context, id string) ([]image.DeleteResponse, error)
}

type docker struct {
	mu sync.Mutex

	cli      *command.DockerCli
	auths    map[string]map[string]cfgtypes.AuthConfig
	builders map[string]*cachedBuilder
}

var _ Client = (*docker)(nil)

var _baseAuth = ""

func newDockerClient() (*docker, error) {
	cli, err := command.NewDockerCli(
		command.WithCombinedStreams(os.Stderr),
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
	// Disable the CLI's tendency to log randomly to stdout.
	logrus.SetLevel(logrus.ErrorLevel)

	d := &docker{
		cli:      cli,
		auths:    map[string]map[string]cfgtypes.AuthConfig{},
		builders: map[string]*cachedBuilder{},
	}

	// Load existing credentials into memory.
	creds, err := cli.ConfigFile().GetAllCredentials()
	if err != nil {
		return nil, err
	}
	for _, cred := range creds {
		err := d.Auth(context.Background(), _baseAuth, properties.RegistryAuth{
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

func (d *docker) Auth(_ context.Context, name string, creds properties.RegistryAuth) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	if _, ok := d.auths[name]; !ok {
		d.auths[name] = map[string]cfgtypes.AuthConfig{}
	}

	// Special handling for legacy DockerHub domains. The OCI-compliant
	// registry is registry-1.docker.io but this is stored in config under the
	// legacy name.
	// https://github.com/docker/cli/issues/3793#issuecomment-1269051403
	key := credentials.ConvertToHostname(creds.Address)
	if key == "registry-1.docker.io" || key == "index.docker.io" || key == "docker.io" {
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

// Build performs a buildkit build.
func (d *docker) Build(
	pctx provider.Context,
	name string,
	opts controllerapi.BuildOptions,
) (*client.SolveResponse, error) {
	ctx := context.Context(pctx)

	// Create a pipe to forward buildx output to our HostClient.
	doneLogging := make(chan struct{}, 1)
	r, w, err := os.Pipe()
	if err != nil {
		return nil, fmt.Errorf("creating pipe: %w", err)
	}
	defer func() {
		_ = w.Close()
		_ = r.Close()
		<-doneLogging
	}()

	go func() {
		defer func() {
			doneLogging <- struct{}{}
			if err := recover(); err != nil {
				fmt.Fprintf(os.Stderr, "Panic recovered: %s", err)
			}
		}()
		s := bufio.NewScanner(r)
		for s.Scan() {
			pctx.LogStatus(diag.Info, s.Text())
		}
		// Print a newline to clear confusing "DONE" statements.
		pctx.LogStatus(diag.Info, "")
	}()

	b, err := d.builder(opts)
	if err != nil {
		return nil, err
	}
	printer, err := progress.NewPrinter(ctx, w,
		progressui.PlainMode,
		progress.WithDesc(
			fmt.Sprintf("building with %q instance using %s driver", b.name, b.driver),
			fmt.Sprintf("%s:%s", b.driver, b.name),
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

	auths := map[string]cfgtypes.AuthConfig{}
	for host, cfg := range d.auths[_baseAuth] {
		auths[host] = cfg
	}
	for host, cfg := range d.auths[name] {
		auths[host] = cfg
	}

	// Perform the build.
	targets, err := build.Build(
		ctx,
		b.nodes,
		map[string]build.Options{
			opts.Target: {
				Inputs: build.Inputs{
					ContextPath:    opts.ContextPath,
					DockerfilePath: opts.DockerfileName,
				},
				BuildArgs: opts.BuildArgs,
				CacheFrom: cacheFrom,
				CacheTo:   cacheTo,
				Exports:   exports,
				NoCache:   opts.NoCache,
				Platforms: platforms,
				Pull:      opts.Pull,
				Tags:      opts.Tags,
				Target:    opts.Target,

				Session: []session.Attachable{
					authprovider.NewDockerAuthProvider(&configfile.ConfigFile{AuthConfigs: auths}, nil),
				},
			},
		},
		dockerutil.NewClient(d.cli),
		filepath.Dir(d.cli.ConfigFile().Filename),
		printer,
	)
	if err != nil {
		return nil, err
	}
	solve := targets[opts.Target]

	if printErr := printer.Wait(); printErr != nil {
		return solve, printErr
	}
	for _, w := range printer.Warnings() {
		b := &bytes.Buffer{}
		fmt.Fprintf(b, "%s", w.Short)
		for _, d := range w.Detail {
			fmt.Fprintf(b, "\n%s", d)
		}
		pctx.Log(diag.Warning, b.String())
	}

	return solve, err
}

// BuildKitEnabled returns true if the client supports buildkit.
func (d *docker) BuildKitEnabled() (bool, error) {
	return d.cli.BuildKitEnabled()
}

// Inspect inspects an image.
func (d *docker) Inspect(ctx context.Context, name string, id string) ([]manifesttypes.ImageManifest, error) {
	ref, err := normalizeReference(id)
	if err != nil {
		return []manifesttypes.ImageManifest{}, err
	}

	// Constructed a RegistryClient which can use our in-memory auth.
	insecure := d.cli.DockerEndpoint().SkipTLSVerify
	resolver := func(_ context.Context, index *registry.IndexInfo) registry.AuthConfig {
		configKey := index.Name
		if index.Official {
			configKey = registryconst.IndexServer
		}

		for _, scope := range []string{name, _baseAuth} {
			auths, ok := d.auths[scope]
			if !ok {
				continue
			}
			if a, ok := auths[configKey]; ok {
				return registry.AuthConfig(a)
			}
		}
		return registry.AuthConfig{}
	}
	rc := registryclient.NewRegistryClient(resolver, command.UserAgent(), insecure)

	manifests, err := rc.GetManifestList(ctx, ref)

	// If the registry doesn't support manifest lists, attempt to fetch an
	// individual one.
	if err != nil && strings.Contains(err.Error(), "unsupported manifest format") {
		manifest, err := rc.GetManifest(ctx, ref)
		manifests = append(manifests, manifest)
		return manifests, err
	}

	return manifests, err
}

// Delete deletes an image with the given ID.
func (d *docker) Delete(ctx context.Context, id string) ([]image.DeleteResponse, error) {
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
) (*cachedBuilder, error) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if b, ok := d.builders[opts.Builder]; ok {
		return b, nil
	}

	contextPathHash := opts.ContextPath
	if absContextPath, err := filepath.Abs(contextPathHash); err == nil {
		contextPathHash = absContextPath
	}
	b, err := builder.New(d.cli,
		builder.WithName(opts.Builder),
		builder.WithContextPathHash(contextPathHash),
	)
	if err != nil {
		return nil, err
	}

	// Need to load nodes in order to determine the builder's driver.
	nodes, err := b.LoadNodes(context.Background())
	if err != nil {
		return nil, err
	}

	cached := &cachedBuilder{name: b.Name, driver: b.Driver, nodes: nodes}
	d.builders[opts.Builder] = cached

	return cached, nil
}

// cachedBuilder caches the builders we've loaded. Repeatedly fetching them can
// sometimes result in EOF errors from the daemon, especially when under load.
type cachedBuilder struct {
	name   string
	driver string
	nodes  []builder.Node
}
