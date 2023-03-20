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
			"github.com/pulumi/pulumi-docker/sdk/v4=../sdk",
		},
		Dir: path.Join(cwd, "nginx-go"),
	})
	integration.ProgramTest(t, &opts)
}

// TestBuildCacheFromGo tests the use of CacheFrom behavior with Image v4 using BuildKit.
//
// This runs `pulumi` multiple times via ProgramTest to evaluate docker build times.
//
// Step 1 creates an ECR registry and a docker image with a 30 second sleep in the first build
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
// used and the each stage uses a small image as the base.
func TestBuildCacheFromGo(t *testing.T) {
	cwd, err := os.Getwd()
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	clearCache := func(imageName string) error {
		if imageName != "" {
			if err := exec.Command("docker", "image", "rm", imageName).Run(); err != nil {
				return fmt.Errorf("error removing image %q: %w", imageName, err)
			}
		}
		if err := exec.Command("docker", "image", "prune").Run(); err != nil {
			return fmt.Errorf("error pruning dangling images: %w", err)
		}
		if err := exec.Command("docker", "builder", "prune", "-a").Run(); err != nil {
			return fmt.Errorf("error pruning build cache: %w", err)
		}
		return nil
	}

	clearCache("")

	step1Start := time.Now()
	var step1DeployTime time.Duration
	var step2Start time.Time
	var step2DeployTime time.Duration

	var imageName string
	var ok bool

	opts := base.With(integration.ProgramTestOptions{
		Dir:              path.Join(cwd, "multi-stage-build-go"),
		Quick:            true,
		SkipRefresh:      true,
		DestroyOnCleanup: true,
		Dependencies: []string{
			"github.com/pulumi/pulumi-docker/sdk/v4=../sdk",
		},
		ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
			step1DeployTime = time.Since(step1Start)

			if imageName, ok = stack.Outputs["imageName"].(string); ok {
				assert.NoError(t, clearCache(imageName))
			} else {
				t.Errorf("expected repositoryUrl output")
			}

			step2Start = time.Now()
		},
		EditDirs: []integration.EditDir{
			{
				Dir:      path.Join(cwd, "multi-stage-build-go", "step2"),
				Additive: true,
				ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
					step2DeployTime = time.Since(step2Start)
					clearCache(imageName)
				},
			},
		},
	})
	integration.ProgramTest(t, &opts)

	assert.Less(t, step2DeployTime, 30*time.Second)
	assert.Greater(t, step1DeployTime, 30*time.Second)
}

func TestDockerfileGo(t *testing.T) {
	cwd, err := os.Getwd()
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	opts := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			"github.com/pulumi/pulumi-docker/sdk/v4=../sdk",
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
		Dir: path.Join(cwd, "azure-container-registry/go"),
		Config: map[string]string{
			"azure:environment": "public",
			"azure:location":    location,
		},
		Dependencies: []string{
			"github.com/pulumi/pulumi-docker/sdk/v4=../sdk",
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
		Dir: path.Join(cwd, "aws-container-registry/go"),
		Config: map[string]string{
			"aws:region": region,
		},
		Dependencies: []string{
			"github.com/pulumi/pulumi-docker/sdk/v4=../sdk",
		},
	})
	integration.ProgramTest(t, &opts)
}

func TestDigitaloceanContainerRegistryGo(t *testing.T) {
	token := os.Getenv("DIGITALOCEAN_TOKEN")
	if token == "" {
		t.Skipf("Skipping test due to missing DIGITALOCEAN_TOKEN environment variable")
	}
	cwd, err := os.Getwd()
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	opts := base.With(integration.ProgramTestOptions{
		Dir: path.Join(cwd, "digitalocean-container-registry/go"),
		Config: map[string]string{
			"digitalocean:token": token,
		},
		Dependencies: []string{
			"github.com/pulumi/pulumi-docker/sdk/v4=../sdk",
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
		Dir: path.Join(getCwd(t), "gcp-container-registry/go"),
		Config: map[string]string{
			"gcp:project": project,
		},
		Dependencies: []string{
			"github.com/pulumi/pulumi-docker/sdk/v4=../sdk",
		},
	})
	integration.ProgramTest(t, &test)
}

func TestDockerContainerRegistryGo(t *testing.T) {
	username := "pulumibot"
	password := os.Getenv("DOCKER_HUB_PASSWORD")
	test := base.With(integration.ProgramTestOptions{
		Dir: path.Join(getCwd(t), "docker-container-registry/go"),
		Config: map[string]string{
			"cbp-docker-go:dockerUsername": username,
		},
		Secrets: map[string]string{
			"cbp-docker-go:dockerPassword": password,
		},
		Dependencies: []string{
			"github.com/pulumi/pulumi-docker/sdk/v4=../sdk",
		},
	})
	integration.ProgramTest(t, &test)
}

var base = integration.ProgramTestOptions{
	ExpectRefreshChanges: true, // Docker resources generally see changes when refreshed.
	// Note: no Config! This package should be usable without any config.
}
