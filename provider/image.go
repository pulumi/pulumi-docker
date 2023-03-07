package provider

import (
	"bufio"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/docker/distribution/reference"
	"github.com/moby/buildkit/session"
	"github.com/moby/moby/registry"
	"net"
	"path/filepath"

	"github.com/docker/cli/cli/config"
	"github.com/docker/cli/cli/config/configfile"
	"github.com/docker/cli/cli/config/credentials"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
	"github.com/docker/docker/pkg/jsonmessage"
	structpb "github.com/golang/protobuf/ptypes/struct"
	controlapi "github.com/moby/buildkit/api/services/control"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"

	buildCmd "github.com/docker/cli/cli/command/image/build"
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

	// make the build context and ensure to exclude dockerignore file patterns
	dockerIgnorePath := filepath.Join(build.Context, ".dockerignore")
	initialIgnorePatterns, err := getIgnore(dockerIgnorePath)
	// un-ignore build files so the docker daemon can use them
	ignorePatterns := buildCmd.TrimBuildFilesFromExcludes(
		initialIgnorePatterns,
		img.Build.Dockerfile,
		false,
	)

	// warn user about accidentally copying build files
	if build.BuilderVersion == defaultBuilder && len(initialIgnorePatterns) != len(ignorePatterns) {
		msg := "It looks like you are trying to dockerignore a build file such as Dockerfile or .dockerignore. " +
			"Due to limitations when running this provider in Buildkit mode, your build files may get copied " +
			"into your image. Please ensure any copied file systems do not include build files."
		err = p.host.Log(ctx, "warning", urn, msg)
		if err != nil {
			return "", nil, err
		}
	}

	if err != nil {
		return "", nil, err
	}

	tar, err := archive.TarWithOptions(img.Build.Context, &archive.TarOptions{
		ExcludePatterns: ignorePatterns,
	})
	if err != nil {
		return "", nil, err
	}

	cfg, err := getDefaultDockerConfig()
	if err != nil {
		return "", nil, err
	}

	authConfigs := make(map[string]types.AuthConfig)
	var regAuth types.AuthConfig

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
		authConfigs[auth.ServerAddress] = auth // for image cache
		regAuth = auth                         // for image push
	}
	// make the build options
	opts := types.ImageBuildOptions{
		Dockerfile: img.Build.Dockerfile,
		Tags:       []string{img.Name}, //this should build the image locally, sans registry info
		Remove:     true,
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
		return "", nil, err
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

func getDefaultDockerConfig() (*configfile.ConfigFile, error) {
	cfg, err := config.Load(config.Dir())
	if err != nil {
		return nil, err
	}
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
