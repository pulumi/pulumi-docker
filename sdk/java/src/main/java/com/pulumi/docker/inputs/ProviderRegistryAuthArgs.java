// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ProviderRegistryAuthArgs extends com.pulumi.resources.ResourceArgs {

    public static final ProviderRegistryAuthArgs Empty = new ProviderRegistryAuthArgs();

    /**
     * Address of the registry
     * 
     */
    @Import(name="address", required=true)
    private Output<String> address;

    /**
     * @return Address of the registry
     * 
     */
    public Output<String> address() {
        return this.address;
    }

    @Import(name="authDisabled")
    private @Nullable Output<Boolean> authDisabled;

    public Optional<Output<Boolean>> authDisabled() {
        return Optional.ofNullable(this.authDisabled);
    }

    /**
     * Path to docker json file for registry auth. Defaults to `~/.docker/config.json`. If `DOCKER_CONFIG` is set, the value of `DOCKER_CONFIG` is used as the path. `config_file` has predencen over all other options.
     * 
     */
    @Import(name="configFile")
    private @Nullable Output<String> configFile;

    /**
     * @return Path to docker json file for registry auth. Defaults to `~/.docker/config.json`. If `DOCKER_CONFIG` is set, the value of `DOCKER_CONFIG` is used as the path. `config_file` has predencen over all other options.
     * 
     */
    public Optional<Output<String>> configFile() {
        return Optional.ofNullable(this.configFile);
    }

    /**
     * Plain content of the docker json file for registry auth. `config_file_content` has precedence over username/password.
     * 
     */
    @Import(name="configFileContent")
    private @Nullable Output<String> configFileContent;

    /**
     * @return Plain content of the docker json file for registry auth. `config_file_content` has precedence over username/password.
     * 
     */
    public Optional<Output<String>> configFileContent() {
        return Optional.ofNullable(this.configFileContent);
    }

    /**
     * Password for the registry. Defaults to `DOCKER_REGISTRY_PASS` env variable if set.
     * 
     */
    @Import(name="password")
    private @Nullable Output<String> password;

    /**
     * @return Password for the registry. Defaults to `DOCKER_REGISTRY_PASS` env variable if set.
     * 
     */
    public Optional<Output<String>> password() {
        return Optional.ofNullable(this.password);
    }

    /**
     * Username for the registry. Defaults to `DOCKER_REGISTRY_USER` env variable if set.
     * 
     */
    @Import(name="username")
    private @Nullable Output<String> username;

    /**
     * @return Username for the registry. Defaults to `DOCKER_REGISTRY_USER` env variable if set.
     * 
     */
    public Optional<Output<String>> username() {
        return Optional.ofNullable(this.username);
    }

    private ProviderRegistryAuthArgs() {}

    private ProviderRegistryAuthArgs(ProviderRegistryAuthArgs $) {
        this.address = $.address;
        this.authDisabled = $.authDisabled;
        this.configFile = $.configFile;
        this.configFileContent = $.configFileContent;
        this.password = $.password;
        this.username = $.username;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProviderRegistryAuthArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProviderRegistryAuthArgs $;

        public Builder() {
            $ = new ProviderRegistryAuthArgs();
        }

        public Builder(ProviderRegistryAuthArgs defaults) {
            $ = new ProviderRegistryAuthArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param address Address of the registry
         * 
         * @return builder
         * 
         */
        public Builder address(Output<String> address) {
            $.address = address;
            return this;
        }

        /**
         * @param address Address of the registry
         * 
         * @return builder
         * 
         */
        public Builder address(String address) {
            return address(Output.of(address));
        }

        public Builder authDisabled(@Nullable Output<Boolean> authDisabled) {
            $.authDisabled = authDisabled;
            return this;
        }

        public Builder authDisabled(Boolean authDisabled) {
            return authDisabled(Output.of(authDisabled));
        }

        /**
         * @param configFile Path to docker json file for registry auth. Defaults to `~/.docker/config.json`. If `DOCKER_CONFIG` is set, the value of `DOCKER_CONFIG` is used as the path. `config_file` has predencen over all other options.
         * 
         * @return builder
         * 
         */
        public Builder configFile(@Nullable Output<String> configFile) {
            $.configFile = configFile;
            return this;
        }

        /**
         * @param configFile Path to docker json file for registry auth. Defaults to `~/.docker/config.json`. If `DOCKER_CONFIG` is set, the value of `DOCKER_CONFIG` is used as the path. `config_file` has predencen over all other options.
         * 
         * @return builder
         * 
         */
        public Builder configFile(String configFile) {
            return configFile(Output.of(configFile));
        }

        /**
         * @param configFileContent Plain content of the docker json file for registry auth. `config_file_content` has precedence over username/password.
         * 
         * @return builder
         * 
         */
        public Builder configFileContent(@Nullable Output<String> configFileContent) {
            $.configFileContent = configFileContent;
            return this;
        }

        /**
         * @param configFileContent Plain content of the docker json file for registry auth. `config_file_content` has precedence over username/password.
         * 
         * @return builder
         * 
         */
        public Builder configFileContent(String configFileContent) {
            return configFileContent(Output.of(configFileContent));
        }

        /**
         * @param password Password for the registry. Defaults to `DOCKER_REGISTRY_PASS` env variable if set.
         * 
         * @return builder
         * 
         */
        public Builder password(@Nullable Output<String> password) {
            $.password = password;
            return this;
        }

        /**
         * @param password Password for the registry. Defaults to `DOCKER_REGISTRY_PASS` env variable if set.
         * 
         * @return builder
         * 
         */
        public Builder password(String password) {
            return password(Output.of(password));
        }

        /**
         * @param username Username for the registry. Defaults to `DOCKER_REGISTRY_USER` env variable if set.
         * 
         * @return builder
         * 
         */
        public Builder username(@Nullable Output<String> username) {
            $.username = username;
            return this;
        }

        /**
         * @param username Username for the registry. Defaults to `DOCKER_REGISTRY_USER` env variable if set.
         * 
         * @return builder
         * 
         */
        public Builder username(String username) {
            return username(Output.of(username));
        }

        public ProviderRegistryAuthArgs build() {
            if ($.address == null) {
                throw new MissingRequiredPropertyException("ProviderRegistryAuthArgs", "address");
            }
            return $;
        }
    }

}
