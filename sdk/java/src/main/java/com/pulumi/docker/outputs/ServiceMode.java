// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.docker.outputs.ServiceModeReplicated;
import java.lang.Boolean;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ServiceMode {
    /**
     * @return When `true`, tasks will run on every worker node. Conflicts with `replicated`
     * 
     */
    private @Nullable Boolean global;
    /**
     * @return The replicated service mode
     * 
     */
    private @Nullable ServiceModeReplicated replicated;

    private ServiceMode() {}
    /**
     * @return When `true`, tasks will run on every worker node. Conflicts with `replicated`
     * 
     */
    public Optional<Boolean> global() {
        return Optional.ofNullable(this.global);
    }
    /**
     * @return The replicated service mode
     * 
     */
    public Optional<ServiceModeReplicated> replicated() {
        return Optional.ofNullable(this.replicated);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ServiceMode defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean global;
        private @Nullable ServiceModeReplicated replicated;
        public Builder() {}
        public Builder(ServiceMode defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.global = defaults.global;
    	      this.replicated = defaults.replicated;
        }

        @CustomType.Setter
        public Builder global(@Nullable Boolean global) {

            this.global = global;
            return this;
        }
        @CustomType.Setter
        public Builder replicated(@Nullable ServiceModeReplicated replicated) {

            this.replicated = replicated;
            return this;
        }
        public ServiceMode build() {
            final var _resultValue = new ServiceMode();
            _resultValue.global = global;
            _resultValue.replicated = replicated;
            return _resultValue;
        }
    }
}
