# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class Service(pulumi.CustomResource):
    """
    This resource manages the lifecycle of a Docker service. By default, the creation, update and delete of services are detached.
    
    With the Converge Config the behavior of the `docker cli` is imitated to guarantee that
    for example, all tasks of a service are running or successfully updated or to inform `terraform` that a service could not
    be updated and was successfully rolled back.
    """
    def __init__(__self__, __name__, __opts__=None, auth=None, converge_config=None, endpoint_spec=None, labels=None, mode=None, name=None, rollback_config=None, task_spec=None, update_config=None):
        """Create a Service resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['auth'] = auth

        __props__['converge_config'] = converge_config

        __props__['endpoint_spec'] = endpoint_spec

        __props__['labels'] = labels

        __props__['mode'] = mode

        if not name:
            raise TypeError('Missing required property name')
        __props__['name'] = name

        __props__['rollback_config'] = rollback_config

        if not task_spec:
            raise TypeError('Missing required property task_spec')
        __props__['task_spec'] = task_spec

        __props__['update_config'] = update_config

        super(Service, __self__).__init__(
            'docker:swarm/service:Service',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

