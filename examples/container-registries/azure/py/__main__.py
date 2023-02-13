import pulumi
import pulumi_azure as azure
import pulumi_azuread as azuread
from pulumi_docker import Image, RegistryArgs, DockerBuildArgs
import pulumi_random as random

# Conditionalize the auth mechanism.
config = pulumi.Config()
use_sp_auth = config.get_bool('useServicePrincipalAuth') or False

# Create a private ACR registry.
rg = azure.core.ResourceGroup('myrg')
registry = azure.containerservice.Registry('myregistry',
    resource_group_name=rg.name,
    admin_enabled=True,
    sku='Basic'
)

# # Get registry info (creds and endpoint).
# image_name = registry.login_server.apply(lambda s: f'{s}/myapp')
# if use_sp_auth:
#     sp = azuread.ServicePrincipal('mysp',
#         application_id=azuread.Application('myspapp').application_id,
#     )
#     sp_password = azuread.ServicePrincipalPassword('mysp-pass',
#         service_principal_id=sp.id,
#         value=random.RandomPassword('mypass',
#             length=32,
#             opts=pulumi.ResourceOptions(additional_secret_outputs=['result'])
#         ).result,
#         end_date_relative='8760h',
#     )
#     sp_auth = azure.authorization.Assignment('myauth',
#         scope=registry.id,
#         role_definition_name='acrpush',
#         principal_id=sp.id,
#     )
#     registry_info = docker.ImageRegistry(
#         server=registry.login_server,
#         username=sp.application_id,
#         password=sp_auth.id.apply(lambda _: sp_password.value),
#         )
# else:
#     registry_info = docker.ImageRegistry(
#         server=registry.login_server,
#         username=registry.admin_username,
#         password=registry.admin_password
#     )

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

# Export the resulting base name in addition to the specific version pushed.
pulumi.export('baseImageName', image.base_image_name)
pulumi.export('imageName', image.image_name)
