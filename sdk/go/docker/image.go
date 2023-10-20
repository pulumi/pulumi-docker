// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package docker

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-docker/sdk/v4/go/docker/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// `Image` builds a Docker image and pushes it Docker and OCI compatible registries.
// This resource enables running Docker builds as part of a Pulumi deployment.
//
// Note: This resource does not delete tags, locally or remotely, when destroyed.
//
// ## Image name
//
// The Image resource uses `imageName` to refer to a fully qualified Docker image name, by the format `repository:tag`.
// Note that this does not include any digest information and thus will not cause any updates when passed to dependencies,
// even when using `latest` tag. To trigger such updates, e.g. when referencing pushed images in container orchestration
// and management resources, please use the `repoDigest` Output instead, which is of the format
// `repository@<algorithm>:<hash>` and unique per build/push.
// Note that `repoDigest` is not available for local Images. For a local Image not pushed to a registry, you may want to
// give `imageName` a unique tag per pulumi update.
//
// ## Cross-platform builds
//
// The Image resource supports cross-platform builds when the [Docker engine has cross-platform support enabled via emulators](https://docs.docker.com/build/building/multi-platform/#building-multi-platform-images).
// The Image resource currently supports providing only a single operating system and architecture in the `platform` field, e.g.: `linux/amd64`.
// To enable this support, you may need to install the emulators in the environment running your Pulumi program.
//
// If you are using Linux, you may be using Docker Engine or Docker Desktop for Linux, depending on how you have installed Docker. The [FAQ for Docker Desktop for Linux](https://docs.docker.com/desktop/faqs/linuxfaqs/#context) describes the differences and how to select which Docker context is in use.
//
// * For local development using Docker Desktop, this is enabled by default.
// * For systems using Docker Engine, install the QEMU binaries and register them with using the docker image from [github.com/tonistiigi/binfmt](https://github.com/tonistiigi/binfmt):
// * In a GitHub Actions workflow, the [docker/setup-qemu-action](https://github.com/docker/setup-qemu-action) can be used instead by adding this step to your workflow file. Example workflow usage:
//
// ## Example Usage
// ### A Docker image build
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-docker/sdk/v4/go/docker"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			demoImage, err := docker.NewImage(ctx, "demo-image", &docker.ImageArgs{
//				Build: &docker.DockerBuildArgs{
//					Args: pulumi.StringMap{
//						"platform": pulumi.String("linux/amd64"),
//					},
//					Context:    pulumi.String("."),
//					Dockerfile: pulumi.String("Dockerfile"),
//				},
//				ImageName: pulumi.String("username/image:tag1"),
//				SkipPush:  pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			ctx.Export("imageName", demoImage.ImageName)
//			return nil
//		})
//	}
//
// ```
// ### A Docker image build and push
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-docker/sdk/v4/go/docker"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			demoPushImage, err := docker.NewImage(ctx, "demo-push-image", &docker.ImageArgs{
//				Build: &docker.DockerBuildArgs{
//					Context:    pulumi.String("."),
//					Dockerfile: pulumi.String("Dockerfile"),
//				},
//				ImageName: pulumi.String("docker.io/username/push-image:tag1"),
//			})
//			if err != nil {
//				return err
//			}
//			ctx.Export("imageName", demoPushImage.ImageName)
//			ctx.Export("repoDigest", demoPushImage.RepoDigest)
//			return nil
//		})
//	}
//
// ```
// ### Docker image build using caching with AWS Elastic Container Registry
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ecr"
//	"github.com/pulumi/pulumi-docker/sdk/v4/go/docker"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			ecrRepository, err := ecr.NewRepository(ctx, "ecr-repository", &ecr.RepositoryArgs{
//				Name: pulumi.String("docker-repository"),
//			})
//			if err != nil {
//				return err
//			}
//			authToken := ecr.GetAuthorizationTokenOutput(ctx, ecr.GetAuthorizationTokenOutputArgs{
//				RegistryId: ecrRepository.RegistryId,
//			}, nil)
//			myAppImage, err := docker.NewImage(ctx, "my-app-image", &docker.ImageArgs{
//				Build: &docker.DockerBuildArgs{
//					Args: pulumi.StringMap{
//						"BUILDKIT_INLINE_CACHE": pulumi.String("1"),
//					},
//					CacheFrom: &docker.CacheFromArgs{
//						Images: pulumi.StringArray{
//							ecrRepository.RepositoryUrl.ApplyT(func(repositoryUrl string) (string, error) {
//								return fmt.Sprintf("%v:latest", repositoryUrl), nil
//							}).(pulumi.StringOutput),
//						},
//					},
//					Context:    pulumi.String("app/"),
//					Dockerfile: pulumi.String("Dockerfile"),
//				},
//				ImageName: ecrRepository.RepositoryUrl.ApplyT(func(repositoryUrl string) (string, error) {
//					return fmt.Sprintf("%v:latest", repositoryUrl), nil
//				}).(pulumi.StringOutput),
//				Registry: &docker.RegistryArgs{
//					Password: pulumi.ToSecret(authToken.ApplyT(func(authToken ecr.GetAuthorizationTokenResult) (*string, error) {
//						return &authToken.Password, nil
//					}).(pulumi.StringPtrOutput)).(*pulumi.StringOutput),
//					Server: ecrRepository.RepositoryUrl,
//				},
//			})
//			if err != nil {
//				return err
//			}
//			ctx.Export("imageName", myAppImage.ImageName)
//			return nil
//		})
//	}
//
// ```
type Image struct {
	pulumi.CustomResourceState

	// The fully qualified image name that was pushed to the registry.
	BaseImageName pulumi.StringOutput `pulumi:"baseImageName"`
	// The path to the build context to use.
	Context pulumi.StringOutput `pulumi:"context"`
	// The location of the Dockerfile relative to the docker build context.
	Dockerfile pulumi.StringOutput `pulumi:"dockerfile"`
	// The fully qualified image name
	ImageName pulumi.StringOutput `pulumi:"imageName"`
	// The name of the registry server hosting the image.
	RegistryServer pulumi.StringOutput `pulumi:"registryServer"`
	// **For pushed images:**
	// The manifest digest of an image pushed to a registry, of the format repository@<algorithm>:<hash>, e.g. `username/demo-image@sha256:a6ae6dd8d39c5bb02320e41abf00cd4cb35905fec540e37d306c878be8d38bd3`.
	// This reference is unique per image build and push.
	// Only available for images pushed to a registry.
	// Use when passing a reference to a pushed image to container management resources.
	//
	// **Local-only images**For local images, this field is the image ID of the built local image, of the format <algorithm>:<hash>, e.g `sha256:826a130323165bb0ccb0374ae774f885c067a951b51a6ee133577f4e5dbc4119`
	RepoDigest pulumi.StringOutput `pulumi:"repoDigest"`
}

// NewImage registers a new resource with the given unique name, arguments, and options.
func NewImage(ctx *pulumi.Context,
	name string, args *ImageArgs, opts ...pulumi.ResourceOption) (*Image, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ImageName == nil {
		return nil, errors.New("invalid value for required argument 'ImageName'")
	}
	if args.SkipPush == nil {
		args.SkipPush = pulumi.BoolPtr(false)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("docker:image:Image"),
		},
	})
	opts = append(opts, aliases)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Image
	err := ctx.RegisterResource("docker:index/image:Image", name, args, &resource, opts...)
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
	err := ctx.ReadResource("docker:index/image:Image", name, id, state, &resource, opts...)
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
	// The Docker build context
	Build *DockerBuild `pulumi:"build"`
	// The image name, of the format repository[:tag], e.g. `docker.io/username/demo-image:v1`.
	// This reference is not unique to each build and push.For the unique manifest SHA of a pushed docker image, or the local image ID, please use `repoDigest`.
	ImageName string `pulumi:"imageName"`
	// The registry to push the image to
	Registry *Registry `pulumi:"registry"`
	// A flag to skip a registry push.
	SkipPush *bool `pulumi:"skipPush"`
}

// The set of arguments for constructing a Image resource.
type ImageArgs struct {
	// The Docker build context
	Build DockerBuildPtrInput
	// The image name, of the format repository[:tag], e.g. `docker.io/username/demo-image:v1`.
	// This reference is not unique to each build and push.For the unique manifest SHA of a pushed docker image, or the local image ID, please use `repoDigest`.
	ImageName pulumi.StringInput
	// The registry to push the image to
	Registry RegistryPtrInput
	// A flag to skip a registry push.
	SkipPush pulumi.BoolPtrInput
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

func (i *Image) ToOutput(ctx context.Context) pulumix.Output[*Image] {
	return pulumix.Output[*Image]{
		OutputState: i.ToImageOutputWithContext(ctx).OutputState,
	}
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

func (i ImageArray) ToOutput(ctx context.Context) pulumix.Output[[]*Image] {
	return pulumix.Output[[]*Image]{
		OutputState: i.ToImageArrayOutputWithContext(ctx).OutputState,
	}
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

func (i ImageMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Image] {
	return pulumix.Output[map[string]*Image]{
		OutputState: i.ToImageMapOutputWithContext(ctx).OutputState,
	}
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

func (o ImageOutput) ToOutput(ctx context.Context) pulumix.Output[*Image] {
	return pulumix.Output[*Image]{
		OutputState: o.OutputState,
	}
}

// The fully qualified image name that was pushed to the registry.
func (o ImageOutput) BaseImageName() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.BaseImageName }).(pulumi.StringOutput)
}

// The path to the build context to use.
func (o ImageOutput) Context() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.Context }).(pulumi.StringOutput)
}

// The location of the Dockerfile relative to the docker build context.
func (o ImageOutput) Dockerfile() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.Dockerfile }).(pulumi.StringOutput)
}

// The fully qualified image name
func (o ImageOutput) ImageName() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.ImageName }).(pulumi.StringOutput)
}

// The name of the registry server hosting the image.
func (o ImageOutput) RegistryServer() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.RegistryServer }).(pulumi.StringOutput)
}

// **For pushed images:**
// The manifest digest of an image pushed to a registry, of the format repository@<algorithm>:<hash>, e.g. `username/demo-image@sha256:a6ae6dd8d39c5bb02320e41abf00cd4cb35905fec540e37d306c878be8d38bd3`.
// This reference is unique per image build and push.
// Only available for images pushed to a registry.
// Use when passing a reference to a pushed image to container management resources.
//
// **Local-only images**For local images, this field is the image ID of the built local image, of the format <algorithm>:<hash>, e.g `sha256:826a130323165bb0ccb0374ae774f885c067a951b51a6ee133577f4e5dbc4119`
func (o ImageOutput) RepoDigest() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.RepoDigest }).(pulumi.StringOutput)
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

func (o ImageArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Image] {
	return pulumix.Output[[]*Image]{
		OutputState: o.OutputState,
	}
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

func (o ImageMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Image] {
	return pulumix.Output[map[string]*Image]{
		OutputState: o.OutputState,
	}
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
