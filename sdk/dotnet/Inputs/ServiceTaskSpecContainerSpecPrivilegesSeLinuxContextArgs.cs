// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Docker.Inputs
{

    public sealed class ServiceTaskSpecContainerSpecPrivilegesSeLinuxContextArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Disable SELinux
        /// </summary>
        [Input("disable")]
        public Input<bool>? Disable { get; set; }

        /// <summary>
        /// SELinux level label
        /// </summary>
        [Input("level")]
        public Input<string>? Level { get; set; }

        /// <summary>
        /// SELinux role label
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// The mount type
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// SELinux user label
        /// </summary>
        [Input("user")]
        public Input<string>? User { get; set; }

        public ServiceTaskSpecContainerSpecPrivilegesSeLinuxContextArgs()
        {
        }
        public static new ServiceTaskSpecContainerSpecPrivilegesSeLinuxContextArgs Empty => new ServiceTaskSpecContainerSpecPrivilegesSeLinuxContextArgs();
    }
}
