// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Docker.Inputs
{

    public sealed class ServiceTaskSpecContainerSpecMountArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional configuration for the `bind` type.
        /// </summary>
        [Input("bindOptions")]
        public Input<Inputs.ServiceTaskSpecContainerSpecMountBindOptionsArgs>? BindOptions { get; set; }

        /// <summary>
        /// Mount the container's root filesystem as read only.
        /// </summary>
        [Input("readOnly")]
        public Input<bool>? ReadOnly { get; set; }

        /// <summary>
        /// The mount source (e.g., a volume name, a host path)
        /// </summary>
        [Input("source")]
        public Input<string>? Source { get; set; }

        /// <summary>
        /// The container path.
        /// </summary>
        [Input("target", required: true)]
        public Input<string> Target { get; set; } = null!;

        /// <summary>
        /// Optional configuration for the `tmpf` type.
        /// </summary>
        [Input("tmpfsOptions")]
        public Input<Inputs.ServiceTaskSpecContainerSpecMountTmpfsOptionsArgs>? TmpfsOptions { get; set; }

        /// <summary>
        /// SELinux type label
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// Optional configuration for the `volume` type.
        /// </summary>
        [Input("volumeOptions")]
        public Input<Inputs.ServiceTaskSpecContainerSpecMountVolumeOptionsArgs>? VolumeOptions { get; set; }

        public ServiceTaskSpecContainerSpecMountArgs()
        {
        }
    }
}
