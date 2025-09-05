package main

import (
	"github.com/pulumi/pulumi-docker/sdk/v4/go/docker"
	"github.com/pulumi/pulumi-gcp/sdk/v9/go/gcp/artifactregistry"
	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a random suffix for the repository name
		randomSuffix, err := random.NewRandomString(ctx, "random-suffix", &random.RandomStringArgs{
			Length:  pulumi.Int(6),
			Special: pulumi.Bool(false),
			Upper:   pulumi.Bool(false),
		})
		if err != nil {
			return err
		}

		// Concatenate repository name with random suffix
		repoName := pulumi.Sprintf("docker-test-repo-%s", randomSuffix.Result)

		// Create a private GCP artifact registry
		registry, err := artifactregistry.NewRepository(ctx, "my-registry", &artifactregistry.RepositoryArgs{
			Format:       pulumi.String("DOCKER"),
			RepositoryId: repoName,
			Location:     pulumi.String("us-central1"), // change to your desired region
		})
		if err != nil {
			return err
		}

		// Form the registry URL
		registryUrl := pulumi.Sprintf("%s-docker.pkg.dev/%s/%s", registry.Location, registry.Project, registry.RepositoryId)

		// Define the image name using the registry URL
		imageName := pulumi.Sprintf("%s/myapp", registryUrl)

		// Build and publish the app image.
		image, err := docker.NewImage(ctx, "my-image", &docker.ImageArgs{
			Build:     &docker.DockerBuildArgs{Context: pulumi.String("app")},
			ImageName: imageName,
		})
		if err != nil {
			return err
		}

		// Export the resulting image name
		ctx.Export("fullImageName", image.ImageName)
		ctx.Export("repoDigest", image.RepoDigest)
		return nil
	})
}
