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
public final class ContainerPort {
    /**
     * @return Port exposed out of the container. If not given a free random port `&gt;= 32768` will be used.
     * 
     */
    private @Nullable Integer external;
    /**
     * @return Port within the container.
     * 
     */
    private Integer internal;
    /**
     * @return IP address/mask that can access this port. Defaults to `0.0.0.0`.
     * 
     */
    private @Nullable String ip;
    /**
     * @return Protocol that can be used over this port. Defaults to `tcp`.
     * 
     */
    private @Nullable String protocol;

    private ContainerPort() {}
    /**
     * @return Port exposed out of the container. If not given a free random port `&gt;= 32768` will be used.
     * 
     */
    public Optional<Integer> external() {
        return Optional.ofNullable(this.external);
    }
    /**
     * @return Port within the container.
     * 
     */
    public Integer internal() {
        return this.internal;
    }
    /**
     * @return IP address/mask that can access this port. Defaults to `0.0.0.0`.
     * 
     */
    public Optional<String> ip() {
        return Optional.ofNullable(this.ip);
    }
    /**
     * @return Protocol that can be used over this port. Defaults to `tcp`.
     * 
     */
    public Optional<String> protocol() {
        return Optional.ofNullable(this.protocol);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ContainerPort defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Integer external;
        private Integer internal;
        private @Nullable String ip;
        private @Nullable String protocol;
        public Builder() {}
        public Builder(ContainerPort defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.external = defaults.external;
    	      this.internal = defaults.internal;
    	      this.ip = defaults.ip;
    	      this.protocol = defaults.protocol;
        }

        @CustomType.Setter
        public Builder external(@Nullable Integer external) {
            this.external = external;
            return this;
        }
        @CustomType.Setter
        public Builder internal(Integer internal) {
            this.internal = Objects.requireNonNull(internal);
            return this;
        }
        @CustomType.Setter
        public Builder ip(@Nullable String ip) {
            this.ip = ip;
            return this;
        }
        @CustomType.Setter
        public Builder protocol(@Nullable String protocol) {
            this.protocol = protocol;
            return this;
        }
        public ContainerPort build() {
            final var _resultValue = new ContainerPort();
            _resultValue.external = external;
            _resultValue.internal = internal;
            _resultValue.ip = ip;
            _resultValue.protocol = protocol;
            return _resultValue;
        }
    }
}
