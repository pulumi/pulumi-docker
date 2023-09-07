// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

/**
 * `Image` builds a Docker image and pushes it Docker and OCI compatible registries.
 * This resource enables running Docker builds as part of a Pulumi deployment.
 *
 * Note: This resource does not delete tags, locally or remotely, when destroyed.
 *
 * ## Image name
 *
 * The Image resource uses `imageName` to refer to a fully qualified Docker image name, by the format `repository:tag`.
 * Note that this does not include any digest information and thus will not cause any updates when passed to dependencies,
 * even when using `latest` tag. To trigger such updates, e.g. when referencing pushed images in container orchestration
 * and management resources, please use the `repoDigest` Output instead, which is of the format
 * `repository@<algorithm>:<hash>` and unique per build/push.
 * Note that `repoDigest` is not available for local Images. For a local Image not pushed to a registry, you may want to
 * give `imageName` a unique tag per pulumi update.
 *
 * ## Cross-platform builds
 *
 * The Image resource supports cross-platform builds when the [Docker engine has cross-platform support enabled via emulators](https://docs.docker.com/build/building/multi-platform/#building-multi-platform-images).
 * The Image resource currently supports providing only a single operating system and architecture in the `platform` field, e.g.: `linux/amd64`.
 * To enable this support, you may need to install the emulators in the environment running your Pulumi program.
 *
 * If you are using Linux, you may be using Docker Engine or Docker Desktop for Linux, depending on how you have installed Docker. The [FAQ for Docker Desktop for Linux](https://docs.docker.com/desktop/faqs/linuxfaqs/#context) describes the differences and how to select which Docker context is in use.
 *
 * * For local development using Docker Desktop, this is enabled by default.
 * * For systems using Docker Engine, install the QEMU binaries and register them with using the docker image from [github.com/tonistiigi/binfmt](https://github.com/tonistiigi/binfmt):
 * * In a GitHub Actions workflow, the [docker/setup-qemu-action](https://github.com/docker/setup-qemu-action) can be used instead by adding this step to your workflow file. Example workflow usage:
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
 *         args: {
 *             platform: "linux/amd64",
 *         },
 *         context: ".",
 *         dockerfile: "Dockerfile",
 *     },
 *     imageName: "username/image:tag1",
 *     skipPush: true,
 * });
 * export const imageName = demoImage.imageName;
 * ```
 * ### A Docker image build and push
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as docker from "@pulumi/docker";
 *
 * const demoPushImage = new docker.Image("demo-push-image", {
 *     build: {
 *         context: ".",
 *         dockerfile: "Dockerfile",
 *     },
 *     imageName: "docker.io/username/push-image:tag1",
 * });
 * export const imageName = demoPushImage.imageName;
 * export const repoDigest = demoPushImage.repoDigest;
 * ```
 * ### Docker image build using caching with AWS Elastic Container Registry
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * import * as docker from "@pulumi/docker";
 *
 * const ecrRepository = new aws.ecr.Repository("ecr-repository", {name: "docker-repository"});
 * const authToken = aws.ecr.getAuthorizationTokenOutput({
 *     registryId: ecrRepository.registryId,
 * });
 * const myAppImage = new docker.Image("my-app-image", {
 *     build: {
 *         args: {
 *             BUILDKIT_INLINE_CACHE: "1",
 *         },
 *         cacheFrom: {
 *             images: [pulumi.interpolate`${ecrRepository.repositoryUrl}:latest`],
 *         },
 *         context: "app/",
 *         dockerfile: "Dockerfile",
 *     },
 *     imageName: pulumi.interpolate`${ecrRepository.repositoryUrl}:latest`,
 *     registry: {
 *         password: pulumi.secret(authToken.apply(authToken => authToken.password)),
 *         server: ecrRepository.repositoryUrl,
 *     },
 * });
 * export const imageName = myAppImage.imageName;
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
    public /*out*/ readonly baseImageName!: pulumi.Output<string>;
    /**
     * The path to the build context to use.
     */
    public /*out*/ readonly context!: pulumi.Output<string>;
    /**
     * The location of the Dockerfile relative to the docker build context.
     */
    public /*out*/ readonly dockerfile!: pulumi.Output<string>;
    /**
     * The fully qualified image name
     */
    public readonly imageName!: pulumi.Output<string>;
    /**
     * The name of the registry server hosting the image.
     */
    public /*out*/ readonly registryServer!: pulumi.Output<string>;
    /**
     * **For pushed images:**
     * The manifest digest of an image pushed to a registry, of the format repository@<algorithm>:<hash>, e.g. `username/demo-image@sha256:a6ae6dd8d39c5bb02320e41abf00cd4cb35905fec540e37d306c878be8d38bd3`.
     * This reference is unique per image build and push. 
     * Only available for images pushed to a registry.
     * Use when passing a reference to a pushed image to container management resources.
     *
     * **Local-only images**For local images, this field is the image ID of the built local image, of the format <algorithm>:<hash>, e.g `sha256:826a130323165bb0ccb0374ae774f885c067a951b51a6ee133577f4e5dbc4119` 
     */
    public /*out*/ readonly repoDigest!: pulumi.Output<string>;

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
            resourceInputs["context"] = undefined /*out*/;
            resourceInputs["dockerfile"] = undefined /*out*/;
            resourceInputs["registryServer"] = undefined /*out*/;
            resourceInputs["repoDigest"] = undefined /*out*/;
        } else {
            resourceInputs["baseImageName"] = undefined /*out*/;
            resourceInputs["context"] = undefined /*out*/;
            resourceInputs["dockerfile"] = undefined /*out*/;
            resourceInputs["imageName"] = undefined /*out*/;
            resourceInputs["registryServer"] = undefined /*out*/;
            resourceInputs["repoDigest"] = undefined /*out*/;
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
     * The image name, of the format repository[:tag], e.g. `docker.io/username/demo-image:v1`.
     * This reference is not unique to each build and push.For the unique manifest SHA of a pushed docker image, or the local image ID, please use `repoDigest`.
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
