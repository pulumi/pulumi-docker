// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package docker

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Reads the image metadata from a Docker Registry. Used in conjunction with the RemoteImage resource to keep an image up to date on the latest available version of the tag.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-docker/sdk/v3/go/docker"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		ubuntuRegistryImage, err := docker.LookupRegistryImage(ctx, &GetRegistryImageArgs{
// 			Name: "ubuntu:precise",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = docker.NewRemoteImage(ctx, "ubuntuRemoteImage", &docker.RemoteImageArgs{
// 			Name: pulumi.String(ubuntuRegistryImage.Name),
// 			PullTriggers: pulumi.StringArray{
// 				pulumi.String(ubuntuRegistryImage.Sha256Digest),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupRegistryImage(ctx *pulumi.Context, args *LookupRegistryImageArgs, opts ...pulumi.InvokeOption) (*LookupRegistryImageResult, error) {
	var rv LookupRegistryImageResult
	err := ctx.Invoke("docker:index/getRegistryImage:getRegistryImage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRegistryImage.
type LookupRegistryImageArgs struct {
	InsecureSkipVerify *bool  `pulumi:"insecureSkipVerify"`
	Name               string `pulumi:"name"`
}

// A collection of values returned by getRegistryImage.
type LookupRegistryImageResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id                 string `pulumi:"id"`
	InsecureSkipVerify *bool  `pulumi:"insecureSkipVerify"`
	Name               string `pulumi:"name"`
	Sha256Digest       string `pulumi:"sha256Digest"`
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
	InsecureSkipVerify pulumi.BoolPtrInput `pulumi:"insecureSkipVerify"`
	Name               pulumi.StringInput  `pulumi:"name"`
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

// The provider-assigned unique ID for this managed resource.
func (o LookupRegistryImageResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryImageResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRegistryImageResultOutput) InsecureSkipVerify() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRegistryImageResult) *bool { return v.InsecureSkipVerify }).(pulumi.BoolPtrOutput)
}

func (o LookupRegistryImageResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryImageResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRegistryImageResultOutput) Sha256Digest() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryImageResult) string { return v.Sha256Digest }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRegistryImageResultOutput{})
}
