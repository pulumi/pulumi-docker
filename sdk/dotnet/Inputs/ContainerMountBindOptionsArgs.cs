// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Docker.Inputs
{

    public sealed class ContainerMountBindOptionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A propagation mode with the value.
        /// </summary>
        [Input("propagation")]
        public Input<string>? Propagation { get; set; }

        public ContainerMountBindOptionsArgs()
        {
        }
        public static new ContainerMountBindOptionsArgs Empty => new ContainerMountBindOptionsArgs();
    }
}
