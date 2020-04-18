// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Docker.Inputs
{

    public sealed class ServiceTaskSpecRestartPolicyArgs : Pulumi.ResourceArgs
    {
        [Input("condition")]
        public Input<string>? Condition { get; set; }

        /// <summary>
        /// Delay between updates `(ns|us|ms|s|m|h)`, e.g. `5s`.
        /// all tasks are up when a service is created, or to check if all tasks are successfully updated on an update. Default: `7s`.
        /// </summary>
        [Input("delay")]
        public Input<string>? Delay { get; set; }

        [Input("maxAttempts")]
        public Input<int>? MaxAttempts { get; set; }

        [Input("window")]
        public Input<string>? Window { get; set; }

        public ServiceTaskSpecRestartPolicyArgs()
        {
        }
    }
}
