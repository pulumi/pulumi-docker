// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ServiceEndpointSpecPort {
    /**
     * @return A random name for the port
     * 
     */
    private @Nullable String name;
    /**
     * @return Rrepresents the protocol of a port: `tcp`, `udp` or `sctp`. Defaults to `tcp`.
     * 
     */
    private @Nullable String protocol;
    /**
     * @return Represents the mode in which the port is to be published: &#39;ingress&#39; or &#39;host&#39;. Defaults to `ingress`.
     * 
     */
    private @Nullable String publishMode;
    /**
     * @return The port on the swarm hosts
     * 
     */
    private @Nullable Integer publishedPort;
    /**
     * @return The port inside the container
     * 
     */
    private Integer targetPort;

    private ServiceEndpointSpecPort() {}
    /**
     * @return A random name for the port
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    /**
     * @return Rrepresents the protocol of a port: `tcp`, `udp` or `sctp`. Defaults to `tcp`.
     * 
     */
    public Optional<String> protocol() {
        return Optional.ofNullable(this.protocol);
    }
    /**
     * @return Represents the mode in which the port is to be published: &#39;ingress&#39; or &#39;host&#39;. Defaults to `ingress`.
     * 
     */
    public Optional<String> publishMode() {
        return Optional.ofNullable(this.publishMode);
    }
    /**
     * @return The port on the swarm hosts
     * 
     */
    public Optional<Integer> publishedPort() {
        return Optional.ofNullable(this.publishedPort);
    }
    /**
     * @return The port inside the container
     * 
     */
    public Integer targetPort() {
        return this.targetPort;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ServiceEndpointSpecPort defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String name;
        private @Nullable String protocol;
        private @Nullable String publishMode;
        private @Nullable Integer publishedPort;
        private Integer targetPort;
        public Builder() {}
        public Builder(ServiceEndpointSpecPort defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.name = defaults.name;
    	      this.protocol = defaults.protocol;
    	      this.publishMode = defaults.publishMode;
    	      this.publishedPort = defaults.publishedPort;
    	      this.targetPort = defaults.targetPort;
        }

        @CustomType.Setter
        public Builder name(@Nullable String name) {

            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder protocol(@Nullable String protocol) {

            this.protocol = protocol;
            return this;
        }
        @CustomType.Setter
        public Builder publishMode(@Nullable String publishMode) {

            this.publishMode = publishMode;
            return this;
        }
        @CustomType.Setter
        public Builder publishedPort(@Nullable Integer publishedPort) {

            this.publishedPort = publishedPort;
            return this;
        }
        @CustomType.Setter
        public Builder targetPort(Integer targetPort) {
            if (targetPort == null) {
              throw new MissingRequiredPropertyException("ServiceEndpointSpecPort", "targetPort");
            }
            this.targetPort = targetPort;
            return this;
        }
        public ServiceEndpointSpecPort build() {
            final var _resultValue = new ServiceEndpointSpecPort();
            _resultValue.name = name;
            _resultValue.protocol = protocol;
            _resultValue.publishMode = publishMode;
            _resultValue.publishedPort = publishedPort;
            _resultValue.targetPort = targetPort;
            return _resultValue;
        }
    }
}
