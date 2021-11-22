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
		"Building image '.'...\n",
		"Successfully tagged pulumi-user/example:v1.0.0\n",
		"Image build succeeded.\n",
	}

	var actualEphemeralDiagnosticsLines []string
	for _, ev := range stack.Events {
		if ev.DiagnosticEvent != nil && ev.DiagnosticEvent.Ephemeral && ev.DiagnosticEvent.Severity == "info" {
			actualEphemeralDiagnosticsLines = append(actualEphemeralDiagnosticsLines, ev.DiagnosticEvent.Message)
		}
	}
	assert.Subset(t, actualEphemeralDiagnosticsLines, expectedEphemeralDiagnosticsLines)
}
