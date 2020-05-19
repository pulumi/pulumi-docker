package main

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ecr"
	"github.com/pulumi/pulumi-docker/sdk/v2/go/docker"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func getRegistryInfo(ctx *pulumi.Context, rid string) (pulumi.Map, error) {
	creds, err := ecr.GetCredentials(ctx, &ecr.GetCredentialsArgs{
		RegistryId: rid,
	})
	if err != nil {
		return nil, err
	}
	decoded, err := base64.StdEncoding.DecodeString(creds.AuthorizationToken)
	if err != nil {
		return nil, err
	}
	parts := strings.Split(string(decoded), ":")
	if len(parts) != 2 {
		return nil, errors.Errorf("Invalid credentials")
	}
	return map[string]pulumi.Input{
		"server":   pulumi.String(creds.ProxyEndpoint),
		"username": pulumi.String(parts[0]),
		"password": pulumi.String(parts[1]),
	}, nil
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

			registry := repo.RegistryId.ApplyString(getRegistryInfo)

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
						Server:   registry.MapIndex(pulumi.String("Server")),
						Username: registry.MapIndex(pulumi.String("Username")),
						Password: registry.MapIndex(pulumi.String("Password")),
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
