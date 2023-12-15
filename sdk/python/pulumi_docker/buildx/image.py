# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ImageArgs', 'Image']

@pulumi.input_type
class ImageArgs:
    def __init__(__self__, *,
                 tags: pulumi.Input[Sequence[pulumi.Input[str]]],
                 context: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 exports: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 file: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Image resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: 
               Name and optionally a tag (format: "name:tag"). If outputting to a
               registry, the name should include the fully qualified registry address.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] context: 
               Contexts to use while building the image. If omitted, an empty context
               is used. If more than one value is specified, they should be of the
               form "name=value".
        :param pulumi.Input[Sequence[pulumi.Input[str]]] exports: 
               Name and optionally a tag (format: "name:tag"). If outputting to a
               registry, the name should include the fully qualified registry address.
        :param pulumi.Input[str] file: 
               Name of the Dockerfile to use (default: "$PATH/Dockerfile").
        """
        pulumi.set(__self__, "tags", tags)
        if context is not None:
            pulumi.set(__self__, "context", context)
        if exports is not None:
            pulumi.set(__self__, "exports", exports)
        if file is None:
            file = 'Dockerfile'
        if file is not None:
            pulumi.set(__self__, "file", file)

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """

        Name and optionally a tag (format: "name:tag"). If outputting to a
        registry, the name should include the fully qualified registry address.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def context(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """

        Contexts to use while building the image. If omitted, an empty context
        is used. If more than one value is specified, they should be of the
        form "name=value".
        """
        return pulumi.get(self, "context")

    @context.setter
    def context(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "context", value)

    @property
    @pulumi.getter
    def exports(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """

        Name and optionally a tag (format: "name:tag"). If outputting to a
        registry, the name should include the fully qualified registry address.
        """
        return pulumi.get(self, "exports")

    @exports.setter
    def exports(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "exports", value)

    @property
    @pulumi.getter
    def file(self) -> Optional[pulumi.Input[str]]:
        """

        Name of the Dockerfile to use (default: "$PATH/Dockerfile").
        """
        return pulumi.get(self, "file")

    @file.setter
    def file(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "file", value)


class Image(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 context: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 exports: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 file: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        A Docker image built using Buildkit

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] context: 
               Contexts to use while building the image. If omitted, an empty context
               is used. If more than one value is specified, they should be of the
               form "name=value".
        :param pulumi.Input[Sequence[pulumi.Input[str]]] exports: 
               Name and optionally a tag (format: "name:tag"). If outputting to a
               registry, the name should include the fully qualified registry address.
        :param pulumi.Input[str] file: 
               Name of the Dockerfile to use (default: "$PATH/Dockerfile").
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: 
               Name and optionally a tag (format: "name:tag"). If outputting to a
               registry, the name should include the fully qualified registry address.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ImageArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A Docker image built using Buildkit

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
                 context: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 exports: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 file: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ImageArgs.__new__(ImageArgs)

            __props__.__dict__["context"] = context
            __props__.__dict__["exports"] = exports
            if file is None:
                file = 'Dockerfile'
            __props__.__dict__["file"] = file
            if tags is None and not opts.urn:
                raise TypeError("Missing required property 'tags'")
            __props__.__dict__["tags"] = tags
            __props__.__dict__["architecture"] = None
            __props__.__dict__["os"] = None
            __props__.__dict__["repo_digests"] = None
            __props__.__dict__["repo_tags"] = None
            __props__.__dict__["size"] = None
        super(Image, __self__).__init__(
            'docker:buildx/image:Image',
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

        __props__.__dict__["architecture"] = None
        __props__.__dict__["context"] = None
        __props__.__dict__["exports"] = None
        __props__.__dict__["file"] = None
        __props__.__dict__["os"] = None
        __props__.__dict__["repo_digests"] = None
        __props__.__dict__["repo_tags"] = None
        __props__.__dict__["size"] = None
        __props__.__dict__["tags"] = None
        return Image(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def architecture(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "architecture")

    @property
    @pulumi.getter
    def context(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """

        Contexts to use while building the image. If omitted, an empty context
        is used. If more than one value is specified, they should be of the
        form "name=value".
        """
        return pulumi.get(self, "context")

    @property
    @pulumi.getter
    def exports(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """

        Name and optionally a tag (format: "name:tag"). If outputting to a
        registry, the name should include the fully qualified registry address.
        """
        return pulumi.get(self, "exports")

    @property
    @pulumi.getter
    def file(self) -> pulumi.Output[Optional[str]]:
        """

        Name of the Dockerfile to use (default: "$PATH/Dockerfile").
        """
        return pulumi.get(self, "file")

    @property
    @pulumi.getter
    def os(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "os")

    @property
    @pulumi.getter(name="repoDigests")
    def repo_digests(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "repo_digests")

    @property
    @pulumi.getter(name="repoTags")
    def repo_tags(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "repo_tags")

    @property
    @pulumi.getter
    def size(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "size")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Sequence[str]]:
        """

        Name and optionally a tag (format: "name:tag"). If outputting to a
        registry, the name should include the fully qualified registry address.
        """
        return pulumi.get(self, "tags")
