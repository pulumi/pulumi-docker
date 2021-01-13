// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Docker.Inputs
{

    public sealed class RegistryImageBuildAuthConfigGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// the auth token
        /// </summary>
        [Input("auth")]
        public Input<string>? Auth { get; set; }

        /// <summary>
        /// the user emal
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        /// <summary>
        /// hostname of the registry
        /// </summary>
        [Input("hostName", required: true)]
        public Input<string> HostName { get; set; } = null!;

        /// <summary>
        /// the identity token
        /// </summary>
        [Input("identityToken")]
        public Input<string>? IdentityToken { get; set; }

        /// <summary>
        /// the registry password
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// the registry token
        /// </summary>
        [Input("registryToken")]
        public Input<string>? RegistryToken { get; set; }

        /// <summary>
        /// the server address
        /// </summary>
        [Input("serverAddress")]
        public Input<string>? ServerAddress { get; set; }

        /// <summary>
        /// the registry user name
        /// </summary>
        [Input("userName")]
        public Input<string>? UserName { get; set; }

        public RegistryImageBuildAuthConfigGetArgs()
        {
        }
    }
}