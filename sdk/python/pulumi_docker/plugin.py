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

__all__ = ['PluginArgs', 'Plugin']

@pulumi.input_type
class PluginArgs:
    def __init__(__self__, *,
                 alias: Optional[pulumi.Input[builtins.str]] = None,
                 enable_timeout: Optional[pulumi.Input[builtins.int]] = None,
                 enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 envs: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 force_destroy: Optional[pulumi.Input[builtins.bool]] = None,
                 force_disable: Optional[pulumi.Input[builtins.bool]] = None,
                 grant_all_permissions: Optional[pulumi.Input[builtins.bool]] = None,
                 grant_permissions: Optional[pulumi.Input[Sequence[pulumi.Input['PluginGrantPermissionArgs']]]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a Plugin resource.
        :param pulumi.Input[builtins.str] alias: Docker Plugin alias
        :param pulumi.Input[builtins.int] enable_timeout: HTTP client timeout to enable the plugin
        :param pulumi.Input[builtins.bool] enabled: If `true` the plugin is enabled. Defaults to `true`
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] envs: The environment variables in the form of `KEY=VALUE`, e.g. `DEBUG=0`
        :param pulumi.Input[builtins.bool] force_destroy: If true, then the plugin is destroyed forcibly
        :param pulumi.Input[builtins.bool] force_disable: If true, then the plugin is disabled forcibly
        :param pulumi.Input[builtins.bool] grant_all_permissions: If true, grant all permissions necessary to run the plugin
        :param pulumi.Input[Sequence[pulumi.Input['PluginGrantPermissionArgs']]] grant_permissions: Grant specific permissions only
        :param pulumi.Input[builtins.str] name: Docker Plugin name
        """
        if alias is not None:
            pulumi.set(__self__, "alias", alias)
        if enable_timeout is not None:
            pulumi.set(__self__, "enable_timeout", enable_timeout)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if envs is not None:
            pulumi.set(__self__, "envs", envs)
        if force_destroy is not None:
            pulumi.set(__self__, "force_destroy", force_destroy)
        if force_disable is not None:
            pulumi.set(__self__, "force_disable", force_disable)
        if grant_all_permissions is not None:
            pulumi.set(__self__, "grant_all_permissions", grant_all_permissions)
        if grant_permissions is not None:
            pulumi.set(__self__, "grant_permissions", grant_permissions)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def alias(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Docker Plugin alias
        """
        return pulumi.get(self, "alias")

    @alias.setter
    def alias(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "alias", value)

    @property
    @pulumi.getter(name="enableTimeout")
    def enable_timeout(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        HTTP client timeout to enable the plugin
        """
        return pulumi.get(self, "enable_timeout")

    @enable_timeout.setter
    def enable_timeout(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "enable_timeout", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        If `true` the plugin is enabled. Defaults to `true`
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def envs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        The environment variables in the form of `KEY=VALUE`, e.g. `DEBUG=0`
        """
        return pulumi.get(self, "envs")

    @envs.setter
    def envs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "envs", value)

    @property
    @pulumi.getter(name="forceDestroy")
    def force_destroy(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        If true, then the plugin is destroyed forcibly
        """
        return pulumi.get(self, "force_destroy")

    @force_destroy.setter
    def force_destroy(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "force_destroy", value)

    @property
    @pulumi.getter(name="forceDisable")
    def force_disable(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        If true, then the plugin is disabled forcibly
        """
        return pulumi.get(self, "force_disable")

    @force_disable.setter
    def force_disable(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "force_disable", value)

    @property
    @pulumi.getter(name="grantAllPermissions")
    def grant_all_permissions(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        If true, grant all permissions necessary to run the plugin
        """
        return pulumi.get(self, "grant_all_permissions")

    @grant_all_permissions.setter
    def grant_all_permissions(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "grant_all_permissions", value)

    @property
    @pulumi.getter(name="grantPermissions")
    def grant_permissions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PluginGrantPermissionArgs']]]]:
        """
        Grant specific permissions only
        """
        return pulumi.get(self, "grant_permissions")

    @grant_permissions.setter
    def grant_permissions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PluginGrantPermissionArgs']]]]):
        pulumi.set(self, "grant_permissions", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Docker Plugin name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _PluginState:
    def __init__(__self__, *,
                 alias: Optional[pulumi.Input[builtins.str]] = None,
                 enable_timeout: Optional[pulumi.Input[builtins.int]] = None,
                 enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 envs: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 force_destroy: Optional[pulumi.Input[builtins.bool]] = None,
                 force_disable: Optional[pulumi.Input[builtins.bool]] = None,
                 grant_all_permissions: Optional[pulumi.Input[builtins.bool]] = None,
                 grant_permissions: Optional[pulumi.Input[Sequence[pulumi.Input['PluginGrantPermissionArgs']]]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 plugin_reference: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering Plugin resources.
        :param pulumi.Input[builtins.str] alias: Docker Plugin alias
        :param pulumi.Input[builtins.int] enable_timeout: HTTP client timeout to enable the plugin
        :param pulumi.Input[builtins.bool] enabled: If `true` the plugin is enabled. Defaults to `true`
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] envs: The environment variables in the form of `KEY=VALUE`, e.g. `DEBUG=0`
        :param pulumi.Input[builtins.bool] force_destroy: If true, then the plugin is destroyed forcibly
        :param pulumi.Input[builtins.bool] force_disable: If true, then the plugin is disabled forcibly
        :param pulumi.Input[builtins.bool] grant_all_permissions: If true, grant all permissions necessary to run the plugin
        :param pulumi.Input[Sequence[pulumi.Input['PluginGrantPermissionArgs']]] grant_permissions: Grant specific permissions only
        :param pulumi.Input[builtins.str] name: Docker Plugin name
        :param pulumi.Input[builtins.str] plugin_reference: Docker Plugin Reference
        """
        if alias is not None:
            pulumi.set(__self__, "alias", alias)
        if enable_timeout is not None:
            pulumi.set(__self__, "enable_timeout", enable_timeout)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if envs is not None:
            pulumi.set(__self__, "envs", envs)
        if force_destroy is not None:
            pulumi.set(__self__, "force_destroy", force_destroy)
        if force_disable is not None:
            pulumi.set(__self__, "force_disable", force_disable)
        if grant_all_permissions is not None:
            pulumi.set(__self__, "grant_all_permissions", grant_all_permissions)
        if grant_permissions is not None:
            pulumi.set(__self__, "grant_permissions", grant_permissions)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if plugin_reference is not None:
            pulumi.set(__self__, "plugin_reference", plugin_reference)

    @property
    @pulumi.getter
    def alias(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Docker Plugin alias
        """
        return pulumi.get(self, "alias")

    @alias.setter
    def alias(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "alias", value)

    @property
    @pulumi.getter(name="enableTimeout")
    def enable_timeout(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        HTTP client timeout to enable the plugin
        """
        return pulumi.get(self, "enable_timeout")

    @enable_timeout.setter
    def enable_timeout(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "enable_timeout", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        If `true` the plugin is enabled. Defaults to `true`
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def envs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        The environment variables in the form of `KEY=VALUE`, e.g. `DEBUG=0`
        """
        return pulumi.get(self, "envs")

    @envs.setter
    def envs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "envs", value)

    @property
    @pulumi.getter(name="forceDestroy")
    def force_destroy(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        If true, then the plugin is destroyed forcibly
        """
        return pulumi.get(self, "force_destroy")

    @force_destroy.setter
    def force_destroy(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "force_destroy", value)

    @property
    @pulumi.getter(name="forceDisable")
    def force_disable(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        If true, then the plugin is disabled forcibly
        """
        return pulumi.get(self, "force_disable")

    @force_disable.setter
    def force_disable(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "force_disable", value)

    @property
    @pulumi.getter(name="grantAllPermissions")
    def grant_all_permissions(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        If true, grant all permissions necessary to run the plugin
        """
        return pulumi.get(self, "grant_all_permissions")

    @grant_all_permissions.setter
    def grant_all_permissions(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "grant_all_permissions", value)

    @property
    @pulumi.getter(name="grantPermissions")
    def grant_permissions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PluginGrantPermissionArgs']]]]:
        """
        Grant specific permissions only
        """
        return pulumi.get(self, "grant_permissions")

    @grant_permissions.setter
    def grant_permissions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PluginGrantPermissionArgs']]]]):
        pulumi.set(self, "grant_permissions", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Docker Plugin name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="pluginReference")
    def plugin_reference(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Docker Plugin Reference
        """
        return pulumi.get(self, "plugin_reference")

    @plugin_reference.setter
    def plugin_reference(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "plugin_reference", value)


class Plugin(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alias: Optional[pulumi.Input[builtins.str]] = None,
                 enable_timeout: Optional[pulumi.Input[builtins.int]] = None,
                 enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 envs: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 force_destroy: Optional[pulumi.Input[builtins.bool]] = None,
                 force_disable: Optional[pulumi.Input[builtins.bool]] = None,
                 grant_all_permissions: Optional[pulumi.Input[builtins.bool]] = None,
                 grant_permissions: Optional[pulumi.Input[Sequence[pulumi.Input[Union['PluginGrantPermissionArgs', 'PluginGrantPermissionArgsDict']]]]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        <!-- Bug: Type and Name are switched -->
        Manages the lifecycle of a Docker plugin.

        ## Import

        #!/bin/bash

        ```sh
        $ pulumi import docker:index/plugin:Plugin sample-volume-plugin "$(docker plugin inspect -f {{.ID}} tiborvass/sample-volume-plugin:latest)"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] alias: Docker Plugin alias
        :param pulumi.Input[builtins.int] enable_timeout: HTTP client timeout to enable the plugin
        :param pulumi.Input[builtins.bool] enabled: If `true` the plugin is enabled. Defaults to `true`
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] envs: The environment variables in the form of `KEY=VALUE`, e.g. `DEBUG=0`
        :param pulumi.Input[builtins.bool] force_destroy: If true, then the plugin is destroyed forcibly
        :param pulumi.Input[builtins.bool] force_disable: If true, then the plugin is disabled forcibly
        :param pulumi.Input[builtins.bool] grant_all_permissions: If true, grant all permissions necessary to run the plugin
        :param pulumi.Input[Sequence[pulumi.Input[Union['PluginGrantPermissionArgs', 'PluginGrantPermissionArgsDict']]]] grant_permissions: Grant specific permissions only
        :param pulumi.Input[builtins.str] name: Docker Plugin name
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[PluginArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        <!-- Bug: Type and Name are switched -->
        Manages the lifecycle of a Docker plugin.

        ## Import

        #!/bin/bash

        ```sh
        $ pulumi import docker:index/plugin:Plugin sample-volume-plugin "$(docker plugin inspect -f {{.ID}} tiborvass/sample-volume-plugin:latest)"
        ```

        :param str resource_name: The name of the resource.
        :param PluginArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PluginArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alias: Optional[pulumi.Input[builtins.str]] = None,
                 enable_timeout: Optional[pulumi.Input[builtins.int]] = None,
                 enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 envs: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 force_destroy: Optional[pulumi.Input[builtins.bool]] = None,
                 force_disable: Optional[pulumi.Input[builtins.bool]] = None,
                 grant_all_permissions: Optional[pulumi.Input[builtins.bool]] = None,
                 grant_permissions: Optional[pulumi.Input[Sequence[pulumi.Input[Union['PluginGrantPermissionArgs', 'PluginGrantPermissionArgsDict']]]]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PluginArgs.__new__(PluginArgs)

            __props__.__dict__["alias"] = alias
            __props__.__dict__["enable_timeout"] = enable_timeout
            __props__.__dict__["enabled"] = enabled
            __props__.__dict__["envs"] = envs
            __props__.__dict__["force_destroy"] = force_destroy
            __props__.__dict__["force_disable"] = force_disable
            __props__.__dict__["grant_all_permissions"] = grant_all_permissions
            __props__.__dict__["grant_permissions"] = grant_permissions
            __props__.__dict__["name"] = name
            __props__.__dict__["plugin_reference"] = None
        super(Plugin, __self__).__init__(
            'docker:index/plugin:Plugin',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            alias: Optional[pulumi.Input[builtins.str]] = None,
            enable_timeout: Optional[pulumi.Input[builtins.int]] = None,
            enabled: Optional[pulumi.Input[builtins.bool]] = None,
            envs: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
            force_destroy: Optional[pulumi.Input[builtins.bool]] = None,
            force_disable: Optional[pulumi.Input[builtins.bool]] = None,
            grant_all_permissions: Optional[pulumi.Input[builtins.bool]] = None,
            grant_permissions: Optional[pulumi.Input[Sequence[pulumi.Input[Union['PluginGrantPermissionArgs', 'PluginGrantPermissionArgsDict']]]]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None,
            plugin_reference: Optional[pulumi.Input[builtins.str]] = None) -> 'Plugin':
        """
        Get an existing Plugin resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] alias: Docker Plugin alias
        :param pulumi.Input[builtins.int] enable_timeout: HTTP client timeout to enable the plugin
        :param pulumi.Input[builtins.bool] enabled: If `true` the plugin is enabled. Defaults to `true`
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] envs: The environment variables in the form of `KEY=VALUE`, e.g. `DEBUG=0`
        :param pulumi.Input[builtins.bool] force_destroy: If true, then the plugin is destroyed forcibly
        :param pulumi.Input[builtins.bool] force_disable: If true, then the plugin is disabled forcibly
        :param pulumi.Input[builtins.bool] grant_all_permissions: If true, grant all permissions necessary to run the plugin
        :param pulumi.Input[Sequence[pulumi.Input[Union['PluginGrantPermissionArgs', 'PluginGrantPermissionArgsDict']]]] grant_permissions: Grant specific permissions only
        :param pulumi.Input[builtins.str] name: Docker Plugin name
        :param pulumi.Input[builtins.str] plugin_reference: Docker Plugin Reference
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PluginState.__new__(_PluginState)

        __props__.__dict__["alias"] = alias
        __props__.__dict__["enable_timeout"] = enable_timeout
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["envs"] = envs
        __props__.__dict__["force_destroy"] = force_destroy
        __props__.__dict__["force_disable"] = force_disable
        __props__.__dict__["grant_all_permissions"] = grant_all_permissions
        __props__.__dict__["grant_permissions"] = grant_permissions
        __props__.__dict__["name"] = name
        __props__.__dict__["plugin_reference"] = plugin_reference
        return Plugin(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def alias(self) -> pulumi.Output[builtins.str]:
        """
        Docker Plugin alias
        """
        return pulumi.get(self, "alias")

    @property
    @pulumi.getter(name="enableTimeout")
    def enable_timeout(self) -> pulumi.Output[Optional[builtins.int]]:
        """
        HTTP client timeout to enable the plugin
        """
        return pulumi.get(self, "enable_timeout")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        If `true` the plugin is enabled. Defaults to `true`
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def envs(self) -> pulumi.Output[Sequence[builtins.str]]:
        """
        The environment variables in the form of `KEY=VALUE`, e.g. `DEBUG=0`
        """
        return pulumi.get(self, "envs")

    @property
    @pulumi.getter(name="forceDestroy")
    def force_destroy(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        If true, then the plugin is destroyed forcibly
        """
        return pulumi.get(self, "force_destroy")

    @property
    @pulumi.getter(name="forceDisable")
    def force_disable(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        If true, then the plugin is disabled forcibly
        """
        return pulumi.get(self, "force_disable")

    @property
    @pulumi.getter(name="grantAllPermissions")
    def grant_all_permissions(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        If true, grant all permissions necessary to run the plugin
        """
        return pulumi.get(self, "grant_all_permissions")

    @property
    @pulumi.getter(name="grantPermissions")
    def grant_permissions(self) -> pulumi.Output[Optional[Sequence['outputs.PluginGrantPermission']]]:
        """
        Grant specific permissions only
        """
        return pulumi.get(self, "grant_permissions")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        Docker Plugin name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="pluginReference")
    def plugin_reference(self) -> pulumi.Output[builtins.str]:
        """
        Docker Plugin Reference
        """
        return pulumi.get(self, "plugin_reference")

