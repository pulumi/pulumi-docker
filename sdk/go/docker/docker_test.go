package docker

import (
	"context"
	"testing"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"github.com/stretchr/testify/assert"
)

func TestRunCommandThatMustSucceed(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		stdout, err := runBasicCommandThatMustSucceed("echo", []string{"-n", "test"}, nil)
		assert.Nil(t, err)
		assert.Equal(t, "test", stdout)
	})

	t.Run("fail", func(t *testing.T) {
		_, err := runBasicCommandThatMustSucceed("cat", []string{"not-a-real-file"}, nil)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "cat not-a-real-file failed with error: exit status 1")
	})
}

func TestRunDockerBuild(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		build := &dockerBuild{
			Context:    ".",
			Dockerfile: "./tests/Dockerfile",
		}
		err := runDockerBuild("test", build, nil, nil, "")
		assert.Nil(t, err)
	})
}

/*
// This test was created for local testing requires valid credentials to run.
// Not intended to be used for CI/CD testing until we can generate those
// credentials within this file.
func TestLoginToRegistry(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		registry := imageRegistry{
			Server:   "insert-your-server",
			Username: "insert-your-username",
			Password: "insert-your-password",
		}

		err := loginToRegistry(registry, nil)
		assert.Nil(t, err)
	})
}
*/

func TestBuildImageAsync(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		build := dockerBuild{
			Dockerfile: "./tests/Dockerfile",
		}
		output, stages, err := buildImageAsync("test", build, nil, nil)
		assert.Nil(t, err)
		assert.NotContains(t, output, ":")
		assert.Len(t, output, 64)
		assert.Empty(t, stages)
	})
}

func TestNewImage(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		ctx, err := pulumi.NewContext(context.Background(), pulumi.RunInfo{
			DryRun: true,
		})
		assert.Nil(t, err)

		// Create the docker image.
		imageArgs := ImageArgs{
			ImageName: pulumi.String("registry3ac31be0.azurecr.io/mynodeapp:v1.0.0"),
			Build: DockerBuild{
				Context: pulumi.String("./app"),
			},
			Registry: &ImageRegistry{
				Password: pulumi.String("jS=jDd/P+fhP39YTlwIhMVpQ9udCPmQ0"),
				Server:   pulumi.String("registry3ac31be0.azurecr.io"),
				Username: pulumi.String("registry3ac31be0"),
			},
		}
		image, err := NewImage(ctx, "node-app", &imageArgs)
		assert.Nil(t, err)
		t.Log(image.ImageName)
	})
}

func TestNewRemoteImage(t *testing.T) {
	ctx, err := pulumi.NewContext(context.Background(), pulumi.RunInfo{
		DryRun: true,
	})
	assert.Nil(t, err)

	imageArgs := &RemoteImageArgs{
		Name:        pulumi.String("nginx"),
		KeepLocally: pulumi.Bool(true),
	}
	_, err = NewRemoteImage(ctx, "nginx-image", imageArgs)
	assert.Nil(t, err)
}
