using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Pulumi;
using Pulumi.Azure.Authorization;
using Pulumi.Azure.Core;
using Pulumi.Azure.ContainerService;
using AAD = Pulumi.AzureAD;
using Docker = Pulumi.Docker;
using Pulumi.Random;

class Program
{
    static Task<int> Main() => Deployment.RunAsync(async () => {
        // Conditionalize the auth mechanism.
        var config = new Config();
        var useServicePrincipalAuth = config.GetBoolean("useServicePrincipalAuth") ?? false;

        // Create a private ACR registry.
        var rg = new ResourceGroup("myrg");
        var registry = new Registry("myregistry", new RegistryArgs
        {
            ResourceGroupName = rg.Name,
            AdminEnabled = !useServicePrincipalAuth,
            Sku = "basic"
        });

        // Get registry info (creds and endpoint) so we can build/publish to it.
        var imageName = Output.Format($"{registry.LoginServer}/myapp");
        Docker.ImageRegistry registryInfo;
        if (useServicePrincipalAuth)
        {
            var sp = new AAD.ServicePrincipal("mysp", new AAD.ServicePrincipalArgs
            {
                ApplicationId = new AAD.Application("myspapp").ApplicationId,
            });
            var spPassword = new AAD.ServicePrincipalPassword("mysp-pass", new AAD.ServicePrincipalPasswordArgs
            {
                ServicePrincipalId = sp.Id,
                Value = new RandomPassword("mypass",
                    new RandomPasswordArgs
                    {
                        Length = 32,
                        },
                    new CustomResourceOptions { AdditionalSecretOutputs = { "result" } }
                ).Result,
                EndDateRelative = "8760h",
            });
            var spAuth = new Assignment("myauth", new AssignmentArgs
            {
                Scope = registry.Id,
                RoleDefinitionName = "acrpush",
                PrincipalId = sp.Id,
            });
            registryInfo = new Docker.ImageRegistry
            {
                Server = registry.LoginServer,
                Username = sp.ApplicationId,
                Password = spAuth.Id.Apply(_ => spPassword.Value),
            };
        }
        else
        {
            registryInfo = new Docker.ImageRegistry
            {
                Server = registry.LoginServer,
                Username = registry.AdminUsername,
                Password = registry.AdminPassword,
            };
        }

        // Build and publish the app image.
        var image = new Docker.Image("my-image", new Docker.ImageArgs
        {
            Build = new Docker.DockerBuild { Context = "app" },
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
