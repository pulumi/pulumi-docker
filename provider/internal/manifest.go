package internal

import (
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Manifest struct {
	Digest   string           `pulumi:"digest"`
	Platform ManifestPlatform `pulumi:"platform"`
	Ref      string           `pulumi:"ref"`
	Size     int64            `pulumi:"size"`
}

type ManifestPlatform struct {
	OS           string `pulumi:"os"`
	Architecture string `pulumi:"architecture"`
}

func (mp *ManifestPlatform) Annotate(a infer.Annotator) {
	a.Describe(&mp.OS, dedent(`The manifest's operating systen.`))
	a.Describe(&mp.Architecture, dedent(`The manifest's architecture.`))
}

func (m *Manifest) Annotate(a infer.Annotator) {
	a.Describe(&m.Ref, "The manifest's canonical ref.")
	a.Describe(&m.Size, dedent(`The size of the manifest in bytes.`))
	a.Describe(&m.Platform, dedent(`The manifest's platform.`))
	a.Describe(&m.Digest, dedent(`The SHA256 digest of the manifest.`))
}

type RegistryAuth struct {
	Address  string `pulumi:"address"`
	Password string `pulumi:"password,optional" provider:"secret"`
	Username string `pulumi:"username,optional"`
}

func (ra *RegistryAuth) Annotate(a infer.Annotator) {
	a.Describe(&ra.Address, `The registry's address (e.g. "docker.io").`)
	a.Describe(&ra.Username, `Username for the registry.`)
	a.Describe(&ra.Password, `Password or token for the registry.`)
}
