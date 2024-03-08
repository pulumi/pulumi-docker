# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['RegistryImageArgs', 'RegistryImage']

@pulumi.input_type
class RegistryImageArgs:
    def __init__(__self__, *,
                 insecure_skip_verify: Optional[pulumi.Input[bool]] = None,
                 keep_remotely: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 triggers: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        The set of arguments for constructing a RegistryImage resource.
        :param pulumi.Input[bool] insecure_skip_verify: If `true`, the verification of TLS certificates of the server/registry is disabled. Defaults to `false`
        :param pulumi.Input[bool] keep_remotely: If true, then the Docker image won't be deleted on destroy operation. If this is false, it will delete the image from the docker registry on destroy operation. Defaults to `false`
        :param pulumi.Input[str] name: The name of the Docker image.
        :param pulumi.Input[Mapping[str, Any]] triggers: A map of arbitrary strings that, when changed, will force the `RegistryImage` resource to be replaced. This can be used to repush a local image
        """
        if insecure_skip_verify is not None:
            pulumi.set(__self__, "insecure_skip_verify", insecure_skip_verify)
        if keep_remotely is not None:
            pulumi.set(__self__, "keep_remotely", keep_remotely)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if triggers is not None:
            pulumi.set(__self__, "triggers", triggers)

    @property
    @pulumi.getter(name="insecureSkipVerify")
    def insecure_skip_verify(self) -> Optional[pulumi.Input[bool]]:
        """
        If `true`, the verification of TLS certificates of the server/registry is disabled. Defaults to `false`
        """
        return pulumi.get(self, "insecure_skip_verify")

    @insecure_skip_verify.setter
    def insecure_skip_verify(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "insecure_skip_verify", value)

    @property
    @pulumi.getter(name="keepRemotely")
    def keep_remotely(self) -> Optional[pulumi.Input[bool]]:
        """
        If true, then the Docker image won't be deleted on destroy operation. If this is false, it will delete the image from the docker registry on destroy operation. Defaults to `false`
        """
        return pulumi.get(self, "keep_remotely")

    @keep_remotely.setter
    def keep_remotely(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "keep_remotely", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Docker image.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def triggers(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        A map of arbitrary strings that, when changed, will force the `RegistryImage` resource to be replaced. This can be used to repush a local image
        """
        return pulumi.get(self, "triggers")

    @triggers.setter
    def triggers(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "triggers", value)


@pulumi.input_type
class _RegistryImageState:
    def __init__(__self__, *,
                 insecure_skip_verify: Optional[pulumi.Input[bool]] = None,
                 keep_remotely: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 sha256_digest: Optional[pulumi.Input[str]] = None,
                 triggers: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        Input properties used for looking up and filtering RegistryImage resources.
        :param pulumi.Input[bool] insecure_skip_verify: If `true`, the verification of TLS certificates of the server/registry is disabled. Defaults to `false`
        :param pulumi.Input[bool] keep_remotely: If true, then the Docker image won't be deleted on destroy operation. If this is false, it will delete the image from the docker registry on destroy operation. Defaults to `false`
        :param pulumi.Input[str] name: The name of the Docker image.
        :param pulumi.Input[str] sha256_digest: The sha256 digest of the image.
        :param pulumi.Input[Mapping[str, Any]] triggers: A map of arbitrary strings that, when changed, will force the `RegistryImage` resource to be replaced. This can be used to repush a local image
        """
        if insecure_skip_verify is not None:
            pulumi.set(__self__, "insecure_skip_verify", insecure_skip_verify)
        if keep_remotely is not None:
            pulumi.set(__self__, "keep_remotely", keep_remotely)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if sha256_digest is not None:
            pulumi.set(__self__, "sha256_digest", sha256_digest)
        if triggers is not None:
            pulumi.set(__self__, "triggers", triggers)

    @property
    @pulumi.getter(name="insecureSkipVerify")
    def insecure_skip_verify(self) -> Optional[pulumi.Input[bool]]:
        """
        If `true`, the verification of TLS certificates of the server/registry is disabled. Defaults to `false`
        """
        return pulumi.get(self, "insecure_skip_verify")

    @insecure_skip_verify.setter
    def insecure_skip_verify(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "insecure_skip_verify", value)

    @property
    @pulumi.getter(name="keepRemotely")
    def keep_remotely(self) -> Optional[pulumi.Input[bool]]:
        """
        If true, then the Docker image won't be deleted on destroy operation. If this is false, it will delete the image from the docker registry on destroy operation. Defaults to `false`
        """
        return pulumi.get(self, "keep_remotely")

    @keep_remotely.setter
    def keep_remotely(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "keep_remotely", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Docker image.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="sha256Digest")
    def sha256_digest(self) -> Optional[pulumi.Input[str]]:
        """
        The sha256 digest of the image.
        """
        return pulumi.get(self, "sha256_digest")

    @sha256_digest.setter
    def sha256_digest(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sha256_digest", value)

    @property
    @pulumi.getter
    def triggers(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        A map of arbitrary strings that, when changed, will force the `RegistryImage` resource to be replaced. This can be used to repush a local image
        """
        return pulumi.get(self, "triggers")

    @triggers.setter
    def triggers(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "triggers", value)


class RegistryImage(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 insecure_skip_verify: Optional[pulumi.Input[bool]] = None,
                 keep_remotely: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 triggers: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None):
        """
        <!-- Bug: Type and Name are switched -->
        Manages the lifecycle of docker image in a registry. You can upload images to a registry (= `docker push`) and also delete them again

        ## Example Usage

        Build an image with the `RemoteImage` resource and then push it to a registry:

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_docker as docker

        helloworld = docker.RegistryImage("helloworld", keep_remotely=True)
        image = docker.RemoteImage("image",
            name="registry.com/somename:1.0",
            build=docker.RemoteImageBuildArgs(
                context=f"{path['cwd']}/absolutePathToContextFolder",
            ))
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] insecure_skip_verify: If `true`, the verification of TLS certificates of the server/registry is disabled. Defaults to `false`
        :param pulumi.Input[bool] keep_remotely: If true, then the Docker image won't be deleted on destroy operation. If this is false, it will delete the image from the docker registry on destroy operation. Defaults to `false`
        :param pulumi.Input[str] name: The name of the Docker image.
        :param pulumi.Input[Mapping[str, Any]] triggers: A map of arbitrary strings that, when changed, will force the `RegistryImage` resource to be replaced. This can be used to repush a local image
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[RegistryImageArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        <!-- Bug: Type and Name are switched -->
        Manages the lifecycle of docker image in a registry. You can upload images to a registry (= `docker push`) and also delete them again

        ## Example Usage

        Build an image with the `RemoteImage` resource and then push it to a registry:

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_docker as docker

        helloworld = docker.RegistryImage("helloworld", keep_remotely=True)
        image = docker.RemoteImage("image",
            name="registry.com/somename:1.0",
            build=docker.RemoteImageBuildArgs(
                context=f"{path['cwd']}/absolutePathToContextFolder",
            ))
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param RegistryImageArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RegistryImageArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 insecure_skip_verify: Optional[pulumi.Input[bool]] = None,
                 keep_remotely: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 triggers: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RegistryImageArgs.__new__(RegistryImageArgs)

            __props__.__dict__["insecure_skip_verify"] = insecure_skip_verify
            __props__.__dict__["keep_remotely"] = keep_remotely
            __props__.__dict__["name"] = name
            __props__.__dict__["triggers"] = triggers
            __props__.__dict__["sha256_digest"] = None
        super(RegistryImage, __self__).__init__(
            'docker:index/registryImage:RegistryImage',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            insecure_skip_verify: Optional[pulumi.Input[bool]] = None,
            keep_remotely: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            sha256_digest: Optional[pulumi.Input[str]] = None,
            triggers: Optional[pulumi.Input[Mapping[str, Any]]] = None) -> 'RegistryImage':
        """
        Get an existing RegistryImage resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] insecure_skip_verify: If `true`, the verification of TLS certificates of the server/registry is disabled. Defaults to `false`
        :param pulumi.Input[bool] keep_remotely: If true, then the Docker image won't be deleted on destroy operation. If this is false, it will delete the image from the docker registry on destroy operation. Defaults to `false`
        :param pulumi.Input[str] name: The name of the Docker image.
        :param pulumi.Input[str] sha256_digest: The sha256 digest of the image.
        :param pulumi.Input[Mapping[str, Any]] triggers: A map of arbitrary strings that, when changed, will force the `RegistryImage` resource to be replaced. This can be used to repush a local image
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RegistryImageState.__new__(_RegistryImageState)

        __props__.__dict__["insecure_skip_verify"] = insecure_skip_verify
        __props__.__dict__["keep_remotely"] = keep_remotely
        __props__.__dict__["name"] = name
        __props__.__dict__["sha256_digest"] = sha256_digest
        __props__.__dict__["triggers"] = triggers
        return RegistryImage(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="insecureSkipVerify")
    def insecure_skip_verify(self) -> pulumi.Output[Optional[bool]]:
        """
        If `true`, the verification of TLS certificates of the server/registry is disabled. Defaults to `false`
        """
        return pulumi.get(self, "insecure_skip_verify")

    @property
    @pulumi.getter(name="keepRemotely")
    def keep_remotely(self) -> pulumi.Output[Optional[bool]]:
        """
        If true, then the Docker image won't be deleted on destroy operation. If this is false, it will delete the image from the docker registry on destroy operation. Defaults to `false`
        """
        return pulumi.get(self, "keep_remotely")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the Docker image.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="sha256Digest")
    def sha256_digest(self) -> pulumi.Output[str]:
        """
        The sha256 digest of the image.
        """
        return pulumi.get(self, "sha256_digest")

    @property
    @pulumi.getter
    def triggers(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        A map of arbitrary strings that, when changed, will force the `RegistryImage` resource to be replaced. This can be used to repush a local image
        """
        return pulumi.get(self, "triggers")

