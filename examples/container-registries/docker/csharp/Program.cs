using System;
using System.Collections.Generic;
using System.Text;
using System.Threading.Tasks;
using Pulumi;
using Pulumi.Docker;

class Program
{
    static Task<int> Main() => Deployment.RunAsync(async () => {
        // Fetch the Docker Hub auth info from config.
        var config = new Pulumi.Config();
        var username = config.Require("dockerUsername");
        var password = config.RequireSecret("dockerPassword");

        // Populate the registry info (creds and endpoint).
        var imageName = $"{username}/myapp";
        var registryInfo = new ImageRegistry
        {
            Server = "docker.io",
            Username = username,
            Password = password,
        };

        // Build and publish the app image.
        var image = new Image("my-image", new ImageArgs
        {
            ImageName = username + "/myapp",
            Build = new DockerBuild { Context = "app" },
            Registry = registryInfo,
        });

        // Export the resulting image name.
        return new Dictionary<string, object>
        {
            { "imageName", image.ImageName },
        };
    });
}
