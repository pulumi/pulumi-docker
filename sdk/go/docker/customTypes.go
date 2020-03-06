// Copyright 2016-2020, Pulumi Corporation.

package docker

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type CacheFrom struct {
	Stages []string `pulumi:"stages`
}

type CacheFromInput interface {
	pulumi.Input

	ToCacheFromOutput() CacheFromOutput
	ToCacheFromOutputWithContext(context.Context) CacheFromOutput
}

type CacheFromArgs struct {
	Stages pulumi.StringArrayInput `pulumi:"stages`
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

type DockerBuild struct {
	Context      string            `pulumi:"context"`
	Dockerfile   string            `pulumi:"dockerfile"`
	Args         map[string]string `pulumi:"args"`
	CacheFrom    *CacheFrom        `pulumi:"cacheFrom"`
	ExtraOptions []string          `pulumi:"extraOptions"`
	Env          map[string]string `pulumi:"env"`
	Target       string            `pulumi:"target"`
}

// DockerBuild may be used to specify detailed instructions about how to build a container.
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
	Args pulumi.MapInput `pulumi:"args"`

	// An optional CacheFrom object with information about the build stages to use for the Docker
	// build cache.  This parameter maps to the --cache-from argument to the Docker CLI. If this
	// parameter is `true`, only the final image will be pulled and passed to --cache-from; if it is
	// a CacheFrom object, the stages named therein will also be pulled and passed to --cache-from.
	CacheFrom CacheFromInput `pulumi:"cacheFrom"`

	// An optional catch-all string to provide extra CLI options to the docker build command.
	// For example, use to specify `--network host`.
	ExtraOptions pulumi.StringArrayInput `pulumi:"extraOptions"`

	// Environment variables to set on the invocation of `docker build`, for example to support
	// `DOCKER_BUILDKIT=1 docker build`.
	Env pulumi.MapInput `pulumi:"env"`

	// The target of the dockerfile to build.
	Target pulumi.StringInput `pulumi:"target"`
}

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
