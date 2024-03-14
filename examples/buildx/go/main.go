package main

import (
	"github.com/pulumi/pulumi-docker/sdk/v4/go/docker/buildx"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")
		dockerHubPassword := cfg.Require("dockerHubPassword")
		multiPlatform, err := buildx.NewImage(ctx, "multiPlatform", &buildx.ImageArgs{
			Dockerfile: &buildx.DockerfileArgs{
				Location: pulumi.String("app/Dockerfile.multiPlatform"),
			},
			Context: &buildx.BuildContextArgs{
				Location: pulumi.String("app"),
			},
			Platforms: buildx.PlatformArray{
				buildx.Platform_Plan9_amd64,
				buildx.Platform_Plan9_386,
			},
		})
		if err != nil {
			return err
		}
		_, err = buildx.NewImage(ctx, "registryPush", &buildx.ImageArgs{
			Context: &buildx.BuildContextArgs{
				Location: pulumi.String("app"),
			},
			Tags: pulumi.StringArray{
				pulumi.String("docker.io/pulumibot/buildkit-e2e:example"),
			},
			Exports: buildx.ExportArray{
				&buildx.ExportArgs{
					Registry: &buildx.ExportRegistryArgs{
						OciMediaTypes: pulumi.Bool(true),
						Push:          pulumi.Bool(false),
					},
				},
			},
			Registries: buildx.RegistryAuthArray{
				&buildx.RegistryAuthArgs{
					Address:  pulumi.String("docker.io"),
					Username: pulumi.String("pulumibot"),
					Password: pulumi.String(dockerHubPassword),
				},
			},
		})
		if err != nil {
			return err
		}
		_, err = buildx.NewImage(ctx, "cached", &buildx.ImageArgs{
			Context: &buildx.BuildContextArgs{
				Location: pulumi.String("app"),
			},
			CacheTo: buildx.CacheToArray{
				&buildx.CacheToArgs{
					Local: &buildx.CacheToLocalArgs{
						Dest: pulumi.String("tmp/cache"),
						Mode: buildx.CacheModeMax,
					},
				},
			},
			CacheFrom: buildx.CacheFromArray{
				&buildx.CacheFromArgs{
					Local: &buildx.CacheFromLocalArgs{
						Src: pulumi.String("tmp/cache"),
					},
				},
			},
		})
		if err != nil {
			return err
		}
		_, err = buildx.NewImage(ctx, "buildArgs", &buildx.ImageArgs{
			Dockerfile: &buildx.DockerfileArgs{
				Location: pulumi.String("app/Dockerfile.buildArgs"),
			},
			Context: &buildx.BuildContextArgs{
				Location: pulumi.String("app"),
			},
			BuildArgs: pulumi.StringMap{
				"SET_ME_TO_TRUE": pulumi.String("true"),
			},
		})
		if err != nil {
			return err
		}
		_, err = buildx.NewImage(ctx, "extraHosts", &buildx.ImageArgs{
			Dockerfile: &buildx.DockerfileArgs{
				Location: pulumi.String("app/Dockerfile.extraHosts"),
			},
			Context: &buildx.BuildContextArgs{
				Location: pulumi.String("app"),
			},
			AddHosts: pulumi.StringArray{
				pulumi.String("metadata.google.internal:169.254.169.254"),
			},
		})
		if err != nil {
			return err
		}
		_, err = buildx.NewImage(ctx, "sshMount", &buildx.ImageArgs{
			Dockerfile: &buildx.DockerfileArgs{
				Location: pulumi.String("app/Dockerfile.sshMount"),
			},
			Context: &buildx.BuildContextArgs{
				Location: pulumi.String("app"),
			},
			Ssh: buildx.SSHArray{
				&buildx.SSHArgs{
					Id: pulumi.String("default"),
				},
			},
		})
		if err != nil {
			return err
		}
		_, err = buildx.NewImage(ctx, "secrets", &buildx.ImageArgs{
			Dockerfile: &buildx.DockerfileArgs{
				Location: pulumi.String("app/Dockerfile.secrets"),
			},
			Context: &buildx.BuildContextArgs{
				Location: pulumi.String("app"),
			},
			Secrets: pulumi.StringMap{
				"password": pulumi.String("hunter2"),
			},
		})
		if err != nil {
			return err
		}
		_, err = buildx.NewImage(ctx, "labels", &buildx.ImageArgs{
			Context: &buildx.BuildContextArgs{
				Location: pulumi.String("app"),
			},
			Labels: pulumi.StringMap{
				"description": pulumi.String("This image will get a descriptive label 👍"),
			},
		})
		if err != nil {
			return err
		}
		_, err = buildx.NewImage(ctx, "target", &buildx.ImageArgs{
			Dockerfile: &buildx.DockerfileArgs{
				Location: pulumi.String("app/Dockerfile.target"),
			},
			Context: &buildx.BuildContextArgs{
				Location: pulumi.String("app"),
			},
			Target: pulumi.String("build-me"),
		})
		if err != nil {
			return err
		}
		_, err = buildx.NewImage(ctx, "namedContexts", &buildx.ImageArgs{
			Dockerfile: &buildx.DockerfileArgs{
				Location: pulumi.String("app/Dockerfile.namedContexts"),
			},
			Context: &buildx.BuildContextArgs{
				Location: pulumi.String("app"),
				Named: buildx.ContextMap{
					"golang:latest": &buildx.ContextArgs{
						Location: pulumi.String("docker-image://golang@sha256:b8e62cf593cdaff36efd90aa3a37de268e6781a2e68c6610940c48f7cdf36984"),
					},
				},
			},
		})
		if err != nil {
			return err
		}
		_, err = buildx.NewImage(ctx, "remoteContext", &buildx.ImageArgs{
			Context: &buildx.BuildContextArgs{
				Location: pulumi.String("https://raw.githubusercontent.com/pulumi/pulumi-docker/api-types/provider/testdata/Dockerfile"),
			},
		})
		if err != nil {
			return err
		}
		_, err = buildx.NewImage(ctx, "remoteContextWithInline", &buildx.ImageArgs{
			Dockerfile: &buildx.DockerfileArgs{
				Inline: pulumi.String("FROM busybox\nCOPY hello.c ./\n"),
			},
			Context: &buildx.BuildContextArgs{
				Location: pulumi.String("https://github.com/docker-library/hello-world.git"),
			},
		})
		if err != nil {
			return err
		}
		_, err = buildx.NewImage(ctx, "inline", &buildx.ImageArgs{
			Dockerfile: &buildx.DockerfileArgs{
				Inline: pulumi.String("FROM alpine\nRUN echo \"This uses an inline Dockerfile! 👍\"\n"),
			},
			Context: &buildx.BuildContextArgs{
				Location: pulumi.String("app"),
			},
		})
		if err != nil {
			return err
		}
		_, err = buildx.NewImage(ctx, "dockerLoad", &buildx.ImageArgs{
			Context: &buildx.BuildContextArgs{
				Location: pulumi.String("app"),
			},
			Exports: buildx.ExportArray{
				&buildx.ExportArgs{
					Docker: &buildx.ExportDockerArgs{
						Tar: pulumi.Bool(true),
					},
				},
			},
		})
		if err != nil {
			return err
		}
		ctx.Export("platforms", multiPlatform.Platforms)
		return nil
	})
}