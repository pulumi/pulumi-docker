// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ServiceEndpointSpecPort {
    /**
     * @return Name of the service
     * 
     */
    private final @Nullable String name;
    private final @Nullable String protocol;
    private final @Nullable String publishMode;
    private final @Nullable Integer publishedPort;
    private final Integer targetPort;

    @CustomType.Constructor
    private ServiceEndpointSpecPort(
        @CustomType.Parameter("name") @Nullable String name,
        @CustomType.Parameter("protocol") @Nullable String protocol,
        @CustomType.Parameter("publishMode") @Nullable String publishMode,
        @CustomType.Parameter("publishedPort") @Nullable Integer publishedPort,
        @CustomType.Parameter("targetPort") Integer targetPort) {
        this.name = name;
        this.protocol = protocol;
        this.publishMode = publishMode;
        this.publishedPort = publishedPort;
        this.targetPort = targetPort;
    }

    /**
     * @return Name of the service
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    public Optional<String> protocol() {
        return Optional.ofNullable(this.protocol);
    }
    public Optional<String> publishMode() {
        return Optional.ofNullable(this.publishMode);
    }
    public Optional<Integer> publishedPort() {
        return Optional.ofNullable(this.publishedPort);
    }
    public Integer targetPort() {
        return this.targetPort;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ServiceEndpointSpecPort defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable String name;
        private @Nullable String protocol;
        private @Nullable String publishMode;
        private @Nullable Integer publishedPort;
        private Integer targetPort;

        public Builder() {
    	      // Empty
        }

        public Builder(ServiceEndpointSpecPort defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.name = defaults.name;
    	      this.protocol = defaults.protocol;
    	      this.publishMode = defaults.publishMode;
    	      this.publishedPort = defaults.publishedPort;
    	      this.targetPort = defaults.targetPort;
        }

        public Builder name(@Nullable String name) {
            this.name = name;
            return this;
        }
        public Builder protocol(@Nullable String protocol) {
            this.protocol = protocol;
            return this;
        }
        public Builder publishMode(@Nullable String publishMode) {
            this.publishMode = publishMode;
            return this;
        }
        public Builder publishedPort(@Nullable Integer publishedPort) {
            this.publishedPort = publishedPort;
            return this;
        }
        public Builder targetPort(Integer targetPort) {
            this.targetPort = Objects.requireNonNull(targetPort);
            return this;
        }        public ServiceEndpointSpecPort build() {
            return new ServiceEndpointSpecPort(name, protocol, publishMode, publishedPort, targetPort);
        }
    }
}