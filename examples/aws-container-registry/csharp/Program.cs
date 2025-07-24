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
            var creds = GetAuthorizationToken.Invoke(new GetAuthorizationTokenInvokeArgs { RegistryId =id });
            return new Pulumi.Docker.Inputs.RegistryArgs
            {
                Server = creds.Apply(credentials => credentials.ProxyEndpoint),
                Username = creds.Apply(credentials => credentials.UserName),
                Password = creds.Apply(credentials => credentials.Password),
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
