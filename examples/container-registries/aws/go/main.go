package main

import (
	"encoding/base64"
	"errors"
	"strings"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ecr"
	"github.com/pulumi/pulumi-docker/sdk/v2/go/docker"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a private ECR registry.
		repo, err := ecr.NewRepository(ctx, "my-repo", nil)
		if err != nil {
			return err
		}

		// Get registry info (creds and endpoint) so we can build/publish to it.
		imageName := repo.RepositoryUrl
		registryInfo := repo.RegistryId.ApplyT(func(id string) (docker.ImageRegistry, error) {
			creds, err := ecr.GetCredentials(ctx, &ecr.GetCredentialsArgs{RegistryId: id})
			if err != nil {
				return docker.ImageRegistry{}, err
			}
			decoded, err := base64.StdEncoding.DecodeString(creds.AuthorizationToken)
			if err != nil {
				return docker.ImageRegistry{}, err
			}
			parts := strings.Split(string(decoded), ":")
			if len(parts) != 2 {
				return docker.ImageRegistry{}, errors.New("Invalid credentials")
			}
			return docker.ImageRegistry{
				Server:   creds.ProxyEndpoint,
				Username: parts[0],
				Password: parts[1],
			}, nil
		}).(docker.ImageRegistryOutput)

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
