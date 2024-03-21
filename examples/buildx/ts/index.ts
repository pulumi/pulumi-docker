import * as pulumi from "@pulumi/pulumi";
import * as docker from "@pulumi/docker";

const config = new pulumi.Config();
const dockerHubPassword = config.require("dockerHubPassword");
const multiPlatform = new docker.buildx.Image("multiPlatform", {
    dockerfile: {
        location: "app/Dockerfile.multiPlatform",
    },
    context: {
        location: "app",
    },
    platforms: [
        "plan9/amd64",
        "plan9/386",
    ],
});
const registryPush = new docker.buildx.Image("registryPush", {
    context: {
        location: "app",
    },
    tags: ["docker.io/pulumibot/buildkit-e2e:example"],
    exports: [{
        registry: {
            ociMediaTypes: true,
            push: false,
        },
    }],
    registries: [{
        address: "docker.io",
        username: "pulumibot",
        password: dockerHubPassword,
    }],
});
const cached = new docker.buildx.Image("cached", {
    context: {
        location: "app",
    },
    cacheTo: [{
        local: {
            dest: "tmp/cache",
            mode: "max",
        },
    }],
    cacheFrom: [{
        local: {
            src: "tmp/cache",
        },
    }],
});
const buildArgs = new docker.buildx.Image("buildArgs", {
    dockerfile: {
        location: "app/Dockerfile.buildArgs",
    },
    context: {
        location: "app",
    },
    buildArgs: {
        SET_ME_TO_TRUE: "true",
    },
});
const extraHosts = new docker.buildx.Image("extraHosts", {
    dockerfile: {
        location: "app/Dockerfile.extraHosts",
    },
    context: {
        location: "app",
    },
    addHosts: ["metadata.google.internal:169.254.169.254"],
});
const sshMount = new docker.buildx.Image("sshMount", {
    dockerfile: {
        location: "app/Dockerfile.sshMount",
    },
    context: {
        location: "app",
    },
    ssh: [{
        id: "default",
    }],
});
const secrets = new docker.buildx.Image("secrets", {
    dockerfile: {
        location: "app/Dockerfile.secrets",
    },
    context: {
        location: "app",
    },
    secrets: {
        password: "hunter2",
    },
});
const labels = new docker.buildx.Image("labels", {
    context: {
        location: "app",
    },
    labels: {
        description: "This image will get a descriptive label 👍",
    },
});
const target = new docker.buildx.Image("target", {
    dockerfile: {
        location: "app/Dockerfile.target",
    },
    context: {
        location: "app",
    },
    target: "build-me",
});
const namedContexts = new docker.buildx.Image("namedContexts", {
    dockerfile: {
        location: "app/Dockerfile.namedContexts",
    },
    context: {
        location: "app",
        named: {
            "golang:latest": {
                location: "docker-image://golang@sha256:b8e62cf593cdaff36efd90aa3a37de268e6781a2e68c6610940c48f7cdf36984",
            },
        },
    },
});
const remoteContext = new docker.buildx.Image("remoteContext", {context: {
    location: "https://raw.githubusercontent.com/pulumi/pulumi-docker/api-types/provider/testdata/Dockerfile",
}});
const remoteContextWithInline = new docker.buildx.Image("remoteContextWithInline", {
    dockerfile: {
        inline: `FROM busybox
COPY hello.c ./
`,
    },
    context: {
        location: "https://github.com/docker-library/hello-world.git",
    },
});
const inline = new docker.buildx.Image("inline", {
    dockerfile: {
        inline: `FROM alpine
RUN echo "This uses an inline Dockerfile! 👍"
`,
    },
    context: {
        location: "app",
    },
});
const dockerLoad = new docker.buildx.Image("dockerLoad", {
    context: {
        location: "app",
    },
    exports: [{
        docker: {
            tar: true,
        },
    }],
});
export const platforms = multiPlatform.platforms;
