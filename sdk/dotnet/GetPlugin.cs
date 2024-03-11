// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Docker
{
    public static class GetPlugin
    {
        /// <summary>
        /// Reads the local Docker plugin. The plugin must be installed locally.
        /// 
        /// ## Example Usage
        /// 
        /// ### With alias
        /// data "docker.Plugin" "by_alias" {
        ///   alias = "sample-volume-plugin:latest"
        /// }
        /// 
        /// ### With ID
        /// data "docker.Plugin" "by_id" {
        ///   id = "e9a9db917b3bfd6706b5d3a66d4bceb9f"
        /// }
        /// ```
        /// </summary>
        public static Task<GetPluginResult> InvokeAsync(GetPluginArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPluginResult>("docker:index/getPlugin:getPlugin", args ?? new GetPluginArgs(), options.WithDefaults());

        /// <summary>
        /// Reads the local Docker plugin. The plugin must be installed locally.
        /// 
        /// ## Example Usage
        /// 
        /// ### With alias
        /// data "docker.Plugin" "by_alias" {
        ///   alias = "sample-volume-plugin:latest"
        /// }
        /// 
        /// ### With ID
        /// data "docker.Plugin" "by_id" {
        ///   id = "e9a9db917b3bfd6706b5d3a66d4bceb9f"
        /// }
        /// ```
        /// </summary>
        public static Output<GetPluginResult> Invoke(GetPluginInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPluginResult>("docker:index/getPlugin:getPlugin", args ?? new GetPluginInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPluginArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The alias of the Docker plugin. If the tag is omitted, `:latest` is complemented to the attribute value.
        /// </summary>
        [Input("alias")]
        public string? Alias { get; set; }

        /// <summary>
        /// The ID of the plugin, which has precedence over the `alias` of both are given
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        public GetPluginArgs()
        {
        }
        public static new GetPluginArgs Empty => new GetPluginArgs();
    }

    public sealed class GetPluginInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The alias of the Docker plugin. If the tag is omitted, `:latest` is complemented to the attribute value.
        /// </summary>
        [Input("alias")]
        public Input<string>? Alias { get; set; }

        /// <summary>
        /// The ID of the plugin, which has precedence over the `alias` of both are given
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        public GetPluginInvokeArgs()
        {
        }
        public static new GetPluginInvokeArgs Empty => new GetPluginInvokeArgs();
    }


    [OutputType]
    public sealed class GetPluginResult
    {
        /// <summary>
        /// The alias of the Docker plugin. If the tag is omitted, `:latest` is complemented to the attribute value.
        /// </summary>
        public readonly string? Alias;
        /// <summary>
        /// If `true` the plugin is enabled
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// The environment variables in the form of `KEY=VALUE`, e.g. `DEBUG=0`
        /// </summary>
        public readonly ImmutableArray<string> Envs;
        /// <summary>
        /// If true, grant all permissions necessary to run the plugin
        /// </summary>
        public readonly bool GrantAllPermissions;
        /// <summary>
        /// The ID of the plugin, which has precedence over the `alias` of both are given
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The plugin name. If the tag is omitted, `:latest` is complemented to the attribute value.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The Docker Plugin Reference
        /// </summary>
        public readonly string PluginReference;

        [OutputConstructor]
        private GetPluginResult(
            string? alias,

            bool enabled,

            ImmutableArray<string> envs,

            bool grantAllPermissions,

            string? id,

            string name,

            string pluginReference)
        {
            Alias = alias;
            Enabled = enabled;
            Envs = envs;
            GrantAllPermissions = grantAllPermissions;
            Id = id;
            Name = name;
            PluginReference = pluginReference;
        }
    }
}
