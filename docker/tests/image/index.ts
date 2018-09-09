// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as aws from "@pulumi/aws";
import * as docker from "@pulumi/docker";
import * as pulumi from "@pulumi/pulumi";

const config = new pulumi.Config();
const dockerUsername = config.require("dockerUsername");
const dockerPassword = config.require("dockerPassword");

// DockerHub
const image = new docker.Image("mynginx", {
    imageName: "hekul/mynginx:v1",
    build: "./mynginx",
    registry: {
        name: "docker.io",
        username: dockerUsername,
        password: dockerPassword,
    },
});
export const dockerImageName = image.imageName;


// AWS ECR
const ecr = new aws.ecr.Repository("ecr");
const creds = ecr.registryId.apply(async (registryId) => {
    const credentials = await aws.ecr.getCredentials({
        registryId: registryId,
    });
    const decodedCredentials = Buffer.from(credentials.authorizationToken, "base64").toString();
    const [username, password] = decodedCredentials.split(":");
    return { name: credentials.proxyEndpoint, username, password };
});
const image2 = new docker.Image("mynginx2", {
    imageName: ecr.repositoryUrl,
    build: "./mynginx",
    registry: creds,
});
export const ecrImageName = image2.imageName;
