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
]

@pulumi.output_type
class Manifest(dict):
    def __init__(__self__, *,
                 digest: str,
                 platform: 'outputs.Platform',
                 ref: str,
                 size: int):
        """
        :param str ref: The manifest's ref
        """
        pulumi.set(__self__, "digest", digest)
        pulumi.set(__self__, "platform", platform)
        pulumi.set(__self__, "ref", ref)
        pulumi.set(__self__, "size", size)

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


