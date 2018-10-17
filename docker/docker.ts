// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import * as pulumi from "@pulumi/pulumi";
import { ResourceError } from "@pulumi/pulumi/errors";

import * as child_process from "child_process";
import * as semver from "semver";

// Registry is the information required to login to a Docker registry.
export interface Registry {
    registry: string;
    username: string;
    password: string;
}

/**
 * CacheFrom may be used to specify build stages to use for the Docker build cache. The final image
 * is always implicitly included.
 */
export interface CacheFrom {
    /**
     * An optional list of build stages to use for caching. Each build stage in this list will be
     * built explicitly and pushed to the target repository. A given stage's image will be tagged as
     * "[stage-name]".
     */
    stages?: string[];
}

/**
 * DockerBuild may be used to specify detailed instructions about how to build a container.
 */
export interface DockerBuild {
    /**
     * context is a path to a directory to use for the Docker build context, usually the directory
     * in which the Dockerfile resides (although dockerfile may be used to choose a custom location
     * independent of this choice). If not specified, the context defaults to the current working
     * directory; if a relative path is used, it is relative to the current working directory that
     * Pulumi is evaluating.
     */
    context?: string;
    /**
     * dockerfile may be used to override the default Dockerfile name and/or location.  By default,
     * it is assumed to be a file named Dockerfile in the root of the build context.
     */
    dockerfile?: string;
    /**
     * An optional map of named build-time argument variables to set during the Docker build.  This
     * flag allows you to pass built-time variables that can be accessed like environment variables
     * inside the `RUN` instruction.
     */
    args?: {
        [key: string]: string;
    };
    /**
     * An optional CacheFrom object with information about the build stages to use for the Docker
     * build cache. This parameter maps to the --cache-from argument to the Docker CLI. If this
     * parameter is `true`, only the final image will be pulled and passed to --cache-from; if it is
     * a CacheFrom object, the stages named therein will also be pulled and passed to --cache-from.
     */
    cacheFrom?: boolean | CacheFrom;
}

let dockerPasswordPromise: Promise<boolean> | undefined;

function useDockerPasswordStdin(logResource: pulumi.Resource) {
    if (!dockerPasswordPromise) {
        dockerPasswordPromise = useDockerPasswordStdinWorker();
    }

    return dockerPasswordPromise;

    async function useDockerPasswordStdinWorker() {
        // Verify that 'docker' is on the PATH and get the client/server versions
        let dockerVersionString: string;
        try {
            dockerVersionString = await runCommandThatMustSucceed(
                "docker", ["version", "-f", "{{json .}}"], logResource);
            // IDEA: In the future we could warn here on out-of-date versions of Docker which may not support key
            // features we want to use.

            pulumi.log.debug(`'docker version' => ${dockerVersionString}`, logResource);
        }
        catch (err) {
            throw new ResourceError("No 'docker' command available on PATH: Please install to use container 'build' mode.", logResource);
        }

        // Decide whether to use --password or --password-stdin based on the client version.
        try {
            const versionData: any = JSON.parse(dockerVersionString!);
            const clientVersion: string = versionData.Client.Version;
            return semver.gte(clientVersion, "17.07.0", true);
        }
        catch (err) {
            pulumi.log.info(`Could not process Docker version (${err})`, logResource);
        }

        return false;
    }
}

/**
 * @deprecated Use [buildAndPushImageAsync] instead.
 */
export function buildAndPushImage(
    imageName: string,
    pathOrBuild: string | DockerBuild,
    repositoryUrl: pulumi.Input<string>,
    logResource: pulumi.Resource,
    connectToRegistry: () => Promise<Registry>): pulumi.Output<string> {

    return pulumi.output(repositoryUrl).apply(repoUrl =>
        buildAndPushImageAsync(imageName, pathOrBuild, repoUrl, logResource, connectToRegistry));
}

function logEphemeral(message: string, logResource: pulumi.Resource) {
    pulumi.log.info(message, logResource, /*streamId:*/ undefined, /*ephemeral:*/ true);
}

// buildAndPushImageAsync will build and push the Dockerfile and context from [buildPath] into the
// requested docker repo [repositoryUrl].  It returns the unique target image name for the image in
// the docker repository.  During preview this will build the image, and return the target image
// name, without pushing. During a normal update, it will do the same, as well as tag and push the
// image.
export async function buildAndPushImageAsync(
    baseImageName: string,
    pathOrBuild: string | DockerBuild,
    repositoryUrl: string,
    logResource: pulumi.Resource,
    connectToRegistry?: () => Promise<Registry>): Promise<string> {

    // Give an initial message indicating what we're about to do.  That way, if anything
    // takes a while, the user has an idea about what's going on.
    logEphemeral("Starting docker build and push...", logResource);

    const result = await buildAndPushImageWorkerAsync(
        baseImageName, pathOrBuild, repositoryUrl, logResource, connectToRegistry);

    // If we got here, then building/pushing didn't throw any errors.  Update the status bar
    // indicating that things worked properly.  That way, the info bar isn't stuck showing the very
    // last thing printed by some subcommand we launched.
    logEphemeral("Successfully pushed to docker", logResource);

    return result;
}

async function buildAndPushImageWorkerAsync(
    baseImageName: string,
    pathOrBuild: string | DockerBuild,
    repositoryUrl: string,
    logResource: pulumi.Resource,
    connectToRegistry: (() => Promise<Registry>) | undefined): Promise<string> {

    const { imageName, tag } = getImageNameAndTag(baseImageName);

    let loggedIn: Promise<void> | undefined;

    // If no `connectToRegistry` function was passed in we simply assume docker is already logged-in to the correct registry (or uses auto-login via credential helpers)
    // hence we resolve loggedIn immediately.
    if (!connectToRegistry) {
        loggedIn = Promise.resolve();
    }
    const login = () => {
        if (!loggedIn) {
            logEphemeral("Logging in to registry...", logResource);
            loggedIn = connectToRegistry!().then(r => loginToRegistry(r, logResource));
        }
        return loggedIn;
    };

    // If the container specified a cacheFrom parameter, first set up the cached stages.
    let cacheFrom: string[] | undefined;
    if (typeof pathOrBuild !== "string" && pathOrBuild && pathOrBuild.cacheFrom) {
        // NOTE: we pull the promise out of the repository URL s.t. we can observe whether or not it
        // exists. Were we to instead hang an apply off of the raw Input<>, we would never end up
        // running the pull if the repository had not yet been created.
        const cacheFromParam = typeof pathOrBuild.cacheFrom === "boolean" ? {} : pathOrBuild.cacheFrom;
        cacheFrom = await pullCacheAsync(imageName, cacheFromParam, login, repositoryUrl, logResource);
    }

    // First build the image.
    const { imageId, stages } = await buildImageAsync(imageName, pathOrBuild, logResource, cacheFrom);

    if (!repositoryUrl) {
        throw new ResourceError("Expected repository URL to be defined during push", logResource);
    }

    // Generate a name that uniquely will identify this built image.  This is similar in purpose to
    // the name@digest form that can be normally be retrieved from a docker repository.  However,
    // this tag doesn't require actually pushing the image, nor does it require communicating with
    // some external system, making it suitable for unique identification, even during preview.
    // This also means that if docker produces a new imageId, we'll get a new name here, ensuring that
    // resources will be appropriately replaced.
    const uniqueTargetName = createTargetName(repositoryUrl, tag, imageId);

    // Use those then push the image.  Then just return the unique target name. as the final result
    // for our caller to use.
    if (!pulumi.runtime.isDryRun()) {
        // Only push the image during an update, do not push during a preview.
        await login();

        // Push the final image first, then push the stage images to use for caching.
        await tagAndPushImageAsync(imageName, repositoryUrl, tag, imageId, logResource);

        for (const stage of stages) {
            await tagAndPushImageAsync(
                localStageImageName(imageName, stage), repositoryUrl, tag, /*imageId:*/ undefined, logResource);
        }
    }

    return uniqueTargetName;
}

function localStageImageName(imageName: string, stage: string) {
    return `${imageName}-${stage}`;
}

function getImageNameAndTag(baseImageName: string): { imageName: string, tag: string | undefined } {
    const lastColon = baseImageName.lastIndexOf(":");
    const imageName = lastColon < 0 ? baseImageName : baseImageName.substr(0, lastColon);
    const tag = lastColon < 0 ? undefined : baseImageName.substr(lastColon + 1);

    return  { imageName, tag };
}

function createTargetName(repositoryUrl: string, tag: string | undefined, imageId: string | undefined): string {
    const pieces: string[] = [];
    if (tag) {
        pieces.push(tag);
    }

    if (imageId) {
        pieces.push(imageId);
    }

    // A tag name must be valid ASCII and may contain lowercase and uppercase letters, digits,
    // underscores, periods and dashes. A tag name may not start with a period or a dash and may
    // contain a maximum of 128 characters.
    const fullTag = pieces.join("-").replace(/[^-_.a-zA-Z0-9]/, "").substr(0, 128);

    return fullTag ? `${repositoryUrl}:${fullTag}` : repositoryUrl;
}

async function pullCacheAsync(
    imageName: string,
    cacheFrom: CacheFrom,
    login: () => Promise<void>,
    repoUrl: string,
    logResource: pulumi.Resource): Promise<string[] | undefined> {

    // Ensure that we have a repository URL. If we don't, we won't be able to pull anything.
    if (!repoUrl) {
        return undefined;
    }

    pulumi.log.debug(`pulling cache for ${imageName} from ${repoUrl}`, logResource);

    // Ensure that we're logged in to the source registry and attempt to pull each stage in turn.
    await login();

    const cacheFromImages: string[] = [];
    const stages = (cacheFrom.stages || []).concat([""]);
    for (const stage of stages) {
        const tag = stage ? `:${stage}` : "";
        const image = `${repoUrl}${tag}`;

        // Try to pull the existing image if it exists.  This may fail if the image does not exist.
        // That's fine, just move onto the next stage.
        const { code } = await runCommandThatCanFail(
            "docker", ["pull", image], logResource,
            /*reportFullCommand:*/ true, /*reportErrorAsWarning:*/ true);
        if (code) {
            continue;
        }

        cacheFromImages.push(image);
    }

    return cacheFromImages;
}

interface BuildResult {
    imageId: string;
    stages: string[];
}

async function buildImageAsync(
    imageName: string,
    pathOrBuild: string | DockerBuild,
    logResource: pulumi.Resource,
    cacheFrom: string[] | undefined): Promise<BuildResult> {

    let build: DockerBuild;
    if (typeof pathOrBuild === "string") {
        build = {
            context: pathOrBuild,
        };
    } else if (pathOrBuild) {
        build = pathOrBuild;
    } else {
        throw new ResourceError(`Cannot build a container with an empty build specification`, logResource);
    }

    // If the build context is missing, default it to the working directory.
    if (!build.context) {
        build.context = ".";
    }

    logEphemeral(
        `Building container image '${imageName}': context=${build.context}` +
            (build.dockerfile ? `, dockerfile=${build.dockerfile}` : "") +
            (build.args ? `, args=${JSON.stringify(build.args)}` : ""), logResource);

    // If the container build specified build stages to cache, build each in turn.
    const stages = [];
    if (build.cacheFrom && typeof build.cacheFrom !== "boolean" && build.cacheFrom.stages) {
        for (const stage of build.cacheFrom.stages) {
            await dockerBuild(localStageImageName(imageName, stage), build, cacheFrom, logResource, stage);
            stages.push(stage);
        }
    }

    // Invoke Docker CLI commands to build.
    await dockerBuild(imageName, build, cacheFrom, logResource);

    // Finally, inspect the image so we can return the SHA digest. Do not forward the output of this
    // command this to the CLI to show the user.
    const inspectResult = await runCommandThatMustSucceed(
        "docker", ["image", "inspect", "-f", "{{.Id}}", imageName], logResource);
    if (!inspectResult) {
       throw new ResourceError(
           `No digest available for image ${imageName}`, logResource);
    }

    // Trim off the hash kind if there is one.
    let imageId = inspectResult.trim();
    const colonIndex = imageId.lastIndexOf(":");
    imageId = colonIndex < 0 ? imageId : imageId.substr(colonIndex + 1);

    return { imageId, stages };
}

async function dockerBuild(
    imageName: string,
    build: DockerBuild,
    cacheFrom: string[] | undefined,
    logResource: pulumi.Resource,
    target?: string): Promise<void> {

    // Prepare the build arguments.
    const buildArgs: string[] = [ "build" ];
    if (build.dockerfile) {
        buildArgs.push(...[ "-f", build.dockerfile ]); // add a custom Dockerfile location.
    }
    if (build.args) {
        for (const arg of Object.keys(build.args)) {
            buildArgs.push(...[ "--build-arg", `${arg}=${build.args[arg]}` ]);
        }
    }
    if (build.cacheFrom) {
        if (cacheFrom && cacheFrom.length) {
            buildArgs.push(...[ "--cache-from", cacheFrom.join() ]);
        }
    }
    buildArgs.push(build.context!); // push the docker build context onto the path.

    buildArgs.push(...[ "-t", imageName ]); // tag the image with the chosen name.
    if (target) {
        buildArgs.push(...[ "--target", target ]);
    }

    await runCommandThatMustSucceed("docker", buildArgs, logResource);
}

// Keep track of registries and users that have been logged in.  If we've already logged into that
// registry with that user, there's no need to do it again.
const loggedInUsers: { registryName: string, username: string }[] = [];

async function loginToRegistry(registry: Registry, logResource: pulumi.Resource): Promise<void> {
    const { registry: registryName, username, password } = registry;

    if (isLoggedIn(registry)) {
        logEphemeral(`Reusing existing login for ${username}@${registryName}`, logResource);
        return;
    }

    loggedInUsers.push({ registryName, username });

    const dockerPasswordStdin = await useDockerPasswordStdin(logResource);

    // pass 'reportFullCommandLine: false' here so that if we fail to login we don't emit the
    // username/password in our logs.  Instead, we'll just say "'docker login' failed with code ..."
    if (dockerPasswordStdin) {
        await runCommandThatMustSucceed(
            "docker", ["login", "-u", username, "--password-stdin", registryName],
            logResource, /*reportFullCommandLine*/ false, password);
    }
    else {
        await runCommandThatMustSucceed(
            "docker", ["login", "-u", username, "-p", password, registryName],
            logResource, /*reportFullCommandLine*/ false);
    }
}

function isLoggedIn(registry: Registry): boolean {
    for (const { registryName, username } of loggedInUsers) {
        if (registryName === registry.registry &&
            username === registry.username) {

            return true;
        }
    }

    return false;
}

async function tagAndPushImageAsync(
        imageName: string, repositoryUrl: string,
        tag: string | undefined, imageId: string | undefined,
        logResource: pulumi.Resource): Promise<void> {

    // Ensure we have a unique target name for this image, and tag and push to that unique target.
    await doTagAndPushAsync(createTargetName(repositoryUrl, tag, imageId));

    if (tag) {
        // user provided a tag themselves (like "x/y:dev").  In this case, also tag and push
        // directly to that 'dev' tag.  This is not going to be a unique location, and future pushes
        // will overwrite this location.  However, that's ok as there's still the unique target we
        // generated above.
        //
        // We don't need to do this for the main image we tagged and pushed as the repo will
        // automatically give is a :latest tag that serves the same purpose.  So this tagged/pushed
        // build will simply act as the 'latest' pointer for this specific tag.
        await doTagAndPushAsync(`${repositoryUrl}:${tag}`);
    }

    return;

    async function doTagAndPushAsync(targetName: string) {
        await runCommandThatMustSucceed("docker", ["tag", imageName, targetName], logResource);
        await runCommandThatMustSucceed("docker", ["push", targetName], logResource);
    }
}

interface CommandResult {
    code: number;
    stdout: string;
}

function getCommandLineMessage(cmd: string, args: string[], reportFullCommandLine: boolean) {
    const argString = reportFullCommandLine ? args.join(" ") : args[0];
    return `'${cmd} ${argString}'`;
}

function getFailureMessage(cmd: string, args: string[], reportFullCommandLine: boolean, code: number) {
    return `${getCommandLineMessage(cmd, args, reportFullCommandLine)} failed with exit code ${code}`;
}

// [reportFullCommandLine] is used to determine if the full command line should be reported
// when an error happens.  In general reporting the full command line is fine.  But it should be set
// to false if it might contain sensitive information (like a username/password)
async function runCommandThatMustSucceed(
    cmd: string,
    args: string[],
    logResource: pulumi.Resource,
    reportFullCommandLine: boolean = true,
    stdin?: string): Promise<string> {

    const { code, stdout } = await runCommandThatCanFail(
        cmd, args, logResource, reportFullCommandLine, /*reportErrorAsWarning:*/ false, stdin);

    if (code !== 0) {
        // Fail the entire build and push.  This includes the full output of the command so that at
        // the end the user can review the full docker message about what the problem was.
        //
        // Note: a message about the command failing will have already been ephemerally reported to
        // the status column.
        throw new ResourceError(
            `${getFailureMessage(cmd, args, reportFullCommandLine, code)}\n${stdout}`, logResource);
    }

    return stdout;
}

// Runs a CLI command in a child process, returning a promise for the process's exit. Both stdout
// and stderr are redirected to process.stdout and process.stder by default.
//
// If the [stdin] argument is defined, it's contents are piped into stdin for the child process.
//
// [logResource] is used to specify the resource to associate command output with. Stderr messages
// are always sent (since they may contain important information about something that's gone wrong).
// Stdout messages will be logged ephemerally to this resource.  This lets the user know there is
// progress, without having that dumped on them at the end.  If an error occurs though, the stdout
// content will be printed.
async function runCommandThatCanFail(
    cmd: string,
    args: string[],
    logResource: pulumi.Resource,
    reportFullCommandLine: boolean,
    reportErrorAsWarning: boolean,
    stdin?: string): Promise<CommandResult> {

    // Let the user ephemerally know the command we're going to execute.
    logEphemeral(`Executing ${getCommandLineMessage(cmd, args, reportFullCommandLine)}`, logResource);

    // Generate a unique stream-ID that we'll associate all the docker output with. This will allow
    // each spawned CLI command's output to associated with 'resource' and also streamed to the UI
    // in pieces so that it can be displayed live.  The stream-ID is so that the UI knows these
    // messages are all related and should be considered as one large message (just one that was
    // sent over in chunks).
    //
    // We use Math.random here in case our package is loaded multiple times in memory (i.e. because
    // different downstream dependencies depend on different versions of us).  By being random we
    // effectively make it completely unlikely that any two cli outputs could map to the same stream
    // id.
    //
    // Pick a reasonably distributed number between 0 and 2^30.  This will fit as an int32
    // which the grpc layer needs.
    const streamID = Math.floor(Math.random() * (1 << 30));

    return new Promise<CommandResult>((resolve, reject) => {
        const p = child_process.spawn(cmd, args);

        // We store the results from stdout in memory and will return them as a string.
        const stdOutChunks: Buffer[] = [];
        let stdErrChunks: Buffer[] = [];

        p.stdout.on("data", (chunk: Buffer) => {
            // Report all stdout messages as ephemeral messages.  That way they show up in the
            // info bar as they're happening.  But they do not overwhelm the user as the end
            // of the run.
            logEphemeral(chunk.toString(), logResource);
            stdOutChunks.push(chunk);
        });

        p.stderr.on("data", (chunk: Buffer) => {
            // We can't stream these stderr messages as we receive them because we don't knows at
            // this point because Docker uses stderr for both errors and warnings.  So, instead, we
            // just collect the messages, and wait for the process to end to decide how to report
            // them.
            stdErrChunks.push(chunk);
        });

        p.on("error", err => {
            stdErrChunks.push(new Buffer(err.message));
            finish(/*code: */ 1);
        });

        p.on("close", code => {
            finish(code);
        });

        if (stdin) {
            p.stdin.end(stdin);
        }

        return;

        function finish(code: number) {
            logStdErrMessages(code);
            const stdout = Buffer.concat(stdOutChunks).toString();

            if (code) {
                // Report an ephemeral message indicating which command failed.  That way the user
                // can immediately see something went wrong, and what command caused it.
                logEphemeral(getFailureMessage(cmd, args, reportFullCommandLine, code), logResource);
            }

            resolve({ code, stdout });
        }

        function logStdErrMessages(code: number) {
            if (stdErrChunks.length > 0) {
                const errorOrWarnings = Buffer.concat(stdErrChunks).toString();
                stdErrChunks = [];

                if (code && !reportErrorAsWarning) {
                    // Command returned non-zero code.  Treat these stderr messages as an error.
                    pulumi.log.error(errorOrWarnings, logResource, streamID);
                }
                else {
                    // command succeeded.  These were just warning.
                    pulumi.log.warn(errorOrWarnings, logResource, streamID);
                }
            }
        }
    });
}
