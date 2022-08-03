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

__all__ = ['ProviderArgs', 'Provider']

@pulumi.input_type
class ProviderArgs:
    def __init__(__self__, *,
                 ca_material: Optional[pulumi.Input[str]] = None,
                 cert_material: Optional[pulumi.Input[str]] = None,
                 cert_path: Optional[pulumi.Input[str]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 key_material: Optional[pulumi.Input[str]] = None,
                 registry_auth: Optional[pulumi.Input[Sequence[pulumi.Input['ProviderRegistryAuthArgs']]]] = None,
                 ssh_opts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Provider resource.
        :param pulumi.Input[str] ca_material: PEM-encoded content of Docker host CA certificate
        :param pulumi.Input[str] cert_material: PEM-encoded content of Docker client certificate
        :param pulumi.Input[str] cert_path: Path to directory with Docker TLS config
        :param pulumi.Input[str] host: The Docker daemon address
        :param pulumi.Input[str] key_material: PEM-encoded content of Docker client private key
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ssh_opts: Additional SSH option flags to be appended when using `ssh://` protocol
        """
        if ca_material is not None:
            pulumi.set(__self__, "ca_material", ca_material)
        if cert_material is not None:
            pulumi.set(__self__, "cert_material", cert_material)
        if cert_path is not None:
            pulumi.set(__self__, "cert_path", cert_path)
        if host is None:
            host = (_utilities.get_env('DOCKER_HOST') or 'unix:///var/run/docker.sock')
        if host is not None:
            pulumi.set(__self__, "host", host)
        if key_material is not None:
            pulumi.set(__self__, "key_material", key_material)
        if registry_auth is not None:
            pulumi.set(__self__, "registry_auth", registry_auth)
        if ssh_opts is not None:
            pulumi.set(__self__, "ssh_opts", ssh_opts)

    @property
    @pulumi.getter(name="caMaterial")
    def ca_material(self) -> Optional[pulumi.Input[str]]:
        """
        PEM-encoded content of Docker host CA certificate
        """
        return pulumi.get(self, "ca_material")

    @ca_material.setter
    def ca_material(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ca_material", value)

    @property
    @pulumi.getter(name="certMaterial")
    def cert_material(self) -> Optional[pulumi.Input[str]]:
        """
        PEM-encoded content of Docker client certificate
        """
        return pulumi.get(self, "cert_material")

    @cert_material.setter
    def cert_material(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cert_material", value)

    @property
    @pulumi.getter(name="certPath")
    def cert_path(self) -> Optional[pulumi.Input[str]]:
        """
        Path to directory with Docker TLS config
        """
        return pulumi.get(self, "cert_path")

    @cert_path.setter
    def cert_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cert_path", value)

    @property
    @pulumi.getter
    def host(self) -> Optional[pulumi.Input[str]]:
        """
        The Docker daemon address
        """
        return pulumi.get(self, "host")

    @host.setter
    def host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host", value)

    @property
    @pulumi.getter(name="keyMaterial")
    def key_material(self) -> Optional[pulumi.Input[str]]:
        """
        PEM-encoded content of Docker client private key
        """
        return pulumi.get(self, "key_material")

    @key_material.setter
    def key_material(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_material", value)

    @property
    @pulumi.getter(name="registryAuth")
    def registry_auth(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ProviderRegistryAuthArgs']]]]:
        return pulumi.get(self, "registry_auth")

    @registry_auth.setter
    def registry_auth(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ProviderRegistryAuthArgs']]]]):
        pulumi.set(self, "registry_auth", value)

    @property
    @pulumi.getter(name="sshOpts")
    def ssh_opts(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Additional SSH option flags to be appended when using `ssh://` protocol
        """
        return pulumi.get(self, "ssh_opts")

    @ssh_opts.setter
    def ssh_opts(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "ssh_opts", value)


class Provider(pulumi.ProviderResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ca_material: Optional[pulumi.Input[str]] = None,
                 cert_material: Optional[pulumi.Input[str]] = None,
                 cert_path: Optional[pulumi.Input[str]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 key_material: Optional[pulumi.Input[str]] = None,
                 registry_auth: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ProviderRegistryAuthArgs']]]]] = None,
                 ssh_opts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        The provider type for the docker package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] ca_material: PEM-encoded content of Docker host CA certificate
        :param pulumi.Input[str] cert_material: PEM-encoded content of Docker client certificate
        :param pulumi.Input[str] cert_path: Path to directory with Docker TLS config
        :param pulumi.Input[str] host: The Docker daemon address
        :param pulumi.Input[str] key_material: PEM-encoded content of Docker client private key
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ssh_opts: Additional SSH option flags to be appended when using `ssh://` protocol
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ProviderArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The provider type for the docker package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param ProviderArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProviderArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ca_material: Optional[pulumi.Input[str]] = None,
                 cert_material: Optional[pulumi.Input[str]] = None,
                 cert_path: Optional[pulumi.Input[str]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 key_material: Optional[pulumi.Input[str]] = None,
                 registry_auth: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ProviderRegistryAuthArgs']]]]] = None,
                 ssh_opts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProviderArgs.__new__(ProviderArgs)

            __props__.__dict__["ca_material"] = ca_material
            __props__.__dict__["cert_material"] = cert_material
            __props__.__dict__["cert_path"] = cert_path
            if host is None:
                host = (_utilities.get_env('DOCKER_HOST') or 'unix:///var/run/docker.sock')
            __props__.__dict__["host"] = host
            __props__.__dict__["key_material"] = key_material
            __props__.__dict__["registry_auth"] = pulumi.Output.from_input(registry_auth).apply(pulumi.runtime.to_json) if registry_auth is not None else None
            __props__.__dict__["ssh_opts"] = pulumi.Output.from_input(ssh_opts).apply(pulumi.runtime.to_json) if ssh_opts is not None else None
        super(Provider, __self__).__init__(
            'docker',
            resource_name,
            __props__,
            opts)

    @property
    @pulumi.getter(name="caMaterial")
    def ca_material(self) -> pulumi.Output[Optional[str]]:
        """
        PEM-encoded content of Docker host CA certificate
        """
        return pulumi.get(self, "ca_material")

    @property
    @pulumi.getter(name="certMaterial")
    def cert_material(self) -> pulumi.Output[Optional[str]]:
        """
        PEM-encoded content of Docker client certificate
        """
        return pulumi.get(self, "cert_material")

    @property
    @pulumi.getter(name="certPath")
    def cert_path(self) -> pulumi.Output[Optional[str]]:
        """
        Path to directory with Docker TLS config
        """
        return pulumi.get(self, "cert_path")

    @property
    @pulumi.getter
    def host(self) -> pulumi.Output[Optional[str]]:
        """
        The Docker daemon address
        """
        return pulumi.get(self, "host")

    @property
    @pulumi.getter(name="keyMaterial")
    def key_material(self) -> pulumi.Output[Optional[str]]:
        """
        PEM-encoded content of Docker client private key
        """
        return pulumi.get(self, "key_material")

