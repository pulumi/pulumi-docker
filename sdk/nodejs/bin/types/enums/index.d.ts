export declare const BuilderVersion: {
    /**
     * The first generation builder for Docker Daemon
     */
    readonly BuilderV1: "BuilderV1";
    /**
     * The builder based on moby/buildkit project
     */
    readonly BuilderBuildKit: "BuilderBuildKit";
};
/**
 * The version of the Docker builder.
 */
export type BuilderVersion = (typeof BuilderVersion)[keyof typeof BuilderVersion];
