// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package docker

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Creates and destroys a volume in Docker. This can be used alongside
// [docker\_container](https://www.terraform.io/docs/providers/docker/r/container.html)
// to prepare volumes that can be shared across containers.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-docker/blob/master/website/docs/r/volume.html.markdown.
type Volume struct {
	pulumi.CustomResourceState

	// Driver type for the volume (defaults to local).
	Driver pulumi.StringOutput `pulumi:"driver"`
	// Options specific to the driver.
	DriverOpts pulumi.MapOutput `pulumi:"driverOpts"`
	// User-defined key/value metadata.
	Labels VolumeLabelArrayOutput `pulumi:"labels"`
	Mountpoint pulumi.StringOutput `pulumi:"mountpoint"`
	// The name of the Docker volume (generated if not
	// provided).
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewVolume registers a new resource with the given unique name, arguments, and options.
func NewVolume(ctx *pulumi.Context,
	name string, args *VolumeArgs, opts ...pulumi.ResourceOption) (*Volume, error) {
	if args == nil {
		args = &VolumeArgs{}
	}
	var resource Volume
	err := ctx.RegisterResource("docker:index/volume:Volume", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVolume gets an existing Volume resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVolume(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VolumeState, opts ...pulumi.ResourceOption) (*Volume, error) {
	var resource Volume
	err := ctx.ReadResource("docker:index/volume:Volume", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Volume resources.
type volumeState struct {
	// Driver type for the volume (defaults to local).
	Driver *string `pulumi:"driver"`
	// Options specific to the driver.
	DriverOpts map[string]interface{} `pulumi:"driverOpts"`
	// User-defined key/value metadata.
	Labels []VolumeLabel `pulumi:"labels"`
	Mountpoint *string `pulumi:"mountpoint"`
	// The name of the Docker volume (generated if not
	// provided).
	Name *string `pulumi:"name"`
}

type VolumeState struct {
	// Driver type for the volume (defaults to local).
	Driver pulumi.StringPtrInput
	// Options specific to the driver.
	DriverOpts pulumi.MapInput
	// User-defined key/value metadata.
	Labels VolumeLabelArrayInput
	Mountpoint pulumi.StringPtrInput
	// The name of the Docker volume (generated if not
	// provided).
	Name pulumi.StringPtrInput
}

func (VolumeState) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeState)(nil)).Elem()
}

type volumeArgs struct {
	// Driver type for the volume (defaults to local).
	Driver *string `pulumi:"driver"`
	// Options specific to the driver.
	DriverOpts map[string]interface{} `pulumi:"driverOpts"`
	// User-defined key/value metadata.
	Labels []VolumeLabel `pulumi:"labels"`
	// The name of the Docker volume (generated if not
	// provided).
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a Volume resource.
type VolumeArgs struct {
	// Driver type for the volume (defaults to local).
	Driver pulumi.StringPtrInput
	// Options specific to the driver.
	DriverOpts pulumi.MapInput
	// User-defined key/value metadata.
	Labels VolumeLabelArrayInput
	// The name of the Docker volume (generated if not
	// provided).
	Name pulumi.StringPtrInput
}

func (VolumeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeArgs)(nil)).Elem()
}

