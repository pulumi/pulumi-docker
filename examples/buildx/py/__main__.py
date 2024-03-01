import pulumi
import pulumi_docker as docker

config = pulumi.Config()
docker_hub_password = config.require("dockerHubPassword")
multi_platform = docker.buildx.Image("multiPlatform",
    dockerfile=docker.buildx.DockerfileArgs(
        location="app/Dockerfile.multiPlatform",
    ),
    context=docker.buildx.BuildContextArgs(
        location="app",
    ),
    platforms=[
        "plan9/amd64",
        "plan9/386",
    ])
registry_push = docker.buildx.Image("registryPush",
    context=docker.buildx.BuildContextArgs(
        location="app",
    ),
    tags=["docker.io/pulumibot/buildkit-e2e:example"],
    exports=[docker.buildx.ExportEntryArgs(
        registry=docker.buildx.ExportRegistryArgs(
            oci_media_types=True,
            push=False,
        ),
    )],
    registries=[docker.buildx.RegistryAuthArgs(
        address="docker.io",
        username="pulumibot",
        password=docker_hub_password,
    )])
cached = docker.buildx.Image("cached",
    context=docker.buildx.BuildContextArgs(
        location="app",
    ),
    cache_to=[docker.buildx.CacheToEntryArgs(
        local=docker.buildx.CacheToLocalArgs(
            dest="tmp/cache",
            mode="max",
        ),
    )],
    cache_from=[docker.buildx.CacheFromEntryArgs(
        local=docker.buildx.CacheFromLocalArgs(
            src="tmp/cache",
        ),
    )])
build_args = docker.buildx.Image("buildArgs",
    dockerfile=docker.buildx.DockerfileArgs(
        location="app/Dockerfile.buildArgs",
    ),
    context=docker.buildx.BuildContextArgs(
        location="app",
    ),
    build_args={
        "SET_ME_TO_TRUE": "true",
    })
extra_hosts = docker.buildx.Image("extraHosts",
    dockerfile=docker.buildx.DockerfileArgs(
        location="app/Dockerfile.extraHosts",
    ),
    context=docker.buildx.BuildContextArgs(
        location="app",
    ),
    add_hosts=["metadata.google.internal:169.254.169.254"])
ssh_mount = docker.buildx.Image("sshMount",
    dockerfile=docker.buildx.DockerfileArgs(
        location="app/Dockerfile.sshMount",
    ),
    context=docker.buildx.BuildContextArgs(
        location="app",
    ),
    ssh=[docker.buildx.SSHArgs(
        id="default",
    )])
secrets = docker.buildx.Image("secrets",
    dockerfile=docker.buildx.DockerfileArgs(
        location="app/Dockerfile.secrets",
    ),
    context=docker.buildx.BuildContextArgs(
        location="app",
    ),
    secrets={
        "password": "hunter2",
    })
labels = docker.buildx.Image("labels",
    context=docker.buildx.BuildContextArgs(
        location="app",
    ),
    labels={
        "description": "This image will get a descriptive label 👍",
    })
targets = docker.buildx.Image("targets",
    dockerfile=docker.buildx.DockerfileArgs(
        location="app/Dockerfile.targets",
    ),
    context=docker.buildx.BuildContextArgs(
        location="app",
    ),
    targets=[
        "build-me",
        "also-build-me",
    ])
named_contexts = docker.buildx.Image("namedContexts",
    dockerfile=docker.buildx.DockerfileArgs(
        location="app/Dockerfile.namedContexts",
    ),
    context=docker.buildx.BuildContextArgs(
        location="app",
        named={
            "golang:latest": docker.buildx.ContextArgs(
                location="docker-image://golang@sha256:b8e62cf593cdaff36efd90aa3a37de268e6781a2e68c6610940c48f7cdf36984",
            ),
        },
    ))
remote_context = docker.buildx.Image("remoteContext", context=docker.buildx.BuildContextArgs(
    location="https://raw.githubusercontent.com/pulumi/pulumi-docker/api-types/provider/testdata/Dockerfile",
))
remote_context_with_inline = docker.buildx.Image("remoteContextWithInline",
    dockerfile=docker.buildx.DockerfileArgs(
        inline="""FROM busybox
COPY hello.c ./
""",
    ),
    context=docker.buildx.BuildContextArgs(
        location="https://github.com/docker-library/hello-world.git",
    ))
inline = docker.buildx.Image("inline",
    dockerfile=docker.buildx.DockerfileArgs(
        inline="""FROM alpine
RUN echo "This uses an inline Dockerfile! 👍"
""",
    ),
    context=docker.buildx.BuildContextArgs(
        location="app",
    ))
docker_load = docker.buildx.Image("dockerLoad",
    context=docker.buildx.BuildContextArgs(
        location="app",
    ),
    exports=[docker.buildx.ExportEntryArgs(
        docker=docker.buildx.ExportDockerArgs(
            tar=True,
        ),
    )])
pulumi.export("platforms", multi_platform.platforms)
