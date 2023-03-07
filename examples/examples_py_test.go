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
//go:build python || all
// +build python all

package examples

import (
	"os"
	"path"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func TestAzureContainerRegistryPy(t *testing.T) {
	location := os.Getenv("AZURE_LOCATION")
	if location == "" {
		t.Skipf("Skipping test due to missing AZURE_LOCATION environment variable")
	}
	test := getPyOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "container-registries/azure/py"),
			Config: map[string]string{
				"azure:environment": "public",
				"azure:location":    location,
			},
		})

	integration.ProgramTest(t, &test)
}

func TestAwsContainerRegistryPy(t *testing.T) {
	region := os.Getenv("AWS_REGION")
	if region == "" {
		t.Skipf("Skipping test due to missing AWS_REGION environment variable")
	}
	test := getPyOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "container-registries/aws/py"),
			Config: map[string]string{
				"aws:region": region,
			},
		})

	integration.ProgramTest(t, &test)
}

func TestDigitaloceanContainerRegistryPy(t *testing.T) {
	token := os.Getenv("DIGITALOCEAN_TOKEN")
	if token == "" {
		t.Skipf("Skipping test due to missing DIGITALOCEAN_TOKEN environment variable")
	}
	test := getPyOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "container-registries/digitalocean/py"),
			Config: map[string]string{
				"digitalocean:token": token,
			},
		})

	integration.ProgramTest(t, &test)
}

func TestGcpContainerRegistryPy(t *testing.T) {
	project := os.Getenv("GOOGLE_PROJECT")
	if project == "" {
		t.Skipf("Skipping test due to missing GOOGLE_PROJECT environment variable")
	}
	test := getPyOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "container-registries/gcp/py"),
			Config: map[string]string{
				"gcp:project": project,
			},
		})

	integration.ProgramTest(t, &test)
}

func TestDockerContainerRegistryPy(t *testing.T) {
	username := "pulumibot"
	password := os.Getenv("DOCKER_HUB_PASSWORD")
	test := getPyOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "container-registries/docker/py"),
			Config: map[string]string{
				"cbp-docker-py:dockerUsername": username,
			},
			Secrets: map[string]string{
				"cbp-docker-py:dockerPassword": password,
			},
		})
	integration.ProgramTest(t, &test)
}
func TestNginxPy(t *testing.T) {
	test := getPyOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "nginx-py"),
		})

	integration.ProgramTest(t, &test)
}

func TestDockerfilePy(t *testing.T) {
	test := getPyOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "dockerfile-py"),
		})

	integration.ProgramTest(t, &test)
}

func getPyOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions()
	basePy := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			path.Join("..", "sdk", "python", "bin"),
		},
	})

	return basePy
}
