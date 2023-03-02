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
//go:build dotnet || all
// +build dotnet all

package examples

import (
	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"os"
	"path"
	"testing"
)

func TestNginxCs(t *testing.T) {
	test := getCsharpBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "nginx-cs"),
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
			Dir: path.Join(getCwd(t), "container-registries/azure/csharp"),
			Config: map[string]string{
				"azure:environment": "public",
				"azure:location":    location,
			},
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
			Dir: path.Join(getCwd(t), "container-registries/aws/csharp"),
			Config: map[string]string{
				"aws:region": region,
			},
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
			Dir: path.Join(getCwd(t), "container-registries/gcp/csharp"),
			Config: map[string]string{
				"gcp:project": project,
			},
		})
	integration.ProgramTest(t, &test)
}

func TestDockerContainerRegistryDotnet(t *testing.T) {
	t.Skipf("smoke screen checck")
	username := "pulumibot"
	password := os.Getenv("DOCKER_HUB_PASSWORD")
	test := getCsharpBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "container-registries/docker/csharp"),
			Config: map[string]string{
				"cbp-docker-csharp:dockerUsername": username,
			},
			Secrets: map[string]string{
				"cbp-docker-csharp:dockerPassword": password,
			},
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
