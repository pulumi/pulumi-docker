using System;
using System.Collections.Generic;
using System.Text;
using System.Threading.Tasks;
using Newtonsoft.Json;
using Pulumi;
using Pulumi.DigitalOcean;
using Pulumi.Docker;
using Pulumi.Docker.Inputs;

class Program
{
    static Task<int> Main() => Deployment.RunAsync(() => {
        // Create a private DigitalOcean Container Registry.
        var registry = new ContainerRegistry("my-reg", new ContainerRegistryArgs
        {
            SubscriptionTierSlug = "starter"
        });

        // Get registry info (creds and endpoint).
        var imageName = Output.Format($"{registry.Endpoint}/myapp");
        var registryCreds = new ContainerRegistryDockerCredentials("my-reg-creds",
            new ContainerRegistryDockerCredentialsArgs
            {
                RegistryName = registry.Name,
                Write = true,
            });
        var registryInfo = Output.All(
            registryCreds.DockerCredentials, registry.ServerUrl).
            Apply(args =>
            {
                var authJson = args[0];
                var serverUrl = args[1];
                dynamic auths = JsonConvert.DeserializeObject(authJson);
                var authToken = auths["auths"][serverUrl]["auth"];
                var decoded = Encoding.ASCII.GetString(authToken);

                var parts = decoded.Split(':');
                if (parts.Length != 2)
                {
                    throw new Exception("Invalid credentials");
                }

                return new RegistryArgs
                {
                    Server = serverUrl,
                    Username = parts[0],
                    Password = parts[1],
                };
            });

        // Build and publish the app image.
        var image = new Image("my-image", new ImageArgs
        {
            Build = new DockerBuildArgs { Context = "app" },
            ImageName = imageName,
            Registry = registryInfo,
        });

        // Export the resulting image name.
        // Export the resulting base name in addition to the specific version pushed.
        return new Dictionary<string, object?>
        {
            { "baseImageName", image.BaseImageName },
            { "fullImageName", image.ImageName },
        };
    });
}
