# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = ['ImageArgs', 'Image']

@pulumi.input_type
class ImageArgs:
    def __init__(__self__, *,
                 tags: pulumi.Input[Sequence[pulumi.Input[str]]],
                 build_args: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 builder: Optional[pulumi.Input[str]] = None,
                 cache_from: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 cache_to: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 context: Optional[pulumi.Input[str]] = None,
                 exports: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 file: Optional[pulumi.Input[str]] = None,
                 platforms: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 pull: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a Image resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: 
               Name and optionally a tag (format: "name:tag"). If outputting to a
               registry, the name should include the fully qualified registry address.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] build_args: 
               An optional map of named build-time argument variables to set during
               the Docker build. This flag allows you to pass build-time variables that
               can be accessed like environment variables inside the RUN
               instruction.
        :param pulumi.Input[str] builder: 
               Build with a specific builder instance
        :param pulumi.Input[Sequence[pulumi.Input[str]]] cache_from: 
               External cache sources (e.g., "user/app:cache", "type=local,src=path/to/dir")
        :param pulumi.Input[Sequence[pulumi.Input[str]]] cache_to: 
               Cache export destinations (e.g., "user/app:cache", "type=local,dest=path/to/dir")
        :param pulumi.Input[str] context: 
               Path to use for build context. If omitted, an empty context is used.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] exports: 
               Name and optionally a tag (format: "name:tag"). If outputting to a
               registry, the name should include the fully qualified registry address.
        :param pulumi.Input[str] file: 
               Name of the Dockerfile to use (defaults to "${context}/Dockerfile").
        :param pulumi.Input[Sequence[pulumi.Input[str]]] platforms: 
               Set target platforms for the build. Defaults to the host's platform
        :param pulumi.Input[bool] pull: 
               Always attempt to pull referenced images.
        """
        pulumi.set(__self__, "tags", tags)
        if build_args is not None:
            pulumi.set(__self__, "build_args", build_args)
        if builder is not None:
            pulumi.set(__self__, "builder", builder)
        if cache_from is not None:
            pulumi.set(__self__, "cache_from", cache_from)
        if cache_to is not None:
            pulumi.set(__self__, "cache_to", cache_to)
        if context is not None:
            pulumi.set(__self__, "context", context)
        if exports is not None:
            pulumi.set(__self__, "exports", exports)
        if file is None:
            file = 'Dockerfile'
        if file is not None:
            pulumi.set(__self__, "file", file)
        if platforms is not None:
            pulumi.set(__self__, "platforms", platforms)
        if pull is not None:
            pulumi.set(__self__, "pull", pull)

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
    @pulumi.getter(name="buildArgs")
    def build_args(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """

        An optional map of named build-time argument variables to set during
        the Docker build. This flag allows you to pass build-time variables that
        can be accessed like environment variables inside the RUN
        instruction.
        """
        return pulumi.get(self, "build_args")

    @build_args.setter
    def build_args(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "build_args", value)

    @property
    @pulumi.getter
    def builder(self) -> Optional[pulumi.Input[str]]:
        """

        Build with a specific builder instance
        """
        return pulumi.get(self, "builder")

    @builder.setter
    def builder(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "builder", value)

    @property
    @pulumi.getter(name="cacheFrom")
    def cache_from(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """

        External cache sources (e.g., "user/app:cache", "type=local,src=path/to/dir")
        """
        return pulumi.get(self, "cache_from")

    @cache_from.setter
    def cache_from(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "cache_from", value)

    @property
    @pulumi.getter(name="cacheTo")
    def cache_to(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """

        Cache export destinations (e.g., "user/app:cache", "type=local,dest=path/to/dir")
        """
        return pulumi.get(self, "cache_to")

    @cache_to.setter
    def cache_to(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "cache_to", value)

    @property
    @pulumi.getter
    def context(self) -> Optional[pulumi.Input[str]]:
        """

        Path to use for build context. If omitted, an empty context is used.
        """
        return pulumi.get(self, "context")

    @context.setter
    def context(self, value: Optional[pulumi.Input[str]]):
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

        Name of the Dockerfile to use (defaults to "${context}/Dockerfile").
        """
        return pulumi.get(self, "file")

    @file.setter
    def file(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "file", value)

    @property
    @pulumi.getter
    def platforms(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """

        Set target platforms for the build. Defaults to the host's platform
        """
        return pulumi.get(self, "platforms")

    @platforms.setter
    def platforms(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "platforms", value)

    @property
    @pulumi.getter
    def pull(self) -> Optional[pulumi.Input[bool]]:
        """

        Always attempt to pull referenced images.
        """
        return pulumi.get(self, "pull")

    @pull.setter
    def pull(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "pull", value)


class Image(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 build_args: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 builder: Optional[pulumi.Input[str]] = None,
                 cache_from: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 cache_to: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 context: Optional[pulumi.Input[str]] = None,
                 exports: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 file: Optional[pulumi.Input[str]] = None,
                 platforms: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 pull: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        A Docker image built using Buildkit

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] build_args: 
               An optional map of named build-time argument variables to set during
               the Docker build. This flag allows you to pass build-time variables that
               can be accessed like environment variables inside the RUN
               instruction.
        :param pulumi.Input[str] builder: 
               Build with a specific builder instance
        :param pulumi.Input[Sequence[pulumi.Input[str]]] cache_from: 
               External cache sources (e.g., "user/app:cache", "type=local,src=path/to/dir")
        :param pulumi.Input[Sequence[pulumi.Input[str]]] cache_to: 
               Cache export destinations (e.g., "user/app:cache", "type=local,dest=path/to/dir")
        :param pulumi.Input[str] context: 
               Path to use for build context. If omitted, an empty context is used.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] exports: 
               Name and optionally a tag (format: "name:tag"). If outputting to a
               registry, the name should include the fully qualified registry address.
        :param pulumi.Input[str] file: 
               Name of the Dockerfile to use (defaults to "${context}/Dockerfile").
        :param pulumi.Input[Sequence[pulumi.Input[str]]] platforms: 
               Set target platforms for the build. Defaults to the host's platform
        :param pulumi.Input[bool] pull: 
               Always attempt to pull referenced images.
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
                 build_args: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 builder: Optional[pulumi.Input[str]] = None,
                 cache_from: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 cache_to: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 context: Optional[pulumi.Input[str]] = None,
                 exports: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 file: Optional[pulumi.Input[str]] = None,
                 platforms: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 pull: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ImageArgs.__new__(ImageArgs)

            __props__.__dict__["build_args"] = build_args
            __props__.__dict__["builder"] = builder
            __props__.__dict__["cache_from"] = cache_from
            __props__.__dict__["cache_to"] = cache_to
            __props__.__dict__["context"] = context
            __props__.__dict__["exports"] = exports
            if file is None:
                file = 'Dockerfile'
            __props__.__dict__["file"] = file
            __props__.__dict__["platforms"] = platforms
            __props__.__dict__["pull"] = pull
            if tags is None and not opts.urn:
                raise TypeError("Missing required property 'tags'")
            __props__.__dict__["tags"] = tags
            __props__.__dict__["manifests"] = None
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

        __props__.__dict__["build_args"] = None
        __props__.__dict__["builder"] = None
        __props__.__dict__["cache_from"] = None
        __props__.__dict__["cache_to"] = None
        __props__.__dict__["context"] = None
        __props__.__dict__["exports"] = None
        __props__.__dict__["file"] = None
        __props__.__dict__["manifests"] = None
        __props__.__dict__["platforms"] = None
        __props__.__dict__["pull"] = None
        __props__.__dict__["tags"] = None
        return Image(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="buildArgs")
    def build_args(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """

        An optional map of named build-time argument variables to set during
        the Docker build. This flag allows you to pass build-time variables that
        can be accessed like environment variables inside the RUN
        instruction.
        """
        return pulumi.get(self, "build_args")

    @property
    @pulumi.getter
    def builder(self) -> pulumi.Output[Optional[str]]:
        """

        Build with a specific builder instance
        """
        return pulumi.get(self, "builder")

    @property
    @pulumi.getter(name="cacheFrom")
    def cache_from(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """

        External cache sources (e.g., "user/app:cache", "type=local,src=path/to/dir")
        """
        return pulumi.get(self, "cache_from")

    @property
    @pulumi.getter(name="cacheTo")
    def cache_to(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """

        Cache export destinations (e.g., "user/app:cache", "type=local,dest=path/to/dir")
        """
        return pulumi.get(self, "cache_to")

    @property
    @pulumi.getter
    def context(self) -> pulumi.Output[Optional[str]]:
        """

        Path to use for build context. If omitted, an empty context is used.
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

        Name of the Dockerfile to use (defaults to "${context}/Dockerfile").
        """
        return pulumi.get(self, "file")

    @property
    @pulumi.getter
    def manifests(self) -> pulumi.Output[Sequence['outputs.Manifest']]:
        return pulumi.get(self, "manifests")

    @property
    @pulumi.getter
    def platforms(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """

        Set target platforms for the build. Defaults to the host's platform
        """
        return pulumi.get(self, "platforms")

    @property
    @pulumi.getter
    def pull(self) -> pulumi.Output[Optional[bool]]:
        """

        Always attempt to pull referenced images.
        """
        return pulumi.get(self, "pull")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Sequence[str]]:
        """

        Name and optionally a tag (format: "name:tag"). If outputting to a
        registry, the name should include the fully qualified registry address.
        """
        return pulumi.get(self, "tags")
