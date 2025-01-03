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

package examples

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/pulumi/providertest/pulumitest"
	"github.com/pulumi/providertest/pulumitest/changesummary"
	"github.com/pulumi/providertest/pulumitest/opttest"
	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/pulumi/pulumi/sdk/v3/go/common/apitype"
	"github.com/stretchr/testify/assert"
)

func getCwd(t *testing.T) string {
	cwd, err := os.Getwd()
	if err != nil {
		t.FailNow()
	}

	return cwd
}

func getBaseOptions() integration.ProgramTestOptions {
	return integration.ProgramTestOptions{
		RunUpdateTest: false,
	}
}

func assertHasRepoDigest(t *testing.T, stack integration.RuntimeValidationStackInfo) {
	repoDigest, ok := stack.Outputs["repoDigest"].(string)
	assert.True(t, ok, "expected repoDigest output")
	assert.NotEmpty(t, repoDigest)
}

func pulumiTest(t *testing.T, dir string, opts ...opttest.Option) *pulumitest.PulumiTest {
	cwd, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}
	opts = append(opts, opttest.LocalProviderPath("docker", filepath.Join(cwd, "..", "bin")))
	ptest := pulumitest.NewPulumiTest(t, dir, opts...)
	return ptest
}

func AssertHasChanges(t *testing.T, preview auto.PreviewResult) {
	t.Helper()

	convertedMap := changesummary.ChangeSummary(preview.ChangeSummary)
	expectedOps := convertedMap.WhereOpEquals(apitype.OpDelete, apitype.OpDeleteReplaced, apitype.OpReplace, apitype.OpUpdate)

	assert.NotEmpty(t, expectedOps, "expected changes, but preview returned no changes: %s", preview.StdOut)
}
