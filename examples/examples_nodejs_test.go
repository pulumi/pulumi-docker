// Copyright 2016-2018, Pulumi Corporation.
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
//go:build nodejs || all
// +build nodejs all

package examples

import (
	"encoding/base64"
	"encoding/json"
	"os"
	"path"
	"testing"
	"time"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"github.com/stretchr/testify/assert"
)

func TestNginxTs(t *testing.T) {
	test := getJsOptions(t).
		With(integration.ProgramTestOptions{
			Dir:                  path.Join(getCwd(t), "nginx"),
			ExpectRefreshChanges: true,
			Quick:                true,
		})

	integration.ProgramTest(t, &test)
}

func TestDockerfileWithTarget(t *testing.T) {
	test := getJsOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "dockerfile-with-target"),
		})

	integration.ProgramTest(t, &test)
}

func TestAzureContainerRegistry(t *testing.T) {
	location := os.Getenv("AZURE_LOCATION")
	if location == "" {
		t.Skipf("Skipping test due to missing AZURE_LOCATION environment variable")
	}
	test := getJsOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "azure-container-registry/ts"),
			Config: map[string]string{
				"azure:environment": "public",
				"azure:location":    location,
			},
			ExpectRefreshChanges: true,
		})

	integration.ProgramTest(t, &test)
}

func TestAwsContainerRegistry(t *testing.T) {
	region := os.Getenv("AWS_REGION")
	if region == "" {
		t.Skipf("Skipping test due to missing AWS_REGION environment variable")
	}
	test := getJsOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "aws-container-registry/ts"),
			Config: map[string]string{
				"aws:region": region,
			},
			ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
				digest, ok := stack.Outputs["repoDigest"].(string)
				assert.True(t, ok)
				assert.NotEmpty(t, digest)
			},
		})

	integration.ProgramTest(t, &test)
}

func TestDigitaloceanContainerRegistry(t *testing.T) {
	t.Skipf("Skipping test due to known storageUsageBytes issue #TODO")

	token := os.Getenv("DIGITALOCEAN_TOKEN")
	if token == "" {
		t.Skipf("Skipping test due to missing DIGITALOCEAN_TOKEN environment variable")
	}
	test := getJsOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "digitalocean-container-registry/ts"),
			Config: map[string]string{
				"digitalocean:token": token,
			},
		})

	integration.ProgramTest(t, &test)
}

func TestGcpContainerRegistry(t *testing.T) {
	project := os.Getenv("GOOGLE_PROJECT")
	if project == "" {
		t.Skipf("Skipping test due to missing GOOGLE_PROJECT environment variable")
	}
	test := getJsOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "gcp-container-registry/ts"),
			Config: map[string]string{
				"gcp:project": project,
			},
		})

	integration.ProgramTest(t, &test)
}

func TestDockerContainerRegistryNode(t *testing.T) {
	username := "pulumibot"
	password := os.Getenv("DOCKER_HUB_PASSWORD")
	test := getJsOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "docker-container-registry/ts"),
			Config: map[string]string{
				"cbp-docker-ts-dev:dockerUsername": username,
			},
			Secrets: map[string]string{
				"cbp-docker-ts-dev:dockerPassword": password,
			},
			ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
				digest, ok := stack.Outputs["repoDigest"].(string)
				assert.True(t, ok)
				assert.NotEmpty(t, digest)
			},
		})
	integration.ProgramTest(t, &test)
}

func TestUnknownInputsNode(t *testing.T) {
	test := getJsOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "test-unknowns", "ts"),
			ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
				randID, ok := stack.Outputs["randnameid"]
				assert.True(t, ok)
				assert.NotEmpty(t, randID)
			},
			Quick:       true,
			SkipRefresh: true,
		})
	integration.ProgramTest(t, &test)
}

func TestSecretsInExplicitProviderNode(t *testing.T) {
	check := func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
		deploymentJSON, err := json.MarshalIndent(stack.Deployment, "", "  ")
		assert.NoError(t, err)

		t.Run("providerWithSecretAddress", func(t *testing.T) {
			assert.NotContainsf(t, string(deploymentJSON), "secret-address",
				"Secrets like 'secret-address' should not be stored in the plain")
		})

		t.Run("providerWithSecretUsername", func(t *testing.T) {
			pw := stack.Outputs["password"].(string)
			realPW, err := base64.StdEncoding.DecodeString(pw)
			assert.NoError(t, err)
			assert.NotContainsf(t, string(deploymentJSON), string(realPW),
				"Secrets like the output of RandomPassword should not be stored in the plain")
		})

		t.Run("providerWithSecretPassword", func(t *testing.T) {
			assert.NotContainsf(t, string(deploymentJSON), "secret-password",
				"Secret properties like RegistryAuth.Password should not be stored in the plain")
		})

		t.Run("noPanics", func(t *testing.T) {
			// Temporary check to rule out panics; needed until pulumi/pulumi#12981 is resolved.
			for _, e := range stack.Events {
				eventsJSON, err := json.MarshalIndent(e, "", "  ")
				assert.NoError(t, err)
				assert.NotContainsf(t, string(eventsJSON), "panic",
					"Unexpected panic recorded in engine events")
			}
		})
	}
	test := getJsOptions(t).With(integration.ProgramTestOptions{
		Dir:                    path.Join(getCwd(t), "test-secrets-in-explicit-provider", "ts"),
		Quick:                  true,
		SkipRefresh:            true,
		ExtraRuntimeValidation: check,
	})
	integration.ProgramTest(t, &test)
}

func TestSSHConnNode(t *testing.T) {
	token := os.Getenv("DIGITALOCEAN_TOKEN")
	if token == "" {
		t.Skipf("Skipping test due to missing DIGITALOCEAN_TOKEN environment variable")
	}
	test := getJsOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "test-ssh-conn", "ts"),
			ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
				ipOutput, ok := stack.Outputs["ipOutput"]
				assert.True(t, ok)
				assert.NotEmpty(t, ipOutput)
				// Due to a bug in the docker client, we add sleep here so our SSH connection does not get refused.
				time.Sleep(20 * time.Second)
			},
			SkipRefresh:      true,
			Quick:            true,
			RetryFailedSteps: true,
			Config: map[string]string{
				"digitalocean:token": token,
			},
			Verbose: true,
		})
	integration.ProgramTest(t, &test)
}

func getJsOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions()
	baseJs := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			"@pulumi/docker",
		},
	})

	return baseJs
}
