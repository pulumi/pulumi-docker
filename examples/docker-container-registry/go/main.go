package main

import (
	"github.com/pulumi/pulumi-docker/sdk/v4/go/docker"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Fetch the Docker Hub auth info from config.
		cfg := config.New(ctx, "")
		username := cfg.Require("dockerUsername")
		password := cfg.RequireSecret("dockerPassword")
		server := "docker.io"

		// Populate the registry info (creds and endpoint).
		imageName := server + "/" + username + "/myapp"

		registryInfo := password.ApplyT(func(pw string) (docker.Registry, error) {
			return docker.Registry{
				Server:   &server,
				Username: &username,
				Password: &pw,
			}, nil
		}).(docker.RegistryOutput)

		// Build and publish the app image.
		image, err := docker.NewImage(ctx, "my-image", &docker.ImageArgs{
			Build:     &docker.DockerBuildArgs{Context: pulumi.String("app")},
			ImageName: pulumi.String(imageName),
			Registry:  registryInfo,
		})
		if err != nil {
			return err
		}

		// Export the resulting image name and tag.
		ctx.Export("imageName", image.ImageName)
		ctx.Export("repoDigest", image.RepoDigest)
		return nil
	})
}
