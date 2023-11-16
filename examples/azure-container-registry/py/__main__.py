import pulumi
import pulumi_azure as azure
from pulumi_docker import Image, RegistryArgs, DockerBuildArgs

# Create a private ACR registry.
rg = azure.core.ResourceGroup('myrg')
registry = azure.containerservice.Registry('myregistry',
    resource_group_name=rg.name,
    admin_enabled=True,
    sku='Basic'
)

# Build and publish the image.
image = Image('my-image',
    build=DockerBuildArgs(context=f'./app'),
    image_name=registry.login_server.apply(lambda s: f'{s}/myapp'),
    registry=RegistryArgs(
        server=registry.login_server,
        username=registry.admin_username,
        password=registry.admin_password
    )
)

# Export the resulting image name
pulumi.export('imageName', image.image_name)
pulumi.export('repoDigest', image.repo_digest)
