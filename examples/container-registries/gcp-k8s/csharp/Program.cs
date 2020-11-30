using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Pulumi;
using Pulumi.Docker;
using Pulumi.Gcp.Container;
using Pulumi.Kubernetes.Types.Inputs.Core.V1;
using Pulumi.Kubernetes.Types.Inputs.Apps.V1;
using Pulumi.Kubernetes.Types.Inputs.Meta.V1;

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

        // Create a load balanced Kubernetes service using this image, and export its IP.
        var appLabels = new InputMap<string>
        {
            { "app", "myapp" }
        };
        var appDep = new Pulumi.Kubernetes.Apps.V1.Deployment("app-dep", new DeploymentArgs
        {
            Spec = new DeploymentSpecArgs
            {
                Selector = new LabelSelectorArgs
                {
                    MatchLabels = appLabels
                },
                Replicas = 3,
                Template = new PodTemplateSpecArgs
                {
                    Metadata = new ObjectMetaArgs
                    {
                        Labels = appLabels,
                    },
                    Spec = new PodSpecArgs
                    {
                        Containers =
                        {
                            new Pulumi.Kubernetes.Types.Inputs.Core.V1.ContainerArgs
                            {
                                Name = "myapp",
                                Image = image.ImageName,
                            }
                        }
                    }
                }
            }
        });
        var appSvc = new Pulumi.Kubernetes.Core.V1.Service("app-svc", new Pulumi.Kubernetes.Types.Inputs.Core.V1.ServiceArgs
        {
            Metadata = new ObjectMetaArgs
            {
                Labels = appLabels
            },
            Spec = new ServiceSpecArgs
            {
                Type = "LoadBalancer",
                Ports =
                {
                    new ServicePortArgs
                    {
                        Port = 80,
                        TargetPort = 80
                    }
                },
                Selector = appLabels
            }
        });

        // Export the resulting base name in addition to the specific version pushed.
        return new Dictionary<string, object>
        {
            { "baseImageName", image.BaseImageName },
            { "fullImageName", image.ImageName },
            { "appIp", appSvc.Status.Apply(status => status.LoadBalancer.Ingress[0].Ip) },
        };
    });
}
