using System;
using System.Collections.Generic;
using System.Text;
using System.Threading.Tasks;
using Newtonsoft.Json;
using Pulumi;
using Pulumi.DigitalOcean;
using Pulumi.Docker;

class Program
{
    static Task<int> Main() => Deployment.RunAsync(async () => {
        // Lookup the DigitalOcean Container Registry.
        var registry = Output.Create(GetContainerRegistry.InvokeAsync(new GetContainerRegistryArgs()
        {
            Name = "development-pulumi-provider"
        }));
        var endpoint = registry.Apply(x => x.Endpoint);
        var serverUrl = registry.Apply(x => x.ServerUrl);
        var registryName = registry.Apply(x => x.Name);

        // Get registry info (creds and endpoint).
        var imageName = Output.Format($"{endpoint}/myapp");
        var registryCreds = new ContainerRegistryDockerCredentials("my-reg-creds",
            new ContainerRegistryDockerCredentialsArgs
            {
                RegistryName = registryName,
                Write = true,
            });
        var registryInfo = Output.All(
            registryCreds.DockerCredentials, serverUrl).
            Apply(args =>
            {
                var authJson = args[0];
                var serverUrl = args[1];
                dynamic auths = JsonConvert.DeserializeObject(authJson);
                var authToken = auths["auths"][serverUrl]["auth"].ToString();
                
                var decodedBytes = Convert.FromBase64String(authToken);
                var decoded = Encoding.UTF8.GetString(decodedBytes);
                
                var parts = decoded.Split(':');
                if (parts.Length != 2)
                {
                    throw new Exception("Invalid credentials");
                }

                return new ImageRegistry
                {
                    Server = serverUrl,
                    Password = parts[0],
                    Username = parts[1],
                };
            });

        // Build and publish the app image.
        var image = new Image("my-image", new ImageArgs
        {
            Build = new DockerBuild { Context = "app" },
            ImageName = imageName,
            Registry = registryInfo,
        });

        // Export the resulting base name in addition to the specific version pushed.
        return new Dictionary<string, object>
        {
            { "baseImageName", image.BaseImageName },
            { "fullImageName", image.ImageName },
        };
    });
}
