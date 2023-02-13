"""An Azure Python Pulumi program"""

import pulumi
from pulumi_azure import core, containerservice
from pulumi_docker import Image, RegistryArgs, DockerBuildArgs

resource_group = core.ResourceGroup('resource_group')

registry = containerservice.Registry(
    "myregistry", admin_enabled="true", resource_group_name=resource_group.name, sku="Basic")

my_image = Image("myimage",
                 image_name=registry.login_server.apply(
                     lambda server: f'{server}/app:v1.0.0'),
                 build=DockerBuildArgs(context=f'./app'),
                 registry=RegistryArgs(
                     server=registry.login_server,
                     username=registry.admin_username,
                     password=registry.admin_password)
                 )
