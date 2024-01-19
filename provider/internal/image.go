package internal

import (
	"context"
	"errors"
	"fmt"

	// These imports are needed to register the drivers with buildkit.
	_ "github.com/docker/buildx/driver/docker-container"
	_ "github.com/docker/buildx/driver/kubernetes"
	_ "github.com/docker/buildx/util/buildflags"

	"github.com/distribution/reference"
	controllerapi "github.com/docker/buildx/controller/pb"
	"github.com/docker/buildx/util/buildflags"
	"github.com/docker/buildx/util/platformutil"
	"github.com/docker/docker/errdefs"
	"github.com/moby/buildkit/exporter/containerimage/exptypes"
	"github.com/muesli/reflow/dedent"

	provider "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/diag"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"

	"github.com/pulumi/pulumi-docker/provider/v4/internal/properties"
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
	BuildArgs  map[string]string         `pulumi:"buildArgs,optional"`
	CacheFrom  []string                  `pulumi:"cacheFrom,optional"`
	CacheTo    []string                  `pulumi:"cacheTo,optional"`
	Context    string                    `pulumi:"context,optional"`
	Exports    []string                  `pulumi:"exports,optional"`
	File       string                    `pulumi:"file,optional"`
	Platforms  []string                  `pulumi:"platforms,optional"`
	Pull       bool                      `pulumi:"pull,optional"`
	Registries []properties.RegistryAuth `pulumi:"registries,optional"`
	Tags       []string                  `pulumi:"tags"`
}

// Annotate describes inputs to the Image resource.
func (ia *ImageArgs) Annotate(a infer.Annotator) {
	a.Describe(&ia.BuildArgs, dedent.String(`
		An optional map of named build-time argument variables to set during
		the Docker build. This flag allows you to pass build-time variables that
		can be accessed like environment variables inside the RUN
		instruction.`,
	))
	a.Describe(&ia.CacheFrom, dedent.String(`
		External cache sources (e.g., "user/app:cache", "type=local,src=path/to/dir")`,
	))
	a.Describe(&ia.CacheTo, dedent.String(`
		Cache export destinations (e.g., "user/app:cache", "type=local,dest=path/to/dir")`,
	))
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
	a.Describe(&ia.Platforms, dedent.String(`
		Set target platforms for the build. Defaults to the host's platform`,
	))
	a.Describe(&ia.Pull, dedent.String(`
		Always attempt to pull all referenced images`,
	))
	a.Describe(&ia.Tags, dedent.String(`
		Name and optionally a tag (format: "name:tag"). If outputting to a
		registry, the name should include the fully qualified registry address.`,
	))
	a.Describe(&ia.Registries, dedent.String(`
		Logins for registry outputs`,
	))

	a.SetDefault(&ia.File, "Dockerfile")
}

// ImageState is serialized to the program's state file.
type ImageState struct {
	ImageArgs

	Manifests []properties.Manifest `pulumi:"manifests" provider:"output"`
}

// Annotate describes outputs of the Image resource.
func (is *ImageState) Annotate(a infer.Annotator) {
	is.ImageArgs.Annotate(a)
}

// Check validates ImageArgs, sets defaults, and ensures our client is
// authenticated.
func (*Image) Check(
	ctx provider.Context,
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

	if _, berr := args.toBuildOptions(); berr != nil {
		errs := berr.(interface{ Unwrap() []error }).Unwrap()
		for _, e := range errs {
			if cf, ok := e.(checkFailure); ok {
				failures = append(failures, cf.CheckFailure)
			}
		}
	}

	// Check is called before every operation except Read, so this ensures
	// we're authenticated in almost all cases.
	cfg := infer.GetConfig[Config](ctx)
	for _, reg := range args.Registries {
		if err = cfg.client.Auth(ctx, reg); err != nil {
			failures = append(failures,
				provider.CheckFailure{Property: "registries", Reason: fmt.Sprintf("unable to authenticate: %s", err.Error())})
		}
	}

	return args, failures, err
}

type checkFailure struct {
	provider.CheckFailure
}

func (cf checkFailure) Error() string {
	return cf.Reason
}

func newCheckFailure(property string, err error) checkFailure {
	return checkFailure{provider.CheckFailure{Property: property, Reason: err.Error()}}
}

func (ia *ImageArgs) toBuildOptions() (controllerapi.BuildOptions, error) {
	var multierr error
	exports, err := buildflags.ParseExports(ia.Exports)
	if err != nil {
		multierr = errors.Join(multierr, newCheckFailure("exports", err))
	}

	_, err = platformutil.Parse(ia.Platforms)
	if err != nil {
		multierr = errors.Join(multierr, newCheckFailure("platforms", err))
	}

	cacheFrom, err := buildflags.ParseCacheEntry(ia.CacheFrom)
	if err != nil {
		multierr = errors.Join(multierr, newCheckFailure("cacheFrom", err))
	}

	cacheTo, err := buildflags.ParseCacheEntry(ia.CacheTo)
	if err != nil {
		multierr = errors.Join(multierr, newCheckFailure("cacheTo", err))
	}

	for _, t := range ia.Tags {
		if _, err := reference.Parse(t); err != nil {
			multierr = errors.Join(multierr, newCheckFailure("tags", err))
		}
	}

	opts := controllerapi.BuildOptions{
		BuildArgs:      ia.BuildArgs,
		CacheFrom:      cacheFrom,
		CacheTo:        cacheTo,
		ContextPath:    ia.Context,
		DockerfileName: ia.File,
		Exports:        exports,
		Platforms:      ia.Platforms,
		Pull:           ia.Pull,
		Tags:           ia.Tags,
	}

	return opts, multierr
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

	opts, err := input.toBuildOptions()
	if err != nil {
		return name, state, fmt.Errorf("validating input: %w", err)
	}

	if preview {
		return name, state, nil
	}

	result, err := cfg.client.Build(ctx.(context.Context), opts)
	if err != nil {
		return name, state, fmt.Errorf("building %q: %w", input.Tags, err)
	}

	var id string
	if digest, ok := result.ExporterResponse["containerimage.digest"]; ok {
		id = digest
	} else if digest, ok := result.ExporterResponse[exptypes.ExporterImageConfigDigestKey]; ok {
		id = digest
	} else if tags, ok := result.ExporterResponse["image.name"]; ok {
		id = tags
	} else {
		id = name
	}

	// TODO: Handle case with no export.
	_, _, state, err = i.Read(ctx, id, input, state)

	return id, state, err
}

// Read attempts to read manifests from an image's exports. An image without
// exports will have no manifests.
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
	opts, err := input.toBuildOptions()
	if err != nil {
		return id, input, state, err
	}

	// Ensure we're authenticated.
	cfg := infer.GetConfig[Config](ctx)
	for _, reg := range input.Registries {
		if err = cfg.client.Auth(ctx, reg); err != nil {
			return id, input, state, err
		}
	}

	manifests := []properties.Manifest{}
	for _, export := range opts.Exports {
		switch export.GetType() {
		case "image":
			if export.GetAttrs()["push"] != "true" {
				// No manifest to read if we didn't push.
				continue
			}
			for _, tag := range input.Tags {
				infos, err := cfg.client.Inspect(ctx, tag)
				if err != nil {
					continue
				}
				for _, m := range infos {
					if m.Descriptor.Platform != nil && m.Descriptor.Platform.Architecture == "unknown" {
						// Ignore cache manifests.
						continue
					}
					if m.Ref == nil {
						// Shouldn't happen, but just in case.
						continue
					}
					manifests = append(manifests, properties.Manifest{
						Digest: m.Descriptor.Digest.String(),
						Platform: properties.Platform{
							OS:           m.Descriptor.Platform.OS,
							Architecture: m.Descriptor.Platform.Architecture,
						},
						Ref:  m.Ref.String(),
						Size: m.Descriptor.Size,
						URLs: m.Descriptor.URLs,
					})
				}
			}
		case "docker":
			//
		default:
			// Other export types (e.g. file) are not supported.
			continue
		}
	}

	state.Manifests = manifests

	return id, input, state, nil
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
