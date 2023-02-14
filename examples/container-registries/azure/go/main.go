//package main

//import (
//	"github.com/pulumi/pulumi-azure/sdk/v2/go/azure/authorization"
//	"github.com/pulumi/pulumi-azure/sdk/v2/go/azure/containerservice"
//	"github.com/pulumi/pulumi-azure/sdk/v2/go/azure/core"
//	"github.com/pulumi/pulumi-azuread/sdk/v2/go/azuread"
//	"github.com/pulumi/pulumi-docker/sdk/v2/go/docker"
//	"github.com/pulumi/pulumi-random/sdk/v2/go/random"
//	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v2/go/pulumi/config"
//)
//
//func main() {
//	pulumi.Run(func(ctx *pulumi.Context) error {
//		// Conditionalize the auth mechanism.
//		//useServicePrincipalAuth := config.GetBool(ctx, "useServicePrincipalAuth")
//
//		// Create a private ACR registry.
//		rg, err := core.NewResourceGroup(ctx, "myrg", nil)
//		if err != nil {
//			return err
//		}
//		registry, err := containerservice.NewRegistry(ctx, "my-repo", &containerservice.RegistryArgs{
//			ResourceGroupName: rg.Name,
//			AdminEnabled:      pulumi.Bool(true),
//			Sku:               pulumi.String("Basic"),
//		})
//		if err != nil {
//			return err
//		}
//
//		// Get registry info (creds and endpoint).
//		imageName := pulumi.Sprintf("%s/myapp", registry.LoginServer)
//		var registryInfo docker.ImageRegistryArgs
//		if useServicePrincipalAuth {
//			spApp, err := azuread.NewApplication(ctx, "myspapp", nil)
//			if err != nil {
//				return err
//			}
//			sp, err := azuread.NewServicePrincipal(ctx, "mysp", &azuread.ServicePrincipalArgs{
//				ApplicationId: spApp.ApplicationId,
//			})
//			if err != nil {
//				return err
//			}
//			password, err := random.NewRandomPassword(ctx, "mypass",
//				&random.RandomPasswordArgs{
//					Length: pulumi.Int32(32),
//				},
//				pulumi.AdditionalSecretOutputs([]string{"result"}),
//			)
//			if err != nil {
//				return err
//			}
//			spPassword, err := azuread.NewServicePrincipalPassword(ctx, "mysp-pass", &azuread.ServicePrincipalPasswordArgs{
//				ServicePrincipalId: sp.ID(),
//				Value:              password.Result,
//				EndDateRelative:    pulumi.String("8760h"),
//			})
//			if err != nil {
//				return err
//			}
//			apAuth, err := authorization.NewAssignment(ctx, "myauth", &authorization.AssignmentArgs{
//				Scope:              registry.ID(),
//				RoleDefinitionName: pulumi.String("acrpush"),
//				PrincipalId:        sp.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			registryInfo = docker.ImageRegistryArgs{
//				Server:   registry.LoginServer,
//				Username: sp.ApplicationId,
//				Password: spPassword.Value,
//			}
//		}
//
//		// Build and publish the app image.
//		image, err := docker.NewImage(ctx, "my-image", &docker.ImageArgs{
//			Build:     &docker.DockerBuildArgs{Context: pulumi.String("app")},
//			ImageName: imageName,
//			Registry:  registryInfo,
//		})
//
//		// Export the resulting base name in addition to the specific version pushed.
//		ctx.Export("baseImageName", image.BaseImageName)
//		ctx.Export("fullImageName", image.ImageName)
//		return nil
//	})
//}

package main

import (
	"github.com/pulumi/pulumi-azure/sdk/v5/go/azure/containerservice"
	"github.com/pulumi/pulumi-azure/sdk/v5/go/azure/core"
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

		return nil
	})
}
