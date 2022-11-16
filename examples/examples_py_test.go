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
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func TestAwsPy(t *testing.T) {
	t.Skipf("https://github.com/pulumi/pulumi-docker/issues/411")
	region := os.Getenv("AWS_REGION")
	if region == "" {
		t.Skipf("Skipping test due to missing AWS_REGION environment variable")
	}
	fmt.Printf("AWS Region: %v\n", region)

	test := getPyOptions(t).
		With(integration.ProgramTestOptions{
			Config: map[string]string{
				"aws:region": region,
			},
			Dir: path.Join(getCwd(t), "aws-py"),
		})

	integration.ProgramTest(t, &test)
}

func TestAzurePy(t *testing.T) {
	t.Skip("Skipping test due to updates in Image resource")
	location := os.Getenv("AZURE_LOCATION")
	if location == "" {
		t.Skipf("Skipping test due to missing AZURE_LOCATION environment variable")
	}
	test := getPyOptions(t).
		With(integration.ProgramTestOptions{
			Config: map[string]string{
				"azure:environment": "public",
				"azure:location":    location,
			},
			Dir: path.Join(getCwd(t), "azure-py"),
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
