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
	ImageName string
	rawBuild  interface{} // string or Build
	Registry  Registry
	SkipPush  bool
	// Tag     string  // Tag appears to be unused?
}

func (img *Image) Build() (build *Build) {
	switch b := img.rawBuild.(type) {
	case string:
		build = &Build{Context: b}
	case Build:
		build = &b
	}
	if build.Context == "" {
		build.Context = "."
	}
	if build.Dockerfile == "" {
		build.Dockerfile = defaultDockerfile
	}
	return
}

func (img *Image) CacheImages() (res []string) {
	switch cacheFrom := img.Build().CacheFrom.(type) {
	case bool:
		// if we specify cacheFrom as True, we pull the latest build+push of our image implicitly, i.e.
		// registry/image
		if cacheFrom {
			latest := img.Registry.Username + "/" + img.ImageName
			res = []string{latest}
		}
	case CacheFrom:
		res = cacheFrom.Stages
	}
	return
}

type Registry struct {
	Server   string `pulumi:"server,optional"`
	Username string `pulumi:"username,optional"`
	Password string `pulumi:"password,optional"`
}

type Build struct {
	Context      string
	Dockerfile   string
	CacheFrom    interface{} // bool or CacheFrom
	Env          map[string]string
	Args         map[string]*string
	ExtraOptions []string
	Target       string
}

type CacheFrom struct {
	Stages []string `pulumi:"stages,optional"`
}

func (p *dockerNativeProvider) dockerBuild(ctx context.Context,
	urn resource.URN,
	props *structpb.Struct) (string, *structpb.Struct, error) {

	inputs, err := plugin.UnmarshalProperties(props, plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true})
	if err != nil {
		return "", nil, err
	}

	img := new(Image)
	if err := img.unmarshalPropertyMap(inputs); err != nil {
		return "", nil, err
	}

	docker, err := client.NewClientWithOpts(client.FromEnv)

	if err != nil {
		return "", nil, err
	}

	err = p.host.Log(ctx, "info", urn, "Building the image")

	if err != nil {
		return "", nil, err
	}

	// make the build context
	tar, err := archive.TarWithOptions(img.Build().Context, &archive.TarOptions{})
	if err != nil {
		return "", nil, err
	}

	// make the build options

	opts := types.ImageBuildOptions{
		Dockerfile: img.Build().Dockerfile,
		Tags:       []string{img.ImageName}, //this should build the image locally, sans registry info
		Remove:     true,
		//CacheFrom:  img.Build.CachedImages, // TODO: this needs a login, so needs to be handled differently.
		BuildArgs: img.Build().Args,
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
			"dockerfile":     img.Build().Dockerfile,
			"context":        img.Build().Context,
			"registryServer": img.Registry.Server,
		}

		pbstruct, err := plugin.MarshalProperties(
			resource.NewPropertyMapFromMap(outputs),
			plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true},
		)
		return img.ImageName, pbstruct, err
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
	pushOutput, err := docker.ImagePush(ctx, img.ImageName, pushOpts)

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
		"dockerfile":     img.Build().Dockerfile,
		"context":        img.Build().Context,
		"baseImageName":  img.ImageName,
		"registryServer": img.Registry.Server,
	}
	pbstruct, err := plugin.MarshalProperties(
		resource.NewPropertyMapFromMap(outputs),
		plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true},
	)
	return img.ImageName, pbstruct, err
}
