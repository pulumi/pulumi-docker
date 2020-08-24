# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from . import _utilities, _tables

__all__ = ['ServiceConfig']


class ServiceConfig(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 data: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a ServiceConfig resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] data: The base64 encoded data of the config.
        :param pulumi.Input[str] name: The name of the Docker config.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if data is None:
                raise TypeError("Missing required property 'data'")
            __props__['data'] = data
            __props__['name'] = name
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
        :param pulumi.Input[str] data: The base64 encoded data of the config.
        :param pulumi.Input[str] name: The name of the Docker config.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["data"] = data
        __props__["name"] = name
        return ServiceConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def data(self) -> str:
        """
        The base64 encoded data of the config.
        """
        return pulumi.get(self, "data")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the Docker config.
        """
        return pulumi.get(self, "name")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

