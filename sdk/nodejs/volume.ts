// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

/**
 * <!-- Bug: Type and Name are switched -->
 *
 * Creates and destroys a volume in Docker. This can be used alongside docker.Container to prepare volumes that can be shared across containers.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as docker from "@pulumi/docker";
 *
 * const sharedVolume = new docker.Volume("shared_volume", {name: "shared_volume"});
 * ```
 *
 * ## Import
 *
 * ### Example
 *
 * Assuming you created a `volume` as follows
 *
 * #!/bin/bash
 *
 * docker volume create
 *
 * prints the long ID
 *
 * 524b0457aa2a87dd2b75c74c3e4e53f406974249e63ab3ed9bf21e5644f9dc7d
 *
 * you provide the definition for the resource as follows
 *
 * terraform
 *
 * resource "docker_volume" "foo" {
 *
 *   name = "524b0457aa2a87dd2b75c74c3e4e53f406974249e63ab3ed9bf21e5644f9dc7d"
 *
 * }
 *
 * then the import command is as follows
 *
 * #!/bin/bash
 *
 * ```sh
 * $ pulumi import docker:index/volume:Volume foo 524b0457aa2a87dd2b75c74c3e4e53f406974249e63ab3ed9bf21e5644f9dc7d
 * ```
 */
export class Volume extends pulumi.CustomResource {
    /**
     * Get an existing Volume resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VolumeState, opts?: pulumi.CustomResourceOptions): Volume {
        return new Volume(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'docker:index/volume:Volume';

    /**
     * Returns true if the given object is an instance of Volume.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Volume {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Volume.__pulumiType;
    }

    /**
     * Driver type for the volume. Defaults to `local`.
     */
    public readonly driver!: pulumi.Output<string>;
    /**
     * Options specific to the driver.
     */
    public readonly driverOpts!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * User-defined key/value metadata
     */
    public readonly labels!: pulumi.Output<outputs.VolumeLabel[] | undefined>;
    /**
     * The mountpoint of the volume.
     */
    public /*out*/ readonly mountpoint!: pulumi.Output<string>;
    /**
     * The name of the Docker volume (will be generated if not provided).
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a Volume resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: VolumeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VolumeArgs | VolumeState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VolumeState | undefined;
            resourceInputs["driver"] = state ? state.driver : undefined;
            resourceInputs["driverOpts"] = state ? state.driverOpts : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["mountpoint"] = state ? state.mountpoint : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as VolumeArgs | undefined;
            resourceInputs["driver"] = args ? args.driver : undefined;
            resourceInputs["driverOpts"] = args ? args.driverOpts : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["mountpoint"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Volume.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Volume resources.
 */
export interface VolumeState {
    /**
     * Driver type for the volume. Defaults to `local`.
     */
    driver?: pulumi.Input<string>;
    /**
     * Options specific to the driver.
     */
    driverOpts?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * User-defined key/value metadata
     */
    labels?: pulumi.Input<pulumi.Input<inputs.VolumeLabel>[]>;
    /**
     * The mountpoint of the volume.
     */
    mountpoint?: pulumi.Input<string>;
    /**
     * The name of the Docker volume (will be generated if not provided).
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Volume resource.
 */
export interface VolumeArgs {
    /**
     * Driver type for the volume. Defaults to `local`.
     */
    driver?: pulumi.Input<string>;
    /**
     * Options specific to the driver.
     */
    driverOpts?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * User-defined key/value metadata
     */
    labels?: pulumi.Input<pulumi.Input<inputs.VolumeLabel>[]>;
    /**
     * The name of the Docker volume (will be generated if not provided).
     */
    name?: pulumi.Input<string>;
}
