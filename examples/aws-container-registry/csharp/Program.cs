using System;
using System.Collections.Generic;
using System.Text;
using System.Threading.Tasks;
using Pulumi;
using Pulumi.Aws.Ecr;
using Pulumi.Docker;
using Pulumi.Docker.Inputs;

class Program
{
    static Task<int> Main() => Deployment.RunAsync(async () => {
        // Create a private ECR registry.
        var repo = new Repository("my-repo", new RepositoryArgs{
            ForceDelete = true,
        });


        // Get registry info (creds and endpoint) so we can build/publish to it.
        var imageName = repo.RepositoryUrl;
        var registryInfo = repo.RegistryId.Apply(async (id) =>
        {
            var creds = await GetCredentials.InvokeAsync(new GetCredentialsArgs { RegistryId = id });
            var decodedData = Convert.FromBase64String(creds.AuthorizationToken);
            var decoded = ASCIIEncoding.ASCII.GetString(decodedData);

            var parts = decoded.Split(':');
            if (parts.Length != 2)
            {
                throw new Exception("Invalid credentials");
            }

            return new Pulumi.Docker.Inputs.RegistryArgs
            {
                Server = creds.ProxyEndpoint,
                Username = parts[0],
                Password = parts[1],
            };
        });

        // Build and publish the app image.
        var image = new Image("my-image", new ImageArgs
        {
            Build = new Pulumi.Docker.Inputs.DockerBuildArgs { Context = "app" },
            ImageName = imageName,
            Registry = registryInfo,
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
