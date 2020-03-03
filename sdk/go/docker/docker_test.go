package docker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunCommandThatMustSucceed(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		stdout, err := RunCommandThatMustSucceed("echo", []string{"-n", "test"}, nil, true, nil)
		assert.Nil(t, err)
		assert.Equal(t, "test", stdout)
	})

	t.Run("fail", func(t *testing.T) {
		_, err := RunCommandThatMustSucceed("cat", []string{"not-a-real-file"}, nil, true, nil)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "cat not-a-real-file failed with error: exit status 1")
	})
}
