package main

import (
	"fmt"

	"github.com/pulumi/pulumi-docker/sdk/v4/go/docker"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Get a reference to the remote image "nginx:1.15.6". Without specifying the repository, the Docker provider will
		// try to download it from the public Docker Hub.
		imageArgs := &docker.RemoteImageArgs{
			Name:        pulumi.String("nginx"),
			KeepLocally: pulumi.Bool(true),
		}
		image, err := docker.NewRemoteImage(ctx, "nginx-image", imageArgs)
		if err != nil {
			return err
		}

		// Launch a container using the nginx image we just downloaded.
		portArgs := docker.ContainerPortArgs{
			Internal: pulumi.Int(80),
		}
		containerArgs := &docker.ContainerArgs{
			Image: image.Latest,
			Ports: docker.ContainerPortArray([]docker.ContainerPortInput{portArgs}),
		}
		container, err := docker.NewContainer(ctx, "nginx", containerArgs)
		if err != nil {
			return err
		}

		// Since the container is auto-named, export the name.
		ctx.Export("name", container.Name)

		// Since the provider picked a random ephemeral port for this container, export the port.
		port := container.Ports.Index(pulumi.Int(0))
		endpoint := port.ApplyT(func(port docker.ContainerPort) string {
			if port.Ip == nil || port.External == nil {
				return ""
			}
			return fmt.Sprintf("%s:%d", *port.Ip, *port.External)
		})
		ctx.Export("endpoint", endpoint)

		return nil
	})
}
