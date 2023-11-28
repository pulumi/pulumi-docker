import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
/**
 * <!-- Bug: Type and Name are switched -->
 * Manages the lifecycle of a Docker container.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as docker from "@pulumi/docker";
 *
 * // Find the latest Ubuntu precise image.
 * const ubuntuRemoteImage = new docker.RemoteImage("ubuntuRemoteImage", {name: "ubuntu:precise"});
 * // Start a container
 * const ubuntuContainer = new docker.Container("ubuntuContainer", {image: ubuntuRemoteImage.imageId});
 * ```
 *
 * ## Import
 *
 * ### Example Assuming you created a `container` as follows #!/bin/bash docker run --name foo -p8080:80 -d nginx
 *
 * prints the container ID
 *
 * 9a550c0f0163d39d77222d3efd58701b625d47676c25c686c95b5b92d1cba6fd you provide the definition for the resource as follows terraform resource "docker_container" "foo" {
 *
 *  name
 *
 * = "foo"
 *
 *  image = "nginx"
 *
 *  ports {
 *
 *  internal = "80"
 *
 *  external = "8080"
 *
 *  } } then the import command is as follows #!/bin/bash
 *
 * ```sh
 *  $ pulumi import docker:index/container:Container foo 9a550c0f0163d39d77222d3efd58701b625d47676c25c686c95b5b92d1cba6fd
 * ```
 */
export declare class Container extends pulumi.CustomResource {
    /**
     * Get an existing Container resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ContainerState, opts?: pulumi.CustomResourceOptions): Container;
    /**
     * Returns true if the given object is an instance of Container.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    static isInstance(obj: any): obj is Container;
    /**
     * If `true` attach to the container after its creation and waits the end of its execution. Defaults to `false`.
     */
    readonly attach: pulumi.Output<boolean | undefined>;
    /**
     * The network bridge of the container as read from its NetworkSettings.
     */
    readonly bridge: pulumi.Output<string>;
    /**
     * Add or drop certrain linux capabilities.
     */
    readonly capabilities: pulumi.Output<outputs.ContainerCapabilities | undefined>;
    /**
     * Cgroup namespace mode to use for the container. Possible values are: `private`, `host`.
     */
    readonly cgroupnsMode: pulumi.Output<string | undefined>;
    /**
     * The command to use to start the container. For example, to run `/usr/bin/myprogram -f baz.conf` set the command to be `["/usr/bin/myprogram","-f","baz.con"]`.
     */
    readonly command: pulumi.Output<string[]>;
    /**
     * The logs of the container if its execution is done (`attach` must be disabled).
     */
    readonly containerLogs: pulumi.Output<string>;
    /**
     * The total number of milliseconds to wait for the container to reach status 'running'
     */
    readonly containerReadRefreshTimeoutMilliseconds: pulumi.Output<number | undefined>;
    /**
     * A comma-separated list or hyphen-separated range of CPUs a container can use, e.g. `0-1`.
     */
    readonly cpuSet: pulumi.Output<string | undefined>;
    /**
     * CPU shares (relative weight) for the container.
     */
    readonly cpuShares: pulumi.Output<number | undefined>;
    /**
     * If defined will attempt to stop the container before destroying. Container will be destroyed after `n` seconds or on successful stop.
     */
    readonly destroyGraceSeconds: pulumi.Output<number | undefined>;
    /**
     * Bind devices to the container.
     */
    readonly devices: pulumi.Output<outputs.ContainerDevice[] | undefined>;
    /**
     * DNS servers to use.
     */
    readonly dns: pulumi.Output<string[] | undefined>;
    /**
     * DNS options used by the DNS provider(s), see `resolv.conf` documentation for valid list of options.
     */
    readonly dnsOpts: pulumi.Output<string[] | undefined>;
    /**
     * DNS search domains that are used when bare unqualified hostnames are used inside of the container.
     */
    readonly dnsSearches: pulumi.Output<string[] | undefined>;
    /**
     * Domain name of the container.
     */
    readonly domainname: pulumi.Output<string | undefined>;
    /**
     * The command to use as the Entrypoint for the container. The Entrypoint allows you to configure a container to run as an executable. For example, to run `/usr/bin/myprogram` when starting a container, set the entrypoint to be `"/usr/bin/myprogra"]`.
     */
    readonly entrypoints: pulumi.Output<string[]>;
    /**
     * Environment variables to set in the form of `KEY=VALUE`, e.g. `DEBUG=0`
     */
    readonly envs: pulumi.Output<string[]>;
    /**
     * The exit code of the container if its execution is done (`mustRun` must be disabled).
     */
    readonly exitCode: pulumi.Output<number>;
    /**
     * GPU devices to add to the container. Currently, only the value `all` is supported. Passing any other value will result in unexpected behavior.
     */
    readonly gpus: pulumi.Output<string | undefined>;
    /**
     * Additional groups for the container user
     */
    readonly groupAdds: pulumi.Output<string[] | undefined>;
    /**
     * A test to perform to check that the container is healthy
     */
    readonly healthcheck: pulumi.Output<outputs.ContainerHealthcheck>;
    /**
     * Hostname of the container.
     */
    readonly hostname: pulumi.Output<string>;
    /**
     * Hostname to add
     */
    readonly hosts: pulumi.Output<outputs.ContainerHost[] | undefined>;
    /**
     * The ID of the image to back this container. The easiest way to get this value is to use the `docker.RemoteImage` resource as is shown in the example.
     */
    readonly image: pulumi.Output<string>;
    /**
     * Configured whether an init process should be injected for this container. If unset this will default to the `dockerd` defaults.
     */
    readonly init: pulumi.Output<boolean>;
    /**
     * IPC sharing mode for the container. Possible values are: `none`, `private`, `shareable`, `container:<name|id>` or `host`.
     */
    readonly ipcMode: pulumi.Output<string>;
    /**
     * User-defined key/value metadata.
     */
    readonly labels: pulumi.Output<outputs.ContainerLabel[]>;
    /**
     * The logging driver to use for the container.
     */
    readonly logDriver: pulumi.Output<string>;
    /**
     * Key/value pairs to use as options for the logging driver.
     */
    readonly logOpts: pulumi.Output<{
        [key: string]: any;
    } | undefined>;
    /**
     * Save the container logs (`attach` must be enabled). Defaults to `false`.
     */
    readonly logs: pulumi.Output<boolean | undefined>;
    /**
     * The maximum amount of times to an attempt a restart when `restart` is set to 'on-failure'.
     */
    readonly maxRetryCount: pulumi.Output<number | undefined>;
    /**
     * The memory limit for the container in MBs.
     */
    readonly memory: pulumi.Output<number | undefined>;
    /**
     * The total memory limit (memory + swap) for the container in MBs. This setting may compute to `-1` after `pulumi up` if the target host doesn't support memory swap, when that is the case docker will use a soft limitation.
     */
    readonly memorySwap: pulumi.Output<number | undefined>;
    /**
     * Specification for mounts to be added to containers created as part of the service.
     */
    readonly mounts: pulumi.Output<outputs.ContainerMount[] | undefined>;
    /**
     * If `true`, then the Docker container will be kept running. If `false`, then as long as the container exists, Terraform
     * assumes it is successful. Defaults to `true`.
     */
    readonly mustRun: pulumi.Output<boolean | undefined>;
    /**
     * The name or id of the network to use. You can use `name` or `id` attribute from a `docker.Network` resource.
     */
    readonly name: pulumi.Output<string>;
    /**
     * The data of the networks the container is connected to.
     */
    readonly networkDatas: pulumi.Output<outputs.ContainerNetworkData[]>;
    /**
     * Network mode of the container.
     */
    readonly networkMode: pulumi.Output<string | undefined>;
    /**
     * The networks the container is attached to
     */
    readonly networksAdvanced: pulumi.Output<outputs.ContainerNetworksAdvanced[] | undefined>;
    /**
     * he PID (Process) Namespace mode for the container. Either `container:<name|id>` or `host`.
     */
    readonly pidMode: pulumi.Output<string | undefined>;
    /**
     * Publish a container's port(s) to the host.
     */
    readonly ports: pulumi.Output<outputs.ContainerPort[] | undefined>;
    /**
     * If `true`, the container runs in privileged mode.
     */
    readonly privileged: pulumi.Output<boolean | undefined>;
    /**
     * Publish all ports of the container.
     */
    readonly publishAllPorts: pulumi.Output<boolean | undefined>;
    /**
     * Whether the mount should be read-only.
     */
    readonly readOnly: pulumi.Output<boolean | undefined>;
    /**
     * If `true`, it will remove anonymous volumes associated with the container. Defaults to `true`.
     */
    readonly removeVolumes: pulumi.Output<boolean | undefined>;
    /**
     * The restart policy for the container. Must be one of 'no', 'on-failure', 'always', 'unless-stopped'. Defaults to `no`.
     */
    readonly restart: pulumi.Output<string | undefined>;
    /**
     * If `true`, then the container will be automatically removed when it exits. Defaults to `false`.
     */
    readonly rm: pulumi.Output<boolean | undefined>;
    /**
     * Runtime to use for the container.
     */
    readonly runtime: pulumi.Output<string>;
    /**
     * List of string values to customize labels for MLS systems, such as SELinux. See https://docs.docker.com/engine/reference/run/#security-configuration.
     */
    readonly securityOpts: pulumi.Output<string[]>;
    /**
     * Size of `/dev/shm` in MBs.
     */
    readonly shmSize: pulumi.Output<number>;
    /**
     * If `true`, then the Docker container will be started after creation. If `false`, then the container is only created. Defaults to `true`.
     */
    readonly start: pulumi.Output<boolean | undefined>;
    /**
     * If `true`, keep STDIN open even if not attached (`docker run -i`). Defaults to `false`.
     */
    readonly stdinOpen: pulumi.Output<boolean | undefined>;
    /**
     * Signal to stop a container (default `SIGTERM`).
     */
    readonly stopSignal: pulumi.Output<string>;
    /**
     * Timeout (in seconds) to stop a container.
     */
    readonly stopTimeout: pulumi.Output<number>;
    /**
     * Key/value pairs for the storage driver options, e.g. `size`: `120G`
     */
    readonly storageOpts: pulumi.Output<{
        [key: string]: any;
    } | undefined>;
    /**
     * A map of kernel parameters (sysctls) to set in the container.
     */
    readonly sysctls: pulumi.Output<{
        [key: string]: any;
    } | undefined>;
    /**
     * A map of container directories which should be replaced by `tmpfs mounts`, and their corresponding mount options.
     */
    readonly tmpfs: pulumi.Output<{
        [key: string]: any;
    } | undefined>;
    /**
     * If `true`, allocate a pseudo-tty (`docker run -t`). Defaults to `false`.
     */
    readonly tty: pulumi.Output<boolean | undefined>;
    /**
     * Ulimit options to add.
     */
    readonly ulimits: pulumi.Output<outputs.ContainerUlimit[] | undefined>;
    /**
     * Specifies files to upload to the container before starting it. Only one of `content` or `contentBase64` can be set and at least one of them has to be set.
     */
    readonly uploads: pulumi.Output<outputs.ContainerUpload[] | undefined>;
    /**
     * User used for run the first process. Format is `user` or `user:group` which user and group can be passed literraly or by name.
     */
    readonly user: pulumi.Output<string | undefined>;
    /**
     * Sets the usernamespace mode for the container when usernamespace remapping option is enabled.
     */
    readonly usernsMode: pulumi.Output<string | undefined>;
    /**
     * Spec for mounting volumes in the container.
     */
    readonly volumes: pulumi.Output<outputs.ContainerVolume[] | undefined>;
    /**
     * If `true`, then the Docker container is waited for being healthy state after creation. If `false`, then the container health state is not checked. Defaults to `false`.
     */
    readonly wait: pulumi.Output<boolean | undefined>;
    /**
     * The timeout in seconds to wait the container to be healthy after creation. Defaults to `60`.
     */
    readonly waitTimeout: pulumi.Output<number | undefined>;
    /**
     * The working directory for commands to run in.
     */
    readonly workingDir: pulumi.Output<string | undefined>;
    /**
     * Create a Container resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ContainerArgs, opts?: pulumi.CustomResourceOptions);
}
/**
 * Input properties used for looking up and filtering Container resources.
 */
export interface ContainerState {
    /**
     * If `true` attach to the container after its creation and waits the end of its execution. Defaults to `false`.
     */
    attach?: pulumi.Input<boolean>;
    /**
     * The network bridge of the container as read from its NetworkSettings.
     */
    bridge?: pulumi.Input<string>;
    /**
     * Add or drop certrain linux capabilities.
     */
    capabilities?: pulumi.Input<inputs.ContainerCapabilities>;
    /**
     * Cgroup namespace mode to use for the container. Possible values are: `private`, `host`.
     */
    cgroupnsMode?: pulumi.Input<string>;
    /**
     * The command to use to start the container. For example, to run `/usr/bin/myprogram -f baz.conf` set the command to be `["/usr/bin/myprogram","-f","baz.con"]`.
     */
    command?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The logs of the container if its execution is done (`attach` must be disabled).
     */
    containerLogs?: pulumi.Input<string>;
    /**
     * The total number of milliseconds to wait for the container to reach status 'running'
     */
    containerReadRefreshTimeoutMilliseconds?: pulumi.Input<number>;
    /**
     * A comma-separated list or hyphen-separated range of CPUs a container can use, e.g. `0-1`.
     */
    cpuSet?: pulumi.Input<string>;
    /**
     * CPU shares (relative weight) for the container.
     */
    cpuShares?: pulumi.Input<number>;
    /**
     * If defined will attempt to stop the container before destroying. Container will be destroyed after `n` seconds or on successful stop.
     */
    destroyGraceSeconds?: pulumi.Input<number>;
    /**
     * Bind devices to the container.
     */
    devices?: pulumi.Input<pulumi.Input<inputs.ContainerDevice>[]>;
    /**
     * DNS servers to use.
     */
    dns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * DNS options used by the DNS provider(s), see `resolv.conf` documentation for valid list of options.
     */
    dnsOpts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * DNS search domains that are used when bare unqualified hostnames are used inside of the container.
     */
    dnsSearches?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Domain name of the container.
     */
    domainname?: pulumi.Input<string>;
    /**
     * The command to use as the Entrypoint for the container. The Entrypoint allows you to configure a container to run as an executable. For example, to run `/usr/bin/myprogram` when starting a container, set the entrypoint to be `"/usr/bin/myprogra"]`.
     */
    entrypoints?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Environment variables to set in the form of `KEY=VALUE`, e.g. `DEBUG=0`
     */
    envs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The exit code of the container if its execution is done (`mustRun` must be disabled).
     */
    exitCode?: pulumi.Input<number>;
    /**
     * GPU devices to add to the container. Currently, only the value `all` is supported. Passing any other value will result in unexpected behavior.
     */
    gpus?: pulumi.Input<string>;
    /**
     * Additional groups for the container user
     */
    groupAdds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A test to perform to check that the container is healthy
     */
    healthcheck?: pulumi.Input<inputs.ContainerHealthcheck>;
    /**
     * Hostname of the container.
     */
    hostname?: pulumi.Input<string>;
    /**
     * Hostname to add
     */
    hosts?: pulumi.Input<pulumi.Input<inputs.ContainerHost>[]>;
    /**
     * The ID of the image to back this container. The easiest way to get this value is to use the `docker.RemoteImage` resource as is shown in the example.
     */
    image?: pulumi.Input<string>;
    /**
     * Configured whether an init process should be injected for this container. If unset this will default to the `dockerd` defaults.
     */
    init?: pulumi.Input<boolean>;
    /**
     * IPC sharing mode for the container. Possible values are: `none`, `private`, `shareable`, `container:<name|id>` or `host`.
     */
    ipcMode?: pulumi.Input<string>;
    /**
     * User-defined key/value metadata.
     */
    labels?: pulumi.Input<pulumi.Input<inputs.ContainerLabel>[]>;
    /**
     * The logging driver to use for the container.
     */
    logDriver?: pulumi.Input<string>;
    /**
     * Key/value pairs to use as options for the logging driver.
     */
    logOpts?: pulumi.Input<{
        [key: string]: any;
    }>;
    /**
     * Save the container logs (`attach` must be enabled). Defaults to `false`.
     */
    logs?: pulumi.Input<boolean>;
    /**
     * The maximum amount of times to an attempt a restart when `restart` is set to 'on-failure'.
     */
    maxRetryCount?: pulumi.Input<number>;
    /**
     * The memory limit for the container in MBs.
     */
    memory?: pulumi.Input<number>;
    /**
     * The total memory limit (memory + swap) for the container in MBs. This setting may compute to `-1` after `pulumi up` if the target host doesn't support memory swap, when that is the case docker will use a soft limitation.
     */
    memorySwap?: pulumi.Input<number>;
    /**
     * Specification for mounts to be added to containers created as part of the service.
     */
    mounts?: pulumi.Input<pulumi.Input<inputs.ContainerMount>[]>;
    /**
     * If `true`, then the Docker container will be kept running. If `false`, then as long as the container exists, Terraform
     * assumes it is successful. Defaults to `true`.
     */
    mustRun?: pulumi.Input<boolean>;
    /**
     * The name or id of the network to use. You can use `name` or `id` attribute from a `docker.Network` resource.
     */
    name?: pulumi.Input<string>;
    /**
     * The data of the networks the container is connected to.
     */
    networkDatas?: pulumi.Input<pulumi.Input<inputs.ContainerNetworkData>[]>;
    /**
     * Network mode of the container.
     */
    networkMode?: pulumi.Input<string>;
    /**
     * The networks the container is attached to
     */
    networksAdvanced?: pulumi.Input<pulumi.Input<inputs.ContainerNetworksAdvanced>[]>;
    /**
     * he PID (Process) Namespace mode for the container. Either `container:<name|id>` or `host`.
     */
    pidMode?: pulumi.Input<string>;
    /**
     * Publish a container's port(s) to the host.
     */
    ports?: pulumi.Input<pulumi.Input<inputs.ContainerPort>[]>;
    /**
     * If `true`, the container runs in privileged mode.
     */
    privileged?: pulumi.Input<boolean>;
    /**
     * Publish all ports of the container.
     */
    publishAllPorts?: pulumi.Input<boolean>;
    /**
     * Whether the mount should be read-only.
     */
    readOnly?: pulumi.Input<boolean>;
    /**
     * If `true`, it will remove anonymous volumes associated with the container. Defaults to `true`.
     */
    removeVolumes?: pulumi.Input<boolean>;
    /**
     * The restart policy for the container. Must be one of 'no', 'on-failure', 'always', 'unless-stopped'. Defaults to `no`.
     */
    restart?: pulumi.Input<string>;
    /**
     * If `true`, then the container will be automatically removed when it exits. Defaults to `false`.
     */
    rm?: pulumi.Input<boolean>;
    /**
     * Runtime to use for the container.
     */
    runtime?: pulumi.Input<string>;
    /**
     * List of string values to customize labels for MLS systems, such as SELinux. See https://docs.docker.com/engine/reference/run/#security-configuration.
     */
    securityOpts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Size of `/dev/shm` in MBs.
     */
    shmSize?: pulumi.Input<number>;
    /**
     * If `true`, then the Docker container will be started after creation. If `false`, then the container is only created. Defaults to `true`.
     */
    start?: pulumi.Input<boolean>;
    /**
     * If `true`, keep STDIN open even if not attached (`docker run -i`). Defaults to `false`.
     */
    stdinOpen?: pulumi.Input<boolean>;
    /**
     * Signal to stop a container (default `SIGTERM`).
     */
    stopSignal?: pulumi.Input<string>;
    /**
     * Timeout (in seconds) to stop a container.
     */
    stopTimeout?: pulumi.Input<number>;
    /**
     * Key/value pairs for the storage driver options, e.g. `size`: `120G`
     */
    storageOpts?: pulumi.Input<{
        [key: string]: any;
    }>;
    /**
     * A map of kernel parameters (sysctls) to set in the container.
     */
    sysctls?: pulumi.Input<{
        [key: string]: any;
    }>;
    /**
     * A map of container directories which should be replaced by `tmpfs mounts`, and their corresponding mount options.
     */
    tmpfs?: pulumi.Input<{
        [key: string]: any;
    }>;
    /**
     * If `true`, allocate a pseudo-tty (`docker run -t`). Defaults to `false`.
     */
    tty?: pulumi.Input<boolean>;
    /**
     * Ulimit options to add.
     */
    ulimits?: pulumi.Input<pulumi.Input<inputs.ContainerUlimit>[]>;
    /**
     * Specifies files to upload to the container before starting it. Only one of `content` or `contentBase64` can be set and at least one of them has to be set.
     */
    uploads?: pulumi.Input<pulumi.Input<inputs.ContainerUpload>[]>;
    /**
     * User used for run the first process. Format is `user` or `user:group` which user and group can be passed literraly or by name.
     */
    user?: pulumi.Input<string>;
    /**
     * Sets the usernamespace mode for the container when usernamespace remapping option is enabled.
     */
    usernsMode?: pulumi.Input<string>;
    /**
     * Spec for mounting volumes in the container.
     */
    volumes?: pulumi.Input<pulumi.Input<inputs.ContainerVolume>[]>;
    /**
     * If `true`, then the Docker container is waited for being healthy state after creation. If `false`, then the container health state is not checked. Defaults to `false`.
     */
    wait?: pulumi.Input<boolean>;
    /**
     * The timeout in seconds to wait the container to be healthy after creation. Defaults to `60`.
     */
    waitTimeout?: pulumi.Input<number>;
    /**
     * The working directory for commands to run in.
     */
    workingDir?: pulumi.Input<string>;
}
/**
 * The set of arguments for constructing a Container resource.
 */
export interface ContainerArgs {
    /**
     * If `true` attach to the container after its creation and waits the end of its execution. Defaults to `false`.
     */
    attach?: pulumi.Input<boolean>;
    /**
     * Add or drop certrain linux capabilities.
     */
    capabilities?: pulumi.Input<inputs.ContainerCapabilities>;
    /**
     * Cgroup namespace mode to use for the container. Possible values are: `private`, `host`.
     */
    cgroupnsMode?: pulumi.Input<string>;
    /**
     * The command to use to start the container. For example, to run `/usr/bin/myprogram -f baz.conf` set the command to be `["/usr/bin/myprogram","-f","baz.con"]`.
     */
    command?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The total number of milliseconds to wait for the container to reach status 'running'
     */
    containerReadRefreshTimeoutMilliseconds?: pulumi.Input<number>;
    /**
     * A comma-separated list or hyphen-separated range of CPUs a container can use, e.g. `0-1`.
     */
    cpuSet?: pulumi.Input<string>;
    /**
     * CPU shares (relative weight) for the container.
     */
    cpuShares?: pulumi.Input<number>;
    /**
     * If defined will attempt to stop the container before destroying. Container will be destroyed after `n` seconds or on successful stop.
     */
    destroyGraceSeconds?: pulumi.Input<number>;
    /**
     * Bind devices to the container.
     */
    devices?: pulumi.Input<pulumi.Input<inputs.ContainerDevice>[]>;
    /**
     * DNS servers to use.
     */
    dns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * DNS options used by the DNS provider(s), see `resolv.conf` documentation for valid list of options.
     */
    dnsOpts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * DNS search domains that are used when bare unqualified hostnames are used inside of the container.
     */
    dnsSearches?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Domain name of the container.
     */
    domainname?: pulumi.Input<string>;
    /**
     * The command to use as the Entrypoint for the container. The Entrypoint allows you to configure a container to run as an executable. For example, to run `/usr/bin/myprogram` when starting a container, set the entrypoint to be `"/usr/bin/myprogra"]`.
     */
    entrypoints?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Environment variables to set in the form of `KEY=VALUE`, e.g. `DEBUG=0`
     */
    envs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * GPU devices to add to the container. Currently, only the value `all` is supported. Passing any other value will result in unexpected behavior.
     */
    gpus?: pulumi.Input<string>;
    /**
     * Additional groups for the container user
     */
    groupAdds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A test to perform to check that the container is healthy
     */
    healthcheck?: pulumi.Input<inputs.ContainerHealthcheck>;
    /**
     * Hostname of the container.
     */
    hostname?: pulumi.Input<string>;
    /**
     * Hostname to add
     */
    hosts?: pulumi.Input<pulumi.Input<inputs.ContainerHost>[]>;
    /**
     * The ID of the image to back this container. The easiest way to get this value is to use the `docker.RemoteImage` resource as is shown in the example.
     */
    image: pulumi.Input<string>;
    /**
     * Configured whether an init process should be injected for this container. If unset this will default to the `dockerd` defaults.
     */
    init?: pulumi.Input<boolean>;
    /**
     * IPC sharing mode for the container. Possible values are: `none`, `private`, `shareable`, `container:<name|id>` or `host`.
     */
    ipcMode?: pulumi.Input<string>;
    /**
     * User-defined key/value metadata.
     */
    labels?: pulumi.Input<pulumi.Input<inputs.ContainerLabel>[]>;
    /**
     * The logging driver to use for the container.
     */
    logDriver?: pulumi.Input<string>;
    /**
     * Key/value pairs to use as options for the logging driver.
     */
    logOpts?: pulumi.Input<{
        [key: string]: any;
    }>;
    /**
     * Save the container logs (`attach` must be enabled). Defaults to `false`.
     */
    logs?: pulumi.Input<boolean>;
    /**
     * The maximum amount of times to an attempt a restart when `restart` is set to 'on-failure'.
     */
    maxRetryCount?: pulumi.Input<number>;
    /**
     * The memory limit for the container in MBs.
     */
    memory?: pulumi.Input<number>;
    /**
     * The total memory limit (memory + swap) for the container in MBs. This setting may compute to `-1` after `pulumi up` if the target host doesn't support memory swap, when that is the case docker will use a soft limitation.
     */
    memorySwap?: pulumi.Input<number>;
    /**
     * Specification for mounts to be added to containers created as part of the service.
     */
    mounts?: pulumi.Input<pulumi.Input<inputs.ContainerMount>[]>;
    /**
     * If `true`, then the Docker container will be kept running. If `false`, then as long as the container exists, Terraform
     * assumes it is successful. Defaults to `true`.
     */
    mustRun?: pulumi.Input<boolean>;
    /**
     * The name or id of the network to use. You can use `name` or `id` attribute from a `docker.Network` resource.
     */
    name?: pulumi.Input<string>;
    /**
     * Network mode of the container.
     */
    networkMode?: pulumi.Input<string>;
    /**
     * The networks the container is attached to
     */
    networksAdvanced?: pulumi.Input<pulumi.Input<inputs.ContainerNetworksAdvanced>[]>;
    /**
     * he PID (Process) Namespace mode for the container. Either `container:<name|id>` or `host`.
     */
    pidMode?: pulumi.Input<string>;
    /**
     * Publish a container's port(s) to the host.
     */
    ports?: pulumi.Input<pulumi.Input<inputs.ContainerPort>[]>;
    /**
     * If `true`, the container runs in privileged mode.
     */
    privileged?: pulumi.Input<boolean>;
    /**
     * Publish all ports of the container.
     */
    publishAllPorts?: pulumi.Input<boolean>;
    /**
     * Whether the mount should be read-only.
     */
    readOnly?: pulumi.Input<boolean>;
    /**
     * If `true`, it will remove anonymous volumes associated with the container. Defaults to `true`.
     */
    removeVolumes?: pulumi.Input<boolean>;
    /**
     * The restart policy for the container. Must be one of 'no', 'on-failure', 'always', 'unless-stopped'. Defaults to `no`.
     */
    restart?: pulumi.Input<string>;
    /**
     * If `true`, then the container will be automatically removed when it exits. Defaults to `false`.
     */
    rm?: pulumi.Input<boolean>;
    /**
     * Runtime to use for the container.
     */
    runtime?: pulumi.Input<string>;
    /**
     * List of string values to customize labels for MLS systems, such as SELinux. See https://docs.docker.com/engine/reference/run/#security-configuration.
     */
    securityOpts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Size of `/dev/shm` in MBs.
     */
    shmSize?: pulumi.Input<number>;
    /**
     * If `true`, then the Docker container will be started after creation. If `false`, then the container is only created. Defaults to `true`.
     */
    start?: pulumi.Input<boolean>;
    /**
     * If `true`, keep STDIN open even if not attached (`docker run -i`). Defaults to `false`.
     */
    stdinOpen?: pulumi.Input<boolean>;
    /**
     * Signal to stop a container (default `SIGTERM`).
     */
    stopSignal?: pulumi.Input<string>;
    /**
     * Timeout (in seconds) to stop a container.
     */
    stopTimeout?: pulumi.Input<number>;
    /**
     * Key/value pairs for the storage driver options, e.g. `size`: `120G`
     */
    storageOpts?: pulumi.Input<{
        [key: string]: any;
    }>;
    /**
     * A map of kernel parameters (sysctls) to set in the container.
     */
    sysctls?: pulumi.Input<{
        [key: string]: any;
    }>;
    /**
     * A map of container directories which should be replaced by `tmpfs mounts`, and their corresponding mount options.
     */
    tmpfs?: pulumi.Input<{
        [key: string]: any;
    }>;
    /**
     * If `true`, allocate a pseudo-tty (`docker run -t`). Defaults to `false`.
     */
    tty?: pulumi.Input<boolean>;
    /**
     * Ulimit options to add.
     */
    ulimits?: pulumi.Input<pulumi.Input<inputs.ContainerUlimit>[]>;
    /**
     * Specifies files to upload to the container before starting it. Only one of `content` or `contentBase64` can be set and at least one of them has to be set.
     */
    uploads?: pulumi.Input<pulumi.Input<inputs.ContainerUpload>[]>;
    /**
     * User used for run the first process. Format is `user` or `user:group` which user and group can be passed literraly or by name.
     */
    user?: pulumi.Input<string>;
    /**
     * Sets the usernamespace mode for the container when usernamespace remapping option is enabled.
     */
    usernsMode?: pulumi.Input<string>;
    /**
     * Spec for mounting volumes in the container.
     */
    volumes?: pulumi.Input<pulumi.Input<inputs.ContainerVolume>[]>;
    /**
     * If `true`, then the Docker container is waited for being healthy state after creation. If `false`, then the container health state is not checked. Defaults to `false`.
     */
    wait?: pulumi.Input<boolean>;
    /**
     * The timeout in seconds to wait the container to be healthy after creation. Defaults to `60`.
     */
    waitTimeout?: pulumi.Input<number>;
    /**
     * The working directory for commands to run in.
     */
    workingDir?: pulumi.Input<string>;
}
