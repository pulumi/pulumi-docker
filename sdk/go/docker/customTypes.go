// Copyright 2016-2020, Pulumi Corporation.

package docker

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// CacheFrom is a copy of CacheFromArgs but without using TInput in types.
type CacheFrom struct {
	Stages []string `pulumi:"stages"`
}

// CacheFromInput is an input type that accepts CacheFromArgs and CacheFromOutput values.
// You can construct a concrete instance of `CacheFromInput` via:
//
//          CacheFromArgs{...}
type CacheFromInput interface {
	pulumi.Input

	ToCacheFromOutput() CacheFromOutput
	ToCacheFromOutputWithContext(context.Context) CacheFromOutput
}

// CacheFromPtrInput is an input type that accepts CacheFromArgs, CacheFromPtr and CacheFromPtrOutput values.
// You can construct a concrete instance of `CacheFromPtrInput` via:
//
//          CacheFromArgs{...}
//
//  or:
//
//          nil
type CacheFromPtrInput interface {
	pulumi.Input

	ToCacheFromPtrOutput() CacheFromPtrOutput
	ToCacheFromPtrOutputWithContext(context.Context) CacheFromPtrOutput
}

type cacheFromPtrType CacheFromArgs

func CacheFromPtr(v *CacheFromArgs) CacheFromPtrInput {
	return (*cacheFromPtrType)(v)
}

func (*cacheFromPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheFrom)(nil)).Elem()
}

func (i *cacheFromPtrType) ToCacheFromPtrOutput() CacheFromPtrOutput {
	return i.ToCacheFromPtrOutputWithContext(context.Background())
}

func (i *cacheFromPtrType) ToCacheFromPtrOutputWithContext(ctx context.Context) CacheFromPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheFromPtrOutput)
}

type CacheFromPtrOutput struct{ *pulumi.OutputState }

func (CacheFromPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheFrom)(nil)).Elem()
}

func (o CacheFromPtrOutput) ToCacheFromPtrOutput() CacheFromPtrOutput {
	return o
}

func (o CacheFromPtrOutput) ToCacheFromPtrOutputWithContext(ctx context.Context) CacheFromPtrOutput {
	return o
}

func (o CacheFromPtrOutput) Elem() CacheFromOutput {
	return o.ApplyT(func(v *CacheFrom) CacheFrom { return *v }).(CacheFromOutput)
}

// CacheFromArgs may be used to specify build stages to use for the Docker build cache.
// The final image is always implicitly included.
type CacheFromArgs struct {
	Stages pulumi.StringArrayInput `pulumi:"stages"`
}

func (CacheFromArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheFrom)(nil)).Elem()
}

func (i CacheFromArgs) ToCacheFromOutput() CacheFromOutput {
	return i.ToCacheFromOutputWithContext(context.Background())
}

func (i CacheFromArgs) ToCacheFromOutputWithContext(ctx context.Context) CacheFromOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheFromOutput)
}

func (i CacheFromArgs) ToCacheFromPtrOutput() CacheFromPtrOutput {
	return i.ToCacheFromPtrOutputWithContext(context.Background())
}

func (i CacheFromArgs) ToCacheFromPtrOutputWithContext(ctx context.Context) CacheFromPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheFromOutput).ToCacheFromPtrOutputWithContext(ctx)
}

type CacheFromOutput struct{ *pulumi.OutputState }

func (CacheFromOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CacheFrom)(nil)).Elem()
}

func (o CacheFromOutput) ToCacheFromOutput() CacheFromOutput {
	return o
}

func (o CacheFromOutput) ToCacheFromOutputWithContext(ctx context.Context) CacheFromOutput {
	return o
}

func (o CacheFromOutput) ToCacheFromPtrOutput() CacheFromPtrOutput {
	return o.ToCacheFromPtrOutputWithContext(context.Background())
}

func (o CacheFromOutput) ToCacheFromPtrOutputWithContext(ctx context.Context) CacheFromPtrOutput {
	return o.ApplyT(func(v CacheFrom) *CacheFrom {
		return &v
	}).(CacheFromPtrOutput)
}

// DockerBuild is a copy of DockerBuildArgs but without using TInput in types.
// nolint:golint
type DockerBuild struct {
	Context      string            `pulumi:"context"`
	Dockerfile   string            `pulumi:"dockerfile"`
	Args         map[string]string `pulumi:"args"`
	CacheFrom    *CacheFrom        `pulumi:"cacheFrom"`
	ExtraOptions []string          `pulumi:"extraOptions"`
	Env          map[string]string `pulumi:"env"`
	Target       string            `pulumi:"target"`
}

// DockerBuildArgs may be used to specify detailed instructions about how to build a container.
// nolint:golint
type DockerBuildArgs struct {
	// Context is a path to a directory to use for the Docker build context, usually the directory
	// in which the Dockerfile resides (although dockerfile may be used to choose a custom location
	// independent of this choice). If not specified, the context defaults to the current working
	// directory; if a relative path is used, it is relative to the current working directory that
	// Pulumi is evaluating.
	Context pulumi.StringInput `pulumi:"context"`

	// Dockerfile may be used to override the default Dockerfile name and/or location.
	// By default, it is assumed to be a file named Dockerfile in the root of the build context.
	Dockerfile pulumi.StringInput `pulumi:"dockerfile"`

	// An optional map of named build-time argument variables to set during the Docker build.
	// This flag allows you to pass built-time variables that can be accessed like environment variables
	// inside the `RUN` instruction.
	Args pulumi.StringMap `pulumi:"args"`

	// An optional CacheFrom object with information about the build stages to use for the Docker
	// build cache.  This parameter maps to the --cache-from argument to the Docker CLI. If this
	// parameter is `true`, only the final image will be pulled and passed to --cache-from; if it is
	// a CacheFrom object, the stages named therein will also be pulled and passed to --cache-from.
	CacheFrom CacheFromPtrInput `pulumi:"cacheFrom"`

	// An optional catch-all list of arguments to provide extra CLI options to the docker build command.
	// For example `{'--network', 'host'}`.
	ExtraOptions pulumi.StringArrayInput `pulumi:"extraOptions"`

	// Environment variables to set on the invocation of `docker build`, for example to support
	// `DOCKER_BUILDKIT=1 docker build`.
	Env pulumi.StringMapInput `pulumi:"env"`

	// The target of the dockerfile to build.
	Target pulumi.StringInput `pulumi:"target"`
}

// nolint:golint
type DockerBuildInput interface {
	pulumi.Input

	ToDockerBuildOutput() DockerBuildOutput
	ToDockerBuildOutputWithContext(context.Context) DockerBuildOutput
}

func (DockerBuildArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DockerBuild)(nil)).Elem()
}

func (i DockerBuildArgs) ToDockerBuildOutput() DockerBuildOutput {
	return i.ToDockerBuildOutputWithContext(context.Background())
}

func (i DockerBuildArgs) ToDockerBuildOutputWithContext(ctx context.Context) DockerBuildOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DockerBuildOutput)
}

// nolint:golint
type DockerBuildOutput struct{ *pulumi.OutputState }

func (DockerBuildOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DockerBuild)(nil)).Elem()
}

func (o DockerBuildOutput) ToDockerBuildOutput() DockerBuildOutput {
	return o
}

func (o DockerBuildOutput) ToDockerBuildOutputWithContext(ctx context.Context) DockerBuildOutput {
	return o
}

// ImageRegistryArgs contains credentials for the docker registry.
type ImageRegistryArgs struct {
	// Docker registry server URL to push to.  Some common values include:
	// DockerHub: `docker.io` or `https://index.docker.io/v1`
	// Azure Container Registry: `<name>.azurecr.io`
	// AWS Elastic Container Registry: `<account>.dkr.ecr.us-east-2.amazonaws.com`
	// Google Container Registry: `<name>.gcr.io`
	Server pulumi.StringInput `pulumi:"server"`

	// Username for login to the target Docker registry.
	Username pulumi.StringInput `pulumi:"username"`

	// Password for login to the target Docker registry.
	Password pulumi.StringInput `pulumi:"password"`
}

// ImageRegistry is a copy of ImageRegistryArgs but without using TInput in types.
type ImageRegistry struct {
	Server   string
	Username string
	Password string
}

type ImageRegistryInput interface {
	pulumi.Input

	ToImageRegistryOutput() ImageRegistryOutput
	ToImageRegistryOutputWithContext(context.Context) ImageRegistryOutput
}

func (ImageRegistryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageRegistry)(nil)).Elem()
}

func (i ImageRegistryArgs) ToImageRegistryOutput() ImageRegistryOutput {
	return i.ToImageRegistryOutputWithContext(context.Background())
}

func (i ImageRegistryArgs) ToImageRegistryOutputWithContext(ctx context.Context) ImageRegistryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageRegistryOutput)
}

type ImageRegistryOutput struct{ *pulumi.OutputState }

func (ImageRegistryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageRegistry)(nil)).Elem()
}

func (o ImageRegistryOutput) ToImageRegistryOutput() ImageRegistryOutput {
	return o
}

func (o ImageRegistryOutput) ToImageRegistryOutputWithContext(ctx context.Context) ImageRegistryOutput {
	return o
}

func (o ImageRegistryOutput) Server() pulumi.StringOutput {
	return o.ApplyT(func(v ImageRegistry) string { return v.Server }).(pulumi.StringOutput)
}

// ImageArgs are the arguments are constructing an Image resource.
type ImageArgs struct {

	// The qualified image name that will be pushed to the remote registry.  Must be a supported
	// image name for the target registry user.  This name can include a tag at the end.  If
	// provided all pushed image resources will contain that tag as well.
	//
	// Either [imageName] or [localImageName] can have a tag.  However, if both have a tag, then
	// those tags must match.
	ImageName pulumi.StringInput

	// The Docker build context, as a folder path or a detailed DockerBuild object.
	Build DockerBuildInput

	// The docker image name to build locally before tagging with imageName.  If not provided, it
	// will be given the value of to [imageName].  This name can include a tag at the end.  If
	// provided all pushed image resources will contain that tag as well.
	//
	// Either [imageName] or [localImageName] can have a tag.  However, if both have a tag, then
	// those tags must match.
	LocalImageName pulumi.StringInput

	// Credentials for the docker registry to push to.
	Registry ImageRegistryInput

	// Skip push flag.
	SkipPush pulumi.BoolInput
}

// imageArgs is a copy of ImageArgs but without using TInput in types.
type imageArgs struct {
	ImageName      string        `pulumi:"imageName"`
	Build          DockerBuild   `pulumi:"build"`
	LocalImageName string        `pulumi:"localImageName"`
	Registry       ImageRegistry `pulumi:"registry"`
	SkipPush       bool          `pulumi:"skipPush"`
}

type ImageArgsInput interface {
	pulumi.Input

	ToImageArgsOutput() ImageArgsOutput
	ToImageArgsOutputWithContext(context.Context) ImageArgsOutput
}

func (ImageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*imageArgs)(nil)).Elem()
}

func (i ImageArgs) ToImageArgsOutput() ImageArgsOutput {
	return i.ToImageArgsOutputWithContext(context.Background())
}

func (i ImageArgs) ToImageArgsOutputWithContext(ctx context.Context) ImageArgsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageArgsOutput)
}

type ImageArgsOutput struct{ *pulumi.OutputState }

func (ImageArgsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*imageArgs)(nil)).Elem()
}

func (o ImageArgsOutput) ToImageArgsOutput() ImageArgsOutput {
	return o
}

func (o ImageArgsOutput) ToImageArgsOutputWithContext(ctx context.Context) ImageArgsOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(CacheFromOutput{})
	pulumi.RegisterOutputType(DockerBuildOutput{})
	pulumi.RegisterOutputType(ImageRegistryOutput{})
	pulumi.RegisterOutputType(ImageArgsOutput{})
}
