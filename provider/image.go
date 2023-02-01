package provider

import (
	"bufio"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/docker/distribution/reference"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
	"golang.org/x/sync/errgroup"
	"io"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/docker/cli/cli/config"
	"github.com/docker/cli/cli/config/credentials"
	clitypes "github.com/docker/cli/cli/config/types"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/jsonmessage"
	structpb "github.com/golang/protobuf/ptypes/struct"
	controlapi "github.com/moby/buildkit/api/services/control"
	buildkitclient "github.com/moby/buildkit/client"
	dockerfile "github.com/moby/buildkit/frontend/dockerfile/builder"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/ryboe/q"
)

const defaultDockerfile = "Dockerfile"
const defaultBuilder = "2"

type Image struct {
	Name     string
	SkipPush bool
	Registry Registry
	Tag      string
	Build    Build
}

type Registry struct {
	Server   string
	Username string
	Password string
}

type Build struct {
	Context        string
	Dockerfile     string
	CachedImages   []string
	Args           map[string]*string
	Target         string
	Platform       string
	BuilderVersion types.BuilderVersion
}

func (p *dockerNativeProvider) dockerBuild(ctx context.Context,
	urn resource.URN,
	props *structpb.Struct) (string, *structpb.Struct, error) {

	inputs, err := plugin.UnmarshalProperties(props, plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true})
	if err != nil {
		return "", nil, err
	}
	reg := marshalRegistry(inputs["registry"])
	skipPush := marshalSkipPush(inputs["skipPush"])
	// read in values to Image
	img := Image{
		Name:     inputs["imageName"].StringValue(),
		SkipPush: skipPush,
		Registry: reg,
	}

	build, err := marshalBuildAndApplyDefaults(inputs["build"])
	if err != nil {
		return "", nil, err
	}
	cache := marshalCachedImages(img, inputs["build"])

	build.CachedImages = cache
	img.Build = build

	docker, err := client.NewClientWithOpts(client.FromEnv)

	if err != nil {
		return "", nil, err
	}

	err = p.host.LogStatus(ctx, "info", urn, "Building the image")

	if err != nil {
		return "", nil, err
	}

	// make the build context
	//tar, err := archive.TarWithOptions(img.Build.Context, &archive.TarOptions{})
	//if err != nil {
	//	return "", nil, err
	//}
	//q.Q("attempt to load the tar file")
	//q.Q(tar)
	//tarReadCloser := io.NopCloser(tar)
	//err = loadDockerTar(tarReadCloser)
	//if err != nil {
	//	return "", nil, err
	//}

	auths, err := getCredentials()
	if err != nil {
		return "", nil, err
	}
	// read auths to a map of authConfigs for the build options to consume
	authConfigs := make(map[string]types.AuthConfig, len(auths))
	for k, auth := range auths {
		authConfigs[k] = types.AuthConfig(auth)
	}

	// make the build options
	//opts := types.ImageBuildOptions{
	//	Dockerfile: img.Build.Dockerfile,
	//	Tags:       []string{img.Name}, //this should build the image locally, sans registry info
	//	Remove:     true,
	//	CacheFrom:  img.Build.CachedImages,
	//	BuildArgs:  build.Args,
	//	Version:    build.BuilderVersion,
	//	Platform:   build.Platform,
	//
	//	AuthConfigs: authConfigs,
	//}

	//// Start a session for BuildKit
	//if build.BuilderVersion == defaultBuilder {
	//	sess, _ := session.NewSession(ctx, "pulumi-docker", "")
	//	dialSession := func(ctx context.Context, proto string, meta map[string][]string) (net.Conn, error) {
	//		return docker.DialHijack(ctx, "/session", proto, meta)
	//	}
	//	go func() {
	//		err := sess.Run(ctx, dialSession)
	//		if err != nil {
	//			return
	//		}
	//	}()
	//	defer sess.Close()
	//	opts.SessionID = sess.ID()
	//}

	q.Q("Starting buildkitclient")

	bkClient, err := buildkitclient.New(ctx, "",
		buildkitclient.WithFailFast(),
		buildkitclient.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return docker.DialHijack(ctx, "/grpc", "h2c", nil)
		}))
	if err != nil {
		return "did not get a new buildkit client", nil, err
	}

	q.Q(bkClient)
	// TODO: figure out what do do with these
	pipeR, pipeW := io.Pipe()

	// TODO looks like we need a "solver" of some sort - what even are words
	q.Q("calling newSolveOpt")
	solveOpt, err := newSolveOpt(build, pipeW)
	if err != nil {
		return "did not get a new solve opts", nil, err
	}
	ch := make(chan *buildkitclient.SolveStatus)
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		var err error
		q.Q("in first eg.Go")
		solvResp, err := bkClient.Build(ctx, *solveOpt, "", dockerfile.Build, ch)
		q.Q(solvResp)
		return fmt.Errorf("hey this thing couldn't be Solved: %q", err)
	})

	//eg.Go(func() error {
	//	var c console.Console
	//	if cn, err := console.ConsoleFromFile(os.Stderr); err == nil {
	//		c = cn
	//	}
	//	// not using shared context to not disrupt display but let is finish reporting errors
	//	_, err = progressui.DisplaySolveStatus(context.TODO(), "", c, os.Stdout, ch)
	//	return err
	//})
	eg.Go(func() error {
		q.Q("about to call loadDockerTar")
		if err := loadDockerTar(pipeR); err != nil { //TODO: maybe use moby/moby to get tar file
			return fmt.Errorf("tar pipe nonsense %q", err)
		}
		q.Q("retruned form loadDorckerTar")
		return pipeR.Close()
	})

	q.Q("this is below the goroutines")
	if err := eg.Wait(); err != nil {
		q.Q("are we hitting the errgroup?")
		return "error group had an error", nil, err
	}
	p.host.Log(ctx, "info", urn, "Loaded the image  to Docker.")

	//bkClient.Build()
	//builder.New()
	//command.Cli()
	//
	//buildxbuild.Build()

	//imgBuildResp, err := docker.ImageBuild(ctx, tar, opts)
	//if err != nil {
	//	return "", nil, err
	//}
	//
	//defer imgBuildResp.Body.Close()
	//
	//// Print build logs to `Info` progress report
	//scanner := bufio.NewScanner(imgBuildResp.Body)
	//for scanner.Scan() {
	//
	//	info, err := processLogLine(scanner.Text())
	//	if err != nil {
	//		return "", nil, err
	//	}
	//	err = p.host.LogStatus(ctx, "info", urn, info)
	//	if err != nil {
	//		return "", nil, err
	//	}
	//}

	// if we are not pushing to the registry, we return after building the local image.
	if img.SkipPush {
		outputs := map[string]interface{}{
			"dockerfile":     img.Build.Dockerfile,
			"context":        img.Build.Context,
			"registryServer": img.Registry.Server,
		}

		pbstruct, err := plugin.MarshalProperties(
			resource.NewPropertyMapFromMap(outputs),
			plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true},
		)
		return img.Name, pbstruct, err
	}

	// TODO: pushing behavior begins here!!!

	err = p.host.LogStatus(ctx, "info", urn, "Pushing Image to the registry")

	if err != nil {
		return "", nil, err
	}

	// authentication for registry push
	// we check if the user set creds in the Pulumi program, and use those preferentially,
	// otherwise we use host machine creds via authConfigs.
	var pushAuthConfig types.AuthConfig

	if img.Registry.Username != "" && img.Registry.Password != "" {
		pushAuthConfig.Username = img.Registry.Username
		pushAuthConfig.Password = img.Registry.Password
		pushAuthConfig.ServerAddress, err = getRegistryAddrForAuth(img.Registry.Server, img.Name)
		if err != nil {
			return "", nil, err
		}

	} else {
		// send warning if user is attempting to use in-program credentials
		if img.Registry.Username == "" && img.Registry.Password != "" {
			msg := "username was not set, although password was; using host credentials file"
			err = p.host.Log(ctx, "warning", urn, msg)
			if err != nil {
				return "", nil, err
			}
		}
		if img.Registry.Password == "" && img.Registry.Username != "" {
			msg := "password was not set, although username was; using host credentials file"
			err = p.host.Log(ctx, "warning", urn, msg)
			if err != nil {
				return "", nil, err
			}
		}

		registryServer, err := getRegistryAddrForAuth(img.Registry.Server, img.Name)
		if err != nil {
			return "", nil, err
		}

		// we use the credentials for the server declared in the program, looking them up from the host authConfigs.
		pushAuthConfig = authConfigs[registryServer]
	}

	authConfigBytes, err := json.Marshal(pushAuthConfig)

	if err != nil {
		return "", nil, fmt.Errorf("error parsing authConfig: %v", err)
	}
	authConfigEncoded := base64.URLEncoding.EncodeToString(authConfigBytes)

	pushOpts := types.ImagePushOptions{RegistryAuth: authConfigEncoded}

	// By default, we push our image with the qualified image name from the input, without extra tagging.
	pushOutput, err := docker.ImagePush(ctx, img.Name, pushOpts)

	if err != nil {
		return "", nil, err
	}

	defer pushOutput.Close()

	// Print push logs to `Info` progress report
	pushScanner := bufio.NewScanner(pushOutput)
	for pushScanner.Scan() {
		info, err := processLogLine(pushScanner.Text())
		if err != nil {
			return "", nil, err
		}
		err = p.host.LogStatus(ctx, "info", urn, info)
		if err != nil {
			return "", nil, err
		}
	}

	outputs := map[string]interface{}{
		"dockerfile":     img.Build.Dockerfile,
		"context":        img.Build.Context,
		"baseImageName":  img.Name,
		"registryServer": img.Registry.Server,
	}
	pbstruct, err := plugin.MarshalProperties(
		resource.NewPropertyMapFromMap(outputs),
		plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true},
	)
	return img.Name, pbstruct, err
}

func marshalBuildAndApplyDefaults(b resource.PropertyValue) (Build, error) {

	// build can be nil, a string or an object; we will also use reasonable defaults here.
	var build Build
	if b.IsNull() {
		// use the default build context
		build.Dockerfile = defaultDockerfile
		build.Context = "."
		return build, nil
	}
	if b.IsString() {
		// use the filepath as context
		build.Context = b.StringValue()
		build.Dockerfile = defaultDockerfile
		return build, nil
	}

	// read in the build type fields
	buildObject := b.ObjectValue()
	// Dockerfile
	if buildObject["dockerfile"].IsNull() {
		// set default
		build.Dockerfile = defaultDockerfile
	} else {
		build.Dockerfile = buildObject["dockerfile"].StringValue()
	}
	// Context
	if buildObject["context"].IsNull() {
		// set default
		build.Context = "."
	} else {
		build.Context = buildObject["context"].StringValue()
	}

	// BuildKit
	version, err := marshalBuilder(buildObject["builderVersion"])

	if err != nil {
		return build, err
	}
	build.BuilderVersion = version

	// Args
	build.Args = marshalArgs(buildObject["args"])

	// Target
	if !buildObject["target"].IsNull() {
		build.Target = buildObject["target"].StringValue()
	}

	// Platform
	if !buildObject["platform"].IsNull() {
		build.Platform = buildObject["platform"].StringValue()
	}
	return build, nil
}

func marshalCachedImages(img Image, b resource.PropertyValue) []string {
	var cacheImages []string
	if b.IsNull() {
		return cacheImages
	}
	c := b.ObjectValue()["cacheFrom"]

	if c.IsNull() {
		return cacheImages
	}

	// if we specify a list of stages, then we only pull those
	cacheFrom := c.ObjectValue()
	stages := cacheFrom["images"].ArrayValue()
	for _, img := range stages {
		stage := img.StringValue()
		cacheImages = append(cacheImages, stage)
	}
	return cacheImages
}

func marshalRegistry(r resource.PropertyValue) Registry {
	var reg Registry
	if !r.IsNull() {
		if !r.ObjectValue()["server"].IsNull() {
			reg.Server = r.ObjectValue()["server"].StringValue()
		}
		if !r.ObjectValue()["username"].IsNull() {
			reg.Username = r.ObjectValue()["username"].StringValue()
		}
		if !r.ObjectValue()["password"].IsNull() {
			reg.Password = r.ObjectValue()["password"].StringValue()
		}
		return reg
	}
	return reg
}

func marshalArgs(a resource.PropertyValue) map[string]*string {
	args := make(map[string]*string)
	if !a.IsNull() {
		for k, v := range a.ObjectValue() {
			key := fmt.Sprintf("%v", k)
			vStr := v.StringValue()
			args[key] = &vStr
		}
	}
	if len(args) == 0 {
		return nil
	}
	return args
}

func marshalBuilder(builder resource.PropertyValue) (types.BuilderVersion, error) {
	var version types.BuilderVersion

	if builder.IsNull() {
		//set default
		return defaultBuilder, nil
	}
	// verify valid input
	switch builder.StringValue() {
	case "BuilderV1":
		return "1", nil
	case "BuilderBuildKit":
		return "2", nil
	default:
		// because the Docker client will default to `BuilderV1`
		// when version isn't set, we return an error
		return version, fmt.Errorf("invalid Docker Builder version")
	}
}

func marshalSkipPush(sp resource.PropertyValue) bool {
	if sp.IsNull() {
		// defaults to false
		return false
	}
	return sp.BoolValue()
}

func getCredentials() (map[string]clitypes.AuthConfig, error) {
	creds, err := config.Load(config.Dir())
	if err != nil {
		return nil, err
	}
	creds.CredentialsStore = credentials.DetectDefaultStore(creds.CredentialsStore)
	auths, err := creds.GetAllCredentials()
	if err != nil {
		return nil, err
	}
	return auths, nil
}

// Because the authConfigs provided by the host include the `https://` prefix in the map keys, `getRegistryAddrForAuth`
// ensures we return a registry address that includes the `https://` scheme.
// While this prefix is not needed for program-provided auth, it is valid regardless, so adding it by default
// keeps the special case handling to a minimum.
func getRegistryAddrForAuth(serverName, imgName string) (string, error) {
	var serverAddr string
	if serverName == "docker.io" {
		// if it's dockerhub, we special case it so host config can find the correct registry
		return "https://index.docker.io/v1/", nil
	}

	if serverName == "" {
		// if there is no servername in the registry input, we attempt to build it from the fully qualified image name.
		addr, err := getRegistryAddrFromImage(imgName)
		if err != nil {
			return "", err
		}

		if addr == "docker.io" {
			return "https://index.docker.io/v1/", nil
		}
		// we need the full server address for the lookup
		serverAddr = "https://" + addr

	} else {
		// check if the provider registry server starts with https://
		if strings.HasPrefix(serverName, "https://") {
			serverAddr = serverName
		} else {
			// courtesy add the prefix so user does not have to explicitly do so
			serverAddr = "https://" + serverName
		}
	}
	return serverAddr, nil
}

func getRegistryAddrFromImage(imgName string) (string, error) {
	named, err := reference.ParseNamed(imgName)
	if err != nil {
		msg := fmt.Errorf("error: %s. This provider requires all image names to be fully qualified.\n"+
			"For example, if you are attempting to push to Dockerhub, prefix your image name with `docker.io`:\n\n"+
			"`docker.io/repository/image:tag`", err)
		return "", msg
	}
	addr := reference.Domain(named)
	return addr, nil

}

func processLogLine(msg string) (string, error) {
	var info string
	var jm jsonmessage.JSONMessage
	err := json.Unmarshal([]byte(msg), &jm)
	if err != nil {
		return info, fmt.Errorf("encountered error unmarshalling: %v", err)
	}
	// process this JSONMessage
	if jm.Error != nil {
		if jm.Error.Code == 401 {
			return info, fmt.Errorf("authentication is required")
		}
		return info, fmt.Errorf(jm.Error.Message)
	}
	if jm.From != "" {
		info += jm.From
	}
	if jm.Progress != nil {
		info += jm.Status + " " + jm.Progress.String()
	} else if jm.Stream != "" {
		info += jm.Stream

	} else {
		info += jm.Status
	}
	if jm.Aux != nil {
		// if we're dealing with buildkit tracer logs, we need to decode
		if jm.ID == "moby.buildkit.trace" {
			// Process the message like the 'tracer.write' method in build_buildkit.go
			// https://github.com/docker/docker-ce/blob/master/components/cli/cli/command/image/build_buildkit.go#L392
			var resp controlapi.StatusResponse
			var infoBytes []byte
			// ignore messages that are not understood
			if err := json.Unmarshal(*jm.Aux, &infoBytes); err != nil {
				info += "failed to parse aux message: " + err.Error()
			}
			if err := (&resp).Unmarshal(infoBytes); err != nil {
				info += "failed to parse aux message: " + err.Error()
			}
			for _, vertex := range resp.Vertexes {
				info += fmt.Sprintf("digest: %+v\n", vertex.Digest)
				info += fmt.Sprintf("%s\n", vertex.Name)
				if vertex.Error != "" {
					info += fmt.Sprintf("error: %s\n", vertex.Error)
				}
			}
			for _, status := range resp.Statuses {
				info += fmt.Sprintf("%s\n", status.GetID())
			}
			for _, log := range resp.Logs {
				info += fmt.Sprintf("%s\n", string(log.Msg))

			}
			for _, warn := range resp.Warnings {
				info += fmt.Sprintf("%s\n", string(warn.Short))
			}

		} else {
			// most other aux messages are secretly a BuildResult
			var result types.BuildResult
			if err := json.Unmarshal(*jm.Aux, &result); err != nil {
				// in the case of non-BuildResult aux messages we print out the whole object.
				infoBytes, err := json.Marshal(jm.Aux)
				if err != nil {
					info += "failed to parse aux message: " + err.Error()
				}
				info += string(infoBytes)
			} else {
				info += result.ID
			}
		}
	}
	return info, nil
}

func newSolveOpt(b Build, w io.WriteCloser) (*buildkitclient.SolveOpt, error) {

	localDirs := map[string]string{
		"context":    b.Context,
		"dockerfile": b.Context,
	}
	q.Q("context in newSolveOpt", b.Context)
	//frontend := "dockerfile.v0" // TODO: use gateway

	file := filepath.Join(b.Context, b.Dockerfile)
	q.Q("filepath joined", file)
	q.Q("filepathbase: ", filepath.Base(file))
	frontendAttrs := map[string]string{
		"filename": filepath.Base(file),
	}
	if b.Target != "" {
		frontendAttrs["target"] = b.Target
	}
	//frontendAttrs["context:"] = "./" // TODO: remove hardcoding
	//if clicontext.Bool("no-cache") {
	//	frontendAttrs["no-cache"] = ""
	//}
	for _, buildArg := range b.Args {
		kv := strings.SplitN(*buildArg, "=", 2)
		if len(kv) != 2 {
			return nil, fmt.Errorf("invalid build-arg value %s", buildArg)
		}
		frontendAttrs["build-arg:"+kv[0]] = kv[1]
	}
	q.Q(frontendAttrs)
	return &buildkitclient.SolveOpt{
		Exports: []buildkitclient.ExportEntry{ //TODO: find out how these behave
			{
				Type: "tar", // TODO: is this the correct type here?
				//Attrs: map[string]string{
				//	"name": "tag", //
				//},
				Output: func(_ map[string]string) (io.WriteCloser, error) {
					return w, nil //TODO hoping this is the build output writer
				},
			},
		},
		LocalDirs:     localDirs,
		Frontend:      "",
		FrontendAttrs: frontendAttrs,
	}, nil
}

func loadDockerTar(r io.Reader) error {
	// TODO: possibly use moby/moby/client here
	q.Q("in loadDockerTar")
	cmd := exec.Command("docker", "load")
	cmd.Stdin = r
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
