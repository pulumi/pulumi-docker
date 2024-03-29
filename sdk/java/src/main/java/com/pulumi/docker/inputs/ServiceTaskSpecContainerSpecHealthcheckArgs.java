// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceTaskSpecContainerSpecHealthcheckArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceTaskSpecContainerSpecHealthcheckArgs Empty = new ServiceTaskSpecContainerSpecHealthcheckArgs();

    /**
     * Time between running the check (ms|s|m|h). Defaults to `0s`.
     * 
     */
    @Import(name="interval")
    private @Nullable Output<String> interval;

    /**
     * @return Time between running the check (ms|s|m|h). Defaults to `0s`.
     * 
     */
    public Optional<Output<String>> interval() {
        return Optional.ofNullable(this.interval);
    }

    /**
     * Consecutive failures needed to report unhealthy. Defaults to `0`
     * 
     */
    @Import(name="retries")
    private @Nullable Output<Integer> retries;

    /**
     * @return Consecutive failures needed to report unhealthy. Defaults to `0`
     * 
     */
    public Optional<Output<Integer>> retries() {
        return Optional.ofNullable(this.retries);
    }

    /**
     * Start period for the container to initialize before counting retries towards unstable (ms|s|m|h). Defaults to `0s`.
     * 
     */
    @Import(name="startPeriod")
    private @Nullable Output<String> startPeriod;

    /**
     * @return Start period for the container to initialize before counting retries towards unstable (ms|s|m|h). Defaults to `0s`.
     * 
     */
    public Optional<Output<String>> startPeriod() {
        return Optional.ofNullable(this.startPeriod);
    }

    /**
     * The test to perform as list
     * 
     */
    @Import(name="tests", required=true)
    private Output<List<String>> tests;

    /**
     * @return The test to perform as list
     * 
     */
    public Output<List<String>> tests() {
        return this.tests;
    }

    /**
     * Maximum time to allow one check to run (ms|s|m|h). Defaults to `0s`.
     * 
     */
    @Import(name="timeout")
    private @Nullable Output<String> timeout;

    /**
     * @return Maximum time to allow one check to run (ms|s|m|h). Defaults to `0s`.
     * 
     */
    public Optional<Output<String>> timeout() {
        return Optional.ofNullable(this.timeout);
    }

    private ServiceTaskSpecContainerSpecHealthcheckArgs() {}

    private ServiceTaskSpecContainerSpecHealthcheckArgs(ServiceTaskSpecContainerSpecHealthcheckArgs $) {
        this.interval = $.interval;
        this.retries = $.retries;
        this.startPeriod = $.startPeriod;
        this.tests = $.tests;
        this.timeout = $.timeout;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceTaskSpecContainerSpecHealthcheckArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceTaskSpecContainerSpecHealthcheckArgs $;

        public Builder() {
            $ = new ServiceTaskSpecContainerSpecHealthcheckArgs();
        }

        public Builder(ServiceTaskSpecContainerSpecHealthcheckArgs defaults) {
            $ = new ServiceTaskSpecContainerSpecHealthcheckArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param interval Time between running the check (ms|s|m|h). Defaults to `0s`.
         * 
         * @return builder
         * 
         */
        public Builder interval(@Nullable Output<String> interval) {
            $.interval = interval;
            return this;
        }

        /**
         * @param interval Time between running the check (ms|s|m|h). Defaults to `0s`.
         * 
         * @return builder
         * 
         */
        public Builder interval(String interval) {
            return interval(Output.of(interval));
        }

        /**
         * @param retries Consecutive failures needed to report unhealthy. Defaults to `0`
         * 
         * @return builder
         * 
         */
        public Builder retries(@Nullable Output<Integer> retries) {
            $.retries = retries;
            return this;
        }

        /**
         * @param retries Consecutive failures needed to report unhealthy. Defaults to `0`
         * 
         * @return builder
         * 
         */
        public Builder retries(Integer retries) {
            return retries(Output.of(retries));
        }

        /**
         * @param startPeriod Start period for the container to initialize before counting retries towards unstable (ms|s|m|h). Defaults to `0s`.
         * 
         * @return builder
         * 
         */
        public Builder startPeriod(@Nullable Output<String> startPeriod) {
            $.startPeriod = startPeriod;
            return this;
        }

        /**
         * @param startPeriod Start period for the container to initialize before counting retries towards unstable (ms|s|m|h). Defaults to `0s`.
         * 
         * @return builder
         * 
         */
        public Builder startPeriod(String startPeriod) {
            return startPeriod(Output.of(startPeriod));
        }

        /**
         * @param tests The test to perform as list
         * 
         * @return builder
         * 
         */
        public Builder tests(Output<List<String>> tests) {
            $.tests = tests;
            return this;
        }

        /**
         * @param tests The test to perform as list
         * 
         * @return builder
         * 
         */
        public Builder tests(List<String> tests) {
            return tests(Output.of(tests));
        }

        /**
         * @param tests The test to perform as list
         * 
         * @return builder
         * 
         */
        public Builder tests(String... tests) {
            return tests(List.of(tests));
        }

        /**
         * @param timeout Maximum time to allow one check to run (ms|s|m|h). Defaults to `0s`.
         * 
         * @return builder
         * 
         */
        public Builder timeout(@Nullable Output<String> timeout) {
            $.timeout = timeout;
            return this;
        }

        /**
         * @param timeout Maximum time to allow one check to run (ms|s|m|h). Defaults to `0s`.
         * 
         * @return builder
         * 
         */
        public Builder timeout(String timeout) {
            return timeout(Output.of(timeout));
        }

        public ServiceTaskSpecContainerSpecHealthcheckArgs build() {
            if ($.tests == null) {
                throw new MissingRequiredPropertyException("ServiceTaskSpecContainerSpecHealthcheckArgs", "tests");
            }
            return $;
        }
    }

}
