package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws/ecr"
	"github.com/pulumi/pulumi-docker/sdk/v4/go/docker"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a private ECR registry.
		repo, err := ecr.NewRepository(ctx, "my-repo", &ecr.RepositoryArgs{
			ForceDelete: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}

		// Get registry info (creds and endpoint) so we can build/publish to it.
		imageName := repo.RepositoryUrl
		registryInfo := repo.RegistryId.ApplyT(func(id string) (docker.Registry, error) {
			creds, err := ecr.GetAuthorizationToken(ctx, &ecr.GetAuthorizationTokenArgs{RegistryId: &id})
			if err != nil {
				return docker.Registry{}, err
			}
			return docker.Registry{
				Server:   &creds.ProxyEndpoint,
				Username: &creds.UserName,
				Password: &creds.Password,
			}, nil
		}).(docker.RegistryOutput)

		// Build and publish the app image.
		image, err := docker.NewImage(ctx, "my-image", &docker.ImageArgs{
			Build:     &docker.DockerBuildArgs{Context: pulumi.String("app")},
			ImageName: imageName,
			Registry:  registryInfo,
		})

		// Export the resulting image name
		ctx.Export("fullImageName", image.ImageName)
		ctx.Export("repoDigest", image.RepoDigest)
		return nil
	})
}
