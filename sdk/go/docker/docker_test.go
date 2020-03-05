package docker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunCommandThatMustSucceed(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		stdout, err := RunCommandThatMustSucceed("echo", []string{"-n", "test"}, nil, true, "", nil)
		assert.Nil(t, err)
		assert.Equal(t, "test", stdout)
	})

	t.Run("fail", func(t *testing.T) {
		_, err := RunCommandThatMustSucceed("cat", []string{"not-a-real-file"}, nil, true, "", nil)
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
