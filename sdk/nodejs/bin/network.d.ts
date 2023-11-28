import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
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
 * ### Example Assuming you created a `network` as follows #!/bin/bash docker network create foo prints the long ID 87b57a9b91ecab2db2a6dbf38df74c67d7c7108cbe479d6576574ec2cd8c2d73 you provide the definition for the resource as follows terraform resource "docker_network" "foo" {
 *
 *  name = "foo" } then the import command is as follows #!/bin/bash
 *
 * ```sh
 *  $ pulumi import docker:index/network:Network foo 87b57a9b91ecab2db2a6dbf38df74c67d7c7108cbe479d6576574ec2cd8c2d73
 * ```
 */
export declare class Network extends pulumi.CustomResource {
    /**
     * Get an existing Network resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NetworkState, opts?: pulumi.CustomResourceOptions): Network;
    /**
     * Returns true if the given object is an instance of Network.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    static isInstance(obj: any): obj is Network;
    /**
     * Enable manual container attachment to the network.
     */
    readonly attachable: pulumi.Output<boolean | undefined>;
    /**
     * Requests daemon to check for networks with same name.
     */
    readonly checkDuplicate: pulumi.Output<boolean | undefined>;
    /**
     * The driver of the Docker network. Possible values are `bridge`, `host`, `overlay`, `macvlan`. See [network docs](https://docs.docker.com/network/#network-drivers) for more details.
     */
    readonly driver: pulumi.Output<string>;
    /**
     * Create swarm routing-mesh network. Defaults to `false`.
     */
    readonly ingress: pulumi.Output<boolean | undefined>;
    /**
     * Whether the network is internal.
     */
    readonly internal: pulumi.Output<boolean>;
    /**
     * The IPAM configuration options
     */
    readonly ipamConfigs: pulumi.Output<outputs.NetworkIpamConfig[]>;
    /**
     * Driver used by the custom IP scheme of the network. Defaults to `default`
     */
    readonly ipamDriver: pulumi.Output<string | undefined>;
    /**
     * Provide explicit options to the IPAM driver. Valid options vary with `ipamDriver` and refer to that driver's documentation for more details.
     */
    readonly ipamOptions: pulumi.Output<{
        [key: string]: any;
    } | undefined>;
    /**
     * Enable IPv6 networking. Defaults to `false`.
     */
    readonly ipv6: pulumi.Output<boolean | undefined>;
    /**
     * User-defined key/value metadata
     */
    readonly labels: pulumi.Output<outputs.NetworkLabel[] | undefined>;
    /**
     * The name of the Docker network.
     */
    readonly name: pulumi.Output<string>;
    /**
     * Only available with bridge networks. See [bridge options docs](https://docs.docker.com/engine/reference/commandline/network_create/#bridge-driver-options) for more details.
     */
    readonly options: pulumi.Output<{
        [key: string]: any;
    }>;
    /**
     * Scope of the network. One of `swarm`, `global`, or `local`.
     */
    readonly scope: pulumi.Output<string>;
    /**
     * Create a Network resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: NetworkArgs, opts?: pulumi.CustomResourceOptions);
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
    ipamOptions?: pulumi.Input<{
        [key: string]: any;
    }>;
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
    options?: pulumi.Input<{
        [key: string]: any;
    }>;
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
    ipamOptions?: pulumi.Input<{
        [key: string]: any;
    }>;
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
    options?: pulumi.Input<{
        [key: string]: any;
    }>;
}
