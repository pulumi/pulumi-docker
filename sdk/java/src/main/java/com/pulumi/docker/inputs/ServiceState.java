// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.docker.inputs.ServiceAuthArgs;
import com.pulumi.docker.inputs.ServiceConvergeConfigArgs;
import com.pulumi.docker.inputs.ServiceEndpointSpecArgs;
import com.pulumi.docker.inputs.ServiceLabelArgs;
import com.pulumi.docker.inputs.ServiceModeArgs;
import com.pulumi.docker.inputs.ServiceRollbackConfigArgs;
import com.pulumi.docker.inputs.ServiceTaskSpecArgs;
import com.pulumi.docker.inputs.ServiceUpdateConfigArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceState extends com.pulumi.resources.ResourceArgs {

    public static final ServiceState Empty = new ServiceState();

    /**
     * Configuration for the authentication for pulling the images of the service
     * 
     */
    @Import(name="auth")
    private @Nullable Output<ServiceAuthArgs> auth;

    /**
     * @return Configuration for the authentication for pulling the images of the service
     * 
     */
    public Optional<Output<ServiceAuthArgs>> auth() {
        return Optional.ofNullable(this.auth);
    }

    /**
     * A configuration to ensure that a service converges aka reaches the desired that of all task up and running
     * 
     */
    @Import(name="convergeConfig")
    private @Nullable Output<ServiceConvergeConfigArgs> convergeConfig;

    /**
     * @return A configuration to ensure that a service converges aka reaches the desired that of all task up and running
     * 
     */
    public Optional<Output<ServiceConvergeConfigArgs>> convergeConfig() {
        return Optional.ofNullable(this.convergeConfig);
    }

    /**
     * Properties that can be configured to access and load balance a service
     * 
     */
    @Import(name="endpointSpec")
    private @Nullable Output<ServiceEndpointSpecArgs> endpointSpec;

    /**
     * @return Properties that can be configured to access and load balance a service
     * 
     */
    public Optional<Output<ServiceEndpointSpecArgs>> endpointSpec() {
        return Optional.ofNullable(this.endpointSpec);
    }

    /**
     * User-defined key/value metadata
     * 
     */
    @Import(name="labels")
    private @Nullable Output<List<ServiceLabelArgs>> labels;

    /**
     * @return User-defined key/value metadata
     * 
     */
    public Optional<Output<List<ServiceLabelArgs>>> labels() {
        return Optional.ofNullable(this.labels);
    }

    /**
     * Scheduling mode for the service
     * 
     */
    @Import(name="mode")
    private @Nullable Output<ServiceModeArgs> mode;

    /**
     * @return Scheduling mode for the service
     * 
     */
    public Optional<Output<ServiceModeArgs>> mode() {
        return Optional.ofNullable(this.mode);
    }

    /**
     * Name of the service
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the service
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Specification for the rollback strategy of the service
     * 
     */
    @Import(name="rollbackConfig")
    private @Nullable Output<ServiceRollbackConfigArgs> rollbackConfig;

    /**
     * @return Specification for the rollback strategy of the service
     * 
     */
    public Optional<Output<ServiceRollbackConfigArgs>> rollbackConfig() {
        return Optional.ofNullable(this.rollbackConfig);
    }

    /**
     * User modifiable task configuration
     * 
     */
    @Import(name="taskSpec")
    private @Nullable Output<ServiceTaskSpecArgs> taskSpec;

    /**
     * @return User modifiable task configuration
     * 
     */
    public Optional<Output<ServiceTaskSpecArgs>> taskSpec() {
        return Optional.ofNullable(this.taskSpec);
    }

    /**
     * Specification for the update strategy of the service
     * 
     */
    @Import(name="updateConfig")
    private @Nullable Output<ServiceUpdateConfigArgs> updateConfig;

    /**
     * @return Specification for the update strategy of the service
     * 
     */
    public Optional<Output<ServiceUpdateConfigArgs>> updateConfig() {
        return Optional.ofNullable(this.updateConfig);
    }

    private ServiceState() {}

    private ServiceState(ServiceState $) {
        this.auth = $.auth;
        this.convergeConfig = $.convergeConfig;
        this.endpointSpec = $.endpointSpec;
        this.labels = $.labels;
        this.mode = $.mode;
        this.name = $.name;
        this.rollbackConfig = $.rollbackConfig;
        this.taskSpec = $.taskSpec;
        this.updateConfig = $.updateConfig;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceState $;

        public Builder() {
            $ = new ServiceState();
        }

        public Builder(ServiceState defaults) {
            $ = new ServiceState(Objects.requireNonNull(defaults));
        }

        /**
         * @param auth Configuration for the authentication for pulling the images of the service
         * 
         * @return builder
         * 
         */
        public Builder auth(@Nullable Output<ServiceAuthArgs> auth) {
            $.auth = auth;
            return this;
        }

        /**
         * @param auth Configuration for the authentication for pulling the images of the service
         * 
         * @return builder
         * 
         */
        public Builder auth(ServiceAuthArgs auth) {
            return auth(Output.of(auth));
        }

        /**
         * @param convergeConfig A configuration to ensure that a service converges aka reaches the desired that of all task up and running
         * 
         * @return builder
         * 
         */
        public Builder convergeConfig(@Nullable Output<ServiceConvergeConfigArgs> convergeConfig) {
            $.convergeConfig = convergeConfig;
            return this;
        }

        /**
         * @param convergeConfig A configuration to ensure that a service converges aka reaches the desired that of all task up and running
         * 
         * @return builder
         * 
         */
        public Builder convergeConfig(ServiceConvergeConfigArgs convergeConfig) {
            return convergeConfig(Output.of(convergeConfig));
        }

        /**
         * @param endpointSpec Properties that can be configured to access and load balance a service
         * 
         * @return builder
         * 
         */
        public Builder endpointSpec(@Nullable Output<ServiceEndpointSpecArgs> endpointSpec) {
            $.endpointSpec = endpointSpec;
            return this;
        }

        /**
         * @param endpointSpec Properties that can be configured to access and load balance a service
         * 
         * @return builder
         * 
         */
        public Builder endpointSpec(ServiceEndpointSpecArgs endpointSpec) {
            return endpointSpec(Output.of(endpointSpec));
        }

        /**
         * @param labels User-defined key/value metadata
         * 
         * @return builder
         * 
         */
        public Builder labels(@Nullable Output<List<ServiceLabelArgs>> labels) {
            $.labels = labels;
            return this;
        }

        /**
         * @param labels User-defined key/value metadata
         * 
         * @return builder
         * 
         */
        public Builder labels(List<ServiceLabelArgs> labels) {
            return labels(Output.of(labels));
        }

        /**
         * @param labels User-defined key/value metadata
         * 
         * @return builder
         * 
         */
        public Builder labels(ServiceLabelArgs... labels) {
            return labels(List.of(labels));
        }

        /**
         * @param mode Scheduling mode for the service
         * 
         * @return builder
         * 
         */
        public Builder mode(@Nullable Output<ServiceModeArgs> mode) {
            $.mode = mode;
            return this;
        }

        /**
         * @param mode Scheduling mode for the service
         * 
         * @return builder
         * 
         */
        public Builder mode(ServiceModeArgs mode) {
            return mode(Output.of(mode));
        }

        /**
         * @param name Name of the service
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the service
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param rollbackConfig Specification for the rollback strategy of the service
         * 
         * @return builder
         * 
         */
        public Builder rollbackConfig(@Nullable Output<ServiceRollbackConfigArgs> rollbackConfig) {
            $.rollbackConfig = rollbackConfig;
            return this;
        }

        /**
         * @param rollbackConfig Specification for the rollback strategy of the service
         * 
         * @return builder
         * 
         */
        public Builder rollbackConfig(ServiceRollbackConfigArgs rollbackConfig) {
            return rollbackConfig(Output.of(rollbackConfig));
        }

        /**
         * @param taskSpec User modifiable task configuration
         * 
         * @return builder
         * 
         */
        public Builder taskSpec(@Nullable Output<ServiceTaskSpecArgs> taskSpec) {
            $.taskSpec = taskSpec;
            return this;
        }

        /**
         * @param taskSpec User modifiable task configuration
         * 
         * @return builder
         * 
         */
        public Builder taskSpec(ServiceTaskSpecArgs taskSpec) {
            return taskSpec(Output.of(taskSpec));
        }

        /**
         * @param updateConfig Specification for the update strategy of the service
         * 
         * @return builder
         * 
         */
        public Builder updateConfig(@Nullable Output<ServiceUpdateConfigArgs> updateConfig) {
            $.updateConfig = updateConfig;
            return this;
        }

        /**
         * @param updateConfig Specification for the update strategy of the service
         * 
         * @return builder
         * 
         */
        public Builder updateConfig(ServiceUpdateConfigArgs updateConfig) {
            return updateConfig(Output.of(updateConfig));
        }

        public ServiceState build() {
            return $;
        }
    }

}