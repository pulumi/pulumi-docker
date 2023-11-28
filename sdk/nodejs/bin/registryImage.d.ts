import * as pulumi from "@pulumi/pulumi";
/**
 * <!-- Bug: Type and Name are switched -->
 * Manages the lifecycle of docker image in a registry. You can upload images to a registry (= `docker push`) and also delete them again
 *
 * ## Example Usage
 *
 * Build an image with the `docker.RemoteImage` resource and then push it to a registry:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as docker from "@pulumi/docker";
 *
 * const helloworld = new docker.RegistryImage("helloworld", {keepRemotely: true});
 * const image = new docker.RemoteImage("image", {
 *     name: "registry.com/somename:1.0",
 *     build: {
 *         context: `${path.cwd}/absolutePathToContextFolder`,
 *     },
 * });
 * ```
 */
export declare class RegistryImage extends pulumi.CustomResource {
    /**
     * Get an existing RegistryImage resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RegistryImageState, opts?: pulumi.CustomResourceOptions): RegistryImage;
    /**
     * Returns true if the given object is an instance of RegistryImage.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    static isInstance(obj: any): obj is RegistryImage;
    /**
     * If `true`, the verification of TLS certificates of the server/registry is disabled. Defaults to `false`
     */
    readonly insecureSkipVerify: pulumi.Output<boolean | undefined>;
    /**
     * If true, then the Docker image won't be deleted on destroy operation. If this is false, it will delete the image from the docker registry on destroy operation. Defaults to `false`
     */
    readonly keepRemotely: pulumi.Output<boolean | undefined>;
    /**
     * The name of the Docker image.
     */
    readonly name: pulumi.Output<string>;
    /**
     * The sha256 digest of the image.
     */
    readonly sha256Digest: pulumi.Output<string>;
    /**
     * A map of arbitrary strings that, when changed, will force the `docker.RegistryImage` resource to be replaced. This can be used to repush a local image
     */
    readonly triggers: pulumi.Output<{
        [key: string]: any;
    } | undefined>;
    /**
     * Create a RegistryImage resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: RegistryImageArgs, opts?: pulumi.CustomResourceOptions);
}
/**
 * Input properties used for looking up and filtering RegistryImage resources.
 */
export interface RegistryImageState {
    /**
     * If `true`, the verification of TLS certificates of the server/registry is disabled. Defaults to `false`
     */
    insecureSkipVerify?: pulumi.Input<boolean>;
    /**
     * If true, then the Docker image won't be deleted on destroy operation. If this is false, it will delete the image from the docker registry on destroy operation. Defaults to `false`
     */
    keepRemotely?: pulumi.Input<boolean>;
    /**
     * The name of the Docker image.
     */
    name?: pulumi.Input<string>;
    /**
     * The sha256 digest of the image.
     */
    sha256Digest?: pulumi.Input<string>;
    /**
     * A map of arbitrary strings that, when changed, will force the `docker.RegistryImage` resource to be replaced. This can be used to repush a local image
     */
    triggers?: pulumi.Input<{
        [key: string]: any;
    }>;
}
/**
 * The set of arguments for constructing a RegistryImage resource.
 */
export interface RegistryImageArgs {
    /**
     * If `true`, the verification of TLS certificates of the server/registry is disabled. Defaults to `false`
     */
    insecureSkipVerify?: pulumi.Input<boolean>;
    /**
     * If true, then the Docker image won't be deleted on destroy operation. If this is false, it will delete the image from the docker registry on destroy operation. Defaults to `false`
     */
    keepRemotely?: pulumi.Input<boolean>;
    /**
     * The name of the Docker image.
     */
    name?: pulumi.Input<string>;
    /**
     * A map of arbitrary strings that, when changed, will force the `docker.RegistryImage` resource to be replaced. This can be used to repush a local image
     */
    triggers?: pulumi.Input<{
        [key: string]: any;
    }>;
}
