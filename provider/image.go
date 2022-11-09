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
	structpb "github.com/golang/protobuf/ptypes/struct"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
)

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
	Args         map[string]string
	ExtraOptions []string
	Target       string
}

func (p *dockerNativeProvider) dockerBuild(ctx context.Context,
	urn resource.URN,
	props *structpb.Struct) (*structpb.Struct, error) {

	fmt.Println("getting inputs...")

	inputs, err := plugin.UnmarshalProperties(props, plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true})
	if err != nil {
		return nil, err
	}

	// read in values to Image
	img := Image{
		Name:     inputs["imageName"].StringValue(),
		SkipPush: inputs["skipPush"].BoolValue(),
		Registry: Registry{
			Server:   inputs["registry"].ObjectValue()["server"].StringValue(),
			Username: inputs["registry"].ObjectValue()["username"].StringValue(),
			Password: inputs["registry"].ObjectValue()["password"].StringValue(),
		},
	}

	build := marshalBuild(inputs["build"])
	cache := getCachedImages(img, inputs["build"].ObjectValue()["cacheFrom"])

	build.CachedImages = cache
	img.Build = build

	fmt.Println("USING THE DOCKER CLIENT NOW")

	docker, err := client.NewClientWithOpts(client.FromEnv)

	if err != nil {
		panic(err)
	}

	fmt.Println("now let's BUILD an image")

	// make the build context
	tar, err := archive.TarWithOptions(img.Build.Context, &archive.TarOptions{})
	if err != nil {
		panic(err)
	}

	// make the build options TODO: this is where we will add the buildkit flags etc
	fmt.Println(img.Build.CachedImages)
	opts := types.ImageBuildOptions{
		Dockerfile: img.Build.Dockerfile,
		Tags:       []string{img.Name}, //this should build the image locally, sans registry info
		Remove:     true,
		CacheFrom:  img.Build.CachedImages,
	}

	imgBuildResp, err := docker.ImageBuild(context.Background(), tar, opts)
	if err != nil {
		panic(err)
	}

	defer imgBuildResp.Body.Close()
	// Print build logs to terminal
	scanner := bufio.NewScanner(imgBuildResp.Body)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
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

	fmt.Println("see if we can push to the registry")

	// Quick and dirty auth; we can also preconfigure the client itself I believe

	var authConfig = types.AuthConfig{
		Username:      img.Registry.Username,
		Password:      img.Registry.Password,
		ServerAddress: img.Registry.Server,
	}

	authConfigBytes, _ := json.Marshal(authConfig)
	authConfigEncoded := base64.URLEncoding.EncodeToString(authConfigBytes)

	pushOpts := types.ImagePushOptions{RegistryAuth: authConfigEncoded}

	// This is a shift away from imageName to have to be of the "registry/image" format.
	registryImageTag := img.Registry.Username + "/" + img.Name
	// tag our image with the qualified image name that includes the registry
	err = docker.ImageTag(context.Background(), img.Name, registryImageTag)
	if err != nil {
		panic(err)
	}
	pushOutput, err := docker.ImagePush(context.Background(), registryImageTag, pushOpts)

	if err != nil {
		panic(err)
	}

	defer pushOutput.Close()

	// Print push logs to terminal
	pushScanner := bufio.NewScanner(pushOutput)
	for pushScanner.Scan() {
		fmt.Println(pushScanner.Text())
	}

	outputs := map[string]interface{}{
		"dockerfile":        img.Build.Dockerfile,
		"context":           img.Build.Context,
		"registryImageName": registryImageTag,
		"registryServer":    img.Registry.Server,
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
		build.Dockerfile = "Dockerfile"
		build.Context = "."
		return build
	}
	if b.IsString() {
		// use the filepath as context
		build.Context = b.StringValue()
		build.Dockerfile = "Dockerfile" // default to Dockerfile
		return build
	}

	// read in the build type fields
	buildObject := b.ObjectValue()
	// Dockerfile
	if buildObject["dockerfile"].IsNull() {
		// set default
		build.Dockerfile = "Dockerfile"
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
	args := make(map[string]string)
	if !buildObject["args"].IsNull() {
		for k, v := range buildObject["args"].ObjectValue() {
			key := fmt.Sprintf("%v", k)
			args[key] = v.StringValue()
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

func getCachedImages(img Image, c resource.PropertyValue) []string {

	var cacheImages []string

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
		// TODO: Verify correct behavior for this feature and how it works with the Docker client
		stage := img.StringValue()
		cacheImages = append(cacheImages, stage)
	}
	return cacheImages
}
