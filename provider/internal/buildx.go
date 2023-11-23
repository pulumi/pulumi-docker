package internal

import (
	provider "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
)

// Config configures the buildx provider.
type Config struct {
	Host string `pulumi:"host,optional"`
}

// Annotate provides user-facing descriptions and defaults for Config's fields.
func (c *Config) Annotate(a infer.Annotator) {
	a.Describe(&c.Host, "The build daemon's address.")
	a.SetDefault(&c.Host, "", "DOCKER_HOST")
}

// Configure validates and processes user-provided configuration values.
func (c *Config) Configure(_ provider.Context) error {
	return nil
}

// NewBuildxProvider returns a new buildx provider.
func NewBuildxProvider() provider.Provider {
	return infer.Provider(
		infer.Options{
			Resources: []infer.InferredResource{
				infer.Resource[*Image, ImageArgs, *ImageState](),
			},
			ModuleMap: map[tokens.ModuleName]tokens.ModuleName{
				"internal": "buildx/image",
			},
			Config: infer.Config[*Config](),
		},
	)
}
