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
    public sealed class ContainerMountTmpfsOptions
    {
        /// <summary>
        /// The permission mode for the tmpfs mount in an integer.
        /// </summary>
        public readonly int? Mode;
        /// <summary>
        /// The size for the tmpfs mount in bytes.
        /// </summary>
        public readonly int? SizeBytes;

        [OutputConstructor]
        private ContainerMountTmpfsOptions(
            int? mode,

            int? sizeBytes)
        {
            Mode = mode;
            SizeBytes = sizeBytes;
        }
    }
}
