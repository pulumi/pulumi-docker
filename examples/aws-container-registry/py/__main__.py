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
    creds = aws.ecr.get_authorization_token(registry_id=rid)

    return RegistryArgs(
        server=creds.proxy_endpoint,
        username=creds.user_name,
        password=creds.password,
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
pulumi.export('repoDigest', image.repo_digest)
