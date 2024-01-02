// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class RemoteImageBuildAuthConfig {
    /**
     * @return the auth token
     * 
     */
    private @Nullable String auth;
    /**
     * @return the user emal
     * 
     */
    private @Nullable String email;
    /**
     * @return hostname of the registry
     * 
     */
    private String hostName;
    /**
     * @return the identity token
     * 
     */
    private @Nullable String identityToken;
    /**
     * @return the registry password
     * 
     */
    private @Nullable String password;
    /**
     * @return the registry token
     * 
     */
    private @Nullable String registryToken;
    /**
     * @return the server address
     * 
     */
    private @Nullable String serverAddress;
    /**
     * @return the registry user name
     * 
     */
    private @Nullable String userName;

    private RemoteImageBuildAuthConfig() {}
    /**
     * @return the auth token
     * 
     */
    public Optional<String> auth() {
        return Optional.ofNullable(this.auth);
    }
    /**
     * @return the user emal
     * 
     */
    public Optional<String> email() {
        return Optional.ofNullable(this.email);
    }
    /**
     * @return hostname of the registry
     * 
     */
    public String hostName() {
        return this.hostName;
    }
    /**
     * @return the identity token
     * 
     */
    public Optional<String> identityToken() {
        return Optional.ofNullable(this.identityToken);
    }
    /**
     * @return the registry password
     * 
     */
    public Optional<String> password() {
        return Optional.ofNullable(this.password);
    }
    /**
     * @return the registry token
     * 
     */
    public Optional<String> registryToken() {
        return Optional.ofNullable(this.registryToken);
    }
    /**
     * @return the server address
     * 
     */
    public Optional<String> serverAddress() {
        return Optional.ofNullable(this.serverAddress);
    }
    /**
     * @return the registry user name
     * 
     */
    public Optional<String> userName() {
        return Optional.ofNullable(this.userName);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(RemoteImageBuildAuthConfig defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String auth;
        private @Nullable String email;
        private String hostName;
        private @Nullable String identityToken;
        private @Nullable String password;
        private @Nullable String registryToken;
        private @Nullable String serverAddress;
        private @Nullable String userName;
        public Builder() {}
        public Builder(RemoteImageBuildAuthConfig defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.auth = defaults.auth;
    	      this.email = defaults.email;
    	      this.hostName = defaults.hostName;
    	      this.identityToken = defaults.identityToken;
    	      this.password = defaults.password;
    	      this.registryToken = defaults.registryToken;
    	      this.serverAddress = defaults.serverAddress;
    	      this.userName = defaults.userName;
        }

        @CustomType.Setter
        public Builder auth(@Nullable String auth) {

            this.auth = auth;
            return this;
        }
        @CustomType.Setter
        public Builder email(@Nullable String email) {

            this.email = email;
            return this;
        }
        @CustomType.Setter
        public Builder hostName(String hostName) {
            if (hostName == null) {
              throw new MissingRequiredPropertyException("RemoteImageBuildAuthConfig", "hostName");
            }
            this.hostName = hostName;
            return this;
        }
        @CustomType.Setter
        public Builder identityToken(@Nullable String identityToken) {

            this.identityToken = identityToken;
            return this;
        }
        @CustomType.Setter
        public Builder password(@Nullable String password) {

            this.password = password;
            return this;
        }
        @CustomType.Setter
        public Builder registryToken(@Nullable String registryToken) {

            this.registryToken = registryToken;
            return this;
        }
        @CustomType.Setter
        public Builder serverAddress(@Nullable String serverAddress) {

            this.serverAddress = serverAddress;
            return this;
        }
        @CustomType.Setter
        public Builder userName(@Nullable String userName) {

            this.userName = userName;
            return this;
        }
        public RemoteImageBuildAuthConfig build() {
            final var _resultValue = new RemoteImageBuildAuthConfig();
            _resultValue.auth = auth;
            _resultValue.email = email;
            _resultValue.hostName = hostName;
            _resultValue.identityToken = identityToken;
            _resultValue.password = password;
            _resultValue.registryToken = registryToken;
            _resultValue.serverAddress = serverAddress;
            _resultValue.userName = userName;
            return _resultValue;
        }
    }
}
