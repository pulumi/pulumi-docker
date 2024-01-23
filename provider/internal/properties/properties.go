package properties

import (
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Manifest struct {
	Digest   string   `pulumi:"digest"`
	Platform Platform `pulumi:"platform"`
	Ref      string   `pulumi:"ref"`
	Size     int64    `pulumi:"size"`
	URLs     []string `pulumi:"urls"`
}

type Platform struct {
	OS           string `pulumi:"os"`
	Architecture string `pulumi:"architecture"`
}

func (m *Manifest) Annotate(a infer.Annotator) {
	// a.SetToken("buildx", "Manifest")
	a.Describe(&m.Ref, "The manifest's ref")
}
