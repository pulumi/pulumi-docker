import * as pulumi from "@pulumi/pulumi";
/**
 * Creates a docker tag. It has the exact same functionality as the `docker tag` command. Deleting the resource will neither delete the source nor target images. The source image must exist on the machine running the docker daemon.
 */
export declare class Tag extends pulumi.CustomResource {
    /**
     * Get an existing Tag resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TagState, opts?: pulumi.CustomResourceOptions): Tag;
    /**
     * Returns true if the given object is an instance of Tag.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    static isInstance(obj: any): obj is Tag;
    /**
     * Name of the source image.
     */
    readonly sourceImage: pulumi.Output<string>;
    /**
     * ImageID of the source image in the format of `sha256:<<ID>>`
     */
    readonly sourceImageId: pulumi.Output<string>;
    /**
     * Name of the target image.
     */
    readonly targetImage: pulumi.Output<string>;
    /**
     * Create a Tag resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TagArgs, opts?: pulumi.CustomResourceOptions);
}
/**
 * Input properties used for looking up and filtering Tag resources.
 */
export interface TagState {
    /**
     * Name of the source image.
     */
    sourceImage?: pulumi.Input<string>;
    /**
     * ImageID of the source image in the format of `sha256:<<ID>>`
     */
    sourceImageId?: pulumi.Input<string>;
    /**
     * Name of the target image.
     */
    targetImage?: pulumi.Input<string>;
}
/**
 * The set of arguments for constructing a Tag resource.
 */
export interface TagArgs {
    /**
     * Name of the source image.
     */
    sourceImage: pulumi.Input<string>;
    /**
     * Name of the target image.
     */
    targetImage: pulumi.Input<string>;
}
