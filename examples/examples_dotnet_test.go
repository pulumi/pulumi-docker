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
// +build dotnet all

package examples

import (
	"os"
	"path"
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
	"github.com/stretchr/testify/assert"
)

func TestNginxCs(t *testing.T) {
	cwd, err := os.Getwd()
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	opts := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			"Pulumi.Docker",
		},
		Dir: path.Join(cwd, "nginx-cs"),
	})
	integration.ProgramTest(t, &opts)
}

func TestDotNet(t *testing.T) {
	cwd, err := os.Getwd()
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	opts := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			"Pulumi.Docker",
		},
		Dir:                    path.Join(cwd, "dotnet"),
		ExtraRuntimeValidation: dockerFileWithDependenciesOutputValidation,
	})
	integration.ProgramTest(t, &opts)
}

func TestAwsContainerRegistryCSharp(t *testing.T) {
	region := os.Getenv("AWS_REGION")
	if region == "" {
		t.Skipf("Skipping test due to missing AWS_REGION environment variable")
	}
	fmt.Printf("AWS Region: %v\n", region)

	cwd, err := os.Getwd()
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	opts := base.With(integration.ProgramTestOptions{
		Config: map[string]string{
			"aws:region": region,
		},
		Dependencies: []string{
			"Pulumi.Docker",
		},
		Dir: path.Join(cwd, "container-registries", "aws", "csharp"),
	})
	integration.ProgramTest(t, &opts)
}
