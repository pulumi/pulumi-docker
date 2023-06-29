package provider

import (
	"context"
	"fmt"
	actualbuild "github.com/docker/buildx/controller/build"
	controllerapi "github.com/docker/buildx/controller/pb"
	_ "github.com/docker/buildx/driver/docker"
	_ "github.com/docker/buildx/driver/docker-container"
	_ "github.com/docker/buildx/driver/kubernetes"
	_ "github.com/docker/buildx/driver/remote"
	"github.com/docker/buildx/util/progress"
	"github.com/docker/cli/cli/command"
	cliflags "github.com/docker/cli/cli/flags"
	"os"
)

func runBuildx(build Build, img Image) {
	fmt.Println("🦉🦉 build", build)
	fmt.Println("🦉🦉image", img)
	ctx := context.Background()

	// some CLI that is just the docker thing
	cli, err := command.NewDockerCli()
	if err != nil {
		fmt.Println(err.Error())
	}

	err = cli.Initialize(cliflags.NewClientOptions())

	fmt.Println(cli.ConfigFile())

	controllerOpts := controllerapi.BuildOptions{
		ContextPath:    build.Context,
		DockerfileName: build.Dockerfile,
		//PrintFunc:      nil,
		//NamedContexts:  nil,
		//Allow:          nil,
		//Attests:        nil,
		//BuildArgs:      nil,
		//CacheFrom:      nil,
		//CacheTo:        nil,
		//CgroupParent:   "",
		//Exports:        nil,
		//ExtraHosts:     nil,
		//Labels:         nil,
		//NetworkMode:    "",
		//NoCacheFilter:  nil,
		Platforms: build.Platform,
		//Secrets:        nil,
		//ShmSize:        0,
		//SSH:            nil,
		Tags: []string{img.Name},
		//Target:         "",
		//Ulimits:        nil,
		Builder: "busy_sammet",
		//NoCache:        false,
		//Pull:           false,
		ExportPush: true, // we want to always push in this case
		//ExportLoad:     false,
		//SourcePolicy:   nil,
	}

	// I got this from here: https://github.com/docker/buildx/blob/master/commands/build.go#L254

	ctx2, cancel := context.WithCancel(context.TODO())

	defer cancel()
	// TODO: implement progress mode
	//progressMode, err := options.toProgress()
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	var printer *progress.Printer
	// TODO: somehow hook this up to pulumi.Info
	printer, err = progress.NewPrinter(ctx2, os.Stderr, os.Stderr, "this is a progress mode",
		progress.WithDesc(
			fmt.Sprintf("building with %q instance using %s driver", "test", "default"), fmt.Sprintf("second print statement"),
		),
		progress.WithOnClose(func() {
			fmt.Sprintf("this is WithOnClose")
		}),
	)
	if err != nil {
		fmt.Println(err.Error())
	}

	resp, res, buildErr := actualbuild.RunBuild(ctx, cli, controllerOpts, cli.In(), printer, true)
	if resp != nil {
		fmt.Println("resp was not nil")
	}
	if res != nil {
		fmt.Println("res was not nil")
	}
	if buildErr != nil {
		fmt.Println("hitting the build error")
		fmt.Println(buildErr.Error())
	}
	fmt.Println("we got to the end sheesh")

}
