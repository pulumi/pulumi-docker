package main

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ecr"
	"github.com/pulumi/pulumi-docker/sdk/v3/go/docker"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		repo, err := ecr.NewRepository(ctx, "foo", &ecr.RepositoryArgs{})
		if err != nil {
			return err
		}

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
		repoPass := repoCreds.Index(pulumi.Int(1))

		image, err := docker.NewImage(ctx, "my-image", &docker.ImageArgs{
			Build: docker.DockerBuildArgs{
				Context:    pulumi.String("./app"),
				Dockerfile: pulumi.String("./app/Dockerfile-multistage"),
				CacheFrom: docker.CacheFromPtr(&docker.CacheFromArgs{
					Stages: pulumi.StringArray{
						pulumi.String("builder"),
					},
				}),
			},
			ImageName: repo.RepositoryUrl,
			Registry: docker.ImageRegistryArgs{
				Server:   repo.RepositoryUrl,
				Username: repoUser,
				Password: repoPass,
			},
		})

		ctx.Export("imageName", image.ImageName)

		return nil
	})
}
