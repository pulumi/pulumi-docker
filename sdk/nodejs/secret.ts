// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * ## Import
 *
 * Docker secret cannot be imported as the secret data, once set, is never exposed again.
 */
export class Secret extends pulumi.CustomResource {
    /**
     * Get an existing Secret resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecretState, opts?: pulumi.CustomResourceOptions): Secret {
        return new Secret(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'docker:index/secret:Secret';

    /**
     * Returns true if the given object is an instance of Secret.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Secret {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Secret.__pulumiType;
    }

    /**
     * The base64 encoded data of the secret.
     */
    public readonly data!: pulumi.Output<string>;
    /**
     * See Labels below for details.
     */
    public readonly labels!: pulumi.Output<outputs.SecretLabel[] | undefined>;
    /**
     * The name of the Docker secret.
     */
    public readonly name!: pulumi.Output<string>;

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
            const state = argsOrState as SecretState | undefined;
            inputs["data"] = state ? state.data : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as SecretArgs | undefined;
            if ((!args || args.data === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'data'");
            }
            inputs["data"] = args ? args.data : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["name"] = args ? args.name : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Secret.__pulumiType, name, inputs, opts);
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
     * See Labels below for details.
     */
    readonly labels?: pulumi.Input<pulumi.Input<inputs.SecretLabel>[]>;
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
     * See Labels below for details.
     */
    readonly labels?: pulumi.Input<pulumi.Input<inputs.SecretLabel>[]>;
    /**
     * The name of the Docker secret.
     */
    readonly name?: pulumi.Input<string>;
}
