// Copyright 2016-2023, Pulumi Corporation.
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
//go:build go || all
// +build go all

package examples

import (
	"encoding/json"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func TestUnknownInputsYAML(t *testing.T) {
	cwd, err := os.Getwd()
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	integration.ProgramTest(t, &integration.ProgramTestOptions{
		Dir:         path.Join(cwd, "test-unknowns", "yaml"),
		Quick:       true,
		SkipRefresh: true,
		ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
			randID, ok := stack.Outputs["randNameId"]
			assert.True(t, ok)
			assert.NotEmpty(t, randID)
		},
	})
}

func TestSecretsYAML(t *testing.T) {
	cwd, err := os.Getwd()
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	integration.ProgramTest(t, &integration.ProgramTestOptions{
		Dir:         path.Join(cwd, "test-secrets", "yaml"),
		Quick:       true,
		SkipRefresh: true,
		ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
			imgName, ok := stack.Outputs["imageName"]
			assert.True(t, ok)
			assert.NotEmpty(t, imgName)

			// imgName is going to be a secret value encoded as a map. Currently ProgramTest lacks the
			// capacity to decrypt it and check the contents. For now we simply ensure that the secret
			// information is not present plaintext in the JSON expansion of this value.
			imgNameStr, err := json.Marshal(imgName)
			assert.NoError(t, err)
			assert.NotContains(t, imgNameStr, "pulumibot/test-secrets:yaml")

			// Make sure that state file does not contain secrets in plain.
			deploymentJSON, err := json.MarshalIndent(stack.Deployment, "", "  ")
			assert.NoError(t, err)
			assert.NotContainsf(t, string(deploymentJSON), "supersecret",
				"Secret should not be stored in the plain state")
		},
	})
}
