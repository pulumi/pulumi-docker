import pulumi
from pulumi_docker import Image, DockerBuildArgs

image = Image(
    "my-image",
    image_name="pulumi-user/example:v1.0.0",
    build=DockerBuildArgs(
        target="dependencies",
        env={'TEST_ENV': '42'},
    ),
    skip_push=True,
)

pulumi.export('deps-image', image.image_name)
