import pulumi

from pulumi_docker import Image, DockerBuild

image = Image(
    "my-image",
    image_name="pulumi-user/example:v1.0.0",
    build=DockerBuild(
        target="dependencies",
    ),
    skip_push=True,
)

pulumi.export('deps-image', image.image_name)
