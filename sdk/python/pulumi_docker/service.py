# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class Service(pulumi.CustomResource):
    auth: pulumi.Output[dict]
    """
    See Auth below for details.

      * `password` (`str`) - The password to use for authenticating to the registry. If this is blank, the `DOCKER_REGISTRY_PASS` is also be checked.
      * `server_address` (`str`) - The address of the registry server
      * `username` (`str`) - The username to use for authenticating to the registry. If this is blank, the `DOCKER_REGISTRY_USER` is also be checked. 
    """
    converge_config: pulumi.Output[dict]
    """
    See Converge Config below for details.

      * `delay` (`str`)
      * `timeout` (`str`)
    """
    endpoint_spec: pulumi.Output[dict]
    """
    See EndpointSpec below for details.

      * `mode` (`str`) - The mode of resolution to use for internal load balancing between tasks. `(vip|dnsrr)`. Default: `vip`.
      * `ports` (`list`) - See Ports below for details.
        * `name` (`str`) - A random name for the port.
        * `protocol` (`str`) - Protocol that can be used over this port: `tcp|udp|sctp`. Default: `tcp`.
        * `publishMode` (`str`) - Represents the mode in which the port is to be published: `ingress|host`
        * `publishedPort` (`float`) - The port on the swarm hosts. If not set the value of `target_port` will be used.
        * `targetPort` (`float`) - Port inside the container.
    """
    labels: pulumi.Output[list]
    """
    User-defined key/value metadata

      * `label` (`str`)
      * `value` (`str`)
    """
    mode: pulumi.Output[dict]
    """
    The mode of resolution to use for internal load balancing between tasks. `(vip|dnsrr)`. Default: `vip`.

      * `global` (`bool`)
      * `replicated` (`dict`)
        * `replicas` (`float`)
    """
    name: pulumi.Output[str]
    """
    A random name for the port.
    """
    rollback_config: pulumi.Output[dict]
    """
    See RollbackConfig below for details.

      * `delay` (`str`)
      * `failureAction` (`str`)
      * `maxFailureRatio` (`str`)
      * `monitor` (`str`)
      * `order` (`str`)
      * `parallelism` (`float`)
    """
    task_spec: pulumi.Output[dict]
    """
    See TaskSpec below for details.

      * `containerSpec` (`dict`)
        * `args` (`list`)
        * `commands` (`list`)
        * `configs` (`list`)
          * `configId` (`str`) - ConfigID represents the ID of the specific config.
          * `configName` (`str`) - The name of the config that this references, but internally it is just provided for lookup/display purposes
          * `fileGid` (`str`) - Represents the file GID. Defaults: `0`
          * `fileMode` (`float`) - Represents the FileMode of the file. Defaults: `0444`
          * `fileName` (`str`) - Represents the final filename in the filesystem. The specific target file that the config data is written within the docker container, e.g. `/root/config/config.json`
          * `fileUid` (`str`) - Represents the file UID. Defaults: `0`

        * `dir` (`str`)
        * `dnsConfig` (`dict`)
          * `nameservers` (`list`)
          * `options` (`list`) - The options for the logging driver, e.g.
          * `searches` (`list`)

        * `env` (`dict`)
        * `groups` (`list`)
        * `healthcheck` (`dict`)
          * `interval` (`str`)
          * `retries` (`float`)
          * `startPeriod` (`str`)
          * `tests` (`list`)
          * `timeout` (`str`)

        * `hostname` (`str`)
        * `hosts` (`list`)
          * `host` (`str`)
          * `ip` (`str`)

        * `image` (`str`)
        * `isolation` (`str`)
        * `labels` (`list`)
          * `label` (`str`)
          * `value` (`str`)

        * `mounts` (`list`)
          * `bindOptions` (`dict`)
            * `propagation` (`str`)

          * `read_only` (`bool`)
          * `source` (`str`)
          * `target` (`str`)
          * `tmpfsOptions` (`dict`)
            * `mode` (`float`) - The mode of resolution to use for internal load balancing between tasks. `(vip|dnsrr)`. Default: `vip`.
            * `sizeBytes` (`float`)

          * `type` (`str`)
          * `volumeOptions` (`dict`)
            * `driverName` (`str`)
            * `driverOptions` (`dict`)
            * `labels` (`list`)
              * `label` (`str`)
              * `value` (`str`)

            * `noCopy` (`bool`)

        * `privileges` (`dict`)
          * `credentialSpec` (`dict`)
            * `file` (`str`)
            * `registry` (`str`)

          * `seLinuxContext` (`dict`)
            * `disable` (`bool`)
            * `level` (`str`)
            * `role` (`str`)
            * `type` (`str`)
            * `user` (`str`)

        * `read_only` (`bool`)
        * `secrets` (`list`)
          * `fileGid` (`str`) - Represents the file GID. Defaults: `0`
          * `fileMode` (`float`) - Represents the FileMode of the file. Defaults: `0444`
          * `fileName` (`str`) - Represents the final filename in the filesystem. The specific target file that the config data is written within the docker container, e.g. `/root/config/config.json`
          * `fileUid` (`str`) - Represents the file UID. Defaults: `0`
          * `secretId` (`str`)
          * `secretName` (`str`)

        * `stopGracePeriod` (`str`)
        * `stopSignal` (`str`)
        * `user` (`str`)

      * `forceUpdate` (`float`)
      * `log_driver` (`dict`)
        * `name` (`str`) - A random name for the port.
        * `options` (`dict`) - The options for the logging driver, e.g.

      * `networks` (`list`)
      * `placement` (`dict`)
        * `constraints` (`list`)
        * `platforms` (`list`)
          * `architecture` (`str`)
          * `os` (`str`)

        * `prefs` (`list`)

      * `resources` (`dict`)
        * `limits` (`dict`)
          * `genericResources` (`dict`)
            * `discreteResourcesSpecs` (`list`)
            * `namedResourcesSpecs` (`list`)

          * `memoryBytes` (`float`)
          * `nanoCpus` (`float`)

        * `reservation` (`dict`)
          * `genericResources` (`dict`)
            * `discreteResourcesSpecs` (`list`)
            * `namedResourcesSpecs` (`list`)

          * `memoryBytes` (`float`)
          * `nanoCpus` (`float`)

      * `restartPolicy` (`dict`)
        * `condition` (`str`)
        * `delay` (`str`)
        * `maxAttempts` (`float`)
        * `window` (`str`)

      * `runtime` (`str`)
    """
    update_config: pulumi.Output[dict]
    """
    See UpdateConfig below for details.

      * `delay` (`str`)
      * `failureAction` (`str`)
      * `maxFailureRatio` (`str`)
      * `monitor` (`str`)
      * `order` (`str`)
      * `parallelism` (`float`)
    """
    def __init__(__self__, resource_name, opts=None, auth=None, converge_config=None, endpoint_spec=None, labels=None, mode=None, name=None, rollback_config=None, task_spec=None, update_config=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a Service resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] auth: See Auth below for details.
        :param pulumi.Input[dict] converge_config: See Converge Config below for details.
        :param pulumi.Input[dict] endpoint_spec: See EndpointSpec below for details.
        :param pulumi.Input[list] labels: User-defined key/value metadata
        :param pulumi.Input[dict] mode: The mode of resolution to use for internal load balancing between tasks. `(vip|dnsrr)`. Default: `vip`.
        :param pulumi.Input[str] name: A random name for the port.
        :param pulumi.Input[dict] rollback_config: See RollbackConfig below for details.
        :param pulumi.Input[dict] task_spec: See TaskSpec below for details.
        :param pulumi.Input[dict] update_config: See UpdateConfig below for details.

        The **auth** object supports the following:

          * `password` (`pulumi.Input[str]`) - The password to use for authenticating to the registry. If this is blank, the `DOCKER_REGISTRY_PASS` is also be checked.
          * `server_address` (`pulumi.Input[str]`) - The address of the registry server
          * `username` (`pulumi.Input[str]`) - The username to use for authenticating to the registry. If this is blank, the `DOCKER_REGISTRY_USER` is also be checked. 

        The **converge_config** object supports the following:

          * `delay` (`pulumi.Input[str]`)
          * `timeout` (`pulumi.Input[str]`)

        The **endpoint_spec** object supports the following:

          * `mode` (`pulumi.Input[str]`) - The mode of resolution to use for internal load balancing between tasks. `(vip|dnsrr)`. Default: `vip`.
          * `ports` (`pulumi.Input[list]`) - See Ports below for details.
            * `name` (`pulumi.Input[str]`) - A random name for the port.
            * `protocol` (`pulumi.Input[str]`) - Protocol that can be used over this port: `tcp|udp|sctp`. Default: `tcp`.
            * `publishMode` (`pulumi.Input[str]`) - Represents the mode in which the port is to be published: `ingress|host`
            * `publishedPort` (`pulumi.Input[float]`) - The port on the swarm hosts. If not set the value of `target_port` will be used.
            * `targetPort` (`pulumi.Input[float]`) - Port inside the container.

        The **labels** object supports the following:

          * `label` (`pulumi.Input[str]`)
          * `value` (`pulumi.Input[str]`)

        The **mode** object supports the following:

          * `global` (`pulumi.Input[bool]`)
          * `replicated` (`pulumi.Input[dict]`)
            * `replicas` (`pulumi.Input[float]`)

        The **rollback_config** object supports the following:

          * `delay` (`pulumi.Input[str]`)
          * `failureAction` (`pulumi.Input[str]`)
          * `maxFailureRatio` (`pulumi.Input[str]`)
          * `monitor` (`pulumi.Input[str]`)
          * `order` (`pulumi.Input[str]`)
          * `parallelism` (`pulumi.Input[float]`)

        The **task_spec** object supports the following:

          * `containerSpec` (`pulumi.Input[dict]`)
            * `args` (`pulumi.Input[list]`)
            * `commands` (`pulumi.Input[list]`)
            * `configs` (`pulumi.Input[list]`)
              * `configId` (`pulumi.Input[str]`) - ConfigID represents the ID of the specific config.
              * `configName` (`pulumi.Input[str]`) - The name of the config that this references, but internally it is just provided for lookup/display purposes
              * `fileGid` (`pulumi.Input[str]`) - Represents the file GID. Defaults: `0`
              * `fileMode` (`pulumi.Input[float]`) - Represents the FileMode of the file. Defaults: `0444`
              * `fileName` (`pulumi.Input[str]`) - Represents the final filename in the filesystem. The specific target file that the config data is written within the docker container, e.g. `/root/config/config.json`
              * `fileUid` (`pulumi.Input[str]`) - Represents the file UID. Defaults: `0`

            * `dir` (`pulumi.Input[str]`)
            * `dnsConfig` (`pulumi.Input[dict]`)
              * `nameservers` (`pulumi.Input[list]`)
              * `options` (`pulumi.Input[list]`) - The options for the logging driver, e.g.
              * `searches` (`pulumi.Input[list]`)

            * `env` (`pulumi.Input[dict]`)
            * `groups` (`pulumi.Input[list]`)
            * `healthcheck` (`pulumi.Input[dict]`)
              * `interval` (`pulumi.Input[str]`)
              * `retries` (`pulumi.Input[float]`)
              * `startPeriod` (`pulumi.Input[str]`)
              * `tests` (`pulumi.Input[list]`)
              * `timeout` (`pulumi.Input[str]`)

            * `hostname` (`pulumi.Input[str]`)
            * `hosts` (`pulumi.Input[list]`)
              * `host` (`pulumi.Input[str]`)
              * `ip` (`pulumi.Input[str]`)

            * `image` (`pulumi.Input[str]`)
            * `isolation` (`pulumi.Input[str]`)
            * `labels` (`pulumi.Input[list]`)
              * `label` (`pulumi.Input[str]`)
              * `value` (`pulumi.Input[str]`)

            * `mounts` (`pulumi.Input[list]`)
              * `bindOptions` (`pulumi.Input[dict]`)
                * `propagation` (`pulumi.Input[str]`)

              * `read_only` (`pulumi.Input[bool]`)
              * `source` (`pulumi.Input[str]`)
              * `target` (`pulumi.Input[str]`)
              * `tmpfsOptions` (`pulumi.Input[dict]`)
                * `mode` (`pulumi.Input[float]`) - The mode of resolution to use for internal load balancing between tasks. `(vip|dnsrr)`. Default: `vip`.
                * `sizeBytes` (`pulumi.Input[float]`)

              * `type` (`pulumi.Input[str]`)
              * `volumeOptions` (`pulumi.Input[dict]`)
                * `driverName` (`pulumi.Input[str]`)
                * `driverOptions` (`pulumi.Input[dict]`)
                * `labels` (`pulumi.Input[list]`)
                  * `label` (`pulumi.Input[str]`)
                  * `value` (`pulumi.Input[str]`)

                * `noCopy` (`pulumi.Input[bool]`)

            * `privileges` (`pulumi.Input[dict]`)
              * `credentialSpec` (`pulumi.Input[dict]`)
                * `file` (`pulumi.Input[str]`)
                * `registry` (`pulumi.Input[str]`)

              * `seLinuxContext` (`pulumi.Input[dict]`)
                * `disable` (`pulumi.Input[bool]`)
                * `level` (`pulumi.Input[str]`)
                * `role` (`pulumi.Input[str]`)
                * `type` (`pulumi.Input[str]`)
                * `user` (`pulumi.Input[str]`)

            * `read_only` (`pulumi.Input[bool]`)
            * `secrets` (`pulumi.Input[list]`)
              * `fileGid` (`pulumi.Input[str]`) - Represents the file GID. Defaults: `0`
              * `fileMode` (`pulumi.Input[float]`) - Represents the FileMode of the file. Defaults: `0444`
              * `fileName` (`pulumi.Input[str]`) - Represents the final filename in the filesystem. The specific target file that the config data is written within the docker container, e.g. `/root/config/config.json`
              * `fileUid` (`pulumi.Input[str]`) - Represents the file UID. Defaults: `0`
              * `secretId` (`pulumi.Input[str]`)
              * `secretName` (`pulumi.Input[str]`)

            * `stopGracePeriod` (`pulumi.Input[str]`)
            * `stopSignal` (`pulumi.Input[str]`)
            * `user` (`pulumi.Input[str]`)

          * `forceUpdate` (`pulumi.Input[float]`)
          * `log_driver` (`pulumi.Input[dict]`)
            * `name` (`pulumi.Input[str]`) - A random name for the port.
            * `options` (`pulumi.Input[dict]`) - The options for the logging driver, e.g.

          * `networks` (`pulumi.Input[list]`)
          * `placement` (`pulumi.Input[dict]`)
            * `constraints` (`pulumi.Input[list]`)
            * `platforms` (`pulumi.Input[list]`)
              * `architecture` (`pulumi.Input[str]`)
              * `os` (`pulumi.Input[str]`)

            * `prefs` (`pulumi.Input[list]`)

          * `resources` (`pulumi.Input[dict]`)
            * `limits` (`pulumi.Input[dict]`)
              * `genericResources` (`pulumi.Input[dict]`)
                * `discreteResourcesSpecs` (`pulumi.Input[list]`)
                * `namedResourcesSpecs` (`pulumi.Input[list]`)

              * `memoryBytes` (`pulumi.Input[float]`)
              * `nanoCpus` (`pulumi.Input[float]`)

            * `reservation` (`pulumi.Input[dict]`)
              * `genericResources` (`pulumi.Input[dict]`)
                * `discreteResourcesSpecs` (`pulumi.Input[list]`)
                * `namedResourcesSpecs` (`pulumi.Input[list]`)

              * `memoryBytes` (`pulumi.Input[float]`)
              * `nanoCpus` (`pulumi.Input[float]`)

          * `restartPolicy` (`pulumi.Input[dict]`)
            * `condition` (`pulumi.Input[str]`)
            * `delay` (`pulumi.Input[str]`)
            * `maxAttempts` (`pulumi.Input[float]`)
            * `window` (`pulumi.Input[str]`)

          * `runtime` (`pulumi.Input[str]`)

        The **update_config** object supports the following:

          * `delay` (`pulumi.Input[str]`)
          * `failureAction` (`pulumi.Input[str]`)
          * `maxFailureRatio` (`pulumi.Input[str]`)
          * `monitor` (`pulumi.Input[str]`)
          * `order` (`pulumi.Input[str]`)
          * `parallelism` (`pulumi.Input[float]`)
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['auth'] = auth
            __props__['converge_config'] = converge_config
            __props__['endpoint_spec'] = endpoint_spec
            __props__['labels'] = labels
            __props__['mode'] = mode
            __props__['name'] = name
            __props__['rollback_config'] = rollback_config
            if task_spec is None:
                raise TypeError("Missing required property 'task_spec'")
            __props__['task_spec'] = task_spec
            __props__['update_config'] = update_config
        super(Service, __self__).__init__(
            'docker:index/service:Service',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, auth=None, converge_config=None, endpoint_spec=None, labels=None, mode=None, name=None, rollback_config=None, task_spec=None, update_config=None):
        """
        Get an existing Service resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] auth: See Auth below for details.
        :param pulumi.Input[dict] converge_config: See Converge Config below for details.
        :param pulumi.Input[dict] endpoint_spec: See EndpointSpec below for details.
        :param pulumi.Input[list] labels: User-defined key/value metadata
        :param pulumi.Input[dict] mode: The mode of resolution to use for internal load balancing between tasks. `(vip|dnsrr)`. Default: `vip`.
        :param pulumi.Input[str] name: A random name for the port.
        :param pulumi.Input[dict] rollback_config: See RollbackConfig below for details.
        :param pulumi.Input[dict] task_spec: See TaskSpec below for details.
        :param pulumi.Input[dict] update_config: See UpdateConfig below for details.

        The **auth** object supports the following:

          * `password` (`pulumi.Input[str]`) - The password to use for authenticating to the registry. If this is blank, the `DOCKER_REGISTRY_PASS` is also be checked.
          * `server_address` (`pulumi.Input[str]`) - The address of the registry server
          * `username` (`pulumi.Input[str]`) - The username to use for authenticating to the registry. If this is blank, the `DOCKER_REGISTRY_USER` is also be checked. 

        The **converge_config** object supports the following:

          * `delay` (`pulumi.Input[str]`)
          * `timeout` (`pulumi.Input[str]`)

        The **endpoint_spec** object supports the following:

          * `mode` (`pulumi.Input[str]`) - The mode of resolution to use for internal load balancing between tasks. `(vip|dnsrr)`. Default: `vip`.
          * `ports` (`pulumi.Input[list]`) - See Ports below for details.
            * `name` (`pulumi.Input[str]`) - A random name for the port.
            * `protocol` (`pulumi.Input[str]`) - Protocol that can be used over this port: `tcp|udp|sctp`. Default: `tcp`.
            * `publishMode` (`pulumi.Input[str]`) - Represents the mode in which the port is to be published: `ingress|host`
            * `publishedPort` (`pulumi.Input[float]`) - The port on the swarm hosts. If not set the value of `target_port` will be used.
            * `targetPort` (`pulumi.Input[float]`) - Port inside the container.

        The **labels** object supports the following:

          * `label` (`pulumi.Input[str]`)
          * `value` (`pulumi.Input[str]`)

        The **mode** object supports the following:

          * `global` (`pulumi.Input[bool]`)
          * `replicated` (`pulumi.Input[dict]`)
            * `replicas` (`pulumi.Input[float]`)

        The **rollback_config** object supports the following:

          * `delay` (`pulumi.Input[str]`)
          * `failureAction` (`pulumi.Input[str]`)
          * `maxFailureRatio` (`pulumi.Input[str]`)
          * `monitor` (`pulumi.Input[str]`)
          * `order` (`pulumi.Input[str]`)
          * `parallelism` (`pulumi.Input[float]`)

        The **task_spec** object supports the following:

          * `containerSpec` (`pulumi.Input[dict]`)
            * `args` (`pulumi.Input[list]`)
            * `commands` (`pulumi.Input[list]`)
            * `configs` (`pulumi.Input[list]`)
              * `configId` (`pulumi.Input[str]`) - ConfigID represents the ID of the specific config.
              * `configName` (`pulumi.Input[str]`) - The name of the config that this references, but internally it is just provided for lookup/display purposes
              * `fileGid` (`pulumi.Input[str]`) - Represents the file GID. Defaults: `0`
              * `fileMode` (`pulumi.Input[float]`) - Represents the FileMode of the file. Defaults: `0444`
              * `fileName` (`pulumi.Input[str]`) - Represents the final filename in the filesystem. The specific target file that the config data is written within the docker container, e.g. `/root/config/config.json`
              * `fileUid` (`pulumi.Input[str]`) - Represents the file UID. Defaults: `0`

            * `dir` (`pulumi.Input[str]`)
            * `dnsConfig` (`pulumi.Input[dict]`)
              * `nameservers` (`pulumi.Input[list]`)
              * `options` (`pulumi.Input[list]`) - The options for the logging driver, e.g.
              * `searches` (`pulumi.Input[list]`)

            * `env` (`pulumi.Input[dict]`)
            * `groups` (`pulumi.Input[list]`)
            * `healthcheck` (`pulumi.Input[dict]`)
              * `interval` (`pulumi.Input[str]`)
              * `retries` (`pulumi.Input[float]`)
              * `startPeriod` (`pulumi.Input[str]`)
              * `tests` (`pulumi.Input[list]`)
              * `timeout` (`pulumi.Input[str]`)

            * `hostname` (`pulumi.Input[str]`)
            * `hosts` (`pulumi.Input[list]`)
              * `host` (`pulumi.Input[str]`)
              * `ip` (`pulumi.Input[str]`)

            * `image` (`pulumi.Input[str]`)
            * `isolation` (`pulumi.Input[str]`)
            * `labels` (`pulumi.Input[list]`)
              * `label` (`pulumi.Input[str]`)
              * `value` (`pulumi.Input[str]`)

            * `mounts` (`pulumi.Input[list]`)
              * `bindOptions` (`pulumi.Input[dict]`)
                * `propagation` (`pulumi.Input[str]`)

              * `read_only` (`pulumi.Input[bool]`)
              * `source` (`pulumi.Input[str]`)
              * `target` (`pulumi.Input[str]`)
              * `tmpfsOptions` (`pulumi.Input[dict]`)
                * `mode` (`pulumi.Input[float]`) - The mode of resolution to use for internal load balancing between tasks. `(vip|dnsrr)`. Default: `vip`.
                * `sizeBytes` (`pulumi.Input[float]`)

              * `type` (`pulumi.Input[str]`)
              * `volumeOptions` (`pulumi.Input[dict]`)
                * `driverName` (`pulumi.Input[str]`)
                * `driverOptions` (`pulumi.Input[dict]`)
                * `labels` (`pulumi.Input[list]`)
                  * `label` (`pulumi.Input[str]`)
                  * `value` (`pulumi.Input[str]`)

                * `noCopy` (`pulumi.Input[bool]`)

            * `privileges` (`pulumi.Input[dict]`)
              * `credentialSpec` (`pulumi.Input[dict]`)
                * `file` (`pulumi.Input[str]`)
                * `registry` (`pulumi.Input[str]`)

              * `seLinuxContext` (`pulumi.Input[dict]`)
                * `disable` (`pulumi.Input[bool]`)
                * `level` (`pulumi.Input[str]`)
                * `role` (`pulumi.Input[str]`)
                * `type` (`pulumi.Input[str]`)
                * `user` (`pulumi.Input[str]`)

            * `read_only` (`pulumi.Input[bool]`)
            * `secrets` (`pulumi.Input[list]`)
              * `fileGid` (`pulumi.Input[str]`) - Represents the file GID. Defaults: `0`
              * `fileMode` (`pulumi.Input[float]`) - Represents the FileMode of the file. Defaults: `0444`
              * `fileName` (`pulumi.Input[str]`) - Represents the final filename in the filesystem. The specific target file that the config data is written within the docker container, e.g. `/root/config/config.json`
              * `fileUid` (`pulumi.Input[str]`) - Represents the file UID. Defaults: `0`
              * `secretId` (`pulumi.Input[str]`)
              * `secretName` (`pulumi.Input[str]`)

            * `stopGracePeriod` (`pulumi.Input[str]`)
            * `stopSignal` (`pulumi.Input[str]`)
            * `user` (`pulumi.Input[str]`)

          * `forceUpdate` (`pulumi.Input[float]`)
          * `log_driver` (`pulumi.Input[dict]`)
            * `name` (`pulumi.Input[str]`) - A random name for the port.
            * `options` (`pulumi.Input[dict]`) - The options for the logging driver, e.g.

          * `networks` (`pulumi.Input[list]`)
          * `placement` (`pulumi.Input[dict]`)
            * `constraints` (`pulumi.Input[list]`)
            * `platforms` (`pulumi.Input[list]`)
              * `architecture` (`pulumi.Input[str]`)
              * `os` (`pulumi.Input[str]`)

            * `prefs` (`pulumi.Input[list]`)

          * `resources` (`pulumi.Input[dict]`)
            * `limits` (`pulumi.Input[dict]`)
              * `genericResources` (`pulumi.Input[dict]`)
                * `discreteResourcesSpecs` (`pulumi.Input[list]`)
                * `namedResourcesSpecs` (`pulumi.Input[list]`)

              * `memoryBytes` (`pulumi.Input[float]`)
              * `nanoCpus` (`pulumi.Input[float]`)

            * `reservation` (`pulumi.Input[dict]`)
              * `genericResources` (`pulumi.Input[dict]`)
                * `discreteResourcesSpecs` (`pulumi.Input[list]`)
                * `namedResourcesSpecs` (`pulumi.Input[list]`)

              * `memoryBytes` (`pulumi.Input[float]`)
              * `nanoCpus` (`pulumi.Input[float]`)

          * `restartPolicy` (`pulumi.Input[dict]`)
            * `condition` (`pulumi.Input[str]`)
            * `delay` (`pulumi.Input[str]`)
            * `maxAttempts` (`pulumi.Input[float]`)
            * `window` (`pulumi.Input[str]`)

          * `runtime` (`pulumi.Input[str]`)

        The **update_config** object supports the following:

          * `delay` (`pulumi.Input[str]`)
          * `failureAction` (`pulumi.Input[str]`)
          * `maxFailureRatio` (`pulumi.Input[str]`)
          * `monitor` (`pulumi.Input[str]`)
          * `order` (`pulumi.Input[str]`)
          * `parallelism` (`pulumi.Input[float]`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["auth"] = auth
        __props__["converge_config"] = converge_config
        __props__["endpoint_spec"] = endpoint_spec
        __props__["labels"] = labels
        __props__["mode"] = mode
        __props__["name"] = name
        __props__["rollback_config"] = rollback_config
        __props__["task_spec"] = task_spec
        __props__["update_config"] = update_config
        return Service(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

