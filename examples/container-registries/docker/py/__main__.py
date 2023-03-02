import pulumi
import pulumi_docker as docker

# Fetch the Docker Hub auth info from config.
config = pulumi.Config()
username = config.require('dockerUsername')
accessToken = config.require_secret('dockerPassword')

# Populate the registry info (creds and endpoint).
image_name=f'{username}/myapp',


def get_registry_info(token):
    return docker.ImageRegistry(
        server='docker.io',
        username=username,
        password=token,
    )


registry_info=accessToken.apply(get_registry_info)

# Build and publish the image.
image = docker.Image('my-image',
    build='app',
    image_name=image_name,
    registry=registry_info,
)

# Export the resulting base name in addition to the specific version pushed.
pulumi.export('baseImageName', image.base_image_name)
pulumi.export('fullImageName', image.image_name)
