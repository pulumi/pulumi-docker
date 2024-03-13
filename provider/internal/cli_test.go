package internal

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExec(t *testing.T) {
	h, err := newHost(nil)
	require.NoError(t, err)
	cli, err := wrap(h)
	require.NoError(t, err)

	err = cli.exec([]string{"buildx", "version"}, nil)
	assert.NoError(t, err)

	out, err := io.ReadAll(cli.r)
	require.NoError(t, err)
	assert.Contains(t, string(out), "github.com/docker/buildx")
}
