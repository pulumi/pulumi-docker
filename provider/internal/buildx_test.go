package internal

import (
	"testing"

	"github.com/blang/semver"
	provider "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/integration"
	"github.com/stretchr/testify/assert"
)

func TestConfigure(t *testing.T) {
	s := newServer()

	err := s.Configure(
		provider.ConfigureRequest{},
	)
	assert.NoError(t, err)
}

func TestAnnotate(t *testing.T) {
	c := &Config{}
	c.Annotate(annotator{})
}

type annotator struct{}

func (annotator) Describe(_ any, _ string)             {}
func (annotator) SetDefault(_ any, _ any, _ ...string) {}
func (annotator) SetToken(_, _ string)                 {}

func newServer() integration.Server {
	provider := NewBuildxProvider()
	return integration.NewServer("docker", semver.Version{Major: 4}, provider)
}
