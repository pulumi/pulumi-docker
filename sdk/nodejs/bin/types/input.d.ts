import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as enums from "../types/enums";
/**
 * Contains a list of images to reference when building using a cache
 */
export interface CacheFrom {
    /**
     * Specifies cached images
     */
    images?: pulumi.Input<pulumi.Input<string>[]>;
}
export interface ContainerCapabilities {
    /**
     * List of linux capabilities to add.
     */
    adds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of linux capabilities to drop.
     */
    drops?: pulumi.Input<pulumi.Input<string>[]>;
}
export interface ContainerDevice {
    /**
     * The path in the container where the device will be bound.
     */
    containerPath?: pulumi.Input<string>;
    /**
     * The path on the host where the device is located.
     */
    hostPath: pulumi.Input<string>;
    /**
     * The cgroup permissions given to the container to access the device. Defaults to `rwm`.
     */
    permissions?: pulumi.Input<string>;
}
export interface ContainerHealthcheck {
    /**
     * Time between running the check (ms|s|m|h). Defaults to `0s`.
     */
    interval?: pulumi.Input<string>;
    /**
     * Consecutive failures needed to report unhealthy. Defaults to `0`.
     */
    retries?: pulumi.Input<number>;
    /**
     * Start period for the container to initialize before counting retries towards unstable (ms|s|m|h). Defaults to `0s`.
     */
    startPeriod?: pulumi.Input<string>;
    /**
     * Command to run to check health. For example, to run `curl -f localhost/health` set the command to be `["CMD", "curl", "-f", "localhost/health"]`.
     */
    tests: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Maximum time to allow one check to run (ms|s|m|h). Defaults to `0s`.
     */
    timeout?: pulumi.Input<string>;
}
export interface ContainerHost {
    /**
     * Hostname to add
     */
    host: pulumi.Input<string>;
    /**
     * IP address this hostname should resolve to.
     */
    ip: pulumi.Input<string>;
}
export interface ContainerLabel {
    /**
     * Name of the label
     */
    label: pulumi.Input<string>;
    /**
     * Value of the label
     */
    value: pulumi.Input<string>;
}
export interface ContainerMount {
    /**
     * Optional configuration for the bind type.
     */
    bindOptions?: pulumi.Input<inputs.ContainerMountBindOptions>;
    /**
     * Whether the mount should be read-only.
     */
    readOnly?: pulumi.Input<boolean>;
    /**
     * Mount source (e.g. a volume name, a host path).
     */
    source?: pulumi.Input<string>;
    /**
     * Container path
     */
    target: pulumi.Input<string>;
    /**
     * Optional configuration for the tmpfs type.
     */
    tmpfsOptions?: pulumi.Input<inputs.ContainerMountTmpfsOptions>;
    /**
     * The mount type
     */
    type: pulumi.Input<string>;
    /**
     * Optional configuration for the volume type.
     */
    volumeOptions?: pulumi.Input<inputs.ContainerMountVolumeOptions>;
}
export interface ContainerMountBindOptions {
    /**
     * A propagation mode with the value.
     */
    propagation?: pulumi.Input<string>;
}
export interface ContainerMountTmpfsOptions {
    /**
     * The permission mode for the tmpfs mount in an integer.
     */
    mode?: pulumi.Input<number>;
    /**
     * The size for the tmpfs mount in bytes.
     */
    sizeBytes?: pulumi.Input<number>;
}
export interface ContainerMountVolumeOptions {
    /**
     * Name of the driver to use to create the volume.
     */
    driverName?: pulumi.Input<string>;
    /**
     * key/value map of driver specific options.
     */
    driverOptions?: pulumi.Input<{
        [key: string]: pulumi.Input<string>;
    }>;
    /**
     * User-defined key/value metadata.
     */
    labels?: pulumi.Input<pulumi.Input<inputs.ContainerMountVolumeOptionsLabel>[]>;
    /**
     * Populate volume with data from the target.
     */
    noCopy?: pulumi.Input<boolean>;
}
export interface ContainerMountVolumeOptionsLabel {
    /**
     * Name of the label
     */
    label: pulumi.Input<string>;
    /**
     * Value of the label
     */
    value: pulumi.Input<string>;
}
export interface ContainerNetworkData {
    gateway?: pulumi.Input<string>;
    globalIpv6Address?: pulumi.Input<string>;
    globalIpv6PrefixLength?: pulumi.Input<number>;
    ipAddress?: pulumi.Input<string>;
    ipPrefixLength?: pulumi.Input<number>;
    ipv6Gateway?: pulumi.Input<string>;
    macAddress?: pulumi.Input<string>;
    networkName?: pulumi.Input<string>;
}
export interface ContainerNetworksAdvanced {
    /**
     * The network aliases of the container in the specific network.
     */
    aliases?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The IPV4 address of the container in the specific network.
     */
    ipv4Address?: pulumi.Input<string>;
    /**
     * The IPV6 address of the container in the specific network.
     */
    ipv6Address?: pulumi.Input<string>;
    /**
     * The name or id of the network to use. You can use `name` or `id` attribute from a `docker.Network` resource.
     */
    name: pulumi.Input<string>;
}
export interface ContainerPort {
    /**
     * Port exposed out of the container. If not given a free random port `>= 32768` will be used.
     */
    external?: pulumi.Input<number>;
    /**
     * Port within the container.
     */
    internal: pulumi.Input<number>;
    /**
     * IP address/mask that can access this port. Defaults to `0.0.0.0`.
     */
    ip?: pulumi.Input<string>;
    /**
     * Protocol that can be used over this port. Defaults to `tcp`.
     */
    protocol?: pulumi.Input<string>;
}
export interface ContainerUlimit {
    /**
     * The hard limit
     */
    hard: pulumi.Input<number>;
    /**
     * The name of the ulimit
     */
    name: pulumi.Input<string>;
    /**
     * The soft limit
     */
    soft: pulumi.Input<number>;
}
export interface ContainerUpload {
    /**
     * Literal string value to use as the object content, which will be uploaded as UTF-8-encoded text. Conflicts with `contentBase64` & `source`
     */
    content?: pulumi.Input<string>;
    contentBase64?: pulumi.Input<string>;
    /**
     * If `true`, the file will be uploaded with user executable permission. Defaults to `false`.
     */
    executable?: pulumi.Input<boolean>;
    /**
     * Path to the file in the container where is upload goes to
     */
    file: pulumi.Input<string>;
    /**
     * A filename that references a file which will be uploaded as the object content. This allows for large file uploads that do not get stored in state. Conflicts with `content` & `contentBase64`
     */
    source?: pulumi.Input<string>;
    /**
     * If using `source`, this will force an update if the file content has updated but the filename has not.
     */
    sourceHash?: pulumi.Input<string>;
}
export interface ContainerVolume {
    /**
     * The path in the container where the volume will be mounted.
     */
    containerPath?: pulumi.Input<string>;
    /**
     * The container where the volume is coming from.
     */
    fromContainer?: pulumi.Input<string>;
    /**
     * The path on the host where the volume is coming from.
     */
    hostPath?: pulumi.Input<string>;
    /**
     * If `true`, this volume will be readonly. Defaults to `false`.
     */
    readOnly?: pulumi.Input<boolean>;
    /**
     * The name of the docker volume which should be mounted.
     */
    volumeName?: pulumi.Input<string>;
}
/**
 * The Docker build context
 */
export interface DockerBuild {
    /**
     * An optional map of named build-time argument variables to set during the Docker build. This flag allows you to pass build-time variables that can be accessed like environment variables inside the RUN instruction.
     */
    args?: pulumi.Input<{
        [key: string]: pulumi.Input<string>;
    }>;
    /**
     * The version of the Docker builder.
     */
    builderVersion?: pulumi.Input<enums.BuilderVersion>;
    /**
     * A list of image names to use as build cache. Images provided must have a cache manifest. Must provide authentication to cache registry.
     */
    cacheFrom?: pulumi.Input<inputs.CacheFrom>;
    /**
     * The path to the build context to use.
     */
    context?: pulumi.Input<string>;
    /**
     * The path to the Dockerfile to use.
     */
    dockerfile?: pulumi.Input<string>;
    /**
     * The architecture of the platform you want to build this image for, e.g. `linux/arm64`.
     */
    platform?: pulumi.Input<string>;
    /**
     * The target of the Dockerfile to build
     */
    target?: pulumi.Input<string>;
}
export interface NetworkIpamConfig {
    /**
     * Auxiliary IPv4 or IPv6 addresses used by Network driver
     */
    auxAddress?: pulumi.Input<{
        [key: string]: any;
    }>;
    /**
     * The IP address of the gateway
     */
    gateway?: pulumi.Input<string>;
    /**
     * The ip range in CIDR form
     */
    ipRange?: pulumi.Input<string>;
    /**
     * The subnet in CIDR form
     */
    subnet?: pulumi.Input<string>;
}
export interface NetworkLabel {
    /**
     * Name of the label
     */
    label: pulumi.Input<string>;
    /**
     * Value of the label
     */
    value: pulumi.Input<string>;
}
export interface PluginGrantPermission {
    /**
     * The name of the permission
     */
    name: pulumi.Input<string>;
    /**
     * The value of the permission
     */
    values: pulumi.Input<pulumi.Input<string>[]>;
}
export interface ProviderRegistryAuth {
    address: pulumi.Input<string>;
    authDisabled?: pulumi.Input<boolean>;
    configFile?: pulumi.Input<string>;
    configFileContent?: pulumi.Input<string>;
    password?: pulumi.Input<string>;
    username?: pulumi.Input<string>;
}
/**
 * Describes a Docker container registry
 */
export interface Registry {
    /**
     * The password to authenticate to the registry. Does not cause image rebuild when changed.
     */
    password?: pulumi.Input<string>;
    /**
     * The URL of the Docker registry server
     */
    server?: pulumi.Input<string>;
    /**
     * The username to authenticate to the registry. Does not cause image rebuild when changed.
     */
    username?: pulumi.Input<string>;
}
export interface RemoteImageBuild {
    /**
     * The configuration for the authentication
     */
    authConfigs?: pulumi.Input<pulumi.Input<inputs.RemoteImageBuildAuthConfig>[]>;
    /**
     * Set build-time variables
     */
    buildArg?: pulumi.Input<{
        [key: string]: pulumi.Input<string>;
    }>;
    /**
     * Pairs for build-time variables in the form TODO
     */
    buildArgs?: pulumi.Input<{
        [key: string]: pulumi.Input<string>;
    }>;
    /**
     * BuildID is an optional identifier that can be passed together with the build request. The same identifier can be used to gracefully cancel the build with the cancel request.
     */
    buildId?: pulumi.Input<string>;
    /**
     * Images to consider as cache sources
     */
    cacheFroms?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Optional parent cgroup for the container
     */
    cgroupParent?: pulumi.Input<string>;
    /**
     * Value to specify the build context. Currently, only a `PATH` context is supported. You can use the helper function '${path.cwd}/context-dir'. Please see https://docs.docker.com/build/building/context/ for more information about build contexts.
     */
    context: pulumi.Input<string>;
    /**
     * The length of a CPU period in microseconds
     */
    cpuPeriod?: pulumi.Input<number>;
    /**
     * Microseconds of CPU time that the container can get in a CPU period
     */
    cpuQuota?: pulumi.Input<number>;
    /**
     * CPUs in which to allow execution (e.g., `0-3`, `0`, `1`)
     */
    cpuSetCpus?: pulumi.Input<string>;
    /**
     * MEMs in which to allow execution (`0-3`, `0`, `1`)
     */
    cpuSetMems?: pulumi.Input<string>;
    /**
     * CPU shares (relative weight)
     */
    cpuShares?: pulumi.Input<number>;
    /**
     * Name of the Dockerfile. Defaults to `Dockerfile`.
     */
    dockerfile?: pulumi.Input<string>;
    /**
     * A list of hostnames/IP mappings to add to the container’s /etc/hosts file. Specified in the form ["hostname:IP"]
     */
    extraHosts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Always remove intermediate containers
     */
    forceRemove?: pulumi.Input<boolean>;
    /**
     * Isolation represents the isolation technology of a container. The supported values are
     */
    isolation?: pulumi.Input<string>;
    /**
     * Set metadata for an image
     */
    label?: pulumi.Input<{
        [key: string]: pulumi.Input<string>;
    }>;
    /**
     * User-defined key/value metadata
     */
    labels?: pulumi.Input<{
        [key: string]: pulumi.Input<string>;
    }>;
    /**
     * Set memory limit for build
     */
    memory?: pulumi.Input<number>;
    /**
     * Total memory (memory + swap), -1 to enable unlimited swap
     */
    memorySwap?: pulumi.Input<number>;
    /**
     * Set the networking mode for the RUN instructions during build
     */
    networkMode?: pulumi.Input<string>;
    /**
     * Do not use the cache when building the image
     */
    noCache?: pulumi.Input<boolean>;
    /**
     * Set platform if server is multi-platform capable
     */
    platform?: pulumi.Input<string>;
    /**
     * Attempt to pull the image even if an older image exists locally
     */
    pullParent?: pulumi.Input<boolean>;
    /**
     * A Git repository URI or HTTP/HTTPS context URI
     */
    remoteContext?: pulumi.Input<string>;
    /**
     * Remove intermediate containers after a successful build. Defaults to `true`.
     */
    remove?: pulumi.Input<boolean>;
    /**
     * The security options
     */
    securityOpts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Set an ID for the build session
     */
    sessionId?: pulumi.Input<string>;
    /**
     * Size of /dev/shm in bytes. The size must be greater than 0
     */
    shmSize?: pulumi.Input<number>;
    /**
     * If true the new layers are squashed into a new image with a single new layer
     */
    squash?: pulumi.Input<boolean>;
    /**
     * Suppress the build output and print image ID on success
     */
    suppressOutput?: pulumi.Input<boolean>;
    /**
     * Name and optionally a tag in the 'name:tag' format
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Set the target build stage to build
     */
    target?: pulumi.Input<string>;
    /**
     * Configuration for ulimits
     */
    ulimits?: pulumi.Input<pulumi.Input<inputs.RemoteImageBuildUlimit>[]>;
    /**
     * Version of the underlying builder to use
     */
    version?: pulumi.Input<string>;
}
export interface RemoteImageBuildAuthConfig {
    /**
     * the auth token
     */
    auth?: pulumi.Input<string>;
    /**
     * the user emal
     */
    email?: pulumi.Input<string>;
    /**
     * hostname of the registry
     */
    hostName: pulumi.Input<string>;
    /**
     * the identity token
     */
    identityToken?: pulumi.Input<string>;
    /**
     * the registry password
     */
    password?: pulumi.Input<string>;
    /**
     * the registry token
     */
    registryToken?: pulumi.Input<string>;
    /**
     * the server address
     */
    serverAddress?: pulumi.Input<string>;
    /**
     * the registry user name
     */
    userName?: pulumi.Input<string>;
}
export interface RemoteImageBuildUlimit {
    /**
     * soft limit
     */
    hard: pulumi.Input<number>;
    /**
     * type of ulimit, e.g. `nofile`
     */
    name: pulumi.Input<string>;
    /**
     * hard limit
     */
    soft: pulumi.Input<number>;
}
export interface SecretLabel {
    /**
     * Name of the label
     */
    label: pulumi.Input<string>;
    /**
     * Value of the label
     */
    value: pulumi.Input<string>;
}
export interface ServiceAuth {
    /**
     * The password
     */
    password?: pulumi.Input<string>;
    /**
     * The address of the server for the authentication
     */
    serverAddress: pulumi.Input<string>;
    /**
     * The username
     */
    username?: pulumi.Input<string>;
}
export interface ServiceConvergeConfig {
    /**
     * The interval to check if the desired state is reached `(ms|s)`. Defaults to `7s`.
     */
    delay?: pulumi.Input<string>;
    /**
     * The timeout of the service to reach the desired state `(s|m)`. Defaults to `3m`
     */
    timeout?: pulumi.Input<string>;
}
export interface ServiceEndpointSpec {
    /**
     * The mode of resolution to use for internal load balancing between tasks
     */
    mode?: pulumi.Input<string>;
    /**
     * List of exposed ports that this service is accessible on from the outside. Ports can only be provided if 'vip' resolution mode is used
     */
    ports?: pulumi.Input<pulumi.Input<inputs.ServiceEndpointSpecPort>[]>;
}
export interface ServiceEndpointSpecPort {
    /**
     * A random name for the port
     */
    name?: pulumi.Input<string>;
    /**
     * Rrepresents the protocol of a port: `tcp`, `udp` or `sctp`. Defaults to `tcp`.
     */
    protocol?: pulumi.Input<string>;
    /**
     * Represents the mode in which the port is to be published: 'ingress' or 'host'. Defaults to `ingress`.
     */
    publishMode?: pulumi.Input<string>;
    /**
     * The port on the swarm hosts
     */
    publishedPort?: pulumi.Input<number>;
    /**
     * The port inside the container
     */
    targetPort: pulumi.Input<number>;
}
export interface ServiceLabel {
    /**
     * Name of the label
     */
    label: pulumi.Input<string>;
    /**
     * Value of the label
     */
    value: pulumi.Input<string>;
}
export interface ServiceMode {
    /**
     * The global service mode. Defaults to `false`
     */
    global?: pulumi.Input<boolean>;
    /**
     * The replicated service mode
     */
    replicated?: pulumi.Input<inputs.ServiceModeReplicated>;
}
export interface ServiceModeReplicated {
    /**
     * The amount of replicas of the service. Defaults to `1`
     */
    replicas?: pulumi.Input<number>;
}
export interface ServiceRollbackConfig {
    /**
     * Delay between task rollbacks (ns|us|ms|s|m|h). Defaults to `0s`.
     */
    delay?: pulumi.Input<string>;
    /**
     * Action on rollback failure: pause | continue. Defaults to `pause`.
     */
    failureAction?: pulumi.Input<string>;
    /**
     * Failure rate to tolerate during a rollback. Defaults to `0.0`.
     */
    maxFailureRatio?: pulumi.Input<string>;
    /**
     * Duration after each task rollback to monitor for failure (ns|us|ms|s|m|h). Defaults to `5s`.
     */
    monitor?: pulumi.Input<string>;
    /**
     * Rollback order: either 'stop-first' or 'start-first'. Defaults to `stop-first`.
     */
    order?: pulumi.Input<string>;
    /**
     * Maximum number of tasks to be rollbacked in one iteration. Defaults to `1`
     */
    parallelism?: pulumi.Input<number>;
}
export interface ServiceTaskSpec {
    /**
     * The spec for each container
     */
    containerSpec: pulumi.Input<inputs.ServiceTaskSpecContainerSpec>;
    /**
     * A counter that triggers an update even if no relevant parameters have been changed. See the [spec](https://github.com/docker/swarmkit/blob/master/api/specs.proto#L126).
     */
    forceUpdate?: pulumi.Input<number>;
    /**
     * Specifies the log driver to use for tasks created from this spec. If not present, the default one for the swarm will be used, finally falling back to the engine default if not specified
     */
    logDriver?: pulumi.Input<inputs.ServiceTaskSpecLogDriver>;
    /**
     * The networks the container is attached to
     */
    networksAdvanceds?: pulumi.Input<pulumi.Input<inputs.ServiceTaskSpecNetworksAdvanced>[]>;
    /**
     * The placement preferences
     */
    placement?: pulumi.Input<inputs.ServiceTaskSpecPlacement>;
    /**
     * Resource requirements which apply to each individual container created as part of the service
     */
    resources?: pulumi.Input<inputs.ServiceTaskSpecResources>;
    /**
     * Specification for the restart policy which applies to containers created as part of this service.
     */
    restartPolicy?: pulumi.Input<inputs.ServiceTaskSpecRestartPolicy>;
    /**
     * Runtime is the type of runtime specified for the task executor. See the [types](https://github.com/moby/moby/blob/master/api/types/swarm/runtime.go).
     */
    runtime?: pulumi.Input<string>;
}
export interface ServiceTaskSpecContainerSpec {
    /**
     * Arguments to the command
     */
    args?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The command/entrypoint to be run in the image. According to the [docker cli](https://github.com/docker/cli/blob/v20.10.7/cli/command/service/opts.go#L705) the override of the entrypoint is also passed to the `command` property and there is no `entrypoint` attribute in the `ContainerSpec` of the service.
     */
    commands?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * References to zero or more configs that will be exposed to the service
     */
    configs?: pulumi.Input<pulumi.Input<inputs.ServiceTaskSpecContainerSpecConfig>[]>;
    /**
     * The working directory for commands to run in
     */
    dir?: pulumi.Input<string>;
    /**
     * Specification for DNS related configurations in resolver configuration file (`resolv.conf`)
     */
    dnsConfig?: pulumi.Input<inputs.ServiceTaskSpecContainerSpecDnsConfig>;
    /**
     * A list of environment variables in the form VAR="value"
     */
    env?: pulumi.Input<{
        [key: string]: pulumi.Input<string>;
    }>;
    /**
     * A list of additional groups that the container process will run as
     */
    groups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A test to perform to check that the container is healthy
     */
    healthcheck?: pulumi.Input<inputs.ServiceTaskSpecContainerSpecHealthcheck>;
    /**
     * The hostname to use for the container, as a valid RFC 1123 hostname
     */
    hostname?: pulumi.Input<string>;
    /**
     * A list of hostname/IP mappings to add to the container's hosts file
     */
    hosts?: pulumi.Input<pulumi.Input<inputs.ServiceTaskSpecContainerSpecHost>[]>;
    /**
     * The image name to use for the containers of the service, like `nginx:1.17.6`. Also use the data-source or resource of `docker.RemoteImage` with the `repoDigest` or `docker.RegistryImage` with the `name` attribute for this, as shown in the examples.
     */
    image: pulumi.Input<string>;
    /**
     * Isolation technology of the containers running the service. (Windows only). Defaults to `default`.
     */
    isolation?: pulumi.Input<string>;
    /**
     * User-defined key/value metadata
     */
    labels?: pulumi.Input<pulumi.Input<inputs.ServiceTaskSpecContainerSpecLabel>[]>;
    /**
     * Specification for mounts to be added to containers created as part of the service
     */
    mounts?: pulumi.Input<pulumi.Input<inputs.ServiceTaskSpecContainerSpecMount>[]>;
    /**
     * Security options for the container
     */
    privileges?: pulumi.Input<inputs.ServiceTaskSpecContainerSpecPrivileges>;
    /**
     * Whether the mount should be read-only
     */
    readOnly?: pulumi.Input<boolean>;
    /**
     * References to zero or more secrets that will be exposed to the service
     */
    secrets?: pulumi.Input<pulumi.Input<inputs.ServiceTaskSpecContainerSpecSecret>[]>;
    /**
     * Amount of time to wait for the container to terminate before forcefully removing it (ms|s|m|h). If not specified or '0s' the destroy will not check if all tasks/containers of the service terminate.
     */
    stopGracePeriod?: pulumi.Input<string>;
    /**
     * Signal to stop the container
     */
    stopSignal?: pulumi.Input<string>;
    /**
     * Sysctls config (Linux only)
     */
    sysctl?: pulumi.Input<{
        [key: string]: any;
    }>;
    /**
     * SELinux user label
     */
    user?: pulumi.Input<string>;
}
export interface ServiceTaskSpecContainerSpecConfig {
    /**
     * ID of the specific config that we're referencing
     */
    configId: pulumi.Input<string>;
    /**
     * Name of the config that this references, but this is just provided for lookup/display purposes. The config in the reference will be identified by its ID
     */
    configName?: pulumi.Input<string>;
    /**
     * Represents the file GID. Defaults to `0`.
     */
    fileGid?: pulumi.Input<string>;
    /**
     * Represents represents the FileMode of the file. Defaults to `0o444`.
     */
    fileMode?: pulumi.Input<number>;
    /**
     * Represents the final filename in the filesystem
     */
    fileName: pulumi.Input<string>;
    /**
     * Represents the file UID. Defaults to `0`.
     */
    fileUid?: pulumi.Input<string>;
}
export interface ServiceTaskSpecContainerSpecDnsConfig {
    /**
     * The IP addresses of the name servers
     */
    nameservers: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of internal resolver variables to be modified (e.g., `debug`, `ndots:3`, etc.)
     */
    options?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A search list for host-name lookup
     */
    searches?: pulumi.Input<pulumi.Input<string>[]>;
}
export interface ServiceTaskSpecContainerSpecHealthcheck {
    /**
     * Time between running the check (ms|s|m|h). Defaults to `0s`.
     */
    interval?: pulumi.Input<string>;
    /**
     * Consecutive failures needed to report unhealthy. Defaults to `0`
     */
    retries?: pulumi.Input<number>;
    /**
     * Start period for the container to initialize before counting retries towards unstable (ms|s|m|h). Defaults to `0s`.
     */
    startPeriod?: pulumi.Input<string>;
    /**
     * The test to perform as list
     */
    tests: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The timeout of the service to reach the desired state `(s|m)`. Defaults to `3m`
     */
    timeout?: pulumi.Input<string>;
}
export interface ServiceTaskSpecContainerSpecHost {
    /**
     * The name of the host
     */
    host: pulumi.Input<string>;
    /**
     * The ip of the host
     */
    ip: pulumi.Input<string>;
}
export interface ServiceTaskSpecContainerSpecLabel {
    /**
     * Name of the label
     */
    label: pulumi.Input<string>;
    /**
     * Value of the label
     */
    value: pulumi.Input<string>;
}
export interface ServiceTaskSpecContainerSpecMount {
    /**
     * Optional configuration for the bind type
     */
    bindOptions?: pulumi.Input<inputs.ServiceTaskSpecContainerSpecMountBindOptions>;
    /**
     * Whether the mount should be read-only
     */
    readOnly?: pulumi.Input<boolean>;
    /**
     * Mount source (e.g. a volume name, a host path)
     */
    source?: pulumi.Input<string>;
    /**
     * Container path
     */
    target: pulumi.Input<string>;
    /**
     * Optional configuration for the tmpfs type
     */
    tmpfsOptions?: pulumi.Input<inputs.ServiceTaskSpecContainerSpecMountTmpfsOptions>;
    /**
     * The mount type
     */
    type: pulumi.Input<string>;
    /**
     * Optional configuration for the volume type
     */
    volumeOptions?: pulumi.Input<inputs.ServiceTaskSpecContainerSpecMountVolumeOptions>;
}
export interface ServiceTaskSpecContainerSpecMountBindOptions {
    /**
     * Bind propagation refers to whether or not mounts created within a given bind-mount or named volume can be propagated to replicas of that mount. See the [docs](https://docs.docker.com/storage/bind-mounts/#configure-bind-propagation) for details. Defaults to `rprivate`
     */
    propagation?: pulumi.Input<string>;
}
export interface ServiceTaskSpecContainerSpecMountTmpfsOptions {
    /**
     * The mode of resolution to use for internal load balancing between tasks
     */
    mode?: pulumi.Input<number>;
    /**
     * The size for the tmpfs mount in bytes
     */
    sizeBytes?: pulumi.Input<number>;
}
export interface ServiceTaskSpecContainerSpecMountVolumeOptions {
    /**
     * Name of the driver to use to create the volume
     */
    driverName?: pulumi.Input<string>;
    /**
     * key/value map of driver specific options
     */
    driverOptions?: pulumi.Input<{
        [key: string]: pulumi.Input<string>;
    }>;
    /**
     * User-defined key/value metadata
     */
    labels?: pulumi.Input<pulumi.Input<inputs.ServiceTaskSpecContainerSpecMountVolumeOptionsLabel>[]>;
    /**
     * Populate volume with data from the target
     */
    noCopy?: pulumi.Input<boolean>;
}
export interface ServiceTaskSpecContainerSpecMountVolumeOptionsLabel {
    /**
     * Name of the label
     */
    label: pulumi.Input<string>;
    /**
     * Value of the label
     */
    value: pulumi.Input<string>;
}
export interface ServiceTaskSpecContainerSpecPrivileges {
    /**
     * CredentialSpec for managed service account (Windows only)
     */
    credentialSpec?: pulumi.Input<inputs.ServiceTaskSpecContainerSpecPrivilegesCredentialSpec>;
    /**
     * SELinux labels of the container
     */
    seLinuxContext?: pulumi.Input<inputs.ServiceTaskSpecContainerSpecPrivilegesSeLinuxContext>;
}
export interface ServiceTaskSpecContainerSpecPrivilegesCredentialSpec {
    /**
     * Load credential spec from this file
     */
    file?: pulumi.Input<string>;
    /**
     * Load credential spec from this value in the Windows registry
     */
    registry?: pulumi.Input<string>;
}
export interface ServiceTaskSpecContainerSpecPrivilegesSeLinuxContext {
    /**
     * Disable SELinux
     */
    disable?: pulumi.Input<boolean>;
    /**
     * SELinux level label
     */
    level?: pulumi.Input<string>;
    /**
     * SELinux role label
     */
    role?: pulumi.Input<string>;
    /**
     * The mount type
     */
    type?: pulumi.Input<string>;
    /**
     * SELinux user label
     */
    user?: pulumi.Input<string>;
}
export interface ServiceTaskSpecContainerSpecSecret {
    /**
     * Represents the file GID. Defaults to `0`.
     */
    fileGid?: pulumi.Input<string>;
    /**
     * Represents represents the FileMode of the file. Defaults to `0o444`.
     */
    fileMode?: pulumi.Input<number>;
    /**
     * Represents the final filename in the filesystem
     */
    fileName: pulumi.Input<string>;
    /**
     * Represents the file UID. Defaults to `0`.
     */
    fileUid?: pulumi.Input<string>;
    /**
     * ID of the specific secret that we're referencing
     */
    secretId: pulumi.Input<string>;
    /**
     * Name of the secret that this references, but this is just provided for lookup/display purposes. The config in the reference will be identified by its ID
     */
    secretName?: pulumi.Input<string>;
}
export interface ServiceTaskSpecLogDriver {
    /**
     * A random name for the port
     */
    name: pulumi.Input<string>;
    /**
     * A list of internal resolver variables to be modified (e.g., `debug`, `ndots:3`, etc.)
     */
    options?: pulumi.Input<{
        [key: string]: pulumi.Input<string>;
    }>;
}
export interface ServiceTaskSpecNetworksAdvanced {
    /**
     * The network aliases of the container in the specific network.
     */
    aliases?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An array of driver options for the network, e.g. `opts1=value`
     */
    driverOpts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A random name for the port
     */
    name: pulumi.Input<string>;
}
export interface ServiceTaskSpecPlacement {
    /**
     * An array of constraints. e.g.: `node.role==manager`
     */
    constraints?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Maximum number of replicas for per node (default value is `0`, which is unlimited)
     */
    maxReplicas?: pulumi.Input<number>;
    /**
     * Platforms stores all the platforms that the service's image can run on
     */
    platforms?: pulumi.Input<pulumi.Input<inputs.ServiceTaskSpecPlacementPlatform>[]>;
    /**
     * Preferences provide a way to make the scheduler aware of factors such as topology. They are provided in order from highest to lowest precedence, e.g.: `spread=node.role.manager`
     */
    prefs?: pulumi.Input<pulumi.Input<string>[]>;
}
export interface ServiceTaskSpecPlacementPlatform {
    /**
     * The architecture, e.g. `amd64`
     */
    architecture: pulumi.Input<string>;
    /**
     * The operation system, e.g. `linux`
     */
    os: pulumi.Input<string>;
}
export interface ServiceTaskSpecResources {
    /**
     * Describes the resources which can be advertised by a node and requested by a task
     */
    limits?: pulumi.Input<inputs.ServiceTaskSpecResourcesLimits>;
    /**
     * An object describing the resources which can be advertised by a node and requested by a task
     */
    reservation?: pulumi.Input<inputs.ServiceTaskSpecResourcesReservation>;
}
export interface ServiceTaskSpecResourcesLimits {
    /**
     * The amounf of memory in bytes the container allocates
     */
    memoryBytes?: pulumi.Input<number>;
    /**
     * CPU shares in units of `1/1e9` (or `10^-9`) of the CPU. Should be at least `1000000`
     */
    nanoCpus?: pulumi.Input<number>;
}
export interface ServiceTaskSpecResourcesReservation {
    /**
     * User-defined resources can be either Integer resources (e.g, `SSD=3`) or String resources (e.g, GPU=UUID1)
     */
    genericResources?: pulumi.Input<inputs.ServiceTaskSpecResourcesReservationGenericResources>;
    /**
     * The amounf of memory in bytes the container allocates
     */
    memoryBytes?: pulumi.Input<number>;
    /**
     * CPU shares in units of `1/1e9` (or `10^-9`) of the CPU. Should be at least `1000000`
     */
    nanoCpus?: pulumi.Input<number>;
}
export interface ServiceTaskSpecResourcesReservationGenericResources {
    /**
     * The Integer resources
     */
    discreteResourcesSpecs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The String resources
     */
    namedResourcesSpecs?: pulumi.Input<pulumi.Input<string>[]>;
}
export interface ServiceTaskSpecRestartPolicy {
    /**
     * Condition for restart
     */
    condition?: pulumi.Input<string>;
    /**
     * The interval to check if the desired state is reached `(ms|s)`. Defaults to `7s`.
     */
    delay?: pulumi.Input<string>;
    /**
     * Maximum attempts to restart a given container before giving up (default value is `0`, which is ignored)
     */
    maxAttempts?: pulumi.Input<number>;
    /**
     * The time window used to evaluate the restart policy (default value is `0`, which is unbounded) (ms|s|m|h)
     */
    window?: pulumi.Input<string>;
}
export interface ServiceUpdateConfig {
    /**
     * Delay between task updates `(ns|us|ms|s|m|h)`. Defaults to `0s`.
     */
    delay?: pulumi.Input<string>;
    /**
     * Action on update failure: `pause`, `continue` or `rollback`. Defaults to `pause`.
     */
    failureAction?: pulumi.Input<string>;
    /**
     * Failure rate to tolerate during an update. Defaults to `0.0`.
     */
    maxFailureRatio?: pulumi.Input<string>;
    /**
     * Duration after each task update to monitor for failure (ns|us|ms|s|m|h). Defaults to `5s`.
     */
    monitor?: pulumi.Input<string>;
    /**
     * Update order: either 'stop-first' or 'start-first'. Defaults to `stop-first`.
     */
    order?: pulumi.Input<string>;
    /**
     * Maximum number of tasks to be updated in one iteration. Defaults to `1`
     */
    parallelism?: pulumi.Input<number>;
}
export interface VolumeLabel {
    /**
     * Name of the label
     */
    label: pulumi.Input<string>;
    /**
     * Value of the label
     */
    value: pulumi.Input<string>;
}
export declare namespace config {
}
