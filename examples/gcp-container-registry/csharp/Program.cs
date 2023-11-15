using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Pulumi;
using Pulumi.Docker;
using Pulumi.Docker.Inputs;
using Pulumi.Gcp.Container;

class Program
{
    static Task<int> Main() => Deployment.RunAsync(async () => {
        // Create a private GCR registry.
        var registry = new Registry("my-registry");
        var registryUrl = registry.Id.Apply(async _ => {
            return (await GetRegistryRepository.InvokeAsync()).RepositoryUrl;
        });

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
            { "repoDigest": image.RepoDigest },
        };
    });
}
