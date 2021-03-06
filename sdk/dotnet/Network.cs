// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Docker
{
    /// <summary>
    /// Manages a Docker Network. This can be used alongside
    /// [docker\_container](https://www.terraform.io/docs/providers/docker/r/container.html)
    /// to create virtual networks within the docker environment.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Docker = Pulumi.Docker;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         // Create a new docker network
    ///         var privateNetwork = new Docker.Network("privateNetwork", new Docker.NetworkArgs
    ///         {
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Docker networks can be imported using the long id, e.g. for a network with the short id `p73jelnrme5f`
    /// 
    /// ```sh
    ///  $ pulumi import docker:index/network:Network foo $(docker network inspect -f {{.ID}} p73)
    /// ```
    /// </summary>
    [DockerResourceType("docker:index/network:Network")]
    public partial class Network : Pulumi.CustomResource
    {
        /// <summary>
        /// Enable manual container attachment to the network.
        /// Defaults to `false`.
        /// </summary>
        [Output("attachable")]
        public Output<bool?> Attachable { get; private set; } = null!;

        /// <summary>
        /// Requests daemon to check for networks
        /// with same name.
        /// </summary>
        [Output("checkDuplicate")]
        public Output<bool?> CheckDuplicate { get; private set; } = null!;

        /// <summary>
        /// Name of the network driver to use. Defaults to
        /// `bridge` driver.
        /// </summary>
        [Output("driver")]
        public Output<string> Driver { get; private set; } = null!;

        /// <summary>
        /// Create swarm routing-mesh network.
        /// Defaults to `false`.
        /// </summary>
        [Output("ingress")]
        public Output<bool?> Ingress { get; private set; } = null!;

        /// <summary>
        /// Restrict external access to the network.
        /// Defaults to `false`.
        /// </summary>
        [Output("internal")]
        public Output<bool> Internal { get; private set; } = null!;

        /// <summary>
        /// See IPAM config below for
        /// details.
        /// </summary>
        [Output("ipamConfigs")]
        public Output<ImmutableArray<Outputs.NetworkIpamConfig>> IpamConfigs { get; private set; } = null!;

        /// <summary>
        /// Driver used by the custom IP scheme of the
        /// network.
        /// </summary>
        [Output("ipamDriver")]
        public Output<string?> IpamDriver { get; private set; } = null!;

        /// <summary>
        /// Enable IPv6 networking.
        /// Defaults to `false`.
        /// </summary>
        [Output("ipv6")]
        public Output<bool?> Ipv6 { get; private set; } = null!;

        /// <summary>
        /// See Labels below for details.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableArray<Outputs.NetworkLabel>> Labels { get; private set; } = null!;

        /// <summary>
        /// The name of the Docker network.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Network specific options to be used by
        /// the drivers.
        /// </summary>
        [Output("options")]
        public Output<ImmutableDictionary<string, object>> Options { get; private set; } = null!;

        [Output("scope")]
        public Output<string> Scope { get; private set; } = null!;


        /// <summary>
        /// Create a Network resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Network(string name, NetworkArgs? args = null, CustomResourceOptions? options = null)
            : base("docker:index/network:Network", name, args ?? new NetworkArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Network(string name, Input<string> id, NetworkState? state = null, CustomResourceOptions? options = null)
            : base("docker:index/network:Network", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Network resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Network Get(string name, Input<string> id, NetworkState? state = null, CustomResourceOptions? options = null)
        {
            return new Network(name, id, state, options);
        }
    }

    public sealed class NetworkArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable manual container attachment to the network.
        /// Defaults to `false`.
        /// </summary>
        [Input("attachable")]
        public Input<bool>? Attachable { get; set; }

        /// <summary>
        /// Requests daemon to check for networks
        /// with same name.
        /// </summary>
        [Input("checkDuplicate")]
        public Input<bool>? CheckDuplicate { get; set; }

        /// <summary>
        /// Name of the network driver to use. Defaults to
        /// `bridge` driver.
        /// </summary>
        [Input("driver")]
        public Input<string>? Driver { get; set; }

        /// <summary>
        /// Create swarm routing-mesh network.
        /// Defaults to `false`.
        /// </summary>
        [Input("ingress")]
        public Input<bool>? Ingress { get; set; }

        /// <summary>
        /// Restrict external access to the network.
        /// Defaults to `false`.
        /// </summary>
        [Input("internal")]
        public Input<bool>? Internal { get; set; }

        [Input("ipamConfigs")]
        private InputList<Inputs.NetworkIpamConfigArgs>? _ipamConfigs;

        /// <summary>
        /// See IPAM config below for
        /// details.
        /// </summary>
        public InputList<Inputs.NetworkIpamConfigArgs> IpamConfigs
        {
            get => _ipamConfigs ?? (_ipamConfigs = new InputList<Inputs.NetworkIpamConfigArgs>());
            set => _ipamConfigs = value;
        }

        /// <summary>
        /// Driver used by the custom IP scheme of the
        /// network.
        /// </summary>
        [Input("ipamDriver")]
        public Input<string>? IpamDriver { get; set; }

        /// <summary>
        /// Enable IPv6 networking.
        /// Defaults to `false`.
        /// </summary>
        [Input("ipv6")]
        public Input<bool>? Ipv6 { get; set; }

        [Input("labels")]
        private InputList<Inputs.NetworkLabelArgs>? _labels;

        /// <summary>
        /// See Labels below for details.
        /// </summary>
        public InputList<Inputs.NetworkLabelArgs> Labels
        {
            get => _labels ?? (_labels = new InputList<Inputs.NetworkLabelArgs>());
            set => _labels = value;
        }

        /// <summary>
        /// The name of the Docker network.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("options")]
        private InputMap<object>? _options;

        /// <summary>
        /// Network specific options to be used by
        /// the drivers.
        /// </summary>
        public InputMap<object> Options
        {
            get => _options ?? (_options = new InputMap<object>());
            set => _options = value;
        }

        public NetworkArgs()
        {
        }
    }

    public sealed class NetworkState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable manual container attachment to the network.
        /// Defaults to `false`.
        /// </summary>
        [Input("attachable")]
        public Input<bool>? Attachable { get; set; }

        /// <summary>
        /// Requests daemon to check for networks
        /// with same name.
        /// </summary>
        [Input("checkDuplicate")]
        public Input<bool>? CheckDuplicate { get; set; }

        /// <summary>
        /// Name of the network driver to use. Defaults to
        /// `bridge` driver.
        /// </summary>
        [Input("driver")]
        public Input<string>? Driver { get; set; }

        /// <summary>
        /// Create swarm routing-mesh network.
        /// Defaults to `false`.
        /// </summary>
        [Input("ingress")]
        public Input<bool>? Ingress { get; set; }

        /// <summary>
        /// Restrict external access to the network.
        /// Defaults to `false`.
        /// </summary>
        [Input("internal")]
        public Input<bool>? Internal { get; set; }

        [Input("ipamConfigs")]
        private InputList<Inputs.NetworkIpamConfigGetArgs>? _ipamConfigs;

        /// <summary>
        /// See IPAM config below for
        /// details.
        /// </summary>
        public InputList<Inputs.NetworkIpamConfigGetArgs> IpamConfigs
        {
            get => _ipamConfigs ?? (_ipamConfigs = new InputList<Inputs.NetworkIpamConfigGetArgs>());
            set => _ipamConfigs = value;
        }

        /// <summary>
        /// Driver used by the custom IP scheme of the
        /// network.
        /// </summary>
        [Input("ipamDriver")]
        public Input<string>? IpamDriver { get; set; }

        /// <summary>
        /// Enable IPv6 networking.
        /// Defaults to `false`.
        /// </summary>
        [Input("ipv6")]
        public Input<bool>? Ipv6 { get; set; }

        [Input("labels")]
        private InputList<Inputs.NetworkLabelGetArgs>? _labels;

        /// <summary>
        /// See Labels below for details.
        /// </summary>
        public InputList<Inputs.NetworkLabelGetArgs> Labels
        {
            get => _labels ?? (_labels = new InputList<Inputs.NetworkLabelGetArgs>());
            set => _labels = value;
        }

        /// <summary>
        /// The name of the Docker network.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("options")]
        private InputMap<object>? _options;

        /// <summary>
        /// Network specific options to be used by
        /// the drivers.
        /// </summary>
        public InputMap<object> Options
        {
            get => _options ?? (_options = new InputMap<object>());
            set => _options = value;
        }

        [Input("scope")]
        public Input<string>? Scope { get; set; }

        public NetworkState()
        {
        }
    }
}
