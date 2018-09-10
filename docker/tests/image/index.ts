// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as aws from "@pulumi/aws";
import * as azure from "@pulumi/azure";
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
        server: "docker.io",
        username: dockerUsername,
        password: dockerPassword,
    },
});
export const dockerImageName = image.imageName;

// AWS ECR
const ecr = new aws.ecr.Repository("ecr");
const ecrCreds = ecr.registryId.apply(async (registryId) => {
    const credentials = await aws.ecr.getCredentials({
        registryId: registryId,
    });
    const decodedCredentials = Buffer.from(credentials.authorizationToken, "base64").toString();
    const [username, password] = decodedCredentials.split(":");
    return { server: credentials.proxyEndpoint, username, password };
});
const image2 = new docker.Image("mynginx2", {
    imageName: ecr.repositoryUrl,
    build: "./mynginx",
    registry: ecrCreds,
});
export const ecrImageName = image2.imageName;

// Azure ACR
const location = "westus";
const resourceGroupName = new azure.core.ResourceGroup("acrrg", { location }).name;
const storageAccountId = new azure.storage.Account("acrstorage", {
    resourceGroupName,
    location,
    accountTier: "Standard",
    accountReplicationType: "LRS",
}).id;
const acr = new azure.containerservice.Registry("acr", {
    resourceGroupName,
    location,
    storageAccountId,
    adminEnabled: true,
});
const image3 = new docker.Image("mynginx3", {
    imageName: acr.loginServer.apply(server => `${server}/mynginx`),
    build: "./mynginx",
    registry: {
        server: acr.loginServer,
        username: acr.adminUsername,
        password: acr.adminPassword,
    },
});
export const acrImage = image3.imageName;
