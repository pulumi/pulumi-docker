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
import { buildAndPushImage, DockerBuild, Registry } from "./docker";

/**
 * Arguments for constructing an Image resource.
 */
export interface ImageArgs {
    /**
     * The qualified image name that will be pushed to the remote registry.  Must be a supported image name for the
     * target registry user.
     */
    imageName: pulumi.Input<string>;
    /**
     * The Docker build context, as a folder path or a detailed DockerBuild object.
     */
    build: pulumi.Input<string | DockerBuild>;
    /**
     * The docker image name to build locally before tagging with imageName.
     */
    localImageName?: pulumi.Input<string>;
    /**
     * Docker registry to push to.
     */
    registry: pulumi.Input<ImageRegistry>;
}

export interface ImageRegistry {
    /**
     * Docker registry to push to.
     */
    name: pulumi.Input<string>;
    /**
     * Username for login to the target Docker registry.
     */
    username: pulumi.Input<string>;
    /**
     * Password for login to the target Docker registry.
     */
    password: pulumi.Input<string>;
}

/**
 * A docker.Image resource represents a Docker image built locally which is published and made available via a remote
 * Docker registry.  This can be used to ensure that a Docker source directory from a local deployment environment is
 * built and pushed to a cloud-hosted Docker registry as part of a Pulumi deployment, so that it can be referenced as an
 * image input from other cloud services that reference Docker images - including Kubernetes Pods, AWS ECS Tasks, and
 * Azure Container Instances.
 */
export class Image extends pulumi.ComponentResource {
    /**
     * The raw image name that was built and pushed.  This does now include the digest annotation, so is not pinned to
     * the specific build performed by this docker.Image.
     */
    public rawImageName: pulumi.Output<string>;
    /**
     * The pinned image name, including digest annotation.
     */
    public imageName: pulumi.Output<string>;
    /**
     * The built image digest.
     */
    public digest: pulumi.Output<string>;

    constructor(name: string, args: ImageArgs, opts?: pulumi.ComponentResourceOptions) {
        super("docker:image:Image", name, args, opts);

        const imageDigest =
            pulumi
            .all([args.imageName, args.build, args.localImageName, args.registry])
            .apply(([imageName, build, localImageName, registry]) =>
                pulumi
                .all([registry.name, registry.username, registry.password])
                .apply(([registryName, username, password]) => {
                    if (!localImageName) {
                        localImageName = imageName;
                    }
                    return buildAndPushImage(localImageName, build, imageName, this, async () => {
                        return {
                            registry: registryName,
                            username: username,
                            password: password,
                        };
                    });
                }));

        this.digest = imageDigest;
        this.rawImageName = pulumi.output(args.imageName);
        this.imageName =
            pulumi
            .all([args.imageName, imageDigest])
            .apply(([imageName, digest]) => `${imageName}@${digest}`);
        this.registerOutputs({
            rawImageName: this.rawImageName,
            imageName: this.imageName,
            digest: this.digest,
        });
    }
}
