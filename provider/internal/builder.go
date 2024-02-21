package internal

import (
	"github.com/pulumi/pulumi-go-provider/infer"
)

var _ = (infer.Annotated)((*BuilderConfig)(nil))

type BuilderConfig struct {
	Name string `pulumi:"name,optional"`
}

func (b *BuilderConfig) Annotate(a infer.Annotator) {
	a.Describe(&b.Name, dedent(`
		Name of an existing buildx builder to use.
		
		Only "docker-container", "kubernetes", or "remote" drivers are
		supported. The legacy "docker" driver is not supported.

		Equivalent to Docker's "--builder" flag.
	`))
}
