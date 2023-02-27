import base64
import pulumi
import pulumi_aws as aws
from pulumi_docker import Image, DockerBuildArgs, RegistryArgs

# Create a private ECR registry.
repo = aws.ecr.Repository(
    'my-repo',
    force_delete=True,

)

# Get registry info (creds and endpoint) so we can build/publish to it.
def getRegistryInfo(rid):
    creds = aws.ecr.get_credentials(registry_id=rid)
    decoded = base64.b64decode(creds.authorization_token).decode()
    parts = decoded.split(':')
    if len(parts) != 2:
        raise Exception("Invalid credentials")
    return RegistryArgs(
        server=creds.proxy_endpoint,
        username=parts[0],
        password=parts[1],
    )


registryInfo = repo.registry_id.apply(getRegistryInfo)

# Build and publish the image.
image = Image(
    'my-image',
    build=DockerBuildArgs(
        context='app',
    ),
    image_name=repo.repository_url,
    registry=registryInfo,
)

# Export the resulting base name in addition to the specific version pushed.
pulumi.export('baseImageName', image.base_image_name)
pulumi.export('imageName', image.image_name)
