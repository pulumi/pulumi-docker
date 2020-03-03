package docker

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/pkg/util/logging"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// RunCommandThatMustSucceed is used to determine if the full command line should be reported
// when an error happens.  In general reporting the full command line is fine.  But it should be set
// to false if it might contain sensitive information (like a username/password)
func RunCommandThatMustSucceed(cmd string, args []string, logResource pulumi.Resource,
	reportFullCommandLine bool, env map[string]string) (string, error) {

	stdout, err := runCommandThatCanFail(cmd, args, logResource, reportFullCommandLine, false, env)

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
	reportErrorAsWarning bool, env map[string]string) (string, error) {

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
