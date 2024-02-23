package provider

import (
	_ "embed" // nolint:golint
)

//go:embed pkg/docs-gen/image-description.md
var docImageDescription string

//go:embed pkg/docs-gen/examples/image.md
var docImageExamples string

var docImage = docImageDescription + "\n\n" + docImageExamples
