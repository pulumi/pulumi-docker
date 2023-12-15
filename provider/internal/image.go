package internal

import (
	"context"
	"fmt"
	"slices"

	// These imports are needed to register the drivers with buildkit.
	_ "github.com/docker/buildx/driver/docker"
	_ "github.com/docker/buildx/driver/docker-container"
	_ "github.com/docker/buildx/driver/kubernetes"
	_ "github.com/docker/buildx/driver/remote"

	controllerapi "github.com/docker/buildx/controller/pb"
	"github.com/docker/buildx/util/buildflags"
	"github.com/docker/docker/errdefs"
	"github.com/moby/buildkit/exporter/containerimage/exptypes"
	"github.com/muesli/reflow/dedent"

	provider "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/diag"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

var (
	_ infer.Annotated                             = (infer.Annotated)((*Image)(nil))
	_ infer.Annotated                             = (infer.Annotated)((*ImageArgs)(nil))
	_ infer.Annotated                             = (infer.Annotated)((*ImageState)(nil))
	_ infer.CustomCheck[ImageArgs]                = (*Image)(nil)
	_ infer.CustomDelete[ImageState]              = (*Image)(nil)
	_ infer.CustomRead[ImageArgs, ImageState]     = (*Image)(nil)
	_ infer.CustomResource[ImageArgs, ImageState] = (*Image)(nil)
)

// Image is a Docker image build using buildkit.
type Image struct{}

// Annotate provides a description of the Image resource.
func (i *Image) Annotate(a infer.Annotator) {
	a.Describe(&i, "A Docker image built using Buildkit")
}

// ImageArgs instantiates a new Image.
type ImageArgs struct {
	Context []string `pulumi:"context,optional"`
	Exports []string `pulumi:"exports,optional"`
	File    string   `pulumi:"file,optional"`
	Tags    []string `pulumi:"tags"`
}

// Annotate describes inputs to the Image resource.
func (ia *ImageArgs) Annotate(a infer.Annotator) {
	a.Describe(&ia.Context, dedent.String(`
		Contexts to use while building the image. If omitted, an empty context
		is used. If more than one value is specified, they should be of the
		form "name=value".`,
	))
	a.Describe(&ia.Exports, dedent.String(`
		Name and optionally a tag (format: "name:tag"). If outputting to a
		registry, the name should include the fully qualified registry address.`,
	))
	a.Describe(&ia.File, dedent.String(`
		Name of the Dockerfile to use (default: "$PATH/Dockerfile").`,
	))
	a.Describe(&ia.Tags, dedent.String(`
		Name and optionally a tag (format: "name:tag"). If outputting to a
		registry, the name should include the fully qualified registry address.`,
	))

	a.SetDefault(&ia.File, "Dockerfile")
	// TODO: SetDefault host platform.
}

// ImageState is serialized to the program's state file.
type ImageState struct {
	ImageArgs

	Architecture string   `pulumi:"architecture,optional"`
	OS           string   `pulumi:"os,optional"`
	RepoDigests  []string `pulumi:"repoDigests,optional"`
	RepoTags     []string `pulumi:"repoTags,optional"`
	Size         int64    `pulumi:"size,optional"`
}

// Annotate describes outputs of the Image resource.
func (is *ImageState) Annotate(a infer.Annotator) {
	is.ImageArgs.Annotate(a)
}

// Check validates ImageArgs and sets defaults.
func (*Image) Check(
	_ provider.Context,
	_ string,
	_ resource.PropertyMap,
	news resource.PropertyMap,
) (ImageArgs, []provider.CheckFailure, error) {
	args, failures, err := infer.DefaultCheck[ImageArgs](news)
	if err != nil || len(failures) != 0 {
		return args, failures, err
	}
	if len(args.Tags) == 0 {
		failures = append(failures,
			provider.CheckFailure{Property: "tags", Reason: "at least one tag is required"},
		)
	}
	if _, err := buildflags.ParseExports(args.Exports); err != nil {
		failures = append(failures,
			provider.CheckFailure{Property: "exports", Reason: err.Error()},
		)
	}

	if args.File == "" {
		args.File = "Dockerfile"
	}

	return args, failures, err
}

// Create builds the image using buildkit.
func (i *Image) Create(
	ctx provider.Context,
	name string,
	input ImageArgs,
	preview bool,
) (string, ImageState, error) {
	cfg := infer.GetConfig[Config](ctx)

	state := ImageState{
		ImageArgs: input,
	}

	ok, err := cfg.client.BuildKitEnabled()
	if err != nil {
		return name, state, fmt.Errorf("checking buildkit compatibility: %w", err)
	}
	if !ok {
		return name, state, fmt.Errorf("buildkit is not supported on this host")
	}

	exports, err := buildflags.ParseExports(input.Exports)
	if err != nil {
		return name, state, fmt.Errorf("parsing exports: %w", err)
	}
	opts := controllerapi.BuildOptions{
		DockerfileName: input.File,
		ContextPath:    input.Context[0],
		Tags:           input.Tags,
		Exports:        exports,
		// CacheFrom:      parsedCacheFrom,
		// CacheTo:        parsedCacheTo,
		// BuildArgs:      args,
		// Platforms:      []string{build.Platform},
		// Target:         build.Target,
	}

	if preview {
		return name, state, nil
	}

	result, err := cfg.client.Build(ctx.(context.Context), opts)
	if err != nil {
		return name, state, err
	}

	var id string
	if digest, ok := result.ExporterResponse[exptypes.ExporterImageConfigDigestKey]; ok {
		id = digest
	} else if tag, ok := result.ExporterResponse["image.name"]; ok {
		id = tag
	} else {
		id = name
	}

	// TODO: Handle case with no export.
	_, _, state, err = i.Read(ctx, id, input, state)

	return id, state, err
}

func (*Image) Read(
	ctx provider.Context,
	id string,
	input ImageArgs,
	state ImageState,
) (
	string, // id
	ImageArgs, // normalized inputs
	ImageState, // normalized state
	error,
) {
	cfg := infer.GetConfig[Config](ctx)
	inspect, err := cfg.client.Inspect(ctx, id)
	if err != nil {
		err = fmt.Errorf("inspecting image: %w", err)
	}
	// TODO: Handle cases where a tag has been deleted.

	state.File = input.File

	state.Architecture = inspect.Architecture
	state.OS = inspect.Os
	state.RepoDigests = inspect.RepoDigests
	state.RepoTags = inspect.RepoTags
	state.Size = inspect.Size

	slices.Sort(state.RepoDigests)
	slices.Sort(state.RepoTags)

	return inspect.ID, input, state, err
}

// Delete deletes an Image. If the Image was already deleted out-of-band it is treated as a success.
//
// Any tags previously pushed to registries will not be deleted.
func (*Image) Delete(
	ctx provider.Context,
	id string,
	_ ImageState,
) error {
	cfg := infer.GetConfig[Config](ctx)

	deletions, err := cfg.client.Delete(ctx.(context.Context), id)
	if errdefs.IsNotFound(err) {
		return nil // Nothing to do.
	}

	for _, d := range deletions {
		if d.Deleted != "" {
			ctx.Log(diag.Info, d.Deleted)
		}
		if d.Untagged != "" {
			ctx.Log(diag.Info, d.Untagged)
		}
	}

	// TODO: Delete tags from registries?

	return err
}
