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
    public sealed class ContainerMountVolumeOptions
    {
        /// <summary>
        /// Name of the driver to use to create the volume.
        /// </summary>
        public readonly string? DriverName;
        /// <summary>
        /// key/value map of driver specific options.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? DriverOptions;
        /// <summary>
        /// User-defined key/value metadata.
        /// </summary>
        public readonly ImmutableArray<Outputs.ContainerMountVolumeOptionsLabel> Labels;
        /// <summary>
        /// Populate volume with data from the target.
        /// </summary>
        public readonly bool? NoCopy;

        [OutputConstructor]
        private ContainerMountVolumeOptions(
            string? driverName,

            ImmutableDictionary<string, string>? driverOptions,

            ImmutableArray<Outputs.ContainerMountVolumeOptionsLabel> labels,

            bool? noCopy)
        {
            DriverName = driverName;
            DriverOptions = driverOptions;
            Labels = labels;
            NoCopy = noCopy;
        }
    }
}
