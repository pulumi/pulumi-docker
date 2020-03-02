// Package docker exports types, functions, subpackages for provisioning docker resources.//
// A Pulumi package for interacting with Docker in Pulumi programs
//
// > This provider is a derived work of the [Terraform Provider](https://github.com/terraform-providers/terraform-provider-docker)
// > distributed under [MPL 2.0](https://www.mozilla.org/en-US/MPL/2.0/). If you encounter a bug or missing feature,
// > first check the [`pulumi/pulumi-docker` repo](https://github.com/pulumi/pulumi-docker/issues); however, if that doesn't turn up anything,
// > please consult the source [`terraform-providers/terraform-provider-docker` repo](https://github.com/terraform-providers/terraform-provider-docker/issues).
//
// nolint: lll
package docker

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os/exec"
	"strings"

	"github.com/pulumi/pulumi/pkg/util/logging"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type commandResult struct {
	code   int
	stdout string
	err    error
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
	reportErrorAsWarning bool, env map[string]string) <-chan commandResult {

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

	r := make(chan commandResult)
	go func() {
		cmd := exec.Command(cmdName, args...)

		stderr, err := cmd.StderrPipe()
		if err != nil {
			logging.Errorf("Error retrieving stderr: %v", err)
			r <- commandResult{err: err}
			return
		}
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			logging.Errorf("Error retrieving stdout: %v", err)
			r <- commandResult{err: err}
			return
		}

		err = cmd.Start()
		if err != nil {
			logging.Errorf("Error starting command: %v", err)
			r <- commandResult{err: err}
			return
		}

		// We store the results from stdout and stderr and will return them as a string.
		stderrBytes, err := ioutil.ReadAll(stderr)
		if err != nil {
			logging.Errorf("Error reading from stderr: %v", err)
			r <- commandResult{err: err}
			return
		}
		stdoutBytes, err := ioutil.ReadAll(stdout)
		if err != nil {
			logging.Errorf("Error reading from stdout: %v", err)
			r <- commandResult{err: err}
			return
		}

		err = cmd.Wait()
		var code int = 0
		// If err is not nil, then there was a non-zero exit code.
		if err != nil {
			code = err.(*exec.ExitError).ExitCode()
		}

		stderrString := string(stderrBytes)
		stdoutString := string(stdoutBytes)

		// If we got any stderr messages, report them as an error/warning depending on the
		// result of the operation.
		if len(stderrString) > 0 {
			if code != 0 && !reportErrorAsWarning {
				logging.Errorf("%s (%d)", stderrString, streamID)
			} else {
				logging.Warningf("%s (%d)", stderrString, streamID)
			}
		}

		r <- commandResult{code: code, stdout: stdoutString}

	}()

	return r
}

func getCommandLineMessage(cmd string, args []string, reportFullCommandLine bool, env map[string]string) string {
	var argString string
	if reportFullCommandLine {
		argString = strings.Join(args, " ")
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
