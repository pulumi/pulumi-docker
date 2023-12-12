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
	"bytes"
	"encoding/json"
	"os"
	"os/exec"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

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
			randID, ok := stack.Outputs["extraArgument"]
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

func TestBuildOnPreviewYAML(t *testing.T) {
	cwd, err := os.Getwd()
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	var outputBuf bytes.Buffer
	integration.ProgramTest(t, &integration.ProgramTestOptions{
		Dir:                      path.Join(cwd, "test-build-on-preview", "yaml"),
		SkipUpdate:               true, //only run Preview
		SkipExportImport:         true,
		Verbose:                  true, //we need this to verify the build output logs
		AllowEmptyPreviewChanges: true,
		Stdout:                   &outputBuf,
		ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
			assert.Contains(t, outputBuf.String(), "Image built successfully, local id")
			assert.Contains(t, outputBuf.String(), "repoDigest:")
		},
	})
}
func TestDockerSwarmYAML(t *testing.T) {
	cwd, err := os.Getwd()
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	// Temporarily make ourselves a swarm manager.
	cmd := exec.Command("docker", "swarm", "init")
	output, err := cmd.CombinedOutput()
	require.NoError(t, err, string(output))
	t.Cleanup(func() {
		require.NoError(t, exec.Command("docker", "swarm", "leave", "--force").Run())
	})
	t.Run("service", func(t *testing.T) {
		integration.ProgramTest(t, &integration.ProgramTestOptions{
			Dir:         path.Join(cwd, "test-swarm", "service"),
			Quick:       true,
			SkipRefresh: true,
		})
	})

	t.Run("service-replicated", func(t *testing.T) {
		integration.ProgramTest(t, &integration.ProgramTestOptions{
			Dir:         path.Join(cwd, "test-swarm", "service-replicated"),
			Quick:       true,
			SkipRefresh: true,
		})
	})

	t.Run("service-global", func(t *testing.T) {
		integration.ProgramTest(t, &integration.ProgramTestOptions{
			Dir:         path.Join(cwd, "test-swarm", "service-global"),
			Quick:       true,
			SkipRefresh: true,
		})
	})
}

func TestUnknownsBuildOnPreviewWarnsYAML(t *testing.T) {
	cwd, err := os.Getwd()
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	var outputBuf bytes.Buffer
	integration.ProgramTest(t, &integration.ProgramTestOptions{
		Dir:                      path.Join(cwd, "test-unknowns", "yaml-build-on-preview"),
		SkipUpdate:               true, //only run Preview
		SkipExportImport:         true,
		Verbose:                  true, //we need this to verify the build output logs
		AllowEmptyPreviewChanges: true,
		Stderr:                   &outputBuf,
		ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
			assert.Contains(t, outputBuf.String(), "Minimum inputs for build are unresolved.")
		},
	})
}

func TestBuilderVersionsYAML(t *testing.T) {
	cwd, err := os.Getwd()
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	t.Run("v1", func(t *testing.T) {
		integration.ProgramTest(t, &integration.ProgramTestOptions{
			Dir:         path.Join(cwd, "test-builder-version", "v1"),
			Quick:       true,
			SkipRefresh: true,
			ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
				platform, ok := stack.Outputs["platform"]
				assert.True(t, ok)
				assert.NotEmpty(t, platform)
			},
		})
	})
	t.Run("v2", func(t *testing.T) {
		integration.ProgramTest(t, &integration.ProgramTestOptions{
			Dir:         path.Join(cwd, "test-builder-version", "v2"),
			Quick:       true,
			SkipRefresh: true,
			ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
				platform, ok := stack.Outputs["platform"]
				assert.True(t, ok)
				assert.NotEmpty(t, platform)
			},
		})
	})
}
