package internal

import (
	provider "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

// Image is a Docker image build using buildkit.
type Image struct{}

// ImageArgs instantiates a new Image.
type ImageArgs struct {
	Tags    []string `pulumi:"tags"`
	File    string   `pulumi:"file,optional"`
	Context []string `pulumi:"context,optional"`
}

// ImageState is serialized to the program's state file.
type ImageState struct {
	ImageArgs
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
) (string, *ImageState, error) {
	state := &ImageState{}

	state.Tags = input.Tags
	state.File = input.File
	state.Context = input.Context

	// TODO(https://github.com/pulumi/pulumi-docker/issues/885): Actually perform the build.

	return name, state, nil
}
