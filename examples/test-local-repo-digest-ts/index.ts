import * as docker from "@pulumi/docker";

const localImg = new docker.Image("my-image", {
    imageName: "pulumi-bot/local-repo-digest",
    skipPush:true,
    build: {
        platform: "linux/amd64",
    }
});

export const imageName = localImg.imageName;

export const repoDigest = localImg.repoDigest

