// Copyright 2016-2020, Pulumi Corporation.

package docker

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"os/exec"
	"strings"

	"github.com/Masterminds/semver"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/pkg/util/logging"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// DockerBuild may be used to specify detailed instructions about how to build a container.
type DockerBuild struct {
	// Context is a path to a directory to use for the Docker build context, usually the directory
	// in which the Dockerfile resides (although dockerfile may be used to choose a custom location
	// independent of this choice). If not specified, the context defaults to the current working
	// directory; if a relative path is used, it is relative to the current working directory that
	// Pulumi is evaluating.
	Context pulumi.StringInput `pulumi:"context"`

	// Dockerfile may be used to override the default Dockerfile name and/or location.
	// By default, it is assumed to be a file named Dockerfile in the root of the build context.
	Dockerfile pulumi.StringInput `pulumi:"dockerfile"`

	// An optional map of named build-time argument variables to set during the Docker build.
	// This flag allows you to pass built-time variables that can be accessed like environment variables
	// inside the `RUN` instruction.
	Args pulumi.MapInput `pulumi:"args"`

	// An optional CacheFrom object with information about the build stages to use for the Docker
	// build cache.  This parameter maps to the --cache-from argument to the Docker CLI. If this
	// parameter is `true`, only the final image will be pulled and passed to --cache-from; if it is
	// a CacheFrom object, the stages named therein will also be pulled and passed to --cache-from.
	CacheFrom *CacheFrom `pulumi:"cacheFrom"`

	// An optional catch-all string to provide extra CLI options to the docker build command.
	// For example, use to specify `--network host`.
	ExtraOptions pulumi.StringArrayInput `pulumi:"extraOptions"`

	// Environment variables to set on the invocation of `docker build`, for example to support
	// `DOCKER_BUILDKIT=1 docker build`.
	Env pulumi.MapInput `pulumi:"env"`

	// The target of the dockerfile to build.
	Target pulumi.StringInput `pulumi:"target"`
}

type dockerBuild struct {
	Context      string
	Dockerfile   string
	Args         map[string]string
	CacheFrom    *cacheFrom
	ExtraOptions []string
	Env          map[string]string
	Target       string
}

// CacheFrom may be used to specify build stages to use for the Docker build cache.
// The final image is always implicitly included
type CacheFrom struct {
	Stages pulumi.StringArrayInput `pulumi:"stages"`
}

type cacheFrom struct {
	Stages []string
}

func buildAndPushImageAsync(ctx *pulumi.Context, imageName string, build dockerBuild, repositoryURL string,
	logResource pulumi.Resource, skipPush bool, registry imageRegistry) (string, error) {

	// Give an initial message indicating what we're about to do.  That way, if anything
	// takes a while, the user has an idea about what's going on.
	logging.Infof("Starting docker build and push...")

	res, err := buildAndPushImageWorkerAsync(ctx, imageName, build, repositoryURL, logResource, skipPush, &registry)
	if err != nil {
		return "", err
	}

	// If we got here, then building/pushing didn't throw any errors.  update the status bar
	// indicating that things worked properly.  that way, the info bar isn't stuck showing the very
	// last thing printed by some subcommand we launched.
	logging.Infof("Successfully pushed to docker.")
	return res, nil
}

func buildAndPushImageWorkerAsync(ctx *pulumi.Context, baseImageName string, build dockerBuild,
	repositoryURL string, logResource pulumi.Resource, skipPush bool, registry *imageRegistry) (string, error) {

	err := checkRepositoryURL(repositoryURL)
	if err != nil {
		return "", err
	}

	_, tag := getImageNameAndTag(baseImageName)

	// login immediately if we're going to have to actually communicate with a remote registry.
	//
	// We know we have to login if:
	//
	//  1. We're doing an update.  In that case, we'll always want to login so we can push our
	//     images to the remote registry.
	//
	// 2. We're in preview or update and the build information contains 'cache from' information. In
	//    that case, we'll want want to pull from the registry and will need to login for that.
	pullFromCache := build.CacheFrom != nil && len(repositoryURL) > 0

	// If no `registry` info was passed in we simply assume docker is already
	// logged-in to the correct registry (or uses auto-login via credential helpers).
	if registry != nil {
		if !ctx.DryRun() || pullFromCache {
			logging.Infof("Logging into registry...")
			err := loginToRegistry(*registry, logResource)
			if err != nil {
				return "", err
			}
		}
	}

	// If the container specified a cacheFrom parameter, first set up the cached stages.
	var cacheFrom []string
	if pullFromCache {
		cacheFrom = pullCacheAsync(baseImageName, *build.CacheFrom, repositoryURL, logResource)
	}

	// Next, build the image.
	imageID, stages, err := buildImageAsync(baseImageName, build, logResource, cacheFrom)
	if err != nil {
		return "", err
	}

	// Generate a name that uniquely will identify this built image.  This is similar in purpose to
	// the name@digest form that can be normally be retrieved from a docker repository.  However,
	// this tag doesn't require actually pushing the image, nor does it require communicating with
	// some external system, making it suitable for unique identification, even during preview.
	// This also means that if docker produces a new imageId, we'll get a new name here, ensuring that
	// resources (like docker.Image and cloud.Service) will be appropriately replaced.
	uniqueTaggedImageName := createTaggedImageName(repositoryURL, tag, imageID)

	// Use those to push the image.  Then just return the unique target name. as the final result
	// for our caller to use. Only push the image during an update, do not push during a preview.
	if !ctx.DryRun() && !skipPush {
		// Push the final image first, then push the stage images to use for caching.

		// First, push with both the optionally-requested-tag *and* imageId (which is guaranteed to
		// be defined).  By using the imageId we give the image a fully unique location that we can
		// successfully pull regardless of whatever else has happened at this repositoryUrl.
		err := tagAndPushImageAsync(baseImageName, repositoryURL, tag, imageID, logResource)
		if err != nil {
			return "", err
		}

		// Next, push only with the optionally-requested-tag.  Users of this API still want to get a
		// nice and simple url that they can reach this image at, without having the explicit imageId
		// hash added to it.  Note: this location is not guaranteed to be idempotent.  For example,
		// pushes on other machines might overwrite that location.
		err = tagAndPushImageAsync(baseImageName, repositoryURL, tag, "", logResource)
		if err != nil {
			return "", err
		}

		for _, stage := range stages {
			err = tagAndPushImageAsync(localStageImageName(baseImageName, stage),
				repositoryURL, stage, "", logResource)
			if err != nil {
				return "", err
			}
		}
	}

	return uniqueTaggedImageName, nil
}

func pullCacheAsync(imageName string, cacheFrom cacheFrom, repoURL string, logResource pulumi.Resource) []string {
	if len(repoURL) == 0 {
		return nil
	}

	logging.Infof("Pulling cache for %s from %s", imageName, repoURL)

	var cacheFromImages []string
	stages := append(cacheFrom.Stages, "")
	for _, stage := range stages {
		tag := stage
		if len(stage) > 0 {
			tag = ":" + stage
		}
		image := repoURL + tag

		// Try to pull the existing image if it exists.  This may fail if the image does not exist.
		// That's fine, just move onto the next stage.  Also, pass along a flag saying that we
		// should print that error as a warning instead.  We don't want the update to succeed but
		// the user to then get a nasty "error:" message at the end.
		_, err := runCommandThatCanFail("docker", []string{"pull", image}, logResource, true, true, "", nil)
		if err == nil {
			continue
		}

		cacheFromImages = append(cacheFromImages, image)
	}

	return cacheFromImages
}

func tagAndPushImageAsync(imageName string, repositoryURL string, tag string,
	imageID string, logResource pulumi.Resource) error {

	doTagAndPushAsync := func(targetName string) error {
		_, err := runBasicCommandThatMustSucceed("docker", []string{"tag", imageName, targetName}, logResource)
		if err != nil {
			return err
		}
		_, err = runBasicCommandThatMustSucceed("docker", []string{"push", targetName}, logResource)
		return err
	}

	// Ensure we have a unique target name for this image, and tag and push to that unique target.
	err := doTagAndPushAsync(createTaggedImageName(repositoryURL, tag, imageID))
	if err != nil {
		return err
	}

	// If the user provided a tag themselves (like "x/y:dev") then also tag and push directly to
	// that 'dev' tag.  This is not going to be a unique location, and future pushes will overwrite
	// this location.  However, that's ok as there's still the unique target we generated above.
	//
	// Note: don't need to do this if imageId was 'undefined' as the above line will have already
	// taken care of things for us.
	if len(tag) > 0 && len(imageID) > 0 {
		err := doTagAndPushAsync(createTaggedImageName(repositoryURL, tag, ""))
		return err
	}

	return nil
}

func createTaggedImageName(repositoryURL string, tag string, imageID string) string {
	var pieces []string
	if len(tag) > 0 {
		pieces = append(pieces, tag)
	}
	if len(imageID) > 0 {
		pieces = append(pieces, imageID)
	}

	// Note: we don't do any validation that the tag is well formed, as per:
	// https://docs.docker.com/engine/reference/commandline/tag
	//
	// If there are any issues with it, we'll just let docker report the problem.
	fullTag := strings.Join(pieces, "-")
	if len(fullTag) > 0 {
		return fmt.Sprintf("%s:%s", repositoryURL, fullTag)
	}
	return repositoryURL
}

// Note: unlike the Typescript and Dotnet implementations, you must pass in a dockerBuild here.
// If you have just the path, `build = &dockerBuild{Context: path}`.
func buildImageAsync(imageName string, build dockerBuild,
	logResource pulumi.Resource, cacheFrom []string) (string, []string, error) {

	// If the build context is missing, default it to the working directory.
	if len(build.Context) == 0 {
		build.Context = "."
	}

	buildInfo := fmt.Sprintf("Building container image %s: context=%s", imageName, build.Context)
	if len(build.Dockerfile) > 0 {
		buildInfo = fmt.Sprintf("%s, dockerfile=%s", buildInfo, build.Dockerfile)
	}
	if len(build.Args) > 0 {
		args, err := json.Marshal(build.Args)
		if err != nil {
			return "", nil, err
		}
		buildInfo = fmt.Sprintf("%s, args=%s", buildInfo, args)
	}
	if len(build.Target) > 0 {
		buildInfo = fmt.Sprintf("%s, target=%s", buildInfo, build.Target)
	}
	logging.InitLogging(true, 7, false)
	logging.Infof(buildInfo)

	// If the container build specified build stages to cache, build each in turn.
	var stages []string
	if build.CacheFrom != nil && build.CacheFrom.Stages != nil {
		for _, stage := range stages {
			err := runDockerBuild(localStageImageName(imageName, stage), &build, cacheFrom, logResource, stage)
			if err != nil {
				return "", nil, err
			}
			stages = append(stages, stage)
		}
	}

	// Invoke Docker CLI commands to build.
	err := runDockerBuild(imageName, &build, cacheFrom, logResource, "")

	// Finally, inspect the image so we can return the SHA digest. Do not forward the output of this
	// command this to the CLI to show the user.
	inspectResult, err := runBasicCommandThatMustSucceed("docker", []string{"image", "inspect", "-f", "{{.Id}}", imageName}, logResource)
	if err != nil {
		return "", nil, err
	}
	if len(inspectResult) == 0 {
		return "", nil, errors.Errorf("No digest available for image %s", imageName)
	}

	// From https://docs.docker.com/registry/spec/api/#content-digests
	//
	// the image id will be a "algorithm:hex" pair.  We don't care about the algorithm part.  All we
	// want is the unique portion we can use elsewhere.  Since we are also going to place this in an
	// image tag, we also don't want the colon, as that's not legal there.  So simply grab the hex
	// portion after the colon and return that.
	imageID := strings.TrimSpace(inspectResult)
	colonIndex := strings.LastIndex(imageID, ":")
	if colonIndex > 0 {
		imageID = imageID[colonIndex+1:]
	}

	return imageID, stages, nil
}

type loginResult struct {
	registryName string
	username     string
}

var loginResults []loginResult = nil

func loginToRegistry(registry imageRegistry, logResource pulumi.Resource) error {
	registryName := registry.Server
	username := registry.Username
	password := registry.Password

	// See if we've issued an outstanding requests to login into this registry.  If so, just
	// await the results of that login request.  Otherwise, create a new request and keep it
	// around so that future login requests will see it.
	for _, loginResult := range loginResults {
		if loginResult.registryName == registryName && loginResult.username == username {
			logging.Infof("Reusing existing login for %s@%s", username, registryName)
			return nil
		}
	}

	dockerPasswordStdin, err := useDockerPasswordStdin(logResource)
	if err != nil {
		return err
	}

	// pass 'reportFullCommandLine: false' here so that if we fail to login we don't emit the
	// username/password in our logs.  Instead, we'll just say "'docker login' failed with code ..."
	if dockerPasswordStdin {
		_, err := runCommandThatMustSucceed("docker", []string{"login", "-u", username,
			"--password-stdin", registryName}, logResource, false, password, nil)
		if err != nil {
			return err
		}

	} else {
		_, err := runCommandThatMustSucceed("docker", []string{"login", "-u", username,
			"-p", password, registryName}, logResource, false, "", nil)
		if err != nil {
			return err
		}
	}

	loginResults = append(loginResults, loginResult{
		registryName: registryName,
		username:     username,
	})
	return nil
}

var dockerPassword *bool = nil

func useDockerPasswordStdin(logResource pulumi.Resource) (bool, error) {
	if dockerPassword != nil {
		return *dockerPassword, nil
	}

	// Verify that 'docker' is on the PATH and get the client/server versions
	dockerVersionString, err := runBasicCommandThatMustSucceed("docker", []string{"version", "-f", "{{json .}}"}, logResource)
	if err != nil {
		return false, errors.Wrap(err, "No 'docker' command available on PATH: Please install to use container 'build' mode.")
	}
	logging.Infof("dockerVersion => %s", dockerVersionString)

	// Decide whether to use --password or --password-stdin based on the client version.
	var versionInterface interface{}
	err = json.Unmarshal([]byte(dockerVersionString), &versionInterface)
	if err != nil {
		return false, err
	}

	versionData := versionInterface.(map[string]interface{})
	clientData := versionData["Client"].(map[string]interface{})
	version := clientData["Version"].(string)
	clientVersion, err := semver.NewVersion(version)
	if err != nil {
		return false, err
	}

	constraint, err := semver.NewConstraint(">= 17.07.0")
	if err != nil {
		return false, err
	}

	return constraint.Check(clientVersion), nil
}

func runDockerBuild(imageName string, build *dockerBuild, cacheFrom []string,
	logResource pulumi.Resource, target string) error {

	// Prepare the build arguments.
	buildArgs := []string{"build"}
	if build != nil {
		// Add a custom Dockerfile location.
		buildArgs = append(buildArgs, "-f", build.Dockerfile)
	}
	if build.Args != nil {
		for k, v := range build.Args {
			buildArgs = append(buildArgs, "--build-arg", fmt.Sprintf("%s=%s", k, v))
		}
	}
	if len(build.Target) > 0 {
		buildArgs = append(buildArgs, "--target", build.Target)
	}
	if build.CacheFrom != nil {
		if cacheFrom != nil && len(cacheFrom) > 0 {
			buildArgs = append(buildArgs, "--cache-from", strings.Join(cacheFrom, ""))
		}
	}
	if build.ExtraOptions != nil {
		buildArgs = append(buildArgs, build.ExtraOptions...)
	}
	// Push the docker build context onto the path.
	buildArgs = append(buildArgs, build.Context)

	buildArgs = append(buildArgs, "-t", imageName)
	if len(target) > 0 {
		buildArgs = append(buildArgs, "--target", target)
	}

	_, err := runCommandThatMustSucceed("docker", buildArgs, logResource, true, "", build.Env)
	return err
}

// runCommandThatMustSucceed is used to determine if the full command line should be reported
// when an error happens.  In general reporting the full command line is fine.  But it should be set
// to false if it might contain sensitive information (like a username/password)
func runCommandThatMustSucceed(cmd string, args []string, logResource pulumi.Resource,
	reportFullCommandLine bool, stdin string, env map[string]string) (string, error) {

	stdout, err := runCommandThatCanFail(cmd, args, logResource, reportFullCommandLine, false, stdin, env)

	if err != nil {
		return "", errors.Wrapf(err, "%s failed with error",
			getCommandLineMessage(cmd, args, reportFullCommandLine, env))
	}

	return stdout, nil
}

func runBasicCommandThatMustSucceed(cmd string, args []string, logResource pulumi.Resource) (string, error) {
	return runCommandThatMustSucceed(cmd, args, logResource, true, "", nil)
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
//
// The promise returned by this function should never reach the rejected state.  Even if the
// underlying spawned command has a problem, this will result in a resolved promise with the
// [CommandResult.code] value set to a non-zero value.
func runCommandThatCanFail(cmdName string, args []string, logResource pulumi.Resource, reportFullCommandLine bool,
	reportErrorAsWarning bool, stdinInput string, env map[string]string) (string, error) {

	logging.InitLogging(true, 7, false)

	// Let the user ephemerally know the command we're going to execute.
	logging.Infof("Executing " + getCommandLineMessage(cmdName, args, reportFullCommandLine, env))

	// Generate a unique stream-ID that we'll associate all the docker output with. This will allow
	// each spawned CLI command's output to associated with 'resource' and also streamed to the UI
	// in pieces so that it can be displayed live.  The stream-ID is so that the UI knows these
	// messages are all related and should be considered as one large message (just one that was
	// sent over in chunks).
	//
	// We use rand here in case our package is loaded multiple times in memory (i.e. because
	// different downstream dependencies depend on different versions of us).  By being random we
	// effectively make it completely unlikely that any two cli outputs could map to the same stream
	// id.
	//
	// Pick a reasonably distributed number between 0 and 2^30.  This will fit as an int32
	// which the grpc layer needs.
	streamID := rand.Int()

	cmd := exec.Command(cmdName, args...)

	stderr, err := cmd.StderrPipe()
	if err != nil {
		logging.Errorf("Error retrieving stderr: %v", err)
		return "", err
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		logging.Errorf("Error retrieving stdout: %v", err)
		return "", err
	}

	if len(stdinInput) > 0 {
		stdin, err := cmd.StdinPipe()
		if err != nil {
			logging.Errorf("Error retreiving stdin: %v", err)
			return "", err
		}
		go func() {
			defer stdin.Close()
			io.WriteString(stdin, stdinInput)
		}()
	}

	err = cmd.Start()
	if err != nil {
		logging.Errorf("Error starting command: %v", err)
		return "", err
	}

	stderrBytes, err := ioutil.ReadAll(stderr)
	if err != nil {
		logging.Errorf("Error reading from stderr: %v", err)
		return "", err
	}

	// We can't stream these stderr messages as we receive them because we don't knows at
	// this point because Docker uses stderr for both errors and warnings.  So, instead, we
	// just collect the messages, and wait for the process to end to decide how to report
	// them.
	stderrString := string(stderrBytes)

	// We store the results from stdout and will return them as a string.
	stdoutBytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		logging.Errorf("Error reading from stdout: %v", err)
		return "", err
	}

	// Report all stdout messages as ephemeral messages.  That way they show up in the
	// info bar as they're happening.  But they do not overwhelm the user as the end
	// of the run.
	stdoutString := string(stdoutBytes)
	logging.Infof(stdoutString)

	err = cmd.Wait()

	// If we got any stderr messages, report them as an error/warning depending on the
	// result of the operation.
	if len(stderrString) > 0 {
		if err != nil && !reportErrorAsWarning {
			// Command returned non-zero code.  Treat these stderr messages as an error.
			logging.Errorf("%s (%d)", stderrString, streamID)
		} else {
			// Command succeeded.  These were just a warning.
			logging.Warningf("%s (%d)", stderrString, streamID)
		}
	}

	// If the command failed report an ephemeral message indicating which command it was.
	// That way the user can immediately see something went wrong in the info bar.  The
	// caller (normally runCommandThatMustSucceed) can choose to also report this
	// non-ephemerally.
	if err != nil {
		logging.Infof("%s failed", getCommandLineMessage(cmdName, args, reportFullCommandLine, env))
	}

	return stdoutString, err
}

func getCommandLineMessage(cmd string, args []string, reportFullCommandLine bool, env map[string]string) string {
	argString := ""
	if len(args) > 0 {
		argString = args[0]
		if reportFullCommandLine {
			argString = strings.Join(args, " ")
		}
	}

	if env != nil {
		envString := ""
		for k, v := range env {
			envString = envString + fmt.Sprintf("%s=%s", k, v)
		}
		return fmt.Sprintf("%s %s %s", envString, cmd, argString)
	}
	return fmt.Sprintf("%s %s", cmd, argString)
}
