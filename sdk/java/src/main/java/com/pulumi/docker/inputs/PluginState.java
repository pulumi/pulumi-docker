// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.docker.inputs.PluginGrantPermissionArgs;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PluginState extends com.pulumi.resources.ResourceArgs {

    public static final PluginState Empty = new PluginState();

    /**
     * Docker Plugin alias
     * 
     */
    @Import(name="alias")
    private @Nullable Output<String> alias;

    /**
     * @return Docker Plugin alias
     * 
     */
    public Optional<Output<String>> alias() {
        return Optional.ofNullable(this.alias);
    }

    /**
     * HTTP client timeout to enable the plugin
     * 
     */
    @Import(name="enableTimeout")
    private @Nullable Output<Integer> enableTimeout;

    /**
     * @return HTTP client timeout to enable the plugin
     * 
     */
    public Optional<Output<Integer>> enableTimeout() {
        return Optional.ofNullable(this.enableTimeout);
    }

    /**
     * If `true` the plugin is enabled. Defaults to `true`
     * 
     */
    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    /**
     * @return If `true` the plugin is enabled. Defaults to `true`
     * 
     */
    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    /**
     * The environment variables in the form of `KEY=VALUE`, e.g. `DEBUG=0`
     * 
     */
    @Import(name="envs")
    private @Nullable Output<List<String>> envs;

    /**
     * @return The environment variables in the form of `KEY=VALUE`, e.g. `DEBUG=0`
     * 
     */
    public Optional<Output<List<String>>> envs() {
        return Optional.ofNullable(this.envs);
    }

    /**
     * If true, then the plugin is destroyed forcibly
     * 
     */
    @Import(name="forceDestroy")
    private @Nullable Output<Boolean> forceDestroy;

    /**
     * @return If true, then the plugin is destroyed forcibly
     * 
     */
    public Optional<Output<Boolean>> forceDestroy() {
        return Optional.ofNullable(this.forceDestroy);
    }

    /**
     * If true, then the plugin is disabled forcibly
     * 
     */
    @Import(name="forceDisable")
    private @Nullable Output<Boolean> forceDisable;

    /**
     * @return If true, then the plugin is disabled forcibly
     * 
     */
    public Optional<Output<Boolean>> forceDisable() {
        return Optional.ofNullable(this.forceDisable);
    }

    /**
     * If true, grant all permissions necessary to run the plugin
     * 
     */
    @Import(name="grantAllPermissions")
    private @Nullable Output<Boolean> grantAllPermissions;

    /**
     * @return If true, grant all permissions necessary to run the plugin
     * 
     */
    public Optional<Output<Boolean>> grantAllPermissions() {
        return Optional.ofNullable(this.grantAllPermissions);
    }

    /**
     * Grant specific permissions only
     * 
     */
    @Import(name="grantPermissions")
    private @Nullable Output<List<PluginGrantPermissionArgs>> grantPermissions;

    /**
     * @return Grant specific permissions only
     * 
     */
    public Optional<Output<List<PluginGrantPermissionArgs>>> grantPermissions() {
        return Optional.ofNullable(this.grantPermissions);
    }

    /**
     * Docker Plugin name
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Docker Plugin name
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Docker Plugin Reference
     * 
     */
    @Import(name="pluginReference")
    private @Nullable Output<String> pluginReference;

    /**
     * @return Docker Plugin Reference
     * 
     */
    public Optional<Output<String>> pluginReference() {
        return Optional.ofNullable(this.pluginReference);
    }

    private PluginState() {}

    private PluginState(PluginState $) {
        this.alias = $.alias;
        this.enableTimeout = $.enableTimeout;
        this.enabled = $.enabled;
        this.envs = $.envs;
        this.forceDestroy = $.forceDestroy;
        this.forceDisable = $.forceDisable;
        this.grantAllPermissions = $.grantAllPermissions;
        this.grantPermissions = $.grantPermissions;
        this.name = $.name;
        this.pluginReference = $.pluginReference;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PluginState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PluginState $;

        public Builder() {
            $ = new PluginState();
        }

        public Builder(PluginState defaults) {
            $ = new PluginState(Objects.requireNonNull(defaults));
        }

        /**
         * @param alias Docker Plugin alias
         * 
         * @return builder
         * 
         */
        public Builder alias(@Nullable Output<String> alias) {
            $.alias = alias;
            return this;
        }

        /**
         * @param alias Docker Plugin alias
         * 
         * @return builder
         * 
         */
        public Builder alias(String alias) {
            return alias(Output.of(alias));
        }

        /**
         * @param enableTimeout HTTP client timeout to enable the plugin
         * 
         * @return builder
         * 
         */
        public Builder enableTimeout(@Nullable Output<Integer> enableTimeout) {
            $.enableTimeout = enableTimeout;
            return this;
        }

        /**
         * @param enableTimeout HTTP client timeout to enable the plugin
         * 
         * @return builder
         * 
         */
        public Builder enableTimeout(Integer enableTimeout) {
            return enableTimeout(Output.of(enableTimeout));
        }

        /**
         * @param enabled If `true` the plugin is enabled. Defaults to `true`
         * 
         * @return builder
         * 
         */
        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled If `true` the plugin is enabled. Defaults to `true`
         * 
         * @return builder
         * 
         */
        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        /**
         * @param envs The environment variables in the form of `KEY=VALUE`, e.g. `DEBUG=0`
         * 
         * @return builder
         * 
         */
        public Builder envs(@Nullable Output<List<String>> envs) {
            $.envs = envs;
            return this;
        }

        /**
         * @param envs The environment variables in the form of `KEY=VALUE`, e.g. `DEBUG=0`
         * 
         * @return builder
         * 
         */
        public Builder envs(List<String> envs) {
            return envs(Output.of(envs));
        }

        /**
         * @param envs The environment variables in the form of `KEY=VALUE`, e.g. `DEBUG=0`
         * 
         * @return builder
         * 
         */
        public Builder envs(String... envs) {
            return envs(List.of(envs));
        }

        /**
         * @param forceDestroy If true, then the plugin is destroyed forcibly
         * 
         * @return builder
         * 
         */
        public Builder forceDestroy(@Nullable Output<Boolean> forceDestroy) {
            $.forceDestroy = forceDestroy;
            return this;
        }

        /**
         * @param forceDestroy If true, then the plugin is destroyed forcibly
         * 
         * @return builder
         * 
         */
        public Builder forceDestroy(Boolean forceDestroy) {
            return forceDestroy(Output.of(forceDestroy));
        }

        /**
         * @param forceDisable If true, then the plugin is disabled forcibly
         * 
         * @return builder
         * 
         */
        public Builder forceDisable(@Nullable Output<Boolean> forceDisable) {
            $.forceDisable = forceDisable;
            return this;
        }

        /**
         * @param forceDisable If true, then the plugin is disabled forcibly
         * 
         * @return builder
         * 
         */
        public Builder forceDisable(Boolean forceDisable) {
            return forceDisable(Output.of(forceDisable));
        }

        /**
         * @param grantAllPermissions If true, grant all permissions necessary to run the plugin
         * 
         * @return builder
         * 
         */
        public Builder grantAllPermissions(@Nullable Output<Boolean> grantAllPermissions) {
            $.grantAllPermissions = grantAllPermissions;
            return this;
        }

        /**
         * @param grantAllPermissions If true, grant all permissions necessary to run the plugin
         * 
         * @return builder
         * 
         */
        public Builder grantAllPermissions(Boolean grantAllPermissions) {
            return grantAllPermissions(Output.of(grantAllPermissions));
        }

        /**
         * @param grantPermissions Grant specific permissions only
         * 
         * @return builder
         * 
         */
        public Builder grantPermissions(@Nullable Output<List<PluginGrantPermissionArgs>> grantPermissions) {
            $.grantPermissions = grantPermissions;
            return this;
        }

        /**
         * @param grantPermissions Grant specific permissions only
         * 
         * @return builder
         * 
         */
        public Builder grantPermissions(List<PluginGrantPermissionArgs> grantPermissions) {
            return grantPermissions(Output.of(grantPermissions));
        }

        /**
         * @param grantPermissions Grant specific permissions only
         * 
         * @return builder
         * 
         */
        public Builder grantPermissions(PluginGrantPermissionArgs... grantPermissions) {
            return grantPermissions(List.of(grantPermissions));
        }

        /**
         * @param name Docker Plugin name
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Docker Plugin name
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param pluginReference Docker Plugin Reference
         * 
         * @return builder
         * 
         */
        public Builder pluginReference(@Nullable Output<String> pluginReference) {
            $.pluginReference = pluginReference;
            return this;
        }

        /**
         * @param pluginReference Docker Plugin Reference
         * 
         * @return builder
         * 
         */
        public Builder pluginReference(String pluginReference) {
            return pluginReference(Output.of(pluginReference));
        }

        public PluginState build() {
            return $;
        }
    }

}
