// Copyright 2016-2020, Pulumi Corporation.

package docker

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// ImageRegistry contains credentials for the docker registry.
type ImageRegistry struct {
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

type imageRegistry struct {
	Server   string
	Username string
	Password string
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
	Build DockerBuild

	// The docker image name to build locally before tagging with imageName.  If not provided, it
	// will be given the value of to [imageName].  This name can include a tag at the end.  If
	// provided all pushed image resources will contain that tag as well.
	//
	// Either [imageName] or [localImageName] can have a tag.  However, if both have a tag, then
	// those tags must match.
	LocalImageName pulumi.StringInput

	// Credentials for the docker registry to push to.
	Registry *ImageRegistry

	// Skip push flag.
	SkipPush pulumi.BoolInput
}

type imageArgs struct {
	ImageName      string        `pulumi:"imageName"`
	Build          dockerBuild   `pulumi:"build"`
	LocalImageName string        `pulumi:"localImageName"`
	Registry       imageRegistry `pulumi:"registry"`
	SkipPush       bool          `pulumi:"skipPush"`
}

func (ImageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*imageArgs)(nil)).Elem()
}

// Image is a resource represents a Docker image built locally which is published and made
// available via a remote Docker registry.  This can be used to ensure that a Docker source
// directory from a local deployment environment is built and pushed to a cloud-hosted Docker
// registry as part of a Pulumi deployment, so that it can be referenced as an image input from
// other cloud services that reference Docker images - including Kubernetes Pods, AWS ECS Tasks, and
// Azure Container Instances.
type Image struct {
	pulumi.ResourceState

	// The base image name that was built and pushed.  This does not include the id annotation, so
	// is not pinned to the specific build performed by this docker.Image.
	BaseImageName pulumi.StringOutput

	// The unique pinned image name on the remote repository.
	ImageName pulumi.StringOutput

	// The server the image is located at.
	RegistryServer pulumi.StringOutput
}

// NewImage registers a new image with the given unique name, arguments, and options.
func NewImage(ctx *pulumi.Context,
	name string, args *ImageArgs, opts ...pulumi.ResourceOption) (*Image, error) {

	resource := &Image{}
	err := ctx.RegisterComponentResource("docker:image:Image", name, resource, opts...)
	if err != nil {
		return nil, err
	}

	resource.ImageName = pulumi.All(args).ApplyT(func(inputArgs []interface{}) (string, error) {
		imageArgs := inputArgs[0].(*imageArgs)
		imageName := imageArgs.ImageName

		// If there is no localImageName set it equal to imageName.  Note: this means
		// that if imageName contains a tag, localImageName will contain the same tag.
		localImageName := imageArgs.LocalImageName
		if len(localImageName) == 0 {
			localImageName = imageName
		}

		skipPush := imageArgs.SkipPush

		// Now break both the localImageName and the imageName into the untagged part and the
		// optional tag.  If both have tags, they must match.  If one or the other has a tag, we
		// just use that as the tag to use.  This allows users to flexibly provide a tag on one
		// option or the other and still have it work out.
		localImageNameWithoutTag, localImageNameTag := getImageNameAndTag(localImageName)
		imageNameWithoutTag, imageNameTag := getImageNameAndTag(imageName)

		if len(localImageNameTag) > 0 && len(imageNameTag) > 0 && localImageNameTag != imageNameTag {
			return "", errors.Errorf("%v and %v had mismatched tags. %s != %s",
				args.LocalImageName, args.ImageName, localImageNameTag, imageNameTag)
		}

		var tag string
		if len(localImageNameTag) == 0 {
			tag = imageNameTag
		} else {
			tag = localImageNameTag
		}

		// BuildAndPushImageAsync expects only the baseImageName to have a tag.  So build that
		// name appropriately if we were given a tag.
		var baseImageName string
		if len(tag) == 0 {
			baseImageName = localImageName
		} else {
			baseImageName = fmt.Sprintf("%s:%s", localImageNameWithoutTag, localImageName)
		}

		return buildAndPushImageAsync(ctx, baseImageName, imageArgs.Build,
			imageNameWithoutTag, resource, skipPush, imageArgs.Registry)
	}).(pulumi.StringOutput)

	resource.RegistryServer = args.Registry.Server.ToStringOutput()
	resource.BaseImageName = args.ImageName.ToStringOutput()

	outputs := pulumi.Map(map[string]pulumi.Input{
		"baseImageName":  resource.BaseImageName,
		"imageName":      resource.ImageName,
		"registryServer": resource.RegistryServer,
	})
	err = ctx.RegisterResourceOutputs(resource, outputs)
	if err != nil {
		return nil, err
	}

	return resource, nil
}
