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
	"github.com/docker/buildx/builder"
	cbuild "github.com/docker/buildx/controller/build"
	controllerapi "github.com/docker/buildx/controller/pb"
	"github.com/docker/buildx/util/progress"
	"github.com/docker/cli/cli/command"
	"github.com/docker/cli/cli/flags"
	manifesttypes "github.com/docker/cli/cli/manifest/types"
	"github.com/docker/docker/api/types"
	"github.com/moby/buildkit/client"
	"github.com/moby/buildkit/util/progress/progressui"

	provider "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi/sdk/v3/go/common/diag"
)

// Client handles all our Docker API calls.
type Client interface {
	Build(ctx provider.Context, opts controllerapi.BuildOptions) (*client.SolveResponse, error)
	BuildKitEnabled() (bool, error)
	Inspect(ctx context.Context, id string) ([]manifesttypes.ImageManifest, error)
	Delete(ctx context.Context, id string) ([]types.ImageDeleteResponseItem, error)
}

type docker struct {
	mu sync.Mutex

	cli      *command.DockerCli
	builders map[string]*cachedBuilder
}

var _ Client = (*docker)(nil)

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

	return &docker{cli: cli, builders: map[string]*cachedBuilder{}}, err
}

// Build performs a buildkit build.
func (d *docker) Build(
	pctx provider.Context,
	opts controllerapi.BuildOptions,
) (*client.SolveResponse, error) {
	// Use a seprate context for the build. We don't want to kill our request's
	// context if the build fails.
	bctx := context.Background()

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

	b, err := d.builder(opts)
	if err != nil {
		return nil, err
	}
	printer, err := progress.NewPrinter(bctx, w,
		progressui.PlainMode,
		progress.WithDesc(
			fmt.Sprintf("building with %q instance using %s driver", b.name, b.driver),
			fmt.Sprintf("%s:%s", b.driver, b.name),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("creating printer: %w", err)
	}

	// Perform the build.
	solve, res, err := cbuild.RunBuild(bctx, d.cli, opts, d.cli.In(), printer, true)
	if res != nil {
		res.Done()
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
