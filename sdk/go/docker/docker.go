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
	CacheFrom pulumi.Input `pulumi:"cacheFrom"`

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
	CacheFrom    interface{}
	ExtraOptions []string
	Env          map[string]string
	Target       string
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
		_, err := RunCommandThatMustSucceed("docker", []string{"login", "-u", username,
			"--password-stdin", registryName}, logResource, false, password, nil)
		if err != nil {
			return err
		}

	} else {
		_, err := RunCommandThatMustSucceed("docker", []string{"login", "-u", username,
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
	dockerVersionString, err := RunCommandThatMustSucceed("docker", []string{"version", "-f", "{{json .}}"}, logResource, true, "", nil)
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

	_, err := RunCommandThatMustSucceed("docker", buildArgs, logResource, true, "", build.Env)
	return err
}

// RunCommandThatMustSucceed is used to determine if the full command line should be reported
// when an error happens.  In general reporting the full command line is fine.  But it should be set
// to false if it might contain sensitive information (like a username/password)
func RunCommandThatMustSucceed(cmd string, args []string, logResource pulumi.Resource,
	reportFullCommandLine bool, stdin string, env map[string]string) (string, error) {

	stdout, err := runCommandThatCanFail(cmd, args, logResource, reportFullCommandLine, false, stdin, env)

	if err != nil {
		return "", errors.Wrapf(err, "%s failed with error",
			getCommandLineMessage(cmd, args, reportFullCommandLine, env))
	}

	return stdout, nil
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
