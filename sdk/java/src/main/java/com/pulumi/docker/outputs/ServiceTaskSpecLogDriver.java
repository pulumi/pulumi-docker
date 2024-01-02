// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class ServiceTaskSpecLogDriver {
    /**
     * @return A random name for the port
     * 
     */
    private String name;
    /**
     * @return A list of internal resolver variables to be modified (e.g., `debug`, `ndots:3`, etc.)
     * 
     */
    private @Nullable Map<String,String> options;

    private ServiceTaskSpecLogDriver() {}
    /**
     * @return A random name for the port
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return A list of internal resolver variables to be modified (e.g., `debug`, `ndots:3`, etc.)
     * 
     */
    public Map<String,String> options() {
        return this.options == null ? Map.of() : this.options;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ServiceTaskSpecLogDriver defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String name;
        private @Nullable Map<String,String> options;
        public Builder() {}
        public Builder(ServiceTaskSpecLogDriver defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.name = defaults.name;
    	      this.options = defaults.options;
        }

        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("ServiceTaskSpecLogDriver", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder options(@Nullable Map<String,String> options) {

            this.options = options;
            return this;
        }
        public ServiceTaskSpecLogDriver build() {
            final var _resultValue = new ServiceTaskSpecLogDriver();
            _resultValue.name = name;
            _resultValue.options = options;
            return _resultValue;
        }
    }
}
