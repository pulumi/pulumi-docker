package internal

import (
	"context"
	"testing"

	"github.com/blang/semver"
	"github.com/stretchr/testify/assert"

	provider "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi-go-provider/integration"
	mwcontext "github.com/pulumi/pulumi-go-provider/middleware/context"
)

func TestConfigure(t *testing.T) {
	s := newServer(nil)

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
		&Index{},
		&IndexArgs{},
		&IndexState{},
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

func newServer(client Client) integration.Server {
	p := NewBuildxProvider()

	// Inject a mock client if provided.
	if client != nil {
		p = mwcontext.Wrap(p, func(ctx provider.Context) provider.Context {
			return provider.CtxWithValue(ctx, _mockClientKey, client)
		})
	}

	return integration.NewServer("docker", semver.Version{Major: 4}, p)
}
