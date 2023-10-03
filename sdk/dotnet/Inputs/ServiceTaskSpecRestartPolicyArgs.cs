// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Docker.Inputs
{

    public sealed class ServiceTaskSpecRestartPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Condition for restart
        /// </summary>
        [Input("condition")]
        public Input<string>? Condition { get; set; }

        /// <summary>
        /// The interval to check if the desired state is reached `(ms|s)`. Defaults to `7s`.
        /// </summary>
        [Input("delay")]
        public Input<string>? Delay { get; set; }

        /// <summary>
        /// Maximum attempts to restart a given container before giving up (default value is `0`, which is ignored)
        /// </summary>
        [Input("maxAttempts")]
        public Input<int>? MaxAttempts { get; set; }

        /// <summary>
        /// The time window used to evaluate the restart policy (default value is `0`, which is unbounded) (ms|s|m|h)
        /// </summary>
        [Input("window")]
        public Input<string>? Window { get; set; }

        public ServiceTaskSpecRestartPolicyArgs()
        {
        }
        public static new ServiceTaskSpecRestartPolicyArgs Empty => new ServiceTaskSpecRestartPolicyArgs();
    }
}
