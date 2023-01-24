// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Reads the local Docker plugin. The plugin must be installed locally.
 *
 * ## Example Usage
 *
 * ### With alias
 * data "docker.Plugin" "byAlias" {
 *   alias = "sample-volume-plugin:latest"
 * }
 */
export function getPlugin(args?: GetPluginArgs, opts?: pulumi.InvokeOptions): Promise<GetPluginResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("docker:index/getPlugin:getPlugin", {
        "alias": args.alias,
        "id": args.id,
    }, opts);
}

/**
 * A collection of arguments for invoking getPlugin.
 */
export interface GetPluginArgs {
    /**
     * The alias of the Docker plugin. If the tag is omitted, `:latest` is complemented to the attribute value.
     */
    alias?: string;
    /**
     * The ID of the plugin, which has precedence over the `alias` of both are given
     */
    id?: string;
}

/**
 * A collection of values returned by getPlugin.
 */
export interface GetPluginResult {
    /**
     * The alias of the Docker plugin. If the tag is omitted, `:latest` is complemented to the attribute value.
     */
    readonly alias?: string;
    /**
     * If `true` the plugin is enabled
     */
    readonly enabled: boolean;
    /**
     * The environment variables in the form of `KEY=VALUE`, e.g. `DEBUG=0`
     */
    readonly envs: string[];
    /**
     * If true, grant all permissions necessary to run the plugin
     */
    readonly grantAllPermissions: boolean;
    /**
     * The ID of the plugin, which has precedence over the `alias` of both are given
     */
    readonly id?: string;
    /**
     * The plugin name. If the tag is omitted, `:latest` is complemented to the attribute value.
     */
    readonly name: string;
    /**
     * The Docker Plugin Reference
     */
    readonly pluginReference: string;
}
/**
 * Reads the local Docker plugin. The plugin must be installed locally.
 *
 * ## Example Usage
 *
 * ### With alias
 * data "docker.Plugin" "byAlias" {
 *   alias = "sample-volume-plugin:latest"
 * }
 */
export function getPluginOutput(args?: GetPluginOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPluginResult> {
    return pulumi.output(args).apply((a: any) => getPlugin(a, opts))
}

/**
 * A collection of arguments for invoking getPlugin.
 */
export interface GetPluginOutputArgs {
    /**
     * The alias of the Docker plugin. If the tag is omitted, `:latest` is complemented to the attribute value.
     */
    alias?: pulumi.Input<string>;
    /**
     * The ID of the plugin, which has precedence over the `alias` of both are given
     */
    id?: pulumi.Input<string>;
}
