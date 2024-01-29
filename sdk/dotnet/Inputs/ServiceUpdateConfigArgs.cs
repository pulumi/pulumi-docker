// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Docker.Inputs
{

    public sealed class ServiceUpdateConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("delay")]
        public Input<string>? Delay { get; set; }

        [Input("failureAction")]
        public Input<string>? FailureAction { get; set; }

        [Input("maxFailureRatio")]
        public Input<string>? MaxFailureRatio { get; set; }

        [Input("monitor")]
        public Input<string>? Monitor { get; set; }

        [Input("order")]
        public Input<string>? Order { get; set; }

        [Input("parallelism")]
        public Input<int>? Parallelism { get; set; }

        public ServiceUpdateConfigArgs()
        {
        }
        public static new ServiceUpdateConfigArgs Empty => new ServiceUpdateConfigArgs();
    }
}
