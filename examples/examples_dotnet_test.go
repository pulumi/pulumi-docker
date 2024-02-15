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
//go:build dotnet || all
// +build dotnet all

package examples

import (
	"encoding/base64"
	"encoding/json"
	"os"
	"path"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"github.com/stretchr/testify/assert"
)

func TestBuildxCs(t *testing.T) {
	test := getCsharpBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "buildx", "csharp"),
			Secrets: map[string]string{
				"dockerHubPassword": os.Getenv("DOCKER_HUB_PASSWORD"),
			},
		})

	integration.ProgramTest(t, &test)
}

func TestNginxCs(t *testing.T) {
	test := getCsharpBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir:                  path.Join(getCwd(t), "nginx-cs"),
			ExpectRefreshChanges: true,
			Quick:                true,
		})

	integration.ProgramTest(t, &test)
}

func TestDotNet(t *testing.T) {
	test := getCsharpBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "dotnet"),
		})

	integration.ProgramTest(t, &test)
}

func TestAzureContainerRegistryDotNet(t *testing.T) {
	location := os.Getenv("AZURE_LOCATION")
	if location == "" {
		t.Skipf("Skipping test due to missing AZURE_LOCATION environment variable")
	}
	test := getCsharpBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "azure-container-registry/csharp"),
			Config: map[string]string{
				"azure:environment": "public",
				"azure:location":    location,
			},
			ExpectRefreshChanges:   true,
			ExtraRuntimeValidation: assertHasRepoDigest,
		})

	integration.ProgramTest(t, &test)
}

func TestAwsContainerRegistryDotnet(t *testing.T) {
	region := os.Getenv("AWS_REGION")
	if region == "" {
		t.Skipf("Skipping test due to missing AWS_REGION environment variable")
	}
	test := getCsharpBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "aws-container-registry/csharp"),
			Config: map[string]string{
				"aws:region": region,
			},
			ExtraRuntimeValidation: assertHasRepoDigest,
		})

	integration.ProgramTest(t, &test)
}

func TestDigitaloceanContainerRegistryDotnet(t *testing.T) {
	t.Skipf("Skipping test due to known storageUsageBytes issue https://github.com/pulumi/pulumi-docker/issues/718")

	token := os.Getenv("DIGITALOCEAN_TOKEN")
	if token == "" {
		t.Skipf("Skipping test due to missing DIGITALOCEAN_TOKEN environment variable")
	}
	test := getCsharpBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "digitalocean-container-registry/csharp"),
			Config: map[string]string{
				"digitalocean:token": token,
			},
			ExtraRuntimeValidation: assertHasRepoDigest,
		})

	integration.ProgramTest(t, &test)
}

func TestGcpContainerRegistryDotnet(t *testing.T) {
	project := os.Getenv("GOOGLE_PROJECT")
	if project == "" {
		t.Skipf("Skipping test due to missing GOOGLE_PROJECT environment variable")
	}
	test := getCsharpBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "gcp-container-registry/csharp"),
			Config: map[string]string{
				"gcp:project": project,
			},
			ExtraRuntimeValidation: assertHasRepoDigest,
		})
	integration.ProgramTest(t, &test)
}

func TestDockerContainerRegistryDotnet(t *testing.T) {
	username := "pulumibot"
	password := os.Getenv("DOCKER_HUB_PASSWORD")
	test := getCsharpBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "docker-container-registry/csharp"),
			Config: map[string]string{
				"cbp-docker-csharp:dockerUsername": username,
			},
			Secrets: map[string]string{
				"cbp-docker-csharp:dockerPassword": password,
			},
			ExtraRuntimeValidation: assertHasRepoDigest,
		})
	integration.ProgramTest(t, &test)
}

func TestSecretsInExplicitProvider(t *testing.T) {
	check := func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
		deploymentJSON, err := json.MarshalIndent(stack.Deployment, "", "  ")
		assert.NoError(t, err)

		t.Run("providerWithSecretAddress", func(t *testing.T) {
			t.Skip("TODO[pulumi/pulumi-docker#643]")
			assert.NotContainsf(t, string(deploymentJSON), "secret-address",
				"Secrets like 'secret-address' should not be stored in the plain")
		})

		t.Run("providerWithSecretUsername", func(t *testing.T) {
			t.Skip("TODO[pulumi/pulumi-docker#643]")
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
	test := getCsharpBaseOptions(t).With(integration.ProgramTestOptions{
		Dir:                    path.Join(getCwd(t), "test-secrets-in-explicit-provider", "csharp"),
		Quick:                  true,
		SkipRefresh:            true,
		ExtraRuntimeValidation: check,
	})
	integration.ProgramTest(t, &test)
}

func getCsharpBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions()
	baseCsharp := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			"Pulumi.Docker",
		},
	})

	return baseCsharp
}
