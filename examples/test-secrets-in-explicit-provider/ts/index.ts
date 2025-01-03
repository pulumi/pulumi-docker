import * as pulumi from "@pulumi/pulumi";
import * as docker from "@pulumi/docker";
import * as random from "@pulumi/random";

const providerWithSecretAddress = new docker.Provider("provider-with-sensitive-address", {
    registryAuth: [{
        address: pulumi.secret("secret-address"),
        username: "some-user",
    }],
})

const passwordResource = new random.RandomPassword("random", {
    length: 16,
    special: false,
});

const providerWithSecretUsername = new docker.Provider("provider-with-sensitive-username", {
    registryAuth: [{
        address: "some-address",
        username: passwordResource.result,
    }],
})

const providerWithSecretPassword = new docker.Provider("provider-with-password", {
    registryAuth: [{
        address: "some-address",
        username: "some-user",
        password: "secret-password",
    }],
})

export const randomPassword = "secret-password-" + Math.random().toString(36).slice(2, 7);
const providerWithRandomPassword = new docker.Provider("provider-with-random-password", {
    registryAuth: [{
        address: "some-address",
        username: passwordResource.result,
        password: randomPassword,
    }],
})

export const password = pulumi.unsecret(passwordResource.result)
    .apply(x => Buffer.from(x).toString('base64'));
