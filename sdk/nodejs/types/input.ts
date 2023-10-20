// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
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
    propagation?: pulumi.Input<string>;
}

export interface ContainerMountTmpfsOptions {
    mode?: pulumi.Input<number>;
    sizeBytes?: pulumi.Input<number>;
}

export interface ContainerMountVolumeOptions {
    driverName?: pulumi.Input<string>;
    driverOptions?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * User-defined key/value metadata
     */
    labels?: pulumi.Input<pulumi.Input<inputs.ContainerMountVolumeOptionsLabel>[]>;
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
     * An optional map of named build-time argument variables to set during the Docker build. This flag allows you to pass build-time variablesthat can be accessed like environment variables inside the RUN instruction.
     */
    args?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
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
    auxAddress?: pulumi.Input<{[key: string]: any}>;
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
    buildArg?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Pairs for build-time variables in the form TODO
     */
    buildArgs?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
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
    label?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * User-defined key/value metadata
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
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
    auth?: pulumi.Input<string>;
    email?: pulumi.Input<string>;
    hostName: pulumi.Input<string>;
    identityToken?: pulumi.Input<string>;
    password?: pulumi.Input<string>;
    registryToken?: pulumi.Input<string>;
    serverAddress?: pulumi.Input<string>;
    userName?: pulumi.Input<string>;
}

export interface RemoteImageBuildUlimit {
    hard: pulumi.Input<number>;
    /**
     * The name of the Docker image, including any tags or SHA256 repo digests.
     */
    name: pulumi.Input<string>;
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
     * Name of the service
     */
    name?: pulumi.Input<string>;
    protocol?: pulumi.Input<string>;
    publishMode?: pulumi.Input<string>;
    publishedPort?: pulumi.Input<number>;
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
    args?: pulumi.Input<pulumi.Input<string>[]>;
    commands?: pulumi.Input<pulumi.Input<string>[]>;
    configs?: pulumi.Input<pulumi.Input<inputs.ServiceTaskSpecContainerSpecConfig>[]>;
    dir?: pulumi.Input<string>;
    dnsConfig?: pulumi.Input<inputs.ServiceTaskSpecContainerSpecDnsConfig>;
    env?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    groups?: pulumi.Input<pulumi.Input<string>[]>;
    healthcheck?: pulumi.Input<inputs.ServiceTaskSpecContainerSpecHealthcheck>;
    hostname?: pulumi.Input<string>;
    hosts?: pulumi.Input<pulumi.Input<inputs.ServiceTaskSpecContainerSpecHost>[]>;
    image: pulumi.Input<string>;
    isolation?: pulumi.Input<string>;
    /**
     * User-defined key/value metadata
     */
    labels?: pulumi.Input<pulumi.Input<inputs.ServiceTaskSpecContainerSpecLabel>[]>;
    mounts?: pulumi.Input<pulumi.Input<inputs.ServiceTaskSpecContainerSpecMount>[]>;
    privileges?: pulumi.Input<inputs.ServiceTaskSpecContainerSpecPrivileges>;
    readOnly?: pulumi.Input<boolean>;
    secrets?: pulumi.Input<pulumi.Input<inputs.ServiceTaskSpecContainerSpecSecret>[]>;
    stopGracePeriod?: pulumi.Input<string>;
    stopSignal?: pulumi.Input<string>;
    sysctl?: pulumi.Input<{[key: string]: any}>;
    user?: pulumi.Input<string>;
}

export interface ServiceTaskSpecContainerSpecConfig {
    configId: pulumi.Input<string>;
    configName?: pulumi.Input<string>;
    fileGid?: pulumi.Input<string>;
    fileMode?: pulumi.Input<number>;
    fileName: pulumi.Input<string>;
    fileUid?: pulumi.Input<string>;
}

export interface ServiceTaskSpecContainerSpecDnsConfig {
    nameservers: pulumi.Input<pulumi.Input<string>[]>;
    options?: pulumi.Input<pulumi.Input<string>[]>;
    searches?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface ServiceTaskSpecContainerSpecHealthcheck {
    interval?: pulumi.Input<string>;
    retries?: pulumi.Input<number>;
    startPeriod?: pulumi.Input<string>;
    tests: pulumi.Input<pulumi.Input<string>[]>;
    timeout?: pulumi.Input<string>;
}

export interface ServiceTaskSpecContainerSpecHost {
    host: pulumi.Input<string>;
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
    bindOptions?: pulumi.Input<inputs.ServiceTaskSpecContainerSpecMountBindOptions>;
    readOnly?: pulumi.Input<boolean>;
    source?: pulumi.Input<string>;
    target: pulumi.Input<string>;
    tmpfsOptions?: pulumi.Input<inputs.ServiceTaskSpecContainerSpecMountTmpfsOptions>;
    type: pulumi.Input<string>;
    volumeOptions?: pulumi.Input<inputs.ServiceTaskSpecContainerSpecMountVolumeOptions>;
}

export interface ServiceTaskSpecContainerSpecMountBindOptions {
    propagation?: pulumi.Input<string>;
}

export interface ServiceTaskSpecContainerSpecMountTmpfsOptions {
    /**
     * Scheduling mode for the service
     */
    mode?: pulumi.Input<number>;
    sizeBytes?: pulumi.Input<number>;
}

export interface ServiceTaskSpecContainerSpecMountVolumeOptions {
    driverName?: pulumi.Input<string>;
    driverOptions?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * User-defined key/value metadata
     */
    labels?: pulumi.Input<pulumi.Input<inputs.ServiceTaskSpecContainerSpecMountVolumeOptionsLabel>[]>;
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
    credentialSpec?: pulumi.Input<inputs.ServiceTaskSpecContainerSpecPrivilegesCredentialSpec>;
    seLinuxContext?: pulumi.Input<inputs.ServiceTaskSpecContainerSpecPrivilegesSeLinuxContext>;
}

export interface ServiceTaskSpecContainerSpecPrivilegesCredentialSpec {
    file?: pulumi.Input<string>;
    registry?: pulumi.Input<string>;
}

export interface ServiceTaskSpecContainerSpecPrivilegesSeLinuxContext {
    disable?: pulumi.Input<boolean>;
    level?: pulumi.Input<string>;
    role?: pulumi.Input<string>;
    type?: pulumi.Input<string>;
    user?: pulumi.Input<string>;
}

export interface ServiceTaskSpecContainerSpecSecret {
    fileGid?: pulumi.Input<string>;
    fileMode?: pulumi.Input<number>;
    fileName: pulumi.Input<string>;
    fileUid?: pulumi.Input<string>;
    secretId: pulumi.Input<string>;
    secretName?: pulumi.Input<string>;
}

export interface ServiceTaskSpecLogDriver {
    /**
     * Name of the service
     */
    name: pulumi.Input<string>;
    options?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

export interface ServiceTaskSpecNetworksAdvanced {
    aliases?: pulumi.Input<pulumi.Input<string>[]>;
    driverOpts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the service
     */
    name: pulumi.Input<string>;
}

export interface ServiceTaskSpecPlacement {
    constraints?: pulumi.Input<pulumi.Input<string>[]>;
    maxReplicas?: pulumi.Input<number>;
    platforms?: pulumi.Input<pulumi.Input<inputs.ServiceTaskSpecPlacementPlatform>[]>;
    prefs?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface ServiceTaskSpecPlacementPlatform {
    architecture: pulumi.Input<string>;
    os: pulumi.Input<string>;
}

export interface ServiceTaskSpecResources {
    limits?: pulumi.Input<inputs.ServiceTaskSpecResourcesLimits>;
    reservation?: pulumi.Input<inputs.ServiceTaskSpecResourcesReservation>;
}

export interface ServiceTaskSpecResourcesLimits {
    memoryBytes?: pulumi.Input<number>;
    nanoCpus?: pulumi.Input<number>;
}

export interface ServiceTaskSpecResourcesReservation {
    genericResources?: pulumi.Input<inputs.ServiceTaskSpecResourcesReservationGenericResources>;
    memoryBytes?: pulumi.Input<number>;
    nanoCpus?: pulumi.Input<number>;
}

export interface ServiceTaskSpecResourcesReservationGenericResources {
    discreteResourcesSpecs?: pulumi.Input<pulumi.Input<string>[]>;
    namedResourcesSpecs?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface ServiceTaskSpecRestartPolicy {
    condition?: pulumi.Input<string>;
    delay?: pulumi.Input<string>;
    maxAttempts?: pulumi.Input<number>;
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
export namespace config {
}
