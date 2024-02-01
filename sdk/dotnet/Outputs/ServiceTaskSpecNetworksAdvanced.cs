// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Docker.Outputs
{

    [OutputType]
    public sealed class ServiceTaskSpecNetworksAdvanced
    {
        /// <summary>
        /// The network aliases of the container in the specific network.
        /// </summary>
        public readonly ImmutableArray<string> Aliases;
        /// <summary>
        /// An array of driver options for the network, e.g. `opts1=value`
        /// </summary>
        public readonly ImmutableArray<string> DriverOpts;
        /// <summary>
        /// The name/id of the network.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private ServiceTaskSpecNetworksAdvanced(
            ImmutableArray<string> aliases,

            ImmutableArray<string> driverOpts,

            string name)
        {
            Aliases = aliases;
            DriverOpts = driverOpts;
            Name = name;
        }
    }
}
