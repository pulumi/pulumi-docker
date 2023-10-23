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
    public sealed class RemoteImageBuildUlimit
    {
        /// <summary>
        /// soft limit
        /// </summary>
        public readonly int Hard;
        /// <summary>
        /// type of ulimit, e.g. `nofile`
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// hard limit
        /// </summary>
        public readonly int Soft;

        [OutputConstructor]
        private RemoteImageBuildUlimit(
            int hard,

            string name,

            int soft)
        {
            Hard = hard;
            Name = name;
            Soft = soft;
        }
    }
}
