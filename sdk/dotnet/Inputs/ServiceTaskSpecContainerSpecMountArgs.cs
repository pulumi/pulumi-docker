// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Docker.Inputs
{

    public sealed class ServiceTaskSpecContainerSpecMountArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional configuration for the bind type
        /// </summary>
        [Input("bindOptions")]
        public Input<Inputs.ServiceTaskSpecContainerSpecMountBindOptionsArgs>? BindOptions { get; set; }

        /// <summary>
        /// Whether the mount should be read-only
        /// </summary>
        [Input("readOnly")]
        public Input<bool>? ReadOnly { get; set; }

        /// <summary>
        /// Mount source (e.g. a volume name, a host path)
        /// </summary>
        [Input("source")]
        public Input<string>? Source { get; set; }

        /// <summary>
        /// Container path
        /// </summary>
        [Input("target", required: true)]
        public Input<string> Target { get; set; } = null!;

        /// <summary>
        /// Optional configuration for the tmpfs type
        /// </summary>
        [Input("tmpfsOptions")]
        public Input<Inputs.ServiceTaskSpecContainerSpecMountTmpfsOptionsArgs>? TmpfsOptions { get; set; }

        /// <summary>
        /// The mount type
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// Optional configuration for the volume type
        /// </summary>
        [Input("volumeOptions")]
        public Input<Inputs.ServiceTaskSpecContainerSpecMountVolumeOptionsArgs>? VolumeOptions { get; set; }

        public ServiceTaskSpecContainerSpecMountArgs()
        {
        }
        public static new ServiceTaskSpecContainerSpecMountArgs Empty => new ServiceTaskSpecContainerSpecMountArgs();
    }
}
