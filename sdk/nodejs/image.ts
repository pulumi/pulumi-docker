// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import * as pulumi from "@pulumi/pulumi";
import * as docker from "./docker";
import * as utils from "./utils";

/**
 * Arguments for constructing an Image resource.
 */
export interface ImageArgs {
    /**
     * The qualified image name that will be pushed to the remote registry.  Must be a supported
     * image name for the target registry user.  This name can include a tag at the end.  If
     * provided all pushed image resources will contain that tag as well.
     *
     * Either [imageName] or [localImageName] can have a tag.  However, if both have a tag, then
     * those tags must match.
     */
    imageName: string;
    /**
     * The Docker build context, as a folder path or a detailed DockerBuild object.
     */
    build: string | docker.DockerBuild;
    /**
     * The docker image name to build locally before tagging with imageName.  If not provided, it
     * will be given the value of to [imageName].  This name can include a tag at the end.  If
     * provided all pushed image resources will contain that tag as well.
     *
     * Either [imageName] or [localImageName] can have a tag.  However, if both have a tag, then
     * those tags must match.
     */
    localImageName?: string;
    /**
     * Credentials for the docker registry to push to.
     */
    registry?: ImageRegistry;
}

export interface ImageRegistry {
    /**
     * Docker registry server URL to push to.  Some common values include:
     * DockerHub: `docker.io` or `https://index.docker.io/v1`
     * Azure Container Registry: `<name>.azurecr.io`
     * AWS Elastic Container Registry: `<account>.dkr.ecr.us-east-2.amazonaws.com`
     * Google Container Registry: `<name>.gcr.io`
     */
    server: string;
    /**
     * Username for login to the target Docker registry.
     */
    username: string;
    /**
     * Password for login to the target Docker registry.
     */
    password: string;
}

/**
 * A docker.Image resource represents a Docker image built locally which is published and made
 * available via a remote Docker registry.  This can be used to ensure that a Docker source
 * directory from a local deployment environment is built and pushed to a cloud-hosted Docker
 * registry as part of a Pulumi deployment, so that it can be referenced as an image input from
 * other cloud services that reference Docker images - including Kubernetes Pods, AWS ECS Tasks, and
 * Azure Container Instances.
 */
export class Image extends pulumi.ComponentResource {
    /**
     * The base image name that was built and pushed.  This does not include the id annotation, so
     * is not pinned to the specific build performed by this docker.Image.
     */
    public baseImageName: pulumi.Output<string>;
    /**
     * The unique pinned image name on the remote repository.
     */
    public imageName: pulumi.Output<string>;
    /**
     * The server the image is located at.
     */
    public registryServer: pulumi.Output<string | undefined>;

    /** @deprecated This will have the same value as [imageName], but will be removed in the future. */
    public id: pulumi.Output<string>;

    /**
     * @deprecated This will have the same value as [imageName], but will be removed in the future.
     * It can be used to get a unique name for this specific image, but is not the actual repository
     * digest value.
     */
    public digest: pulumi.Output<string | undefined>;

    constructor(name: string, args: pulumi.WrappedObject<ImageArgs>, opts?: pulumi.ComponentResourceOptions) {
        super("docker:image:Image", name, {}, opts);

        const imageData = pulumi.output(args).apply(async (imageArgs) => {
            const imageName = imageArgs.imageName;

            // If there is no localImageName set it equal to imageName.  Note: this means
            // that if imageName contains a tag, localImageName will contain the same tag.
            const localImageName = imageArgs.localImageName || imageName;

            // Now break both the localImageName and the imageName into the untagged part and the
            // optional tag.  If both have tags, they must match.  If one or the other has a tag, we
            // just use that as the tag to use.  This allows users to flexibly provide a tag on one
            // option or the other and still have it work out.
            const { imageName: localImageNameWithoutTag, tag: localImageNameTag } = utils.getImageNameAndTag(localImageName);
            const { imageName: imageNameWithoutTag, tag: imageNameTag } = utils.getImageNameAndTag(imageName);

            const tag = localImageNameTag || imageNameTag;

            checkTag(localImageNameTag);
            checkTag(imageNameTag);

            // buildAndPushImageAsync expects only the baseImageName to have a tag.  So build that
            // name appropriately if we were given a tag.
            const baseImageName = tag ? `${localImageNameWithoutTag}:${tag}` : localImageName;

            // buildAndPushImageAsync does not want the repositoryUrl to have a tag.  This is just
            // the base url where the images will be pushed to.  All tagging will be taken care of
            // inside that api.
            const repositoryUrl = imageNameWithoutTag;

            const registry = imageArgs.registry;
            const uniqueTargetName = await docker.buildAndPushImageAsync(
                baseImageName,
                imageArgs.build,
                repositoryUrl,
                /*logResource:*/ this,
                registry && (async () => {
                    return {
                        registry: registry.server,
                        username: registry.username,
                        password: registry.password,
                    };
                }),
            );

            return { uniqueTargetName, registryServer: registry && registry.server };

            function checkTag(t: string | undefined) {
                if (t && (t !== tag)) {
                    throw new Error(`[localImageName] and [imageName] had mismatched tags.
    ${JSON.stringify(localImageNameTag)} !== ${JSON.stringify(imageNameTag)}`);
                }
            }
        });

        this.imageName = imageData.apply(d => d.uniqueTargetName);
        this.id = this.imageName;
        this.digest = this.imageName;
        this.registryServer = imageData.apply(d => d.registryServer);
        this.baseImageName = pulumi.output(args.imageName);

        this.registerOutputs({
            baseImageName: this.baseImageName,
            imageName: this.imageName,
            id: this.id,
            digest: this.digest,
            registryServer: this.registryServer,
        });
    }
}
