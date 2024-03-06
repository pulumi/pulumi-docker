// Copyright 2024, Pulumi Corporation.
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
//go:build java || all
// +build java all

package examples

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/engine"
	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func TestBuildxJava(t *testing.T) {
	t.Skip("not working")

	test := getJavaBase(t, path.Join("buildx", "java"), integration.ProgramTestOptions{
		Secrets: map[string]string{
			"dockerHubPassword": os.Getenv("DOCKER_HUB_PASSWORD"),
		},
	})

	integration.ProgramTest(t, &test)
}

func getJavaBase(t *testing.T, dir string, testSpecificOptions integration.ProgramTestOptions) integration.ProgramTestOptions {
	repoRoot, err := filepath.Abs(filepath.Join("..", ".."))
	if err != nil {
		panic(err)
	}
	opts := integration.ProgramTestOptions{
		Dir: filepath.Join(getCwd(t), dir),
		Env: []string{fmt.Sprintf("PULUMI_REPO_ROOT=%s", repoRoot)},
		PrepareProject: func(*engine.Projinfo) error {
			return nil // needed because defaultPrepareProject does not know about java
		},
	}
	opts = opts.With(getBaseOptions()).With(testSpecificOptions)
	return opts
}
