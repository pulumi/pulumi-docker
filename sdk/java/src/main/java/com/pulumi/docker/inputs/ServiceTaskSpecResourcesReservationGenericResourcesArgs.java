// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceTaskSpecResourcesReservationGenericResourcesArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceTaskSpecResourcesReservationGenericResourcesArgs Empty = new ServiceTaskSpecResourcesReservationGenericResourcesArgs();

    /**
     * The Integer resources
     * 
     */
    @Import(name="discreteResourcesSpecs")
    private @Nullable Output<List<String>> discreteResourcesSpecs;

    /**
     * @return The Integer resources
     * 
     */
    public Optional<Output<List<String>>> discreteResourcesSpecs() {
        return Optional.ofNullable(this.discreteResourcesSpecs);
    }

    /**
     * The String resources
     * 
     */
    @Import(name="namedResourcesSpecs")
    private @Nullable Output<List<String>> namedResourcesSpecs;

    /**
     * @return The String resources
     * 
     */
    public Optional<Output<List<String>>> namedResourcesSpecs() {
        return Optional.ofNullable(this.namedResourcesSpecs);
    }

    private ServiceTaskSpecResourcesReservationGenericResourcesArgs() {}

    private ServiceTaskSpecResourcesReservationGenericResourcesArgs(ServiceTaskSpecResourcesReservationGenericResourcesArgs $) {
        this.discreteResourcesSpecs = $.discreteResourcesSpecs;
        this.namedResourcesSpecs = $.namedResourcesSpecs;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceTaskSpecResourcesReservationGenericResourcesArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceTaskSpecResourcesReservationGenericResourcesArgs $;

        public Builder() {
            $ = new ServiceTaskSpecResourcesReservationGenericResourcesArgs();
        }

        public Builder(ServiceTaskSpecResourcesReservationGenericResourcesArgs defaults) {
            $ = new ServiceTaskSpecResourcesReservationGenericResourcesArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param discreteResourcesSpecs The Integer resources
         * 
         * @return builder
         * 
         */
        public Builder discreteResourcesSpecs(@Nullable Output<List<String>> discreteResourcesSpecs) {
            $.discreteResourcesSpecs = discreteResourcesSpecs;
            return this;
        }

        /**
         * @param discreteResourcesSpecs The Integer resources
         * 
         * @return builder
         * 
         */
        public Builder discreteResourcesSpecs(List<String> discreteResourcesSpecs) {
            return discreteResourcesSpecs(Output.of(discreteResourcesSpecs));
        }

        /**
         * @param discreteResourcesSpecs The Integer resources
         * 
         * @return builder
         * 
         */
        public Builder discreteResourcesSpecs(String... discreteResourcesSpecs) {
            return discreteResourcesSpecs(List.of(discreteResourcesSpecs));
        }

        /**
         * @param namedResourcesSpecs The String resources
         * 
         * @return builder
         * 
         */
        public Builder namedResourcesSpecs(@Nullable Output<List<String>> namedResourcesSpecs) {
            $.namedResourcesSpecs = namedResourcesSpecs;
            return this;
        }

        /**
         * @param namedResourcesSpecs The String resources
         * 
         * @return builder
         * 
         */
        public Builder namedResourcesSpecs(List<String> namedResourcesSpecs) {
            return namedResourcesSpecs(Output.of(namedResourcesSpecs));
        }

        /**
         * @param namedResourcesSpecs The String resources
         * 
         * @return builder
         * 
         */
        public Builder namedResourcesSpecs(String... namedResourcesSpecs) {
            return namedResourcesSpecs(List.of(namedResourcesSpecs));
        }

        public ServiceTaskSpecResourcesReservationGenericResourcesArgs build() {
            return $;
        }
    }

}
