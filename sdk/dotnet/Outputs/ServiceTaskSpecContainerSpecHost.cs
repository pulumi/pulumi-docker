// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Docker.Outputs
{

    [OutputType]
    public sealed class ServiceTaskSpecContainerSpecHost
    {
        /// <summary>
        /// The name of the host
        /// </summary>
        public readonly string Host;
        /// <summary>
        /// The ip of the host
        /// </summary>
        public readonly string Ip;

        [OutputConstructor]
        private ServiceTaskSpecContainerSpecHost(
            string host,

            string ip)
        {
            Host = host;
            Ip = ip;
        }
    }
}
