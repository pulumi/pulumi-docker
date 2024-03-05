package provider

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeProvider(t *testing.T) {
	_, err := makeProvider(nil, "docker", "v4", []byte{})
	assert.NoError(t, err)
}
