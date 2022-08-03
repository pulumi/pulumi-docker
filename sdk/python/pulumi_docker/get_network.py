# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetNetworkResult',
    'AwaitableGetNetworkResult',
    'get_network',
    'get_network_output',
]

@pulumi.output_type
class GetNetworkResult:
    """
    A collection of values returned by getNetwork.
    """
    def __init__(__self__, driver=None, id=None, internal=None, ipam_configs=None, name=None, options=None, scope=None):
        if driver and not isinstance(driver, str):
            raise TypeError("Expected argument 'driver' to be a str")
        pulumi.set(__self__, "driver", driver)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if internal and not isinstance(internal, bool):
            raise TypeError("Expected argument 'internal' to be a bool")
        pulumi.set(__self__, "internal", internal)
        if ipam_configs and not isinstance(ipam_configs, list):
            raise TypeError("Expected argument 'ipam_configs' to be a list")
        pulumi.set(__self__, "ipam_configs", ipam_configs)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if options and not isinstance(options, dict):
            raise TypeError("Expected argument 'options' to be a dict")
        pulumi.set(__self__, "options", options)
        if scope and not isinstance(scope, str):
            raise TypeError("Expected argument 'scope' to be a str")
        pulumi.set(__self__, "scope", scope)

    @property
    @pulumi.getter
    def driver(self) -> str:
        return pulumi.get(self, "driver")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def internal(self) -> bool:
        return pulumi.get(self, "internal")

    @property
    @pulumi.getter(name="ipamConfigs")
    def ipam_configs(self) -> Sequence['outputs.GetNetworkIpamConfigResult']:
        return pulumi.get(self, "ipam_configs")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def options(self) -> Mapping[str, Any]:
        return pulumi.get(self, "options")

    @property
    @pulumi.getter
    def scope(self) -> str:
        return pulumi.get(self, "scope")


class AwaitableGetNetworkResult(GetNetworkResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNetworkResult(
            driver=self.driver,
            id=self.id,
            internal=self.internal,
            ipam_configs=self.ipam_configs,
            name=self.name,
            options=self.options,
            scope=self.scope)


def get_network(name: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNetworkResult:
    """
    `Network` provides details about a specific Docker Network.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_docker as docker

    main = docker.get_network(name="main")
    ```
    """
    __args__ = dict()
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('docker:index/getNetwork:getNetwork', __args__, opts=opts, typ=GetNetworkResult).value

    return AwaitableGetNetworkResult(
        driver=__ret__.driver,
        id=__ret__.id,
        internal=__ret__.internal,
        ipam_configs=__ret__.ipam_configs,
        name=__ret__.name,
        options=__ret__.options,
        scope=__ret__.scope)


@_utilities.lift_output_func(get_network)
def get_network_output(name: Optional[pulumi.Input[str]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetNetworkResult]:
    """
    `Network` provides details about a specific Docker Network.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_docker as docker

    main = docker.get_network(name="main")
    ```
    """
    ...
