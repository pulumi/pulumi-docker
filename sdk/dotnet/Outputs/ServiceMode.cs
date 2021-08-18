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
    public sealed class ServiceMode
    {
        public readonly bool? Global;
        public readonly Outputs.ServiceModeReplicated? Replicated;

        [OutputConstructor]
        private ServiceMode(
            bool? global,

            Outputs.ServiceModeReplicated? replicated)
        {
            Global = global;
            Replicated = replicated;
        }
    }
}
