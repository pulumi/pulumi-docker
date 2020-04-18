// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Docker.Outputs
{

    [OutputType]
    public sealed class ServiceTaskSpecContainerSpecPrivilegesSeLinuxContext
    {
        /// <summary>
        /// Disable SELinux
        /// </summary>
        public readonly bool? Disable;
        /// <summary>
        /// SELinux level label
        /// </summary>
        public readonly string? Level;
        /// <summary>
        /// SELinux role label
        /// </summary>
        public readonly string? Role;
        /// <summary>
        /// SELinux type label
        /// </summary>
        public readonly string? Type;
        /// <summary>
        /// The user inside the container.
        /// </summary>
        public readonly string? User;

        [OutputConstructor]
        private ServiceTaskSpecContainerSpecPrivilegesSeLinuxContext(
            bool? disable,

            string? level,

            string? role,

            string? type,

            string? user)
        {
            Disable = disable;
            Level = level;
            Role = role;
            Type = type;
            User = user;
        }
    }
}
