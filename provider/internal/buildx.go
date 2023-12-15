package internal

import (
	provider "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
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
				imageResource(),
			},
			ModuleMap: map[tokens.ModuleName]tokens.ModuleName{
				"internal": "buildx/image",
			},
			Config: infer.Config[*Config](),
		},
	)
}

// ImageSchema returns Image's schema for SDK generation.
func ImageSchema() schema.ResourceSpec {
	r := imageResource()
	s, err := r.GetSchema(nil) // We don't have subtypes.
	contract.AssertNoErrorf(err, "missing schema")
	return s
}

func imageResource() infer.InferredResource {
	return infer.Resource[*Image, ImageArgs, ImageState]()
}
