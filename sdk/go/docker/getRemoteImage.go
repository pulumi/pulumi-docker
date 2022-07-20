// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package docker

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `RemoteImage` provides details about a specific Docker Image which need to be presend on the Docker Host
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
// 		_, err := docker.LookupRemoteImage(ctx, &GetRemoteImageArgs{
// 			Name: "nginx",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = docker.LookupRemoteImage(ctx, &GetRemoteImageArgs{
// 			Name: "nginx:1.17.6",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = docker.LookupRemoteImage(ctx, &GetRemoteImageArgs{
// 			Name: "nginx@sha256:36b74457bccb56fbf8b05f79c85569501b721d4db813b684391d63e02287c0b2",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = docker.LookupRemoteImage(ctx, &GetRemoteImageArgs{
// 			Name: "nginx:1.19.1@sha256:36b74457bccb56fbf8b05f79c85569501b721d4db813b684391d63e02287c0b2",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupRemoteImage(ctx *pulumi.Context, args *LookupRemoteImageArgs, opts ...pulumi.InvokeOption) (*LookupRemoteImageResult, error) {
	var rv LookupRemoteImageResult
	err := ctx.Invoke("docker:index/getRemoteImage:getRemoteImage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRemoteImage.
type LookupRemoteImageArgs struct {
	Name string `pulumi:"name"`
}

// A collection of values returned by getRemoteImage.
type LookupRemoteImageResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string `pulumi:"id"`
	Name       string `pulumi:"name"`
	RepoDigest string `pulumi:"repoDigest"`
}

func LookupRemoteImageOutput(ctx *pulumi.Context, args LookupRemoteImageOutputArgs, opts ...pulumi.InvokeOption) LookupRemoteImageResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRemoteImageResult, error) {
			args := v.(LookupRemoteImageArgs)
			r, err := LookupRemoteImage(ctx, &args, opts...)
			var s LookupRemoteImageResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRemoteImageResultOutput)
}

// A collection of arguments for invoking getRemoteImage.
type LookupRemoteImageOutputArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupRemoteImageOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRemoteImageArgs)(nil)).Elem()
}

// A collection of values returned by getRemoteImage.
type LookupRemoteImageResultOutput struct{ *pulumi.OutputState }

func (LookupRemoteImageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRemoteImageResult)(nil)).Elem()
}

func (o LookupRemoteImageResultOutput) ToLookupRemoteImageResultOutput() LookupRemoteImageResultOutput {
	return o
}

func (o LookupRemoteImageResultOutput) ToLookupRemoteImageResultOutputWithContext(ctx context.Context) LookupRemoteImageResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRemoteImageResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteImageResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRemoteImageResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteImageResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRemoteImageResultOutput) RepoDigest() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteImageResult) string { return v.RepoDigest }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRemoteImageResultOutput{})
}
