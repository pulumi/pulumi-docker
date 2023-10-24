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

// Reads the image metadata from a Docker Registry. Used in conjunction with the RemoteImage resource to keep an image up to date on the latest available version of the tag.
func LookupRegistryImage(ctx *pulumi.Context, args *LookupRegistryImageArgs, opts ...pulumi.InvokeOption) (*LookupRegistryImageResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRegistryImageResult
	err := ctx.Invoke("docker:index/getRegistryImage:getRegistryImage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRegistryImage.
type LookupRegistryImageArgs struct {
	// If `true`, the verification of TLS certificates of the server/registry is disabled. Defaults to `false`
	InsecureSkipVerify *bool `pulumi:"insecureSkipVerify"`
	// The name of the Docker image, including any tags. e.g. `alpine:latest`
	Name string `pulumi:"name"`
}

// A collection of values returned by getRegistryImage.
type LookupRegistryImageResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// If `true`, the verification of TLS certificates of the server/registry is disabled. Defaults to `false`
	InsecureSkipVerify *bool `pulumi:"insecureSkipVerify"`
	// The name of the Docker image, including any tags. e.g. `alpine:latest`
	Name string `pulumi:"name"`
	// The content digest of the image, as stored in the registry.
	Sha256Digest string `pulumi:"sha256Digest"`
}

func LookupRegistryImageOutput(ctx *pulumi.Context, args LookupRegistryImageOutputArgs, opts ...pulumi.InvokeOption) LookupRegistryImageResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRegistryImageResult, error) {
			args := v.(LookupRegistryImageArgs)
			r, err := LookupRegistryImage(ctx, &args, opts...)
			var s LookupRegistryImageResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRegistryImageResultOutput)
}

// A collection of arguments for invoking getRegistryImage.
type LookupRegistryImageOutputArgs struct {
	// If `true`, the verification of TLS certificates of the server/registry is disabled. Defaults to `false`
	InsecureSkipVerify pulumi.BoolPtrInput `pulumi:"insecureSkipVerify"`
	// The name of the Docker image, including any tags. e.g. `alpine:latest`
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupRegistryImageOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistryImageArgs)(nil)).Elem()
}

// A collection of values returned by getRegistryImage.
type LookupRegistryImageResultOutput struct{ *pulumi.OutputState }

func (LookupRegistryImageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistryImageResult)(nil)).Elem()
}

func (o LookupRegistryImageResultOutput) ToLookupRegistryImageResultOutput() LookupRegistryImageResultOutput {
	return o
}

func (o LookupRegistryImageResultOutput) ToLookupRegistryImageResultOutputWithContext(ctx context.Context) LookupRegistryImageResultOutput {
	return o
}

func (o LookupRegistryImageResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupRegistryImageResult] {
	return pulumix.Output[LookupRegistryImageResult]{
		OutputState: o.OutputState,
	}
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRegistryImageResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryImageResult) string { return v.Id }).(pulumi.StringOutput)
}

// If `true`, the verification of TLS certificates of the server/registry is disabled. Defaults to `false`
func (o LookupRegistryImageResultOutput) InsecureSkipVerify() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRegistryImageResult) *bool { return v.InsecureSkipVerify }).(pulumi.BoolPtrOutput)
}

// The name of the Docker image, including any tags. e.g. `alpine:latest`
func (o LookupRegistryImageResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryImageResult) string { return v.Name }).(pulumi.StringOutput)
}

// The content digest of the image, as stored in the registry.
func (o LookupRegistryImageResultOutput) Sha256Digest() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryImageResult) string { return v.Sha256Digest }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRegistryImageResultOutput{})
}
