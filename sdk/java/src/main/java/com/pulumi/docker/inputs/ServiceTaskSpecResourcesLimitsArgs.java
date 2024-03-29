// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceTaskSpecResourcesLimitsArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceTaskSpecResourcesLimitsArgs Empty = new ServiceTaskSpecResourcesLimitsArgs();

    /**
     * The amounf of memory in bytes the container allocates
     * 
     */
    @Import(name="memoryBytes")
    private @Nullable Output<Integer> memoryBytes;

    /**
     * @return The amounf of memory in bytes the container allocates
     * 
     */
    public Optional<Output<Integer>> memoryBytes() {
        return Optional.ofNullable(this.memoryBytes);
    }

    /**
     * CPU shares in units of `1/1e9` (or `10^-9`) of the CPU. Should be at least `1000000`
     * 
     */
    @Import(name="nanoCpus")
    private @Nullable Output<Integer> nanoCpus;

    /**
     * @return CPU shares in units of `1/1e9` (or `10^-9`) of the CPU. Should be at least `1000000`
     * 
     */
    public Optional<Output<Integer>> nanoCpus() {
        return Optional.ofNullable(this.nanoCpus);
    }

    private ServiceTaskSpecResourcesLimitsArgs() {}

    private ServiceTaskSpecResourcesLimitsArgs(ServiceTaskSpecResourcesLimitsArgs $) {
        this.memoryBytes = $.memoryBytes;
        this.nanoCpus = $.nanoCpus;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceTaskSpecResourcesLimitsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceTaskSpecResourcesLimitsArgs $;

        public Builder() {
            $ = new ServiceTaskSpecResourcesLimitsArgs();
        }

        public Builder(ServiceTaskSpecResourcesLimitsArgs defaults) {
            $ = new ServiceTaskSpecResourcesLimitsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param memoryBytes The amounf of memory in bytes the container allocates
         * 
         * @return builder
         * 
         */
        public Builder memoryBytes(@Nullable Output<Integer> memoryBytes) {
            $.memoryBytes = memoryBytes;
            return this;
        }

        /**
         * @param memoryBytes The amounf of memory in bytes the container allocates
         * 
         * @return builder
         * 
         */
        public Builder memoryBytes(Integer memoryBytes) {
            return memoryBytes(Output.of(memoryBytes));
        }

        /**
         * @param nanoCpus CPU shares in units of `1/1e9` (or `10^-9`) of the CPU. Should be at least `1000000`
         * 
         * @return builder
         * 
         */
        public Builder nanoCpus(@Nullable Output<Integer> nanoCpus) {
            $.nanoCpus = nanoCpus;
            return this;
        }

        /**
         * @param nanoCpus CPU shares in units of `1/1e9` (or `10^-9`) of the CPU. Should be at least `1000000`
         * 
         * @return builder
         * 
         */
        public Builder nanoCpus(Integer nanoCpus) {
            return nanoCpus(Output.of(nanoCpus));
        }

        public ServiceTaskSpecResourcesLimitsArgs build() {
            return $;
        }
    }

}
