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
//go:build go || all
// +build go all

package examples

import (
	"os"
	"path"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"github.com/stretchr/testify/assert"
)

func TestNginxGo(t *testing.T) {

	cwd, err := os.Getwd()
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	opts := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			"github.com/pulumi/pulumi-docker/sdk=../sdk",
		},
		Dir: path.Join(cwd, "nginx-go"),
	})
	integration.ProgramTest(t, &opts)
}

func TestBuildCacheFromGo(t *testing.T) {
	t.Skip("ignoring due to major version change")
	cwd, err := os.Getwd()
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	opts := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			"github.com/pulumi/pulumi-docker/sdk=../sdk",
			"github.com/pulumi/pulumi-aws/sdk/v5",
		},
		Dir: path.Join(cwd, "build-cache-from-go"),
	})
	integration.ProgramTest(t, &opts)
}

func TestDockerfileGo(t *testing.T) {
	cwd, err := os.Getwd()
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	opts := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			"github.com/pulumi/pulumi-docker/sdk=../sdk",
		},
		Dir: path.Join(cwd, "dockerfile-go"),
	})
	integration.ProgramTest(t, &opts)
}

func TestAzureContainerRegistryGo(t *testing.T) {
	location := os.Getenv("AZURE_LOCATION")
	if location == "" {
		t.Skipf("Skipping test due to missing AZURE_LOCATION environment variable")
	}
	cwd, err := os.Getwd()
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	opts := base.With(integration.ProgramTestOptions{
		Dir: path.Join(cwd, "container-registries/azure/go"),
		Config: map[string]string{
			"azure:environment": "public",
			"azure:location":    location,
		},
	})
	integration.ProgramTest(t, &opts)
}

func TestAwsContainerRegistryGo(t *testing.T) {
	region := os.Getenv("AWS_REGION")
	if region == "" {
		t.Skipf("Skipping test due to missing AWS_REGION environment variable")
	}
	cwd, err := os.Getwd()
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	opts := base.With(integration.ProgramTestOptions{
		Dir: path.Join(cwd, "container-registries/aws/go"),
		Config: map[string]string{
			"aws:region": region,
		},
	})
	integration.ProgramTest(t, &opts)
}

func TestGcpContainerRegistryGo(t *testing.T) {
	project := os.Getenv("GOOGLE_PROJECT")
	if project == "" {
		t.Skipf("Skipping test due to missing GOOGLE_PROJECT environment variable")
	}
	test := base.With(integration.ProgramTestOptions{
		Dir: path.Join(getCwd(t), "container-registries/gcp/go"),
		Config: map[string]string{
			"gcp:project": project,
		},
	})
	integration.ProgramTest(t, &test)
}

var base = integration.ProgramTestOptions{
	ExpectRefreshChanges: true, // Docker resources generally see changes when refreshed.
	// Note: no Config! This package should be usable without any config.
}
