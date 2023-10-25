# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['ServiceArgs', 'Service']

@pulumi.input_type
class ServiceArgs:
    def __init__(__self__, *,
                 task_spec: pulumi.Input['ServiceTaskSpecArgs'],
                 auth: Optional[pulumi.Input['ServiceAuthArgs']] = None,
                 converge_config: Optional[pulumi.Input['ServiceConvergeConfigArgs']] = None,
                 endpoint_spec: Optional[pulumi.Input['ServiceEndpointSpecArgs']] = None,
                 labels: Optional[pulumi.Input[Sequence[pulumi.Input['ServiceLabelArgs']]]] = None,
                 mode: Optional[pulumi.Input['ServiceModeArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rollback_config: Optional[pulumi.Input['ServiceRollbackConfigArgs']] = None,
                 update_config: Optional[pulumi.Input['ServiceUpdateConfigArgs']] = None):
        """
        The set of arguments for constructing a Service resource.
        :param pulumi.Input['ServiceTaskSpecArgs'] task_spec: User modifiable task configuration
        :param pulumi.Input['ServiceAuthArgs'] auth: Configuration for the authentication for pulling the images of the service
        :param pulumi.Input['ServiceConvergeConfigArgs'] converge_config: A configuration to ensure that a service converges aka reaches the desired that of all task up and running
        :param pulumi.Input['ServiceEndpointSpecArgs'] endpoint_spec: Properties that can be configured to access and load balance a service
        :param pulumi.Input[Sequence[pulumi.Input['ServiceLabelArgs']]] labels: User-defined key/value metadata
        :param pulumi.Input['ServiceModeArgs'] mode: The mode of resolution to use for internal load balancing between tasks
        :param pulumi.Input[str] name: A random name for the port
        :param pulumi.Input['ServiceRollbackConfigArgs'] rollback_config: Specification for the rollback strategy of the service
        :param pulumi.Input['ServiceUpdateConfigArgs'] update_config: Specification for the update strategy of the service
        """
        ServiceArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            task_spec=task_spec,
            auth=auth,
            converge_config=converge_config,
            endpoint_spec=endpoint_spec,
            labels=labels,
            mode=mode,
            name=name,
            rollback_config=rollback_config,
            update_config=update_config,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             task_spec: Optional[pulumi.Input['ServiceTaskSpecArgs']] = None,
             auth: Optional[pulumi.Input['ServiceAuthArgs']] = None,
             converge_config: Optional[pulumi.Input['ServiceConvergeConfigArgs']] = None,
             endpoint_spec: Optional[pulumi.Input['ServiceEndpointSpecArgs']] = None,
             labels: Optional[pulumi.Input[Sequence[pulumi.Input['ServiceLabelArgs']]]] = None,
             mode: Optional[pulumi.Input['ServiceModeArgs']] = None,
             name: Optional[pulumi.Input[str]] = None,
             rollback_config: Optional[pulumi.Input['ServiceRollbackConfigArgs']] = None,
             update_config: Optional[pulumi.Input['ServiceUpdateConfigArgs']] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if task_spec is None and 'taskSpec' in kwargs:
            task_spec = kwargs['taskSpec']
        if task_spec is None:
            raise TypeError("Missing 'task_spec' argument")
        if converge_config is None and 'convergeConfig' in kwargs:
            converge_config = kwargs['convergeConfig']
        if endpoint_spec is None and 'endpointSpec' in kwargs:
            endpoint_spec = kwargs['endpointSpec']
        if rollback_config is None and 'rollbackConfig' in kwargs:
            rollback_config = kwargs['rollbackConfig']
        if update_config is None and 'updateConfig' in kwargs:
            update_config = kwargs['updateConfig']

        _setter("task_spec", task_spec)
        if auth is not None:
            _setter("auth", auth)
        if converge_config is not None:
            _setter("converge_config", converge_config)
        if endpoint_spec is not None:
            _setter("endpoint_spec", endpoint_spec)
        if labels is not None:
            _setter("labels", labels)
        if mode is not None:
            _setter("mode", mode)
        if name is not None:
            _setter("name", name)
        if rollback_config is not None:
            _setter("rollback_config", rollback_config)
        if update_config is not None:
            _setter("update_config", update_config)

    @property
    @pulumi.getter(name="taskSpec")
    def task_spec(self) -> pulumi.Input['ServiceTaskSpecArgs']:
        """
        User modifiable task configuration
        """
        return pulumi.get(self, "task_spec")

    @task_spec.setter
    def task_spec(self, value: pulumi.Input['ServiceTaskSpecArgs']):
        pulumi.set(self, "task_spec", value)

    @property
    @pulumi.getter
    def auth(self) -> Optional[pulumi.Input['ServiceAuthArgs']]:
        """
        Configuration for the authentication for pulling the images of the service
        """
        return pulumi.get(self, "auth")

    @auth.setter
    def auth(self, value: Optional[pulumi.Input['ServiceAuthArgs']]):
        pulumi.set(self, "auth", value)

    @property
    @pulumi.getter(name="convergeConfig")
    def converge_config(self) -> Optional[pulumi.Input['ServiceConvergeConfigArgs']]:
        """
        A configuration to ensure that a service converges aka reaches the desired that of all task up and running
        """
        return pulumi.get(self, "converge_config")

    @converge_config.setter
    def converge_config(self, value: Optional[pulumi.Input['ServiceConvergeConfigArgs']]):
        pulumi.set(self, "converge_config", value)

    @property
    @pulumi.getter(name="endpointSpec")
    def endpoint_spec(self) -> Optional[pulumi.Input['ServiceEndpointSpecArgs']]:
        """
        Properties that can be configured to access and load balance a service
        """
        return pulumi.get(self, "endpoint_spec")

    @endpoint_spec.setter
    def endpoint_spec(self, value: Optional[pulumi.Input['ServiceEndpointSpecArgs']]):
        pulumi.set(self, "endpoint_spec", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ServiceLabelArgs']]]]:
        """
        User-defined key/value metadata
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ServiceLabelArgs']]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def mode(self) -> Optional[pulumi.Input['ServiceModeArgs']]:
        """
        The mode of resolution to use for internal load balancing between tasks
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: Optional[pulumi.Input['ServiceModeArgs']]):
        pulumi.set(self, "mode", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A random name for the port
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="rollbackConfig")
    def rollback_config(self) -> Optional[pulumi.Input['ServiceRollbackConfigArgs']]:
        """
        Specification for the rollback strategy of the service
        """
        return pulumi.get(self, "rollback_config")

    @rollback_config.setter
    def rollback_config(self, value: Optional[pulumi.Input['ServiceRollbackConfigArgs']]):
        pulumi.set(self, "rollback_config", value)

    @property
    @pulumi.getter(name="updateConfig")
    def update_config(self) -> Optional[pulumi.Input['ServiceUpdateConfigArgs']]:
        """
        Specification for the update strategy of the service
        """
        return pulumi.get(self, "update_config")

    @update_config.setter
    def update_config(self, value: Optional[pulumi.Input['ServiceUpdateConfigArgs']]):
        pulumi.set(self, "update_config", value)


@pulumi.input_type
class _ServiceState:
    def __init__(__self__, *,
                 auth: Optional[pulumi.Input['ServiceAuthArgs']] = None,
                 converge_config: Optional[pulumi.Input['ServiceConvergeConfigArgs']] = None,
                 endpoint_spec: Optional[pulumi.Input['ServiceEndpointSpecArgs']] = None,
                 labels: Optional[pulumi.Input[Sequence[pulumi.Input['ServiceLabelArgs']]]] = None,
                 mode: Optional[pulumi.Input['ServiceModeArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rollback_config: Optional[pulumi.Input['ServiceRollbackConfigArgs']] = None,
                 task_spec: Optional[pulumi.Input['ServiceTaskSpecArgs']] = None,
                 update_config: Optional[pulumi.Input['ServiceUpdateConfigArgs']] = None):
        """
        Input properties used for looking up and filtering Service resources.
        :param pulumi.Input['ServiceAuthArgs'] auth: Configuration for the authentication for pulling the images of the service
        :param pulumi.Input['ServiceConvergeConfigArgs'] converge_config: A configuration to ensure that a service converges aka reaches the desired that of all task up and running
        :param pulumi.Input['ServiceEndpointSpecArgs'] endpoint_spec: Properties that can be configured to access and load balance a service
        :param pulumi.Input[Sequence[pulumi.Input['ServiceLabelArgs']]] labels: User-defined key/value metadata
        :param pulumi.Input['ServiceModeArgs'] mode: The mode of resolution to use for internal load balancing between tasks
        :param pulumi.Input[str] name: A random name for the port
        :param pulumi.Input['ServiceRollbackConfigArgs'] rollback_config: Specification for the rollback strategy of the service
        :param pulumi.Input['ServiceTaskSpecArgs'] task_spec: User modifiable task configuration
        :param pulumi.Input['ServiceUpdateConfigArgs'] update_config: Specification for the update strategy of the service
        """
        _ServiceState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            auth=auth,
            converge_config=converge_config,
            endpoint_spec=endpoint_spec,
            labels=labels,
            mode=mode,
            name=name,
            rollback_config=rollback_config,
            task_spec=task_spec,
            update_config=update_config,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             auth: Optional[pulumi.Input['ServiceAuthArgs']] = None,
             converge_config: Optional[pulumi.Input['ServiceConvergeConfigArgs']] = None,
             endpoint_spec: Optional[pulumi.Input['ServiceEndpointSpecArgs']] = None,
             labels: Optional[pulumi.Input[Sequence[pulumi.Input['ServiceLabelArgs']]]] = None,
             mode: Optional[pulumi.Input['ServiceModeArgs']] = None,
             name: Optional[pulumi.Input[str]] = None,
             rollback_config: Optional[pulumi.Input['ServiceRollbackConfigArgs']] = None,
             task_spec: Optional[pulumi.Input['ServiceTaskSpecArgs']] = None,
             update_config: Optional[pulumi.Input['ServiceUpdateConfigArgs']] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if converge_config is None and 'convergeConfig' in kwargs:
            converge_config = kwargs['convergeConfig']
        if endpoint_spec is None and 'endpointSpec' in kwargs:
            endpoint_spec = kwargs['endpointSpec']
        if rollback_config is None and 'rollbackConfig' in kwargs:
            rollback_config = kwargs['rollbackConfig']
        if task_spec is None and 'taskSpec' in kwargs:
            task_spec = kwargs['taskSpec']
        if update_config is None and 'updateConfig' in kwargs:
            update_config = kwargs['updateConfig']

        if auth is not None:
            _setter("auth", auth)
        if converge_config is not None:
            _setter("converge_config", converge_config)
        if endpoint_spec is not None:
            _setter("endpoint_spec", endpoint_spec)
        if labels is not None:
            _setter("labels", labels)
        if mode is not None:
            _setter("mode", mode)
        if name is not None:
            _setter("name", name)
        if rollback_config is not None:
            _setter("rollback_config", rollback_config)
        if task_spec is not None:
            _setter("task_spec", task_spec)
        if update_config is not None:
            _setter("update_config", update_config)

    @property
    @pulumi.getter
    def auth(self) -> Optional[pulumi.Input['ServiceAuthArgs']]:
        """
        Configuration for the authentication for pulling the images of the service
        """
        return pulumi.get(self, "auth")

    @auth.setter
    def auth(self, value: Optional[pulumi.Input['ServiceAuthArgs']]):
        pulumi.set(self, "auth", value)

    @property
    @pulumi.getter(name="convergeConfig")
    def converge_config(self) -> Optional[pulumi.Input['ServiceConvergeConfigArgs']]:
        """
        A configuration to ensure that a service converges aka reaches the desired that of all task up and running
        """
        return pulumi.get(self, "converge_config")

    @converge_config.setter
    def converge_config(self, value: Optional[pulumi.Input['ServiceConvergeConfigArgs']]):
        pulumi.set(self, "converge_config", value)

    @property
    @pulumi.getter(name="endpointSpec")
    def endpoint_spec(self) -> Optional[pulumi.Input['ServiceEndpointSpecArgs']]:
        """
        Properties that can be configured to access and load balance a service
        """
        return pulumi.get(self, "endpoint_spec")

    @endpoint_spec.setter
    def endpoint_spec(self, value: Optional[pulumi.Input['ServiceEndpointSpecArgs']]):
        pulumi.set(self, "endpoint_spec", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ServiceLabelArgs']]]]:
        """
        User-defined key/value metadata
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ServiceLabelArgs']]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def mode(self) -> Optional[pulumi.Input['ServiceModeArgs']]:
        """
        The mode of resolution to use for internal load balancing between tasks
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: Optional[pulumi.Input['ServiceModeArgs']]):
        pulumi.set(self, "mode", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A random name for the port
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="rollbackConfig")
    def rollback_config(self) -> Optional[pulumi.Input['ServiceRollbackConfigArgs']]:
        """
        Specification for the rollback strategy of the service
        """
        return pulumi.get(self, "rollback_config")

    @rollback_config.setter
    def rollback_config(self, value: Optional[pulumi.Input['ServiceRollbackConfigArgs']]):
        pulumi.set(self, "rollback_config", value)

    @property
    @pulumi.getter(name="taskSpec")
    def task_spec(self) -> Optional[pulumi.Input['ServiceTaskSpecArgs']]:
        """
        User modifiable task configuration
        """
        return pulumi.get(self, "task_spec")

    @task_spec.setter
    def task_spec(self, value: Optional[pulumi.Input['ServiceTaskSpecArgs']]):
        pulumi.set(self, "task_spec", value)

    @property
    @pulumi.getter(name="updateConfig")
    def update_config(self) -> Optional[pulumi.Input['ServiceUpdateConfigArgs']]:
        """
        Specification for the update strategy of the service
        """
        return pulumi.get(self, "update_config")

    @update_config.setter
    def update_config(self, value: Optional[pulumi.Input['ServiceUpdateConfigArgs']]):
        pulumi.set(self, "update_config", value)


class Service(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth: Optional[pulumi.Input[pulumi.InputType['ServiceAuthArgs']]] = None,
                 converge_config: Optional[pulumi.Input[pulumi.InputType['ServiceConvergeConfigArgs']]] = None,
                 endpoint_spec: Optional[pulumi.Input[pulumi.InputType['ServiceEndpointSpecArgs']]] = None,
                 labels: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServiceLabelArgs']]]]] = None,
                 mode: Optional[pulumi.Input[pulumi.InputType['ServiceModeArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rollback_config: Optional[pulumi.Input[pulumi.InputType['ServiceRollbackConfigArgs']]] = None,
                 task_spec: Optional[pulumi.Input[pulumi.InputType['ServiceTaskSpecArgs']]] = None,
                 update_config: Optional[pulumi.Input[pulumi.InputType['ServiceUpdateConfigArgs']]] = None,
                 __props__=None):
        """
        ## Import

        ### Example Assuming you created a `service` as follows #!/bin/bash docker service create --name foo -p 8080:80 nginx prints th ID 4pcphbxkfn2rffhbhe6czytgi you provide the definition for the resource as follows terraform resource "docker_service" "foo" {

         name = "foo"

         task_spec {

         container_spec {

         image = "nginx"

         }

         }

         endpoint_spec {

         ports {

         target_port

        = "80"

         published_port = "8080"

         }

         } } then the import command is as follows #!/bin/bash

        ```sh
         $ pulumi import docker:index/service:Service foo 4pcphbxkfn2rffhbhe6czytgi
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['ServiceAuthArgs']] auth: Configuration for the authentication for pulling the images of the service
        :param pulumi.Input[pulumi.InputType['ServiceConvergeConfigArgs']] converge_config: A configuration to ensure that a service converges aka reaches the desired that of all task up and running
        :param pulumi.Input[pulumi.InputType['ServiceEndpointSpecArgs']] endpoint_spec: Properties that can be configured to access and load balance a service
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServiceLabelArgs']]]] labels: User-defined key/value metadata
        :param pulumi.Input[pulumi.InputType['ServiceModeArgs']] mode: The mode of resolution to use for internal load balancing between tasks
        :param pulumi.Input[str] name: A random name for the port
        :param pulumi.Input[pulumi.InputType['ServiceRollbackConfigArgs']] rollback_config: Specification for the rollback strategy of the service
        :param pulumi.Input[pulumi.InputType['ServiceTaskSpecArgs']] task_spec: User modifiable task configuration
        :param pulumi.Input[pulumi.InputType['ServiceUpdateConfigArgs']] update_config: Specification for the update strategy of the service
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServiceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Import

        ### Example Assuming you created a `service` as follows #!/bin/bash docker service create --name foo -p 8080:80 nginx prints th ID 4pcphbxkfn2rffhbhe6czytgi you provide the definition for the resource as follows terraform resource "docker_service" "foo" {

         name = "foo"

         task_spec {

         container_spec {

         image = "nginx"

         }

         }

         endpoint_spec {

         ports {

         target_port

        = "80"

         published_port = "8080"

         }

         } } then the import command is as follows #!/bin/bash

        ```sh
         $ pulumi import docker:index/service:Service foo 4pcphbxkfn2rffhbhe6czytgi
        ```

        :param str resource_name: The name of the resource.
        :param ServiceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServiceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            ServiceArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth: Optional[pulumi.Input[pulumi.InputType['ServiceAuthArgs']]] = None,
                 converge_config: Optional[pulumi.Input[pulumi.InputType['ServiceConvergeConfigArgs']]] = None,
                 endpoint_spec: Optional[pulumi.Input[pulumi.InputType['ServiceEndpointSpecArgs']]] = None,
                 labels: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServiceLabelArgs']]]]] = None,
                 mode: Optional[pulumi.Input[pulumi.InputType['ServiceModeArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rollback_config: Optional[pulumi.Input[pulumi.InputType['ServiceRollbackConfigArgs']]] = None,
                 task_spec: Optional[pulumi.Input[pulumi.InputType['ServiceTaskSpecArgs']]] = None,
                 update_config: Optional[pulumi.Input[pulumi.InputType['ServiceUpdateConfigArgs']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ServiceArgs.__new__(ServiceArgs)

            auth = _utilities.configure(auth, ServiceAuthArgs, True)
            __props__.__dict__["auth"] = auth
            converge_config = _utilities.configure(converge_config, ServiceConvergeConfigArgs, True)
            __props__.__dict__["converge_config"] = converge_config
            endpoint_spec = _utilities.configure(endpoint_spec, ServiceEndpointSpecArgs, True)
            __props__.__dict__["endpoint_spec"] = endpoint_spec
            __props__.__dict__["labels"] = labels
            mode = _utilities.configure(mode, ServiceModeArgs, True)
            __props__.__dict__["mode"] = mode
            __props__.__dict__["name"] = name
            rollback_config = _utilities.configure(rollback_config, ServiceRollbackConfigArgs, True)
            __props__.__dict__["rollback_config"] = rollback_config
            task_spec = _utilities.configure(task_spec, ServiceTaskSpecArgs, True)
            if task_spec is None and not opts.urn:
                raise TypeError("Missing required property 'task_spec'")
            __props__.__dict__["task_spec"] = task_spec
            update_config = _utilities.configure(update_config, ServiceUpdateConfigArgs, True)
            __props__.__dict__["update_config"] = update_config
        super(Service, __self__).__init__(
            'docker:index/service:Service',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auth: Optional[pulumi.Input[pulumi.InputType['ServiceAuthArgs']]] = None,
            converge_config: Optional[pulumi.Input[pulumi.InputType['ServiceConvergeConfigArgs']]] = None,
            endpoint_spec: Optional[pulumi.Input[pulumi.InputType['ServiceEndpointSpecArgs']]] = None,
            labels: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServiceLabelArgs']]]]] = None,
            mode: Optional[pulumi.Input[pulumi.InputType['ServiceModeArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            rollback_config: Optional[pulumi.Input[pulumi.InputType['ServiceRollbackConfigArgs']]] = None,
            task_spec: Optional[pulumi.Input[pulumi.InputType['ServiceTaskSpecArgs']]] = None,
            update_config: Optional[pulumi.Input[pulumi.InputType['ServiceUpdateConfigArgs']]] = None) -> 'Service':
        """
        Get an existing Service resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['ServiceAuthArgs']] auth: Configuration for the authentication for pulling the images of the service
        :param pulumi.Input[pulumi.InputType['ServiceConvergeConfigArgs']] converge_config: A configuration to ensure that a service converges aka reaches the desired that of all task up and running
        :param pulumi.Input[pulumi.InputType['ServiceEndpointSpecArgs']] endpoint_spec: Properties that can be configured to access and load balance a service
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServiceLabelArgs']]]] labels: User-defined key/value metadata
        :param pulumi.Input[pulumi.InputType['ServiceModeArgs']] mode: The mode of resolution to use for internal load balancing between tasks
        :param pulumi.Input[str] name: A random name for the port
        :param pulumi.Input[pulumi.InputType['ServiceRollbackConfigArgs']] rollback_config: Specification for the rollback strategy of the service
        :param pulumi.Input[pulumi.InputType['ServiceTaskSpecArgs']] task_spec: User modifiable task configuration
        :param pulumi.Input[pulumi.InputType['ServiceUpdateConfigArgs']] update_config: Specification for the update strategy of the service
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ServiceState.__new__(_ServiceState)

        __props__.__dict__["auth"] = auth
        __props__.__dict__["converge_config"] = converge_config
        __props__.__dict__["endpoint_spec"] = endpoint_spec
        __props__.__dict__["labels"] = labels
        __props__.__dict__["mode"] = mode
        __props__.__dict__["name"] = name
        __props__.__dict__["rollback_config"] = rollback_config
        __props__.__dict__["task_spec"] = task_spec
        __props__.__dict__["update_config"] = update_config
        return Service(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def auth(self) -> pulumi.Output[Optional['outputs.ServiceAuth']]:
        """
        Configuration for the authentication for pulling the images of the service
        """
        return pulumi.get(self, "auth")

    @property
    @pulumi.getter(name="convergeConfig")
    def converge_config(self) -> pulumi.Output[Optional['outputs.ServiceConvergeConfig']]:
        """
        A configuration to ensure that a service converges aka reaches the desired that of all task up and running
        """
        return pulumi.get(self, "converge_config")

    @property
    @pulumi.getter(name="endpointSpec")
    def endpoint_spec(self) -> pulumi.Output['outputs.ServiceEndpointSpec']:
        """
        Properties that can be configured to access and load balance a service
        """
        return pulumi.get(self, "endpoint_spec")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Sequence['outputs.ServiceLabel']]:
        """
        User-defined key/value metadata
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def mode(self) -> pulumi.Output['outputs.ServiceMode']:
        """
        The mode of resolution to use for internal load balancing between tasks
        """
        return pulumi.get(self, "mode")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A random name for the port
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="rollbackConfig")
    def rollback_config(self) -> pulumi.Output[Optional['outputs.ServiceRollbackConfig']]:
        """
        Specification for the rollback strategy of the service
        """
        return pulumi.get(self, "rollback_config")

    @property
    @pulumi.getter(name="taskSpec")
    def task_spec(self) -> pulumi.Output['outputs.ServiceTaskSpec']:
        """
        User modifiable task configuration
        """
        return pulumi.get(self, "task_spec")

    @property
    @pulumi.getter(name="updateConfig")
    def update_config(self) -> pulumi.Output[Optional['outputs.ServiceUpdateConfig']]:
        """
        Specification for the update strategy of the service
        """
        return pulumi.get(self, "update_config")

