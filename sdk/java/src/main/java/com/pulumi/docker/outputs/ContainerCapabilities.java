// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class ContainerCapabilities {
    /**
     * @return List of linux capabilities to add.
     * 
     */
    private @Nullable List<String> adds;
    /**
     * @return List of linux capabilities to drop.
     * 
     */
    private @Nullable List<String> drops;

    private ContainerCapabilities() {}
    /**
     * @return List of linux capabilities to add.
     * 
     */
    public List<String> adds() {
        return this.adds == null ? List.of() : this.adds;
    }
    /**
     * @return List of linux capabilities to drop.
     * 
     */
    public List<String> drops() {
        return this.drops == null ? List.of() : this.drops;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ContainerCapabilities defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<String> adds;
        private @Nullable List<String> drops;
        public Builder() {}
        public Builder(ContainerCapabilities defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.adds = defaults.adds;
    	      this.drops = defaults.drops;
        }

        @CustomType.Setter
        public Builder adds(@Nullable List<String> adds) {
            this.adds = adds;
            return this;
        }
        public Builder adds(String... adds) {
            return adds(List.of(adds));
        }
        @CustomType.Setter
        public Builder drops(@Nullable List<String> drops) {
            this.drops = drops;
            return this;
        }
        public Builder drops(String... drops) {
            return drops(List.of(drops));
        }
        public ContainerCapabilities build() {
            final var _resultValue = new ContainerCapabilities();
            _resultValue.adds = adds;
            _resultValue.drops = drops;
            return _resultValue;
        }
    }
}
