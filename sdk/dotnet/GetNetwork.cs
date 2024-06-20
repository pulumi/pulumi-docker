// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Docker
{
    public static class GetNetwork
    {
        /// <summary>
        /// `docker.Network` provides details about a specific Docker Network.
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Docker = Pulumi.Docker;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var main = Docker.GetNetwork.Invoke(new()
        ///     {
        ///         Name = "main",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetNetworkResult> InvokeAsync(GetNetworkArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNetworkResult>("docker:index/getNetwork:getNetwork", args ?? new GetNetworkArgs(), options.WithDefaults());

        /// <summary>
        /// `docker.Network` provides details about a specific Docker Network.
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Docker = Pulumi.Docker;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var main = Docker.GetNetwork.Invoke(new()
        ///     {
        ///         Name = "main",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetNetworkResult> Invoke(GetNetworkInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNetworkResult>("docker:index/getNetwork:getNetwork", args ?? new GetNetworkInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNetworkArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Docker network.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetNetworkArgs()
        {
        }
        public static new GetNetworkArgs Empty => new GetNetworkArgs();
    }

    public sealed class GetNetworkInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Docker network.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetNetworkInvokeArgs()
        {
        }
        public static new GetNetworkInvokeArgs Empty => new GetNetworkInvokeArgs();
    }


    [OutputType]
    public sealed class GetNetworkResult
    {
        /// <summary>
        /// The driver of the Docker network. Possible values are `bridge`, `host`, `overlay`, `macvlan`. See [network docs](https://docs.docker.com/network/#network-drivers) for more details.
        /// </summary>
        public readonly string Driver;
        /// <summary>
        /// The ID of this resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// If `true`, the network is internal.
        /// </summary>
        public readonly bool Internal;
        /// <summary>
        /// The IPAM configuration options
        /// </summary>
        public readonly ImmutableArray<Outputs.GetNetworkIpamConfigResult> IpamConfigs;
        /// <summary>
        /// The name of the Docker network.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Only available with bridge networks. See [bridge options docs](https://docs.docker.com/engine/reference/commandline/network_create/#bridge-driver-options) for more details.
        /// </summary>
        public readonly ImmutableDictionary<string, object> Options;
        /// <summary>
        /// Scope of the network. One of `swarm`, `global`, or `local`.
        /// </summary>
        public readonly string Scope;

        [OutputConstructor]
        private GetNetworkResult(
            string driver,

            string id,

            bool @internal,

            ImmutableArray<Outputs.GetNetworkIpamConfigResult> ipamConfigs,

            string name,

            ImmutableDictionary<string, object> options,

            string scope)
        {
            Driver = driver;
            Id = id;
            Internal = @internal;
            IpamConfigs = ipamConfigs;
            Name = name;
            Options = options;
            Scope = scope;
        }
    }
}
