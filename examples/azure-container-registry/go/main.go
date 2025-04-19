package main

import (
	"github.com/pulumi/pulumi-azure/sdk/v6/go/azure/containerservice"
	"github.com/pulumi/pulumi-azure/sdk/v6/go/azure/core"
	"github.com/pulumi/pulumi-docker/sdk/v4/go/docker"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a private ACR registry.
		rg, err := core.NewResourceGroup(ctx, "myrg", nil)
		if err != nil {
			return err
		}
		registry, err := containerservice.NewRegistry(ctx, "myregistry", &containerservice.RegistryArgs{
			ResourceGroupName: rg.Name,
			AdminEnabled:      pulumi.Bool(true),
			Sku:               pulumi.String("Basic"),
		})
		if err != nil {
			return err
		}
		imageArgs := &docker.ImageArgs{
			ImageName: pulumi.Sprintf("%s/myimage", registry.LoginServer),
			Build: docker.DockerBuildArgs{
				Context: pulumi.String("./app"),
			},
			SkipPush: pulumi.Bool(false),
			Registry: &docker.RegistryArgs{
				Server:   registry.LoginServer,
				Username: registry.AdminUsername,
				Password: registry.AdminPassword,
			},
		}
		image, err := docker.NewImage(ctx, "myimage", imageArgs)
		if err != nil {
			return err
		}

		ctx.Export("deps-image", image.ImageName)
		ctx.Export("repoDigest", image.RepoDigest)

		return nil
	})
}
