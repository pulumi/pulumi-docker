// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class ServiceTaskSpecContainerSpecLabel {
    private String label;
    private String value;

    private ServiceTaskSpecContainerSpecLabel() {}
    public String label() {
        return this.label;
    }
    public String value() {
        return this.value;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ServiceTaskSpecContainerSpecLabel defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String label;
        private String value;
        public Builder() {}
        public Builder(ServiceTaskSpecContainerSpecLabel defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.label = defaults.label;
    	      this.value = defaults.value;
        }

        @CustomType.Setter
        public Builder label(String label) {
            if (label == null) {
              throw new MissingRequiredPropertyException("ServiceTaskSpecContainerSpecLabel", "label");
            }
            this.label = label;
            return this;
        }
        @CustomType.Setter
        public Builder value(String value) {
            if (value == null) {
              throw new MissingRequiredPropertyException("ServiceTaskSpecContainerSpecLabel", "value");
            }
            this.value = value;
            return this;
        }
        public ServiceTaskSpecContainerSpecLabel build() {
            final var _resultValue = new ServiceTaskSpecContainerSpecLabel();
            _resultValue.label = label;
            _resultValue.value = value;
            return _resultValue;
        }
    }
}
