//go:generate go run go.uber.org/mock/mockgen -package mock -source client.go -destination mock/client.go
package internal

import (
	"context"
	"fmt"
	"os"

	cbuild "github.com/docker/buildx/controller/build"
	controllerapi "github.com/docker/buildx/controller/pb"
	"github.com/docker/buildx/util/progress"
	"github.com/docker/cli/cli/command"
	"github.com/docker/cli/cli/flags"
	"github.com/docker/docker/api/types"
	"github.com/moby/buildkit/client"
	"github.com/moby/buildkit/util/progress/progressui"
)

// Client handles all our Docker API calls.
type Client interface {
	Build(ctx context.Context, opts controllerapi.BuildOptions) (*client.SolveResponse, error)
	BuildKitEnabled() (bool, error)
	Inspect(ctx context.Context, id string) (types.ImageInspect, error)
	Delete(ctx context.Context, id string) ([]types.ImageDeleteResponseItem, error)
}

type docker struct {
	cli *command.DockerCli
}

func newDockerClient() (Client, error) {
	cli, err := command.NewDockerCli(
		command.WithCombinedStreams(os.Stdout),
	)
	if err != nil {
		return nil, err
	}

	err = cli.Initialize(flags.NewClientOptions())
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
//
// TODO: Inspect the manifest instead?
func (d *docker) Inspect(ctx context.Context, id string) (types.ImageInspect, error) {
	inspect, _, err := d.cli.Client().ImageInspectWithRaw(ctx, id)
	return inspect, err
}

// Delete deletes an image with the given ID.
func (d *docker) Delete(ctx context.Context, id string) ([]types.ImageDeleteResponseItem, error) {
	return d.cli.Client().ImageRemove(ctx, id, types.ImageRemoveOptions{
		Force: true, // Needed in case the image has multiple tags.
	})
}
