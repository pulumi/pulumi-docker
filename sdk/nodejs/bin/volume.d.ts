import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
/**
 * <!-- Bug: Type and Name are switched -->
 * Creates and destroys a volume in Docker. This can be used alongside docker.Container to prepare volumes that can be shared across containers.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as docker from "@pulumi/docker";
 *
 * const sharedVolume = new docker.Volume("sharedVolume", {});
 * ```
 *
 * ## Import
 *
 * ### Example Assuming you created a `volume` as follows #!/bin/bash docker volume create prints the long ID 524b0457aa2a87dd2b75c74c3e4e53f406974249e63ab3ed9bf21e5644f9dc7d you provide the definition for the resource as follows terraform resource "docker_volume" "foo" {
 *
 *  name = "524b0457aa2a87dd2b75c74c3e4e53f406974249e63ab3ed9bf21e5644f9dc7d" } then the import command is as follows #!/bin/bash
 *
 * ```sh
 *  $ pulumi import docker:index/volume:Volume foo 524b0457aa2a87dd2b75c74c3e4e53f406974249e63ab3ed9bf21e5644f9dc7d
 * ```
 */
export declare class Volume extends pulumi.CustomResource {
    /**
     * Get an existing Volume resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VolumeState, opts?: pulumi.CustomResourceOptions): Volume;
    /**
     * Returns true if the given object is an instance of Volume.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    static isInstance(obj: any): obj is Volume;
    /**
     * Driver type for the volume. Defaults to `local`.
     */
    readonly driver: pulumi.Output<string>;
    /**
     * Options specific to the driver.
     */
    readonly driverOpts: pulumi.Output<{
        [key: string]: any;
    } | undefined>;
    /**
     * User-defined key/value metadata
     */
    readonly labels: pulumi.Output<outputs.VolumeLabel[] | undefined>;
    /**
     * The mountpoint of the volume.
     */
    readonly mountpoint: pulumi.Output<string>;
    /**
     * The name of the Docker volume (will be generated if not provided).
     */
    readonly name: pulumi.Output<string>;
    /**
     * Create a Volume resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: VolumeArgs, opts?: pulumi.CustomResourceOptions);
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
    driverOpts?: pulumi.Input<{
        [key: string]: any;
    }>;
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
    driverOpts?: pulumi.Input<{
        [key: string]: any;
    }>;
    /**
     * User-defined key/value metadata
     */
    labels?: pulumi.Input<pulumi.Input<inputs.VolumeLabel>[]>;
    /**
     * The name of the Docker volume (will be generated if not provided).
     */
    name?: pulumi.Input<string>;
}
