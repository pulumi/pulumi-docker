package main

import (
	"github.com/pulumi/pulumi-docker/sdk/v2/go/docker"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Fetch the Docker Hub auth info from config.
		cfg := config.New(ctx, "")
		username := cfg.Require("dockerUsername")
		password := cfg.RequireSecret("dockerPassword")

		// Populate the registry info (creds and endpoint).
		imageName := pulumi.String(username + "/myapp")
		registryInfo := docker.ImageRegistryArgs{
			Server:   pulumi.String("docker.io"),
			Username: pulumi.String(username),
			Password: password.(pulumi.StringOutput),
		}

		// Build and publish the app image.
		image, err := docker.NewImage(ctx, "my-image", &docker.ImageArgs{
			Build:     &docker.DockerBuildArgs{Context: pulumi.String("app")},
			ImageName: imageName,
			Registry:  registryInfo,
		})
		if err != nil {
			return err
		}

		// Export the resulting image name and tag.
		ctx.Export("imageName", image.ImageName)
		return nil
	})
}
