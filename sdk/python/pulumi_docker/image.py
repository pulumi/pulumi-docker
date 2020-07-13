# Copyright 2016-2020, Pulumi Corporation.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

import pulumi
from typing import Optional, Union

from .docker import build_and_push_image, DockerBuild, Registry
from .utils import get_image_name_and_tag


class ImageRegistry:
    server: pulumi.Input[str]
    """
    Docker registry server URL to push to. 
    
    Some common values include:
        - DockerHub: `docker.io` or `https://index.docker.io/v1`
        - Azure Container Registry: `<name>.azurecr.io`
        - AWS Elastic Container Registry: `<account>.dkr.ecr.us-east-2.amazonaws.com`
        - Google Container Registry: `<name>.gcr.io`
    """

    username: pulumi.Input[str]
    """
    Username for login to the target Docker registry.
    """

    password: pulumi.Input[str]
    """
    Password for login to the target Docker registry.
    """

    def __init__(self, server: pulumi.Input[str], username: pulumi.Input[str], password: pulumi.Input[str]):
        """
        Image is a resource that represents a Docker image built locally which is published and made
        available via a remote Docker registry.  This can be used to ensure that a Docker source
        directory from a local deployment environment is built and pushed to a cloud-hosted Docker
        registry as part of a Pulumi deployment, so that it can be referenced as an image input from
        other cloud services that reference Docker images - including Kubernetes Pods, AWS ECS Tasks, and
        Azure Container Instances.


        :param pulumi.Input[str] server: Docker registry server URL to push to.
        :param pulumi.Input[str] username: Username for login to the target Docker registry.
        :param pulumi.Input[str] password: Password for login to the target Docker registry.
        """
        self.server = server
        self.username = username
        self.password = password


class _ImageArgs:
    image_name: str
    """
    The qualified image name that will be pushed to the remote registry.  Must be a supported
    image name for the target registry user.  This name can include a tag at the end.  If
    provided all pushed image resources will contain that tag as well.

    Either [image_name] or [localImageName] can have a tag.  However, if both have a tag, then
    those tags must match.
    """

    build: Union[str, DockerBuild]
    """
    The Docker build context, as a folder path or a detailed DockerBuild object.
    """

    local_image_name: str
    """
    The docker image name to build locally before tagging with image_name.  If not provided, it
    will be given the value of to [image_name].  This name can include a tag at the end.  If
    provided all pushed image resources will contain that tag as well.

    Either [image_name] or [localImageName] can have a tag.  However, if both have a tag, then
    those tags must match.
    """

    registry: Optional[ImageRegistry]
    """
    Credentials for the docker registry to push to.
    """

    skip_push: Optional[bool]
    """
    Skip push flag.
    """

    def __init__(self,
                 image_name: str,
                 build: Union[str, DockerBuild],
                 local_image_name: str,
                 registry_server: Optional[str],
                 registry_username: Optional[str],
                 registry_password: Optional[str],
                 skip_push: Optional[bool]):
        """
        Arguments for constructing an Image resource.


        :param str image_name: The qualified image name that will be pushed to the remote registry.  Must be a supported
            image name for the target registry user.  This name can include a tag at the end.  If
            provided all pushed image resources will contain that tag as well.
            Either [image_name] or [localImageName] can have a tag.  However, if both have a tag, then
            those tags must match.
        :param Union[str, DockerBuild] build: The Docker build context, as a folder path or a detailed DockerBuild
            object.
        :param str local_image_name: The docker image name to build locally before tagging with image_name.  If not
            provided, it will be given the value of to [image_name].  This name can include a tag at the end.  If
            provided all pushed image resources will contain that tag as well.
            Either [image_name] or [localImageName] can have a tag.  However, if both have a tag, then
            those tags must match.
        :param Optional[str] registry_server: Docker registry to push to.
        :param Optional[str] registry_username:  Username for the registry.
        :param Optional[str] registry_password: Password for the registry.
        :param Optional[bool]) skip_push: Skip push flag.
        """
        self.image_name = image_name
        self.build = build
        self.local_image_name = local_image_name
        self.registry_server = registry_server
        self.registry_username = registry_username
        self.registry_password = registry_password
        self.skip_push = skip_push


class Image(pulumi.ComponentResource):
    base_image_name: pulumi.Output[str]
    """
    The base image name that was built and pushed.  This does not include the id annotation, so
    is not pinned to the specific build performed by this docker.Image.
    """

    image_name: pulumi.Output[str]
    """
    The unique pinned image name on the remote repository."""

    registry_server: pulumi.Output[Optional[str]]
    """
    The server the image is located at.
    """

    def __init__(self, name: str,
                 image_name: pulumi.Input[str],
                 build: Union[pulumi.Input[str], DockerBuild],
                 local_image_name: Optional[pulumi.Input[str]] = None,
                 registry: Optional[pulumi.Input[ImageRegistry]] = None,
                 skip_push: Optional[bool] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A docker.Image resource represents a Docker image built locally which is published and made
        available via a remote Docker registry.  This can be used to ensure that a Docker source
        directory from a local deployment environment is built and pushed to a cloud-hosted Docker
        registry as part of a Pulumi deployment, so that it can be referenced as an image input from
        other cloud services that reference Docker images - including Kubernetes Pods, AWS ECS Tasks, and
        Azure Container Instances.


        :param str name: The name of the resource.
        :param pulumi.Input[str] image_name: The qualified image name that will be pushed to the remote registry.
            Must be a supported image name for the target registry user.  This name can include a tag at the end.  If
            provided all pushed image resources will contain that tag as well.
            Either [imageName] or [localImageName] can have a tag.  However, if both have a tag, then
            those tags must match.
        :param Union[pulumi.Input[str], DockerBuild] build: The Docker build context, as a folder path or a detailed
            DockerBuild object.
        :param Optional[pulumi.Input[str]] local_image_name: The docker image name to build locally before tagging with
            imageName.  If not provided, it will be given the value of to [imageName].  This name can include a tag at
            the end.  If provided all pushed image resources will contain that tag as well.
            Either [imageName] or [localImageName] can have a tag.  However, if both have a tag, then
            those tags must match.
        :param Optional[pulumi.Input[ImageRegistry]] registry: Credentials for the docker registry to push to.
        :param Optional[bool] skip_push: Skip push flag.
        :param Optional[pulumi.ResourceOptions] opts: Options for the resource.
        """
        super().__init__("docker:image:Image", name, {}, opts)

        def get_image_data(image_args: _ImageArgs):
            image_name = image_args.image_name

            # If there is no local_image_name set it equal to image_name.  Note: this means
            # that if image_name contains a tag, local_image_name will contain the same tag.
            local_image_name = image_args.local_image_name or image_name

            # Now break both the localImageName and the image_name into the untagged part and the
            # optional tag.  If both have tags, they must match.  If one or the other has a tag, we
            # just use that as the tag to use.  This allows users to flexibly provide a tag on one
            # option or the other and still have it work out.
            local_image_name_without_tag, local_image_name_tag = get_image_name_and_tag(local_image_name)
            image_name_without_tag, image_name_tag = get_image_name_and_tag(image_name)

            tag = local_image_name_tag or image_name_tag

            def check_tag(t: Optional[str]):
                if t and (t != tag):
                    raise Exception(f'[local_image_name] and [image_name]'
                                    ' had mismatched tags. {local_image_name_tag} != {image_name_tag}')

            check_tag(local_image_name_tag)
            check_tag(image_name_tag)

            # build_and_push_image expects only the base_image_name to have a tag.  So build that
            # name appropriately if we were given a tag.
            base_image_name = f'{local_image_name_without_tag}:{tag}' if tag else local_image_name

            # build_and_push_image does not want the repository_url to have a tag.  This is just
            # the base url where the images will be pushed to.  All tagging will be taken care of
            # inside that api.
            repository_url = image_name_without_tag

            def get_registry():
                registry_server = image_args.registry_server
                registry_password = image_args.registry_password
                registry_username = image_args.registry_username

                if registry_server and registry_username and registry_password:
                    return Registry(registry_server, registry_username, registry_password)
                else:
                    return None

            unique_target_name = build_and_push_image(
                base_image_name,
                image_args.build,
                repository_url,
                self,
                get_registry(),
                image_args.skip_push,
            )

            return {
                'image_name': unique_target_name,
                'registry_server': image_args.registry_server or None
            }

        image_data = pulumi.Output.all(
            image_name,
            build,
            local_image_name,
            registry.server if registry else None,
            registry.username if registry else None,
            registry.password if registry else None,
            skip_push,
        ).apply(
            lambda args: get_image_data(_ImageArgs(*args))
        )

        self.image_name, self.registry_server = image_data['image_name'], image_data['registry_server']

        self.base_image_name = pulumi.Output.from_input(image_name)

        self.register_outputs({
            'base_image_name': self.image_name,
            'image_name': self.image_name,
            'registry_server': self.registry_server,
        })
