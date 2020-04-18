// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Docker.Inputs
{

    public sealed class ServiceTaskSpecContainerSpecPrivilegesCredentialSpecArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Load credential spec from this file.
        /// </summary>
        [Input("file")]
        public Input<string>? File { get; set; }

        /// <summary>
        /// Load credential spec from this value in the Windows registry.
        /// </summary>
        [Input("registry")]
        public Input<string>? Registry { get; set; }

        public ServiceTaskSpecContainerSpecPrivilegesCredentialSpecArgs()
        {
        }
    }
}
