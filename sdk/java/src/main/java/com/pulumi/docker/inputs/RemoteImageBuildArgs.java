// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.docker.inputs.RemoteImageBuildAuthConfigArgs;
import com.pulumi.docker.inputs.RemoteImageBuildUlimitArgs;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RemoteImageBuildArgs extends com.pulumi.resources.ResourceArgs {

    public static final RemoteImageBuildArgs Empty = new RemoteImageBuildArgs();

    /**
     * The configuration for the authentication
     * 
     */
    @Import(name="authConfigs")
    private @Nullable Output<List<RemoteImageBuildAuthConfigArgs>> authConfigs;

    /**
     * @return The configuration for the authentication
     * 
     */
    public Optional<Output<List<RemoteImageBuildAuthConfigArgs>>> authConfigs() {
        return Optional.ofNullable(this.authConfigs);
    }

    /**
     * Set build-time variables
     * 
     */
    @Import(name="buildArg")
    private @Nullable Output<Map<String,String>> buildArg;

    /**
     * @return Set build-time variables
     * 
     */
    public Optional<Output<Map<String,String>>> buildArg() {
        return Optional.ofNullable(this.buildArg);
    }

    /**
     * Pairs for build-time variables in the form TODO
     * 
     */
    @Import(name="buildArgs")
    private @Nullable Output<Map<String,String>> buildArgs;

    /**
     * @return Pairs for build-time variables in the form TODO
     * 
     */
    public Optional<Output<Map<String,String>>> buildArgs() {
        return Optional.ofNullable(this.buildArgs);
    }

    /**
     * BuildID is an optional identifier that can be passed together with the build request. The same identifier can be used to gracefully cancel the build with the cancel request.
     * 
     */
    @Import(name="buildId")
    private @Nullable Output<String> buildId;

    /**
     * @return BuildID is an optional identifier that can be passed together with the build request. The same identifier can be used to gracefully cancel the build with the cancel request.
     * 
     */
    public Optional<Output<String>> buildId() {
        return Optional.ofNullable(this.buildId);
    }

    /**
     * Images to consider as cache sources
     * 
     */
    @Import(name="cacheFroms")
    private @Nullable Output<List<String>> cacheFroms;

    /**
     * @return Images to consider as cache sources
     * 
     */
    public Optional<Output<List<String>>> cacheFroms() {
        return Optional.ofNullable(this.cacheFroms);
    }

    /**
     * Optional parent cgroup for the container
     * 
     */
    @Import(name="cgroupParent")
    private @Nullable Output<String> cgroupParent;

    /**
     * @return Optional parent cgroup for the container
     * 
     */
    public Optional<Output<String>> cgroupParent() {
        return Optional.ofNullable(this.cgroupParent);
    }

    /**
     * Value to specify the build context. Currently, only a `PATH` context is supported. You can use the helper function &#39;${path.cwd}/context-dir&#39;. Please see https://docs.docker.com/build/building/context/ for more information about build contexts.
     * 
     */
    @Import(name="context", required=true)
    private Output<String> context;

    /**
     * @return Value to specify the build context. Currently, only a `PATH` context is supported. You can use the helper function &#39;${path.cwd}/context-dir&#39;. Please see https://docs.docker.com/build/building/context/ for more information about build contexts.
     * 
     */
    public Output<String> context() {
        return this.context;
    }

    /**
     * The length of a CPU period in microseconds
     * 
     */
    @Import(name="cpuPeriod")
    private @Nullable Output<Integer> cpuPeriod;

    /**
     * @return The length of a CPU period in microseconds
     * 
     */
    public Optional<Output<Integer>> cpuPeriod() {
        return Optional.ofNullable(this.cpuPeriod);
    }

    /**
     * Microseconds of CPU time that the container can get in a CPU period
     * 
     */
    @Import(name="cpuQuota")
    private @Nullable Output<Integer> cpuQuota;

    /**
     * @return Microseconds of CPU time that the container can get in a CPU period
     * 
     */
    public Optional<Output<Integer>> cpuQuota() {
        return Optional.ofNullable(this.cpuQuota);
    }

    /**
     * CPUs in which to allow execution (e.g., `0-3`, `0`, `1`)
     * 
     */
    @Import(name="cpuSetCpus")
    private @Nullable Output<String> cpuSetCpus;

    /**
     * @return CPUs in which to allow execution (e.g., `0-3`, `0`, `1`)
     * 
     */
    public Optional<Output<String>> cpuSetCpus() {
        return Optional.ofNullable(this.cpuSetCpus);
    }

    /**
     * MEMs in which to allow execution (`0-3`, `0`, `1`)
     * 
     */
    @Import(name="cpuSetMems")
    private @Nullable Output<String> cpuSetMems;

    /**
     * @return MEMs in which to allow execution (`0-3`, `0`, `1`)
     * 
     */
    public Optional<Output<String>> cpuSetMems() {
        return Optional.ofNullable(this.cpuSetMems);
    }

    /**
     * CPU shares (relative weight)
     * 
     */
    @Import(name="cpuShares")
    private @Nullable Output<Integer> cpuShares;

    /**
     * @return CPU shares (relative weight)
     * 
     */
    public Optional<Output<Integer>> cpuShares() {
        return Optional.ofNullable(this.cpuShares);
    }

    /**
     * Name of the Dockerfile. Defaults to `Dockerfile`.
     * 
     */
    @Import(name="dockerfile")
    private @Nullable Output<String> dockerfile;

    /**
     * @return Name of the Dockerfile. Defaults to `Dockerfile`.
     * 
     */
    public Optional<Output<String>> dockerfile() {
        return Optional.ofNullable(this.dockerfile);
    }

    /**
     * A list of hostnames/IP mappings to add to the container’s /etc/hosts file. Specified in the form [&#34;hostname:IP&#34;]
     * 
     */
    @Import(name="extraHosts")
    private @Nullable Output<List<String>> extraHosts;

    /**
     * @return A list of hostnames/IP mappings to add to the container’s /etc/hosts file. Specified in the form [&#34;hostname:IP&#34;]
     * 
     */
    public Optional<Output<List<String>>> extraHosts() {
        return Optional.ofNullable(this.extraHosts);
    }

    /**
     * Always remove intermediate containers
     * 
     */
    @Import(name="forceRemove")
    private @Nullable Output<Boolean> forceRemove;

    /**
     * @return Always remove intermediate containers
     * 
     */
    public Optional<Output<Boolean>> forceRemove() {
        return Optional.ofNullable(this.forceRemove);
    }

    /**
     * Isolation represents the isolation technology of a container. The supported values are
     * 
     */
    @Import(name="isolation")
    private @Nullable Output<String> isolation;

    /**
     * @return Isolation represents the isolation technology of a container. The supported values are
     * 
     */
    public Optional<Output<String>> isolation() {
        return Optional.ofNullable(this.isolation);
    }

    /**
     * Set metadata for an image
     * 
     */
    @Import(name="label")
    private @Nullable Output<Map<String,String>> label;

    /**
     * @return Set metadata for an image
     * 
     */
    public Optional<Output<Map<String,String>>> label() {
        return Optional.ofNullable(this.label);
    }

    /**
     * User-defined key/value metadata
     * 
     */
    @Import(name="labels")
    private @Nullable Output<Map<String,String>> labels;

    /**
     * @return User-defined key/value metadata
     * 
     */
    public Optional<Output<Map<String,String>>> labels() {
        return Optional.ofNullable(this.labels);
    }

    /**
     * Set memory limit for build
     * 
     */
    @Import(name="memory")
    private @Nullable Output<Integer> memory;

    /**
     * @return Set memory limit for build
     * 
     */
    public Optional<Output<Integer>> memory() {
        return Optional.ofNullable(this.memory);
    }

    /**
     * Total memory (memory + swap), -1 to enable unlimited swap
     * 
     */
    @Import(name="memorySwap")
    private @Nullable Output<Integer> memorySwap;

    /**
     * @return Total memory (memory + swap), -1 to enable unlimited swap
     * 
     */
    public Optional<Output<Integer>> memorySwap() {
        return Optional.ofNullable(this.memorySwap);
    }

    /**
     * Set the networking mode for the RUN instructions during build
     * 
     */
    @Import(name="networkMode")
    private @Nullable Output<String> networkMode;

    /**
     * @return Set the networking mode for the RUN instructions during build
     * 
     */
    public Optional<Output<String>> networkMode() {
        return Optional.ofNullable(this.networkMode);
    }

    /**
     * Do not use the cache when building the image
     * 
     */
    @Import(name="noCache")
    private @Nullable Output<Boolean> noCache;

    /**
     * @return Do not use the cache when building the image
     * 
     */
    public Optional<Output<Boolean>> noCache() {
        return Optional.ofNullable(this.noCache);
    }

    /**
     * Set platform if server is multi-platform capable
     * 
     */
    @Import(name="platform")
    private @Nullable Output<String> platform;

    /**
     * @return Set platform if server is multi-platform capable
     * 
     */
    public Optional<Output<String>> platform() {
        return Optional.ofNullable(this.platform);
    }

    /**
     * Attempt to pull the image even if an older image exists locally
     * 
     */
    @Import(name="pullParent")
    private @Nullable Output<Boolean> pullParent;

    /**
     * @return Attempt to pull the image even if an older image exists locally
     * 
     */
    public Optional<Output<Boolean>> pullParent() {
        return Optional.ofNullable(this.pullParent);
    }

    /**
     * A Git repository URI or HTTP/HTTPS context URI
     * 
     */
    @Import(name="remoteContext")
    private @Nullable Output<String> remoteContext;

    /**
     * @return A Git repository URI or HTTP/HTTPS context URI
     * 
     */
    public Optional<Output<String>> remoteContext() {
        return Optional.ofNullable(this.remoteContext);
    }

    /**
     * Remove intermediate containers after a successful build. Defaults to `true`.
     * 
     */
    @Import(name="remove")
    private @Nullable Output<Boolean> remove;

    /**
     * @return Remove intermediate containers after a successful build. Defaults to `true`.
     * 
     */
    public Optional<Output<Boolean>> remove() {
        return Optional.ofNullable(this.remove);
    }

    /**
     * The security options
     * 
     */
    @Import(name="securityOpts")
    private @Nullable Output<List<String>> securityOpts;

    /**
     * @return The security options
     * 
     */
    public Optional<Output<List<String>>> securityOpts() {
        return Optional.ofNullable(this.securityOpts);
    }

    /**
     * Set an ID for the build session
     * 
     */
    @Import(name="sessionId")
    private @Nullable Output<String> sessionId;

    /**
     * @return Set an ID for the build session
     * 
     */
    public Optional<Output<String>> sessionId() {
        return Optional.ofNullable(this.sessionId);
    }

    /**
     * Size of /dev/shm in bytes. The size must be greater than 0
     * 
     */
    @Import(name="shmSize")
    private @Nullable Output<Integer> shmSize;

    /**
     * @return Size of /dev/shm in bytes. The size must be greater than 0
     * 
     */
    public Optional<Output<Integer>> shmSize() {
        return Optional.ofNullable(this.shmSize);
    }

    /**
     * If true the new layers are squashed into a new image with a single new layer
     * 
     */
    @Import(name="squash")
    private @Nullable Output<Boolean> squash;

    /**
     * @return If true the new layers are squashed into a new image with a single new layer
     * 
     */
    public Optional<Output<Boolean>> squash() {
        return Optional.ofNullable(this.squash);
    }

    /**
     * Suppress the build output and print image ID on success
     * 
     */
    @Import(name="suppressOutput")
    private @Nullable Output<Boolean> suppressOutput;

    /**
     * @return Suppress the build output and print image ID on success
     * 
     */
    public Optional<Output<Boolean>> suppressOutput() {
        return Optional.ofNullable(this.suppressOutput);
    }

    /**
     * Name and optionally a tag in the &#39;name:tag&#39; format
     * 
     */
    @Import(name="tags")
    private @Nullable Output<List<String>> tags;

    /**
     * @return Name and optionally a tag in the &#39;name:tag&#39; format
     * 
     */
    public Optional<Output<List<String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * Set the target build stage to build
     * 
     */
    @Import(name="target")
    private @Nullable Output<String> target;

    /**
     * @return Set the target build stage to build
     * 
     */
    public Optional<Output<String>> target() {
        return Optional.ofNullable(this.target);
    }

    /**
     * Configuration for ulimits
     * 
     */
    @Import(name="ulimits")
    private @Nullable Output<List<RemoteImageBuildUlimitArgs>> ulimits;

    /**
     * @return Configuration for ulimits
     * 
     */
    public Optional<Output<List<RemoteImageBuildUlimitArgs>>> ulimits() {
        return Optional.ofNullable(this.ulimits);
    }

    /**
     * Version of the underlying builder to use
     * 
     */
    @Import(name="version")
    private @Nullable Output<String> version;

    /**
     * @return Version of the underlying builder to use
     * 
     */
    public Optional<Output<String>> version() {
        return Optional.ofNullable(this.version);
    }

    private RemoteImageBuildArgs() {}

    private RemoteImageBuildArgs(RemoteImageBuildArgs $) {
        this.authConfigs = $.authConfigs;
        this.buildArg = $.buildArg;
        this.buildArgs = $.buildArgs;
        this.buildId = $.buildId;
        this.cacheFroms = $.cacheFroms;
        this.cgroupParent = $.cgroupParent;
        this.context = $.context;
        this.cpuPeriod = $.cpuPeriod;
        this.cpuQuota = $.cpuQuota;
        this.cpuSetCpus = $.cpuSetCpus;
        this.cpuSetMems = $.cpuSetMems;
        this.cpuShares = $.cpuShares;
        this.dockerfile = $.dockerfile;
        this.extraHosts = $.extraHosts;
        this.forceRemove = $.forceRemove;
        this.isolation = $.isolation;
        this.label = $.label;
        this.labels = $.labels;
        this.memory = $.memory;
        this.memorySwap = $.memorySwap;
        this.networkMode = $.networkMode;
        this.noCache = $.noCache;
        this.platform = $.platform;
        this.pullParent = $.pullParent;
        this.remoteContext = $.remoteContext;
        this.remove = $.remove;
        this.securityOpts = $.securityOpts;
        this.sessionId = $.sessionId;
        this.shmSize = $.shmSize;
        this.squash = $.squash;
        this.suppressOutput = $.suppressOutput;
        this.tags = $.tags;
        this.target = $.target;
        this.ulimits = $.ulimits;
        this.version = $.version;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RemoteImageBuildArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RemoteImageBuildArgs $;

        public Builder() {
            $ = new RemoteImageBuildArgs();
        }

        public Builder(RemoteImageBuildArgs defaults) {
            $ = new RemoteImageBuildArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param authConfigs The configuration for the authentication
         * 
         * @return builder
         * 
         */
        public Builder authConfigs(@Nullable Output<List<RemoteImageBuildAuthConfigArgs>> authConfigs) {
            $.authConfigs = authConfigs;
            return this;
        }

        /**
         * @param authConfigs The configuration for the authentication
         * 
         * @return builder
         * 
         */
        public Builder authConfigs(List<RemoteImageBuildAuthConfigArgs> authConfigs) {
            return authConfigs(Output.of(authConfigs));
        }

        /**
         * @param authConfigs The configuration for the authentication
         * 
         * @return builder
         * 
         */
        public Builder authConfigs(RemoteImageBuildAuthConfigArgs... authConfigs) {
            return authConfigs(List.of(authConfigs));
        }

        /**
         * @param buildArg Set build-time variables
         * 
         * @return builder
         * 
         */
        public Builder buildArg(@Nullable Output<Map<String,String>> buildArg) {
            $.buildArg = buildArg;
            return this;
        }

        /**
         * @param buildArg Set build-time variables
         * 
         * @return builder
         * 
         */
        public Builder buildArg(Map<String,String> buildArg) {
            return buildArg(Output.of(buildArg));
        }

        /**
         * @param buildArgs Pairs for build-time variables in the form TODO
         * 
         * @return builder
         * 
         */
        public Builder buildArgs(@Nullable Output<Map<String,String>> buildArgs) {
            $.buildArgs = buildArgs;
            return this;
        }

        /**
         * @param buildArgs Pairs for build-time variables in the form TODO
         * 
         * @return builder
         * 
         */
        public Builder buildArgs(Map<String,String> buildArgs) {
            return buildArgs(Output.of(buildArgs));
        }

        /**
         * @param buildId BuildID is an optional identifier that can be passed together with the build request. The same identifier can be used to gracefully cancel the build with the cancel request.
         * 
         * @return builder
         * 
         */
        public Builder buildId(@Nullable Output<String> buildId) {
            $.buildId = buildId;
            return this;
        }

        /**
         * @param buildId BuildID is an optional identifier that can be passed together with the build request. The same identifier can be used to gracefully cancel the build with the cancel request.
         * 
         * @return builder
         * 
         */
        public Builder buildId(String buildId) {
            return buildId(Output.of(buildId));
        }

        /**
         * @param cacheFroms Images to consider as cache sources
         * 
         * @return builder
         * 
         */
        public Builder cacheFroms(@Nullable Output<List<String>> cacheFroms) {
            $.cacheFroms = cacheFroms;
            return this;
        }

        /**
         * @param cacheFroms Images to consider as cache sources
         * 
         * @return builder
         * 
         */
        public Builder cacheFroms(List<String> cacheFroms) {
            return cacheFroms(Output.of(cacheFroms));
        }

        /**
         * @param cacheFroms Images to consider as cache sources
         * 
         * @return builder
         * 
         */
        public Builder cacheFroms(String... cacheFroms) {
            return cacheFroms(List.of(cacheFroms));
        }

        /**
         * @param cgroupParent Optional parent cgroup for the container
         * 
         * @return builder
         * 
         */
        public Builder cgroupParent(@Nullable Output<String> cgroupParent) {
            $.cgroupParent = cgroupParent;
            return this;
        }

        /**
         * @param cgroupParent Optional parent cgroup for the container
         * 
         * @return builder
         * 
         */
        public Builder cgroupParent(String cgroupParent) {
            return cgroupParent(Output.of(cgroupParent));
        }

        /**
         * @param context Value to specify the build context. Currently, only a `PATH` context is supported. You can use the helper function &#39;${path.cwd}/context-dir&#39;. Please see https://docs.docker.com/build/building/context/ for more information about build contexts.
         * 
         * @return builder
         * 
         */
        public Builder context(Output<String> context) {
            $.context = context;
            return this;
        }

        /**
         * @param context Value to specify the build context. Currently, only a `PATH` context is supported. You can use the helper function &#39;${path.cwd}/context-dir&#39;. Please see https://docs.docker.com/build/building/context/ for more information about build contexts.
         * 
         * @return builder
         * 
         */
        public Builder context(String context) {
            return context(Output.of(context));
        }

        /**
         * @param cpuPeriod The length of a CPU period in microseconds
         * 
         * @return builder
         * 
         */
        public Builder cpuPeriod(@Nullable Output<Integer> cpuPeriod) {
            $.cpuPeriod = cpuPeriod;
            return this;
        }

        /**
         * @param cpuPeriod The length of a CPU period in microseconds
         * 
         * @return builder
         * 
         */
        public Builder cpuPeriod(Integer cpuPeriod) {
            return cpuPeriod(Output.of(cpuPeriod));
        }

        /**
         * @param cpuQuota Microseconds of CPU time that the container can get in a CPU period
         * 
         * @return builder
         * 
         */
        public Builder cpuQuota(@Nullable Output<Integer> cpuQuota) {
            $.cpuQuota = cpuQuota;
            return this;
        }

        /**
         * @param cpuQuota Microseconds of CPU time that the container can get in a CPU period
         * 
         * @return builder
         * 
         */
        public Builder cpuQuota(Integer cpuQuota) {
            return cpuQuota(Output.of(cpuQuota));
        }

        /**
         * @param cpuSetCpus CPUs in which to allow execution (e.g., `0-3`, `0`, `1`)
         * 
         * @return builder
         * 
         */
        public Builder cpuSetCpus(@Nullable Output<String> cpuSetCpus) {
            $.cpuSetCpus = cpuSetCpus;
            return this;
        }

        /**
         * @param cpuSetCpus CPUs in which to allow execution (e.g., `0-3`, `0`, `1`)
         * 
         * @return builder
         * 
         */
        public Builder cpuSetCpus(String cpuSetCpus) {
            return cpuSetCpus(Output.of(cpuSetCpus));
        }

        /**
         * @param cpuSetMems MEMs in which to allow execution (`0-3`, `0`, `1`)
         * 
         * @return builder
         * 
         */
        public Builder cpuSetMems(@Nullable Output<String> cpuSetMems) {
            $.cpuSetMems = cpuSetMems;
            return this;
        }

        /**
         * @param cpuSetMems MEMs in which to allow execution (`0-3`, `0`, `1`)
         * 
         * @return builder
         * 
         */
        public Builder cpuSetMems(String cpuSetMems) {
            return cpuSetMems(Output.of(cpuSetMems));
        }

        /**
         * @param cpuShares CPU shares (relative weight)
         * 
         * @return builder
         * 
         */
        public Builder cpuShares(@Nullable Output<Integer> cpuShares) {
            $.cpuShares = cpuShares;
            return this;
        }

        /**
         * @param cpuShares CPU shares (relative weight)
         * 
         * @return builder
         * 
         */
        public Builder cpuShares(Integer cpuShares) {
            return cpuShares(Output.of(cpuShares));
        }

        /**
         * @param dockerfile Name of the Dockerfile. Defaults to `Dockerfile`.
         * 
         * @return builder
         * 
         */
        public Builder dockerfile(@Nullable Output<String> dockerfile) {
            $.dockerfile = dockerfile;
            return this;
        }

        /**
         * @param dockerfile Name of the Dockerfile. Defaults to `Dockerfile`.
         * 
         * @return builder
         * 
         */
        public Builder dockerfile(String dockerfile) {
            return dockerfile(Output.of(dockerfile));
        }

        /**
         * @param extraHosts A list of hostnames/IP mappings to add to the container’s /etc/hosts file. Specified in the form [&#34;hostname:IP&#34;]
         * 
         * @return builder
         * 
         */
        public Builder extraHosts(@Nullable Output<List<String>> extraHosts) {
            $.extraHosts = extraHosts;
            return this;
        }

        /**
         * @param extraHosts A list of hostnames/IP mappings to add to the container’s /etc/hosts file. Specified in the form [&#34;hostname:IP&#34;]
         * 
         * @return builder
         * 
         */
        public Builder extraHosts(List<String> extraHosts) {
            return extraHosts(Output.of(extraHosts));
        }

        /**
         * @param extraHosts A list of hostnames/IP mappings to add to the container’s /etc/hosts file. Specified in the form [&#34;hostname:IP&#34;]
         * 
         * @return builder
         * 
         */
        public Builder extraHosts(String... extraHosts) {
            return extraHosts(List.of(extraHosts));
        }

        /**
         * @param forceRemove Always remove intermediate containers
         * 
         * @return builder
         * 
         */
        public Builder forceRemove(@Nullable Output<Boolean> forceRemove) {
            $.forceRemove = forceRemove;
            return this;
        }

        /**
         * @param forceRemove Always remove intermediate containers
         * 
         * @return builder
         * 
         */
        public Builder forceRemove(Boolean forceRemove) {
            return forceRemove(Output.of(forceRemove));
        }

        /**
         * @param isolation Isolation represents the isolation technology of a container. The supported values are
         * 
         * @return builder
         * 
         */
        public Builder isolation(@Nullable Output<String> isolation) {
            $.isolation = isolation;
            return this;
        }

        /**
         * @param isolation Isolation represents the isolation technology of a container. The supported values are
         * 
         * @return builder
         * 
         */
        public Builder isolation(String isolation) {
            return isolation(Output.of(isolation));
        }

        /**
         * @param label Set metadata for an image
         * 
         * @return builder
         * 
         */
        public Builder label(@Nullable Output<Map<String,String>> label) {
            $.label = label;
            return this;
        }

        /**
         * @param label Set metadata for an image
         * 
         * @return builder
         * 
         */
        public Builder label(Map<String,String> label) {
            return label(Output.of(label));
        }

        /**
         * @param labels User-defined key/value metadata
         * 
         * @return builder
         * 
         */
        public Builder labels(@Nullable Output<Map<String,String>> labels) {
            $.labels = labels;
            return this;
        }

        /**
         * @param labels User-defined key/value metadata
         * 
         * @return builder
         * 
         */
        public Builder labels(Map<String,String> labels) {
            return labels(Output.of(labels));
        }

        /**
         * @param memory Set memory limit for build
         * 
         * @return builder
         * 
         */
        public Builder memory(@Nullable Output<Integer> memory) {
            $.memory = memory;
            return this;
        }

        /**
         * @param memory Set memory limit for build
         * 
         * @return builder
         * 
         */
        public Builder memory(Integer memory) {
            return memory(Output.of(memory));
        }

        /**
         * @param memorySwap Total memory (memory + swap), -1 to enable unlimited swap
         * 
         * @return builder
         * 
         */
        public Builder memorySwap(@Nullable Output<Integer> memorySwap) {
            $.memorySwap = memorySwap;
            return this;
        }

        /**
         * @param memorySwap Total memory (memory + swap), -1 to enable unlimited swap
         * 
         * @return builder
         * 
         */
        public Builder memorySwap(Integer memorySwap) {
            return memorySwap(Output.of(memorySwap));
        }

        /**
         * @param networkMode Set the networking mode for the RUN instructions during build
         * 
         * @return builder
         * 
         */
        public Builder networkMode(@Nullable Output<String> networkMode) {
            $.networkMode = networkMode;
            return this;
        }

        /**
         * @param networkMode Set the networking mode for the RUN instructions during build
         * 
         * @return builder
         * 
         */
        public Builder networkMode(String networkMode) {
            return networkMode(Output.of(networkMode));
        }

        /**
         * @param noCache Do not use the cache when building the image
         * 
         * @return builder
         * 
         */
        public Builder noCache(@Nullable Output<Boolean> noCache) {
            $.noCache = noCache;
            return this;
        }

        /**
         * @param noCache Do not use the cache when building the image
         * 
         * @return builder
         * 
         */
        public Builder noCache(Boolean noCache) {
            return noCache(Output.of(noCache));
        }

        /**
         * @param platform Set platform if server is multi-platform capable
         * 
         * @return builder
         * 
         */
        public Builder platform(@Nullable Output<String> platform) {
            $.platform = platform;
            return this;
        }

        /**
         * @param platform Set platform if server is multi-platform capable
         * 
         * @return builder
         * 
         */
        public Builder platform(String platform) {
            return platform(Output.of(platform));
        }

        /**
         * @param pullParent Attempt to pull the image even if an older image exists locally
         * 
         * @return builder
         * 
         */
        public Builder pullParent(@Nullable Output<Boolean> pullParent) {
            $.pullParent = pullParent;
            return this;
        }

        /**
         * @param pullParent Attempt to pull the image even if an older image exists locally
         * 
         * @return builder
         * 
         */
        public Builder pullParent(Boolean pullParent) {
            return pullParent(Output.of(pullParent));
        }

        /**
         * @param remoteContext A Git repository URI or HTTP/HTTPS context URI
         * 
         * @return builder
         * 
         */
        public Builder remoteContext(@Nullable Output<String> remoteContext) {
            $.remoteContext = remoteContext;
            return this;
        }

        /**
         * @param remoteContext A Git repository URI or HTTP/HTTPS context URI
         * 
         * @return builder
         * 
         */
        public Builder remoteContext(String remoteContext) {
            return remoteContext(Output.of(remoteContext));
        }

        /**
         * @param remove Remove intermediate containers after a successful build. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder remove(@Nullable Output<Boolean> remove) {
            $.remove = remove;
            return this;
        }

        /**
         * @param remove Remove intermediate containers after a successful build. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder remove(Boolean remove) {
            return remove(Output.of(remove));
        }

        /**
         * @param securityOpts The security options
         * 
         * @return builder
         * 
         */
        public Builder securityOpts(@Nullable Output<List<String>> securityOpts) {
            $.securityOpts = securityOpts;
            return this;
        }

        /**
         * @param securityOpts The security options
         * 
         * @return builder
         * 
         */
        public Builder securityOpts(List<String> securityOpts) {
            return securityOpts(Output.of(securityOpts));
        }

        /**
         * @param securityOpts The security options
         * 
         * @return builder
         * 
         */
        public Builder securityOpts(String... securityOpts) {
            return securityOpts(List.of(securityOpts));
        }

        /**
         * @param sessionId Set an ID for the build session
         * 
         * @return builder
         * 
         */
        public Builder sessionId(@Nullable Output<String> sessionId) {
            $.sessionId = sessionId;
            return this;
        }

        /**
         * @param sessionId Set an ID for the build session
         * 
         * @return builder
         * 
         */
        public Builder sessionId(String sessionId) {
            return sessionId(Output.of(sessionId));
        }

        /**
         * @param shmSize Size of /dev/shm in bytes. The size must be greater than 0
         * 
         * @return builder
         * 
         */
        public Builder shmSize(@Nullable Output<Integer> shmSize) {
            $.shmSize = shmSize;
            return this;
        }

        /**
         * @param shmSize Size of /dev/shm in bytes. The size must be greater than 0
         * 
         * @return builder
         * 
         */
        public Builder shmSize(Integer shmSize) {
            return shmSize(Output.of(shmSize));
        }

        /**
         * @param squash If true the new layers are squashed into a new image with a single new layer
         * 
         * @return builder
         * 
         */
        public Builder squash(@Nullable Output<Boolean> squash) {
            $.squash = squash;
            return this;
        }

        /**
         * @param squash If true the new layers are squashed into a new image with a single new layer
         * 
         * @return builder
         * 
         */
        public Builder squash(Boolean squash) {
            return squash(Output.of(squash));
        }

        /**
         * @param suppressOutput Suppress the build output and print image ID on success
         * 
         * @return builder
         * 
         */
        public Builder suppressOutput(@Nullable Output<Boolean> suppressOutput) {
            $.suppressOutput = suppressOutput;
            return this;
        }

        /**
         * @param suppressOutput Suppress the build output and print image ID on success
         * 
         * @return builder
         * 
         */
        public Builder suppressOutput(Boolean suppressOutput) {
            return suppressOutput(Output.of(suppressOutput));
        }

        /**
         * @param tags Name and optionally a tag in the &#39;name:tag&#39; format
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<List<String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Name and optionally a tag in the &#39;name:tag&#39; format
         * 
         * @return builder
         * 
         */
        public Builder tags(List<String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tags Name and optionally a tag in the &#39;name:tag&#39; format
         * 
         * @return builder
         * 
         */
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }

        /**
         * @param target Set the target build stage to build
         * 
         * @return builder
         * 
         */
        public Builder target(@Nullable Output<String> target) {
            $.target = target;
            return this;
        }

        /**
         * @param target Set the target build stage to build
         * 
         * @return builder
         * 
         */
        public Builder target(String target) {
            return target(Output.of(target));
        }

        /**
         * @param ulimits Configuration for ulimits
         * 
         * @return builder
         * 
         */
        public Builder ulimits(@Nullable Output<List<RemoteImageBuildUlimitArgs>> ulimits) {
            $.ulimits = ulimits;
            return this;
        }

        /**
         * @param ulimits Configuration for ulimits
         * 
         * @return builder
         * 
         */
        public Builder ulimits(List<RemoteImageBuildUlimitArgs> ulimits) {
            return ulimits(Output.of(ulimits));
        }

        /**
         * @param ulimits Configuration for ulimits
         * 
         * @return builder
         * 
         */
        public Builder ulimits(RemoteImageBuildUlimitArgs... ulimits) {
            return ulimits(List.of(ulimits));
        }

        /**
         * @param version Version of the underlying builder to use
         * 
         * @return builder
         * 
         */
        public Builder version(@Nullable Output<String> version) {
            $.version = version;
            return this;
        }

        /**
         * @param version Version of the underlying builder to use
         * 
         * @return builder
         * 
         */
        public Builder version(String version) {
            return version(Output.of(version));
        }

        public RemoteImageBuildArgs build() {
            if ($.context == null) {
                throw new MissingRequiredPropertyException("RemoteImageBuildArgs", "context");
            }
            return $;
        }
    }

}
