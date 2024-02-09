// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Docker
{
    /// <summary>
    /// &lt;!-- Bug: Type and Name are switched --&gt;
    /// Manages the lifecycle of a Docker plugin.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Docker = Pulumi.Docker;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var sample_volume_plugin = new Docker.Plugin("sample-volume-plugin", new()
    ///     {
    ///         Alias = "sample-volume-plugin",
    ///         EnableTimeout = 60,
    ///         Enabled = false,
    ///         Envs = new[]
    ///         {
    ///             "DEBUG=1",
    ///         },
    ///         ForceDestroy = true,
    ///         ForceDisable = true,
    ///         GrantAllPermissions = true,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// #!/bin/bash
    /// 
    /// ```sh
    /// $ pulumi import docker:index/plugin:Plugin sample-volume-plugin "$(docker plugin inspect -f {{.ID}} tiborvass/sample-volume-plugin:latest)"
    /// ```
    /// </summary>
    [DockerResourceType("docker:index/plugin:Plugin")]
    public partial class Plugin : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Docker Plugin alias
        /// </summary>
        [Output("alias")]
        public Output<string> Alias { get; private set; } = null!;

        /// <summary>
        /// HTTP client timeout to enable the plugin
        /// </summary>
        [Output("enableTimeout")]
        public Output<int?> EnableTimeout { get; private set; } = null!;

        /// <summary>
        /// If `true` the plugin is enabled. Defaults to `true`
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// The environment variables in the form of `KEY=VALUE`, e.g. `DEBUG=0`
        /// </summary>
        [Output("envs")]
        public Output<ImmutableArray<string>> Envs { get; private set; } = null!;

        /// <summary>
        /// If true, then the plugin is destroyed forcibly
        /// </summary>
        [Output("forceDestroy")]
        public Output<bool?> ForceDestroy { get; private set; } = null!;

        /// <summary>
        /// If true, then the plugin is disabled forcibly
        /// </summary>
        [Output("forceDisable")]
        public Output<bool?> ForceDisable { get; private set; } = null!;

        /// <summary>
        /// If true, grant all permissions necessary to run the plugin
        /// </summary>
        [Output("grantAllPermissions")]
        public Output<bool?> GrantAllPermissions { get; private set; } = null!;

        /// <summary>
        /// Grant specific permissions only
        /// </summary>
        [Output("grantPermissions")]
        public Output<ImmutableArray<Outputs.PluginGrantPermission>> GrantPermissions { get; private set; } = null!;

        /// <summary>
        /// The name of the permission
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Docker Plugin Reference
        /// </summary>
        [Output("pluginReference")]
        public Output<string> PluginReference { get; private set; } = null!;


        /// <summary>
        /// Create a Plugin resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Plugin(string name, PluginArgs? args = null, CustomResourceOptions? options = null)
            : base("docker:index/plugin:Plugin", name, args ?? new PluginArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Plugin(string name, Input<string> id, PluginState? state = null, CustomResourceOptions? options = null)
            : base("docker:index/plugin:Plugin", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Plugin resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Plugin Get(string name, Input<string> id, PluginState? state = null, CustomResourceOptions? options = null)
        {
            return new Plugin(name, id, state, options);
        }
    }

    public sealed class PluginArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Docker Plugin alias
        /// </summary>
        [Input("alias")]
        public Input<string>? Alias { get; set; }

        /// <summary>
        /// HTTP client timeout to enable the plugin
        /// </summary>
        [Input("enableTimeout")]
        public Input<int>? EnableTimeout { get; set; }

        /// <summary>
        /// If `true` the plugin is enabled. Defaults to `true`
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("envs")]
        private InputList<string>? _envs;

        /// <summary>
        /// The environment variables in the form of `KEY=VALUE`, e.g. `DEBUG=0`
        /// </summary>
        public InputList<string> Envs
        {
            get => _envs ?? (_envs = new InputList<string>());
            set => _envs = value;
        }

        /// <summary>
        /// If true, then the plugin is destroyed forcibly
        /// </summary>
        [Input("forceDestroy")]
        public Input<bool>? ForceDestroy { get; set; }

        /// <summary>
        /// If true, then the plugin is disabled forcibly
        /// </summary>
        [Input("forceDisable")]
        public Input<bool>? ForceDisable { get; set; }

        /// <summary>
        /// If true, grant all permissions necessary to run the plugin
        /// </summary>
        [Input("grantAllPermissions")]
        public Input<bool>? GrantAllPermissions { get; set; }

        [Input("grantPermissions")]
        private InputList<Inputs.PluginGrantPermissionArgs>? _grantPermissions;

        /// <summary>
        /// Grant specific permissions only
        /// </summary>
        public InputList<Inputs.PluginGrantPermissionArgs> GrantPermissions
        {
            get => _grantPermissions ?? (_grantPermissions = new InputList<Inputs.PluginGrantPermissionArgs>());
            set => _grantPermissions = value;
        }

        /// <summary>
        /// The name of the permission
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public PluginArgs()
        {
        }
        public static new PluginArgs Empty => new PluginArgs();
    }

    public sealed class PluginState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Docker Plugin alias
        /// </summary>
        [Input("alias")]
        public Input<string>? Alias { get; set; }

        /// <summary>
        /// HTTP client timeout to enable the plugin
        /// </summary>
        [Input("enableTimeout")]
        public Input<int>? EnableTimeout { get; set; }

        /// <summary>
        /// If `true` the plugin is enabled. Defaults to `true`
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("envs")]
        private InputList<string>? _envs;

        /// <summary>
        /// The environment variables in the form of `KEY=VALUE`, e.g. `DEBUG=0`
        /// </summary>
        public InputList<string> Envs
        {
            get => _envs ?? (_envs = new InputList<string>());
            set => _envs = value;
        }

        /// <summary>
        /// If true, then the plugin is destroyed forcibly
        /// </summary>
        [Input("forceDestroy")]
        public Input<bool>? ForceDestroy { get; set; }

        /// <summary>
        /// If true, then the plugin is disabled forcibly
        /// </summary>
        [Input("forceDisable")]
        public Input<bool>? ForceDisable { get; set; }

        /// <summary>
        /// If true, grant all permissions necessary to run the plugin
        /// </summary>
        [Input("grantAllPermissions")]
        public Input<bool>? GrantAllPermissions { get; set; }

        [Input("grantPermissions")]
        private InputList<Inputs.PluginGrantPermissionGetArgs>? _grantPermissions;

        /// <summary>
        /// Grant specific permissions only
        /// </summary>
        public InputList<Inputs.PluginGrantPermissionGetArgs> GrantPermissions
        {
            get => _grantPermissions ?? (_grantPermissions = new InputList<Inputs.PluginGrantPermissionGetArgs>());
            set => _grantPermissions = value;
        }

        /// <summary>
        /// The name of the permission
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Docker Plugin Reference
        /// </summary>
        [Input("pluginReference")]
        public Input<string>? PluginReference { get; set; }

        public PluginState()
        {
        }
        public static new PluginState Empty => new PluginState();
    }
}
