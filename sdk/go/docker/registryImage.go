// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package docker

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-docker/sdk/v4/go/docker/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// <!-- Bug: Type and Name are switched -->
// Manages the lifecycle of docker image in a registry. You can upload images to a registry (= `docker push`) and also delete them again
type RegistryImage struct {
	pulumi.CustomResourceState

	// If `true`, the verification of TLS certificates of the server/registry is disabled. Defaults to `false`
	InsecureSkipVerify pulumi.BoolPtrOutput `pulumi:"insecureSkipVerify"`
	// If true, then the Docker image won't be deleted on destroy operation. If this is false, it will delete the image from the docker registry on destroy operation. Defaults to `false`
	KeepRemotely pulumi.BoolPtrOutput `pulumi:"keepRemotely"`
	// The name of the Docker image.
	Name pulumi.StringOutput `pulumi:"name"`
	// The sha256 digest of the image.
	Sha256Digest pulumi.StringOutput `pulumi:"sha256Digest"`
	// A map of arbitrary strings that, when changed, will force the `RegistryImage` resource to be replaced. This can be used to repush a local image
	Triggers pulumi.MapOutput `pulumi:"triggers"`
}

// NewRegistryImage registers a new resource with the given unique name, arguments, and options.
func NewRegistryImage(ctx *pulumi.Context,
	name string, args *RegistryImageArgs, opts ...pulumi.ResourceOption) (*RegistryImage, error) {
	if args == nil {
		args = &RegistryImageArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RegistryImage
	err := ctx.RegisterResource("docker:index/registryImage:RegistryImage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegistryImage gets an existing RegistryImage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegistryImage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegistryImageState, opts ...pulumi.ResourceOption) (*RegistryImage, error) {
	var resource RegistryImage
	err := ctx.ReadResource("docker:index/registryImage:RegistryImage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegistryImage resources.
type registryImageState struct {
	// If `true`, the verification of TLS certificates of the server/registry is disabled. Defaults to `false`
	InsecureSkipVerify *bool `pulumi:"insecureSkipVerify"`
	// If true, then the Docker image won't be deleted on destroy operation. If this is false, it will delete the image from the docker registry on destroy operation. Defaults to `false`
	KeepRemotely *bool `pulumi:"keepRemotely"`
	// The name of the Docker image.
	Name *string `pulumi:"name"`
	// The sha256 digest of the image.
	Sha256Digest *string `pulumi:"sha256Digest"`
	// A map of arbitrary strings that, when changed, will force the `RegistryImage` resource to be replaced. This can be used to repush a local image
	Triggers map[string]interface{} `pulumi:"triggers"`
}

type RegistryImageState struct {
	// If `true`, the verification of TLS certificates of the server/registry is disabled. Defaults to `false`
	InsecureSkipVerify pulumi.BoolPtrInput
	// If true, then the Docker image won't be deleted on destroy operation. If this is false, it will delete the image from the docker registry on destroy operation. Defaults to `false`
	KeepRemotely pulumi.BoolPtrInput
	// The name of the Docker image.
	Name pulumi.StringPtrInput
	// The sha256 digest of the image.
	Sha256Digest pulumi.StringPtrInput
	// A map of arbitrary strings that, when changed, will force the `RegistryImage` resource to be replaced. This can be used to repush a local image
	Triggers pulumi.MapInput
}

func (RegistryImageState) ElementType() reflect.Type {
	return reflect.TypeOf((*registryImageState)(nil)).Elem()
}

type registryImageArgs struct {
	// If `true`, the verification of TLS certificates of the server/registry is disabled. Defaults to `false`
	InsecureSkipVerify *bool `pulumi:"insecureSkipVerify"`
	// If true, then the Docker image won't be deleted on destroy operation. If this is false, it will delete the image from the docker registry on destroy operation. Defaults to `false`
	KeepRemotely *bool `pulumi:"keepRemotely"`
	// The name of the Docker image.
	Name *string `pulumi:"name"`
	// A map of arbitrary strings that, when changed, will force the `RegistryImage` resource to be replaced. This can be used to repush a local image
	Triggers map[string]interface{} `pulumi:"triggers"`
}

// The set of arguments for constructing a RegistryImage resource.
type RegistryImageArgs struct {
	// If `true`, the verification of TLS certificates of the server/registry is disabled. Defaults to `false`
	InsecureSkipVerify pulumi.BoolPtrInput
	// If true, then the Docker image won't be deleted on destroy operation. If this is false, it will delete the image from the docker registry on destroy operation. Defaults to `false`
	KeepRemotely pulumi.BoolPtrInput
	// The name of the Docker image.
	Name pulumi.StringPtrInput
	// A map of arbitrary strings that, when changed, will force the `RegistryImage` resource to be replaced. This can be used to repush a local image
	Triggers pulumi.MapInput
}

func (RegistryImageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registryImageArgs)(nil)).Elem()
}

type RegistryImageInput interface {
	pulumi.Input

	ToRegistryImageOutput() RegistryImageOutput
	ToRegistryImageOutputWithContext(ctx context.Context) RegistryImageOutput
}

func (*RegistryImage) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistryImage)(nil)).Elem()
}

func (i *RegistryImage) ToRegistryImageOutput() RegistryImageOutput {
	return i.ToRegistryImageOutputWithContext(context.Background())
}

func (i *RegistryImage) ToRegistryImageOutputWithContext(ctx context.Context) RegistryImageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryImageOutput)
}

func (i *RegistryImage) ToOutput(ctx context.Context) pulumix.Output[*RegistryImage] {
	return pulumix.Output[*RegistryImage]{
		OutputState: i.ToRegistryImageOutputWithContext(ctx).OutputState,
	}
}

// RegistryImageArrayInput is an input type that accepts RegistryImageArray and RegistryImageArrayOutput values.
// You can construct a concrete instance of `RegistryImageArrayInput` via:
//
//	RegistryImageArray{ RegistryImageArgs{...} }
type RegistryImageArrayInput interface {
	pulumi.Input

	ToRegistryImageArrayOutput() RegistryImageArrayOutput
	ToRegistryImageArrayOutputWithContext(context.Context) RegistryImageArrayOutput
}

type RegistryImageArray []RegistryImageInput

func (RegistryImageArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RegistryImage)(nil)).Elem()
}

func (i RegistryImageArray) ToRegistryImageArrayOutput() RegistryImageArrayOutput {
	return i.ToRegistryImageArrayOutputWithContext(context.Background())
}

func (i RegistryImageArray) ToRegistryImageArrayOutputWithContext(ctx context.Context) RegistryImageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryImageArrayOutput)
}

func (i RegistryImageArray) ToOutput(ctx context.Context) pulumix.Output[[]*RegistryImage] {
	return pulumix.Output[[]*RegistryImage]{
		OutputState: i.ToRegistryImageArrayOutputWithContext(ctx).OutputState,
	}
}

// RegistryImageMapInput is an input type that accepts RegistryImageMap and RegistryImageMapOutput values.
// You can construct a concrete instance of `RegistryImageMapInput` via:
//
//	RegistryImageMap{ "key": RegistryImageArgs{...} }
type RegistryImageMapInput interface {
	pulumi.Input

	ToRegistryImageMapOutput() RegistryImageMapOutput
	ToRegistryImageMapOutputWithContext(context.Context) RegistryImageMapOutput
}

type RegistryImageMap map[string]RegistryImageInput

func (RegistryImageMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RegistryImage)(nil)).Elem()
}

func (i RegistryImageMap) ToRegistryImageMapOutput() RegistryImageMapOutput {
	return i.ToRegistryImageMapOutputWithContext(context.Background())
}

func (i RegistryImageMap) ToRegistryImageMapOutputWithContext(ctx context.Context) RegistryImageMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryImageMapOutput)
}

func (i RegistryImageMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*RegistryImage] {
	return pulumix.Output[map[string]*RegistryImage]{
		OutputState: i.ToRegistryImageMapOutputWithContext(ctx).OutputState,
	}
}

type RegistryImageOutput struct{ *pulumi.OutputState }

func (RegistryImageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistryImage)(nil)).Elem()
}

func (o RegistryImageOutput) ToRegistryImageOutput() RegistryImageOutput {
	return o
}

func (o RegistryImageOutput) ToRegistryImageOutputWithContext(ctx context.Context) RegistryImageOutput {
	return o
}

func (o RegistryImageOutput) ToOutput(ctx context.Context) pulumix.Output[*RegistryImage] {
	return pulumix.Output[*RegistryImage]{
		OutputState: o.OutputState,
	}
}

// If `true`, the verification of TLS certificates of the server/registry is disabled. Defaults to `false`
func (o RegistryImageOutput) InsecureSkipVerify() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RegistryImage) pulumi.BoolPtrOutput { return v.InsecureSkipVerify }).(pulumi.BoolPtrOutput)
}

// If true, then the Docker image won't be deleted on destroy operation. If this is false, it will delete the image from the docker registry on destroy operation. Defaults to `false`
func (o RegistryImageOutput) KeepRemotely() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RegistryImage) pulumi.BoolPtrOutput { return v.KeepRemotely }).(pulumi.BoolPtrOutput)
}

// The name of the Docker image.
func (o RegistryImageOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryImage) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The sha256 digest of the image.
func (o RegistryImageOutput) Sha256Digest() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryImage) pulumi.StringOutput { return v.Sha256Digest }).(pulumi.StringOutput)
}

// A map of arbitrary strings that, when changed, will force the `RegistryImage` resource to be replaced. This can be used to repush a local image
func (o RegistryImageOutput) Triggers() pulumi.MapOutput {
	return o.ApplyT(func(v *RegistryImage) pulumi.MapOutput { return v.Triggers }).(pulumi.MapOutput)
}

type RegistryImageArrayOutput struct{ *pulumi.OutputState }

func (RegistryImageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RegistryImage)(nil)).Elem()
}

func (o RegistryImageArrayOutput) ToRegistryImageArrayOutput() RegistryImageArrayOutput {
	return o
}

func (o RegistryImageArrayOutput) ToRegistryImageArrayOutputWithContext(ctx context.Context) RegistryImageArrayOutput {
	return o
}

func (o RegistryImageArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*RegistryImage] {
	return pulumix.Output[[]*RegistryImage]{
		OutputState: o.OutputState,
	}
}

func (o RegistryImageArrayOutput) Index(i pulumi.IntInput) RegistryImageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RegistryImage {
		return vs[0].([]*RegistryImage)[vs[1].(int)]
	}).(RegistryImageOutput)
}

type RegistryImageMapOutput struct{ *pulumi.OutputState }

func (RegistryImageMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RegistryImage)(nil)).Elem()
}

func (o RegistryImageMapOutput) ToRegistryImageMapOutput() RegistryImageMapOutput {
	return o
}

func (o RegistryImageMapOutput) ToRegistryImageMapOutputWithContext(ctx context.Context) RegistryImageMapOutput {
	return o
}

func (o RegistryImageMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*RegistryImage] {
	return pulumix.Output[map[string]*RegistryImage]{
		OutputState: o.OutputState,
	}
}

func (o RegistryImageMapOutput) MapIndex(k pulumi.StringInput) RegistryImageOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RegistryImage {
		return vs[0].(map[string]*RegistryImage)[vs[1].(string)]
	}).(RegistryImageOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RegistryImageInput)(nil)).Elem(), &RegistryImage{})
	pulumi.RegisterInputType(reflect.TypeOf((*RegistryImageArrayInput)(nil)).Elem(), RegistryImageArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RegistryImageMapInput)(nil)).Elem(), RegistryImageMap{})
	pulumi.RegisterOutputType(RegistryImageOutput{})
	pulumi.RegisterOutputType(RegistryImageArrayOutput{})
	pulumi.RegisterOutputType(RegistryImageMapOutput{})
}
