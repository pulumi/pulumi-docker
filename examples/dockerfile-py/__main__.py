import pulumi

import pulumi_docker as pd

image = pd.Image(
    "my-image",
    image_name="pulumi-user/example:v1.0.0",
    build=pd.DockerBuild(
        target="dependencies",
    ),
    skip_push=True,
)

pulumi.export('deps-image', image.image_name)
