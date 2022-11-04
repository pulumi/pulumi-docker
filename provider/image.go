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

func (p *dockerNativeProvider) dockerBuild(ctx context.Context,
	urn resource.URN,
	props *structpb.Struct) (*structpb.Struct, error) {
	fmt.Println("resource URN", urn)

	fmt.Println("getting inputs...")

	inputs, err := plugin.UnmarshalProperties(props, plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true})
	if err != nil {
		return nil, err
	}

	imageName := inputs["imageName"].StringValue()
	//tag := inputs["tag"].StringValue()
	registry := inputs["registry"].ObjectValue()
	username := registry["username"].StringValue()
	password := registry["password"].StringValue()
	server := registry["server"].StringValue()

	// build can be a string or an object; we will also use reasonable defaults here.
	var buildContext string
	var dockerfile string

	if inputs["build"].IsNull() {
		// use the filepath and the default "Dockerfile" for the Dockerfile
		buildContext = "."
		dockerfile = "Dockerfile"
	} else if inputs["build"].IsString() {
		// use the filepath and the default "Dockerfile" for the Dockerfile
		buildContext = inputs["build"].StringValue()
		dockerfile = "Dockerfile" // default to Dockerfile
	} else {
		// read in the build type fields
		build := inputs["build"].ObjectValue()
		dockerfile = build["dockerfile"].StringValue()
		buildContext = build["context"].StringValue()
	}

	fmt.Println("USING THE DOCKER CLIENT NOW")

	cli, err := client.NewClientWithOpts(client.FromEnv)

	if err != nil {
		panic(err)
	}
	//fmt.Println("Get list of  existing images")
	//imgs, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println("List image labels")
	//for _, img := range imgs {
	//	fmt.Println("whole image: ", img.RepoTags) // RepoTags gives a list of tags that correspond to the REPOSITORY and TAGS headers of the "docker images" output
	//}

	fmt.Println("now let's BUILD an image")
	imageTag := username + "/" + imageName

	// make the build context
	tar, err := archive.TarWithOptions(buildContext, &archive.TarOptions{})
	if err != nil {
		panic(err)
	}

	// make the build options
	opts := types.ImageBuildOptions{
		Dockerfile: dockerfile,
		Tags:       []string{imageTag},
		Remove:     true,
	}

	imgBuildResp, err := cli.ImageBuild(context.Background(), tar, opts)
	if err != nil {
		panic(err)
	}
	// Print build logs to terminal
	scanner := bufio.NewScanner(imgBuildResp.Body)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	fmt.Println("see if we can push to the registry")

	// Quick and dirty auth; we can also preconfigure the client itself I believe

	var authConfig = types.AuthConfig{
		Username:      username,
		Password:      password,
		ServerAddress: server,
	}

	authConfigBytes, _ := json.Marshal(authConfig)
	authConfigEncoded := base64.URLEncoding.EncodeToString(authConfigBytes)

	pushOpts := types.ImagePushOptions{RegistryAuth: authConfigEncoded}

	pushOutput, err := cli.ImagePush(context.Background(), imageTag, pushOpts)

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
		"dockerfile":     dockerfile,
		"context":        buildContext,
		"name":           imageTag,
		"registryServer": server,
	}
	return plugin.MarshalProperties(
		resource.NewPropertyMapFromMap(outputs),
		plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true},
	)
}
