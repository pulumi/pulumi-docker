// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class ServiceLabel {
    /**
     * @return Name of the label
     * 
     */
    private String label;
    /**
     * @return Value of the label
     * 
     */
    private String value;

    private ServiceLabel() {}
    /**
     * @return Name of the label
     * 
     */
    public String label() {
        return this.label;
    }
    /**
     * @return Value of the label
     * 
     */
    public String value() {
        return this.value;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ServiceLabel defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String label;
        private String value;
        public Builder() {}
        public Builder(ServiceLabel defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.label = defaults.label;
    	      this.value = defaults.value;
        }

        @CustomType.Setter
        public Builder label(String label) {
            if (label == null) {
              throw new MissingRequiredPropertyException("ServiceLabel", "label");
            }
            this.label = label;
            return this;
        }
        @CustomType.Setter
        public Builder value(String value) {
            if (value == null) {
              throw new MissingRequiredPropertyException("ServiceLabel", "value");
            }
            this.value = value;
            return this;
        }
        public ServiceLabel build() {
            final var _resultValue = new ServiceLabel();
            _resultValue.label = label;
            _resultValue.value = value;
            return _resultValue;
        }
    }
}
