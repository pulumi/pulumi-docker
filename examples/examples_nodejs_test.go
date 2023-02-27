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
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func TestAws(t *testing.T) {
	t.Skip("Skipping test due to updates in Image resource")
	region := os.Getenv("AWS_REGION")
	if region == "" {
		t.Skipf("Skipping test due to missing AWS_REGION environment variable")
	}
	fmt.Printf("AWS Region: %v\n", region)

	test := getJsOptions(t).
		With(integration.ProgramTestOptions{
			Config: map[string]string{
				"aws:region": region,
			},
			Dir: path.Join(getCwd(t), "aws"),
		})

	integration.ProgramTest(t, &test)
}

func TestNginx(t *testing.T) {
	test := getJsOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "nginx"),
		})

	integration.ProgramTest(t, &test)
}

func TestDockerfileWithMultipleTargets(t *testing.T) {

	test := getJsOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "dockerfile-with-targets"),
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
			Dir: path.Join(getCwd(t), "container-registries/azure/ts"),
			Config: map[string]string{
				"azure:environment": "public",
				"azure:location":    location,
			},
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
			Dir: path.Join(getCwd(t), "container-registries/aws/ts"),
			Config: map[string]string{
				"aws:region": region,
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
			Dir: path.Join(getCwd(t), "container-registries/gcp/ts"),
			Config: map[string]string{
				"gcp:project": project,
			},
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
