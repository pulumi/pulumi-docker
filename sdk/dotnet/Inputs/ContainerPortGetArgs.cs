// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Docker.Inputs
{

    public sealed class ContainerPortGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Port exposed out of the container. If not given a free random port `&gt;= 32768` will be used.
        /// </summary>
        [Input("external")]
        public Input<int>? External { get; set; }

        /// <summary>
        /// Port within the container.
        /// </summary>
        [Input("internal", required: true)]
        public Input<int> Internal { get; set; } = null!;

        /// <summary>
        /// IP address/mask that can access this port. Defaults to `0.0.0.0`.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// Protocol that can be used over this port. Defaults to `tcp`.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        public ContainerPortGetArgs()
        {
        }
    }
}
