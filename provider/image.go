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
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/mapper"
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
	Context      string             `pulumi:"context,optional"`
	Dockerfile   string             `pulumi:"dockerfile,optional"`
	CachedImages []string           `pulumi:"cachedImages,optional"`
	Env          map[string]string  `pulumi:"env,optional"`
	Args         map[string]*string `pulumi:"args,optional"`
	ExtraOptions []string           `pulumi:"extraOptions,optional"`
	Target       string             `pulumi:"target,optional"`
}

func (p *dockerNativeProvider) dockerBuild(ctx context.Context,
	urn resource.URN,
	props *structpb.Struct) (string, *structpb.Struct, error) {

	inputs, err := plugin.UnmarshalProperties(props, plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true})
	if err != nil {
		return "", nil, err
	}

	reg := setRegistry(inputs["registry"])
	// read in values to Image
	img := Image{
		Name:     inputs["imageName"].StringValue(),
		SkipPush: inputs["skipPush"].BoolValue(),
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

	err = p.host.Log(ctx, "info", urn, "Building the image")

	if err != nil {
		return "", nil, err
	}

	// make the build context
	tar, err := archive.TarWithOptions(img.Build.Context, &archive.TarOptions{})
	if err != nil {
		return "", nil, err
	}

	// make the build options

	opts := types.ImageBuildOptions{
		Dockerfile: img.Build.Dockerfile,
		Tags:       []string{img.Name}, //this should build the image locally, sans registry info
		Remove:     true,
		//CacheFrom:  img.Build.CachedImages, // TODO: this needs a login, so needs to be handled differently.
		BuildArgs: build.Args,
		//Version:   types.BuilderBuildKit, // TODO: parse this setting from the `env` input
	}

	imgBuildResp, err := docker.ImageBuild(ctx, tar, opts)
	if err != nil {
		return "", nil, err
	}

	defer imgBuildResp.Body.Close()
	// Print build logs to terminal
	scanner := bufio.NewScanner(imgBuildResp.Body)
	for scanner.Scan() {
		err := p.host.Log(ctx, "info", urn, scanner.Text())
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

	err = p.host.Log(ctx, "info", urn, "Pushing Image to the registry")

	if err != nil {
		return "", nil, err
	}
	// Quick and dirty auth; we can also preconfigure the client itself I believe

	var authConfig = types.AuthConfig{
		Username:      img.Registry.Username,
		Password:      img.Registry.Password,
		ServerAddress: img.Registry.Server,
	}

	authConfigBytes, err := json.Marshal(authConfig)

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

	// Print push logs to terminal
	pushScanner := bufio.NewScanner(pushOutput)
	for pushScanner.Scan() {
		msg := pushScanner.Text()
		var jsmsg jsonmessage.JSONMessage
		err := json.Unmarshal([]byte(msg), &jsmsg)
		if err != nil {
			return "", nil, errors.Wrapf(err, "encountered error unmarshalling:")
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
					return "", nil, err
				}
			}
		}

		if jsmsg.Error != nil {
			return "", nil, errors.Errorf(jsmsg.Error.Message)
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

func buildToPropertyValue(build Build) (resource.PropertyValue, error) {
	pm, err := buildToPropertyMap(build)
	if err != nil {
		return resource.PropertyValue{}, err
	}
	return resource.NewObjectProperty(pm), nil
}

func buildToPropertyMap(build Build) (resource.PropertyMap, error) {
	mapper := mapper.New(nil)
	m, err := mapper.Encode(build)
	if err != nil {
		return nil, err
	}
	return resource.NewPropertyMapFromMap(m), nil
}

func buildFromPropertyMap(pm resource.PropertyMap) (Build, error) {
	var build Build
	err := mapper.Map(pm.Mappable(), &build)
	if err != nil {
		return Build{}, err
	}
	return build, nil
}

func buildFromString(context string) (Build, error) {
	return Build{Context: context}, nil
}

func buildFromPropertyValue(pv resource.PropertyValue) (Build, error) {
	if pv.IsNull() {
		return Build{}, nil
	}
	if pv.IsString() {
		return buildFromString(pv.StringValue())
	}
	if pv.IsObject() {
		return buildFromPropertyMap(pv.ObjectValue())
	}
	return Build{}, fmt.Errorf("Cannot recognize Build from: %v", pv)
}

func applyBuildDefaults(build Build) Build {
	if build.Context == "" {
		build.Context = "."
	}
	if build.Dockerfile == "" {
		build.Dockerfile = defaultDockerfile
	}
	return build
}

func marshalBuildAndApplyDefaults(pv resource.PropertyValue) (Build, error) {
	build, err := buildFromPropertyValue(pv)
	if err != nil {
		return build, err
	}
	return applyBuildDefaults(build), nil
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

func setRegistry(r resource.PropertyValue) Registry {
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
