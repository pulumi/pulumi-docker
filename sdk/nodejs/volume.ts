// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Creates and destroys a volume in Docker. This can be used alongside
 * [docker\_container](https://www.terraform.io/docs/providers/docker/r/container.html)
 * to prepare volumes that can be shared across containers.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as docker from "@pulumi/docker";
 * 
 * // Creates a docker volume "sharedVolume".
 * const sharedVolume = new docker.Volume("sharedVolume", {});
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-docker/blob/master/website/docs/r/volume.html.markdown.
 */
export class Volume extends pulumi.CustomResource {
    /**
     * Get an existing Volume resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
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
     * Driver type for the volume (defaults to local).
     */
    public readonly driver!: pulumi.Output<string>;
    /**
     * Options specific to the driver.
     */
    public readonly driverOpts!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * User-defined key/value metadata.
     */
    public readonly labels!: pulumi.Output<outputs.VolumeLabel[] | undefined>;
    public /*out*/ readonly mountpoint!: pulumi.Output<string>;
    /**
     * The name of the Docker volume (generated if not
     * provided).
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
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as VolumeState | undefined;
            inputs["driver"] = state ? state.driver : undefined;
            inputs["driverOpts"] = state ? state.driverOpts : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["mountpoint"] = state ? state.mountpoint : undefined;
            inputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as VolumeArgs | undefined;
            inputs["driver"] = args ? args.driver : undefined;
            inputs["driverOpts"] = args ? args.driverOpts : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["mountpoint"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Volume.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Volume resources.
 */
export interface VolumeState {
    /**
     * Driver type for the volume (defaults to local).
     */
    readonly driver?: pulumi.Input<string>;
    /**
     * Options specific to the driver.
     */
    readonly driverOpts?: pulumi.Input<{[key: string]: any}>;
    /**
     * User-defined key/value metadata.
     */
    readonly labels?: pulumi.Input<pulumi.Input<inputs.VolumeLabel>[]>;
    readonly mountpoint?: pulumi.Input<string>;
    /**
     * The name of the Docker volume (generated if not
     * provided).
     */
    readonly name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Volume resource.
 */
export interface VolumeArgs {
    /**
     * Driver type for the volume (defaults to local).
     */
    readonly driver?: pulumi.Input<string>;
    /**
     * Options specific to the driver.
     */
    readonly driverOpts?: pulumi.Input<{[key: string]: any}>;
    /**
     * User-defined key/value metadata.
     */
    readonly labels?: pulumi.Input<pulumi.Input<inputs.VolumeLabel>[]>;
    /**
     * The name of the Docker volume (generated if not
     * provided).
     */
    readonly name?: pulumi.Input<string>;
}
