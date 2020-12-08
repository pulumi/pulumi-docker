using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Pulumi;
using Pulumi.Docker;
using Pulumi.Gcp.Container;

class Program
{
    static Task<int> Main() => Deployment.RunAsync(async () => {
        // Create a private GCR registry.
        var registry = new Registry("my-registry");
        var registryUrl = registry.Id.Apply(async _ => {
            return (await GetRegistryRepository.InvokeAsync()).RepositoryUrl;
        }); 

        // Get registry info (creds and endpoint).
        var imageName = Output.Format($"{registryUrl}/myapp");
        //var registryInfo = new ImageRegistry(); // use gcloud for authentication.

        // Build and publish the app image.
        var image = new Image("my-image", new ImageArgs
        {
            Build = new DockerBuild { Context = "app" },
            ImageName = imageName,
            //Registry = registryInfo,
        });

        // Export the resulting base name in addition to the specific version pushed.
        return new Dictionary<string, object>
        {
            { "baseImageName", image.BaseImageName },
            { "fullImageName", image.ImageName },
        };
    });
}
