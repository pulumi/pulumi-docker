import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
/**
 * ## Import
 *
 * #!/bin/bash Docker secret cannot be imported as the secret data, once set, is never exposed again.
 */
export declare class Secret extends pulumi.CustomResource {
    /**
     * Get an existing Secret resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecretState, opts?: pulumi.CustomResourceOptions): Secret;
    /**
     * Returns true if the given object is an instance of Secret.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    static isInstance(obj: any): obj is Secret;
    /**
     * Base64-url-safe-encoded secret data
     */
    readonly data: pulumi.Output<string>;
    /**
     * User-defined key/value metadata
     */
    readonly labels: pulumi.Output<outputs.SecretLabel[] | undefined>;
    /**
     * User-defined name of the secret
     */
    readonly name: pulumi.Output<string>;
    /**
     * Create a Secret resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecretArgs, opts?: pulumi.CustomResourceOptions);
}
/**
 * Input properties used for looking up and filtering Secret resources.
 */
export interface SecretState {
    /**
     * Base64-url-safe-encoded secret data
     */
    data?: pulumi.Input<string>;
    /**
     * User-defined key/value metadata
     */
    labels?: pulumi.Input<pulumi.Input<inputs.SecretLabel>[]>;
    /**
     * User-defined name of the secret
     */
    name?: pulumi.Input<string>;
}
/**
 * The set of arguments for constructing a Secret resource.
 */
export interface SecretArgs {
    /**
     * Base64-url-safe-encoded secret data
     */
    data: pulumi.Input<string>;
    /**
     * User-defined key/value metadata
     */
    labels?: pulumi.Input<pulumi.Input<inputs.SecretLabel>[]>;
    /**
     * User-defined name of the secret
     */
    name?: pulumi.Input<string>;
}
