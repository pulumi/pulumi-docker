import * as pulumi from "@pulumi/pulumi";
/**
 * `docker.RemoteImage` provides details about a specific Docker Image which need to be presend on the Docker Host
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as docker from "@pulumi/docker";
 *
 * const latest = docker.getRemoteImage({
 *     name: "nginx",
 * });
 * const specific = docker.getRemoteImage({
 *     name: "nginx:1.17.6",
 * });
 * const digest = docker.getRemoteImage({
 *     name: "nginx@sha256:36b74457bccb56fbf8b05f79c85569501b721d4db813b684391d63e02287c0b2",
 * });
 * const tagAndDigest = docker.getRemoteImage({
 *     name: "nginx:1.19.1@sha256:36b74457bccb56fbf8b05f79c85569501b721d4db813b684391d63e02287c0b2",
 * });
 * ```
 */
export declare function getRemoteImage(args: GetRemoteImageArgs, opts?: pulumi.InvokeOptions): Promise<GetRemoteImageResult>;
/**
 * A collection of arguments for invoking getRemoteImage.
 */
export interface GetRemoteImageArgs {
    /**
     * The name of the Docker image, including any tags or SHA256 repo digests.
     */
    name: string;
}
/**
 * A collection of values returned by getRemoteImage.
 */
export interface GetRemoteImageResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The name of the Docker image, including any tags or SHA256 repo digests.
     */
    readonly name: string;
    /**
     * The image sha256 digest in the form of `repo[:tag]@sha256:<hash>`. It may be empty in the edge case where the local image was pulled from a repo, tagged locally, and then referred to in the data source by that local name/tag.
     */
    readonly repoDigest: string;
}
/**
 * `docker.RemoteImage` provides details about a specific Docker Image which need to be presend on the Docker Host
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as docker from "@pulumi/docker";
 *
 * const latest = docker.getRemoteImage({
 *     name: "nginx",
 * });
 * const specific = docker.getRemoteImage({
 *     name: "nginx:1.17.6",
 * });
 * const digest = docker.getRemoteImage({
 *     name: "nginx@sha256:36b74457bccb56fbf8b05f79c85569501b721d4db813b684391d63e02287c0b2",
 * });
 * const tagAndDigest = docker.getRemoteImage({
 *     name: "nginx:1.19.1@sha256:36b74457bccb56fbf8b05f79c85569501b721d4db813b684391d63e02287c0b2",
 * });
 * ```
 */
export declare function getRemoteImageOutput(args: GetRemoteImageOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRemoteImageResult>;
/**
 * A collection of arguments for invoking getRemoteImage.
 */
export interface GetRemoteImageOutputArgs {
    /**
     * The name of the Docker image, including any tags or SHA256 repo digests.
     */
    name: pulumi.Input<string>;
}
