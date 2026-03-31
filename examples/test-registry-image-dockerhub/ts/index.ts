import * as docker from "@pulumi/docker";
import * as pulumi from "@pulumi/pulumi";

const config = new pulumi.Config();
const username = config.require("dockerUsername");
const password = config.requireSecret("dockerPassword");

const provider = new docker.Provider("docker-hub", {
    registryAuth: [{
        address: "registry-1.docker.io",
        username: username,
        password: password,
    }],
});

const registryImage = docker.getRegistryImageOutput({
    name: "postgres:18",
}, { provider });

export const digest = registryImage.sha256Digest;
