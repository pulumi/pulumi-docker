// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages the secrets of a Docker service in a swarm.
 */
export class Secret extends pulumi.CustomResource {
    /**
     * Get an existing Secret resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecretState): Secret {
        return new Secret(name, <any>state, { id });
    }

    /**
     * The base64 encoded data of the secret.
     */
    public readonly data: pulumi.Output<string>;
    /**
     * User-defined key/value metadata.
     */
    public readonly labels: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The name of the Docker secret.
     */
    public readonly name: pulumi.Output<string>;

    /**
     * Create a Secret resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecretArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecretArgs | SecretState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: SecretState = argsOrState as SecretState | undefined;
            inputs["data"] = state ? state.data : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as SecretArgs | undefined;
            if (!args || args.data === undefined) {
                throw new Error("Missing required property 'data'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            inputs["data"] = args ? args.data : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["name"] = args ? args.name : undefined;
        }
        super("docker:swarm/secret:Secret", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Secret resources.
 */
export interface SecretState {
    /**
     * The base64 encoded data of the secret.
     */
    readonly data?: pulumi.Input<string>;
    /**
     * User-defined key/value metadata.
     */
    readonly labels?: pulumi.Input<{[key: string]: any}>;
    /**
     * The name of the Docker secret.
     */
    readonly name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Secret resource.
 */
export interface SecretArgs {
    /**
     * The base64 encoded data of the secret.
     */
    readonly data: pulumi.Input<string>;
    /**
     * User-defined key/value metadata.
     */
    readonly labels?: pulumi.Input<{[key: string]: any}>;
    /**
     * The name of the Docker secret.
     */
    readonly name: pulumi.Input<string>;
}
