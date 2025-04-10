using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Pulumi;
using Pulumi.Docker;
using Pulumi.Docker.Inputs;
using Pulumi.Gcp.ArtifactRegistry;
using Pulumi.Random;

class Program
{
    static Task<int> Main() => Deployment.RunAsync(async () => {
        // Create a random suffix for the repository name
        var randomSuffix = new RandomString("random-suffix", new RandomStringArgs
        {
            Length = 6,
            Special = false,
            Upper = false,
        });

        // Concatenate repository name with random suffix
        var repoName = Output.Format($"docker-test-repo-{randomSuffix.Result}");

        / Create a private GCP artifact registry
        var registry = new Repository("my-registry", new RepositoryArgs
        {
            Format = "DOCKER",
            RepositoryId = repoName,
            Location = "us-central1", // change to your desired region
            DockerConfig = new RepositoryDockerConfigArgs
            {
                ImmutableTags = false,
            },
        });

        // Form the registry URL
        var registryUrl = Output.Format($"{registry.Location}-docker.pkg.dev/{registry.Project}/{registry.RepositoryId}");


        // Get image name
        var imageName = Output.Format($"{registryUrl}/myapp");

        // Build and publish the app image.
        var image = new Image("my-image", new ImageArgs
        {
            Build = new Pulumi.Docker.Inputs.DockerBuildArgs { Context = "app" },
            ImageName = imageName,
        });

        // Export the resulting base name in addition to the specific version pushed.
        return new Dictionary<string, object>
        {
            { "baseImageName", image.BaseImageName },
            { "fullImageName", image.ImageName },
            { "repoDigest", image.RepoDigest },
        };
    });
}
