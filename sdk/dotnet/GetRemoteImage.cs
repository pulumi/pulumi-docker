// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Docker
{
    public static class GetRemoteImage
    {
        /// <summary>
        /// `docker.RemoteImage` provides details about a specific Docker Image which need to be presend on the Docker Host
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Docker = Pulumi.Docker;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var latest = Output.Create(Docker.GetRemoteImage.InvokeAsync(new Docker.GetRemoteImageArgs
        ///         {
        ///             Name = "nginx",
        ///         }));
        ///         var specific = Output.Create(Docker.GetRemoteImage.InvokeAsync(new Docker.GetRemoteImageArgs
        ///         {
        ///             Name = "nginx:1.17.6",
        ///         }));
        ///         var digest = Output.Create(Docker.GetRemoteImage.InvokeAsync(new Docker.GetRemoteImageArgs
        ///         {
        ///             Name = "nginx@sha256:36b74457bccb56fbf8b05f79c85569501b721d4db813b684391d63e02287c0b2",
        ///         }));
        ///         var tagAndDigest = Output.Create(Docker.GetRemoteImage.InvokeAsync(new Docker.GetRemoteImageArgs
        ///         {
        ///             Name = "nginx:1.19.1@sha256:36b74457bccb56fbf8b05f79c85569501b721d4db813b684391d63e02287c0b2",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetRemoteImageResult> InvokeAsync(GetRemoteImageArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRemoteImageResult>("docker:index/getRemoteImage:getRemoteImage", args ?? new GetRemoteImageArgs(), options.WithDefaults());

        /// <summary>
        /// `docker.RemoteImage` provides details about a specific Docker Image which need to be presend on the Docker Host
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Docker = Pulumi.Docker;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var latest = Output.Create(Docker.GetRemoteImage.InvokeAsync(new Docker.GetRemoteImageArgs
        ///         {
        ///             Name = "nginx",
        ///         }));
        ///         var specific = Output.Create(Docker.GetRemoteImage.InvokeAsync(new Docker.GetRemoteImageArgs
        ///         {
        ///             Name = "nginx:1.17.6",
        ///         }));
        ///         var digest = Output.Create(Docker.GetRemoteImage.InvokeAsync(new Docker.GetRemoteImageArgs
        ///         {
        ///             Name = "nginx@sha256:36b74457bccb56fbf8b05f79c85569501b721d4db813b684391d63e02287c0b2",
        ///         }));
        ///         var tagAndDigest = Output.Create(Docker.GetRemoteImage.InvokeAsync(new Docker.GetRemoteImageArgs
        ///         {
        ///             Name = "nginx:1.19.1@sha256:36b74457bccb56fbf8b05f79c85569501b721d4db813b684391d63e02287c0b2",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetRemoteImageResult> Invoke(GetRemoteImageInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetRemoteImageResult>("docker:index/getRemoteImage:getRemoteImage", args ?? new GetRemoteImageInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRemoteImageArgs : Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetRemoteImageArgs()
        {
        }
    }

    public sealed class GetRemoteImageInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetRemoteImageInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRemoteImageResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string RepoDigest;

        [OutputConstructor]
        private GetRemoteImageResult(
            string id,

            string name,

            string repoDigest)
        {
            Id = id;
            Name = name;
            RepoDigest = repoDigest;
        }
    }
}
