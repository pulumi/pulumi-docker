// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Docker.Inputs
{

    public sealed class ContainerMountTmpfsOptionsGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("mode")]
        public Input<int>? Mode { get; set; }

        [Input("sizeBytes")]
        public Input<int>? SizeBytes { get; set; }

        public ContainerMountTmpfsOptionsGetArgs()
        {
        }
        public static new ContainerMountTmpfsOptionsGetArgs Empty => new ContainerMountTmpfsOptionsGetArgs();
    }
}
