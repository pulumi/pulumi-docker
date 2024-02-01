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
    public sealed class ServiceEndpointSpec
    {
        /// <summary>
        /// The mode of resolution to use for internal load balancing between tasks
        /// </summary>
        public readonly string? Mode;
        /// <summary>
        /// List of exposed ports that this service is accessible on from the outside. Ports can only be provided if 'vip' resolution mode is used
        /// </summary>
        public readonly ImmutableArray<Outputs.ServiceEndpointSpecPort> Ports;

        [OutputConstructor]
        private ServiceEndpointSpec(
            string? mode,

            ImmutableArray<Outputs.ServiceEndpointSpecPort> ports)
        {
            Mode = mode;
            Ports = ports;
        }
    }
}
