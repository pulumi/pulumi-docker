// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package buildx

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-docker/sdk/v4/go/docker/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A Docker image built using Buildkit
type Image struct {
	pulumi.CustomResourceState

	// An optional map of named build-time argument variables to set during
	// the Docker build. This flag allows you to pass build-time variables that
	// can be accessed like environment variables inside the RUN
	// instruction.
	BuildArgs pulumi.StringMapOutput `pulumi:"buildArgs"`
	// External cache sources (e.g., "user/app:cache", "type=local,src=path/to/dir")
	CacheFrom pulumi.StringArrayOutput `pulumi:"cacheFrom"`
	// Cache export destinations (e.g., "user/app:cache", "type=local,dest=path/to/dir")
	CacheTo pulumi.StringArrayOutput `pulumi:"cacheTo"`
	// Contexts to use while building the image. If omitted, an empty context
	// is used. If more than one value is specified, they should be of the
	// form "name=value".
	Context pulumi.StringPtrOutput `pulumi:"context"`
	// Name and optionally a tag (format: "name:tag"). If outputting to a
	// registry, the name should include the fully qualified registry address.
	Exports pulumi.StringArrayOutput `pulumi:"exports"`
	// Name of the Dockerfile to use (default: "$PATH/Dockerfile").
	File      pulumi.StringPtrOutput `pulumi:"file"`
	Manifests ManifestArrayOutput    `pulumi:"manifests"`
	// Set target platforms for the build. Defaults to the host's platform
	Platforms pulumi.StringArrayOutput `pulumi:"platforms"`
	// Always attempt to pull all referenced images
	Pull pulumi.BoolPtrOutput `pulumi:"pull"`
	// Name and optionally a tag (format: "name:tag"). If outputting to a
	// registry, the name should include the fully qualified registry address.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
}

// NewImage registers a new resource with the given unique name, arguments, and options.
func NewImage(ctx *pulumi.Context,
	name string, args *ImageArgs, opts ...pulumi.ResourceOption) (*Image, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Tags == nil {
		return nil, errors.New("invalid value for required argument 'Tags'")
	}
	if args.File == nil {
		args.File = pulumi.StringPtr("Dockerfile")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Image
	err := ctx.RegisterResource("docker:buildx/image:Image", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetImage gets an existing Image resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetImage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ImageState, opts ...pulumi.ResourceOption) (*Image, error) {
	var resource Image
	err := ctx.ReadResource("docker:buildx/image:Image", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Image resources.
type imageState struct {
}

type ImageState struct {
}

func (ImageState) ElementType() reflect.Type {
	return reflect.TypeOf((*imageState)(nil)).Elem()
}

type imageArgs struct {
	// An optional map of named build-time argument variables to set during
	// the Docker build. This flag allows you to pass build-time variables that
	// can be accessed like environment variables inside the RUN
	// instruction.
	BuildArgs map[string]string `pulumi:"buildArgs"`
	// External cache sources (e.g., "user/app:cache", "type=local,src=path/to/dir")
	CacheFrom []string `pulumi:"cacheFrom"`
	// Cache export destinations (e.g., "user/app:cache", "type=local,dest=path/to/dir")
	CacheTo []string `pulumi:"cacheTo"`
	// Contexts to use while building the image. If omitted, an empty context
	// is used. If more than one value is specified, they should be of the
	// form "name=value".
	Context *string `pulumi:"context"`
	// Name and optionally a tag (format: "name:tag"). If outputting to a
	// registry, the name should include the fully qualified registry address.
	Exports []string `pulumi:"exports"`
	// Name of the Dockerfile to use (default: "$PATH/Dockerfile").
	File *string `pulumi:"file"`
	// Set target platforms for the build. Defaults to the host's platform
	Platforms []string `pulumi:"platforms"`
	// Always attempt to pull all referenced images
	Pull *bool `pulumi:"pull"`
	// Name and optionally a tag (format: "name:tag"). If outputting to a
	// registry, the name should include the fully qualified registry address.
	Tags []string `pulumi:"tags"`
}

// The set of arguments for constructing a Image resource.
type ImageArgs struct {
	// An optional map of named build-time argument variables to set during
	// the Docker build. This flag allows you to pass build-time variables that
	// can be accessed like environment variables inside the RUN
	// instruction.
	BuildArgs pulumi.StringMapInput
	// External cache sources (e.g., "user/app:cache", "type=local,src=path/to/dir")
	CacheFrom pulumi.StringArrayInput
	// Cache export destinations (e.g., "user/app:cache", "type=local,dest=path/to/dir")
	CacheTo pulumi.StringArrayInput
	// Contexts to use while building the image. If omitted, an empty context
	// is used. If more than one value is specified, they should be of the
	// form "name=value".
	Context pulumi.StringPtrInput
	// Name and optionally a tag (format: "name:tag"). If outputting to a
	// registry, the name should include the fully qualified registry address.
	Exports pulumi.StringArrayInput
	// Name of the Dockerfile to use (default: "$PATH/Dockerfile").
	File pulumi.StringPtrInput
	// Set target platforms for the build. Defaults to the host's platform
	Platforms pulumi.StringArrayInput
	// Always attempt to pull all referenced images
	Pull pulumi.BoolPtrInput
	// Name and optionally a tag (format: "name:tag"). If outputting to a
	// registry, the name should include the fully qualified registry address.
	Tags pulumi.StringArrayInput
}

func (ImageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*imageArgs)(nil)).Elem()
}

type ImageInput interface {
	pulumi.Input

	ToImageOutput() ImageOutput
	ToImageOutputWithContext(ctx context.Context) ImageOutput
}

func (*Image) ElementType() reflect.Type {
	return reflect.TypeOf((**Image)(nil)).Elem()
}

func (i *Image) ToImageOutput() ImageOutput {
	return i.ToImageOutputWithContext(context.Background())
}

func (i *Image) ToImageOutputWithContext(ctx context.Context) ImageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageOutput)
}

// ImageArrayInput is an input type that accepts ImageArray and ImageArrayOutput values.
// You can construct a concrete instance of `ImageArrayInput` via:
//
//	ImageArray{ ImageArgs{...} }
type ImageArrayInput interface {
	pulumi.Input

	ToImageArrayOutput() ImageArrayOutput
	ToImageArrayOutputWithContext(context.Context) ImageArrayOutput
}

type ImageArray []ImageInput

func (ImageArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Image)(nil)).Elem()
}

func (i ImageArray) ToImageArrayOutput() ImageArrayOutput {
	return i.ToImageArrayOutputWithContext(context.Background())
}

func (i ImageArray) ToImageArrayOutputWithContext(ctx context.Context) ImageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageArrayOutput)
}

// ImageMapInput is an input type that accepts ImageMap and ImageMapOutput values.
// You can construct a concrete instance of `ImageMapInput` via:
//
//	ImageMap{ "key": ImageArgs{...} }
type ImageMapInput interface {
	pulumi.Input

	ToImageMapOutput() ImageMapOutput
	ToImageMapOutputWithContext(context.Context) ImageMapOutput
}

type ImageMap map[string]ImageInput

func (ImageMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Image)(nil)).Elem()
}

func (i ImageMap) ToImageMapOutput() ImageMapOutput {
	return i.ToImageMapOutputWithContext(context.Background())
}

func (i ImageMap) ToImageMapOutputWithContext(ctx context.Context) ImageMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageMapOutput)
}

type ImageOutput struct{ *pulumi.OutputState }

func (ImageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Image)(nil)).Elem()
}

func (o ImageOutput) ToImageOutput() ImageOutput {
	return o
}

func (o ImageOutput) ToImageOutputWithContext(ctx context.Context) ImageOutput {
	return o
}

// An optional map of named build-time argument variables to set during
// the Docker build. This flag allows you to pass build-time variables that
// can be accessed like environment variables inside the RUN
// instruction.
func (o ImageOutput) BuildArgs() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Image) pulumi.StringMapOutput { return v.BuildArgs }).(pulumi.StringMapOutput)
}

// External cache sources (e.g., "user/app:cache", "type=local,src=path/to/dir")
func (o ImageOutput) CacheFrom() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Image) pulumi.StringArrayOutput { return v.CacheFrom }).(pulumi.StringArrayOutput)
}

// Cache export destinations (e.g., "user/app:cache", "type=local,dest=path/to/dir")
func (o ImageOutput) CacheTo() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Image) pulumi.StringArrayOutput { return v.CacheTo }).(pulumi.StringArrayOutput)
}

// Contexts to use while building the image. If omitted, an empty context
// is used. If more than one value is specified, they should be of the
// form "name=value".
func (o ImageOutput) Context() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Image) pulumi.StringPtrOutput { return v.Context }).(pulumi.StringPtrOutput)
}

// Name and optionally a tag (format: "name:tag"). If outputting to a
// registry, the name should include the fully qualified registry address.
func (o ImageOutput) Exports() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Image) pulumi.StringArrayOutput { return v.Exports }).(pulumi.StringArrayOutput)
}

// Name of the Dockerfile to use (default: "$PATH/Dockerfile").
func (o ImageOutput) File() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Image) pulumi.StringPtrOutput { return v.File }).(pulumi.StringPtrOutput)
}

func (o ImageOutput) Manifests() ManifestArrayOutput {
	return o.ApplyT(func(v *Image) ManifestArrayOutput { return v.Manifests }).(ManifestArrayOutput)
}

// Set target platforms for the build. Defaults to the host's platform
func (o ImageOutput) Platforms() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Image) pulumi.StringArrayOutput { return v.Platforms }).(pulumi.StringArrayOutput)
}

// Always attempt to pull all referenced images
func (o ImageOutput) Pull() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Image) pulumi.BoolPtrOutput { return v.Pull }).(pulumi.BoolPtrOutput)
}

// Name and optionally a tag (format: "name:tag"). If outputting to a
// registry, the name should include the fully qualified registry address.
func (o ImageOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Image) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

type ImageArrayOutput struct{ *pulumi.OutputState }

func (ImageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Image)(nil)).Elem()
}

func (o ImageArrayOutput) ToImageArrayOutput() ImageArrayOutput {
	return o
}

func (o ImageArrayOutput) ToImageArrayOutputWithContext(ctx context.Context) ImageArrayOutput {
	return o
}

func (o ImageArrayOutput) Index(i pulumi.IntInput) ImageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Image {
		return vs[0].([]*Image)[vs[1].(int)]
	}).(ImageOutput)
}

type ImageMapOutput struct{ *pulumi.OutputState }

func (ImageMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Image)(nil)).Elem()
}

func (o ImageMapOutput) ToImageMapOutput() ImageMapOutput {
	return o
}

func (o ImageMapOutput) ToImageMapOutputWithContext(ctx context.Context) ImageMapOutput {
	return o
}

func (o ImageMapOutput) MapIndex(k pulumi.StringInput) ImageOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Image {
		return vs[0].(map[string]*Image)[vs[1].(string)]
	}).(ImageOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ImageInput)(nil)).Elem(), &Image{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageArrayInput)(nil)).Elem(), ImageArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageMapInput)(nil)).Elem(), ImageMap{})
	pulumi.RegisterOutputType(ImageOutput{})
	pulumi.RegisterOutputType(ImageArrayOutput{})
	pulumi.RegisterOutputType(ImageMapOutput{})
}
