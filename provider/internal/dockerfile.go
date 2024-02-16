package internal

import (
	"github.com/muesli/reflow/dedent"

	"github.com/pulumi/pulumi-go-provider/infer"
)

type Dockerfile struct {
	Location string `pulumi:"location,optional"`
	Inline   string `pulumi:"inline,optional"`
}

func (d *Dockerfile) Annotate(a infer.Annotator) {
	a.Describe(&d.Location, dedent.String(`
		Name of the Dockerfile to use (defaults to "${context}/Dockerfile").
		Conflicts with "inline".`,
	))
	a.Describe(&d.Inline, dedent.String(`
		Raw Dockerfile contents. Conflicts with "location".`,
	))
}
