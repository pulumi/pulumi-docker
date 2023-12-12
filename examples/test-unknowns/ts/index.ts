import * as pulumi from "@pulumi/pulumi";
import * as docker from "@pulumi/docker";
import * as command from "@pulumi/command";



const randArg = new command.local.Command ("arg", {
   create: "echo setMyArg"
});

const img = new docker.Image("docker-565-one", {
    imageName: "pulumibot/test-image:with-build-args",

    build: {
        args: {
            "RANDOM_ARG": randArg.stdout
        },
    },
    skipPush: true,
});

export const randArgument = randArg.stdout
