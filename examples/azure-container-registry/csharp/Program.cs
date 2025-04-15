using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Pulumi;
using Pulumi.Azure.Authorization;
using Pulumi.Azure.Core;
using Azure = Pulumi.Azure.ContainerService;
using Docker = Pulumi.Docker;
using Pulumi.Docker.Inputs;

class Program
{
    static Task<int> Main() => Deployment.RunAsync(async () => {
        // Create a private ACR registry.
        var rg = new ResourceGroup("myrg");
        var registry = new Azure.Registry("myregistry", new Azure.RegistryArgs
        {
            ResourceGroupName = rg.Name,
        });

        // Get registry info (creds and endpoint) so we can build/publish to it.
        var imageName = Output.Format($"{registry.LoginServer}/myapp");

        // Build and publish the app image.
        var image = new Docker.Image("my-image", new Docker.ImageArgs
        {
            Build = new Docker.Inputs.DockerBuildArgs { Context = "app" },
            ImageName = imageName,
            Registry = new Docker.Inputs.RegistryArgs
            {
                Server = registry.LoginServer,
                Username = registry.AdminUsername,
                Password = registry.AdminPassword,
            },
        });

        // Export the resulting image name
        return new Dictionary<string, object>
        {
            { "baseImageName", image.BaseImageName },
            { "fullImageName", image.ImageName },
            { "repoDigest", image.RepoDigest },
        };
    });
}
