// Copyright 2016-2020, Pulumi Corporation.

package docker

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"os/exec"
	"strings"

	"github.com/Masterminds/semver"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func buildAndPushImage(ctx *pulumi.Context, baseImageName string, build *DockerBuild,
	repositoryURL string, logResource pulumi.Resource, skipPush bool, registry *ImageRegistry) (string, error) {

	// Give an initial message indicating what we're about to do.  That way, if anything
	// takes a while, the user has an idea about what's going on.
	logDebug(ctx, "Starting docker build and push...", logResource)

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
	if registry.Server != "" {
		if !ctx.DryRun() || pullFromCache {
			logDebug(ctx, "Logging into registry...", logResource)
			err := loginToRegistry(ctx, *registry, logResource)
			if err != nil {
				return "", err
			}
		}
	}

	// If the container specified a cacheFrom parameter, first set up the cached stages.
	var cacheFrom []string
	if pullFromCache {
		cacheFrom = pullCache(ctx, baseImageName, *build.CacheFrom, repositoryURL, logResource)
	}

	// If the build context is missing, default it to the working directory.
	if build.Context == "" {
		build.Context = "."
	}

	// Next, build the image.
	logEphemeral(ctx, fmt.Sprintf("Building image '%s'...", build.Context), logResource)
	imageID, stages, err := buildImage(ctx, baseImageName, build, logResource, cacheFrom)
	if err != nil {
		return "", err
	}
	logEphemeral(ctx, "Image build succeeded.", logResource)

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
		logEphemeral(ctx, fmt.Sprintf("Pushing image '%s'...", baseImageName), logResource)

		// First, push with both the optionally-requested-tag *and* imageId (which is guaranteed to
		// be defined).  By using the imageId we give the image a fully unique location that we can
		// successfully pull regardless of whatever else has happened at this repositoryUrl.
		err := tagAndPushImage(ctx, baseImageName, repositoryURL, tag, imageID, logResource)
		if err != nil {
			return "", err
		}

		// Next, push only with the optionally-requested-tag.  Users of this API still want to get a
		// nice and simple url that they can reach this image at, without having the explicit imageId
		// hash added to it.  Note: this location is not guaranteed to be idempotent.  For example,
		// pushes on other machines might overwrite that location.
		err = tagAndPushImage(ctx, baseImageName, repositoryURL, tag, "", logResource)
		if err != nil {
			return "", err
		}

		for _, stage := range stages {
			err = tagAndPushImage(ctx, localStageImageName(baseImageName, stage),
				repositoryURL, stage, "", logResource)
			if err != nil {
				return "", err
			}
		}

		logEphemeral(ctx, "Image push succeeded.", logResource)
	}

	return uniqueTaggedImageName, nil
}

func pullCache(ctx *pulumi.Context, imageName string, cacheFrom CacheFrom, repoURL string, logResource pulumi.Resource) []string { // nolint:lll
	if len(repoURL) == 0 {
		return nil
	}

	logDebug(ctx, fmt.Sprintf("Pulling cache for %s from %s", imageName, repoURL), logResource)

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
		_, err := runCommandThatCanFail(ctx, "docker", []string{"pull", image}, logResource, true, true, "", nil)
		if err == nil {
			continue
		}

		cacheFromImages = append(cacheFromImages, image)
	}

	return cacheFromImages
}

func tagAndPushImage(ctx *pulumi.Context, imageName string, repositoryURL string, tag string,
	imageID string, logResource pulumi.Resource) error {

	doTagAndPush := func(targetName string) error {
		_, err := runBasicCommandThatMustSucceed(ctx, "docker", []string{"tag", imageName, targetName}, logResource)
		if err != nil {
			return err
		}
		_, err = runBasicCommandThatMustSucceed(ctx, "docker", []string{"push", targetName}, logResource)
		return err
	}

	// Ensure we have a unique target name for this image, and tag and push to that unique target.
	err := doTagAndPush(createTaggedImageName(repositoryURL, tag, imageID))
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
		err := doTagAndPush(createTaggedImageName(repositoryURL, tag, ""))
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

// Note: unlike the Typescript and Dotnet implementations, you must pass in a DockerBuild here.
// If you have just the path, `build = &DockerBuild{Context: path}`.
func buildImage(ctx *pulumi.Context, imageName string, build *DockerBuild,
	logResource pulumi.Resource, cacheFrom []string) (string, []string, error) {

	buildInfo := fmt.Sprintf("Building container image %s: context=%s", imageName, build.Context)
	if build.Dockerfile != "" {
		buildInfo = fmt.Sprintf("%s, dockerfile=%s", buildInfo, build.Dockerfile)
	}
	if len(build.Args) > 0 {
		args, err := json.Marshal(build.Args)
		if err != nil {
			return "", nil, err
		}
		buildInfo = fmt.Sprintf("%s, args=%s", buildInfo, args)
	}
	if build.Target != "" {
		buildInfo = fmt.Sprintf("%s, target=%s", buildInfo, build.Target)
	}
	logDebug(ctx, buildInfo, logResource)

	// If the container build specified build stages to cache, build each in turn.
	var stages []string
	if build.CacheFrom != nil && build.CacheFrom.Stages != nil {
		for _, stage := range build.CacheFrom.Stages {
			err := runDockerBuild(ctx, localStageImageName(imageName, stage), build, cacheFrom, logResource, stage)
			if err != nil {
				return "", nil, err
			}
			stages = append(stages, stage)
		}
	}

	// Invoke Docker CLI commands to build.
	err := runDockerBuild(ctx, imageName, build, cacheFrom, logResource, "")
	if err != nil {
		return "", nil, err
	}

	// Finally, inspect the image so we can return the SHA digest. Do not forward the output of this
	// command this to the CLI to show the user.
	inspectResult, err := runBasicCommandThatMustSucceed(ctx, "docker",
		[]string{"image", "inspect", "-f", "{{.Id}}", imageName}, logResource)
	if err != nil {
		return "", nil, err
	}
	if inspectResult == "" {
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

func loginToRegistry(ctx *pulumi.Context, registry ImageRegistry, logResource pulumi.Resource) error {
	registryName := registry.Server
	username := registry.Username
	password := registry.Password

	// See if we've issued an outstanding requests to login into this registry.  If so, just
	// await the results of that login request.  Otherwise, create a new request and keep it
	// around so that future login requests will see it.
	for _, loginResult := range loginResults {
		if loginResult.registryName == registryName && loginResult.username == username {
			logDebug(ctx, fmt.Sprintf("Reusing existing login for %s@%s", username, registryName), logResource)
			return nil
		}
	}

	dockerPasswordStdin, err := useDockerPasswordStdin(ctx, logResource)
	if err != nil {
		return err
	}

	// pass 'reportFullCommandLine: false' here so that if we fail to login we don't emit the
	// username/password in our logs.  Instead, we'll just say "'docker login' failed with code ..."
	if dockerPasswordStdin {
		_, err := runCommandThatMustSucceed(ctx, "docker", []string{"login", "-u", username,
			"--password-stdin", registryName}, logResource, false, password, nil)
		if err != nil {
			return err
		}

	} else {
		_, err := runCommandThatMustSucceed(ctx, "docker", []string{"login", "-u", username,
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

func useDockerPasswordStdin(ctx *pulumi.Context, logResource pulumi.Resource) (bool, error) {
	if dockerPassword != nil {
		return *dockerPassword, nil
	}

	// Verify that 'docker' is on the PATH and get the client/server versions
	dockerVersionString, err := runBasicCommandThatMustSucceed(ctx, "docker",
		[]string{"version", "-f", "{{json .}}"}, logResource)
	if err != nil {
		return false, errors.Wrap(err, "no 'docker' command available on PATH: Please install to use container 'build' mode.")
	}
	logDebug(ctx, fmt.Sprintf("'docker version' => %s", dockerVersionString), logResource)

	// Decide whether to use --password or --password-stdin based on the client version.
	var versionInterface interface{}
	err = json.Unmarshal([]byte(dockerVersionString), &versionInterface)
	if err != nil {
		return false, errors.Wrapf(err, "unable to parse docker version output '%s'", dockerVersionString)
	}

	versionData, ok := versionInterface.(map[string]interface{})
	if !ok {
		return false, errors.New("unable to extract the docker version")
	}

	clientData, ok := versionData["Client"].(map[string]interface{})
	if !ok {
		return false, errors.New("unable to extract the docker client version")
	}

	version, ok := clientData["Version"].(string)
	if !ok {
		return false, errors.New("unable to extract the docker client version")
	}

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

func runDockerBuild(ctx *pulumi.Context, imageName string, build *DockerBuild, cacheFrom []string,
	logResource pulumi.Resource, target string) error {
	if build == nil {
		return fmt.Errorf("build field required for running docker build")
	}
	// Prepare the build arguments.
	buildArgs := []string{"build"}
	buildArgs = append(buildArgs, "-f", build.Dockerfile)
	if build.Args != nil {
		for k, v := range build.Args {
			buildArgs = append(buildArgs, "--build-arg", fmt.Sprintf("%s=%s", k, v))
		}
	}
	if build.Target != "" {
		buildArgs = append(buildArgs, "--target", build.Target)
	}
	if build.CacheFrom != nil {
		if len(cacheFrom) > 0 {
			buildArgs = append(buildArgs, "--cache-from", strings.Join(cacheFrom, ""))
		}
	}
	if build.ExtraOptions != nil {
		buildArgs = append(buildArgs, build.ExtraOptions...)
	}
	// Push the docker build context onto the path.
	buildArgs = append(buildArgs, build.Context)

	buildArgs = append(buildArgs, "-t", imageName)
	if target != "" {
		buildArgs = append(buildArgs, "--target", target)
	}

	_, err := runCommandThatMustSucceed(ctx, "docker", buildArgs, logResource, true, "", build.Env)
	return err
}

// runCommandThatMustSucceed is used to determine if the full command line should be reported
// when an error happens.  In general reporting the full command line is fine.  But it should be set
// to false if it might contain sensitive information (like a username/password)
func runCommandThatMustSucceed(ctx *pulumi.Context, cmd string, args []string, logResource pulumi.Resource,
	reportFullCommandLine bool, stdin string, env map[string]string) (string, error) {

	stdout, err := runCommandThatCanFail(ctx, cmd, args, logResource, reportFullCommandLine, false, stdin, env)

	if err != nil {
		return "", errors.Wrapf(err, "%s failed with error",
			getCommandLineMessage(cmd, args, reportFullCommandLine, env))
	}

	return stdout, nil
}

func runBasicCommandThatMustSucceed(ctx *pulumi.Context, cmd string, args []string,
	logResource pulumi.Resource) (string, error) {

	return runCommandThatMustSucceed(ctx, cmd, args, logResource, true, "", nil)
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
func runCommandThatCanFail(ctx *pulumi.Context, cmdName string, args []string, logResource pulumi.Resource,
	reportFullCommandLine bool, reportErrorAsWarning bool, stdinInput string, env map[string]string) (string, error) {

	// Let the user ephemerally know the command we're going to execute.
	logDebug(ctx, "Executing "+getCommandLineMessage(cmdName, args, reportFullCommandLine, env), logResource)

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
	streamID := rand.Int31() // #nosec

	// nolint:errcheck
	logErrorf := func(format string, a ...interface{}) {
		ctx.Log.Error(fmt.Sprintf(format, a...), &pulumi.LogArgs{
			Resource: logResource,
			StreamID: streamID,
		})
	}

	cmd := exec.Command(cmdName, args...)

	stderr, err := cmd.StderrPipe()
	if err != nil {
		logErrorf("Error retrieving stderr: %v", err)
		return "", err
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		logErrorf("Error retrieving stdout: %v", err)
		return "", err
	}

	if stdinInput != "" {
		stdin, err := cmd.StdinPipe()
		if err != nil {
			logErrorf("Error retrieving stdin: %v", err)
			return "", err
		}
		go func() {
			defer stdin.Close()
			// nolint:errcheck
			io.WriteString(stdin, stdinInput)
		}()
	}

	err = cmd.Start()
	if err != nil {
		logErrorf("Error starting command: %v", err)
		return "", err
	}

	// Kick off some goroutines to stream the output asynchronously. We store the results from stdout in
	// memory and will return them as a string.
	var stdoutBuffer bytes.Buffer
	var stderrBuffer bytes.Buffer
	go func() {
		outs := bufio.NewScanner(stdout)
		for outs.Scan() {
			text := outs.Text()
			stdoutBuffer.WriteString(text + "\n")
			// Also report all stdout messages as ephemeral messages.  That way they show up in the
			// info bar as they're happening.  But they do not overwhelm the user as the end
			// of the run.
			logEphemeral(ctx, text, logResource)

		}
	}()
	go func() {
		errs := bufio.NewScanner(stderr)
		for errs.Scan() {
			text := errs.Text()
			stderrBuffer.WriteString(text + "\n")
		}
	}()

	// Wait for the command to exit
	err = cmd.Wait()

	// If we got any stderr messages, report them as an error/warning depending on the
	// result of the operation.
	stderrString := stderrBuffer.String()
	if stderrString != "" {
		if err != nil && !reportErrorAsWarning {
			// Command returned non-zero code.  Treat these stderr messages as an error.
			logErrorf(stderrString)
		} else {
			// Command succeeded.  These were just a warning.
			// nolint:errcheck
			ctx.Log.Warn(stderrString, &pulumi.LogArgs{
				Resource:  logResource,
				StreamID:  streamID,
				Ephemeral: true,
			})
		}
	}

	// If the command failed report an ephemeral message indicating which command it was.
	// That way the user can immediately see something went wrong in the info bar.  The
	// caller (normally runCommandThatMustSucceed) can choose to also report this
	// non-ephemerally.
	if err != nil {
		logEphemeral(ctx, getCommandLineMessage(cmdName, args, reportFullCommandLine, env)+" failed", logResource)
	}

	stdoutString := stdoutBuffer.String()
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

// nolint:errcheck
func logEphemeral(ctx *pulumi.Context, msg string, logResource pulumi.Resource) {
	ctx.Log.Info(msg, &pulumi.LogArgs{
		Resource:  logResource,
		Ephemeral: true,
	})
}

// nolint:errcheck
func logDebug(ctx *pulumi.Context, msg string, logResource pulumi.Resource) {
	ctx.Log.Debug(msg, &pulumi.LogArgs{
		Resource:  logResource,
		Ephemeral: true,
	})
}
