// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Docker.Inputs
{

    public sealed class ServiceTaskSpecContainerSpecPrivilegesGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("credentialSpec")]
        public Input<Inputs.ServiceTaskSpecContainerSpecPrivilegesCredentialSpecGetArgs>? CredentialSpec { get; set; }

        [Input("seLinuxContext")]
        public Input<Inputs.ServiceTaskSpecContainerSpecPrivilegesSeLinuxContextGetArgs>? SeLinuxContext { get; set; }

        public ServiceTaskSpecContainerSpecPrivilegesGetArgs()
        {
        }
        public static new ServiceTaskSpecContainerSpecPrivilegesGetArgs Empty => new ServiceTaskSpecContainerSpecPrivilegesGetArgs();
    }
}
