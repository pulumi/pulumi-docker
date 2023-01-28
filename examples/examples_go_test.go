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
	"fmt"
	"os"
	"os/exec"
	"path"
	"testing"
	"time"

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

// TestBuildCacheFromGo tests the use of CacheFrom behavior with Image v4 using BuildKit.
//
// This runs `pulumi` multiple times via ProgramTest to evaluate docker build times.
//
// Step 1 creates an ECR registry and a docker image with a 10 second sleep in the first build
// stage.
//
// ExtraRuntimeValidation after step 1 clears deletes the image locally, then prunes docker's build
// cache.
//
// Step 2 modifies a second-stage input and rebuilds.
//
// We expect Step 1 to take more time than the sleep duration, ensuring we don't have a local cache.
//
// With `--cache-from` and a decent network connection, we expect Step 2 to take less time than the
// sleep duration.
//
// There is some measurement error here, as we're measuring Docker via `pulumi up` and step 2 will
// have to download the first stage from the repository. To minimize this the same Pulumi program is
// used and the initial stage uses busybox as a base.
func TestBuildCacheFromGo(t *testing.T) {
	cwd, err := os.Getwd()
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	clearCache := func(imageName string) error {
		if err := exec.Command("docker", "image", "rm", imageName).Run(); err != nil {
			return fmt.Errorf("error removing image %q: %w", imageName, err)
		}
		if err := exec.Command("docker", "image", "prune").Run(); err != nil {
			return fmt.Errorf("error pruning dangling images: %w", err)
		}
		if err := exec.Command("docker", "builder", "prune").Run(); err != nil {
			return fmt.Errorf("error pruning build cache: %w", err)
		}
		return nil
	}

	step1Start := time.Now()
	var step1DeployTime time.Duration
	var step2Start time.Time
	var step2DeployTime time.Duration

	opts := base.With(integration.ProgramTestOptions{
		Dir:         path.Join(cwd, "build-cache-from-go"),
		Quick:       true,
		SkipRefresh: true,
		Dependencies: []string{
			"github.com/pulumi/pulumi-docker/sdk=../sdk",
		},
		ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
			step1DeployTime = time.Since(step1Start)
			if repoUrl, ok := stack.Outputs["repositoryUrl"].(string); ok {
				assert.NoError(t, clearCache(repoUrl))
			} else {
				t.Errorf("expected repositoryUrl output")
			}

			step2Start = time.Now()
		},
		EditDirs: []integration.EditDir{
			{
				Dir:      "step2",
				Additive: true,
				ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
					step2DeployTime = time.Since(step2Start)
					if repoUrl, ok := stack.Outputs["repositoryUrl"].(string); ok {
						assert.NoError(t, clearCache(repoUrl))
					} else {
						t.Errorf("expected repositoryUrl output")
					}
				},
			},
		},
	})
	integration.ProgramTest(t, &opts)

	assert.Greater(t, step1DeployTime, 30*time.Second)
	assert.Less(t, step2DeployTime, 30*time.Second)
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
