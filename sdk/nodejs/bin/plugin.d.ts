import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
/**
 * <!-- Bug: Type and Name are switched -->
 * Manages the lifecycle of a Docker plugin.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as docker from "@pulumi/docker";
 *
 * const sample_volume_plugin = new docker.Plugin("sample-volume-plugin", {
 *     alias: "sample-volume-plugin",
 *     enableTimeout: 60,
 *     enabled: false,
 *     envs: ["DEBUG=1"],
 *     forceDestroy: true,
 *     forceDisable: true,
 *     grantAllPermissions: true,
 * });
 * ```
 *
 * ## Import
 *
 * #!/bin/bash
 *
 * ```sh
 *  $ pulumi import docker:index/plugin:Plugin sample-volume-plugin "$(docker plugin inspect -f {{.ID}} tiborvass/sample-volume-plugin:latest)"
 * ```
 */
export declare class Plugin extends pulumi.CustomResource {
    /**
     * Get an existing Plugin resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PluginState, opts?: pulumi.CustomResourceOptions): Plugin;
    /**
     * Returns true if the given object is an instance of Plugin.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    static isInstance(obj: any): obj is Plugin;
    /**
     * Docker Plugin alias
     */
    readonly alias: pulumi.Output<string>;
    /**
     * HTTP client timeout to enable the plugin
     */
    readonly enableTimeout: pulumi.Output<number | undefined>;
    /**
     * If `true` the plugin is enabled. Defaults to `true`
     */
    readonly enabled: pulumi.Output<boolean | undefined>;
    /**
     * The environment variables in the form of `KEY=VALUE`, e.g. `DEBUG=0`
     */
    readonly envs: pulumi.Output<string[]>;
    /**
     * If true, then the plugin is destroyed forcibly
     */
    readonly forceDestroy: pulumi.Output<boolean | undefined>;
    /**
     * If true, then the plugin is disabled forcibly
     */
    readonly forceDisable: pulumi.Output<boolean | undefined>;
    /**
     * If true, grant all permissions necessary to run the plugin
     */
    readonly grantAllPermissions: pulumi.Output<boolean | undefined>;
    /**
     * Grant specific permissions only
     */
    readonly grantPermissions: pulumi.Output<outputs.PluginGrantPermission[] | undefined>;
    /**
     * The name of the permission
     */
    readonly name: pulumi.Output<string>;
    /**
     * Docker Plugin Reference
     */
    readonly pluginReference: pulumi.Output<string>;
    /**
     * Create a Plugin resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: PluginArgs, opts?: pulumi.CustomResourceOptions);
}
/**
 * Input properties used for looking up and filtering Plugin resources.
 */
export interface PluginState {
    /**
     * Docker Plugin alias
     */
    alias?: pulumi.Input<string>;
    /**
     * HTTP client timeout to enable the plugin
     */
    enableTimeout?: pulumi.Input<number>;
    /**
     * If `true` the plugin is enabled. Defaults to `true`
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The environment variables in the form of `KEY=VALUE`, e.g. `DEBUG=0`
     */
    envs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If true, then the plugin is destroyed forcibly
     */
    forceDestroy?: pulumi.Input<boolean>;
    /**
     * If true, then the plugin is disabled forcibly
     */
    forceDisable?: pulumi.Input<boolean>;
    /**
     * If true, grant all permissions necessary to run the plugin
     */
    grantAllPermissions?: pulumi.Input<boolean>;
    /**
     * Grant specific permissions only
     */
    grantPermissions?: pulumi.Input<pulumi.Input<inputs.PluginGrantPermission>[]>;
    /**
     * The name of the permission
     */
    name?: pulumi.Input<string>;
    /**
     * Docker Plugin Reference
     */
    pluginReference?: pulumi.Input<string>;
}
/**
 * The set of arguments for constructing a Plugin resource.
 */
export interface PluginArgs {
    /**
     * Docker Plugin alias
     */
    alias?: pulumi.Input<string>;
    /**
     * HTTP client timeout to enable the plugin
     */
    enableTimeout?: pulumi.Input<number>;
    /**
     * If `true` the plugin is enabled. Defaults to `true`
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The environment variables in the form of `KEY=VALUE`, e.g. `DEBUG=0`
     */
    envs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If true, then the plugin is destroyed forcibly
     */
    forceDestroy?: pulumi.Input<boolean>;
    /**
     * If true, then the plugin is disabled forcibly
     */
    forceDisable?: pulumi.Input<boolean>;
    /**
     * If true, grant all permissions necessary to run the plugin
     */
    grantAllPermissions?: pulumi.Input<boolean>;
    /**
     * Grant specific permissions only
     */
    grantPermissions?: pulumi.Input<pulumi.Input<inputs.PluginGrantPermission>[]>;
    /**
     * The name of the permission
     */
    name?: pulumi.Input<string>;
}
