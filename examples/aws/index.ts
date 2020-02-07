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

import * as docker from "@pulumi/docker";
import * as aws from "@pulumi/aws";

function getImageRegistry(repo: aws.ecr.Repository) {
    return repo.registryId.apply(async registryId => {
        if (!registryId) {
            throw new Error("Expected registry ID to be defined during push");
        }
        const credentials = await aws.ecr.getCredentials({ registryId: registryId });
        const decodedCredentials = Buffer.from(credentials.authorizationToken, "base64").toString();
        const [username, password] = decodedCredentials.split(":");
        if (!password || !username) {
            throw new Error("Invalid credentials");
        }
        return {
            server: credentials.proxyEndpoint,
            username: username,
            password: password,
        };
    });
}

const ecr1 = new aws.ecr.Repository("build-cached");
const image1 = new docker.Image("build-cached", {
    imageName: ecr1.repositoryUrl,
    build: {
        context: "./app",
        cacheFrom: true,
        env: { DOCKER_BUILDKIT: "1" },
    },
    registry: getImageRegistry(ecr1),
});

const ecr2 = new aws.ecr.Repository("build-multistage");
const image2 = new docker.Image("build-multistage", {
    imageName: ecr2.repositoryUrl,
    build: {
        context: "./app",
        dockerfile: "./app/Dockerfile-multistage",
        cacheFrom: { stages: ["build"] },
    },
    registry: getImageRegistry(ecr2),
});

const ecr3 = new aws.ecr.Repository("path-example");
const image3 = new docker.Image("path-example", {
    imageName: ecr3.repositoryUrl,
    build: "./app",
    registry: getImageRegistry(ecr3),
});

export const image1Name = image1.imageName;
export const image2Name = image2.imageName;
export const image3Name = image3.imageName;

export const image4Name = docker.buildAndPushImage(
    "test-name", "./app", /*repositoryUrl:*/ undefined, /*logResource:*/ undefined);
