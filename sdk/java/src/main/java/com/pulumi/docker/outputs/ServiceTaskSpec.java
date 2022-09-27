// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.docker.outputs.ServiceTaskSpecContainerSpec;
import com.pulumi.docker.outputs.ServiceTaskSpecLogDriver;
import com.pulumi.docker.outputs.ServiceTaskSpecPlacement;
import com.pulumi.docker.outputs.ServiceTaskSpecResources;
import com.pulumi.docker.outputs.ServiceTaskSpecRestartPolicy;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ServiceTaskSpec {
    /**
     * @return The spec for each container
     * 
     */
    private ServiceTaskSpecContainerSpec containerSpec;
    /**
     * @return A counter that triggers an update even if no relevant parameters have been changed. See the [spec](https://github.com/docker/swarmkit/blob/master/api/specs.proto#L126).
     * 
     */
    private @Nullable Integer forceUpdate;
    /**
     * @return Specifies the log driver to use for tasks created from this spec. If not present, the default one for the swarm will be used, finally falling back to the engine default if not specified
     * 
     */
    private @Nullable ServiceTaskSpecLogDriver logDriver;
    /**
     * @return Ids of the networks in which the  container will be put in
     * 
     */
    private @Nullable List<String> networks;
    /**
     * @return The placement preferences
     * 
     */
    private @Nullable ServiceTaskSpecPlacement placement;
    /**
     * @return Resource requirements which apply to each individual container created as part of the service
     * 
     */
    private @Nullable ServiceTaskSpecResources resources;
    /**
     * @return Specification for the restart policy which applies to containers created as part of this service.
     * 
     */
    private @Nullable ServiceTaskSpecRestartPolicy restartPolicy;
    /**
     * @return Runtime is the type of runtime specified for the task executor. See the [types](https://github.com/moby/moby/blob/master/api/types/swarm/runtime.go).
     * 
     */
    private @Nullable String runtime;

    private ServiceTaskSpec() {}
    /**
     * @return The spec for each container
     * 
     */
    public ServiceTaskSpecContainerSpec containerSpec() {
        return this.containerSpec;
    }
    /**
     * @return A counter that triggers an update even if no relevant parameters have been changed. See the [spec](https://github.com/docker/swarmkit/blob/master/api/specs.proto#L126).
     * 
     */
    public Optional<Integer> forceUpdate() {
        return Optional.ofNullable(this.forceUpdate);
    }
    /**
     * @return Specifies the log driver to use for tasks created from this spec. If not present, the default one for the swarm will be used, finally falling back to the engine default if not specified
     * 
     */
    public Optional<ServiceTaskSpecLogDriver> logDriver() {
        return Optional.ofNullable(this.logDriver);
    }
    /**
     * @return Ids of the networks in which the  container will be put in
     * 
     */
    public List<String> networks() {
        return this.networks == null ? List.of() : this.networks;
    }
    /**
     * @return The placement preferences
     * 
     */
    public Optional<ServiceTaskSpecPlacement> placement() {
        return Optional.ofNullable(this.placement);
    }
    /**
     * @return Resource requirements which apply to each individual container created as part of the service
     * 
     */
    public Optional<ServiceTaskSpecResources> resources() {
        return Optional.ofNullable(this.resources);
    }
    /**
     * @return Specification for the restart policy which applies to containers created as part of this service.
     * 
     */
    public Optional<ServiceTaskSpecRestartPolicy> restartPolicy() {
        return Optional.ofNullable(this.restartPolicy);
    }
    /**
     * @return Runtime is the type of runtime specified for the task executor. See the [types](https://github.com/moby/moby/blob/master/api/types/swarm/runtime.go).
     * 
     */
    public Optional<String> runtime() {
        return Optional.ofNullable(this.runtime);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ServiceTaskSpec defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private ServiceTaskSpecContainerSpec containerSpec;
        private @Nullable Integer forceUpdate;
        private @Nullable ServiceTaskSpecLogDriver logDriver;
        private @Nullable List<String> networks;
        private @Nullable ServiceTaskSpecPlacement placement;
        private @Nullable ServiceTaskSpecResources resources;
        private @Nullable ServiceTaskSpecRestartPolicy restartPolicy;
        private @Nullable String runtime;
        public Builder() {}
        public Builder(ServiceTaskSpec defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.containerSpec = defaults.containerSpec;
    	      this.forceUpdate = defaults.forceUpdate;
    	      this.logDriver = defaults.logDriver;
    	      this.networks = defaults.networks;
    	      this.placement = defaults.placement;
    	      this.resources = defaults.resources;
    	      this.restartPolicy = defaults.restartPolicy;
    	      this.runtime = defaults.runtime;
        }

        @CustomType.Setter
        public Builder containerSpec(ServiceTaskSpecContainerSpec containerSpec) {
            this.containerSpec = Objects.requireNonNull(containerSpec);
            return this;
        }
        @CustomType.Setter
        public Builder forceUpdate(@Nullable Integer forceUpdate) {
            this.forceUpdate = forceUpdate;
            return this;
        }
        @CustomType.Setter
        public Builder logDriver(@Nullable ServiceTaskSpecLogDriver logDriver) {
            this.logDriver = logDriver;
            return this;
        }
        @CustomType.Setter
        public Builder networks(@Nullable List<String> networks) {
            this.networks = networks;
            return this;
        }
        public Builder networks(String... networks) {
            return networks(List.of(networks));
        }
        @CustomType.Setter
        public Builder placement(@Nullable ServiceTaskSpecPlacement placement) {
            this.placement = placement;
            return this;
        }
        @CustomType.Setter
        public Builder resources(@Nullable ServiceTaskSpecResources resources) {
            this.resources = resources;
            return this;
        }
        @CustomType.Setter
        public Builder restartPolicy(@Nullable ServiceTaskSpecRestartPolicy restartPolicy) {
            this.restartPolicy = restartPolicy;
            return this;
        }
        @CustomType.Setter
        public Builder runtime(@Nullable String runtime) {
            this.runtime = runtime;
            return this;
        }
        public ServiceTaskSpec build() {
            final var o = new ServiceTaskSpec();
            o.containerSpec = containerSpec;
            o.forceUpdate = forceUpdate;
            o.logDriver = logDriver;
            o.networks = networks;
            o.placement = placement;
            o.resources = resources;
            o.restartPolicy = restartPolicy;
            o.runtime = runtime;
            return o;
        }
    }
}
