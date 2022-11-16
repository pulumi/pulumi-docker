package provider

import (
	"bufio"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
	"github.com/docker/docker/pkg/jsonmessage"
	structpb "github.com/golang/protobuf/ptypes/struct"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
)

const defaultDockerfile = "Dockerfile"

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
	Context      string
	Dockerfile   string
	CachedImages []string
	Env          map[string]string
	Args         map[string]*string
	ExtraOptions []string
	Target       string
}

func (p *dockerNativeProvider) dockerBuild(ctx context.Context,
	urn resource.URN,
	props *structpb.Struct) (*structpb.Struct, error) {

	inputs, err := plugin.UnmarshalProperties(props, plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true})
	if err != nil {
		return nil, err
	}

	//set Registry
	var reg Registry
	if !inputs["registry"].IsNull() {
		reg = Registry{
			Server:   inputs["registry"].ObjectValue()["server"].StringValue(),
			Username: inputs["registry"].ObjectValue()["username"].StringValue(),
			Password: inputs["registry"].ObjectValue()["password"].StringValue(),
		}
	}

	// read in values to Image
	img := Image{
		Name:     inputs["imageName"].StringValue(),
		SkipPush: inputs["skipPush"].BoolValue(),
		Registry: reg,
	}

	build := marshalBuild(inputs["build"])
	cache := getCachedImages(img, inputs["build"])

	build.CachedImages = cache
	img.Build = build

	docker, err := client.NewClientWithOpts(client.FromEnv)

	if err != nil {
		panic(err)
	}

	err = p.host.Log(ctx, "info", urn, "Building the image")

	if err != nil {
		return nil, err
	}

	// make the build context
	tar, err := archive.TarWithOptions(img.Build.Context, &archive.TarOptions{})
	if err != nil {
		return nil, err
	}

	// make the build options

	opts := types.ImageBuildOptions{
		Dockerfile: img.Build.Dockerfile,
		Tags:       []string{img.Name}, //this should build the image locally, sans registry info
		Remove:     true,
		//CacheFrom:  img.Build.CachedImages, // TODO: this needs a login, so needs to be handled differently.
		BuildArgs: build.Args,
		Version:   types.BuilderBuildKit,
	}

	imgBuildResp, err := docker.ImageBuild(context.Background(), tar, opts)
	if err != nil {
		return nil, err
	}

	defer imgBuildResp.Body.Close()
	// Print build logs to terminal
	scanner := bufio.NewScanner(imgBuildResp.Body)
	for scanner.Scan() {
		err := p.host.Log(ctx, "info", urn, scanner.Text())
		if err != nil {
			return nil, err
		}
	}

	// if we are not pushing to the registry, we return after building the local image.
	if img.SkipPush {
		outputs := map[string]interface{}{
			"dockerfile":     img.Build.Dockerfile,
			"context":        img.Build.Context,
			"registryServer": img.Registry.Server,
		}
		return plugin.MarshalProperties(
			resource.NewPropertyMapFromMap(outputs),
			plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true},
		)
	}

	err = imgBuildResp.Body.Close()

	if err != nil {
		return nil, err
	}

	err = p.host.Log(ctx, "info", urn, "Pushing Image to the registry")

	if err != nil {
		return nil, err
	}
	// Quick and dirty auth; we can also preconfigure the client itself I believe

	var authConfig = types.AuthConfig{
		Username:      img.Registry.Username,
		Password:      img.Registry.Password,
		ServerAddress: img.Registry.Server,
	}

	authConfigBytes, err := json.Marshal(authConfig)

	if err != nil {
		return nil, errors.Wrap(err, "Error parsing authConfig")
	}
	authConfigEncoded := base64.URLEncoding.EncodeToString(authConfigBytes)

	pushOpts := types.ImagePushOptions{RegistryAuth: authConfigEncoded}

	// By default, we push our image with the qualified image name from the input, without extra tagging.
	pushOutput, err := docker.ImagePush(context.Background(), img.Name, pushOpts)

	if err != nil {
		return nil, err
	}

	defer pushOutput.Close()

	// Print push logs to terminal
	pushScanner := bufio.NewScanner(pushOutput)
	for pushScanner.Scan() {
		msg := pushScanner.Text()
		var jsmsg jsonmessage.JSONMessage
		err := json.Unmarshal([]byte(msg), &jsmsg)
		if err != nil {
			return nil, errors.Wrapf(err, "encountered error unmarshalling:")
		}
		if jsmsg.Status != "" {
			if jsmsg.Status != "Pushing" {
				var info string
				if jsmsg.ID != "" {
					info = fmt.Sprintf("%s: %s", jsmsg.ID, jsmsg.Status)
				} else {
					info = jsmsg.Status

				}
				err := p.host.Log(ctx, "info", urn, info)
				if err != nil {
					return nil, err
				}
			}
		}

		if jsmsg.Error != nil {
			return nil, errors.Errorf(jsmsg.Error.Message)
		}

	}

	outputs := map[string]interface{}{
		"dockerfile":     img.Build.Dockerfile,
		"context":        img.Build.Context,
		"baseImageName":  img.Name,
		"registryServer": img.Registry.Server,
	}
	return plugin.MarshalProperties(
		resource.NewPropertyMapFromMap(outputs),
		plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true},
	)
}

func marshalBuild(b resource.PropertyValue) Build {

	// build can be nil, a string or an object; we will also use reasonable defaults here.
	var build Build

	if b.IsNull() {
		// use the default build context
		build.Dockerfile = defaultDockerfile
		build.Context = "."
		return build
	}
	if b.IsString() {
		// use the filepath as context
		build.Context = b.StringValue()
		build.Dockerfile = defaultDockerfile
		return build
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
	// Envs
	envs := make(map[string]string)
	if !buildObject["env"].IsNull() {
		for k, v := range buildObject["env"].ObjectValue() {
			key := fmt.Sprintf("%v", k)
			envs[key] = v.StringValue()
		}
	}
	build.Env = envs
	// Args
	args := make(map[string]*string)
	if !buildObject["args"].IsNull() {
		for k, v := range buildObject["args"].ObjectValue() {
			key := fmt.Sprintf("%v", k)
			vStr := v.StringValue()
			args[key] = &vStr
		}
	}
	build.Args = args

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
	return build
}

func getCachedImages(img Image, b resource.PropertyValue) []string {

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
