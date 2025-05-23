// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class RegistryImageAuthConfig {
    /**
     * @return The address of the Docker registry.
     * 
     */
    private String address;
    /**
     * @return The password for the Docker registry.
     * 
     */
    private String password;
    /**
     * @return The username for the Docker registry.
     * 
     */
    private String username;

    private RegistryImageAuthConfig() {}
    /**
     * @return The address of the Docker registry.
     * 
     */
    public String address() {
        return this.address;
    }
    /**
     * @return The password for the Docker registry.
     * 
     */
    public String password() {
        return this.password;
    }
    /**
     * @return The username for the Docker registry.
     * 
     */
    public String username() {
        return this.username;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(RegistryImageAuthConfig defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String address;
        private String password;
        private String username;
        public Builder() {}
        public Builder(RegistryImageAuthConfig defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.address = defaults.address;
    	      this.password = defaults.password;
    	      this.username = defaults.username;
        }

        @CustomType.Setter
        public Builder address(String address) {
            if (address == null) {
              throw new MissingRequiredPropertyException("RegistryImageAuthConfig", "address");
            }
            this.address = address;
            return this;
        }
        @CustomType.Setter
        public Builder password(String password) {
            if (password == null) {
              throw new MissingRequiredPropertyException("RegistryImageAuthConfig", "password");
            }
            this.password = password;
            return this;
        }
        @CustomType.Setter
        public Builder username(String username) {
            if (username == null) {
              throw new MissingRequiredPropertyException("RegistryImageAuthConfig", "username");
            }
            this.username = username;
            return this;
        }
        public RegistryImageAuthConfig build() {
            final var _resultValue = new RegistryImageAuthConfig();
            _resultValue.address = address;
            _resultValue.password = password;
            _resultValue.username = username;
            return _resultValue;
        }
    }
}
