//go:generate go run go.uber.org/mock/mockgen -typed -package mock -source client.go -destination mock/client.go
package internal

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/distribution/reference"
	cbuild "github.com/docker/buildx/controller/build"
	controllerapi "github.com/docker/buildx/controller/pb"
	"github.com/docker/buildx/util/progress"
	"github.com/docker/cli/cli/command"
	"github.com/docker/cli/cli/config"
	cfgtypes "github.com/docker/cli/cli/config/types"
	"github.com/docker/cli/cli/flags"
	manifesttypes "github.com/docker/cli/cli/manifest/types"
	"github.com/docker/docker/api/types"
	registrytypes "github.com/docker/docker/api/types/registry"
	"github.com/moby/buildkit/client"
	"github.com/moby/buildkit/util/progress/progressui"
	cp "github.com/otiai10/copy"

	provider "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi/sdk/v3/go/common/diag"

	"github.com/pulumi/pulumi-docker/provider/v4/internal/properties"
)

// Client handles all our Docker API calls.
type Client interface {
	Auth(ctx context.Context, creds properties.RegistryAuth) error
	Build(ctx provider.Context, opts controllerapi.BuildOptions) (*client.SolveResponse, error)
	Close(ctx context.Context) error
	BuildKitEnabled() (bool, error)
	Inspect(ctx context.Context, id string) ([]manifesttypes.ImageManifest, error)
	Delete(ctx context.Context, id string) ([]types.ImageDeleteResponseItem, error)
}

type docker struct {
	mu sync.Mutex // Guards changes to configs.

	cli *command.DockerCli
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

	return &docker{cli: cli, dir: dir}, err
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
	ctx provider.Context,
	opts controllerapi.BuildOptions,
) (*client.SolveResponse, error) {
	// Create a pipe to forward buildx output to our HostClient.
	r, w, err := os.Pipe()
	if err != nil {
		return nil, fmt.Errorf("creating pipe: %w", err)
	}
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ctx.LogStatus(diag.Info, s.Text())
		}
	}()
	printer, err := progress.NewPrinter(ctx, w, progressui.PlainMode)
	if err != nil {
		return nil, fmt.Errorf("creating printer: %w", err)
	}

	// Perform the build.
	solve, res, err := cbuild.RunBuild(ctx, d.cli, opts, d.cli.In(), printer, true)
	if res != nil {
		res.Done()
	}
	return solve, err
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
