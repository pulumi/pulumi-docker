// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

/**
 * ## Import
 *
 * ### Example
 *
 * Assuming you created a `service` as follows
 *
 * #!/bin/bash
 *
 * docker service create --name foo -p 8080:80 nginx
 *
 * prints th ID
 *
 * 4pcphbxkfn2rffhbhe6czytgi
 *
 * you provide the definition for the resource as follows
 *
 * terraform
 *
 * resource "docker_service" "foo" {
 *
 *   name = "foo"
 *
 *   task_spec {
 *
 *     container_spec {
 *     
 *       image = "nginx"
 *     
 *     }
 *
 *   }
 *
 *   endpoint_spec {
 *
 *     ports {
 *     
 *       target_port    = "80"
 *     
 *       published_port = "8080"
 *     
 *     }
 *
 *   }
 *
 * }
 *
 * then the import command is as follows
 *
 * #!/bin/bash
 *
 * ```sh
 * $ pulumi import docker:index/service:Service foo 4pcphbxkfn2rffhbhe6czytgi
 * ```
 */
export class Service extends pulumi.CustomResource {
    /**
     * Get an existing Service resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceState, opts?: pulumi.CustomResourceOptions): Service {
        return new Service(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'docker:index/service:Service';

    /**
     * Returns true if the given object is an instance of Service.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Service {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Service.__pulumiType;
    }

    /**
     * Configuration for the authentication for pulling the images of the service
     */
    public readonly auth!: pulumi.Output<outputs.ServiceAuth | undefined>;
    /**
     * A configuration to ensure that a service converges aka reaches the desired that of all task up and running
     */
    public readonly convergeConfig!: pulumi.Output<outputs.ServiceConvergeConfig | undefined>;
    /**
     * Properties that can be configured to access and load balance a service
     */
    public readonly endpointSpec!: pulumi.Output<outputs.ServiceEndpointSpec>;
    /**
     * User-defined key/value metadata
     */
    public readonly labels!: pulumi.Output<outputs.ServiceLabel[]>;
    /**
     * Scheduling mode for the service
     */
    public readonly mode!: pulumi.Output<outputs.ServiceMode>;
    /**
     * Name of the service
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specification for the rollback strategy of the service
     */
    public readonly rollbackConfig!: pulumi.Output<outputs.ServiceRollbackConfig | undefined>;
    /**
     * User modifiable task configuration
     */
    public readonly taskSpec!: pulumi.Output<outputs.ServiceTaskSpec>;
    /**
     * Specification for the update strategy of the service
     */
    public readonly updateConfig!: pulumi.Output<outputs.ServiceUpdateConfig | undefined>;

    /**
     * Create a Service resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceArgs | ServiceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServiceState | undefined;
            resourceInputs["auth"] = state ? state.auth : undefined;
            resourceInputs["convergeConfig"] = state ? state.convergeConfig : undefined;
            resourceInputs["endpointSpec"] = state ? state.endpointSpec : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["mode"] = state ? state.mode : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["rollbackConfig"] = state ? state.rollbackConfig : undefined;
            resourceInputs["taskSpec"] = state ? state.taskSpec : undefined;
            resourceInputs["updateConfig"] = state ? state.updateConfig : undefined;
        } else {
            const args = argsOrState as ServiceArgs | undefined;
            if ((!args || args.taskSpec === undefined) && !opts.urn) {
                throw new Error("Missing required property 'taskSpec'");
            }
            resourceInputs["auth"] = args ? args.auth : undefined;
            resourceInputs["convergeConfig"] = args ? args.convergeConfig : undefined;
            resourceInputs["endpointSpec"] = args ? args.endpointSpec : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["mode"] = args ? args.mode : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["rollbackConfig"] = args ? args.rollbackConfig : undefined;
            resourceInputs["taskSpec"] = args ? args.taskSpec : undefined;
            resourceInputs["updateConfig"] = args ? args.updateConfig : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Service.__pulumiType, name, resourceInputs, opts);
    }
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
     * Scheduling mode for the service
     */
    mode?: pulumi.Input<inputs.ServiceMode>;
    /**
     * Name of the service
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
     * Scheduling mode for the service
     */
    mode?: pulumi.Input<inputs.ServiceMode>;
    /**
     * Name of the service
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
