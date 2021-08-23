// Copyright 2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package examples

import (
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"github.com/stretchr/testify/assert"
)

func dockerFileWithDependenciesOutputValidation(t *testing.T, stack integration.RuntimeValidationStackInfo) {
	// The full set of lines that are currently produced  are below, but some are commented out as they may not be
	// stable going forward in case of changes to Docker.  To be safe, we just test a subset of this output that is
	// enough to verify no structural regressions in the output.
	expectedEphemeralDiagnosticsLines := []string{
		"<{%reset%}>Building image '.'...<{%reset%}>\n",
		//"<{%reset%}>Sending build context to Docker daemon  66.88MB<{%reset%}>\n",
		// "<{%reset%}><{%reset%}>\n",
		"<{%reset%}>Step 1/3 : FROM python:3.6 AS base<{%reset%}>\n",
		// "<{%reset%}> ---> bd4a91d81d7e<{%reset%}>\n",
		// "<{%reset%}>Step 2/3 : FROM base AS dependencies<{%reset%}>\n",
		// "<{%reset%}> ---> bd4a91d81d7e<{%reset%}>\n",
		// "<{%reset%}>Step 3/3 : RUN pip install gunicorn<{%reset%}>\n",
		// "<{%reset%}> ---> Using cache<{%reset%}>\n",
		// "<{%reset%}> ---> 584bd8f8844a<{%reset%}>\n",
		// "<{%reset%}>Successfully built 584bd8f8844a<{%reset%}>\n",
		"<{%reset%}>Successfully tagged pulumi-user/example:v1.0.0<{%reset%}>\n",
		// "<{%reset%}>sha256:584bd8f8844aeb5bc815c8c416b586a4334a577d49b1a4552a93765537c39151<{%reset%}>\n",
		"<{%reset%}>Image build succeeded.<{%reset%}>\n",
	}

	var actualEphemeralDiagnosticsLines []string
	for _, ev := range stack.Events {
		if ev.DiagnosticEvent != nil && ev.DiagnosticEvent.Ephemeral && ev.DiagnosticEvent.Severity == "info" {
			actualEphemeralDiagnosticsLines = append(actualEphemeralDiagnosticsLines, ev.DiagnosticEvent.Message)
		}
	}
	assert.Subset(t, actualEphemeralDiagnosticsLines, expectedEphemeralDiagnosticsLines)
}
