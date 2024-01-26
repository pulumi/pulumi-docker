// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * A Docker image built using Buildkit
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
     *
     * An optional map of named build-time argument variables to set during
     * the Docker build. This flag allows you to pass build-time variables that
     * can be accessed like environment variables inside the RUN
     * instruction.
     */
    public readonly buildArgs!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     *
     * External cache sources (e.g., "user/app:cache", "type=local,src=path/to/dir")
     */
    public readonly cacheFrom!: pulumi.Output<string[] | undefined>;
    /**
     *
     * Cache export destinations (e.g., "user/app:cache", "type=local,dest=path/to/dir")
     */
    public readonly cacheTo!: pulumi.Output<string[] | undefined>;
    /**
     *
     * Contexts to use while building the image. If omitted, an empty context
     * is used. If more than one value is specified, they should be of the
     * form "name=value".
     */
    public readonly context!: pulumi.Output<string | undefined>;
    public /*out*/ readonly contextHash!: pulumi.Output<string | undefined>;
    /**
     *
     * Name and optionally a tag (format: "name:tag"). If outputting to a
     * registry, the name should include the fully qualified registry address.
     */
    public readonly exports!: pulumi.Output<string[] | undefined>;
    /**
     *
     * Name of the Dockerfile to use (default: "$PATH/Dockerfile").
     */
    public readonly file!: pulumi.Output<string | undefined>;
    public /*out*/ readonly manifests!: pulumi.Output<outputs.buildx.Manifest[]>;
    /**
     *
     * Set target platforms for the build. Defaults to the host's platform
     */
    public readonly platforms!: pulumi.Output<string[] | undefined>;
    /**
     *
     * Always attempt to pull all referenced images
     */
    public readonly pull!: pulumi.Output<boolean | undefined>;
    /**
     *
     * Logins for registry outputs
     */
    public readonly registries!: pulumi.Output<outputs.buildx.RegistryAuth[] | undefined>;
    /**
     *
     * Name and optionally a tag (format: "name:tag"). If outputting to a
     * registry, the name should include the fully qualified registry address.
     */
    public readonly tags!: pulumi.Output<string[]>;

    /**
     * Create a Image resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ImageArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.tags === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tags'");
            }
            resourceInputs["buildArgs"] = args ? args.buildArgs : undefined;
            resourceInputs["cacheFrom"] = args ? args.cacheFrom : undefined;
            resourceInputs["cacheTo"] = args ? args.cacheTo : undefined;
            resourceInputs["context"] = args ? args.context : undefined;
            resourceInputs["exports"] = args ? args.exports : undefined;
            resourceInputs["file"] = (args ? args.file : undefined) ?? "Dockerfile";
            resourceInputs["platforms"] = args ? args.platforms : undefined;
            resourceInputs["pull"] = args ? args.pull : undefined;
            resourceInputs["registries"] = args ? args.registries : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["contextHash"] = undefined /*out*/;
            resourceInputs["manifests"] = undefined /*out*/;
        } else {
            resourceInputs["buildArgs"] = undefined /*out*/;
            resourceInputs["cacheFrom"] = undefined /*out*/;
            resourceInputs["cacheTo"] = undefined /*out*/;
            resourceInputs["context"] = undefined /*out*/;
            resourceInputs["contextHash"] = undefined /*out*/;
            resourceInputs["exports"] = undefined /*out*/;
            resourceInputs["file"] = undefined /*out*/;
            resourceInputs["manifests"] = undefined /*out*/;
            resourceInputs["platforms"] = undefined /*out*/;
            resourceInputs["pull"] = undefined /*out*/;
            resourceInputs["registries"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Image.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Image resource.
 */
export interface ImageArgs {
    /**
     *
     * An optional map of named build-time argument variables to set during
     * the Docker build. This flag allows you to pass build-time variables that
     * can be accessed like environment variables inside the RUN
     * instruction.
     */
    buildArgs?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     *
     * External cache sources (e.g., "user/app:cache", "type=local,src=path/to/dir")
     */
    cacheFrom?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     *
     * Cache export destinations (e.g., "user/app:cache", "type=local,dest=path/to/dir")
     */
    cacheTo?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     *
     * Contexts to use while building the image. If omitted, an empty context
     * is used. If more than one value is specified, they should be of the
     * form "name=value".
     */
    context?: pulumi.Input<string>;
    /**
     *
     * Name and optionally a tag (format: "name:tag"). If outputting to a
     * registry, the name should include the fully qualified registry address.
     */
    exports?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     *
     * Name of the Dockerfile to use (default: "$PATH/Dockerfile").
     */
    file?: pulumi.Input<string>;
    /**
     *
     * Set target platforms for the build. Defaults to the host's platform
     */
    platforms?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     *
     * Always attempt to pull all referenced images
     */
    pull?: pulumi.Input<boolean>;
    /**
     *
     * Logins for registry outputs
     */
    registries?: pulumi.Input<pulumi.Input<inputs.buildx.RegistryAuth>[]>;
    /**
     *
     * Name and optionally a tag (format: "name:tag"). If outputting to a
     * registry, the name should include the fully qualified registry address.
     */
    tags: pulumi.Input<pulumi.Input<string>[]>;
}
