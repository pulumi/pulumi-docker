import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
/**
 * <!-- Bug: Type and Name are switched -->
 * Pulls a Docker image to a given Docker host from a Docker Registry.
 *  This resource will *not* pull new layers of the image automatically unless used in conjunction with docker.RegistryImage data source to update the `pullTriggers` field.
 *
 * ## Example Usage
 * ### Basic
 *
 * Finds and downloads the latest `ubuntu:precise` image but does not check
 * for further updates of the image
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as docker from "@pulumi/docker";
 *
 * const ubuntu = new docker.RemoteImage("ubuntu", {name: "ubuntu:precise"});
 * ```
 * ### Dynamic updates
 *
 * To be able to update an image dynamically when the `sha256` sum changes,
 * you need to use it in combination with `docker.RegistryImage` as follows:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as docker from "@pulumi/docker";
 *
 * const ubuntuRegistryImage = docker.getRegistryImage({
 *     name: "ubuntu:precise",
 * });
 * const ubuntuRemoteImage = new docker.RemoteImage("ubuntuRemoteImage", {
 *     name: ubuntuRegistryImage.then(ubuntuRegistryImage => ubuntuRegistryImage.name),
 *     pullTriggers: [ubuntuRegistryImage.then(ubuntuRegistryImage => ubuntuRegistryImage.sha256Digest)],
 * });
 * ```
 */
export declare class RemoteImage extends pulumi.CustomResource {
    /**
     * Get an existing RemoteImage resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RemoteImageState, opts?: pulumi.CustomResourceOptions): RemoteImage;
    /**
     * Returns true if the given object is an instance of RemoteImage.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    static isInstance(obj: any): obj is RemoteImage;
    /**
     * Configuration to build an image. Please see [docker build command reference](https://docs.docker.com/engine/reference/commandline/build/#options) too.
     */
    readonly build: pulumi.Output<outputs.RemoteImageBuild | undefined>;
    /**
     * Always remove intermediate containers
     */
    readonly forceRemove: pulumi.Output<boolean | undefined>;
    /**
     * The ID of the image (as seen when executing `docker inspect` on the image). Can be used to reference the image via its ID in other resources.
     */
    readonly imageId: pulumi.Output<string>;
    /**
     * If true, then the Docker image won't be deleted on destroy operation. If this is false, it will delete the image from the docker local storage on destroy operation.
     */
    readonly keepLocally: pulumi.Output<boolean | undefined>;
    /**
     * type of ulimit, e.g. `nofile`
     */
    readonly name: pulumi.Output<string>;
    /**
     * Set platform if server is multi-platform capable
     */
    readonly platform: pulumi.Output<string | undefined>;
    /**
     * List of values which cause an image pull when changed. This is used to store the image digest from the registry when using the docker*registry*image.
     */
    readonly pullTriggers: pulumi.Output<string[] | undefined>;
    /**
     * The image sha256 digest in the form of `repo[:tag]@sha256:<hash>`.
     */
    readonly repoDigest: pulumi.Output<string>;
    /**
     * A map of arbitrary strings that, when changed, will force the `docker.RemoteImage` resource to be replaced. This can be used to rebuild an image when contents of source code folders change
     */
    readonly triggers: pulumi.Output<{
        [key: string]: any;
    } | undefined>;
    /**
     * Create a RemoteImage resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RemoteImageArgs, opts?: pulumi.CustomResourceOptions);
}
/**
 * Input properties used for looking up and filtering RemoteImage resources.
 */
export interface RemoteImageState {
    /**
     * Configuration to build an image. Please see [docker build command reference](https://docs.docker.com/engine/reference/commandline/build/#options) too.
     */
    build?: pulumi.Input<inputs.RemoteImageBuild>;
    /**
     * Always remove intermediate containers
     */
    forceRemove?: pulumi.Input<boolean>;
    /**
     * The ID of the image (as seen when executing `docker inspect` on the image). Can be used to reference the image via its ID in other resources.
     */
    imageId?: pulumi.Input<string>;
    /**
     * If true, then the Docker image won't be deleted on destroy operation. If this is false, it will delete the image from the docker local storage on destroy operation.
     */
    keepLocally?: pulumi.Input<boolean>;
    /**
     * type of ulimit, e.g. `nofile`
     */
    name?: pulumi.Input<string>;
    /**
     * Set platform if server is multi-platform capable
     */
    platform?: pulumi.Input<string>;
    /**
     * List of values which cause an image pull when changed. This is used to store the image digest from the registry when using the docker*registry*image.
     */
    pullTriggers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The image sha256 digest in the form of `repo[:tag]@sha256:<hash>`.
     */
    repoDigest?: pulumi.Input<string>;
    /**
     * A map of arbitrary strings that, when changed, will force the `docker.RemoteImage` resource to be replaced. This can be used to rebuild an image when contents of source code folders change
     */
    triggers?: pulumi.Input<{
        [key: string]: any;
    }>;
}
/**
 * The set of arguments for constructing a RemoteImage resource.
 */
export interface RemoteImageArgs {
    /**
     * Configuration to build an image. Please see [docker build command reference](https://docs.docker.com/engine/reference/commandline/build/#options) too.
     */
    build?: pulumi.Input<inputs.RemoteImageBuild>;
    /**
     * Always remove intermediate containers
     */
    forceRemove?: pulumi.Input<boolean>;
    /**
     * If true, then the Docker image won't be deleted on destroy operation. If this is false, it will delete the image from the docker local storage on destroy operation.
     */
    keepLocally?: pulumi.Input<boolean>;
    /**
     * type of ulimit, e.g. `nofile`
     */
    name: pulumi.Input<string>;
    /**
     * Set platform if server is multi-platform capable
     */
    platform?: pulumi.Input<string>;
    /**
     * List of values which cause an image pull when changed. This is used to store the image digest from the registry when using the docker*registry*image.
     */
    pullTriggers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A map of arbitrary strings that, when changed, will force the `docker.RemoteImage` resource to be replaced. This can be used to rebuild an image when contents of source code folders change
     */
    triggers?: pulumi.Input<{
        [key: string]: any;
    }>;
}
