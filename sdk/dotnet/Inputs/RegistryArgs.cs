// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Docker.Inputs
{

    /// <summary>
    /// Describes a Docker container registry
    /// </summary>
    public sealed class RegistryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The password to authenticate to the registry
        /// </summary>
        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        /// <summary>
        /// The URL of the Docker registry server
        /// </summary>
        [Input("serverURL", required: true)]
        public Input<string> ServerURL { get; set; } = null!;

        /// <summary>
        /// The username to authenticate to the registry
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public RegistryArgs()
        {
        }
        public static new RegistryArgs Empty => new RegistryArgs();
    }
}