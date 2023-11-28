import * as pulumi from "@pulumi/pulumi";
/**
 * `docker.getLogs` provides logs from specific container
 */
export declare function getLogs(args: GetLogsArgs, opts?: pulumi.InvokeOptions): Promise<GetLogsResult>;
/**
 * A collection of arguments for invoking getLogs.
 */
export interface GetLogsArgs {
    details?: boolean;
    /**
     * Discard headers that docker appends to each log entry
     */
    discardHeaders?: boolean;
    follow?: boolean;
    /**
     * If true populate computed value `logsListString`
     */
    logsListStringEnabled?: boolean;
    /**
     * The name of the Docker Container
     */
    name: string;
    showStderr?: boolean;
    showStdout?: boolean;
    since?: string;
    tail?: string;
    timestamps?: boolean;
    until?: string;
}
/**
 * A collection of values returned by getLogs.
 */
export interface GetLogsResult {
    readonly details?: boolean;
    /**
     * Discard headers that docker appends to each log entry
     */
    readonly discardHeaders?: boolean;
    readonly follow?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * If true populate computed value `logsListString`
     */
    readonly logsListStringEnabled?: boolean;
    /**
     * List of container logs, each element is a line.
     */
    readonly logsListStrings: string[];
    /**
     * The name of the Docker Container
     */
    readonly name: string;
    readonly showStderr?: boolean;
    readonly showStdout?: boolean;
    readonly since?: string;
    readonly tail?: string;
    readonly timestamps?: boolean;
    readonly until?: string;
}
/**
 * `docker.getLogs` provides logs from specific container
 */
export declare function getLogsOutput(args: GetLogsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLogsResult>;
/**
 * A collection of arguments for invoking getLogs.
 */
export interface GetLogsOutputArgs {
    details?: pulumi.Input<boolean>;
    /**
     * Discard headers that docker appends to each log entry
     */
    discardHeaders?: pulumi.Input<boolean>;
    follow?: pulumi.Input<boolean>;
    /**
     * If true populate computed value `logsListString`
     */
    logsListStringEnabled?: pulumi.Input<boolean>;
    /**
     * The name of the Docker Container
     */
    name: pulumi.Input<string>;
    showStderr?: pulumi.Input<boolean>;
    showStdout?: pulumi.Input<boolean>;
    since?: pulumi.Input<string>;
    tail?: pulumi.Input<string>;
    timestamps?: pulumi.Input<boolean>;
    until?: pulumi.Input<string>;
}
