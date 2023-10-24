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
    public sealed class ServiceTaskSpecLogDriver
    {
        /// <summary>
        /// A random name for the port
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// A list of internal resolver variables to be modified (e.g., `debug`, `ndots:3`, etc.)
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Options;

        [OutputConstructor]
        private ServiceTaskSpecLogDriver(
            string name,

            ImmutableDictionary<string, string>? options)
        {
            Name = name;
            Options = options;
        }
    }
}
