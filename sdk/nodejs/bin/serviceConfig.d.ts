import * as pulumi from "@pulumi/pulumi";
/**
 * ## Import
 *
 * ### Example Assuming you created a `config` as follows #!/bin/bash printf '{"a":"b"}' | docker config create foo - prints the id
 *
 * 08c26c477474478d971139f750984775a7f019dbe8a2e7f09d66a187c009e66d you provide the definition for the resource as follows terraform resource "docker_config" "foo" {
 *
 *  name = "foo"
 *
 *  data = base64encode("{\"a\"\"b\"}") } then the import command is as follows #!/bin/bash
 *
 * ```sh
 *  $ pulumi import docker:index/serviceConfig:ServiceConfig foo 08c26c477474478d971139f750984775a7f019dbe8a2e7f09d66a187c009e66d
 * ```
 */
export declare class ServiceConfig extends pulumi.CustomResource {
    /**
     * Get an existing ServiceConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceConfigState, opts?: pulumi.CustomResourceOptions): ServiceConfig;
    /**
     * Returns true if the given object is an instance of ServiceConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    static isInstance(obj: any): obj is ServiceConfig;
    /**
     * Base64-url-safe-encoded config data
     */
    readonly data: pulumi.Output<string>;
    /**
     * User-defined name of the config
     */
    readonly name: pulumi.Output<string>;
    /**
     * Create a ServiceConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceConfigArgs, opts?: pulumi.CustomResourceOptions);
}
/**
 * Input properties used for looking up and filtering ServiceConfig resources.
 */
export interface ServiceConfigState {
    /**
     * Base64-url-safe-encoded config data
     */
    data?: pulumi.Input<string>;
    /**
     * User-defined name of the config
     */
    name?: pulumi.Input<string>;
}
/**
 * The set of arguments for constructing a ServiceConfig resource.
 */
export interface ServiceConfigArgs {
    /**
     * Base64-url-safe-encoded config data
     */
    data: pulumi.Input<string>;
    /**
     * User-defined name of the config
     */
    name?: pulumi.Input<string>;
}
