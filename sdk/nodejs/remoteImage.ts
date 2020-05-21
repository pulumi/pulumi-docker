// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Pulls a Docker image to a given Docker host from a Docker Registry.
 *
 * This resource will *not* pull new layers of the image automatically unless used in
 * conjunction with [`docker..getRegistryImage`](https://www.terraform.io/docs/providers/docker/d/registry_image.html)
 * data source to update the `pullTriggers` field.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as docker from "@pulumi/docker";
 *
 * // Find the latest Ubuntu precise image.
 * const ubuntu = new docker.RemoteImage("ubuntu", {
 *     name: "ubuntu:precise",
 * });
 * ```
 *
 * ### Dynamic image
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as docker from "@pulumi/docker";
 *
 * const ubuntuRegistryImage = pulumi.output(docker.getRegistryImage({
 *     name: "ubuntu:precise",
 * }, { async: true }));
 * const ubuntuRemoteImage = new docker.RemoteImage("ubuntu", {
 *     name: ubuntuRegistryImage.name!,
 *     pullTriggers: [ubuntuRegistryImage.sha256Digest],
 * });
 * ```
 */
export class RemoteImage extends pulumi.CustomResource {
    /**
     * Get an existing RemoteImage resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RemoteImageState, opts?: pulumi.CustomResourceOptions): RemoteImage {
        return new RemoteImage(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'docker:index/remoteImage:RemoteImage';

    /**
     * Returns true if the given object is an instance of RemoteImage.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RemoteImage {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RemoteImage.__pulumiType;
    }

    /**
     * If true, then the Docker image won't be
     * deleted on destroy operation. If this is false, it will delete the image from
     * the docker local storage on destroy operation.
     */
    public readonly keepLocally!: pulumi.Output<boolean | undefined>;
    public /*out*/ readonly latest!: pulumi.Output<string>;
    /**
     * The name of the Docker image, including any tags or SHA256 repo digests.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * **Deprecated**, use `pullTriggers` instead.
     */
    public readonly pullTrigger!: pulumi.Output<string | undefined>;
    /**
     * List of values which cause an
     * image pull when changed. This is used to store the image digest from the
     * registry when using the `docker..getRegistryImage` [data source](https://www.terraform.io/docs/providers/docker/d/registry_image.html)
     * to trigger an image update.
     */
    public readonly pullTriggers!: pulumi.Output<string[] | undefined>;

    /**
     * Create a RemoteImage resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RemoteImageArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RemoteImageArgs | RemoteImageState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as RemoteImageState | undefined;
            inputs["keepLocally"] = state ? state.keepLocally : undefined;
            inputs["latest"] = state ? state.latest : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["pullTrigger"] = state ? state.pullTrigger : undefined;
            inputs["pullTriggers"] = state ? state.pullTriggers : undefined;
        } else {
            const args = argsOrState as RemoteImageArgs | undefined;
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            inputs["keepLocally"] = args ? args.keepLocally : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["pullTrigger"] = args ? args.pullTrigger : undefined;
            inputs["pullTriggers"] = args ? args.pullTriggers : undefined;
            inputs["latest"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(RemoteImage.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RemoteImage resources.
 */
export interface RemoteImageState {
    /**
     * If true, then the Docker image won't be
     * deleted on destroy operation. If this is false, it will delete the image from
     * the docker local storage on destroy operation.
     */
    readonly keepLocally?: pulumi.Input<boolean>;
    readonly latest?: pulumi.Input<string>;
    /**
     * The name of the Docker image, including any tags or SHA256 repo digests.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * **Deprecated**, use `pullTriggers` instead.
     * @deprecated Use field pull_triggers instead
     */
    readonly pullTrigger?: pulumi.Input<string>;
    /**
     * List of values which cause an
     * image pull when changed. This is used to store the image digest from the
     * registry when using the `docker..getRegistryImage` [data source](https://www.terraform.io/docs/providers/docker/d/registry_image.html)
     * to trigger an image update.
     */
    readonly pullTriggers?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a RemoteImage resource.
 */
export interface RemoteImageArgs {
    /**
     * If true, then the Docker image won't be
     * deleted on destroy operation. If this is false, it will delete the image from
     * the docker local storage on destroy operation.
     */
    readonly keepLocally?: pulumi.Input<boolean>;
    /**
     * The name of the Docker image, including any tags or SHA256 repo digests.
     */
    readonly name: pulumi.Input<string>;
    /**
     * **Deprecated**, use `pullTriggers` instead.
     * @deprecated Use field pull_triggers instead
     */
    readonly pullTrigger?: pulumi.Input<string>;
    /**
     * List of values which cause an
     * image pull when changed. This is used to store the image digest from the
     * registry when using the `docker..getRegistryImage` [data source](https://www.terraform.io/docs/providers/docker/d/registry_image.html)
     * to trigger an image update.
     */
    readonly pullTriggers?: pulumi.Input<pulumi.Input<string>[]>;
}
