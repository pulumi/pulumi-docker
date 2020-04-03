// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Docker.Inputs
{

    public sealed class ServiceTaskSpecResourcesLimitsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// User-defined resources can be either Integer resources (e.g, SSD=3) or String resources (e.g, GPU=UUID1)
        /// </summary>
        [Input("genericResources")]
        public Input<Inputs.ServiceTaskSpecResourcesLimitsGenericResourcesArgs>? GenericResources { get; set; }

        /// <summary>
        /// The amount of memory in bytes the container allocates
        /// </summary>
        [Input("memoryBytes")]
        public Input<int>? MemoryBytes { get; set; }

        /// <summary>
        /// CPU shares in units of 1/1e9 (or 10^-9) of the CPU. Should be at least 1000000
        /// </summary>
        [Input("nanoCpus")]
        public Input<int>? NanoCpus { get; set; }

        public ServiceTaskSpecResourcesLimitsArgs()
        {
        }
    }
}