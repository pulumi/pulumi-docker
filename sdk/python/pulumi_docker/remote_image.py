# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['RemoteImageArgs', 'RemoteImage']

@pulumi.input_type
class RemoteImageArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[builtins.str],
                 build: Optional[pulumi.Input['RemoteImageBuildArgs']] = None,
                 force_remove: Optional[pulumi.Input[builtins.bool]] = None,
                 keep_locally: Optional[pulumi.Input[builtins.bool]] = None,
                 platform: Optional[pulumi.Input[builtins.str]] = None,
                 pull_triggers: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 triggers: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None):
        """
        The set of arguments for constructing a RemoteImage resource.
        :param pulumi.Input[builtins.str] name: The name of the Docker image, including any tags or SHA256 repo digests.
        :param pulumi.Input['RemoteImageBuildArgs'] build: Configuration to build an image. Please see [docker build command reference](https://docs.docker.com/engine/reference/commandline/build/#options) too.
        :param pulumi.Input[builtins.bool] force_remove: If true, then the image is removed forcibly when the resource is destroyed.
        :param pulumi.Input[builtins.bool] keep_locally: If true, then the Docker image won't be deleted on destroy operation. If this is false, it will delete the image from the docker local storage on destroy operation.
        :param pulumi.Input[builtins.str] platform: The platform to use when pulling the image. Defaults to the platform of the current machine.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] pull_triggers: List of values which cause an image pull when changed. This is used to store the image digest from the registry when using the docker*registry*image.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] triggers: A map of arbitrary strings that, when changed, will force the `RemoteImage` resource to be replaced. This can be used to rebuild an image when contents of source code folders change
        """
        pulumi.set(__self__, "name", name)
        if build is not None:
            pulumi.set(__self__, "build", build)
        if force_remove is not None:
            pulumi.set(__self__, "force_remove", force_remove)
        if keep_locally is not None:
            pulumi.set(__self__, "keep_locally", keep_locally)
        if platform is not None:
            pulumi.set(__self__, "platform", platform)
        if pull_triggers is not None:
            pulumi.set(__self__, "pull_triggers", pull_triggers)
        if triggers is not None:
            pulumi.set(__self__, "triggers", triggers)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the Docker image, including any tags or SHA256 repo digests.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def build(self) -> Optional[pulumi.Input['RemoteImageBuildArgs']]:
        """
        Configuration to build an image. Please see [docker build command reference](https://docs.docker.com/engine/reference/commandline/build/#options) too.
        """
        return pulumi.get(self, "build")

    @build.setter
    def build(self, value: Optional[pulumi.Input['RemoteImageBuildArgs']]):
        pulumi.set(self, "build", value)

    @property
    @pulumi.getter(name="forceRemove")
    def force_remove(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        If true, then the image is removed forcibly when the resource is destroyed.
        """
        return pulumi.get(self, "force_remove")

    @force_remove.setter
    def force_remove(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "force_remove", value)

    @property
    @pulumi.getter(name="keepLocally")
    def keep_locally(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        If true, then the Docker image won't be deleted on destroy operation. If this is false, it will delete the image from the docker local storage on destroy operation.
        """
        return pulumi.get(self, "keep_locally")

    @keep_locally.setter
    def keep_locally(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "keep_locally", value)

    @property
    @pulumi.getter
    def platform(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The platform to use when pulling the image. Defaults to the platform of the current machine.
        """
        return pulumi.get(self, "platform")

    @platform.setter
    def platform(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "platform", value)

    @property
    @pulumi.getter(name="pullTriggers")
    def pull_triggers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        List of values which cause an image pull when changed. This is used to store the image digest from the registry when using the docker*registry*image.
        """
        return pulumi.get(self, "pull_triggers")

    @pull_triggers.setter
    def pull_triggers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "pull_triggers", value)

    @property
    @pulumi.getter
    def triggers(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        """
        A map of arbitrary strings that, when changed, will force the `RemoteImage` resource to be replaced. This can be used to rebuild an image when contents of source code folders change
        """
        return pulumi.get(self, "triggers")

    @triggers.setter
    def triggers(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "triggers", value)


@pulumi.input_type
class _RemoteImageState:
    def __init__(__self__, *,
                 build: Optional[pulumi.Input['RemoteImageBuildArgs']] = None,
                 force_remove: Optional[pulumi.Input[builtins.bool]] = None,
                 image_id: Optional[pulumi.Input[builtins.str]] = None,
                 keep_locally: Optional[pulumi.Input[builtins.bool]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 platform: Optional[pulumi.Input[builtins.str]] = None,
                 pull_triggers: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 repo_digest: Optional[pulumi.Input[builtins.str]] = None,
                 triggers: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None):
        """
        Input properties used for looking up and filtering RemoteImage resources.
        :param pulumi.Input['RemoteImageBuildArgs'] build: Configuration to build an image. Please see [docker build command reference](https://docs.docker.com/engine/reference/commandline/build/#options) too.
        :param pulumi.Input[builtins.bool] force_remove: If true, then the image is removed forcibly when the resource is destroyed.
        :param pulumi.Input[builtins.str] image_id: The ID of the image (as seen when executing `docker inspect` on the image). Can be used to reference the image via its ID in other resources.
        :param pulumi.Input[builtins.bool] keep_locally: If true, then the Docker image won't be deleted on destroy operation. If this is false, it will delete the image from the docker local storage on destroy operation.
        :param pulumi.Input[builtins.str] name: The name of the Docker image, including any tags or SHA256 repo digests.
        :param pulumi.Input[builtins.str] platform: The platform to use when pulling the image. Defaults to the platform of the current machine.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] pull_triggers: List of values which cause an image pull when changed. This is used to store the image digest from the registry when using the docker*registry*image.
        :param pulumi.Input[builtins.str] repo_digest: The image sha256 digest in the form of `repo[:tag]@sha256:<hash>`.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] triggers: A map of arbitrary strings that, when changed, will force the `RemoteImage` resource to be replaced. This can be used to rebuild an image when contents of source code folders change
        """
        if build is not None:
            pulumi.set(__self__, "build", build)
        if force_remove is not None:
            pulumi.set(__self__, "force_remove", force_remove)
        if image_id is not None:
            pulumi.set(__self__, "image_id", image_id)
        if keep_locally is not None:
            pulumi.set(__self__, "keep_locally", keep_locally)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if platform is not None:
            pulumi.set(__self__, "platform", platform)
        if pull_triggers is not None:
            pulumi.set(__self__, "pull_triggers", pull_triggers)
        if repo_digest is not None:
            pulumi.set(__self__, "repo_digest", repo_digest)
        if triggers is not None:
            pulumi.set(__self__, "triggers", triggers)

    @property
    @pulumi.getter
    def build(self) -> Optional[pulumi.Input['RemoteImageBuildArgs']]:
        """
        Configuration to build an image. Please see [docker build command reference](https://docs.docker.com/engine/reference/commandline/build/#options) too.
        """
        return pulumi.get(self, "build")

    @build.setter
    def build(self, value: Optional[pulumi.Input['RemoteImageBuildArgs']]):
        pulumi.set(self, "build", value)

    @property
    @pulumi.getter(name="forceRemove")
    def force_remove(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        If true, then the image is removed forcibly when the resource is destroyed.
        """
        return pulumi.get(self, "force_remove")

    @force_remove.setter
    def force_remove(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "force_remove", value)

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The ID of the image (as seen when executing `docker inspect` on the image). Can be used to reference the image via its ID in other resources.
        """
        return pulumi.get(self, "image_id")

    @image_id.setter
    def image_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "image_id", value)

    @property
    @pulumi.getter(name="keepLocally")
    def keep_locally(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        If true, then the Docker image won't be deleted on destroy operation. If this is false, it will delete the image from the docker local storage on destroy operation.
        """
        return pulumi.get(self, "keep_locally")

    @keep_locally.setter
    def keep_locally(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "keep_locally", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the Docker image, including any tags or SHA256 repo digests.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def platform(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The platform to use when pulling the image. Defaults to the platform of the current machine.
        """
        return pulumi.get(self, "platform")

    @platform.setter
    def platform(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "platform", value)

    @property
    @pulumi.getter(name="pullTriggers")
    def pull_triggers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        List of values which cause an image pull when changed. This is used to store the image digest from the registry when using the docker*registry*image.
        """
        return pulumi.get(self, "pull_triggers")

    @pull_triggers.setter
    def pull_triggers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "pull_triggers", value)

    @property
    @pulumi.getter(name="repoDigest")
    def repo_digest(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The image sha256 digest in the form of `repo[:tag]@sha256:<hash>`.
        """
        return pulumi.get(self, "repo_digest")

    @repo_digest.setter
    def repo_digest(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "repo_digest", value)

    @property
    @pulumi.getter
    def triggers(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        """
        A map of arbitrary strings that, when changed, will force the `RemoteImage` resource to be replaced. This can be used to rebuild an image when contents of source code folders change
        """
        return pulumi.get(self, "triggers")

    @triggers.setter
    def triggers(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "triggers", value)


class RemoteImage(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 build: Optional[pulumi.Input[Union['RemoteImageBuildArgs', 'RemoteImageBuildArgsDict']]] = None,
                 force_remove: Optional[pulumi.Input[builtins.bool]] = None,
                 keep_locally: Optional[pulumi.Input[builtins.bool]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 platform: Optional[pulumi.Input[builtins.str]] = None,
                 pull_triggers: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 triggers: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        """
        <!-- Bug: Type and Name are switched -->
        Pulls a Docker image to a given Docker host from a Docker Registry.
         This resource will *not* pull new layers of the image automatically unless used in conjunction with RegistryImage data source to update the `pull_triggers` field.

        ## Example Usage

        ### Basic

        Finds and downloads the latest `ubuntu:precise` image but does not check
        for further updates of the image

        ```python
        import pulumi
        import pulumi_docker as docker

        ubuntu = docker.RemoteImage("ubuntu", name="ubuntu:precise")
        ```

        ### Dynamic updates

        To be able to update an image dynamically when the `sha256` sum changes,
        you need to use it in combination with `RegistryImage` as follows:

        ```python
        import pulumi
        import pulumi_docker as docker

        ubuntu = docker.get_registry_image(name="ubuntu:precise")
        ubuntu_remote_image = docker.RemoteImage("ubuntu",
            name=ubuntu.name,
            pull_triggers=[ubuntu.sha256_digest])
        ```

        ### Build

        You can also use the resource to build an image.
        In this case the image "zoo" and "zoo:develop" are built.

        ```python
        import pulumi
        import pulumi_docker as docker

        zoo = docker.RemoteImage("zoo",
            name="zoo",
            build={
                "context": ".",
                "tags": ["zoo:develop"],
                "build_arg": {
                    "foo": "zoo",
                },
                "label": {
                    "author": "zoo",
                },
            })
        ```

        You can use the `triggers` argument to specify when the image should be rebuild. This is for example helpful when you want to rebuild the docker image whenever the source code changes.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['RemoteImageBuildArgs', 'RemoteImageBuildArgsDict']] build: Configuration to build an image. Please see [docker build command reference](https://docs.docker.com/engine/reference/commandline/build/#options) too.
        :param pulumi.Input[builtins.bool] force_remove: If true, then the image is removed forcibly when the resource is destroyed.
        :param pulumi.Input[builtins.bool] keep_locally: If true, then the Docker image won't be deleted on destroy operation. If this is false, it will delete the image from the docker local storage on destroy operation.
        :param pulumi.Input[builtins.str] name: The name of the Docker image, including any tags or SHA256 repo digests.
        :param pulumi.Input[builtins.str] platform: The platform to use when pulling the image. Defaults to the platform of the current machine.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] pull_triggers: List of values which cause an image pull when changed. This is used to store the image digest from the registry when using the docker*registry*image.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] triggers: A map of arbitrary strings that, when changed, will force the `RemoteImage` resource to be replaced. This can be used to rebuild an image when contents of source code folders change
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RemoteImageArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        <!-- Bug: Type and Name are switched -->
        Pulls a Docker image to a given Docker host from a Docker Registry.
         This resource will *not* pull new layers of the image automatically unless used in conjunction with RegistryImage data source to update the `pull_triggers` field.

        ## Example Usage

        ### Basic

        Finds and downloads the latest `ubuntu:precise` image but does not check
        for further updates of the image

        ```python
        import pulumi
        import pulumi_docker as docker

        ubuntu = docker.RemoteImage("ubuntu", name="ubuntu:precise")
        ```

        ### Dynamic updates

        To be able to update an image dynamically when the `sha256` sum changes,
        you need to use it in combination with `RegistryImage` as follows:

        ```python
        import pulumi
        import pulumi_docker as docker

        ubuntu = docker.get_registry_image(name="ubuntu:precise")
        ubuntu_remote_image = docker.RemoteImage("ubuntu",
            name=ubuntu.name,
            pull_triggers=[ubuntu.sha256_digest])
        ```

        ### Build

        You can also use the resource to build an image.
        In this case the image "zoo" and "zoo:develop" are built.

        ```python
        import pulumi
        import pulumi_docker as docker

        zoo = docker.RemoteImage("zoo",
            name="zoo",
            build={
                "context": ".",
                "tags": ["zoo:develop"],
                "build_arg": {
                    "foo": "zoo",
                },
                "label": {
                    "author": "zoo",
                },
            })
        ```

        You can use the `triggers` argument to specify when the image should be rebuild. This is for example helpful when you want to rebuild the docker image whenever the source code changes.

        :param str resource_name: The name of the resource.
        :param RemoteImageArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RemoteImageArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 build: Optional[pulumi.Input[Union['RemoteImageBuildArgs', 'RemoteImageBuildArgsDict']]] = None,
                 force_remove: Optional[pulumi.Input[builtins.bool]] = None,
                 keep_locally: Optional[pulumi.Input[builtins.bool]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 platform: Optional[pulumi.Input[builtins.str]] = None,
                 pull_triggers: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 triggers: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RemoteImageArgs.__new__(RemoteImageArgs)

            __props__.__dict__["build"] = build
            __props__.__dict__["force_remove"] = force_remove
            __props__.__dict__["keep_locally"] = keep_locally
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__.__dict__["name"] = name
            __props__.__dict__["platform"] = platform
            __props__.__dict__["pull_triggers"] = pull_triggers
            __props__.__dict__["triggers"] = triggers
            __props__.__dict__["image_id"] = None
            __props__.__dict__["repo_digest"] = None
        super(RemoteImage, __self__).__init__(
            'docker:index/remoteImage:RemoteImage',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            build: Optional[pulumi.Input[Union['RemoteImageBuildArgs', 'RemoteImageBuildArgsDict']]] = None,
            force_remove: Optional[pulumi.Input[builtins.bool]] = None,
            image_id: Optional[pulumi.Input[builtins.str]] = None,
            keep_locally: Optional[pulumi.Input[builtins.bool]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None,
            platform: Optional[pulumi.Input[builtins.str]] = None,
            pull_triggers: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
            repo_digest: Optional[pulumi.Input[builtins.str]] = None,
            triggers: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None) -> 'RemoteImage':
        """
        Get an existing RemoteImage resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['RemoteImageBuildArgs', 'RemoteImageBuildArgsDict']] build: Configuration to build an image. Please see [docker build command reference](https://docs.docker.com/engine/reference/commandline/build/#options) too.
        :param pulumi.Input[builtins.bool] force_remove: If true, then the image is removed forcibly when the resource is destroyed.
        :param pulumi.Input[builtins.str] image_id: The ID of the image (as seen when executing `docker inspect` on the image). Can be used to reference the image via its ID in other resources.
        :param pulumi.Input[builtins.bool] keep_locally: If true, then the Docker image won't be deleted on destroy operation. If this is false, it will delete the image from the docker local storage on destroy operation.
        :param pulumi.Input[builtins.str] name: The name of the Docker image, including any tags or SHA256 repo digests.
        :param pulumi.Input[builtins.str] platform: The platform to use when pulling the image. Defaults to the platform of the current machine.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] pull_triggers: List of values which cause an image pull when changed. This is used to store the image digest from the registry when using the docker*registry*image.
        :param pulumi.Input[builtins.str] repo_digest: The image sha256 digest in the form of `repo[:tag]@sha256:<hash>`.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] triggers: A map of arbitrary strings that, when changed, will force the `RemoteImage` resource to be replaced. This can be used to rebuild an image when contents of source code folders change
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RemoteImageState.__new__(_RemoteImageState)

        __props__.__dict__["build"] = build
        __props__.__dict__["force_remove"] = force_remove
        __props__.__dict__["image_id"] = image_id
        __props__.__dict__["keep_locally"] = keep_locally
        __props__.__dict__["name"] = name
        __props__.__dict__["platform"] = platform
        __props__.__dict__["pull_triggers"] = pull_triggers
        __props__.__dict__["repo_digest"] = repo_digest
        __props__.__dict__["triggers"] = triggers
        return RemoteImage(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def build(self) -> pulumi.Output[Optional['outputs.RemoteImageBuild']]:
        """
        Configuration to build an image. Please see [docker build command reference](https://docs.docker.com/engine/reference/commandline/build/#options) too.
        """
        return pulumi.get(self, "build")

    @property
    @pulumi.getter(name="forceRemove")
    def force_remove(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        If true, then the image is removed forcibly when the resource is destroyed.
        """
        return pulumi.get(self, "force_remove")

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> pulumi.Output[builtins.str]:
        """
        The ID of the image (as seen when executing `docker inspect` on the image). Can be used to reference the image via its ID in other resources.
        """
        return pulumi.get(self, "image_id")

    @property
    @pulumi.getter(name="keepLocally")
    def keep_locally(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        If true, then the Docker image won't be deleted on destroy operation. If this is false, it will delete the image from the docker local storage on destroy operation.
        """
        return pulumi.get(self, "keep_locally")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the Docker image, including any tags or SHA256 repo digests.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def platform(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The platform to use when pulling the image. Defaults to the platform of the current machine.
        """
        return pulumi.get(self, "platform")

    @property
    @pulumi.getter(name="pullTriggers")
    def pull_triggers(self) -> pulumi.Output[Optional[Sequence[builtins.str]]]:
        """
        List of values which cause an image pull when changed. This is used to store the image digest from the registry when using the docker*registry*image.
        """
        return pulumi.get(self, "pull_triggers")

    @property
    @pulumi.getter(name="repoDigest")
    def repo_digest(self) -> pulumi.Output[builtins.str]:
        """
        The image sha256 digest in the form of `repo[:tag]@sha256:<hash>`.
        """
        return pulumi.get(self, "repo_digest")

    @property
    @pulumi.getter
    def triggers(self) -> pulumi.Output[Optional[Mapping[str, builtins.str]]]:
        """
        A map of arbitrary strings that, when changed, will force the `RemoteImage` resource to be replaced. This can be used to rebuild an image when contents of source code folders change
        """
        return pulumi.get(self, "triggers")

