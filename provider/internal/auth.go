package internal

import "github.com/pulumi/pulumi-go-provider/infer"

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
