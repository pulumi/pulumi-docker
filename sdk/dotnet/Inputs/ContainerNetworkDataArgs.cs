// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Docker.Inputs
{

    public sealed class ContainerNetworkDataArgs : Pulumi.ResourceArgs
    {
        [Input("gateway")]
        public Input<string>? Gateway { get; set; }

        [Input("globalIpv6Address")]
        public Input<string>? GlobalIpv6Address { get; set; }

        [Input("globalIpv6PrefixLength")]
        public Input<int>? GlobalIpv6PrefixLength { get; set; }

        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        [Input("ipPrefixLength")]
        public Input<int>? IpPrefixLength { get; set; }

        [Input("ipv6Gateway")]
        public Input<string>? Ipv6Gateway { get; set; }

        [Input("networkName")]
        public Input<string>? NetworkName { get; set; }

        public ContainerNetworkDataArgs()
        {
        }
    }
}
