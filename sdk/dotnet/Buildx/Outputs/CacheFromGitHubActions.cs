// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Docker.Buildx.Outputs
{

    [OutputType]
    public sealed class CacheFromGitHubActions
    {
        /// <summary>
        /// Which scope cache object belongs to.
        /// </summary>
        public readonly string? Scope;
        /// <summary>
        /// Access token
        /// </summary>
        public readonly string? Token;
        /// <summary>
        /// Cache server URL
        /// </summary>
        public readonly string? Url;

        [OutputConstructor]
        private CacheFromGitHubActions(
            string? scope,

            string? token,

            string? url)
        {
            Scope = scope;
            Token = token;
            Url = url;
        }
    }
}