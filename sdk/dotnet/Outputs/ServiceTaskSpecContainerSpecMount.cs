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
    public sealed class ServiceTaskSpecContainerSpecMount
    {
        /// <summary>
        /// Optional configuration for the bind type
        /// </summary>
        public readonly Outputs.ServiceTaskSpecContainerSpecMountBindOptions? BindOptions;
        /// <summary>
        /// Whether the mount should be read-only
        /// </summary>
        public readonly bool? ReadOnly;
        /// <summary>
        /// Mount source (e.g. a volume name, a host path)
        /// </summary>
        public readonly string? Source;
        /// <summary>
        /// Container path
        /// </summary>
        public readonly string Target;
        /// <summary>
        /// Optional configuration for the tmpfs type
        /// </summary>
        public readonly Outputs.ServiceTaskSpecContainerSpecMountTmpfsOptions? TmpfsOptions;
        /// <summary>
        /// The mount type
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Optional configuration for the volume type
        /// </summary>
        public readonly Outputs.ServiceTaskSpecContainerSpecMountVolumeOptions? VolumeOptions;

        [OutputConstructor]
        private ServiceTaskSpecContainerSpecMount(
            Outputs.ServiceTaskSpecContainerSpecMountBindOptions? bindOptions,

            bool? readOnly,

            string? source,

            string target,

            Outputs.ServiceTaskSpecContainerSpecMountTmpfsOptions? tmpfsOptions,

            string type,

            Outputs.ServiceTaskSpecContainerSpecMountVolumeOptions? volumeOptions)
        {
            BindOptions = bindOptions;
            ReadOnly = readOnly;
            Source = source;
            Target = target;
            TmpfsOptions = tmpfsOptions;
            Type = type;
            VolumeOptions = volumeOptions;
        }
    }
}
