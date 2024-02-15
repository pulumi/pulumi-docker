// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker.buildx.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class RegistryAuth {
    /**
     * @return The registry&#39;s address (e.g. &#34;docker.io&#34;).
     * 
     */
    private String address;
    /**
     * @return Password or token for the registry.
     * 
     */
    private @Nullable String password;
    /**
     * @return Username for the registry.
     * 
     */
    private @Nullable String username;

    private RegistryAuth() {}
    /**
     * @return The registry&#39;s address (e.g. &#34;docker.io&#34;).
     * 
     */
    public String address() {
        return this.address;
    }
    /**
     * @return Password or token for the registry.
     * 
     */
    public Optional<String> password() {
        return Optional.ofNullable(this.password);
    }
    /**
     * @return Username for the registry.
     * 
     */
    public Optional<String> username() {
        return Optional.ofNullable(this.username);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(RegistryAuth defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String address;
        private @Nullable String password;
        private @Nullable String username;
        public Builder() {}
        public Builder(RegistryAuth defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.address = defaults.address;
    	      this.password = defaults.password;
    	      this.username = defaults.username;
        }

        @CustomType.Setter
        public Builder address(String address) {
            if (address == null) {
              throw new MissingRequiredPropertyException("RegistryAuth", "address");
            }
            this.address = address;
            return this;
        }
        @CustomType.Setter
        public Builder password(@Nullable String password) {

            this.password = password;
            return this;
        }
        @CustomType.Setter
        public Builder username(@Nullable String username) {

            this.username = username;
            return this;
        }
        public RegistryAuth build() {
            final var _resultValue = new RegistryAuth();
            _resultValue.address = address;
            _resultValue.password = password;
            _resultValue.username = username;
            return _resultValue;
        }
    }
}
