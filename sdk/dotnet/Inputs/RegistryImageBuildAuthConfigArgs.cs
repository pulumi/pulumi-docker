// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Docker.Inputs
{

    public sealed class RegistryImageBuildAuthConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("auth")]
        public Input<string>? Auth { get; set; }

        [Input("email")]
        public Input<string>? Email { get; set; }

        [Input("hostName", required: true)]
        public Input<string> HostName { get; set; } = null!;

        [Input("identityToken")]
        public Input<string>? IdentityToken { get; set; }

        [Input("password")]
        public Input<string>? Password { get; set; }

        [Input("registryToken")]
        public Input<string>? RegistryToken { get; set; }

        [Input("serverAddress")]
        public Input<string>? ServerAddress { get; set; }

        [Input("userName")]
        public Input<string>? UserName { get; set; }

        public RegistryImageBuildAuthConfigArgs()
        {
        }
        public static new RegistryImageBuildAuthConfigArgs Empty => new RegistryImageBuildAuthConfigArgs();
    }
}
