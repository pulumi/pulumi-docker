package main

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ecr"
	"github.com/pulumi/pulumi-docker/sdk/go/docker"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		repo, err := ecr.NewRepository(ctx, "docker-provider-test", &ecr.RepositoryArgs{
			ForceDelete: pulumi.BoolPtr(true),
		})
		if err != nil {
			return err
		}

		ctx.Export("repositoryUrl", repo.RepositoryUrl)

		repoCreds := repo.RegistryId.ApplyT(func(rid string) ([]string, error) {
			creds, err := ecr.GetCredentials(ctx, &ecr.GetCredentialsArgs{
				RegistryId: rid,
			})
			if err != nil {
				return nil, err
			}
			data, err := base64.StdEncoding.DecodeString(creds.AuthorizationToken)
			if err != nil {
				fmt.Println("error:", err)
				return nil, err
			}

			return strings.Split(string(data), ":"), nil
		}).(pulumi.StringArrayOutput)
		repoUser := repoCreds.Index(pulumi.Int(0))
		repoPass := pulumi.ToSecret(repoCreds.Index(pulumi.Int(1))).(pulumi.StringOutput)

		image, err := docker.NewImage(ctx, "build-cache-from-go", &docker.ImageArgs{
			ImageName: repo.RepositoryUrl,
			Build: docker.DockerBuildArgs{
				Context: pulumi.String("./app"),
				CacheFrom: &docker.CacheFromArgs{
					Images: pulumi.StringArray{
						pulumi.Sprintf("%v:latest", repo.RepositoryUrl),
					},
				},
			},
			Registry: docker.RegistryArgs{
				Server:   repo.RepositoryUrl,
				Username: repoUser,
				Password: repoPass,
			},
		})
		if err != nil {
			return err
		}

		ctx.Export("imageName", image.ImageName)
		return nil
	})
}
