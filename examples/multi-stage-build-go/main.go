package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ecr"
	"github.com/pulumi/pulumi-docker/sdk/v4/go/docker"
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

		token := ecr.GetAuthorizationTokenOutput(ctx, ecr.GetAuthorizationTokenOutputArgs{
			RegistryId: repo.RegistryId,
		})

		imageName := pulumi.Sprintf("%v:latest", repo.RepositoryUrl)

		image, err := docker.NewImage(ctx, "build-cache-from-go", &docker.ImageArgs{
			ImageName: imageName,
			Build: docker.DockerBuildArgs{
				Context: pulumi.String("./app"),
				CacheFrom: &docker.CacheFromArgs{
					Images: pulumi.StringArray{
						pulumi.Sprintf("%v:latest", repo.RepositoryUrl),
					},
				},
				Args: pulumi.StringMap{
					"BUILDKIT_INLINE_CACHE": pulumi.String("1"),
				},
			},
			Registry: docker.RegistryArgs{
				Server:   repo.RepositoryUrl,
				Username: token.UserName(),
				Password: pulumi.ToSecret(token.Password()).(pulumi.StringOutput),
			},
		})
		if err != nil {
			return err
		}

		ctx.Export("imageName", image.RepoDigest)
		return nil
	})
}
