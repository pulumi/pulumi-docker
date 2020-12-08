package main

import (
	"github.com/pulumi/pulumi-docker/sdk/v2/go/docker"
	"github.com/pulumi/pulumi-gcp/sdk/v2/go/gcp/container"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a private GCR registry.
		registry, err := container.NewRegistry(ctx, "my-registry", nil)
		if err != nil {
			return err
		}
		registryUrl := registry.ID().ApplyString(func(_ string) (string, error) {
			rep, err := container.GetRegistryRepository(ctx, nil)
			if err != nil {
				return "", err
			}
			return rep.RepositoryUrl, nil
		})

		// Get registry info (creds and endpoint).
		imageName := pulumi.Sprintf("%s/myapp", registryUrl)
		registryInfo := docker.ImageRegistryArgs{} // use gcloud

		// Build and publish the app image.
		image, err := docker.NewImage(ctx, "my-image", &docker.ImageArgs{
			Build:     &docker.DockerBuildArgs{Context: pulumi.String("app")},
			ImageName: imageName,
			Registry:  registryInfo,
		})

		// Export the resulting base name in addition to the specific version pushed.
		ctx.Export("baseImageName", image.BaseImageName)
		ctx.Export("fullImageName", image.ImageName)
		return nil
	})
}
