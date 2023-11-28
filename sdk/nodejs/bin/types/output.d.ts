import * as outputs from "../types/output";
export interface ContainerCapabilities {
    /**
     * List of linux capabilities to add.
     */
    adds?: string[];
    /**
     * List of linux capabilities to drop.
     */
    drops?: string[];
}
export interface ContainerDevice {
    /**
     * The path in the container where the device will be bound.
     */
    containerPath?: string;
    /**
     * The path on the host where the device is located.
     */
    hostPath: string;
    /**
     * The cgroup permissions given to the container to access the device. Defaults to `rwm`.
     */
    permissions?: string;
}
export interface ContainerHealthcheck {
    /**
     * Time between running the check (ms|s|m|h). Defaults to `0s`.
     */
    interval?: string;
    /**
     * Consecutive failures needed to report unhealthy. Defaults to `0`.
     */
    retries?: number;
    /**
     * Start period for the container to initialize before counting retries towards unstable (ms|s|m|h). Defaults to `0s`.
     */
    startPeriod?: string;
    /**
     * Command to run to check health. For example, to run `curl -f localhost/health` set the command to be `["CMD", "curl", "-f", "localhost/health"]`.
     */
    tests: string[];
    /**
     * Maximum time to allow one check to run (ms|s|m|h). Defaults to `0s`.
     */
    timeout?: string;
}
export interface ContainerHost {
    /**
     * Hostname to add
     */
    host: string;
    /**
     * IP address this hostname should resolve to.
     */
    ip: string;
}
export interface ContainerLabel {
    /**
     * Name of the label
     */
    label: string;
    /**
     * Value of the label
     */
    value: string;
}
export interface ContainerMount {
    /**
     * Optional configuration for the bind type.
     */
    bindOptions?: outputs.ContainerMountBindOptions;
    /**
     * Whether the mount should be read-only.
     */
    readOnly?: boolean;
    /**
     * Mount source (e.g. a volume name, a host path).
     */
    source?: string;
    /**
     * Container path
     */
    target: string;
    /**
     * Optional configuration for the tmpfs type.
     */
    tmpfsOptions?: outputs.ContainerMountTmpfsOptions;
    /**
     * The mount type
     */
    type: string;
    /**
     * Optional configuration for the volume type.
     */
    volumeOptions?: outputs.ContainerMountVolumeOptions;
}
export interface ContainerMountBindOptions {
    /**
     * A propagation mode with the value.
     */
    propagation?: string;
}
export interface ContainerMountTmpfsOptions {
    /**
     * The permission mode for the tmpfs mount in an integer.
     */
    mode?: number;
    /**
     * The size for the tmpfs mount in bytes.
     */
    sizeBytes?: number;
}
export interface ContainerMountVolumeOptions {
    /**
     * Name of the driver to use to create the volume.
     */
    driverName?: string;
    /**
     * key/value map of driver specific options.
     */
    driverOptions?: {
        [key: string]: string;
    };
    /**
     * User-defined key/value metadata.
     */
    labels?: outputs.ContainerMountVolumeOptionsLabel[];
    /**
     * Populate volume with data from the target.
     */
    noCopy?: boolean;
}
export interface ContainerMountVolumeOptionsLabel {
    /**
     * Name of the label
     */
    label: string;
    /**
     * Value of the label
     */
    value: string;
}
export interface ContainerNetworkData {
    gateway: string;
    globalIpv6Address: string;
    globalIpv6PrefixLength: number;
    ipAddress: string;
    ipPrefixLength: number;
    ipv6Gateway: string;
    macAddress: string;
    networkName: string;
}
export interface ContainerNetworksAdvanced {
    /**
     * The network aliases of the container in the specific network.
     */
    aliases?: string[];
    /**
     * The IPV4 address of the container in the specific network.
     */
    ipv4Address?: string;
    /**
     * The IPV6 address of the container in the specific network.
     */
    ipv6Address?: string;
    /**
     * The name or id of the network to use. You can use `name` or `id` attribute from a `docker.Network` resource.
     */
    name: string;
}
export interface ContainerPort {
    /**
     * Port exposed out of the container. If not given a free random port `>= 32768` will be used.
     */
    external: number;
    /**
     * Port within the container.
     */
    internal: number;
    /**
     * IP address/mask that can access this port. Defaults to `0.0.0.0`.
     */
    ip?: string;
    /**
     * Protocol that can be used over this port. Defaults to `tcp`.
     */
    protocol?: string;
}
export interface ContainerUlimit {
    /**
     * The hard limit
     */
    hard: number;
    /**
     * The name of the ulimit
     */
    name: string;
    /**
     * The soft limit
     */
    soft: number;
}
export interface ContainerUpload {
    /**
     * Literal string value to use as the object content, which will be uploaded as UTF-8-encoded text. Conflicts with `contentBase64` & `source`
     */
    content?: string;
    contentBase64?: string;
    /**
     * If `true`, the file will be uploaded with user executable permission. Defaults to `false`.
     */
    executable?: boolean;
    /**
     * Path to the file in the container where is upload goes to
     */
    file: string;
    /**
     * A filename that references a file which will be uploaded as the object content. This allows for large file uploads that do not get stored in state. Conflicts with `content` & `contentBase64`
     */
    source?: string;
    /**
     * If using `source`, this will force an update if the file content has updated but the filename has not.
     */
    sourceHash?: string;
}
export interface ContainerVolume {
    /**
     * The path in the container where the volume will be mounted.
     */
    containerPath?: string;
    /**
     * The container where the volume is coming from.
     */
    fromContainer?: string;
    /**
     * The path on the host where the volume is coming from.
     */
    hostPath?: string;
    /**
     * If `true`, this volume will be readonly. Defaults to `false`.
     */
    readOnly?: boolean;
    /**
     * The name of the docker volume which should be mounted.
     */
    volumeName?: string;
}
export interface GetNetworkIpamConfig {
    auxAddress?: {
        [key: string]: any;
    };
    gateway?: string;
    ipRange?: string;
    subnet?: string;
}
export interface NetworkIpamConfig {
    /**
     * Auxiliary IPv4 or IPv6 addresses used by Network driver
     */
    auxAddress?: {
        [key: string]: any;
    };
    /**
     * The IP address of the gateway
     */
    gateway?: string;
    /**
     * The ip range in CIDR form
     */
    ipRange?: string;
    /**
     * The subnet in CIDR form
     */
    subnet?: string;
}
export interface NetworkLabel {
    /**
     * Name of the label
     */
    label: string;
    /**
     * Value of the label
     */
    value: string;
}
export interface PluginGrantPermission {
    /**
     * The name of the permission
     */
    name: string;
    /**
     * The value of the permission
     */
    values: string[];
}
export interface RemoteImageBuild {
    /**
     * The configuration for the authentication
     */
    authConfigs?: outputs.RemoteImageBuildAuthConfig[];
    /**
     * Set build-time variables
     */
    buildArg?: {
        [key: string]: string;
    };
    /**
     * Pairs for build-time variables in the form TODO
     */
    buildArgs?: {
        [key: string]: string;
    };
    /**
     * BuildID is an optional identifier that can be passed together with the build request. The same identifier can be used to gracefully cancel the build with the cancel request.
     */
    buildId?: string;
    /**
     * Images to consider as cache sources
     */
    cacheFroms?: string[];
    /**
     * Optional parent cgroup for the container
     */
    cgroupParent?: string;
    /**
     * Value to specify the build context. Currently, only a `PATH` context is supported. You can use the helper function '${path.cwd}/context-dir'. Please see https://docs.docker.com/build/building/context/ for more information about build contexts.
     */
    context: string;
    /**
     * The length of a CPU period in microseconds
     */
    cpuPeriod?: number;
    /**
     * Microseconds of CPU time that the container can get in a CPU period
     */
    cpuQuota?: number;
    /**
     * CPUs in which to allow execution (e.g., `0-3`, `0`, `1`)
     */
    cpuSetCpus?: string;
    /**
     * MEMs in which to allow execution (`0-3`, `0`, `1`)
     */
    cpuSetMems?: string;
    /**
     * CPU shares (relative weight)
     */
    cpuShares?: number;
    /**
     * Name of the Dockerfile. Defaults to `Dockerfile`.
     */
    dockerfile?: string;
    /**
     * A list of hostnames/IP mappings to add to the container’s /etc/hosts file. Specified in the form ["hostname:IP"]
     */
    extraHosts?: string[];
    /**
     * Always remove intermediate containers
     */
    forceRemove?: boolean;
    /**
     * Isolation represents the isolation technology of a container. The supported values are
     */
    isolation?: string;
    /**
     * Set metadata for an image
     */
    label?: {
        [key: string]: string;
    };
    /**
     * User-defined key/value metadata
     */
    labels?: {
        [key: string]: string;
    };
    /**
     * Set memory limit for build
     */
    memory?: number;
    /**
     * Total memory (memory + swap), -1 to enable unlimited swap
     */
    memorySwap?: number;
    /**
     * Set the networking mode for the RUN instructions during build
     */
    networkMode?: string;
    /**
     * Do not use the cache when building the image
     */
    noCache?: boolean;
    /**
     * Set platform if server is multi-platform capable
     */
    platform?: string;
    /**
     * Attempt to pull the image even if an older image exists locally
     */
    pullParent?: boolean;
    /**
     * A Git repository URI or HTTP/HTTPS context URI
     */
    remoteContext?: string;
    /**
     * Remove intermediate containers after a successful build. Defaults to `true`.
     */
    remove?: boolean;
    /**
     * The security options
     */
    securityOpts?: string[];
    /**
     * Set an ID for the build session
     */
    sessionId?: string;
    /**
     * Size of /dev/shm in bytes. The size must be greater than 0
     */
    shmSize?: number;
    /**
     * If true the new layers are squashed into a new image with a single new layer
     */
    squash?: boolean;
    /**
     * Suppress the build output and print image ID on success
     */
    suppressOutput?: boolean;
    /**
     * Name and optionally a tag in the 'name:tag' format
     */
    tags?: string[];
    /**
     * Set the target build stage to build
     */
    target?: string;
    /**
     * Configuration for ulimits
     */
    ulimits?: outputs.RemoteImageBuildUlimit[];
    /**
     * Version of the underlying builder to use
     */
    version?: string;
}
export interface RemoteImageBuildAuthConfig {
    /**
     * the auth token
     */
    auth?: string;
    /**
     * the user emal
     */
    email?: string;
    /**
     * hostname of the registry
     */
    hostName: string;
    /**
     * the identity token
     */
    identityToken?: string;
    /**
     * the registry password
     */
    password?: string;
    /**
     * the registry token
     */
    registryToken?: string;
    /**
     * the server address
     */
    serverAddress?: string;
    /**
     * the registry user name
     */
    userName?: string;
}
export interface RemoteImageBuildUlimit {
    /**
     * soft limit
     */
    hard: number;
    /**
     * type of ulimit, e.g. `nofile`
     */
    name: string;
    /**
     * hard limit
     */
    soft: number;
}
export interface SecretLabel {
    /**
     * Name of the label
     */
    label: string;
    /**
     * Value of the label
     */
    value: string;
}
export interface ServiceAuth {
    /**
     * The password
     */
    password?: string;
    /**
     * The address of the server for the authentication
     */
    serverAddress: string;
    /**
     * The username
     */
    username?: string;
}
export interface ServiceConvergeConfig {
    /**
     * The interval to check if the desired state is reached `(ms|s)`. Defaults to `7s`.
     */
    delay?: string;
    /**
     * The timeout of the service to reach the desired state `(s|m)`. Defaults to `3m`
     */
    timeout?: string;
}
export interface ServiceEndpointSpec {
    /**
     * The mode of resolution to use for internal load balancing between tasks
     */
    mode: string;
    /**
     * List of exposed ports that this service is accessible on from the outside. Ports can only be provided if 'vip' resolution mode is used
     */
    ports?: outputs.ServiceEndpointSpecPort[];
}
export interface ServiceEndpointSpecPort {
    /**
     * A random name for the port
     */
    name?: string;
    /**
     * Rrepresents the protocol of a port: `tcp`, `udp` or `sctp`. Defaults to `tcp`.
     */
    protocol?: string;
    /**
     * Represents the mode in which the port is to be published: 'ingress' or 'host'. Defaults to `ingress`.
     */
    publishMode?: string;
    /**
     * The port on the swarm hosts
     */
    publishedPort: number;
    /**
     * The port inside the container
     */
    targetPort: number;
}
export interface ServiceLabel {
    /**
     * Name of the label
     */
    label: string;
    /**
     * Value of the label
     */
    value: string;
}
export interface ServiceMode {
    /**
     * The global service mode. Defaults to `false`
     */
    global?: boolean;
    /**
     * The replicated service mode
     */
    replicated: outputs.ServiceModeReplicated;
}
export interface ServiceModeReplicated {
    /**
     * The amount of replicas of the service. Defaults to `1`
     */
    replicas?: number;
}
export interface ServiceRollbackConfig {
    /**
     * Delay between task rollbacks (ns|us|ms|s|m|h). Defaults to `0s`.
     */
    delay?: string;
    /**
     * Action on rollback failure: pause | continue. Defaults to `pause`.
     */
    failureAction?: string;
    /**
     * Failure rate to tolerate during a rollback. Defaults to `0.0`.
     */
    maxFailureRatio?: string;
    /**
     * Duration after each task rollback to monitor for failure (ns|us|ms|s|m|h). Defaults to `5s`.
     */
    monitor?: string;
    /**
     * Rollback order: either 'stop-first' or 'start-first'. Defaults to `stop-first`.
     */
    order?: string;
    /**
     * Maximum number of tasks to be rollbacked in one iteration. Defaults to `1`
     */
    parallelism?: number;
}
export interface ServiceTaskSpec {
    /**
     * The spec for each container
     */
    containerSpec: outputs.ServiceTaskSpecContainerSpec;
    /**
     * A counter that triggers an update even if no relevant parameters have been changed. See the [spec](https://github.com/docker/swarmkit/blob/master/api/specs.proto#L126).
     */
    forceUpdate: number;
    /**
     * Specifies the log driver to use for tasks created from this spec. If not present, the default one for the swarm will be used, finally falling back to the engine default if not specified
     */
    logDriver?: outputs.ServiceTaskSpecLogDriver;
    /**
     * The networks the container is attached to
     */
    networksAdvanceds?: outputs.ServiceTaskSpecNetworksAdvanced[];
    /**
     * The placement preferences
     */
    placement: outputs.ServiceTaskSpecPlacement;
    /**
     * Resource requirements which apply to each individual container created as part of the service
     */
    resources: outputs.ServiceTaskSpecResources;
    /**
     * Specification for the restart policy which applies to containers created as part of this service.
     */
    restartPolicy: outputs.ServiceTaskSpecRestartPolicy;
    /**
     * Runtime is the type of runtime specified for the task executor. See the [types](https://github.com/moby/moby/blob/master/api/types/swarm/runtime.go).
     */
    runtime: string;
}
export interface ServiceTaskSpecContainerSpec {
    /**
     * Arguments to the command
     */
    args?: string[];
    /**
     * The command/entrypoint to be run in the image. According to the [docker cli](https://github.com/docker/cli/blob/v20.10.7/cli/command/service/opts.go#L705) the override of the entrypoint is also passed to the `command` property and there is no `entrypoint` attribute in the `ContainerSpec` of the service.
     */
    commands?: string[];
    /**
     * References to zero or more configs that will be exposed to the service
     */
    configs?: outputs.ServiceTaskSpecContainerSpecConfig[];
    /**
     * The working directory for commands to run in
     */
    dir?: string;
    /**
     * Specification for DNS related configurations in resolver configuration file (`resolv.conf`)
     */
    dnsConfig: outputs.ServiceTaskSpecContainerSpecDnsConfig;
    /**
     * A list of environment variables in the form VAR="value"
     */
    env?: {
        [key: string]: string;
    };
    /**
     * A list of additional groups that the container process will run as
     */
    groups?: string[];
    /**
     * A test to perform to check that the container is healthy
     */
    healthcheck: outputs.ServiceTaskSpecContainerSpecHealthcheck;
    /**
     * The hostname to use for the container, as a valid RFC 1123 hostname
     */
    hostname?: string;
    /**
     * A list of hostname/IP mappings to add to the container's hosts file
     */
    hosts?: outputs.ServiceTaskSpecContainerSpecHost[];
    /**
     * The image name to use for the containers of the service, like `nginx:1.17.6`. Also use the data-source or resource of `docker.RemoteImage` with the `repoDigest` or `docker.RegistryImage` with the `name` attribute for this, as shown in the examples.
     */
    image: string;
    /**
     * Isolation technology of the containers running the service. (Windows only). Defaults to `default`.
     */
    isolation?: string;
    /**
     * User-defined key/value metadata
     */
    labels?: outputs.ServiceTaskSpecContainerSpecLabel[];
    /**
     * Specification for mounts to be added to containers created as part of the service
     */
    mounts?: outputs.ServiceTaskSpecContainerSpecMount[];
    /**
     * Security options for the container
     */
    privileges?: outputs.ServiceTaskSpecContainerSpecPrivileges;
    /**
     * Whether the mount should be read-only
     */
    readOnly?: boolean;
    /**
     * References to zero or more secrets that will be exposed to the service
     */
    secrets?: outputs.ServiceTaskSpecContainerSpecSecret[];
    /**
     * Amount of time to wait for the container to terminate before forcefully removing it (ms|s|m|h). If not specified or '0s' the destroy will not check if all tasks/containers of the service terminate.
     */
    stopGracePeriod: string;
    /**
     * Signal to stop the container
     */
    stopSignal?: string;
    /**
     * Sysctls config (Linux only)
     */
    sysctl?: {
        [key: string]: any;
    };
    /**
     * SELinux user label
     */
    user?: string;
}
export interface ServiceTaskSpecContainerSpecConfig {
    /**
     * ID of the specific config that we're referencing
     */
    configId: string;
    /**
     * Name of the config that this references, but this is just provided for lookup/display purposes. The config in the reference will be identified by its ID
     */
    configName?: string;
    /**
     * Represents the file GID. Defaults to `0`.
     */
    fileGid?: string;
    /**
     * Represents represents the FileMode of the file. Defaults to `0o444`.
     */
    fileMode?: number;
    /**
     * Represents the final filename in the filesystem
     */
    fileName: string;
    /**
     * Represents the file UID. Defaults to `0`.
     */
    fileUid?: string;
}
export interface ServiceTaskSpecContainerSpecDnsConfig {
    /**
     * The IP addresses of the name servers
     */
    nameservers: string[];
    /**
     * A list of internal resolver variables to be modified (e.g., `debug`, `ndots:3`, etc.)
     */
    options?: string[];
    /**
     * A search list for host-name lookup
     */
    searches?: string[];
}
export interface ServiceTaskSpecContainerSpecHealthcheck {
    /**
     * Time between running the check (ms|s|m|h). Defaults to `0s`.
     */
    interval?: string;
    /**
     * Consecutive failures needed to report unhealthy. Defaults to `0`
     */
    retries?: number;
    /**
     * Start period for the container to initialize before counting retries towards unstable (ms|s|m|h). Defaults to `0s`.
     */
    startPeriod?: string;
    /**
     * The test to perform as list
     */
    tests: string[];
    /**
     * The timeout of the service to reach the desired state `(s|m)`. Defaults to `3m`
     */
    timeout?: string;
}
export interface ServiceTaskSpecContainerSpecHost {
    /**
     * The name of the host
     */
    host: string;
    /**
     * The ip of the host
     */
    ip: string;
}
export interface ServiceTaskSpecContainerSpecLabel {
    /**
     * Name of the label
     */
    label: string;
    /**
     * Value of the label
     */
    value: string;
}
export interface ServiceTaskSpecContainerSpecMount {
    /**
     * Optional configuration for the bind type
     */
    bindOptions?: outputs.ServiceTaskSpecContainerSpecMountBindOptions;
    /**
     * Whether the mount should be read-only
     */
    readOnly?: boolean;
    /**
     * Mount source (e.g. a volume name, a host path)
     */
    source?: string;
    /**
     * Container path
     */
    target: string;
    /**
     * Optional configuration for the tmpfs type
     */
    tmpfsOptions?: outputs.ServiceTaskSpecContainerSpecMountTmpfsOptions;
    /**
     * The mount type
     */
    type: string;
    /**
     * Optional configuration for the volume type
     */
    volumeOptions?: outputs.ServiceTaskSpecContainerSpecMountVolumeOptions;
}
export interface ServiceTaskSpecContainerSpecMountBindOptions {
    /**
     * Bind propagation refers to whether or not mounts created within a given bind-mount or named volume can be propagated to replicas of that mount. See the [docs](https://docs.docker.com/storage/bind-mounts/#configure-bind-propagation) for details. Defaults to `rprivate`
     */
    propagation?: string;
}
export interface ServiceTaskSpecContainerSpecMountTmpfsOptions {
    /**
     * The mode of resolution to use for internal load balancing between tasks
     */
    mode?: number;
    /**
     * The size for the tmpfs mount in bytes
     */
    sizeBytes?: number;
}
export interface ServiceTaskSpecContainerSpecMountVolumeOptions {
    /**
     * Name of the driver to use to create the volume
     */
    driverName?: string;
    /**
     * key/value map of driver specific options
     */
    driverOptions?: {
        [key: string]: string;
    };
    /**
     * User-defined key/value metadata
     */
    labels?: outputs.ServiceTaskSpecContainerSpecMountVolumeOptionsLabel[];
    /**
     * Populate volume with data from the target
     */
    noCopy?: boolean;
}
export interface ServiceTaskSpecContainerSpecMountVolumeOptionsLabel {
    /**
     * Name of the label
     */
    label: string;
    /**
     * Value of the label
     */
    value: string;
}
export interface ServiceTaskSpecContainerSpecPrivileges {
    /**
     * CredentialSpec for managed service account (Windows only)
     */
    credentialSpec?: outputs.ServiceTaskSpecContainerSpecPrivilegesCredentialSpec;
    /**
     * SELinux labels of the container
     */
    seLinuxContext?: outputs.ServiceTaskSpecContainerSpecPrivilegesSeLinuxContext;
}
export interface ServiceTaskSpecContainerSpecPrivilegesCredentialSpec {
    /**
     * Load credential spec from this file
     */
    file?: string;
    /**
     * Load credential spec from this value in the Windows registry
     */
    registry?: string;
}
export interface ServiceTaskSpecContainerSpecPrivilegesSeLinuxContext {
    /**
     * Disable SELinux
     */
    disable?: boolean;
    /**
     * SELinux level label
     */
    level?: string;
    /**
     * SELinux role label
     */
    role?: string;
    /**
     * The mount type
     */
    type?: string;
    /**
     * SELinux user label
     */
    user?: string;
}
export interface ServiceTaskSpecContainerSpecSecret {
    /**
     * Represents the file GID. Defaults to `0`.
     */
    fileGid?: string;
    /**
     * Represents represents the FileMode of the file. Defaults to `0o444`.
     */
    fileMode?: number;
    /**
     * Represents the final filename in the filesystem
     */
    fileName: string;
    /**
     * Represents the file UID. Defaults to `0`.
     */
    fileUid?: string;
    /**
     * ID of the specific secret that we're referencing
     */
    secretId: string;
    /**
     * Name of the secret that this references, but this is just provided for lookup/display purposes. The config in the reference will be identified by its ID
     */
    secretName?: string;
}
export interface ServiceTaskSpecLogDriver {
    /**
     * A random name for the port
     */
    name: string;
    /**
     * A list of internal resolver variables to be modified (e.g., `debug`, `ndots:3`, etc.)
     */
    options?: {
        [key: string]: string;
    };
}
export interface ServiceTaskSpecNetworksAdvanced {
    /**
     * The network aliases of the container in the specific network.
     */
    aliases?: string[];
    /**
     * An array of driver options for the network, e.g. `opts1=value`
     */
    driverOpts?: string[];
    /**
     * A random name for the port
     */
    name: string;
}
export interface ServiceTaskSpecPlacement {
    /**
     * An array of constraints. e.g.: `node.role==manager`
     */
    constraints?: string[];
    /**
     * Maximum number of replicas for per node (default value is `0`, which is unlimited)
     */
    maxReplicas?: number;
    /**
     * Platforms stores all the platforms that the service's image can run on
     */
    platforms?: outputs.ServiceTaskSpecPlacementPlatform[];
    /**
     * Preferences provide a way to make the scheduler aware of factors such as topology. They are provided in order from highest to lowest precedence, e.g.: `spread=node.role.manager`
     */
    prefs?: string[];
}
export interface ServiceTaskSpecPlacementPlatform {
    /**
     * The architecture, e.g. `amd64`
     */
    architecture: string;
    /**
     * The operation system, e.g. `linux`
     */
    os: string;
}
export interface ServiceTaskSpecResources {
    /**
     * Describes the resources which can be advertised by a node and requested by a task
     */
    limits?: outputs.ServiceTaskSpecResourcesLimits;
    /**
     * An object describing the resources which can be advertised by a node and requested by a task
     */
    reservation?: outputs.ServiceTaskSpecResourcesReservation;
}
export interface ServiceTaskSpecResourcesLimits {
    /**
     * The amounf of memory in bytes the container allocates
     */
    memoryBytes?: number;
    /**
     * CPU shares in units of `1/1e9` (or `10^-9`) of the CPU. Should be at least `1000000`
     */
    nanoCpus?: number;
}
export interface ServiceTaskSpecResourcesReservation {
    /**
     * User-defined resources can be either Integer resources (e.g, `SSD=3`) or String resources (e.g, GPU=UUID1)
     */
    genericResources?: outputs.ServiceTaskSpecResourcesReservationGenericResources;
    /**
     * The amounf of memory in bytes the container allocates
     */
    memoryBytes?: number;
    /**
     * CPU shares in units of `1/1e9` (or `10^-9`) of the CPU. Should be at least `1000000`
     */
    nanoCpus?: number;
}
export interface ServiceTaskSpecResourcesReservationGenericResources {
    /**
     * The Integer resources
     */
    discreteResourcesSpecs?: string[];
    /**
     * The String resources
     */
    namedResourcesSpecs?: string[];
}
export interface ServiceTaskSpecRestartPolicy {
    /**
     * Condition for restart
     */
    condition?: string;
    /**
     * The interval to check if the desired state is reached `(ms|s)`. Defaults to `7s`.
     */
    delay?: string;
    /**
     * Maximum attempts to restart a given container before giving up (default value is `0`, which is ignored)
     */
    maxAttempts?: number;
    /**
     * The time window used to evaluate the restart policy (default value is `0`, which is unbounded) (ms|s|m|h)
     */
    window?: string;
}
export interface ServiceUpdateConfig {
    /**
     * Delay between task updates `(ns|us|ms|s|m|h)`. Defaults to `0s`.
     */
    delay?: string;
    /**
     * Action on update failure: `pause`, `continue` or `rollback`. Defaults to `pause`.
     */
    failureAction?: string;
    /**
     * Failure rate to tolerate during an update. Defaults to `0.0`.
     */
    maxFailureRatio?: string;
    /**
     * Duration after each task update to monitor for failure (ns|us|ms|s|m|h). Defaults to `5s`.
     */
    monitor?: string;
    /**
     * Update order: either 'stop-first' or 'start-first'. Defaults to `stop-first`.
     */
    order?: string;
    /**
     * Maximum number of tasks to be updated in one iteration. Defaults to `1`
     */
    parallelism?: number;
}
export interface VolumeLabel {
    /**
     * Name of the label
     */
    label: string;
    /**
     * Value of the label
     */
    value: string;
}
export declare namespace config {
    interface RegistryAuth {
        address: string;
        authDisabled?: boolean;
        configFile?: string;
        configFileContent?: string;
        password?: string;
        username?: string;
    }
}
