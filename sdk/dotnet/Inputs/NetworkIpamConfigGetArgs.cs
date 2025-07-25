// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Docker.Inputs
{

    public sealed class NetworkIpamConfigGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("auxAddress")]
        private InputMap<string>? _auxAddress;

        /// <summary>
        /// Auxiliary IPv4 or IPv6 addresses used by Network driver
        /// </summary>
        public InputMap<string> AuxAddress
        {
            get => _auxAddress ?? (_auxAddress = new InputMap<string>());
            set => _auxAddress = value;
        }

        /// <summary>
        /// The IP address of the gateway
        /// </summary>
        [Input("gateway")]
        public Input<string>? Gateway { get; set; }

        /// <summary>
        /// The ip range in CIDR form
        /// </summary>
        [Input("ipRange")]
        public Input<string>? IpRange { get; set; }

        /// <summary>
        /// The subnet in CIDR form
        /// </summary>
        [Input("subnet")]
        public Input<string>? Subnet { get; set; }

        public NetworkIpamConfigGetArgs()
        {
        }
        public static new NetworkIpamConfigGetArgs Empty => new NetworkIpamConfigGetArgs();
    }
}
