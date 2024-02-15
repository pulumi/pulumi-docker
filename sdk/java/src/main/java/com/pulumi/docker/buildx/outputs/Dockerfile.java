// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker.buildx.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class Dockerfile {
    /**
     * @return Raw Dockerfile contents.
     * 
     * Conflicts with `location`.
     * 
     */
    private @Nullable String inline;
    /**
     * @return Location of the Dockerfile to use.
     * 
     * Can be a relative or absolute path to a local file, or a remote URL.
     * 
     * Conflicts with `inline`.
     * 
     */
    private @Nullable String location;

    private Dockerfile() {}
    /**
     * @return Raw Dockerfile contents.
     * 
     * Conflicts with `location`.
     * 
     */
    public Optional<String> inline() {
        return Optional.ofNullable(this.inline);
    }
    /**
     * @return Location of the Dockerfile to use.
     * 
     * Can be a relative or absolute path to a local file, or a remote URL.
     * 
     * Conflicts with `inline`.
     * 
     */
    public Optional<String> location() {
        return Optional.ofNullable(this.location);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(Dockerfile defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String inline;
        private @Nullable String location;
        public Builder() {}
        public Builder(Dockerfile defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.inline = defaults.inline;
    	      this.location = defaults.location;
        }

        @CustomType.Setter
        public Builder inline(@Nullable String inline) {

            this.inline = inline;
            return this;
        }
        @CustomType.Setter
        public Builder location(@Nullable String location) {

            this.location = location;
            return this;
        }
        public Dockerfile build() {
            final var _resultValue = new Dockerfile();
            _resultValue.inline = inline;
            _resultValue.location = location;
            return _resultValue;
        }
    }
}
