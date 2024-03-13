package internal

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/docker/buildx/commands"
	"github.com/docker/cli/cli-plugins/manager"
	"github.com/docker/cli/cli/command"
	cfgtypes "github.com/docker/cli/cli/config/types"
	"github.com/docker/cli/cli/streams"
	"github.com/moby/buildkit/client"
	cp "github.com/otiai10/copy"
	"github.com/regclient/regclient"
	"github.com/regclient/regclient/config"

	provider "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi/sdk/v3/go/common/diag"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

// cli wraps a DockerCLI instance with scoped auth credentials. It satisfies
// the Cli interface so it can be used with Docker's cobra.Commands directly.
//
// It buffers stdout/stderr, and layers temporary auth configs on top of the
// host's existing auth.
type cli struct {
	command.Cli

	auths map[string]cfgtypes.AuthConfig
	host  *host

	in   string        // stdin
	r, w *os.File      // stdout
	err  bytes.Buffer  // stderr
	done chan struct{} // signaled when all logs have been forwarded to the engine.
}

// wrap creates a new cli client with auth configs layered on top of our host's
// auth.
func wrap(host *host, registries ...RegistryAuth) (*cli, error) {
	// We need to create a new DockerCLI instance because we don't want the
	// auth changes we make to the ConfigFile to leak to the host.
	docker, err := newDockerCLI(host.config)
	if err != nil {
		return nil, err
	}

	auths := map[string]cfgtypes.AuthConfig{}
	for k, v := range host.auths {
		auths[k] = v
	}

	for _, r := range registries {
		// HostNewName takes care of DockerHub's special-casing for us.
		h := config.HostNewName(r.Address)
		auths[h.CredHost] = cfgtypes.AuthConfig{
			ServerAddress: h.Hostname,
			Username:      r.Username,
			Password:      r.Password,
		}
	}

	// Override our config's auth and disable any credential helpers. Auth
	// lookups will now only return whatever we have in memory.
	cfg := docker.ConfigFile()
	cfg.AuthConfigs = auths
	cfg.CredentialHelpers = nil
	cfg.CredentialsStore = ""

	r, w, err := os.Pipe()
	if err != nil {
		return nil, err
	}

	wrapped := &cli{
		Cli:   docker,
		host:  host,
		auths: auths,
		r:     r,
		w:     w,
	}

	return wrapped, nil
}

func (c *cli) In() *streams.In {
	return streams.NewIn(io.NopCloser(strings.NewReader(c.in)))
}

func (c *cli) Out() *streams.Out {
	return streams.NewOut(c.w)
}

func (c *cli) Err() io.Writer {
	return &c.err
}

// rc returns a registry client with matching auth.
func (c *cli) rc() *regclient.RegClient {
	hosts := []config.Host{}
	for k, v := range c.auths {
		h := config.HostNewName(k)
		h.User = v.Username
		h.Pass = v.Password
		hosts = append(hosts, *h)
	}
	return regclient.New(
		regclient.WithConfigHost(hosts...),
	)
}

// tail is meant to be called as a goroutine and will pipe output from the CLI
// back to the Pulumi engine. Requires a corresponding call to close.
func (c *cli) tail(ctx provider.Context) {
	c.done = make(chan struct{}, 1)
	defer func() {
		c.done <- struct{}{}
		if err := recover(); err != nil {
			fmt.Fprintf(os.Stderr, "recovered: %s\n", err)
		}
	}()

	s := bufio.NewScanner(c.r)
	for s.Scan() {
		ctx.LogStatus(diag.Info, s.Text())
	}
	ctx.LogStatus(diag.Info, "") // clear confusing "DONE" statements.
}

// close flushes any outstanding logs and cleans up resources.
func (c *cli) close() error {
	err := c.w.Close()
	err = errors.Join(err, c.r.Close())
	if c.done != nil {
		<-c.done
	}
	return err
}

// execBuild performs a build by os.Exec'ing the docker-buildx binary.
// Credentials are communicated to docker-buildx via a temporary directory.
// Secrets are communicated via dynamic environment variables.
func (c *cli) execBuild(b Build) (*client.SolveResponse, error) {
	// Setup a temporary directory for auth, and clean it up when we're done.
	tmp, err := os.MkdirTemp("", "pulumi-docker-")
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(tmp)

	opts := b.BuildOptions()

	builder, err := c.host.builderFor(b)
	if err != nil {
		return nil, err
	}

	// Docker expects a "$DOCKER_CONFIG/contexts" directory in addition to
	// "$DOCKER_CONFIG/config.json", so we attempt to copy this from the host
	// to our temporary directory.
	hostConfigDir := filepath.Dir(c.ConfigFile().Filename)
	err = cp.Copy(
		filepath.Join(hostConfigDir, "contexts"),
		filepath.Join(tmp, "contexts"),
	)
	if err != nil {
		return nil, err
	}

	// Save our temporary credentials to $tmp/config.json.
	tmpCfg := filepath.Join(tmp, filepath.Base(c.ConfigFile().Filename))
	c.ConfigFile().Filename = tmpCfg
	err = c.ConfigFile().Save()
	if err != nil {
		return nil, err
	}

	// We will spawn docker-buildx with DOCKER_CONFIG set to our temporary
	// directory for auth, but BUILDX_CONFIG will point to the host. There's a
	// bunch of builder state in there that we want to preserve.
	env := []string{
		fmt.Sprintf("DOCKER_CONFIG=%s", tmp),
		fmt.Sprintf("BUILDX_CONFIG=%s", filepath.Join(hostConfigDir, "buildx")),
	}

	// We need to write to this file in order to recover information about the
	// build, like the digest.
	metadata := filepath.Join(tmp, "metadata.json")
	args := []string{
		"buildx",
		"build",
		"--progress", "plain",
		"--metadata-file", metadata,
		"--builder", builder.name,
	}

	// TODO: --allow
	// TODO: --annotation
	// TODO: --attest
	// TODO: --cgroup-parent

	for k, v := range opts.BuildArgs {
		args = append(args, "--build-arg", fmt.Sprintf("%s=%s", k, v))
	}
	if opts.Builder != "" {
		args = append(args, "--builder", opts.Builder)
	}
	for _, c := range opts.CacheFrom {
		args = append(args, "--cache-from", attrcsv(c.Type, c.Attrs))
	}
	for _, c := range opts.CacheTo {
		args = append(args, "--cache-to", attrcsv(c.Type, c.Attrs))
	}
	if opts.ExportLoad {
		args = append(args, "--load")
	}
	if opts.ExportPush {
		args = append(args, "--push")
	}
	for _, e := range opts.Exports {
		args = append(args, "--output", attrcsv(e.Type, e.Attrs))
	}
	for _, h := range opts.ExtraHosts {
		args = append(args, "--add-host", h)
	}
	for k, v := range opts.NamedContexts {
		args = append(args, "--build-context", fmt.Sprintf("%s=%s", k, v))
	}
	for k, v := range opts.Labels {
		args = append(args, "--label", fmt.Sprintf("%s=%s", k, v))
	}
	if opts.NetworkMode != "" {
		args = append(args, "--network", opts.NetworkMode)
	}
	if opts.NoCache {
		args = append(args, "--no-cache")
	}
	for _, p := range opts.Platforms {
		args = append(args, "--platform", p)
	}
	if opts.Pull {
		args = append(args, "--pull")
	}
	for _, ssh := range opts.SSH {
		s := ssh.ID
		if len(ssh.Paths) > 0 {
			s += "=" + strings.Join(ssh.Paths, ",")
		}
		args = append(args, "--ssh", s)
	}
	for _, t := range opts.Tags {
		args = append(args, "--tag", t)
	}
	if opts.Target != "" {
		args = append(args, "--target", opts.Target)
	}
	if opts.DockerfileName != "" {
		args = append(args, "-f", opts.DockerfileName)
	}
	if in := b.Inline(); in != "" {
		c.in = in
		args = append(args, "-f", "-")
	}
	if opts.ContextPath != "" {
		args = append(args, opts.ContextPath)
	}

	// We pass secrets by value via dynamic PULUMI_DOCKER_* environment
	// variables.
	for _, s := range opts.Secrets {
		envvar, err := resource.NewUniqueHex("PULUMI_DOCKER_", 0, 0)
		if err != nil {
			return nil, err
		}
		// We abuse the pb.Secret proto by stuffing the secret's value in
		// XXX_unrecognized. We never serialize this proto so this is tolerable.
		env = append(env, fmt.Sprintf("%s=%s", envvar, s.XXX_unrecognized))
		args = append(args, "--secret", fmt.Sprintf("id=%s,env=%s", s.ID, envvar))
	}

	// Invoke docker-buildx.
	err = c.exec(args, env)
	if err != nil {
		return nil, err
	}

	// Read the metadata file and transform it back into the map[string]string
	// structure originally returned by the exporter.
	_, err = os.Stat(metadata)
	if err != nil {
		return nil, fmt.Errorf("missing metadata: %w", err)
	}
	out, err := os.ReadFile(metadata)
	if err != nil {
		return nil, err
	}
	var raw map[string]any
	err = json.Unmarshal(out, &raw)
	if err != nil {
		return nil, err
	}
	resp := map[string]string{}
	for k, v := range raw {
		switch vv := v.(type) {
		case string:
			resp[k] = vv
		default:
			out, err := json.Marshal(v)
			if err != nil {
				continue
			}
			resp[k] = string(out)
		}
	}

	return &client.SolveResponse{ExporterResponse: resp}, nil
}

// exec invokes a Docker plugin binary. The first argument should be the name
// of the plugin's subcommand, e.g. "buildx".
func (c *cli) exec(args []string, extraEnv []string) error {
	if len(args) == 0 {
		return fmt.Errorf("args must be non-empty")
	}
	name := args[0]

	root := commands.NewRootCmd(name, false, c)
	plug, err := manager.GetPlugin(name, c, root)
	if err != nil {
		return err
	}
	if plug.Err != nil {
		return plug.Err
	}

	defer c.w.Close()

	cmd, err := manager.PluginRunCommand(c, name, root)
	if err != nil {
		return err
	}
	cmd.Args = append([]string{cmd.Args[0]}, args...)
	cmd.Stderr = c.Err() // TODO: This is build output...
	cmd.Stdout = c.Out()
	cmd.Stdin = c.In()

	cmd.Env = append(cmd.Env, extraEnv...)

	return cmd.Run()
}

// attrcsv transforms key/values into a CSV: key1=value1,key2=value2,...
func attrcsv(typ string, m map[string]string) string {
	s := []string{fmt.Sprintf("type=%s", typ)}
	for k, v := range m {
		s = append(s, fmt.Sprintf("%s=%s", k, v))
	}
	return strings.Join(s, ",")
}
