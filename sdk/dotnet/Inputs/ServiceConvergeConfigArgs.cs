// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Docker.Inputs
{

    public sealed class ServiceConvergeConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The interval to check if the desired state is reached `(ms|s)`. Defaults to `7s`.
        /// </summary>
        [Input("delay")]
        public Input<string>? Delay { get; set; }

        /// <summary>
        /// The timeout of the service to reach the desired state `(s|m)`. Defaults to `3m`
        /// </summary>
        [Input("timeout")]
        public Input<string>? Timeout { get; set; }

        public ServiceConvergeConfigArgs()
        {
        }
    }
}
