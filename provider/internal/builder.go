package internal

import (
	"github.com/muesli/reflow/dedent"

	"github.com/pulumi/pulumi-go-provider/infer"
)

var _ = (infer.Annotated)((*BuilderConfig)(nil))

type BuilderConfig struct {
	Name string `pulumi:"name,optional"`
}

func (b *BuilderConfig) Annotate(a infer.Annotator) {
	a.Describe(&b.Name, dedent.String(`
		Name of an existing builder to use.
	`))
}
