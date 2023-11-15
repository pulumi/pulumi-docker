package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"

	"github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean"
	"github.com/pulumi/pulumi-docker/sdk/v4/go/docker"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a private DigitalOcean Container Registry.
		registry, err := digitalocean.NewContainerRegistry(ctx, "my-reg",
			&digitalocean.ContainerRegistryArgs{
				SubscriptionTierSlug: pulumi.String("starter"),
			})
		if err != nil {
			return err
		}

		// Get registry info (creds and endpoint).
		imageName := pulumi.Sprintf("%s/myapp", registry.Endpoint)
		creds, err := digitalocean.NewContainerRegistryDockerCredentials(ctx, "my-reg-creds",
			&digitalocean.ContainerRegistryDockerCredentialsArgs{
				RegistryName: registry.Name,
				Write:        pulumi.Bool(true),
			},
		)
		if err != nil {
			return err
		}

		registryInfo := pulumi.All(creds.DockerCredentials, registry.ServerUrl).ApplyT(
			func(args []interface{}) (docker.Registry, error) {
				// We are given a Docker creds file; parse it to find the temp username/password.
				authJson := args[0].(string)
				serverUrl := args[1].(string)
				var auths map[string]interface{}
				if err := json.Unmarshal([]byte(authJson), &auths); err != nil {
					return docker.Registry{}, err
				}
				authMap := auths["auths"].(map[string]interface{})
				authToken := authMap[serverUrl].(map[string]interface{})["auth"].(string)
				decoded, err := base64.StdEncoding.DecodeString(authToken)
				if err != nil {
					return docker.Registry{}, err
				}
				parts := strings.Split(string(decoded), ":")
				if len(parts) != 2 {
					return docker.Registry{}, errors.New("Invalid credentials")
				}
				return docker.Registry{
					Server:   &serverUrl,
					Username: &parts[0],
					Password: &parts[1],
				}, nil
			},
		).(docker.RegistryOutput)

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
