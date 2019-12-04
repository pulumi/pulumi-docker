// Copyright 2016-2019, Pulumi Corporation.

using System;
using System.Collections.Concurrent;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Diagnostics;
using System.Linq;
using System.Text;
using System.Text.Json;
using System.Text.RegularExpressions;
using System.Threading.Tasks;

using Pulumi.Serialization;
using Semver;

namespace Pulumi.Docker
{
    /// <summary>
    /// <see cref="DockerBuild"/> may be used to specify detailed instructions about how to build a container.
    /// </summary>
    public class DockerBuild : ResourceArgs
    {
        /// <summary>
        /// <see cref="Context"/> is a path to a directory to use for the Docker build context, usually the directory
        /// in which the Dockerfile resides (although dockerfile may be used to choose a custom location
        /// independent of this choice). If not specified, the context defaults to the current working
        /// directory; if a relative path is used, it is relative to the current working directory that
        /// Pulumi is evaluating.
        /// </summary>
        [Input("context")]
        public Input<string>? Context { get; set; }

        /// <summary>
        /// <see cref="Dockerfile"/> may be used to override the default Dockerfile name and/or location.
        /// By default, it is assumed to be a file named Dockerfile in the root of the build context.
        /// </summary>
        [Input("dockerfile")]
        public Input<string>? Dockerfile { get; set; }

        /// <summary>
        /// An optional map of named build-time argument variables to set during the Docker build.
        /// This flag allows you to pass built-time variables that can be accessed like environment variables
        /// inside the `RUN` instruction.
        /// </summary>
        [Input("args")]
        public InputMap<string>? Args { get; set; }

        /// <summary>
        /// An optional <see cref="CacheFrom"/> object with information about the build stages to use for the Docker
        /// build cache.  This parameter maps to the --cache-from argument to the Docker CLI. If this
        /// parameter is `true`, only the final image will be pulled and passed to --cache-from; if it is
        /// a <see cref="CacheFrom"/> object, the stages named therein will also be pulled and passed to --cache-from.
        /// </summary>
        [Input("cacheFrom")]
        public InputUnion<bool, CacheFrom>? CacheFrom { get; set; }

        /// <summary>
        /// An optional catch-all string to provide extra CLI options to the docker build command.
        /// For example, use to specify `--network host`.
        /// </summary>
        [Input("extraOptions")]
        public InputList<string>? ExtraOptions { get; set; }

        /// <summary>
        /// Environment variables to set on the invocation of `docker build`, for example to support
        /// `DOCKER_BUILDKIT=1 docker build`.
        /// </summary>
        [Input("env")]
        public InputMap<string>? Env { get; set; }
    }

    /// <summary>
    /// A copy of <see cref="DockerBuild"/> but without using <see cref="Input{T}"/> in types.
    /// </summary>
    internal class DockerBuildUnwrap
    {
        public string? Context { get; set; }
        public string? Dockerfile { get; set; }
        public ImmutableDictionary<string, string>? Args { get; set; }
        public CacheFromUnwrap? CacheFrom { get; set; }
        public ImmutableArray<string>? ExtraOptions { get; set; }
        public ImmutableDictionary<string, string>? Env { get; set; }
    }

    /// <summary>
    /// <see cref="CacheFrom"/> may be used to specify build stages to use for the Docker build cache.
    /// The final image is always implicitly included.
    /// </summary>
    public class CacheFrom : ResourceArgs
    {
        /// <summary>
        /// An optional list of build stages to use for caching. Each build stage in this list will be 
        /// built explicitly and pushed to the target repository.A given stage's image will be tagged as
        /// "[stage-name]".
        /// </summary>
        [Input("stages")]
        public InputList<string>? Stages { get; set; }
    }

    /// <summary>
    /// A copy of <see cref="CacheFrom"/> but without using <see cref="Input{T}"/> in types.
    /// </summary>
    internal class CacheFromUnwrap
    {
        public ImmutableArray<string>? Stages { get; set; }
    }

    internal static class Docker
    {
        /// <summary>
        /// Build and push the Dockerfile and context from a path into the requested docker 
        /// repo.  It returns the unique target image name for the image in the docker 
        /// repository. During preview this will build the image, and return the target image
        /// name, without pushing.During a normal update, it will do the same, as well as tag 
        /// and push the image.
        /// </summary>
        /// <param name="imageName">Docker image name.</param>
        /// <param name="pathOrBuild">Image build parameters.</param>
        /// <param name="repositoryUrl">Target repository URL.</param>
        /// <param name="logResource">Associated Pulumi resource.</param>
        /// <param name="registry">Optional image registry properties.</param>
        /// <returns></returns>
        public static Output<string> BuildAndPushImageAsync(string imageName, InputUnion<string, DockerBuild> pathOrBuild,
            Input<string> repositoryUrl, Resource logResource, Input<ImageRegistry>? registry)
        {
            return Output.Tuple(pathOrBuild.Unwrap(), repositoryUrl.ToOutput(), registry.Unwrap()).Apply(async v =>
            {
                var buildVal = v.Item1;
                var repositoryUrlVal = v.Item2;
                var registryVal = v.Item3;

                // Give an initial message indicating what we're about to do.  That way, if anything
                // takes a while, the user has an idea about what's going on.
                Log.Info("Starting docker build and push...", logResource, ephemeral: true);

                var result = await BuildAndPushImageWorkerAsync(
                    imageName, buildVal, repositoryUrlVal, logResource, registryVal).ConfigureAwait(false);

                // If we got here, then building/pushing didn't throw any errors.  update the status bar
                // indicating that things worked properly.  that way, the info bar isn't stuck showing the very
                // last thing printed by some subcommand we launched.
                Log.Info("successfully pushed to docker", logResource, ephemeral: true);

                return result;
            });
        }

        internal static void CheckRepositoryUrl(string repositoryUrl)
        {
            (_, string? tag) = Utils.GetImageNameAndTag(repositoryUrl);

            // We want to report an advisory error to users so that they don't accidentally include a 'tag'
            // in the repo url they supply.  i.e. their repo url can be:
            //
            //      docker.mycompany.com/namespace/myimage
            //
            // but should not be:
            //
            //      docker.mycompany.com/namespace/myimage:latest
            //
            // We could consider removing this check entirely.  However, it is likely valuable to catch
            // clear mistakes where a tag was included in a repo url inappropriately.
            //
            // However, since we do have the check, we need to ensure that we do allow the user to specify
            // a *port* on their repository that the are communicating with.  i.e. it's fine to have:
            //
            //      docker.mycompany.com:5000 or
            //      docker.mycompany.com:5000/namespace/myimage
            //
            // So check if this actually does look like a port, and don't report an error in that case.
            //
            // From: https://www.w3.org/Addressing/URL/url-spec.txt
            //
            //      port        digits
            //
            // Regex = any number of digits, optionally followed by / and any remainder.
            if (tag != null && !Regex.IsMatch(tag, @"^\d+(\/.*)?"))
            {
                throw new ArgumentException($"[{nameof(repositoryUrl)}] should not contain a tag: {tag}");
            }
        }

        private async static Task<string> BuildAndPushImageWorkerAsync(string baseImageName,
            Union<string, DockerBuildUnwrap> pathOrBuild, string repositoryUrl, Resource logResource, ImageRegistryUnwrap? registry)
        {
            CheckRepositoryUrl(repositoryUrl);

            var (_, tag) = Utils.GetImageNameAndTag(baseImageName);

            // login immediately if we're going to have to actually communicate with a remote registry.
            //
            // We know we have to login if:
            //
            //  1. We're doing an update.  In that case, we'll always want to login so we can push our
            //     images to the remote registry.
            //
            // 2. We're in preview or update and the build information contains 'cache from' information. In
            //    that case, we'll want want to pull from the registry and will need to login for that.

            var pullFromCache = pathOrBuild.IsT1 && pathOrBuild.AsT1.CacheFrom != null && !string.IsNullOrEmpty(repositoryUrl);

            // If no `registry` info was passed in we simply assume docker is already
            // logged-in to the correct registry (or uses auto-login via credential helpers).
            if (registry != null)
            {
                if (!Deployment.Instance.IsDryRun || pullFromCache)
                {
                    Log.Info("Logging in to registry...", logResource, ephemeral: true);
                    await LoginToRegistry(registry, logResource).ConfigureAwait(false);
                }
            }

            // If the container specified a cacheFrom parameter, first set up the cached stages.
            string[]? cacheFrom = null;
            if (pullFromCache)
            {
                var cacheFromParam = pathOrBuild.AsT1.CacheFrom;
                cacheFrom = await PullCacheAsync(baseImageName, cacheFromParam, repositoryUrl, logResource).ConfigureAwait(false);
            }

            // Next, build the image.
            (string imageId, string[] stages) = await BuildImageAsync(baseImageName, pathOrBuild, logResource, cacheFrom).ConfigureAwait(false);
            if (imageId == null)
            {
                throw new Exception("Internal error: docker build did not produce an imageId.");
            }

            // Generate a name that uniquely will identify this built image.  This is similar in purpose to
            // the name@digest form that can be normally be retrieved from a docker repository.  However,
            // this tag doesn't require actually pushing the image, nor does it require communicating with
            // some external system, making it suitable for unique identification, even during preview.
            // This also means that if docker produces a new imageId, we'll get a new name here, ensuring that
            // resources (like docker.Image and cloud.Service) will be appropriately replaced.
            var uniqueTaggedImageName = CreateTaggedImageName(repositoryUrl, tag, imageId);

            // Use those to push the image.  Then just return the unique target name. as the final result
            // for our caller to use. Only push the image during an update, do not push during a preview.
            if (!Deployment.Instance.IsDryRun)
            {
                // Push the final image first, then push the stage images to use for caching.

                // First, push with both the optionally-requested-tag *and* imageId (which is guaranteed to
                // be defined).  By using the imageId we give the image a fully unique location that we can
                // successfully pull regardless of whatever else has happened at this repositoryUrl.

                // Next, push only with the optionally-requested-tag.  Users of this API still want to get a
                // nice and simple url that they can reach this image at, without having the explicit imageId
                // hash added to it.  Note: this location is not guaranteed to be idempotent.  For example,
                // pushes on other machines might overwrite that location.
                await TagAndPushImageAsync(baseImageName, repositoryUrl, tag, imageId, logResource).ConfigureAwait(false);
                await TagAndPushImageAsync(baseImageName, repositoryUrl, tag, imageId: null, logResource).ConfigureAwait(false);

                foreach (var stage in stages)
                {
                    await TagAndPushImageAsync(
                        LocalStageImageName(baseImageName, stage), repositoryUrl, stage, imageId: null, logResource).ConfigureAwait(false);
                }
            }

            return uniqueTaggedImageName;
        }

        private static string LocalStageImageName(string imageName, string stage) => $"{imageName}-{stage}";

        private static string CreateTaggedImageName(string repositoryUrl, string? tag, string? imageId)
        {
            var pieces = new List<string>();
            if (tag != null)
            {
                pieces.Add(tag);
            }

            if (imageId != null)
            {
                pieces.Add(imageId);
            }

            // Note: we don't do any validation that the tag is well formed, as per:
            // https://docs.docker.com/engine/reference/commandline/tag
            //
            // If there are any issues with it, we'll just let docker report the problem.
            var fullTag = string.Join("-", pieces);
            return pieces.Count > 0 ? $"{repositoryUrl}:{fullTag}" : repositoryUrl;
        }

        private static async Task<string[]?> PullCacheAsync(string imageName, CacheFromUnwrap? cacheFrom, string repoUrl, Resource logResource)
        {
            // Ensure that we have a repository URL. If we don't, we won't be able to pull anything.
            if (string.IsNullOrEmpty(repoUrl))
            {
                return null;
            }

            Log.Debug($"pulling cache for {imageName} from {repoUrl}", logResource);

            var cacheFromImages = new List<string>();
            var stages = new List<string>();
            if (cacheFrom?.Stages != null)
                stages.AddRange(cacheFrom.Stages);
            stages.Add("");
            foreach (var stage in stages)
            {
                var tag = !string.IsNullOrEmpty(stage) ? $":{stage}" : "";
                var image = $"{repoUrl}{tag}";

                // Try to pull the existing image if it exists.  This may fail if the image does not exist.
                // That's fine, just move onto the next stage.  Also, pass along a flag saying that we
                // should print that error as a warning instead.  We don't want the update to succeed but
                // the user to then get a nasty "error:" message at the end.
                (int code, _) = await RunCommandThatCanFail("docker", new[] { "pull", image }, logResource,
                     reportFullCommandLine: true, reportErrorAsWarning: true).ConfigureAwait(false);
                if (code > 0)
                {
                    continue;
                }

                cacheFromImages.Add(image);
            }

            return cacheFromImages.ToArray();
        }

        private static async Task<(string imageId, string[] stages)> BuildImageAsync(
            string imageName, Union<string, DockerBuildUnwrap> pathOrBuild, Resource logResource, string[]? cacheFrom)
        {
            DockerBuildUnwrap build = pathOrBuild.IsT0 ? new DockerBuildUnwrap { Context = pathOrBuild.AsT0 } : pathOrBuild.AsT1;

            // If the build context is missing, default it to the working directory.
            if (build.Context == null)
            {
                build.Context = ".";
            }

            Log.Info(
                $"Building container image '{imageName}': context={build.Context}" +
                (build.Dockerfile != null ? $", dockerfile={build.Dockerfile}" : "") +
                (build.Args != null ? $", args={JsonSerializer.Serialize(build.Args)}" : ""),
                logResource, ephemeral: true);

            // If the container build specified build stages to cache, build each in turn.
            var stages = new List<string>();
            if (build.CacheFrom?.Stages != null)
            {
                foreach (var stage in build.CacheFrom.Stages)
                {
                    await RunDockerBuild(LocalStageImageName(imageName, stage),
                        build, cacheFrom, logResource, stage).ConfigureAwait(false);
                    stages.Add(stage);
                }
            }

            //// Invoke Docker CLI commands to build.
            await RunDockerBuild(imageName, build, cacheFrom, logResource).ConfigureAwait(false);

            // Finally, inspect the image so we can return the SHA digest. Do not forward the output of this
            // command this to the CLI to show the user.
            var inspectResult = await RunCommandThatMustSucceed(
                "docker", new[] { "image", "inspect", "-f", "{{.Id}}", imageName }, logResource).ConfigureAwait(false);
            if (string.IsNullOrEmpty(inspectResult))
            {
                throw new ResourceException($"No digest available for image {imageName}", logResource);
            }

            // From https://docs.docker.com/registry/spec/api/#content-digests
            //
            // the image id will be a "algorithm:hex" pair.  We don't care about the algorithm part.  All we
            // want is the unique portion we can use elsewhere.  Since we are also going to place this in an
            // image tag, we also don't want the colon, as that's not legal there.  So simply grab the hex
            // portion after the colon and return that.

            var imageId = inspectResult.Trim();
            var colonIndex = imageId.LastIndexOf(":");
            imageId = colonIndex < 0 ? imageId : imageId.Substring(colonIndex + 1);

            return (imageId, stages.ToArray());
        }

        private static async Task RunDockerBuild(
            string imageName, DockerBuildUnwrap build, string[]? cacheFrom, Resource logResource, string? target = null)
        {
            // Prepare the build arguments.
            var buildArgs = new List<string>(new[] { "build" });
            if (!string.IsNullOrEmpty(build.Dockerfile))
            {
                buildArgs.AddRange(new[] { "-f", build.Dockerfile }); // add a custom Dockerfile location.
            }
            if (build.Args != null)
            {
                foreach (var arg in build.Args.Keys)
                {
                    buildArgs.AddRange(new[] { "--build-arg", $"{arg}={build.Args[arg]}" });
                }
            }
            if (build.CacheFrom != null)
            {
                var cacheFromImages = cacheFrom;
                if (cacheFromImages != null && cacheFromImages.Length > 0)
                {
                    buildArgs.AddRange(new[] { "--cache-from", string.Join("", cacheFromImages) });
                }
            }
            if (build.ExtraOptions != null)
            {
                buildArgs.AddRange(build.ExtraOptions);
            }
            buildArgs.Add(build.Context!); // push the docker build context onto the path.

            buildArgs.AddRange(new[] { "-t", imageName }); // tag the image with the chosen name.
            if (target != null)
            {
                buildArgs.AddRange(new[] { "--target", target });
            }

            await RunCommandThatMustSucceed("docker", buildArgs.ToArray(), logResource,
                reportFullCommandLine: true, stdin: null, build.Env).ConfigureAwait(false);
        }

        private static readonly ConcurrentDictionary<(string, string), Task> loginResults = new ConcurrentDictionary<(string, string), Task>();

        private static Task LoginToRegistry(ImageRegistryUnwrap registry, Resource logResource)
        {
            var registryName = registry.Server;
            var username = registry.Username;
            var password = registry.Password;

            return loginResults.GetOrAdd((registryName, username), async _ =>
            {
                var dockerPasswordStdin = await UseDockerPasswordStdin(logResource).ConfigureAwait(false);

                // pass 'reportFullCommandLine: false' here so that if we fail to login we don't emit the
                // username/password in our logs.  Instead, we'll just say "'docker login' failed with code ..."
                if (dockerPasswordStdin)
                {
                    await RunCommandThatMustSucceed(
                        "docker", new[] { "login", "-u", username, "--password-stdin", registryName },
                        logResource, reportFullCommandLine: false, password).ConfigureAwait(false);
                }
                else
                {
                    await RunCommandThatMustSucceed(
                        "docker", new[] { "login", "-u", username, "-p", password, registryName },
                        logResource, reportFullCommandLine: false).ConfigureAwait(false);
                }
            });
        }

        private static Task<bool>? dockerPasswordTask;

        private static Task<bool> UseDockerPasswordStdin(Resource logResource)
        {
            return dockerPasswordTask ??= detectPasswordStdin();

            async Task<bool> detectPasswordStdin()
            {
                // Verify that 'docker' is on the PATH and get the client/server versions
                string dockerVersionString;
                try
                {
                    dockerVersionString = await RunCommandThatMustSucceed(
                        "docker", new[] { "version", "-f", "{{json .}}" }, logResource).ConfigureAwait(false);
                    // IDEA: In the future we could warn here on out-of-date versions of Docker which may not support key
                    // features we want to use.
                    Log.Debug($"'docker version' => {dockerVersionString}", logResource);
                }
                catch
                {
                    throw new ResourceException("No 'docker' command available on PATH: Please install to use container 'build' mode.", logResource);
                }

                // Decide whether to use --password or --password-stdin based on the client version.
                try
                {
                    var versionData = JsonDocument.Parse(dockerVersionString!).RootElement;
                    var clientVersion = SemVersion.Parse(versionData.GetProperty("Client").GetProperty("Version").GetString());
                    return clientVersion.CompareByPrecedence(SemVersion.Parse("17.07.0")) > 0;
                }
                catch (Exception ex)
                {
                    Log.Info($"Could not process Docker version({ex})", logResource);
                }

                return false;
            }
        }

        private static async Task TagAndPushImageAsync(
            string imageName, string repositoryUrl, string? tag, string? imageId, Resource logResource)
        {
            // Ensure we have a unique target name for this image, and tag and push to that unique target.
            await doTagAndPushAsync(CreateTaggedImageName(repositoryUrl, tag, imageId)).ConfigureAwait(false);

            // If the user provided a tag themselves (like "x/y:dev") then also tag and push directly to
            // that 'dev' tag.  This is not going to be a unique location, and future pushes will overwrite
            // this location.  However, that's ok as there's still the unique target we generated above.
            //
            // Note: don't need to do this if imageId was 'undefined' as the above line will have already
            // taken care of things for us.
            if (tag != null && imageId != null)
            {
                await doTagAndPushAsync(CreateTaggedImageName(repositoryUrl, tag, imageId: null)).ConfigureAwait(false);
            }

            async Task doTagAndPushAsync(string targetName)
            {
                await RunCommandThatMustSucceed("docker", new[] { "tag", imageName, targetName }, logResource).ConfigureAwait(false);
                await RunCommandThatMustSucceed("docker", new[] { "push", targetName }, logResource).ConfigureAwait(false);
            }
        }

        private static string GetCommandLineMessage(
            string cmd, string[] args, bool reportFullCommandLine, ImmutableDictionary<string, string>? env)
        {
            var argString = reportFullCommandLine ? string.Join(" ", args) : args[0];
            var envString = env == null ? "" : string.Join(" ", env.Keys.Select(k => $"{k}={env[k]}"));
            return $"'{envString} {cmd} {argString}'";
        }

        private static string GetFailureMessage(
            string cmd, string[] args, bool reportFullCommandLine, int code, ImmutableDictionary<string, string>? env = null)
            => $"{GetCommandLineMessage(cmd, args, reportFullCommandLine, env)} failed with exit code {code}";

        /// <summary>
        /// <see cref="RunCommandThatMustSucceed"/> is used to determine if the full command line should be reported
        /// when an error happens.  In general reporting the full command line is fine.  But it should be set
        /// to false if it might contain sensitive information (like a username/password)
        /// </summary>
        private static async Task<string> RunCommandThatMustSucceed(
            string cmd,
            string[] args,
            Resource logResource,
            bool reportFullCommandLine = true,
            string? stdin = null,
            ImmutableDictionary<string, string>? env = null)
        {
            (int code, string stdout) = await RunCommandThatCanFail(
                cmd, args, logResource, reportFullCommandLine, reportErrorAsWarning: false, stdin, env).ConfigureAwait(false);

            if (code != 0)
            {
                // Fail the entire build and push.  This includes the full output of the command so that at
                // the end the user can review the full docker message about what the problem was.
                //
                // Note: a message about the command failing will have already been ephemerally reported to
                // the status column.
                throw new ResourceException($"{GetFailureMessage(cmd, args, reportFullCommandLine, code)}\n{stdout}", logResource);
            }

            return stdout;
        }

        /// <summary>
        /// Runs a CLI command in a child process, returning a task for the process's exit. Both stdout
        /// and stderr are redirected to process.stdout and process.stder by default.
        /// </summary>
        /// <param name="cmd">Command name to run.</param>
        /// <param name="args">An array of command arguments.</param>
        /// <param name="logResource">Used to specify the resource to associate command output with. Stderr messages
        /// are always sent (since they may contain important information about something that's gone wrong).
        /// Stdout messages will be logged ephemerally to this resource.  This lets the user know there is
        /// progress, without having that dumped on them at the end.  If an error occurs though, the stdout
        /// content will be printed.</param>
        /// <param name="reportFullCommandLine"></param>
        /// <param name="reportErrorAsWarning"></param>
        /// <param name="stdin">If defined, its contents are piped into stdin for the child process.</param>
        /// <param name="env"></param>
        /// <returns></returns>
        private static async Task<(int, string)> RunCommandThatCanFail(
            string cmd,
            string[] args,
            Resource logResource,
            bool reportFullCommandLine,
            bool reportErrorAsWarning,
            string? stdin = null,
            ImmutableDictionary<string, string>? env = null)
        {
            // Let the user ephemerally know the command we're going to execute.
            Log.Info($"Executing {GetCommandLineMessage(cmd, args, reportFullCommandLine, env)}", logResource, ephemeral: true);
            var streamId = Utils.RandomInt();

            using var process = new Process();
            process.StartInfo.FileName = cmd;
            process.StartInfo.Arguments = Utils.EscapeArguments(args);
            process.StartInfo.RedirectStandardInput = stdin != null;

            var stdout = new StringBuilder();
            process.StartInfo.RedirectStandardOutput = true;
            process.OutputDataReceived += new DataReceivedEventHandler((object sender, DataReceivedEventArgs e) =>
            {
                if (e.Data != null) // null would indicate the end of stream
                {
                    Log.Info(e.Data, logResource, streamId, ephemeral: true);
                    stdout.Append(e.Data);
                }
            });

            var stderr = new StringBuilder();
            process.StartInfo.RedirectStandardError = true;
            process.ErrorDataReceived += new DataReceivedEventHandler((object sender, DataReceivedEventArgs e) =>
            {
                // We can't stream these stderr messages as we receive them because we don't knows at
                // this point because Docker uses stderr for both errors and warnings.  So, instead, we
                // just collect the messages, and wait for the process to end to decide how to report them.
                stderr.Append(e.Data ?? "");
            });

            try
            {
                process.Start();
                process.BeginOutputReadLine();
                process.BeginErrorReadLine();

                if (stdin != null)
                {
                    process.StandardInput.Write(stdin);
                    process.StandardInput.Close();
                }

                process.WaitForExit();
                var code = process.ExitCode;

                // If we got any stderr messages, report them as an error/warning depending on the
                // result of the operation.
                if (stderr.Length > 0)
                {
                    if (code > 0 && !reportErrorAsWarning)
                    {
                        // Command returned non-zero code.  Treat these stderr messages as an error.
                        Log.Error(stderr.ToString(), logResource);
                    }
                    else
                    {
                        // command succeeded.  These were just warning.
                        Log.Warn(stderr.ToString(), logResource);
                    }
                }

                // If the command failed report an ephemeral message indicating which command it was.
                // That way the user can immediately see something went wrong in the info bar.  The
                // caller (normally runCommandThatMustSucceed) can choose to also report this
                // non-ephemerally.
                if (code > 0)
                {
                    Log.Info(GetFailureMessage(cmd, args, reportFullCommandLine, code), logResource, ephemeral: true);
                }

                return (code, stdout.ToString());
            }
            catch (Exception ex)
            {
                // This shouldn't normally happen, but we want to be sure Process doesn't throw.
                // In case it does, return an error with the exception printout.
                return (1, ex.ToString());
            }
        }
    }
}