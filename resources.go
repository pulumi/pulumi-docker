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

package docker

import (
	"unicode"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-providers/terraform-provider-docker/docker"

	"github.com/pulumi/pulumi-terraform/pkg/tfbridge"
	"github.com/pulumi/pulumi/pkg/tokens"
)

const (
	dockerPkg = "docker"
	dockerMod = "index"
)

// dockerMember manufactures a type token for the docker package and the given module and type.
func dockerMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(dockerPkg + ":" + mod + ":" + mem)
}

// dockerType manufactures a type token for the docker package and the given module and type.
func dockerType(mod string, typ string) tokens.Type {
	return tokens.Type(dockerMember(mod, typ))
}

// dockerResource manufactures a standard resource token given a module and resource name.  It automatically uses the
// docker package and names the file by simply lower casing the resource's first character.
func dockerResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return dockerType(mod+"/"+fn, res)
}

func Provider() tfbridge.ProviderInfo {
	return tfbridge.ProviderInfo{
		P:           docker.Provider().(*schema.Provider),
		Name:        "docker",
		Description: "A Pulumi package for interacting with Docker in Pulumi programs",
		Keywords:    []string{"pulumi", "docker"},
		License:     "Apache-2.0",
		Homepage:    "https://pulumi.io",
		Repository:  "https://github.com/pulumi/pulumi-docker",
		Resources: map[string]*tfbridge.ResourceInfo{
			"docker_container": {
				Tok: dockerResource(dockerMod, "Container"),
			},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			Dependencies: map[string]string{
				"@pulumi/pulumi": "dev",
				"semver":         "^5.4.0",
			},
			DevDependencies: map[string]string{
				"@types/node":   "^8.0.26",
				"@types/semver": "^5.4.0",
			},
		},
		Python: &tfbridge.PythonInfo{
			Requires: map[string]string{
				"pulumi": ">=0.16.4",
			},
		},
	}
}
