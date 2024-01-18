package internal

import (
	"context"
	"testing"

	"github.com/blang/semver"
	"github.com/stretchr/testify/assert"

	provider "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi-go-provider/integration"
)

func TestConfigure(t *testing.T) {
	s := newServer()

	err := s.Configure(
		provider.ConfigureRequest{},
	)
	assert.NoError(t, err)
}

// TestAnnotate sanity checks that our annotations don't panic.
func TestAnnotate(_ *testing.T) {
	for _, tt := range []infer.Annotated{
		&Config{},
		&Image{},
		&ImageArgs{},
		&ImageState{},
	} {
		tt.Annotate(annotator{})
	}
}

// TestSchema sanity checks that our schema doesn't panic.
func TestSchema(_ *testing.T) {
	Schema(context.Background(), "v4")
}

type annotator struct{}

func (annotator) Describe(_ any, _ string)             {}
func (annotator) SetDefault(_ any, _ any, _ ...string) {}
func (annotator) SetToken(_, _ string)                 {}

func newServer() integration.Server {
	provider := NewBuildxProvider()
	return integration.NewServer("docker", semver.Version{Major: 4}, provider)
}
