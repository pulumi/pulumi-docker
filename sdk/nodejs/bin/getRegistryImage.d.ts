import * as pulumi from "@pulumi/pulumi";
/**
 * Reads the image metadata from a Docker Registry. Used in conjunction with the docker.RemoteImage resource to keep an image up to date on the latest available version of the tag.
 *
 * ## Example Usage
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
export declare function getRegistryImage(args: GetRegistryImageArgs, opts?: pulumi.InvokeOptions): Promise<GetRegistryImageResult>;
/**
 * A collection of arguments for invoking getRegistryImage.
 */
export interface GetRegistryImageArgs {
    /**
     * If `true`, the verification of TLS certificates of the server/registry is disabled. Defaults to `false`
     */
    insecureSkipVerify?: boolean;
    /**
     * The name of the Docker image, including any tags. e.g. `alpine:latest`
     */
    name: string;
}
/**
 * A collection of values returned by getRegistryImage.
 */
export interface GetRegistryImageResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * If `true`, the verification of TLS certificates of the server/registry is disabled. Defaults to `false`
     */
    readonly insecureSkipVerify?: boolean;
    /**
     * The name of the Docker image, including any tags. e.g. `alpine:latest`
     */
    readonly name: string;
    /**
     * The content digest of the image, as stored in the registry.
     */
    readonly sha256Digest: string;
}
/**
 * Reads the image metadata from a Docker Registry. Used in conjunction with the docker.RemoteImage resource to keep an image up to date on the latest available version of the tag.
 *
 * ## Example Usage
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
export declare function getRegistryImageOutput(args: GetRegistryImageOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRegistryImageResult>;
/**
 * A collection of arguments for invoking getRegistryImage.
 */
export interface GetRegistryImageOutputArgs {
    /**
     * If `true`, the verification of TLS certificates of the server/registry is disabled. Defaults to `false`
     */
    insecureSkipVerify?: pulumi.Input<boolean>;
    /**
     * The name of the Docker image, including any tags. e.g. `alpine:latest`
     */
    name: pulumi.Input<string>;
}
