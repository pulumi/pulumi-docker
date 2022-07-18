// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.docker.outputs.ServiceTaskSpecResourcesReservationGenericResources;
import java.lang.Integer;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ServiceTaskSpecResourcesReservation {
    private final @Nullable ServiceTaskSpecResourcesReservationGenericResources genericResources;
    private final @Nullable Integer memoryBytes;
    private final @Nullable Integer nanoCpus;

    @CustomType.Constructor
    private ServiceTaskSpecResourcesReservation(
        @CustomType.Parameter("genericResources") @Nullable ServiceTaskSpecResourcesReservationGenericResources genericResources,
        @CustomType.Parameter("memoryBytes") @Nullable Integer memoryBytes,
        @CustomType.Parameter("nanoCpus") @Nullable Integer nanoCpus) {
        this.genericResources = genericResources;
        this.memoryBytes = memoryBytes;
        this.nanoCpus = nanoCpus;
    }

    public Optional<ServiceTaskSpecResourcesReservationGenericResources> genericResources() {
        return Optional.ofNullable(this.genericResources);
    }
    public Optional<Integer> memoryBytes() {
        return Optional.ofNullable(this.memoryBytes);
    }
    public Optional<Integer> nanoCpus() {
        return Optional.ofNullable(this.nanoCpus);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ServiceTaskSpecResourcesReservation defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable ServiceTaskSpecResourcesReservationGenericResources genericResources;
        private @Nullable Integer memoryBytes;
        private @Nullable Integer nanoCpus;

        public Builder() {
    	      // Empty
        }

        public Builder(ServiceTaskSpecResourcesReservation defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.genericResources = defaults.genericResources;
    	      this.memoryBytes = defaults.memoryBytes;
    	      this.nanoCpus = defaults.nanoCpus;
        }

        public Builder genericResources(@Nullable ServiceTaskSpecResourcesReservationGenericResources genericResources) {
            this.genericResources = genericResources;
            return this;
        }
        public Builder memoryBytes(@Nullable Integer memoryBytes) {
            this.memoryBytes = memoryBytes;
            return this;
        }
        public Builder nanoCpus(@Nullable Integer nanoCpus) {
            this.nanoCpus = nanoCpus;
            return this;
        }        public ServiceTaskSpecResourcesReservation build() {
            return new ServiceTaskSpecResourcesReservation(genericResources, memoryBytes, nanoCpus);
        }
    }
}