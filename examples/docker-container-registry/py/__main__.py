import pulumi
from pulumi_docker import Image, DockerBuildArgs, RegistryArgs

# Fetch the Docker Hub auth info from config.
config = pulumi.Config()
username = config.require('dockerUsername')
access_token = config.require_secret('dockerPassword')

# Populate the registry info (creds and endpoint).
image_name=f'docker.io/{username}/myapp'


def get_registry_info(token):
    return RegistryArgs(
        server='docker.io',
        username=username,
        password=token,
    )


registry_info=access_token.apply(get_registry_info)

# Build and publish the image.
image = Image(
    'my-image',
    build=DockerBuildArgs(
        context='app',
    ),
    image_name=image_name,
    registry=registry_info,
)

# Export the resulting image name
pulumi.export('fullImageName', image.image_name)
pulumi.export('repoDigest', image.repo_digest)
