// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Docker.Inputs
{

    public sealed class ContainerMountArgs : Pulumi.ResourceArgs
    {
        [Input("bindOptions")]
        public Input<Inputs.ContainerMountBindOptionsArgs>? BindOptions { get; set; }

        [Input("readOnly")]
        public Input<bool>? ReadOnly { get; set; }

        [Input("source")]
        public Input<string>? Source { get; set; }

        [Input("target", required: true)]
        public Input<string> Target { get; set; } = null!;

        [Input("tmpfsOptions")]
        public Input<Inputs.ContainerMountTmpfsOptionsArgs>? TmpfsOptions { get; set; }

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        [Input("volumeOptions")]
        public Input<Inputs.ContainerMountVolumeOptionsArgs>? VolumeOptions { get; set; }

        public ContainerMountArgs()
        {
        }
    }
}
