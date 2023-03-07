import base64
import json
import pulumi
import pulumi_digitalocean as digitalocean
from pulumi_docker import Image, DockerBuildArgs, RegistryArgs

# Create a private DigitalOcean Container Registry.
registry = digitalocean.ContainerRegistry('my-reg',
    subscription_tier_slug='starter',
)

# Get registry info (creds and endpoint).
image_name = registry.endpoint.apply(lambda s: f'{s}/myapp')
registry_creds = digitalocean.ContainerRegistryDockerCredentials('my-reg-creds',
    registry_name=registry.name,
    write=True,
)
def getRegistryInfo(info):
    # We are given a Docker creds file; parse it to find the temp username/password.
    auth_json = info[0]
    auths = json.loads(auth_json)
    server_url = info[1]
    auth_token = auths['auths'][server_url]['auth']
    decoded = base64.b64decode(auth_token).decode()
    parts = decoded.split(':')
    if len(parts) != 2:
        raise Exception('Invalid credentials')
    return RegistryArgs(
        server=server_url,
        username=parts[0],
        password=parts[1],
    )
registry_info = pulumi.Output.all(
    registry_creds.docker_credentials, registry.server_url).apply(getRegistryInfo)

# Build and publish the image.
image = Image(
    'my-image',
    build=DockerBuildArgs(
        context='app',
    ),
    image_name=image_name,
    registry=registry_info,
)

# Export the resulting base name in addition to the specific version pushed.
pulumi.export('baseImageName', image.base_image_name)
pulumi.export('fullImageName', image.image_name)
