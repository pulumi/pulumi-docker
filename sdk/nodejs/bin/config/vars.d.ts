import * as outputs from "../types/output";
/**
 * PEM-encoded content of Docker host CA certificate
 */
export declare const caMaterial: string | undefined;
/**
 * PEM-encoded content of Docker client certificate
 */
export declare const certMaterial: string | undefined;
/**
 * Path to directory with Docker TLS config
 */
export declare const certPath: string | undefined;
/**
 * The Docker daemon address
 */
export declare const host: string | undefined;
/**
 * PEM-encoded content of Docker client private key
 */
export declare const keyMaterial: string | undefined;
export declare const registryAuth: outputs.config.RegistryAuth[] | undefined;
/**
 * Additional SSH option flags to be appended when using `ssh://` protocol
 */
export declare const sshOpts: string[] | undefined;
