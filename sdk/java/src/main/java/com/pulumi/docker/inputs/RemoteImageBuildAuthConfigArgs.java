// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RemoteImageBuildAuthConfigArgs extends com.pulumi.resources.ResourceArgs {

    public static final RemoteImageBuildAuthConfigArgs Empty = new RemoteImageBuildAuthConfigArgs();

    /**
     * the auth token
     * 
     */
    @Import(name="auth")
    private @Nullable Output<String> auth;

    /**
     * @return the auth token
     * 
     */
    public Optional<Output<String>> auth() {
        return Optional.ofNullable(this.auth);
    }

    /**
     * the user emal
     * 
     */
    @Import(name="email")
    private @Nullable Output<String> email;

    /**
     * @return the user emal
     * 
     */
    public Optional<Output<String>> email() {
        return Optional.ofNullable(this.email);
    }

    /**
     * hostname of the registry
     * 
     */
    @Import(name="hostName", required=true)
    private Output<String> hostName;

    /**
     * @return hostname of the registry
     * 
     */
    public Output<String> hostName() {
        return this.hostName;
    }

    /**
     * the identity token
     * 
     */
    @Import(name="identityToken")
    private @Nullable Output<String> identityToken;

    /**
     * @return the identity token
     * 
     */
    public Optional<Output<String>> identityToken() {
        return Optional.ofNullable(this.identityToken);
    }

    /**
     * the registry password
     * 
     */
    @Import(name="password")
    private @Nullable Output<String> password;

    /**
     * @return the registry password
     * 
     */
    public Optional<Output<String>> password() {
        return Optional.ofNullable(this.password);
    }

    /**
     * the registry token
     * 
     */
    @Import(name="registryToken")
    private @Nullable Output<String> registryToken;

    /**
     * @return the registry token
     * 
     */
    public Optional<Output<String>> registryToken() {
        return Optional.ofNullable(this.registryToken);
    }

    /**
     * the server address
     * 
     */
    @Import(name="serverAddress")
    private @Nullable Output<String> serverAddress;

    /**
     * @return the server address
     * 
     */
    public Optional<Output<String>> serverAddress() {
        return Optional.ofNullable(this.serverAddress);
    }

    /**
     * the registry user name
     * 
     */
    @Import(name="userName")
    private @Nullable Output<String> userName;

    /**
     * @return the registry user name
     * 
     */
    public Optional<Output<String>> userName() {
        return Optional.ofNullable(this.userName);
    }

    private RemoteImageBuildAuthConfigArgs() {}

    private RemoteImageBuildAuthConfigArgs(RemoteImageBuildAuthConfigArgs $) {
        this.auth = $.auth;
        this.email = $.email;
        this.hostName = $.hostName;
        this.identityToken = $.identityToken;
        this.password = $.password;
        this.registryToken = $.registryToken;
        this.serverAddress = $.serverAddress;
        this.userName = $.userName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RemoteImageBuildAuthConfigArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RemoteImageBuildAuthConfigArgs $;

        public Builder() {
            $ = new RemoteImageBuildAuthConfigArgs();
        }

        public Builder(RemoteImageBuildAuthConfigArgs defaults) {
            $ = new RemoteImageBuildAuthConfigArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param auth the auth token
         * 
         * @return builder
         * 
         */
        public Builder auth(@Nullable Output<String> auth) {
            $.auth = auth;
            return this;
        }

        /**
         * @param auth the auth token
         * 
         * @return builder
         * 
         */
        public Builder auth(String auth) {
            return auth(Output.of(auth));
        }

        /**
         * @param email the user emal
         * 
         * @return builder
         * 
         */
        public Builder email(@Nullable Output<String> email) {
            $.email = email;
            return this;
        }

        /**
         * @param email the user emal
         * 
         * @return builder
         * 
         */
        public Builder email(String email) {
            return email(Output.of(email));
        }

        /**
         * @param hostName hostname of the registry
         * 
         * @return builder
         * 
         */
        public Builder hostName(Output<String> hostName) {
            $.hostName = hostName;
            return this;
        }

        /**
         * @param hostName hostname of the registry
         * 
         * @return builder
         * 
         */
        public Builder hostName(String hostName) {
            return hostName(Output.of(hostName));
        }

        /**
         * @param identityToken the identity token
         * 
         * @return builder
         * 
         */
        public Builder identityToken(@Nullable Output<String> identityToken) {
            $.identityToken = identityToken;
            return this;
        }

        /**
         * @param identityToken the identity token
         * 
         * @return builder
         * 
         */
        public Builder identityToken(String identityToken) {
            return identityToken(Output.of(identityToken));
        }

        /**
         * @param password the registry password
         * 
         * @return builder
         * 
         */
        public Builder password(@Nullable Output<String> password) {
            $.password = password;
            return this;
        }

        /**
         * @param password the registry password
         * 
         * @return builder
         * 
         */
        public Builder password(String password) {
            return password(Output.of(password));
        }

        /**
         * @param registryToken the registry token
         * 
         * @return builder
         * 
         */
        public Builder registryToken(@Nullable Output<String> registryToken) {
            $.registryToken = registryToken;
            return this;
        }

        /**
         * @param registryToken the registry token
         * 
         * @return builder
         * 
         */
        public Builder registryToken(String registryToken) {
            return registryToken(Output.of(registryToken));
        }

        /**
         * @param serverAddress the server address
         * 
         * @return builder
         * 
         */
        public Builder serverAddress(@Nullable Output<String> serverAddress) {
            $.serverAddress = serverAddress;
            return this;
        }

        /**
         * @param serverAddress the server address
         * 
         * @return builder
         * 
         */
        public Builder serverAddress(String serverAddress) {
            return serverAddress(Output.of(serverAddress));
        }

        /**
         * @param userName the registry user name
         * 
         * @return builder
         * 
         */
        public Builder userName(@Nullable Output<String> userName) {
            $.userName = userName;
            return this;
        }

        /**
         * @param userName the registry user name
         * 
         * @return builder
         * 
         */
        public Builder userName(String userName) {
            return userName(Output.of(userName));
        }

        public RemoteImageBuildAuthConfigArgs build() {
            $.hostName = Objects.requireNonNull($.hostName, "expected parameter 'hostName' to be non-null");
            return $;
        }
    }

}
