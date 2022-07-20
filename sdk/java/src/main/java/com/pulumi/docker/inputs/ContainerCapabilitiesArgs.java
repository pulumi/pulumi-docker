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


public final class ContainerCapabilitiesArgs extends com.pulumi.resources.ResourceArgs {

    public static final ContainerCapabilitiesArgs Empty = new ContainerCapabilitiesArgs();

    @Import(name="adds")
    private @Nullable Output<List<String>> adds;

    public Optional<Output<List<String>>> adds() {
        return Optional.ofNullable(this.adds);
    }

    @Import(name="drops")
    private @Nullable Output<List<String>> drops;

    public Optional<Output<List<String>>> drops() {
        return Optional.ofNullable(this.drops);
    }

    private ContainerCapabilitiesArgs() {}

    private ContainerCapabilitiesArgs(ContainerCapabilitiesArgs $) {
        this.adds = $.adds;
        this.drops = $.drops;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ContainerCapabilitiesArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ContainerCapabilitiesArgs $;

        public Builder() {
            $ = new ContainerCapabilitiesArgs();
        }

        public Builder(ContainerCapabilitiesArgs defaults) {
            $ = new ContainerCapabilitiesArgs(Objects.requireNonNull(defaults));
        }

        public Builder adds(@Nullable Output<List<String>> adds) {
            $.adds = adds;
            return this;
        }

        public Builder adds(List<String> adds) {
            return adds(Output.of(adds));
        }

        public Builder adds(String... adds) {
            return adds(List.of(adds));
        }

        public Builder drops(@Nullable Output<List<String>> drops) {
            $.drops = drops;
            return this;
        }

        public Builder drops(List<String> drops) {
            return drops(Output.of(drops));
        }

        public Builder drops(String... drops) {
            return drops(List.of(drops));
        }

        public ContainerCapabilitiesArgs build() {
            return $;
        }
    }

}
