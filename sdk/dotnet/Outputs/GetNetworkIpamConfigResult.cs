// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Docker.Outputs
{

    [OutputType]
    public sealed class GetNetworkIpamConfigResult
    {
        /// <summary>
        /// Auxiliary IPv4 or IPv6 addresses used by Network driver
        /// </summary>
        public readonly ImmutableDictionary<string, string>? AuxAddress;
        /// <summary>
        /// The IP address of the gateway
        /// </summary>
        public readonly string? Gateway;
        /// <summary>
        /// The ip range in CIDR form
        /// </summary>
        public readonly string? IpRange;
        /// <summary>
        /// The subnet in CIDR form
        /// </summary>
        public readonly string? Subnet;

        [OutputConstructor]
        private GetNetworkIpamConfigResult(
            ImmutableDictionary<string, string>? auxAddress,

            string? gateway,

            string? ipRange,

            string? subnet)
        {
            AuxAddress = auxAddress;
            Gateway = gateway;
            IpRange = ipRange;
            Subnet = subnet;
        }
    }
}
