// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.docker.ServiceArgs;
import com.pulumi.docker.Utilities;
import com.pulumi.docker.inputs.ServiceState;
import com.pulumi.docker.outputs.ServiceAuth;
import com.pulumi.docker.outputs.ServiceConvergeConfig;
import com.pulumi.docker.outputs.ServiceEndpointSpec;
import com.pulumi.docker.outputs.ServiceLabel;
import com.pulumi.docker.outputs.ServiceMode;
import com.pulumi.docker.outputs.ServiceRollbackConfig;
import com.pulumi.docker.outputs.ServiceTaskSpec;
import com.pulumi.docker.outputs.ServiceUpdateConfig;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * &lt;!-- Bug: Type and Name are switched --&gt;
 * This resource manages the lifecycle of a Docker service. By default, the creation, update and delete of services are detached.
 *  With the Converge Config Name of the service
 * - `task_spec` (Block List, Min: 1, Max: 1) User modifiable task configuration (see below for nested schema)
 * 
 * ## Import
 * 
 * ### Example Assuming you created a `service` as follows #!/bin/bash docker service create --name foo -p 8080:80 nginx prints th ID 4pcphbxkfn2rffhbhe6czytgi you provide the definition for the resource as follows terraform resource &#34;docker_service&#34; &#34;foo&#34; {
 * 
 *  name = &#34;foo&#34;
 * 
 *  task_spec {
 * 
 *  container_spec {
 * 
 *  image = &#34;nginx&#34;
 * 
 *  }
 * 
 *  }
 * 
 *  endpoint_spec {
 * 
 *  ports {
 * 
 *  target_port
 * 
 * = &#34;80&#34;
 * 
 *  published_port = &#34;8080&#34;
 * 
 *  }
 * 
 *  } } then the import command is as follows #!/bin/bash
 * 
 * ```sh
 *  $ pulumi import docker:index/service:Service foo 4pcphbxkfn2rffhbhe6czytgi
 * ```
 * 
 */
@ResourceType(type="docker:index/service:Service")
public class Service extends com.pulumi.resources.CustomResource {
    /**
     * Configuration for the authentication for pulling the images of the service
     * 
     */
    @Export(name="auth", refs={ServiceAuth.class}, tree="[0]")
    private Output</* @Nullable */ ServiceAuth> auth;

    /**
     * @return Configuration for the authentication for pulling the images of the service
     * 
     */
    public Output<Optional<ServiceAuth>> auth() {
        return Codegen.optional(this.auth);
    }
    /**
     * A configuration to ensure that a service converges aka reaches the desired that of all task up and running
     * 
     */
    @Export(name="convergeConfig", refs={ServiceConvergeConfig.class}, tree="[0]")
    private Output</* @Nullable */ ServiceConvergeConfig> convergeConfig;

    /**
     * @return A configuration to ensure that a service converges aka reaches the desired that of all task up and running
     * 
     */
    public Output<Optional<ServiceConvergeConfig>> convergeConfig() {
        return Codegen.optional(this.convergeConfig);
    }
    /**
     * Properties that can be configured to access and load balance a service
     * 
     */
    @Export(name="endpointSpec", refs={ServiceEndpointSpec.class}, tree="[0]")
    private Output<ServiceEndpointSpec> endpointSpec;

    /**
     * @return Properties that can be configured to access and load balance a service
     * 
     */
    public Output<ServiceEndpointSpec> endpointSpec() {
        return this.endpointSpec;
    }
    /**
     * User-defined key/value metadata
     * 
     */
    @Export(name="labels", refs={List.class,ServiceLabel.class}, tree="[0,1]")
    private Output<List<ServiceLabel>> labels;

    /**
     * @return User-defined key/value metadata
     * 
     */
    public Output<List<ServiceLabel>> labels() {
        return this.labels;
    }
    /**
     * Scheduling mode for the service
     * 
     */
    @Export(name="mode", refs={ServiceMode.class}, tree="[0]")
    private Output<ServiceMode> mode;

    /**
     * @return Scheduling mode for the service
     * 
     */
    public Output<ServiceMode> mode() {
        return this.mode;
    }
    /**
     * Name of the service
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the service
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Specification for the rollback strategy of the service
     * 
     */
    @Export(name="rollbackConfig", refs={ServiceRollbackConfig.class}, tree="[0]")
    private Output</* @Nullable */ ServiceRollbackConfig> rollbackConfig;

    /**
     * @return Specification for the rollback strategy of the service
     * 
     */
    public Output<Optional<ServiceRollbackConfig>> rollbackConfig() {
        return Codegen.optional(this.rollbackConfig);
    }
    /**
     * User modifiable task configuration
     * 
     */
    @Export(name="taskSpec", refs={ServiceTaskSpec.class}, tree="[0]")
    private Output<ServiceTaskSpec> taskSpec;

    /**
     * @return User modifiable task configuration
     * 
     */
    public Output<ServiceTaskSpec> taskSpec() {
        return this.taskSpec;
    }
    /**
     * Specification for the update strategy of the service
     * 
     */
    @Export(name="updateConfig", refs={ServiceUpdateConfig.class}, tree="[0]")
    private Output</* @Nullable */ ServiceUpdateConfig> updateConfig;

    /**
     * @return Specification for the update strategy of the service
     * 
     */
    public Output<Optional<ServiceUpdateConfig>> updateConfig() {
        return Codegen.optional(this.updateConfig);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Service(String name) {
        this(name, ServiceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Service(String name, ServiceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Service(String name, ServiceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("docker:index/service:Service", name, args == null ? ServiceArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Service(String name, Output<String> id, @Nullable ServiceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("docker:index/service:Service", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static Service get(String name, Output<String> id, @Nullable ServiceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Service(name, id, state, options);
    }
}
