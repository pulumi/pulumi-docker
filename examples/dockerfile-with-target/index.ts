import * as docker from "@pulumi/docker";

const myDependenciesImage = new docker.Image("my-image", {
    imageName: "pulumi-user/example:v1.0.0",
    build: {
        target: "dependencies",
        args: { "TEST_ARG": "42" },
    },
    skipPush: true,
});

export const depsImage = myDependenciesImage.imageName;
