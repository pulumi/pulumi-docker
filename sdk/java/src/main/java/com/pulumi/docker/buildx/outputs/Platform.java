// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker.buildx.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class Platform {
    private String architecture;
    private String os;

    private Platform() {}
    public String architecture() {
        return this.architecture;
    }
    public String os() {
        return this.os;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(Platform defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String architecture;
        private String os;
        public Builder() {}
        public Builder(Platform defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.architecture = defaults.architecture;
    	      this.os = defaults.os;
        }

        @CustomType.Setter
        public Builder architecture(String architecture) {
            if (architecture == null) {
              throw new MissingRequiredPropertyException("Platform", "architecture");
            }
            this.architecture = architecture;
            return this;
        }
        @CustomType.Setter
        public Builder os(String os) {
            if (os == null) {
              throw new MissingRequiredPropertyException("Platform", "os");
            }
            this.os = os;
            return this;
        }
        public Platform build() {
            final var _resultValue = new Platform();
            _resultValue.architecture = architecture;
            _resultValue.os = os;
            return _resultValue;
        }
    }
}