# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from ._inputs import *

__all__ = ['ImageArgs', 'Image']

@pulumi.input_type
class ImageArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 registry: pulumi.Input['RegistryArgs'],
                 registry_url: pulumi.Input[str],
                 context: Optional[pulumi.Input[str]] = None,
                 dockerfile: Optional[pulumi.Input[str]] = None,
                 tag: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Image resource.
        :param pulumi.Input[str] name: The image name
        :param pulumi.Input['RegistryArgs'] registry: The registry to push the image to
        :param pulumi.Input[str] registry_url: The URL of the registry server hosting the image.
        :param pulumi.Input[str] context: The path to the build context to use.
        :param pulumi.Input[str] dockerfile: The path to the Dockerfile to use.
        :param pulumi.Input[str] tag: The image tag.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "registry", registry)
        pulumi.set(__self__, "registry_url", registry_url)
        if context is not None:
            pulumi.set(__self__, "context", context)
        if dockerfile is not None:
            pulumi.set(__self__, "dockerfile", dockerfile)
        if tag is None:
            tag = 'latest'
        if tag is not None:
            pulumi.set(__self__, "tag", tag)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        The image name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def registry(self) -> pulumi.Input['RegistryArgs']:
        """
        The registry to push the image to
        """
        return pulumi.get(self, "registry")

    @registry.setter
    def registry(self, value: pulumi.Input['RegistryArgs']):
        pulumi.set(self, "registry", value)

    @property
    @pulumi.getter(name="registryURL")
    def registry_url(self) -> pulumi.Input[str]:
        """
        The URL of the registry server hosting the image.
        """
        return pulumi.get(self, "registry_url")

    @registry_url.setter
    def registry_url(self, value: pulumi.Input[str]):
        pulumi.set(self, "registry_url", value)

    @property
    @pulumi.getter
    def context(self) -> Optional[pulumi.Input[str]]:
        """
        The path to the build context to use.
        """
        return pulumi.get(self, "context")

    @context.setter
    def context(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "context", value)

    @property
    @pulumi.getter
    def dockerfile(self) -> Optional[pulumi.Input[str]]:
        """
        The path to the Dockerfile to use.
        """
        return pulumi.get(self, "dockerfile")

    @dockerfile.setter
    def dockerfile(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dockerfile", value)

    @property
    @pulumi.getter
    def tag(self) -> Optional[pulumi.Input[str]]:
        """
        The image tag.
        """
        return pulumi.get(self, "tag")

    @tag.setter
    def tag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tag", value)


class Image(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 context: Optional[pulumi.Input[str]] = None,
                 dockerfile: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 registry: Optional[pulumi.Input[pulumi.InputType['RegistryArgs']]] = None,
                 registry_url: Optional[pulumi.Input[str]] = None,
                 tag: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        A real CRUD docker image we hope

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] context: The path to the build context to use.
        :param pulumi.Input[str] dockerfile: The path to the Dockerfile to use.
        :param pulumi.Input[str] name: The image name
        :param pulumi.Input[pulumi.InputType['RegistryArgs']] registry: The registry to push the image to
        :param pulumi.Input[str] registry_url: The URL of the registry server hosting the image.
        :param pulumi.Input[str] tag: The image tag.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ImageArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A real CRUD docker image we hope

        :param str resource_name: The name of the resource.
        :param ImageArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ImageArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 context: Optional[pulumi.Input[str]] = None,
                 dockerfile: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 registry: Optional[pulumi.Input[pulumi.InputType['RegistryArgs']]] = None,
                 registry_url: Optional[pulumi.Input[str]] = None,
                 tag: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ImageArgs.__new__(ImageArgs)

            __props__.__dict__["context"] = context
            __props__.__dict__["dockerfile"] = dockerfile
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__.__dict__["name"] = name
            if registry is None and not opts.urn:
                raise TypeError("Missing required property 'registry'")
            __props__.__dict__["registry"] = registry
            if registry_url is None and not opts.urn:
                raise TypeError("Missing required property 'registry_url'")
            __props__.__dict__["registry_url"] = registry_url
            if tag is None:
                tag = 'latest'
            __props__.__dict__["tag"] = tag
        super(Image, __self__).__init__(
            'docker:index/image:Image',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Image':
        """
        Get an existing Image resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ImageArgs.__new__(ImageArgs)

        __props__.__dict__["context"] = None
        __props__.__dict__["dockerfile"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["registry_url"] = None
        __props__.__dict__["tag"] = None
        return Image(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def context(self) -> pulumi.Output[str]:
        """
        The path to the build context to use.
        """
        return pulumi.get(self, "context")

    @property
    @pulumi.getter
    def dockerfile(self) -> pulumi.Output[str]:
        """
        The path to the Dockerfile to use.
        """
        return pulumi.get(self, "dockerfile")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The image name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="registryURL")
    def registry_url(self) -> pulumi.Output[str]:
        """
        The URL of the registry server hosting the image.
        """
        return pulumi.get(self, "registry_url")

    @property
    @pulumi.getter
    def tag(self) -> pulumi.Output[Optional[str]]:
        """
        The image tag.
        """
        return pulumi.get(self, "tag")

