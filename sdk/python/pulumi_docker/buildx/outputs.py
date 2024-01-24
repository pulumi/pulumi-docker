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

__all__ = [
    'Manifest',
    'Platform',
    'RegistryAuth',
]

@pulumi.output_type
class Manifest(dict):
    def __init__(__self__, *,
                 digest: str,
                 platform: 'outputs.Platform',
                 ref: str,
                 size: int,
                 urls: Optional[Sequence[str]] = None):
        """
        :param str ref: The manifest's ref
        """
        pulumi.set(__self__, "digest", digest)
        pulumi.set(__self__, "platform", platform)
        pulumi.set(__self__, "ref", ref)
        pulumi.set(__self__, "size", size)
        if urls is not None:
            pulumi.set(__self__, "urls", urls)

    @property
    @pulumi.getter
    def digest(self) -> str:
        return pulumi.get(self, "digest")

    @property
    @pulumi.getter
    def platform(self) -> 'outputs.Platform':
        return pulumi.get(self, "platform")

    @property
    @pulumi.getter
    def ref(self) -> str:
        """
        The manifest's ref
        """
        return pulumi.get(self, "ref")

    @property
    @pulumi.getter
    def size(self) -> int:
        return pulumi.get(self, "size")

    @property
    @pulumi.getter
    def urls(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "urls")


@pulumi.output_type
class Platform(dict):
    def __init__(__self__, *,
                 architecture: str,
                 os: str):
        pulumi.set(__self__, "architecture", architecture)
        pulumi.set(__self__, "os", os)

    @property
    @pulumi.getter
    def architecture(self) -> str:
        return pulumi.get(self, "architecture")

    @property
    @pulumi.getter
    def os(self) -> str:
        return pulumi.get(self, "os")


@pulumi.output_type
class RegistryAuth(dict):
    def __init__(__self__, *,
                 address: str,
                 password: Optional[str] = None,
                 username: Optional[str] = None):
        """
        :param str address: The registry's address (e.g. "docker.io")
        :param str password: Password or token for the registry
        :param str username: Username for the registry
        """
        pulumi.set(__self__, "address", address)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter
    def address(self) -> str:
        """
        The registry's address (e.g. "docker.io")
        """
        return pulumi.get(self, "address")

    @property
    @pulumi.getter
    def password(self) -> Optional[str]:
        """
        Password or token for the registry
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def username(self) -> Optional[str]:
        """
        Username for the registry
        """
        return pulumi.get(self, "username")


