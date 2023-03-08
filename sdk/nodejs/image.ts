// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

/**
 * Builds a Docker Image and pushes to a Docker registry.
 *
 * ## Example Usage
 * ### A Docker image build
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as docker from "@pulumi/docker";
 *
 * const demoImage = new docker.Image("demo-image", {
 *     build: {
 *         context: ".",
 *         dockerfile: "Dockerfile",
 *     },
 *     imageName: "username/image:tag1",
 *     skipPush: true,
 * });
 * export const imageName = demoImage.imageName;
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
    public static readonly __pulumiType = 'docker:index/image:Image';

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
     * The fully qualified image name that was pushed to the registry.
     */
    public /*out*/ readonly baseImageName!: pulumi.Output<string | undefined>;
    /**
     * The fully qualified image name
     */
    public readonly imageName!: pulumi.Output<string | undefined>;
    /**
     * The name of the registry server hosting the image.
     */
    public /*out*/ readonly registryServer!: pulumi.Output<string | undefined>;

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
            if ((!args || args.imageName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'imageName'");
            }
            resourceInputs["build"] = args ? (args.build ? pulumi.output(args.build).apply(inputs.dockerBuildProvideDefaults) : undefined) : undefined;
            resourceInputs["imageName"] = args ? args.imageName : undefined;
            resourceInputs["registry"] = args ? args.registry : undefined;
            resourceInputs["skipPush"] = (args ? args.skipPush : undefined) ?? false;
            resourceInputs["baseImageName"] = undefined /*out*/;
            resourceInputs["registryServer"] = undefined /*out*/;
        } else {
            resourceInputs["baseImageName"] = undefined /*out*/;
            resourceInputs["imageName"] = undefined /*out*/;
            resourceInputs["registryServer"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "docker:image:Image" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(Image.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Image resource.
 */
export interface ImageArgs {
    /**
     * The Docker build context
     */
    build?: pulumi.Input<inputs.DockerBuild>;
    /**
     * The image name
     */
    imageName: pulumi.Input<string>;
    /**
     * The registry to push the image to
     */
    registry?: pulumi.Input<inputs.Registry>;
    /**
     * A flag to skip a registry push.
     */
    skipPush?: pulumi.Input<boolean>;
}
