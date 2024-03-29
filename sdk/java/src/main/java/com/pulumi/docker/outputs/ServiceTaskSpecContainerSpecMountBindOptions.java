// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ServiceTaskSpecContainerSpecMountBindOptions {
    /**
     * @return Bind propagation refers to whether or not mounts created within a given bind-mount or named volume can be propagated to replicas of that mount. See the [docs](https://docs.docker.com/storage/bind-mounts/#configure-bind-propagation) for details. Defaults to `rprivate`
     * 
     */
    private @Nullable String propagation;

    private ServiceTaskSpecContainerSpecMountBindOptions() {}
    /**
     * @return Bind propagation refers to whether or not mounts created within a given bind-mount or named volume can be propagated to replicas of that mount. See the [docs](https://docs.docker.com/storage/bind-mounts/#configure-bind-propagation) for details. Defaults to `rprivate`
     * 
     */
    public Optional<String> propagation() {
        return Optional.ofNullable(this.propagation);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ServiceTaskSpecContainerSpecMountBindOptions defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String propagation;
        public Builder() {}
        public Builder(ServiceTaskSpecContainerSpecMountBindOptions defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.propagation = defaults.propagation;
        }

        @CustomType.Setter
        public Builder propagation(@Nullable String propagation) {

            this.propagation = propagation;
            return this;
        }
        public ServiceTaskSpecContainerSpecMountBindOptions build() {
            final var _resultValue = new ServiceTaskSpecContainerSpecMountBindOptions();
            _resultValue.propagation = propagation;
            return _resultValue;
        }
    }
}
