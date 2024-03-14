// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * An index (or manifest list) referencing one or more existing images.
 *
 * Useful for crafting a multi-platform image from several
 * platform-specific images.
 *
 * This creates an OCI image index or a Docker manifest list depending on
 * the media types of the source images.
 *
 * ## Example Usage
 * ### Multi-platform registry caching
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as docker from "@pulumi/docker";
 *
 * const amd64 = new docker.buildx.Image("amd64", {
 *     cacheFrom: [{
 *         registry: {
 *             ref: "docker.io/pulumi/pulumi:cache-amd64",
 *         },
 *     }],
 *     cacheTo: [{
 *         registry: {
 *             mode: docker.buildx.image.CacheMode.Max,
 *             ref: "docker.io/pulumi/pulumi:cache-amd64",
 *         },
 *     }],
 *     context: {
 *         location: "app",
 *     },
 *     platforms: [docker.buildx.image.Platform.Linux_amd64],
 *     tags: ["docker.io/pulumi/pulumi:3.107.0-amd64"],
 * });
 * const arm64 = new docker.buildx.Image("arm64", {
 *     cacheFrom: [{
 *         registry: {
 *             ref: "docker.io/pulumi/pulumi:cache-arm64",
 *         },
 *     }],
 *     cacheTo: [{
 *         registry: {
 *             mode: docker.buildx.image.CacheMode.Max,
 *             ref: "docker.io/pulumi/pulumi:cache-arm64",
 *         },
 *     }],
 *     context: {
 *         location: "app",
 *     },
 *     platforms: [docker.buildx.image.Platform.Linux_arm64],
 *     tags: ["docker.io/pulumi/pulumi:3.107.0-arm64"],
 * });
 * const index = new docker.buildx.Index("index", {
 *     sources: [
 *         amd64.ref,
 *         arm64.ref,
 *     ],
 *     tag: "docker.io/pulumi/pulumi:3.107.0",
 * });
 * export const ref = index.ref;
 * ```
 */
export class Index extends pulumi.CustomResource {
    /**
     * Get an existing Index resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Index {
        return new Index(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'docker:buildx/image:Index';

    /**
     * Returns true if the given object is an instance of Index.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Index {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Index.__pulumiType;
    }

    /**
     * If true, push the index to the target registry.
     *
     * Defaults to `true`.
     */
    public readonly push!: pulumi.Output<boolean | undefined>;
    /**
     * The pushed tag with digest.
     *
     * Identical to the tag if the index was not pushed.
     */
    public /*out*/ readonly ref!: pulumi.Output<string>;
    /**
     * Authentication for the registry where the tagged index will be pushed.
     *
     * Credentials can also be included with the provider's configuration.
     */
    public readonly registry!: pulumi.Output<outputs.buildx.RegistryAuth | undefined>;
    /**
     * Existing images to include in the index.
     */
    public readonly sources!: pulumi.Output<string[]>;
    /**
     * The tag to apply to the index.
     */
    public readonly tag!: pulumi.Output<string>;

    /**
     * Create a Index resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IndexArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.sources === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sources'");
            }
            if ((!args || args.tag === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tag'");
            }
            resourceInputs["push"] = (args ? args.push : undefined) ?? true;
            resourceInputs["registry"] = args ? args.registry : undefined;
            resourceInputs["sources"] = args ? args.sources : undefined;
            resourceInputs["tag"] = args ? args.tag : undefined;
            resourceInputs["ref"] = undefined /*out*/;
        } else {
            resourceInputs["push"] = undefined /*out*/;
            resourceInputs["ref"] = undefined /*out*/;
            resourceInputs["registry"] = undefined /*out*/;
            resourceInputs["sources"] = undefined /*out*/;
            resourceInputs["tag"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Index.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Index resource.
 */
export interface IndexArgs {
    /**
     * If true, push the index to the target registry.
     *
     * Defaults to `true`.
     */
    push?: pulumi.Input<boolean>;
    /**
     * Authentication for the registry where the tagged index will be pushed.
     *
     * Credentials can also be included with the provider's configuration.
     */
    registry?: pulumi.Input<inputs.buildx.RegistryAuth>;
    /**
     * Existing images to include in the index.
     */
    sources: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The tag to apply to the index.
     */
    tag: pulumi.Input<string>;
}
