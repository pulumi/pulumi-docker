package internal

import (
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Dockerfile struct {
	Location string `pulumi:"location,optional"`
	Inline   string `pulumi:"inline,optional"`
}

func (d *Dockerfile) Annotate(a infer.Annotator) {
	a.Describe(&d.Location, dedent(`
		Location of the Dockerfile to use.

		Can be a relative or absolute path to a local file, or a remote URL.
		
		Defaults to "${context.location}/Dockerfile" if context is on-disk.

		Conflicts with "inline".
	`))
	a.Describe(&d.Inline, dedent(`
		Raw Dockerfile contents.
		
		Conflicts with "location".

		Equivalent to invoking Docker with "-f -".
	`))
}
