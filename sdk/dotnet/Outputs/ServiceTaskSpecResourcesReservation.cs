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
    public sealed class ServiceTaskSpecResourcesReservation
    {
        /// <summary>
        /// User-defined resources can be either Integer resources (e.g, `SSD=3`) or String resources (e.g, GPU=UUID1)
        /// </summary>
        public readonly Outputs.ServiceTaskSpecResourcesReservationGenericResources? GenericResources;
        /// <summary>
        /// The amounf of memory in bytes the container allocates
        /// </summary>
        public readonly int? MemoryBytes;
        /// <summary>
        /// CPU shares in units of 1/1e9 (or 10^-9) of the CPU. Should be at least `1000000`
        /// </summary>
        public readonly int? NanoCpus;

        [OutputConstructor]
        private ServiceTaskSpecResourcesReservation(
            Outputs.ServiceTaskSpecResourcesReservationGenericResources? genericResources,

            int? memoryBytes,

            int? nanoCpus)
        {
            GenericResources = genericResources;
            MemoryBytes = memoryBytes;
            NanoCpus = nanoCpus;
        }
    }
}
