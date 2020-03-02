package docker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunCommandThatCanFail(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		res := <-runCommandThatCanFail("echo", []string{"-n", "test"}, nil, true, true, nil)
		assert.Equal(t, res.stdout, "test")
		assert.Equal(t, res.code, 0)
	})

	t.Run("fail", func(t *testing.T) {
		res := <-runCommandThatCanFail("cat", []string{"not-a-real-file"}, nil, true, false, nil)
		assert.Empty(t, res.stdout)
		assert.Equal(t, res.code, 1)
	})
}
