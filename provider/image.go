package provider

import (
	"bufio"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
	"path"
	"path/filepath"
	"strings"

	buildCmd "github.com/docker/cli/cli/command/image/build"
	clibuild "github.com/docker/cli/cli/command/image/build"
	"github.com/docker/cli/cli/config"
	"github.com/docker/cli/cli/config/configfile"
	"github.com/docker/cli/cli/config/credentials"
	clitypes "github.com/docker/cli/cli/config/types"
	"github.com/docker/distribution/reference"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
	"github.com/docker/docker/pkg/idtools"
	"github.com/docker/docker/pkg/jsonmessage"
	structpb "github.com/golang/protobuf/ptypes/struct"
	controlapi "github.com/moby/buildkit/api/services/control"
	"github.com/moby/buildkit/session"
	"github.com/moby/buildkit/session/auth/authprovider"
	"github.com/moby/moby/registry"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
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
	q.Q(inputs)
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
	cache, err := marshalCachedImages(inputs["build"])
	if err != nil {
		return "", nil, err
	}

	build.CachedImages = cache
	img.Build = build

	ca := "../aws-k8s-dockerd-ts/certs/ca.pem"
	cert := "../aws-k8s-dockerd-ts/certs/cert.pem"
	key := "../aws-k8s-dockerd-ts/certs/key.pem"

	docker, err := client.NewClientWithOpts(
		client.FromEnv,
		client.WithAPIVersionNegotiation(),
		client.WithTLSClientConfig(ca, cert, key),
	)

	if err != nil {
		return "", nil, err
	}

	err = p.host.LogStatus(ctx, "info", urn, "Building the image")

	if err != nil {
		return "", nil, err
	}

	// make the build context and ensure to exclude dockerignore file patterns
	dockerIgnorePath := filepath.Join(build.Context, ".dockerignore")
	initialIgnorePatterns, err := getIgnore(dockerIgnorePath)
	if err != nil {
		return "", nil, fmt.Errorf("error reading ignore file: %w", err)
	}

	absDockerfile, err := filepath.Abs(build.Dockerfile)
	if err != nil {
		return "", nil, fmt.Errorf("absDockerfile error: %s", err)
	}
	absBuildpath, err := filepath.Abs(build.Context)
	if err != nil {
		return "", nil, fmt.Errorf("absBuildPath error: %s", err)
	}
	relDockerfile, err := filepath.Rel(absBuildpath, absDockerfile)
	if err != nil {
		return "", nil, fmt.Errorf("relDockerfile error: %s", err)
	}

	// if the dockerfile is in the context it will be something like "./Dockerfile" or ".\sub\dir\Dockerfile"
	// if the dockerfile is out of the context it will begin with "../"
	dockerfileInContext := true
	if strings.HasPrefix(relDockerfile, ".."+string(filepath.Separator)) {
		dockerfileInContext = false
	}

	contextDir, err := clibuild.ResolveAndValidateContextPath(build.Context)
	if err != nil {
		return "", nil, fmt.Errorf("error resolving context: %w", err)
	}

	if err := clibuild.ValidateContextDirectory(contextDir, initialIgnorePatterns); err != nil {
		return "", nil, fmt.Errorf("error validating context: %w", err)
	}

	// un-ignore build files so the docker daemon can use them
	ignorePatterns := buildCmd.TrimBuildFilesFromExcludes(
		initialIgnorePatterns,
		relDockerfile,
		false,
	)

	// warn user about accidentally copying build files
	if build.BuilderVersion == defaultBuilder && len(initialIgnorePatterns) != len(ignorePatterns) {
		msg := "It looks like you are trying to dockerignore a build file such as `Dockerfile` or `.dockerignore`. " +
			"To avoid accidentally copying these files to your image, please ensure any copied file systems do not " +
			"include `Dockerfile` or `.dockerignore`."
		err = p.host.Log(ctx, "warning", urn, msg)
		if err != nil {
			return "", nil, err
		}
	}

	userIdentity := idtools.CurrentIdentity()
	q.Q(userIdentity)

	tar, err := archive.TarWithOptions(contextDir, &archive.TarOptions{
		ExcludePatterns: ignorePatterns,
		ChownOpts:       &idtools.Identity{UID: 0, GID: 0}, // TODO: this does not seem to affect anything in TarWithOptions.
	})
	//TODO: the error "Can't add ... to tar file" may still be a permissions error - just not one that is controlled by ChownOpts.
	if err != nil {
		return "", nil, err
	}

	// add dockerfile to tarball if it's not in the build context
	replaceDockerfile := relDockerfile
	if !dockerfileInContext {
		// Handle Dockerfile from outside of build context folder
		var dockerfileCtx io.ReadCloser
		dockerfileCtx, err = os.Open(build.Dockerfile)
		if err != nil {
			return "", nil, err
		}
		tar, replaceDockerfile, err = clibuild.AddDockerfileToBuildContext(dockerfileCtx, tar)
		if err != nil {
			return "", nil, err
		}
	}

	q.Q("do we still need even more configs? did we make it here?")

	cfg, err := getDefaultDockerConfig()
	if err != nil {
		return "", nil, err
	}

	authConfigs := make(map[string]types.AuthConfig)
	var regAuth types.AuthConfig

	auths, err := cfg.GetAllCredentials()
	if err != nil {
		return "", nil, err
	}
	for k, auth := range auths {
		authConfigs[k] = types.AuthConfig(auth)
	}

	// sign into registry if we're pushing or setting CacheFrom
	// TODO: add functionality for additional registry caches not associated with the stack image
	// See: https://github.com/pulumi/pulumi-docker/issues/497
	if len(img.Build.CachedImages) > 0 || !img.SkipPush {
		auth, msg, err := getRegistryAuth(img, cfg)
		if err != nil {
			return "", nil, err
		}
		if msg != "" {
			err = p.host.Log(ctx, "warning", urn, msg)
			if err != nil {
				return "", nil, err
			}
		}
		authConfigs[auth.ServerAddress] = auth                          // for image cache
		cfg.AuthConfigs[auth.ServerAddress] = clitypes.AuthConfig(auth) // for buildkit cache using session auth
		regAuth = auth                                                  // for image push
	}
	// make the build options
	opts := types.ImageBuildOptions{
		Dockerfile: replaceDockerfile,
		Tags:       []string{img.Name}, //this should build the image locally, sans registry info
		CacheFrom:  img.Build.CachedImages,
		BuildArgs:  build.Args,
		Version:    build.BuilderVersion,
		Platform:   build.Platform,
		Target:     build.Target,

		AuthConfigs: authConfigs,
	}

	// Start a session for BuildKit
	if build.BuilderVersion == defaultBuilder {
		sess, _ := session.NewSession(ctx, "pulumi-docker", "")

		dockerAuthProvider := authprovider.NewDockerAuthProvider(cfg)
		sess.Allow(dockerAuthProvider)

		dialSession := func(ctx context.Context, proto string, meta map[string][]string) (net.Conn, error) {
			return docker.DialHijack(ctx, "/session", proto, meta)
		}
		go func() {
			err := sess.Run(ctx, dialSession)
			if err != nil {
				return
			}
		}()
		defer sess.Close()
		opts.SessionID = sess.ID()
	}

	imgBuildResp, err := docker.ImageBuild(ctx, tar, opts)
	if err != nil {
		q.Q("we hit this error after Build")
		return "", nil, err
	} else {
		q.Q("we did NOT hit an error after build")
	}

	defer imgBuildResp.Body.Close()

	// Print build logs to `Info` progress report
	scanner := bufio.NewScanner(imgBuildResp.Body)
	for scanner.Scan() {

		info, err := processLogLine(scanner.Text())
		if err != nil {
			return "", nil, err
		}
		err = p.host.LogStatus(ctx, "info", urn, info)
		if err != nil {
			return "", nil, err
		}
	}

	outputs := map[string]interface{}{
		"dockerfile":     relDockerfile,
		"context":        img.Build.Context,
		"baseImageName":  img.Name,
		"registryServer": img.Registry.Server,
		"imageName":      img.Name,
	}

	// if we are not pushing to the registry, we return after building the local image.
	if img.SkipPush {
		pbstruct, err := plugin.MarshalProperties(
			resource.NewPropertyMapFromMap(outputs),
			plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true},
		)
		return img.Name, pbstruct, err
	}

	err = p.host.LogStatus(ctx, "info", urn, "Pushing Image to the registry")

	if err != nil {
		return "", nil, err
	}

	authConfigBytes, err := json.Marshal(regAuth)

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

	dist, _, err := docker.ImageInspectWithRaw(ctx, img.Name)
	if err != nil {
		return "", nil, err
	}

	// The repoDigest should be populated after a push. Clients may choose to throw an error or coerce
	// this to a non-optional value.
	if len(dist.RepoDigests) > 0 {
		outputs["repoDigest"] = dist.RepoDigests[0]
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
		build.Context = "."
		build.Dockerfile = defaultDockerfile
		return build, nil
	}
	// read in the build type fields
	buildObject := b.ObjectValue()

	// Context
	if buildObject["context"].IsNull() {
		// set default
		build.Context = "."
	} else {
		build.Context = buildObject["context"].StringValue()
	}

	// Dockerfile
	if buildObject["dockerfile"].IsNull() {
		// set default
		build.Dockerfile = path.Join(build.Context, defaultDockerfile)
	} else {
		build.Dockerfile = buildObject["dockerfile"].StringValue()
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

func marshalCachedImages(b resource.PropertyValue) ([]string, error) {
	var cacheImages []string
	if b.IsNull() {
		return cacheImages, nil
	}
	c := b.ObjectValue()["cacheFrom"]

	if c.IsNull() {
		return cacheImages, nil
	}

	// if we specify a list of stages, then we only pull those
	cacheFrom := c.ObjectValue()
	if cacheFrom["images"].IsNull() {
		return cacheImages, fmt.Errorf("if you want to use cacheFrom, you must have `images` set")
	}
	stages := cacheFrom["images"].ArrayValue()
	for _, img := range stages {
		stage := img.StringValue()
		cacheImages = append(cacheImages, stage)
	}
	return cacheImages, nil
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

func getDefaultDockerConfig() (*configfile.ConfigFile, error) {
	q.Q("in docker config")
	cfg, err := config.Load(config.Dir())
	q.Q(cfg)
	if err != nil {
		return nil, err
	}
	q.Q(cfg.CredentialsStore)
	cfg.CredentialsStore = credentials.DetectDefaultStore(cfg.CredentialsStore)
	return cfg, nil
}

func getRegistryAuth(img Image, cfg *configfile.ConfigFile) (types.AuthConfig, string, error) {
	// authentication for registry push or cache pull
	// we check if the user set creds in the Pulumi program, and use those preferentially,
	// otherwise we use host machine creds via authConfigs.
	var regAuthConfig types.AuthConfig
	var msg string

	if img.Registry.Username != "" && img.Registry.Password != "" {
		regAuthConfig.Username = img.Registry.Username
		regAuthConfig.Password = img.Registry.Password
		serverAddr, err := getRegistryAddrForAuth(img.Registry.Server, img.Name)
		if err != nil {
			return regAuthConfig, msg, err
		}
		regAuthConfig.ServerAddress = serverAddr

	} else {
		// send warning if user is attempting to use in-program credentials
		if img.Registry.Username == "" && img.Registry.Password != "" {
			msg = "username was not set, although password was; using host credentials file"
		}
		if img.Registry.Password == "" && img.Registry.Username != "" {
			msg = "password was not set, although username was; using host credentials file"
		}

		registryServer, err := getRegistryAddrForAuth(img.Registry.Server, img.Name)
		if err != nil {
			return regAuthConfig, msg, err
		}

		cliPushAuthConfig, err := cfg.GetAuthConfig(registryServer)
		if err != nil {
			return regAuthConfig, msg, err
		}

		regAuthConfig = types.AuthConfig(cliPushAuthConfig)
	}
	return regAuthConfig, msg, nil
}

// Because the authConfigs provided by the host may return URIs with the `https://` scheme in the
// map keys, `getRegistryAddrForAuth` ensures we return either the legacy Docker IndexServer's URI,
// which is special cased, or a registry hostname.
func getRegistryAddrForAuth(serverName, imgName string) (string, error) {
	var hostname string

	if serverName == "" {
		// if there is no servername in the registry input, we attempt to build it from the fully qualified image name.
		var err error
		hostname, err = getRegistryAddrFromImage(imgName)
		if err != nil {
			return "", err
		}
	} else {
		hostname = registry.ConvertToHostname(serverName)
	}

	switch hostname {
	// handle historically permitted names, mapping them to the v1 registry hostname
	case registry.IndexHostname, registry.IndexName, registry.DefaultV2Registry.Host:
		return registry.IndexServer, nil
	}
	return hostname, nil
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
		if jm.Error.Message == "EOF" {
			return info, fmt.Errorf("%s\n: This error is most likely due to incorrect or mismatched registry "+
				"credentials. Please double check you are using the correct credentials and registry name.",
				jm.Error.Message)
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
