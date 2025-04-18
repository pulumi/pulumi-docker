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
from .. import _utilities
from . import outputs

caMaterial: Optional[str]
"""
PEM-encoded content of Docker host CA certificate
"""

certMaterial: Optional[str]
"""
PEM-encoded content of Docker client certificate
"""

certPath: Optional[str]
"""
Path to directory with Docker TLS config
"""

host: Optional[str]
"""
The Docker daemon address
"""

keyMaterial: Optional[str]
"""
PEM-encoded content of Docker client private key
"""

registryAuth: Optional[str]

sshOpts: Optional[str]
"""
Additional SSH option flags to be appended when using `ssh://` protocol
"""

