// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * A Docker image built using buildx -- Docker's interface to the improved
 * BuildKit backend.
 *
 * **This resource is experimental and subject to change.**
 *
 * API types are unstable. Subsequent releases _may_ require manual edits
 * to your state file(s) in order to adopt API changes.
 *
 * Only use this resource if you understand and accept the risks.
 *
 * ## Example Usage
 * ### Multi-platform image
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as docker from "@pulumi/docker";
 *
 * const image = new docker.buildx.Image("image", {
 *     context: {
 *         location: "app",
 *     },
 *     dockerfile: {
 *         location: "app/Dockerfile",
 *     },
 *     platforms: [
 *         docker.buildx.image.Platform.Plan9_amd64,
 *         docker.buildx.image.Platform.Plan9_386,
 *     ],
 * });
 * ```
 * ### Registry export
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as docker from "@pulumi/docker";
 *
 * const image = new docker.buildx.Image("image", {
 *     context: {
 *         location: "app",
 *     },
 *     dockerfile: {
 *         location: "app/Dockerfile",
 *     },
 *     exports: [{
 *         registry: {
 *             ociMediaTypes: true,
 *         },
 *     }],
 *     registries: [{
 *         address: "docker.io",
 *         password: dockerHubPassword,
 *         username: "pulumibot",
 *     }],
 * });
 * ```
 * ### Caching
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as docker from "@pulumi/docker";
 *
 * const image = new docker.buildx.Image("image", {
 *     cacheFrom: [{
 *         local: {
 *             src: "tmp/cache",
 *         },
 *     }],
 *     cacheTo: [{
 *         local: {
 *             dest: "tmp/cache",
 *             mode: docker.buildx.image.CacheMode.Max,
 *         },
 *     }],
 *     context: {
 *         location: "app",
 *     },
 *     dockerfile: {
 *         location: "app/Dockerfile",
 *     },
 * });
 * ```
 * ### Build arguments
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as docker from "@pulumi/docker";
 *
 * const image = new docker.buildx.Image("image", {
 *     buildArgs: {
 *         SET_ME_TO_TRUE: "true",
 *     },
 *     context: {
 *         location: "app",
 *     },
 *     dockerfile: {
 *         location: "app/Dockerfile",
 *     },
 * });
 * ```
 * ### Build targets
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as docker from "@pulumi/docker";
 *
 * const image = new docker.buildx.Image("image", {
 *     context: {
 *         location: "app",
 *     },
 *     dockerfile: {
 *         location: "app/Dockerfile",
 *     },
 *     targets: [
 *         "build-me",
 *         "also-build-me",
 *     ],
 * });
 * ```
 * ### Named contexts
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as docker from "@pulumi/docker";
 *
 * const image = new docker.buildx.Image("image", {
 *     context: {
 *         location: "app",
 *         named: {
 *             "golang:latest": {
 *                 location: "docker-image://golang@sha256:b8e62cf593cdaff36efd90aa3a37de268e6781a2e68c6610940c48f7cdf36984",
 *             },
 *         },
 *     },
 *     dockerfile: {
 *         location: "app/Dockerfile",
 *     },
 * });
 * ```
 * ### Remote context
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as docker from "@pulumi/docker";
 *
 * const image = new docker.buildx.Image("image", {context: {
 *     location: "https://raw.githubusercontent.com/pulumi/pulumi-docker/api-types/provider/testdata/Dockerfile",
 * }});
 * ```
 * ### Inline Dockerfile
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as docker from "@pulumi/docker";
 *
 * const image = new docker.buildx.Image("image", {
 *     context: {
 *         location: "app",
 *     },
 *     dockerfile: {
 *         inline: `FROM busybox
 * COPY hello.c ./
 * `,
 *     },
 * });
 * ```
 * ### Remote context
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as docker from "@pulumi/docker";
 *
 * const image = new docker.buildx.Image("image", {
 *     context: {
 *         location: "https://github.com/docker-library/hello-world.git",
 *     },
 *     dockerfile: {
 *         location: "app/Dockerfile",
 *     },
 * });
 * ```
 * ### Local export
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as docker from "@pulumi/docker";
 *
 * const image = new docker.buildx.Image("image", {
 *     context: {
 *         location: "app",
 *     },
 *     dockerfile: {
 *         location: "app/Dockerfile",
 *     },
 *     exports: [{
 *         docker: {
 *             tar: true,
 *         },
 *     }],
 * });
 * ```
 */
export class Image extends pulumi.CustomResource {
    /**
     * Get an existing Image resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Image {
        return new Image(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'docker:buildx/image:Image';

    /**
     * Returns true if the given object is an instance of Image.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Image {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Image.__pulumiType;
    }

    /**
     * `ARG` names and values to set during the build.
     *
     * These variables are accessed like environment variables inside `RUN`
     * instructions.
     *
     * Build arguments are persisted in the image, so you should use `secrets`
     * if these arguments are sensitive.
     */
    public readonly buildArgs!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * When `true`, attempt to build the image during previews. The image will
     * not be pushed to registries, however caches will still populated.
     */
    public readonly buildOnPreview!: pulumi.Output<boolean | undefined>;
    /**
     * Builder configuration.
     */
    public readonly builder!: pulumi.Output<outputs.buildx.BuilderConfig | undefined>;
    /**
     * External cache configuration.
     */
    public readonly cacheFrom!: pulumi.Output<outputs.buildx.CacheFromEntry[] | undefined>;
    /**
     * Cache export configuration.
     */
    public readonly cacheTo!: pulumi.Output<outputs.buildx.CacheToEntry[] | undefined>;
    /**
     * Build context settings.
     */
    public readonly context!: pulumi.Output<outputs.buildx.BuildContext | undefined>;
    /**
     * A preliminary hash of the image's build context.
     *
     * Pulumi uses this to determine if an image _may_ need to be re-built.
     */
    public /*out*/ readonly contextHash!: pulumi.Output<string | undefined>;
    /**
     * A mapping of platform type to refs which were pushed to registries.
     */
    public /*out*/ readonly digests!: pulumi.Output<{[key: string]: string[]} | undefined>;
    /**
     * Dockerfile settings.
     */
    public readonly dockerfile!: pulumi.Output<outputs.buildx.Dockerfile | undefined>;
    /**
     * Controls where images are persisted after building.
     *
     * Images are only stored in the local cache unless `exports` are
     * explicitly configured.
     */
    public readonly exports!: pulumi.Output<outputs.buildx.ExportEntry[] | undefined>;
    /**
     * Attach arbitrary key/value metadata to the image.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Set target platform(s) for the build. Defaults to the host's platform
     */
    public readonly platforms!: pulumi.Output<enums.buildx.Platform[] | undefined>;
    /**
     * Always pull referenced images.
     */
    public readonly pull!: pulumi.Output<boolean | undefined>;
    /**
     * Registry credentials. Required if reading or exporting to private
     * repositories.
     */
    public readonly registries!: pulumi.Output<outputs.buildx.RegistryAuth[] | undefined>;
    /**
     * A mapping of secret names to their corresponding values.
     *
     * Unlike the Docker CLI, these can be passed by value and do not need to
     * exist on-disk or in environment variables.
     *
     * Build arguments and environment variables are persistent in the final
     * image, so you should use this for sensitive values.
     */
    public readonly secrets!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Name and optionally a tag (format: `name:tag`).
     *
     * If exporting to a registry, the name should include the fully qualified
     * registry address (e.g. `docker.io/pulumi/pulumi:latest`).
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * Set the target build stage(s) to build.
     *
     * If not specified all targets will be built by default.
     */
    public readonly targets!: pulumi.Output<string[] | undefined>;

    /**
     * Create a Image resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ImageArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["buildArgs"] = args ? args.buildArgs : undefined;
            resourceInputs["buildOnPreview"] = args ? args.buildOnPreview : undefined;
            resourceInputs["builder"] = args ? args.builder : undefined;
            resourceInputs["cacheFrom"] = args ? args.cacheFrom : undefined;
            resourceInputs["cacheTo"] = args ? args.cacheTo : undefined;
            resourceInputs["context"] = args ? args.context : undefined;
            resourceInputs["dockerfile"] = args ? args.dockerfile : undefined;
            resourceInputs["exports"] = args ? args.exports : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["platforms"] = args ? args.platforms : undefined;
            resourceInputs["pull"] = args ? args.pull : undefined;
            resourceInputs["registries"] = args ? args.registries : undefined;
            resourceInputs["secrets"] = args?.secrets ? pulumi.secret(args.secrets) : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["targets"] = args ? args.targets : undefined;
            resourceInputs["contextHash"] = undefined /*out*/;
            resourceInputs["digests"] = undefined /*out*/;
        } else {
            resourceInputs["buildArgs"] = undefined /*out*/;
            resourceInputs["buildOnPreview"] = undefined /*out*/;
            resourceInputs["builder"] = undefined /*out*/;
            resourceInputs["cacheFrom"] = undefined /*out*/;
            resourceInputs["cacheTo"] = undefined /*out*/;
            resourceInputs["context"] = undefined /*out*/;
            resourceInputs["contextHash"] = undefined /*out*/;
            resourceInputs["digests"] = undefined /*out*/;
            resourceInputs["dockerfile"] = undefined /*out*/;
            resourceInputs["exports"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["platforms"] = undefined /*out*/;
            resourceInputs["pull"] = undefined /*out*/;
            resourceInputs["registries"] = undefined /*out*/;
            resourceInputs["secrets"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["targets"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["secrets"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Image.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Image resource.
 */
export interface ImageArgs {
    /**
     * `ARG` names and values to set during the build.
     *
     * These variables are accessed like environment variables inside `RUN`
     * instructions.
     *
     * Build arguments are persisted in the image, so you should use `secrets`
     * if these arguments are sensitive.
     */
    buildArgs?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * When `true`, attempt to build the image during previews. The image will
     * not be pushed to registries, however caches will still populated.
     */
    buildOnPreview?: pulumi.Input<boolean>;
    /**
     * Builder configuration.
     */
    builder?: pulumi.Input<inputs.buildx.BuilderConfig>;
    /**
     * External cache configuration.
     */
    cacheFrom?: pulumi.Input<pulumi.Input<inputs.buildx.CacheFromEntry>[]>;
    /**
     * Cache export configuration.
     */
    cacheTo?: pulumi.Input<pulumi.Input<inputs.buildx.CacheToEntry>[]>;
    /**
     * Build context settings.
     */
    context?: pulumi.Input<inputs.buildx.BuildContext>;
    /**
     * Dockerfile settings.
     */
    dockerfile?: pulumi.Input<inputs.buildx.Dockerfile>;
    /**
     * Controls where images are persisted after building.
     *
     * Images are only stored in the local cache unless `exports` are
     * explicitly configured.
     */
    exports?: pulumi.Input<pulumi.Input<inputs.buildx.ExportEntry>[]>;
    /**
     * Attach arbitrary key/value metadata to the image.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Set target platform(s) for the build. Defaults to the host's platform
     */
    platforms?: pulumi.Input<pulumi.Input<enums.buildx.Platform>[]>;
    /**
     * Always pull referenced images.
     */
    pull?: pulumi.Input<boolean>;
    /**
     * Registry credentials. Required if reading or exporting to private
     * repositories.
     */
    registries?: pulumi.Input<pulumi.Input<inputs.buildx.RegistryAuth>[]>;
    /**
     * A mapping of secret names to their corresponding values.
     *
     * Unlike the Docker CLI, these can be passed by value and do not need to
     * exist on-disk or in environment variables.
     *
     * Build arguments and environment variables are persistent in the final
     * image, so you should use this for sensitive values.
     */
    secrets?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Name and optionally a tag (format: `name:tag`).
     *
     * If exporting to a registry, the name should include the fully qualified
     * registry address (e.g. `docker.io/pulumi/pulumi:latest`).
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Set the target build stage(s) to build.
     *
     * If not specified all targets will be built by default.
     */
    targets?: pulumi.Input<pulumi.Input<string>[]>;
}
