package internal

import (
	"context"
	"fmt"

	provider "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	pschema "github.com/pulumi/pulumi-go-provider/middleware/schema"
	gen "github.com/pulumi/pulumi/pkg/v3/codegen/go"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
)

var (
	_ infer.CustomConfigure = (*Config)(nil)
	_ infer.Annotated       = (infer.Annotated)((*Config)(nil))
	_ infer.Annotated       = (infer.Annotated)((*RegistryAuth)(nil))
)

// Config configures the buildx provider.
type Config struct {
	Host         string         `pulumi:"host,optional"`
	RegistryAuth []RegistryAuth `pulumi:"registryAuth,optional"`

	host *host
}

// _mockClientKey is used by tests to inject a mock Docker client.
var _mockClientKey struct{}

// Annotate provides user-facing descriptions and defaults for Config's fields.
func (c *Config) Annotate(a infer.Annotator) {
	a.Describe(&c.Host, "The build daemon's address.")
	a.SetDefault(&c.Host, "", "DOCKER_HOST")
}

// Configure validates and processes user-provided configuration values.
func (c *Config) Configure(_ provider.Context) error {
	h, err := newHost(c)
	if err != nil {
		return fmt.Errorf("getting host: %w", err)
	}
	c.host = h
	return nil
}

// NewBuildxProvider returns a new buildx provider.
func NewBuildxProvider() provider.Provider {
	return infer.Provider(
		infer.Options{
			Metadata: pschema.Metadata{
				Keywords: []string{"docker", "buildkit", "buildx"},
				LanguageMap: map[string]any{
					"go": gen.GoPackageInfo{
						Generics: gen.GenericsSettingGenericsOnly,
					},
				},
			},
			Resources: []infer.InferredResource{
				infer.Resource[*Image](),
				infer.Resource[*Index](),
			},
			ModuleMap: map[tokens.ModuleName]tokens.ModuleName{
				"internal": "buildx/image",
			},
			Config: infer.Config[*Config](),
		},
	)
}

// Schema returns our package specification.
func Schema(ctx context.Context, version string) schema.PackageSpec {
	p := NewBuildxProvider()
	spec, err := provider.GetSchema(ctx, "docker", version, p)
	contract.AssertNoErrorf(err, "missing schema")
	return spec
}
