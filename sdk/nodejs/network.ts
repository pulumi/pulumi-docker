// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

/**
 * <!-- Bug: Type and Name are switched -->
 * `docker.Network` provides a docker network resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as docker from "@pulumi/docker";
 *
 * const privateNetwork = new docker.Network("privateNetwork", {});
 * ```
 *
 * ## Import
 *
 * ### Example
 *
 *  Assuming you created a `network` as follows
 *
 *  #!/bin/bash
 *
 *  docker network create foo
 *
 *  prints the long ID
 *
 *  87b57a9b91ecab2db2a6dbf38df74c67d7c7108cbe479d6576574ec2cd8c2d73
 *
 *  you provide the definition for the resource as follows
 *
 *  terraform
 *
 *  resource "docker_network" "foo" {
 *
 *  name = "foo"
 *
 *  }
 *
 *  then the import command is as follows
 *
 *  #!/bin/bash
 *
 * ```sh
 * $ pulumi import docker:index/network:Network foo 87b57a9b91ecab2db2a6dbf38df74c67d7c7108cbe479d6576574ec2cd8c2d73
 * ```
 */
export class Network extends pulumi.CustomResource {
    /**
     * Get an existing Network resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NetworkState, opts?: pulumi.CustomResourceOptions): Network {
        return new Network(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'docker:index/network:Network';

    /**
     * Returns true if the given object is an instance of Network.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Network {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Network.__pulumiType;
    }

    /**
     * Enable manual container attachment to the network.
     */
    public readonly attachable!: pulumi.Output<boolean | undefined>;
    /**
     * Requests daemon to check for networks with same name.
     */
    public readonly checkDuplicate!: pulumi.Output<boolean | undefined>;
    /**
     * The driver of the Docker network. Possible values are `bridge`, `host`, `overlay`, `macvlan`. See [network docs](https://docs.docker.com/network/#network-drivers) for more details.
     */
    public readonly driver!: pulumi.Output<string>;
    /**
     * Create swarm routing-mesh network. Defaults to `false`.
     */
    public readonly ingress!: pulumi.Output<boolean | undefined>;
    /**
     * Whether the network is internal.
     */
    public readonly internal!: pulumi.Output<boolean>;
    /**
     * The IPAM configuration options
     */
    public readonly ipamConfigs!: pulumi.Output<outputs.NetworkIpamConfig[]>;
    /**
     * Driver used by the custom IP scheme of the network. Defaults to `default`
     */
    public readonly ipamDriver!: pulumi.Output<string | undefined>;
    /**
     * Provide explicit options to the IPAM driver. Valid options vary with `ipamDriver` and refer to that driver's documentation for more details.
     */
    public readonly ipamOptions!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Enable IPv6 networking. Defaults to `false`.
     */
    public readonly ipv6!: pulumi.Output<boolean | undefined>;
    /**
     * User-defined key/value metadata
     */
    public readonly labels!: pulumi.Output<outputs.NetworkLabel[] | undefined>;
    /**
     * The name of the Docker network.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Only available with bridge networks. See [bridge options docs](https://docs.docker.com/engine/reference/commandline/network_create/#bridge-driver-options) for more details.
     */
    public readonly options!: pulumi.Output<{[key: string]: any}>;
    /**
     * Scope of the network. One of `swarm`, `global`, or `local`.
     */
    public /*out*/ readonly scope!: pulumi.Output<string>;

    /**
     * Create a Network resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: NetworkArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NetworkArgs | NetworkState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NetworkState | undefined;
            resourceInputs["attachable"] = state ? state.attachable : undefined;
            resourceInputs["checkDuplicate"] = state ? state.checkDuplicate : undefined;
            resourceInputs["driver"] = state ? state.driver : undefined;
            resourceInputs["ingress"] = state ? state.ingress : undefined;
            resourceInputs["internal"] = state ? state.internal : undefined;
            resourceInputs["ipamConfigs"] = state ? state.ipamConfigs : undefined;
            resourceInputs["ipamDriver"] = state ? state.ipamDriver : undefined;
            resourceInputs["ipamOptions"] = state ? state.ipamOptions : undefined;
            resourceInputs["ipv6"] = state ? state.ipv6 : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["options"] = state ? state.options : undefined;
            resourceInputs["scope"] = state ? state.scope : undefined;
        } else {
            const args = argsOrState as NetworkArgs | undefined;
            resourceInputs["attachable"] = args ? args.attachable : undefined;
            resourceInputs["checkDuplicate"] = args ? args.checkDuplicate : undefined;
            resourceInputs["driver"] = args ? args.driver : undefined;
            resourceInputs["ingress"] = args ? args.ingress : undefined;
            resourceInputs["internal"] = args ? args.internal : undefined;
            resourceInputs["ipamConfigs"] = args ? args.ipamConfigs : undefined;
            resourceInputs["ipamDriver"] = args ? args.ipamDriver : undefined;
            resourceInputs["ipamOptions"] = args ? args.ipamOptions : undefined;
            resourceInputs["ipv6"] = args ? args.ipv6 : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["options"] = args ? args.options : undefined;
            resourceInputs["scope"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Network.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Network resources.
 */
export interface NetworkState {
    /**
     * Enable manual container attachment to the network.
     */
    attachable?: pulumi.Input<boolean>;
    /**
     * Requests daemon to check for networks with same name.
     */
    checkDuplicate?: pulumi.Input<boolean>;
    /**
     * The driver of the Docker network. Possible values are `bridge`, `host`, `overlay`, `macvlan`. See [network docs](https://docs.docker.com/network/#network-drivers) for more details.
     */
    driver?: pulumi.Input<string>;
    /**
     * Create swarm routing-mesh network. Defaults to `false`.
     */
    ingress?: pulumi.Input<boolean>;
    /**
     * Whether the network is internal.
     */
    internal?: pulumi.Input<boolean>;
    /**
     * The IPAM configuration options
     */
    ipamConfigs?: pulumi.Input<pulumi.Input<inputs.NetworkIpamConfig>[]>;
    /**
     * Driver used by the custom IP scheme of the network. Defaults to `default`
     */
    ipamDriver?: pulumi.Input<string>;
    /**
     * Provide explicit options to the IPAM driver. Valid options vary with `ipamDriver` and refer to that driver's documentation for more details.
     */
    ipamOptions?: pulumi.Input<{[key: string]: any}>;
    /**
     * Enable IPv6 networking. Defaults to `false`.
     */
    ipv6?: pulumi.Input<boolean>;
    /**
     * User-defined key/value metadata
     */
    labels?: pulumi.Input<pulumi.Input<inputs.NetworkLabel>[]>;
    /**
     * The name of the Docker network.
     */
    name?: pulumi.Input<string>;
    /**
     * Only available with bridge networks. See [bridge options docs](https://docs.docker.com/engine/reference/commandline/network_create/#bridge-driver-options) for more details.
     */
    options?: pulumi.Input<{[key: string]: any}>;
    /**
     * Scope of the network. One of `swarm`, `global`, or `local`.
     */
    scope?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Network resource.
 */
export interface NetworkArgs {
    /**
     * Enable manual container attachment to the network.
     */
    attachable?: pulumi.Input<boolean>;
    /**
     * Requests daemon to check for networks with same name.
     */
    checkDuplicate?: pulumi.Input<boolean>;
    /**
     * The driver of the Docker network. Possible values are `bridge`, `host`, `overlay`, `macvlan`. See [network docs](https://docs.docker.com/network/#network-drivers) for more details.
     */
    driver?: pulumi.Input<string>;
    /**
     * Create swarm routing-mesh network. Defaults to `false`.
     */
    ingress?: pulumi.Input<boolean>;
    /**
     * Whether the network is internal.
     */
    internal?: pulumi.Input<boolean>;
    /**
     * The IPAM configuration options
     */
    ipamConfigs?: pulumi.Input<pulumi.Input<inputs.NetworkIpamConfig>[]>;
    /**
     * Driver used by the custom IP scheme of the network. Defaults to `default`
     */
    ipamDriver?: pulumi.Input<string>;
    /**
     * Provide explicit options to the IPAM driver. Valid options vary with `ipamDriver` and refer to that driver's documentation for more details.
     */
    ipamOptions?: pulumi.Input<{[key: string]: any}>;
    /**
     * Enable IPv6 networking. Defaults to `false`.
     */
    ipv6?: pulumi.Input<boolean>;
    /**
     * User-defined key/value metadata
     */
    labels?: pulumi.Input<pulumi.Input<inputs.NetworkLabel>[]>;
    /**
     * The name of the Docker network.
     */
    name?: pulumi.Input<string>;
    /**
     * Only available with bridge networks. See [bridge options docs](https://docs.docker.com/engine/reference/commandline/network_create/#bridge-driver-options) for more details.
     */
    options?: pulumi.Input<{[key: string]: any}>;
}
