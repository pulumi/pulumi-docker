package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDedent(t *testing.T) {
	tests := []struct {
		name  string
		given string
		want  string
	}{
		{
			name: "simple case",
			given: `
		An optional map of named build-time argument variables to set during
		the Docker build. This flag allows you to pass build-time variables that
		can be accessed like environment variables inside the "RUN"
		instruction.`,
			want: `An optional map of named build-time argument variables to set during
the Docker build. This flag allows you to pass build-time variables that
can be accessed like environment variables inside the ` + "`RUN`\n" + `instruction.`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := dedent(tt.given)
			assert.Equal(t, tt.want, actual)
		})
	}
}
