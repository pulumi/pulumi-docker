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
    public sealed class ServiceTaskSpecResourcesReservationGenericResources
    {
        /// <summary>
        /// The Integer resources
        /// </summary>
        public readonly ImmutableArray<string> DiscreteResourcesSpecs;
        /// <summary>
        /// The String resources
        /// </summary>
        public readonly ImmutableArray<string> NamedResourcesSpecs;

        [OutputConstructor]
        private ServiceTaskSpecResourcesReservationGenericResources(
            ImmutableArray<string> discreteResourcesSpecs,

            ImmutableArray<string> namedResourcesSpecs)
        {
            DiscreteResourcesSpecs = discreteResourcesSpecs;
            NamedResourcesSpecs = namedResourcesSpecs;
        }
    }
}
