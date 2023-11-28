import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
/**
 * ## Import
 *
 * ### Example Assuming you created a `service` as follows #!/bin/bash docker service create --name foo -p 8080:80 nginx prints th ID 4pcphbxkfn2rffhbhe6czytgi you provide the definition for the resource as follows terraform resource "docker_service" "foo" {
 *
 *  name = "foo"
 *
 *  task_spec {
 *
 *  container_spec {
 *
 *  image = "nginx"
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
 * = "80"
 *
 *  published_port = "8080"
 *
 *  }
 *
 *  } } then the import command is as follows #!/bin/bash
 *
 * ```sh
 *  $ pulumi import docker:index/service:Service foo 4pcphbxkfn2rffhbhe6czytgi
 * ```
 */
export declare class Service extends pulumi.CustomResource {
    /**
     * Get an existing Service resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceState, opts?: pulumi.CustomResourceOptions): Service;
    /**
     * Returns true if the given object is an instance of Service.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    static isInstance(obj: any): obj is Service;
    /**
     * Configuration for the authentication for pulling the images of the service
     */
    readonly auth: pulumi.Output<outputs.ServiceAuth | undefined>;
    /**
     * A configuration to ensure that a service converges aka reaches the desired that of all task up and running
     */
    readonly convergeConfig: pulumi.Output<outputs.ServiceConvergeConfig | undefined>;
    /**
     * Properties that can be configured to access and load balance a service
     */
    readonly endpointSpec: pulumi.Output<outputs.ServiceEndpointSpec>;
    /**
     * User-defined key/value metadata
     */
    readonly labels: pulumi.Output<outputs.ServiceLabel[]>;
    /**
     * The mode of resolution to use for internal load balancing between tasks
     */
    readonly mode: pulumi.Output<outputs.ServiceMode>;
    /**
     * A random name for the port
     */
    readonly name: pulumi.Output<string>;
    /**
     * Specification for the rollback strategy of the service
     */
    readonly rollbackConfig: pulumi.Output<outputs.ServiceRollbackConfig | undefined>;
    /**
     * User modifiable task configuration
     */
    readonly taskSpec: pulumi.Output<outputs.ServiceTaskSpec>;
    /**
     * Specification for the update strategy of the service
     */
    readonly updateConfig: pulumi.Output<outputs.ServiceUpdateConfig | undefined>;
    /**
     * Create a Service resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceArgs, opts?: pulumi.CustomResourceOptions);
}
/**
 * Input properties used for looking up and filtering Service resources.
 */
export interface ServiceState {
    /**
     * Configuration for the authentication for pulling the images of the service
     */
    auth?: pulumi.Input<inputs.ServiceAuth>;
    /**
     * A configuration to ensure that a service converges aka reaches the desired that of all task up and running
     */
    convergeConfig?: pulumi.Input<inputs.ServiceConvergeConfig>;
    /**
     * Properties that can be configured to access and load balance a service
     */
    endpointSpec?: pulumi.Input<inputs.ServiceEndpointSpec>;
    /**
     * User-defined key/value metadata
     */
    labels?: pulumi.Input<pulumi.Input<inputs.ServiceLabel>[]>;
    /**
     * The mode of resolution to use for internal load balancing between tasks
     */
    mode?: pulumi.Input<inputs.ServiceMode>;
    /**
     * A random name for the port
     */
    name?: pulumi.Input<string>;
    /**
     * Specification for the rollback strategy of the service
     */
    rollbackConfig?: pulumi.Input<inputs.ServiceRollbackConfig>;
    /**
     * User modifiable task configuration
     */
    taskSpec?: pulumi.Input<inputs.ServiceTaskSpec>;
    /**
     * Specification for the update strategy of the service
     */
    updateConfig?: pulumi.Input<inputs.ServiceUpdateConfig>;
}
/**
 * The set of arguments for constructing a Service resource.
 */
export interface ServiceArgs {
    /**
     * Configuration for the authentication for pulling the images of the service
     */
    auth?: pulumi.Input<inputs.ServiceAuth>;
    /**
     * A configuration to ensure that a service converges aka reaches the desired that of all task up and running
     */
    convergeConfig?: pulumi.Input<inputs.ServiceConvergeConfig>;
    /**
     * Properties that can be configured to access and load balance a service
     */
    endpointSpec?: pulumi.Input<inputs.ServiceEndpointSpec>;
    /**
     * User-defined key/value metadata
     */
    labels?: pulumi.Input<pulumi.Input<inputs.ServiceLabel>[]>;
    /**
     * The mode of resolution to use for internal load balancing between tasks
     */
    mode?: pulumi.Input<inputs.ServiceMode>;
    /**
     * A random name for the port
     */
    name?: pulumi.Input<string>;
    /**
     * Specification for the rollback strategy of the service
     */
    rollbackConfig?: pulumi.Input<inputs.ServiceRollbackConfig>;
    /**
     * User modifiable task configuration
     */
    taskSpec: pulumi.Input<inputs.ServiceTaskSpec>;
    /**
     * Specification for the update strategy of the service
     */
    updateConfig?: pulumi.Input<inputs.ServiceUpdateConfig>;
}
