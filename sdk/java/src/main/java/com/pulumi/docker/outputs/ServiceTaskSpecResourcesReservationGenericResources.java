// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class ServiceTaskSpecResourcesReservationGenericResources {
    /**
     * @return The Integer resources
     * 
     */
    private @Nullable List<String> discreteResourcesSpecs;
    /**
     * @return The String resources
     * 
     */
    private @Nullable List<String> namedResourcesSpecs;

    private ServiceTaskSpecResourcesReservationGenericResources() {}
    /**
     * @return The Integer resources
     * 
     */
    public List<String> discreteResourcesSpecs() {
        return this.discreteResourcesSpecs == null ? List.of() : this.discreteResourcesSpecs;
    }
    /**
     * @return The String resources
     * 
     */
    public List<String> namedResourcesSpecs() {
        return this.namedResourcesSpecs == null ? List.of() : this.namedResourcesSpecs;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ServiceTaskSpecResourcesReservationGenericResources defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<String> discreteResourcesSpecs;
        private @Nullable List<String> namedResourcesSpecs;
        public Builder() {}
        public Builder(ServiceTaskSpecResourcesReservationGenericResources defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.discreteResourcesSpecs = defaults.discreteResourcesSpecs;
    	      this.namedResourcesSpecs = defaults.namedResourcesSpecs;
        }

        @CustomType.Setter
        public Builder discreteResourcesSpecs(@Nullable List<String> discreteResourcesSpecs) {

            this.discreteResourcesSpecs = discreteResourcesSpecs;
            return this;
        }
        public Builder discreteResourcesSpecs(String... discreteResourcesSpecs) {
            return discreteResourcesSpecs(List.of(discreteResourcesSpecs));
        }
        @CustomType.Setter
        public Builder namedResourcesSpecs(@Nullable List<String> namedResourcesSpecs) {

            this.namedResourcesSpecs = namedResourcesSpecs;
            return this;
        }
        public Builder namedResourcesSpecs(String... namedResourcesSpecs) {
            return namedResourcesSpecs(List.of(namedResourcesSpecs));
        }
        public ServiceTaskSpecResourcesReservationGenericResources build() {
            final var _resultValue = new ServiceTaskSpecResourcesReservationGenericResources();
            _resultValue.discreteResourcesSpecs = discreteResourcesSpecs;
            _resultValue.namedResourcesSpecs = namedResourcesSpecs;
            return _resultValue;
        }
    }
}
