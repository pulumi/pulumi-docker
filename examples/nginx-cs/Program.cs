using System.Collections.Generic;
using System.Threading.Tasks;

using Pulumi;
using Pulumi.Docker;
using Pulumi.Docker.Inputs;

class Program
{
    static Task<int> Main()
    {
        return Deployment.RunAsync(() =>
        {
            // Get a reference to the remote image "nginx:1.15.6". Without specifying the repository, the Docker provider will
            // try to download it from the public Docker Hub.
            var remoteImage = new RemoteImage("nginx-image", new RemoteImageArgs
            {
                Name = "nginx",
                KeepLocally = true, // don't delete the image from the local cache when deleting this resource
            });

            // Launch a container using the nginx image we just downloaded.
            var container = new Container("nginx", new ContainerArgs
            {
                Image = remoteImage.Latest,
                Ports =
                {
                    new ContainerPortArgs
                    {
                        Internal = 80,
                        // external: defaults to an open ephemeral port
                        // protocol: defaults to TCP
                        // ip: defaults to 0.0.0.0
                    },
                },
            });

            var endpoints = container.Ports.Apply(ports => $"{ports![0].Ip}:{ports![0].External}");
            return new Dictionary<string, object?>
            {
                // Since the container is auto-named, export the name.
                { "name", container.Name },
                // Since the provider picked a random ephemeral port for this container, export the endpoint.
                { "endpoints", endpoints },
            };
        });
    }
}
