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
    public sealed class ServiceTaskSpecContainerSpecMountVolumeOptionsLabel
    {
        /// <summary>
        /// Name of the label
        /// </summary>
        public readonly string Label;
        /// <summary>
        /// Value of the label
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private ServiceTaskSpecContainerSpecMountVolumeOptionsLabel(
            string label,

            string value)
        {
            Label = label;
            Value = value;
        }
    }
}
