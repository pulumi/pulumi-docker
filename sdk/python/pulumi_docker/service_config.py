# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ServiceConfigArgs', 'ServiceConfig']

@pulumi.input_type
class ServiceConfigArgs:
    def __init__(__self__, *,
                 data: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ServiceConfig resource.
        :param pulumi.Input[str] data: Base64-url-safe-encoded config data
        :param pulumi.Input[str] name: User-defined name of the config
        """
        pulumi.set(__self__, "data", data)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def data(self) -> pulumi.Input[str]:
        """
        Base64-url-safe-encoded config data
        """
        return pulumi.get(self, "data")

    @data.setter
    def data(self, value: pulumi.Input[str]):
        pulumi.set(self, "data", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        User-defined name of the config
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _ServiceConfigState:
    def __init__(__self__, *,
                 data: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ServiceConfig resources.
        :param pulumi.Input[str] data: Base64-url-safe-encoded config data
        :param pulumi.Input[str] name: User-defined name of the config
        """
        if data is not None:
            pulumi.set(__self__, "data", data)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def data(self) -> Optional[pulumi.Input[str]]:
        """
        Base64-url-safe-encoded config data
        """
        return pulumi.get(self, "data")

    @data.setter
    def data(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        User-defined name of the config
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class ServiceConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 data: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Import

        ### Example Assuming you created a `config` as follows #!/bin/bash printf '{"a":"b"}' | docker config create foo - # prints the id

        08c26c477474478d971139f750984775a7f019dbe8a2e7f09d66a187c009e66d you provide the definition for the resource as follows terraform resource "docker_config" "foo" {

         name = "foo"

         data = base64encode("{\"a\"\"b\"}") } then the import command is as follows #!/bin/bash

        ```sh
         $ pulumi import docker:index/serviceConfig:ServiceConfig foo 08c26c477474478d971139f750984775a7f019dbe8a2e7f09d66a187c009e66d
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] data: Base64-url-safe-encoded config data
        :param pulumi.Input[str] name: User-defined name of the config
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServiceConfigArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Import

        ### Example Assuming you created a `config` as follows #!/bin/bash printf '{"a":"b"}' | docker config create foo - # prints the id

        08c26c477474478d971139f750984775a7f019dbe8a2e7f09d66a187c009e66d you provide the definition for the resource as follows terraform resource "docker_config" "foo" {

         name = "foo"

         data = base64encode("{\"a\"\"b\"}") } then the import command is as follows #!/bin/bash

        ```sh
         $ pulumi import docker:index/serviceConfig:ServiceConfig foo 08c26c477474478d971139f750984775a7f019dbe8a2e7f09d66a187c009e66d
        ```

        :param str resource_name: The name of the resource.
        :param ServiceConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServiceConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 data: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ServiceConfigArgs.__new__(ServiceConfigArgs)

            if data is None and not opts.urn:
                raise TypeError("Missing required property 'data'")
            __props__.__dict__["data"] = data
            __props__.__dict__["name"] = name
        super(ServiceConfig, __self__).__init__(
            'docker:index/serviceConfig:ServiceConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            data: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'ServiceConfig':
        """
        Get an existing ServiceConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] data: Base64-url-safe-encoded config data
        :param pulumi.Input[str] name: User-defined name of the config
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ServiceConfigState.__new__(_ServiceConfigState)

        __props__.__dict__["data"] = data
        __props__.__dict__["name"] = name
        return ServiceConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def data(self) -> pulumi.Output[str]:
        """
        Base64-url-safe-encoded config data
        """
        return pulumi.get(self, "data")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        User-defined name of the config
        """
        return pulumi.get(self, "name")

