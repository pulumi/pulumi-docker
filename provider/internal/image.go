package internal

import (
	// Docker providers are registered in init()
	_ "github.com/docker/buildx/driver/docker"
	_ "github.com/docker/buildx/driver/docker-container"
	_ "github.com/docker/buildx/driver/kubernetes"
	_ "github.com/docker/buildx/driver/remote"

	"github.com/muesli/reflow/dedent"

	provider "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

var (
	_ infer.CustomResource[ImageArgs, ImageState] = (*Image)(nil)
	_ infer.CustomCheck[ImageArgs]                = (*Image)(nil)
	_ infer.Annotated                             = (infer.Annotated)((*Image)(nil))
	_ infer.Annotated                             = (infer.Annotated)((*ImageArgs)(nil))
	_ infer.Annotated                             = (infer.Annotated)((*ImageState)(nil))
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

	if args.File == "" {
		args.File = "Dockerfile"
	}

	return args, failures, err
}

// Create builds the image using buildkit.
func (*Image) Create(
	_ provider.Context,
	name string,
	input ImageArgs,
	_ bool,
) (string, ImageState, error) {
	state := ImageState{}

	state.Tags = input.Tags
	state.File = input.File
	state.Context = input.Context

	// TODO(https://github.com/pulumi/pulumi-docker/issues/885): Actually perform the build.

	return name, state, nil
}
