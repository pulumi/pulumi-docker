import * as pulumi from "@pulumi/pulumi";
import * as outputs from "./types/output";
/**
 * `docker.Network` provides details about a specific Docker Network.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as docker from "@pulumi/docker";
 *
 * const main = docker.getNetwork({
 *     name: "main",
 * });
 * ```
 */
export declare function getNetwork(args: GetNetworkArgs, opts?: pulumi.InvokeOptions): Promise<GetNetworkResult>;
/**
 * A collection of arguments for invoking getNetwork.
 */
export interface GetNetworkArgs {
    /**
     * The name of the Docker network.
     */
    name: string;
}
/**
 * A collection of values returned by getNetwork.
 */
export interface GetNetworkResult {
    /**
     * The driver of the Docker network. Possible values are `bridge`, `host`, `overlay`, `macvlan`. See [network docs](https://docs.docker.com/network/#network-drivers) for more details.
     */
    readonly driver: string;
    /**
     * The ID of this resource.
     */
    readonly id: string;
    /**
     * If `true`, the network is internal.
     */
    readonly internal: boolean;
    /**
     * The IPAM configuration options
     */
    readonly ipamConfigs: outputs.GetNetworkIpamConfig[];
    /**
     * The name of the Docker network.
     */
    readonly name: string;
    /**
     * Only available with bridge networks. See [bridge options docs](https://docs.docker.com/engine/reference/commandline/network_create/#bridge-driver-options) for more details.
     */
    readonly options: {
        [key: string]: any;
    };
    /**
     * Scope of the network. One of `swarm`, `global`, or `local`.
     */
    readonly scope: string;
}
/**
 * `docker.Network` provides details about a specific Docker Network.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as docker from "@pulumi/docker";
 *
 * const main = docker.getNetwork({
 *     name: "main",
 * });
 * ```
 */
export declare function getNetworkOutput(args: GetNetworkOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNetworkResult>;
/**
 * A collection of arguments for invoking getNetwork.
 */
export interface GetNetworkOutputArgs {
    /**
     * The name of the Docker network.
     */
    name: pulumi.Input<string>;
}
