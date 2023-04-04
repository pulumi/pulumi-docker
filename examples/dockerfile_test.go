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
	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"github.com/stretchr/testify/assert"
	"os"
	"path"
	"testing"
)

func TestDockerfileDefaultYAML(t *testing.T) {
	cwd, err := os.Getwd()
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	integration.ProgramTest(t, &integration.ProgramTestOptions{
		Dir:         path.Join(cwd, "test-dockerfile", "dockerfile-default"),
		Quick:       true,
		SkipRefresh: true,
	})
}

func TestDockerfileDefaultContextYAML(t *testing.T) {
	cwd, err := os.Getwd()
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	integration.ProgramTest(t, &integration.ProgramTestOptions{
		Dir:         path.Join(cwd, "test-dockerfile", "dockerfile-default-context"),
		Quick:       true,
		SkipRefresh: true,
	})
}

func TestDockerfileExternalYAML(t *testing.T) {
	cwd, err := os.Getwd()
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	integration.ProgramTest(t, &integration.ProgramTestOptions{
		Dir:         path.Join(cwd, "test-dockerfile", "dockerfile-external"),
		Quick:       true,
		SkipRefresh: true,
	})
}

func TestDockerfileInContextYAML(t *testing.T) {
	cwd, err := os.Getwd()
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	integration.ProgramTest(t, &integration.ProgramTestOptions{
		Dir:         path.Join(cwd, "test-dockerfile", "dockerfile-in-context"),
		Quick:       true,
		SkipRefresh: true,
	})
}

func TestDockerignoreDefaultYAML(t *testing.T) {
	cwd, err := os.Getwd()
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	integration.ProgramTest(t, &integration.ProgramTestOptions{
		Dir:         path.Join(cwd, "test-dockerfile", "dockerignore-default"),
		Quick:       true,
		SkipRefresh: true,
	})
}

func TestDockerignoreSpecifiedYAML(t *testing.T) {
	cwd, err := os.Getwd()
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	integration.ProgramTest(t, &integration.ProgramTestOptions{
		Dir:         path.Join(cwd, "test-dockerfile", "dockerignore-specified"),
		Quick:       true,
		SkipRefresh: true,
	})
}

func TestDockerignoreDefaultFailYAML(t *testing.T) {
	cwd, err := os.Getwd()
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	integration.ProgramTest(t, &integration.ProgramTestOptions{
		Dir:           path.Join(cwd, "test-dockerfile", "dockerignore-default-fail"),
		Quick:         true,
		SkipRefresh:   true,
		ExpectFailure: true,
	})
}

func TestDockerignoreNoMappingYAML(t *testing.T) {
	cwd, err := os.Getwd()
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	// we expect this test to succeed, as we test that the ignore.txt file does in fact _not_ get ignored
	// the ignore.txt file does not get ignored, as  .dockerignore does not map to Mockerfile.
	// The RUN command in Mockerfile therefore succeeds.
	integration.ProgramTest(t, &integration.ProgramTestOptions{
		Dir:         path.Join(cwd, "test-dockerfile", "dockerignore-no-mapping"),
		Quick:       true,
		SkipRefresh: true,
	})
}

func TestDockerignoreWithExternalDockerfileYAML(t *testing.T) {
	cwd, err := os.Getwd()
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	integration.ProgramTest(t, &integration.ProgramTestOptions{
		Dir:         path.Join(cwd, "test-dockerfile", "dockerignore-with-external-dockerfile"),
		Quick:       true,
		SkipRefresh: true,
	})
}
