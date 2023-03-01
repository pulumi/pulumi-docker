package main

import (
	"github.com/pulumi/pulumi-docker/sdk/v4/go/docker"
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/container"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a private GCR registry.
		registry, err := container.NewRegistry(ctx, "my-registry", nil)
		if err != nil {
			return err
		}
		registryUrl := registry.ID().ApplyT(func(_ string) (string, error) {
			rep, err := container.GetRegistryRepository(ctx, nil)
			if err != nil {
				return "", err
			}
			return rep.RepositoryUrl, nil
		})

		imageName := pulumi.Sprintf("%s/myapp", registryUrl)

		// Build and publish the app image.
		image, err := docker.NewImage(ctx, "my-image", &docker.ImageArgs{
			Build:     &docker.DockerBuildArgs{Context: pulumi.String("app")},
			ImageName: imageName,
		})

		// Export the resulting image name
		ctx.Export("fullImageName", image.ImageName)
		return nil
	})
}
