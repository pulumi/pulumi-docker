# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import builtins
from enum import Enum

__all__ = [
    'BuilderVersion',
]


class BuilderVersion(builtins.str, Enum):
    """
    The version of the Docker builder.
    """
    BUILDER_V1 = "BuilderV1"
    """
    The first generation builder for Docker Daemon
    """
    BUILDER_BUILD_KIT = "BuilderBuildKit"
    """
    The builder based on moby/buildkit project
    """
