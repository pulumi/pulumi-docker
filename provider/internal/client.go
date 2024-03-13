//go:generate go run go.uber.org/mock/mockgen -typed -package internal -source client.go -destination mockclient_test.go --self_package github.com/pulumi/pulumi-docker/provider/v4/internal
package internal

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/distribution/reference"
	buildx "github.com/docker/buildx/build"
	"github.com/docker/buildx/commands"
	controllerapi "github.com/docker/buildx/controller/pb"
	"github.com/docker/buildx/util/dockerutil"
	"github.com/docker/buildx/util/platformutil"
	"github.com/docker/buildx/util/progress"
	"github.com/docker/cli/cli/command"
	"github.com/docker/cli/cli/flags"
	manifesttypes "github.com/docker/cli/cli/manifest/types"
	registryclient "github.com/docker/cli/cli/registry/client"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/registry"
	registryconst "github.com/docker/docker/registry"
	"github.com/moby/buildkit/client"
	"github.com/moby/buildkit/session"
	"github.com/moby/buildkit/session/auth/authprovider"
	"github.com/moby/buildkit/util/progress/progressui"
	"github.com/regclient/regclient/types/errs"
	"github.com/regclient/regclient/types/ref"
	"github.com/sirupsen/logrus"

	provider "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi/sdk/v3/go/common/diag"
)

// Client handles all our Docker API calls.
type Client interface {
	Build(ctx provider.Context, b Build) (*client.SolveResponse, error)
	BuildKitEnabled() (bool, error)
	Inspect(ctx context.Context, id string) ([]manifesttypes.ImageManifest, error)
	Delete(ctx context.Context, id string) ([]image.DeleteResponse, error)

	ManifestCreate(ctx provider.Context, push bool, target string, refs ...string) error
	ManifestInspect(ctx provider.Context, target string) (string, error)
	ManifestDelete(ctx provider.Context, target string) error
}

type Build interface {
	BuildOptions() controllerapi.BuildOptions
	Inline() string
	ShouldExec() bool
	Secrets() session.Attachable
}

var _ Client = (*cli)(nil)

func newDockerCLI(config *Config) (*command.DockerCli, error) {
	cli, err := command.NewDockerCli(
		command.WithDefaultContextStoreConfig(),
		command.WithContentTrustFromEnv(),
	)
	if err != nil {
		return nil, err
	}

	opts := flags.NewClientOptions()
	if config != nil && config.Host != "" {
		opts.Hosts = append(opts.Hosts, config.Host)
	}
	err = cli.Initialize(opts)
	if err != nil {
		return nil, err
	}

	// TODO: Log some version information for debugging.

	// Disable the CLI's tendency to log randomly to stdout.
	logrus.SetOutput(io.Discard)

	return cli, nil
}

// Build performs a BuildKit build. Returns a map of target names (or one name,
// "default", if no targets were specified) to SolveResponses, which capture
// the build's digest and tags (if any).
func (c *cli) Build(
	pctx provider.Context,
	build Build,
) (*client.SolveResponse, error) {
	ctx := context.Context(pctx)
	opts := build.BuildOptions()

	go c.tail(pctx)
	defer c.close()

	if build.ShouldExec() {
		return c.execBuild(build)
	}

	b, err := c.host.builderFor(build)
	if err != nil {
		return nil, err
	}
	printer, err := progress.NewPrinter(ctx, c.w,
		progressui.PlainMode,
		progress.WithDesc(
			fmt.Sprintf("building with %q instance using %s driver", b.name, b.driver),
			fmt.Sprintf("%s:%s", b.driver, b.name),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("creating printer: %w", err)
	}

	cacheFrom := []client.CacheOptionsEntry{}
	for _, c := range opts.CacheFrom {
		cacheFrom = append(cacheFrom, client.CacheOptionsEntry{
			Type:  c.Type,
			Attrs: c.Attrs,
		})
	}
	cacheTo := []client.CacheOptionsEntry{}
	for _, c := range opts.CacheTo {
		cacheTo = append(cacheTo, client.CacheOptionsEntry{
			Type:  c.Type,
			Attrs: c.Attrs,
		})
	}
	exports := []client.ExportEntry{}
	for _, e := range opts.Exports {
		exports = append(exports, client.ExportEntry{
			Type:      e.Type,
			Attrs:     e.Attrs,
			OutputDir: e.Destination,
		})
	}
	platforms, _ := platformutil.Parse(opts.Platforms)
	platforms = platformutil.Dedupe(platforms)

	namedContexts := map[string]buildx.NamedContext{}
	for k, v := range opts.NamedContexts {
		ref, err := reference.ParseNormalizedNamed(k)
		if err != nil {
			return nil, err
		}
		name := strings.TrimSuffix(reference.FamiliarString(ref), ":latest")
		namedContexts[name] = buildx.NamedContext{Path: v}
	}

	ssh, err := controllerapi.CreateSSH(opts.SSH)
	if err != nil {
		return nil, err
	}

	target := opts.Target
	if target == "" {
		target = "default"
	}
	payload := map[string]buildx.Options{
		target: {
			Inputs: buildx.Inputs{
				ContextPath:      opts.ContextPath,
				DockerfilePath:   opts.DockerfileName,
				DockerfileInline: build.Inline(),
				NamedContexts:    namedContexts,
				InStream:         strings.NewReader(""),
			},
			// Disable default provenance for now. Docker's `manifest create`
			// doesn't handle manifests with provenance included; more reason
			// to use imagetools instead.
			Attests:     map[string]*string{"provenance": nil},
			BuildArgs:   opts.BuildArgs,
			CacheFrom:   cacheFrom,
			CacheTo:     cacheTo,
			Exports:     exports,
			ExtraHosts:  opts.ExtraHosts,
			NetworkMode: opts.NetworkMode,
			NoCache:     opts.NoCache,
			Labels:      opts.Labels,
			Platforms:   platforms,
			Pull:        opts.Pull,
			Tags:        opts.Tags,
			Target:      opts.Target,

			Session: []session.Attachable{
				ssh,
				authprovider.NewDockerAuthProvider(c.ConfigFile(), nil),
				build.Secrets(),
			},
		},
	}

	// Perform the build.
	results, err := buildx.Build(
		ctx,
		b.nodes,
		payload,
		dockerutil.NewClient(c),
		filepath.Dir(c.ConfigFile().Filename),
		printer,
	)
	if err != nil {
		return nil, err
	}

	if printErr := printer.Wait(); printErr != nil {
		return results[target], printErr
	}
	for _, w := range printer.Warnings() {
		b := &bytes.Buffer{}
		fmt.Fprintf(b, "%s", w.Short)
		for _, d := range w.Detail {
			fmt.Fprintf(b, "\n%s", d)
		}
		pctx.Log(diag.Warning, b.String())
	}

	return results[target], err
}

// BuildKitEnabled returns true if the client supports buildkit.
func (c *cli) BuildKitEnabled() (bool, error) {
	return c.Cli.BuildKitEnabled()
}

func (c *cli) ManifestCreate(ctx provider.Context, push bool, target string, refs ...string) error {
	// TODO: Create this manifest with regclient or imagetools.

	go c.tail(ctx)
	defer c.close()

	args := []string{
		// "buildx",
		"imagetools",
		"create",
		"--tag", target,
	}

	if !push {
		args = append(args, "--dry-run")
	}

	args = append(args, refs...)

	cmd := commands.NewRootCmd(os.Args[0], false, c)

	cmd.SetArgs(args)
	return cmd.ExecuteContext(ctx)

	/*
		cmd := manifest.NewManifestCommand(c)
		cmd.SilenceUsage = true
		createArgs := []string{"create", target, "--amend"}
		createArgs = append(createArgs, refs...)

		ctx.LogStatus(diag.Info, "Creating manifest...")

		go c.tail(ctx)
		defer c.close()

		cmd.SetArgs(createArgs)
		err := cmd.ExecuteContext(ctx)
		if err != nil {
			return fmt.Errorf("creating: %w", err)
		}

		if !push {
			return nil
		}

		ctx.LogStatus(diag.Info, "Pushing manifest...")
		pushArgs := []string{"push", target}
		cmd.SetArgs(pushArgs)
		err = cmd.ExecuteContext(ctx)
		if err != nil {
			return fmt.Errorf("pushing: %w", err)
		}
		return nil
	*/
}

func (c *cli) ManifestInspect(ctx provider.Context, target string) (string, error) {
	rc := c.rc()

	ref, err := ref.New(target)
	if err != nil {
		return "", err
	}

	m, err := rc.ManifestHead(ctx, ref)
	if err != nil {
		return "", fmt.Errorf("fetching head: %w", err)
	}

	return string(m.GetDescriptor().Digest), nil

	/*
	   cmd := manifest.NewManifestCommand(CLI)
	   inspectArgs := []string{"inspect", target}
	   cmd.SetArgs(inspectArgs)
	   err = cmd.ExecuteContext(ctx)

	   	if err != nil {
	   		return "", fmt.Errorf("inspecting: %w", err)
	   	}

	   data := buf.Bytes()
	   ml := manifestlist.DeserializedManifestList{}
	   err = ml.UnmarshalJSON(data)
	   	if err != nil {
	   		return "", err
	   	}
	   _, payload, _ := ml.Payload()
	   digest := digest.FromBytes(payload)

	   return string(digest), nil
	*/
}

func (c *cli) ManifestDelete(ctx provider.Context, target string) error {
	rc := c.rc()

	ref, err := ref.New(target)
	if err != nil {
		return err
	}

	err = rc.ManifestDelete(context.Context(ctx), ref)
	if errors.Is(err, errs.ErrHTTPStatus) {
		ctx.Log(diag.Warning, "this registry does not support deletions")
		return nil
	}
	if err != nil {
		return fmt.Errorf("fetching head: %w", err)
	}

	return nil
	/*
		go c.tail(ctx)
		defer c.close()

		cmd := manifest.NewManifestCommand(c)
		cmd.SilenceUsage = true
		deleteArgs := []string{"rm", target}

		cmd.SetArgs(deleteArgs)
		err := cmd.ExecuteContext(ctx)

		if strings.Contains(err.Error(), "No such manifest:") {
			return nil
		}

		if err != nil {
			return fmt.Errorf("deleting: %w", err)
		}

		// TODO: Actually delete this with regclient.

		return nil
	*/
}

// Inspect inspects an image.
func (c *cli) Inspect(ctx context.Context, id string) ([]manifesttypes.ImageManifest, error) {
	ref, err := normalizeReference(id)
	if err != nil {
		return []manifesttypes.ImageManifest{}, err
	}

	// Constructed a RegistryClient which can use our in-memory auth.
	insecure := c.DockerEndpoint().SkipTLSVerify
	resolver := func(_ context.Context, index *registry.IndexInfo) registry.AuthConfig {
		configKey := index.Name
		if index.Official {
			configKey = registryconst.IndexServer
		}
		return registry.AuthConfig(c.auths[configKey])
	}
	rc := registryclient.NewRegistryClient(resolver, command.UserAgent(), insecure)

	manifests, err := rc.GetManifestList(ctx, ref)

	// If the registry doesn't support manifest lists, attempt to fetch an
	// individual one.
	if err != nil && strings.Contains(err.Error(), "unsupported manifest format") {
		manifest, err := rc.GetManifest(ctx, ref)
		manifests = append(manifests, manifest)
		return manifests, err
	}

	return manifests, err
}

// Delete deletes an image with the given ID.
func (c *cli) Delete(ctx context.Context, id string) ([]image.DeleteResponse, error) {
	return c.Client().ImageRemove(ctx, id, types.ImageRemoveOptions{
		Force: true, // Needed in case the image has multiple tags.
	})
}

func normalizeReference(ref string) (reference.Named, error) {
	namedRef, err := reference.ParseNormalizedNamed(ref)
	if err != nil {
		return nil, err
	}
	if _, isDigested := namedRef.(reference.Canonical); !isDigested {
		return reference.TagNameOnly(namedRef), nil
	}
	return namedRef, nil
}
