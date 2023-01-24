package provider

import (
	"bufio"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/moby/buildkit/session"
	"net"
	"strings"

	"github.com/docker/cli/cli/config"
	"github.com/docker/cli/cli/config/credentials"
	clitypes "github.com/docker/cli/cli/config/types"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
	"github.com/docker/docker/pkg/jsonmessage"
	structpb "github.com/golang/protobuf/ptypes/struct"
	controlapi "github.com/moby/buildkit/api/services/control"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
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
	Env            map[string]string
	Args           map[string]*string
	ExtraOptions   []string
	Target         string
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
	tar, err := archive.TarWithOptions(img.Build.Context, &archive.TarOptions{})
	if err != nil {
		return "", nil, err
	}

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
	opts := types.ImageBuildOptions{
		Dockerfile: img.Build.Dockerfile,
		Tags:       []string{img.Name}, //this should build the image locally, sans registry info
		Remove:     true,
		//CacheFrom:  img.Build.CachedImages, // TODO: this needs a login, so needs to be handled differently.
		BuildArgs: build.Args,
		Version:   build.BuilderVersion,

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
		return "", nil, errors.Wrap(err, "Error parsing authConfig")
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

	// Envs
	build.Env = marshalEnvs(buildObject["env"])

	// Args
	build.Args = marshalArgs(buildObject["args"])

	// ExtraOptions
	if !buildObject["extraOptions"].IsNull() {
		opts := buildObject["extraOptions"].ArrayValue()
		for _, v := range opts {
			build.ExtraOptions = append(build.ExtraOptions, v.StringValue())
		}
	}

	// Target
	if !buildObject["target"].IsNull() {
		build.Target = buildObject["target"].StringValue()
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
	latest := img.Registry.Username + "/" + img.Name
	// if we specify cacheFrom as True, we pull the latest build+push of our image implicitly, i.e. registry/image
	if c.IsBool() {
		useCache := c.BoolValue()
		if useCache {
			cacheImages = append(cacheImages, latest)
		}
		return cacheImages
	}
	// if we specify a list of stages, then we only pull those
	cacheFrom := c.ObjectValue()
	stages := cacheFrom["stages"].ArrayValue()
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

func marshalEnvs(e resource.PropertyValue) map[string]string {
	envs := make(map[string]string)
	if !e.IsNull() {
		for k, v := range e.ObjectValue() {
			key := fmt.Sprintf("%v", k)
			envs[key] = v.StringValue()
		}
	}
	if len(envs) == 0 {
		return nil
	}
	return envs
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
		return version, errors.Errorf("Invalid Docker Builder version")
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
		addr, _, found := strings.Cut(imgName, "/")
		if !found {
			return "", errors.Errorf("image name must be fully qualified: %s", imgName)
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
		// TODO: this is where we would default to Docker if we wanted to do so.
		// A valid registry address includes a `.` so we could add additional helper checks here.
	}
	return serverAddr, nil
}

func processLogLine(msg string) (string, error) {
	var info string
	var jm jsonmessage.JSONMessage
	err := json.Unmarshal([]byte(msg), &jm)
	if err != nil {
		return info, errors.Wrapf(err, "encountered error unmarshalling:")
	}
	// process this JSONMessage
	if jm.Error != nil {
		if jm.Error.Code == 401 {
			return info, fmt.Errorf("authentication is required")
		}
		return info, errors.Errorf(jm.Error.Message)
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
				info += fmt.Sprintf("layer: %+v\n", vertex.Digest)
			}
			for _, status := range resp.Statuses {
				info += fmt.Sprintf("status: %s\n", status.GetID())
			}
			for _, log := range resp.Logs {
				info += fmt.Sprintf("log: %+v\n", log.GetMsg())
			}
			for _, warn := range resp.Warnings {
				info += fmt.Sprintf("warning: %+v\n", warn.GetShort())
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
