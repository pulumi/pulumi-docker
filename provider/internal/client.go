//go:generate go run go.uber.org/mock/mockgen -package mock -source client.go -destination mock/client.go
package internal

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/distribution/reference"
	cbuild "github.com/docker/buildx/controller/build"
	controllerapi "github.com/docker/buildx/controller/pb"
	"github.com/docker/buildx/util/progress"
	"github.com/docker/cli/cli/command"
	"github.com/docker/cli/cli/flags"
	manifesttypes "github.com/docker/cli/cli/manifest/types"
	"github.com/docker/docker/api/types"
	"github.com/moby/buildkit/client"
	"github.com/moby/buildkit/util/progress/progressui"
)

// Client handles all our Docker API calls.
type Client interface {
	Build(ctx context.Context, opts controllerapi.BuildOptions) (*client.SolveResponse, error)
	BuildKitEnabled() (bool, error)
	Inspect(ctx context.Context, id string) ([]manifesttypes.ImageManifest, error)
	Delete(ctx context.Context, id string) ([]types.ImageDeleteResponseItem, error)
}

type docker struct {
	cli *command.DockerCli
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

	return &docker{cli: cli}, err
}

// Build performs a buildkit build.
func (d *docker) Build(
	ctx context.Context,
	opts controllerapi.BuildOptions,
) (*client.SolveResponse, error) {
	printer, err := progress.NewPrinter(ctx, os.Stdout, progressui.PlainMode)
	if err != nil {
		return nil, fmt.Errorf("creating printer: %w", err)
	}
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
