package internal

import (
	"strings"

	"github.com/pulumi/pulumi-go-provider/infer"
)

type SSH struct {
	ID    string   `pulumi:"id"`
	Paths []string `pulumi:"paths,optional"`
}

func (s *SSH) Annotate(a infer.Annotator) {
	a.Describe(&s.ID, dedent(`
		Useful for distinguishing different servers that are part of the same
		build.

		A value of "default" is appropriate if only dealing with a single host.
	`))
	a.Describe(&s.Paths, dedent(`
		SSH agent socket or private keys to expose to the build under the given
		identifier.

		Defaults to "[$SSH_AUTH_SOCK]".

		Note that your keys are **not** automatically added when using an
		agent. Run "ssh-add -l" locally to confirm which public keys are
		visible to the agent; these will be exposed to your build.
	`))
}

func (s SSH) String() string {
	if s.ID == "" {
		return ""
	}

	r := s.ID

	if len(s.Paths) > 0 {
		r += "=" + strings.Join(s.Paths, ",")
	}

	return r
}
