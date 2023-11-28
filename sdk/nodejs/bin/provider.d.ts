import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
/**
 * The provider type for the docker package. By default, resources use package-wide configuration
 * settings, however an explicit `Provider` instance may be created and passed during resource
 * construction to achieve fine-grained programmatic control over provider settings. See the
 * [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
 */
export declare class Provider extends pulumi.ProviderResource {
    /**
     * Returns true if the given object is an instance of Provider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    static isInstance(obj: any): obj is Provider;
    /**
     * PEM-encoded content of Docker host CA certificate
     */
    readonly caMaterial: pulumi.Output<string | undefined>;
    /**
     * PEM-encoded content of Docker client certificate
     */
    readonly certMaterial: pulumi.Output<string | undefined>;
    /**
     * Path to directory with Docker TLS config
     */
    readonly certPath: pulumi.Output<string | undefined>;
    /**
     * The Docker daemon address
     */
    readonly host: pulumi.Output<string | undefined>;
    /**
     * PEM-encoded content of Docker client private key
     */
    readonly keyMaterial: pulumi.Output<string | undefined>;
    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ProviderArgs, opts?: pulumi.ResourceOptions);
}
/**
 * The set of arguments for constructing a Provider resource.
 */
export interface ProviderArgs {
    /**
     * PEM-encoded content of Docker host CA certificate
     */
    caMaterial?: pulumi.Input<string>;
    /**
     * PEM-encoded content of Docker client certificate
     */
    certMaterial?: pulumi.Input<string>;
    /**
     * Path to directory with Docker TLS config
     */
    certPath?: pulumi.Input<string>;
    /**
     * The Docker daemon address
     */
    host?: pulumi.Input<string>;
    /**
     * PEM-encoded content of Docker client private key
     */
    keyMaterial?: pulumi.Input<string>;
    registryAuth?: pulumi.Input<pulumi.Input<inputs.ProviderRegistryAuth>[]>;
    /**
     * Additional SSH option flags to be appended when using `ssh://` protocol
     */
    sshOpts?: pulumi.Input<pulumi.Input<string>[]>;
}
