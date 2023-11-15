// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package docker

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-docker/sdk/v4/go/docker/internal"
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
//
//	"github.com/pulumi/pulumi-docker/sdk/v4/go/docker"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := docker.LookupRemoteImage(ctx, &docker.LookupRemoteImageArgs{
//				Name: "nginx",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = docker.LookupRemoteImage(ctx, &docker.LookupRemoteImageArgs{
//				Name: "nginx:1.17.6",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = docker.LookupRemoteImage(ctx, &docker.LookupRemoteImageArgs{
//				Name: "nginx@sha256:36b74457bccb56fbf8b05f79c85569501b721d4db813b684391d63e02287c0b2",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = docker.LookupRemoteImage(ctx, &docker.LookupRemoteImageArgs{
//				Name: "nginx:1.19.1@sha256:36b74457bccb56fbf8b05f79c85569501b721d4db813b684391d63e02287c0b2",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupRemoteImage(ctx *pulumi.Context, args *LookupRemoteImageArgs, opts ...pulumi.InvokeOption) (*LookupRemoteImageResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRemoteImageResult
	err := ctx.Invoke("docker:index/getRemoteImage:getRemoteImage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRemoteImage.
type LookupRemoteImageArgs struct {
	// The name of the Docker image, including any tags or SHA256 repo digests.
	Name string `pulumi:"name"`
}

// A collection of values returned by getRemoteImage.
type LookupRemoteImageResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The name of the Docker image, including any tags or SHA256 repo digests.
	Name string `pulumi:"name"`
	// The image sha256 digest in the form of `repo[:tag]@sha256:<hash>`. It may be empty in the edge case where the local image was pulled from a repo, tagged locally, and then referred to in the data source by that local name/tag.
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
	// The name of the Docker image, including any tags or SHA256 repo digests.
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

// The name of the Docker image, including any tags or SHA256 repo digests.
func (o LookupRemoteImageResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteImageResult) string { return v.Name }).(pulumi.StringOutput)
}

// The image sha256 digest in the form of `repo[:tag]@sha256:<hash>`. It may be empty in the edge case where the local image was pulled from a repo, tagged locally, and then referred to in the data source by that local name/tag.
func (o LookupRemoteImageResultOutput) RepoDigest() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteImageResult) string { return v.RepoDigest }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRemoteImageResultOutput{})
}
