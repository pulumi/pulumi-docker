package internal

import (
	"strings"

	dd "github.com/muesli/reflow/dedent"
)

func dedent(s string) string {
	return strings.TrimSpace(dd.String(
		strings.Replace(s, `"`, "`", -1),
	))
}
