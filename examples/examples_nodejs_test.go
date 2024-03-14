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
	"fmt"
	"math/rand"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"github.com/stretchr/testify/assert"
)

func TestBuildxTs(t *testing.T) {
	test := getJsOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "buildx", "ts"),
			Secrets: map[string]string{
				"dockerHubPassword": os.Getenv("DOCKER_HUB_PASSWORD"),
			},
		})

	integration.ProgramTest(t, &test)
}

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
	dir := path.Join(getCwd(t), "azure-container-registry/ts")
	test := getJsOptions(t).
		With(integration.ProgramTestOptions{
			Dir: dir,
			Config: map[string]string{
				"azure:environment": "public",
				"azure:location":    location,
			},
			ExpectRefreshChanges: true,
			EditDirs: []integration.EditDir{
				{
					Dir:      path.Join(dir, "step2"),
					Additive: true,
				},
			},
			ExtraRuntimeValidation: assertHasRepoDigest,
		})

	integration.ProgramTest(t, &test)
}

func TestAwsContainerRegistryNode(t *testing.T) {
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
			ExtraRuntimeValidation: assertHasRepoDigest,
		})

	integration.ProgramTest(t, &test)
}

func TestDigitaloceanContainerRegistry(t *testing.T) {
	t.Skipf("Skipping test due to known storageUsageBytes issue https://github.com/pulumi/pulumi-docker/issues/718")

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
			ExtraRuntimeValidation: assertHasRepoDigest,
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
			ExtraRuntimeValidation: assertHasRepoDigest,
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
			ExtraRuntimeValidation: assertHasRepoDigest,
		})
	integration.ProgramTest(t, &test)
}

func TestUnknownInputsNode(t *testing.T) {
	test := getJsOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "test-unknowns", "ts"),
			ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
				randID, ok := stack.Outputs["randArgument"]
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

func TestLocalRepoDigestNode(t *testing.T) {
	test := getJsOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "test-local-repo-digest-ts"),
			ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
				repoDigest, ok := stack.Outputs["repoDigest"]
				assert.True(t, ok)
				assert.NotEmpty(t, repoDigest)
			},
			Quick:       true,
			SkipRefresh: true,
		})
	integration.ProgramTest(t, &test)
}

// TestBuildxCaching simulates a slow multi-platform build with --cache-to
// enabled. We aren't able to directly detect cache hits, so we re-run the
// update and confirm it took less time than the image originally took to
// build.
//
// This is a moderately slow test because we need to "build" (i.e., sleep)
// longer than it would take for cache layer uploads under slow network
// conditions.
func TestBuildxCaching(t *testing.T) {
	t.Parallel()

	sleep := 20.0 // seconds

	// Provision ECR outside of our stack, because the cache needs to be shared
	// across updates.
	ecr, ecrOK := tmpEcr(t)

	localCache := t.TempDir()

	tests := []struct {
		name string
		skip bool

		cacheTo   string
		cacheFrom string
		address   string
		username  string
		password  string
	}{
		{
			name:      "local",
			cacheTo:   fmt.Sprintf("type=local,mode=max,oci-mediatypes=true,dest=%s", localCache),
			cacheFrom: fmt.Sprintf("type=local,src=%s", localCache),
		},
		{
			name:      "gha",
			skip:      os.Getenv("ACTIONS_CACHE_URL") == "",
			cacheTo:   "type=gha,mode=max,scope=cache-test",
			cacheFrom: "type=gha,scope=cache-test",
		},
		{
			name:      "dockerhub",
			skip:      os.Getenv("DOCKER_HUB_PASSWORD") == "",
			cacheTo:   "type=registry,mode=max,ref=docker.io/pulumibot/myapp:cache",
			cacheFrom: "type=registry,ref=docker.io/pulumibot/myapp:cache",
			address:   "docker.io",
			username:  "pulumibot",
			password:  os.Getenv("DOCKER_HUB_PASSWORD"),
		},
		{
			name:      "ecr",
			skip:      !ecrOK,
			cacheTo:   fmt.Sprintf("type=registry,mode=max,image-manifest=true,oci-mediatypes=true,ref=%s:cache", ecr.address),
			cacheFrom: fmt.Sprintf("type=registry,ref=%s:cache", ecr.address),
			address:   ecr.address,
			username:  ecr.username,
			password:  ecr.password,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if tt.skip {
				t.Skip("Missing environment variables")
			}

			sleepFuzzed := sleep + rand.Float64() // Add some fuzz to bust any prior build caches.

			test := getJsOptions(t).
				With(integration.ProgramTestOptions{
					Dir: path.Join(getCwd(t), "test-buildx", "caching", "ts"),
					ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
						duration, ok := stack.Outputs["durationSeconds"]
						assert.True(t, ok)
						assert.Greater(t, duration.(float64), sleepFuzzed)
					},
					Config: map[string]string{
						"SLEEP_SECONDS": fmt.Sprint(sleepFuzzed),
						"cacheTo":       tt.cacheTo,
						"cacheFrom":     tt.cacheFrom,
						"name":          tt.name,
						"address":       tt.address,
						"username":      tt.username,
					},
					Secrets: map[string]string{
						"password": tt.password,
					},
					NoParallel:  true,
					Quick:       true,
					SkipPreview: true,
					SkipRefresh: true,
					Verbose:     true,
				})

			// First run should be un-cached.
			integration.ProgramTest(t, &test)

			// Now run again and confirm our build was faster due to a cache hit.
			test.ExtraRuntimeValidation = func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
				duration, ok := stack.Outputs["durationSeconds"]
				assert.True(t, ok)
				assert.Less(t, duration.(float64), sleepFuzzed)
			}
			test.Config["name"] += "-cached"
			integration.ProgramTest(t, &test)
		})
	}
}

func TestBuildxIndex(t *testing.T) {
	password := os.Getenv("DOCKER_HUB_PASSWORD")
	if password == "" {
		t.Skip("missing DOCKER_HUB_PASSWORD")
	}

	test := getJsOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "test-buildx", "index", "ts"),
			Config: map[string]string{
				"tag":      "docker.io/pulumibot/buildkit-e2e:manifest",
				"address":  "docker.io",
				"username": "pulumibot",
			},
			Secrets: map[string]string{
				"password": password,
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

type ECR struct {
	address  string
	username string
	password string
}

// tmpEcr creates a new ECR repo and cleans it up after the test concludes.
func tmpEcr(t *testing.T) (ECR, bool) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	})
	if err != nil {
		return ECR{}, false
	}

	svc := ecr.New(sess)
	name := strings.ToLower(t.Name())

	// Always attempt to delete pre-existing repos, in case our cleanup didn't
	// run.
	_, _ = svc.DeleteRepository(&ecr.DeleteRepositoryInput{
		Force:          aws.Bool(true),
		RepositoryName: aws.String(name),
	})

	params := &ecr.CreateRepositoryInput{
		RepositoryName: aws.String(name),
	}
	resp, err := svc.CreateRepository(params)
	if err != nil {
		return ECR{}, false
	}
	repo := resp.Repository
	t.Cleanup(func() {
		svc.DeleteRepository(&ecr.DeleteRepositoryInput{
			Force:          aws.Bool(true),
			RegistryId:     repo.RegistryId,
			RepositoryName: repo.RepositoryName,
		})
	})

	// Now grab auth for the repo.
	auth, err := svc.GetAuthorizationToken(&ecr.GetAuthorizationTokenInput{})
	if err != nil {
		return ECR{}, false
	}
	b64token := auth.AuthorizationData[0].AuthorizationToken
	token, err := base64.StdEncoding.DecodeString(*b64token)
	if err != nil {
		return ECR{}, false
	}
	parts := strings.SplitN(string(token), ":", 2)
	if len(parts) != 2 {
		return ECR{}, false
	}
	username := parts[0]
	password := parts[1]

	return ECR{
		address:  *repo.RepositoryUri,
		username: username,
		password: password,
	}, true
}
