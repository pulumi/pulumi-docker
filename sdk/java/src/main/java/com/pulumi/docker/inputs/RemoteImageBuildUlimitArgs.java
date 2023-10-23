// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;


public final class RemoteImageBuildUlimitArgs extends com.pulumi.resources.ResourceArgs {

    public static final RemoteImageBuildUlimitArgs Empty = new RemoteImageBuildUlimitArgs();

    /**
     * soft limit
     * 
     */
    @Import(name="hard", required=true)
    private Output<Integer> hard;

    /**
     * @return soft limit
     * 
     */
    public Output<Integer> hard() {
        return this.hard;
    }

    /**
     * type of ulimit, e.g. `nofile`
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return type of ulimit, e.g. `nofile`
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     * hard limit
     * 
     */
    @Import(name="soft", required=true)
    private Output<Integer> soft;

    /**
     * @return hard limit
     * 
     */
    public Output<Integer> soft() {
        return this.soft;
    }

    private RemoteImageBuildUlimitArgs() {}

    private RemoteImageBuildUlimitArgs(RemoteImageBuildUlimitArgs $) {
        this.hard = $.hard;
        this.name = $.name;
        this.soft = $.soft;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RemoteImageBuildUlimitArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RemoteImageBuildUlimitArgs $;

        public Builder() {
            $ = new RemoteImageBuildUlimitArgs();
        }

        public Builder(RemoteImageBuildUlimitArgs defaults) {
            $ = new RemoteImageBuildUlimitArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param hard soft limit
         * 
         * @return builder
         * 
         */
        public Builder hard(Output<Integer> hard) {
            $.hard = hard;
            return this;
        }

        /**
         * @param hard soft limit
         * 
         * @return builder
         * 
         */
        public Builder hard(Integer hard) {
            return hard(Output.of(hard));
        }

        /**
         * @param name type of ulimit, e.g. `nofile`
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name type of ulimit, e.g. `nofile`
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param soft hard limit
         * 
         * @return builder
         * 
         */
        public Builder soft(Output<Integer> soft) {
            $.soft = soft;
            return this;
        }

        /**
         * @param soft hard limit
         * 
         * @return builder
         * 
         */
        public Builder soft(Integer soft) {
            return soft(Output.of(soft));
        }

        public RemoteImageBuildUlimitArgs build() {
            $.hard = Objects.requireNonNull($.hard, "expected parameter 'hard' to be non-null");
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            $.soft = Objects.requireNonNull($.soft, "expected parameter 'soft' to be non-null");
            return $;
        }
    }

}
