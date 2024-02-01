// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceUpdateConfigArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceUpdateConfigArgs Empty = new ServiceUpdateConfigArgs();

    /**
     * Delay between task updates `(ns|us|ms|s|m|h)`. Defaults to `0s`.
     * 
     */
    @Import(name="delay")
    private @Nullable Output<String> delay;

    /**
     * @return Delay between task updates `(ns|us|ms|s|m|h)`. Defaults to `0s`.
     * 
     */
    public Optional<Output<String>> delay() {
        return Optional.ofNullable(this.delay);
    }

    /**
     * Action on update failure: `pause`, `continue` or `rollback`. Defaults to `pause`.
     * 
     */
    @Import(name="failureAction")
    private @Nullable Output<String> failureAction;

    /**
     * @return Action on update failure: `pause`, `continue` or `rollback`. Defaults to `pause`.
     * 
     */
    public Optional<Output<String>> failureAction() {
        return Optional.ofNullable(this.failureAction);
    }

    /**
     * Failure rate to tolerate during an update. Defaults to `0.0`.
     * 
     */
    @Import(name="maxFailureRatio")
    private @Nullable Output<String> maxFailureRatio;

    /**
     * @return Failure rate to tolerate during an update. Defaults to `0.0`.
     * 
     */
    public Optional<Output<String>> maxFailureRatio() {
        return Optional.ofNullable(this.maxFailureRatio);
    }

    /**
     * Duration after each task update to monitor for failure (ns|us|ms|s|m|h). Defaults to `5s`.
     * 
     */
    @Import(name="monitor")
    private @Nullable Output<String> monitor;

    /**
     * @return Duration after each task update to monitor for failure (ns|us|ms|s|m|h). Defaults to `5s`.
     * 
     */
    public Optional<Output<String>> monitor() {
        return Optional.ofNullable(this.monitor);
    }

    /**
     * Update order: either &#39;stop-first&#39; or &#39;start-first&#39;. Defaults to `stop-first`.
     * 
     */
    @Import(name="order")
    private @Nullable Output<String> order;

    /**
     * @return Update order: either &#39;stop-first&#39; or &#39;start-first&#39;. Defaults to `stop-first`.
     * 
     */
    public Optional<Output<String>> order() {
        return Optional.ofNullable(this.order);
    }

    /**
     * Maximum number of tasks to be updated in one iteration. Defaults to `1`
     * 
     */
    @Import(name="parallelism")
    private @Nullable Output<Integer> parallelism;

    /**
     * @return Maximum number of tasks to be updated in one iteration. Defaults to `1`
     * 
     */
    public Optional<Output<Integer>> parallelism() {
        return Optional.ofNullable(this.parallelism);
    }

    private ServiceUpdateConfigArgs() {}

    private ServiceUpdateConfigArgs(ServiceUpdateConfigArgs $) {
        this.delay = $.delay;
        this.failureAction = $.failureAction;
        this.maxFailureRatio = $.maxFailureRatio;
        this.monitor = $.monitor;
        this.order = $.order;
        this.parallelism = $.parallelism;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceUpdateConfigArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceUpdateConfigArgs $;

        public Builder() {
            $ = new ServiceUpdateConfigArgs();
        }

        public Builder(ServiceUpdateConfigArgs defaults) {
            $ = new ServiceUpdateConfigArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param delay Delay between task updates `(ns|us|ms|s|m|h)`. Defaults to `0s`.
         * 
         * @return builder
         * 
         */
        public Builder delay(@Nullable Output<String> delay) {
            $.delay = delay;
            return this;
        }

        /**
         * @param delay Delay between task updates `(ns|us|ms|s|m|h)`. Defaults to `0s`.
         * 
         * @return builder
         * 
         */
        public Builder delay(String delay) {
            return delay(Output.of(delay));
        }

        /**
         * @param failureAction Action on update failure: `pause`, `continue` or `rollback`. Defaults to `pause`.
         * 
         * @return builder
         * 
         */
        public Builder failureAction(@Nullable Output<String> failureAction) {
            $.failureAction = failureAction;
            return this;
        }

        /**
         * @param failureAction Action on update failure: `pause`, `continue` or `rollback`. Defaults to `pause`.
         * 
         * @return builder
         * 
         */
        public Builder failureAction(String failureAction) {
            return failureAction(Output.of(failureAction));
        }

        /**
         * @param maxFailureRatio Failure rate to tolerate during an update. Defaults to `0.0`.
         * 
         * @return builder
         * 
         */
        public Builder maxFailureRatio(@Nullable Output<String> maxFailureRatio) {
            $.maxFailureRatio = maxFailureRatio;
            return this;
        }

        /**
         * @param maxFailureRatio Failure rate to tolerate during an update. Defaults to `0.0`.
         * 
         * @return builder
         * 
         */
        public Builder maxFailureRatio(String maxFailureRatio) {
            return maxFailureRatio(Output.of(maxFailureRatio));
        }

        /**
         * @param monitor Duration after each task update to monitor for failure (ns|us|ms|s|m|h). Defaults to `5s`.
         * 
         * @return builder
         * 
         */
        public Builder monitor(@Nullable Output<String> monitor) {
            $.monitor = monitor;
            return this;
        }

        /**
         * @param monitor Duration after each task update to monitor for failure (ns|us|ms|s|m|h). Defaults to `5s`.
         * 
         * @return builder
         * 
         */
        public Builder monitor(String monitor) {
            return monitor(Output.of(monitor));
        }

        /**
         * @param order Update order: either &#39;stop-first&#39; or &#39;start-first&#39;. Defaults to `stop-first`.
         * 
         * @return builder
         * 
         */
        public Builder order(@Nullable Output<String> order) {
            $.order = order;
            return this;
        }

        /**
         * @param order Update order: either &#39;stop-first&#39; or &#39;start-first&#39;. Defaults to `stop-first`.
         * 
         * @return builder
         * 
         */
        public Builder order(String order) {
            return order(Output.of(order));
        }

        /**
         * @param parallelism Maximum number of tasks to be updated in one iteration. Defaults to `1`
         * 
         * @return builder
         * 
         */
        public Builder parallelism(@Nullable Output<Integer> parallelism) {
            $.parallelism = parallelism;
            return this;
        }

        /**
         * @param parallelism Maximum number of tasks to be updated in one iteration. Defaults to `1`
         * 
         * @return builder
         * 
         */
        public Builder parallelism(Integer parallelism) {
            return parallelism(Output.of(parallelism));
        }

        public ServiceUpdateConfigArgs build() {
            return $;
        }
    }

}
