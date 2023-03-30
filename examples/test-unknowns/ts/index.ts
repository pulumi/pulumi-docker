import * as pulumi from "@pulumi/pulumi";
import * as docker from "@pulumi/docker";
import * as random from "@pulumi/random";



const randName = new random.RandomString("random", {
    length: 10,

});

const img = new docker.Image("docker-565-one", {
    imageName: "pulumibot/test-image:with-build-args",

    build: {
        args: {
            "RANDOM_ARG": randName.id
        },
    },
    skipPush: true,
});

export const randnameid = randName.id
