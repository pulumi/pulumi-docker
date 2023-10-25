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
public final class ServiceTaskSpecRestartPolicy {
    /**
     * @return Condition for restart
     * 
     */
    private @Nullable String condition;
    /**
     * @return The interval to check if the desired state is reached `(ms|s)`. Defaults to `7s`.
     * 
     */
    private @Nullable String delay;
    /**
     * @return Maximum attempts to restart a given container before giving up (default value is `0`, which is ignored)
     * 
     */
    private @Nullable Integer maxAttempts;
    /**
     * @return The time window used to evaluate the restart policy (default value is `0`, which is unbounded) (ms|s|m|h)
     * 
     */
    private @Nullable String window;

    private ServiceTaskSpecRestartPolicy() {}
    /**
     * @return Condition for restart
     * 
     */
    public Optional<String> condition() {
        return Optional.ofNullable(this.condition);
    }
    /**
     * @return The interval to check if the desired state is reached `(ms|s)`. Defaults to `7s`.
     * 
     */
    public Optional<String> delay() {
        return Optional.ofNullable(this.delay);
    }
    /**
     * @return Maximum attempts to restart a given container before giving up (default value is `0`, which is ignored)
     * 
     */
    public Optional<Integer> maxAttempts() {
        return Optional.ofNullable(this.maxAttempts);
    }
    /**
     * @return The time window used to evaluate the restart policy (default value is `0`, which is unbounded) (ms|s|m|h)
     * 
     */
    public Optional<String> window() {
        return Optional.ofNullable(this.window);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ServiceTaskSpecRestartPolicy defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String condition;
        private @Nullable String delay;
        private @Nullable Integer maxAttempts;
        private @Nullable String window;
        public Builder() {}
        public Builder(ServiceTaskSpecRestartPolicy defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.condition = defaults.condition;
    	      this.delay = defaults.delay;
    	      this.maxAttempts = defaults.maxAttempts;
    	      this.window = defaults.window;
        }

        @CustomType.Setter
        public Builder condition(@Nullable String condition) {
            this.condition = condition;
            return this;
        }
        @CustomType.Setter
        public Builder delay(@Nullable String delay) {
            this.delay = delay;
            return this;
        }
        @CustomType.Setter
        public Builder maxAttempts(@Nullable Integer maxAttempts) {
            this.maxAttempts = maxAttempts;
            return this;
        }
        @CustomType.Setter
        public Builder window(@Nullable String window) {
            this.window = window;
            return this;
        }
        public ServiceTaskSpecRestartPolicy build() {
            final var o = new ServiceTaskSpecRestartPolicy();
            o.condition = condition;
            o.delay = delay;
            o.maxAttempts = maxAttempts;
            o.window = window;
            return o;
        }
    }
}
