// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Docker.Buildx.Inputs
{

    public sealed class CacheToGitHubActionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Ignore errors caused by failed cache exports.
        /// </summary>
        [Input("ignoreError")]
        public Input<bool>? IgnoreError { get; set; }

        [Input("mode")]
        public Input<Pulumi.Docker.Buildx.CacheMode>? Mode { get; set; }

        /// <summary>
        /// Which scope cache object belongs to.
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        [Input("token")]
        private Input<string>? _token;

        /// <summary>
        /// Access token
        /// </summary>
        public Input<string>? Token
        {
            get => _token;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _token = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Cache server URL
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public CacheToGitHubActionsArgs()
        {
            IgnoreError = false;
            Mode = Pulumi.Docker.Buildx.CacheMode.Min;
            Scope = Utilities.GetEnv("buildkit") ?? "";
            Token = Utilities.GetEnv("ACTIONS_RUNTIME_TOKEN") ?? "";
            Url = Utilities.GetEnv("ACTIONS_RUNTIME_URL") ?? "";
        }
        public static new CacheToGitHubActionsArgs Empty => new CacheToGitHubActionsArgs();
    }
}