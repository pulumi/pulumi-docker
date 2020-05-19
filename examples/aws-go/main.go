package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ecr"
	"github.com/pulumi/pulumi-docker/sdk/v2/go/docker"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func getRegistryInfo(ctx *pulumi.Context, rid string) (map[string]pulumi.Input, error) {
	creds, err := ecr.GetCredentials(ctx, &ecr.GetCredentialsArgs{
		RegistryId: rid,
	})
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		loop := 0
		for loop < 4 {
			repoName := fmt.Sprintf("my-repo-%d", loop)
			imageName := fmt.Sprintf("image-%d", loop)
			repo, err := ecr.NewRepository(ctx, repoName, &ecr.RepositoryArgs{
				Name: pulumi.String(imageName),
			})
			if err != nil {
				return err
			}

			registry := repo.RegistryId.ApplyT(func(rid string) map[string]pulumi.Input {
				registryInfo, err := getRegistryInfo(ctx, rid)
				if err != nil {
					return map[string]pulumi.Input{}
				}
				return registryInfo
			})

			_, err = docker.NewImage(
				ctx,
				fmt.Sprintf("my-%s", imageName),
				&docker.ImageArgs{
					ImageName: repo.RepositoryUrl,
					Build: docker.DockerBuildArgs{
						Context: pulumi.String("app"),
						Args: pulumi.Map{
							"parameter": pulumi.String(loop),
						},
					},
					Registry: docker.ImageRegistryArgs{
						Server:   registry["server"],
						Username: registry["username"],
						Password: registry["password"],
					},
				},
			)
			if err != nil {
				return err
			}
			loop += 1
		}
		return nil
	})
}
