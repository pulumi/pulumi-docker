package main

import (
	"github.com/pulumi/pulumi-docker/sdk/v4/go/docker"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		imageArgs := &docker.ImageArgs{
			ImageName: pulumi.String("pulumi-user/example:v1.0.0"),
			Build: docker.DockerBuildArgs{

				Context: pulumi.String("app"),
			},

			SkipPush: pulumi.Bool(true),
		}
		image, err := docker.NewImage(ctx, "my-image", imageArgs)
		if err != nil {
			return err
		}

		ctx.Export("deps-image", image.ImageName)

		return nil
	})
}
