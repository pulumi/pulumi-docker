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
	// dockerPkg is the root package name for the Docker package.
	dockerPkg = "docker"

	// dockerMod is the root module for unparented Docker resources.
	dockerMod = "index"

	// dockerSwarmMod is the module for all Swarm-related resources.
	dockerSwarmMod = "swarm"
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

// dockerDataSource manufactures a standard resource token given a module and resource name.  It automatically uses the
// Docker package and names the file by simply lower casing the data source's first character.
func dockerDataSource(mod string, res string) tokens.ModuleMember {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return dockerMember(mod+"/"+fn, res)
}

func Provider() tfbridge.ProviderInfo {
	p := docker.Provider().(*schema.Provider)
	prov := tfbridge.ProviderInfo{
		P:           p,
		Name:        "docker",
		Description: "A Pulumi package for interacting with Docker in Pulumi programs",
		Keywords:    []string{"pulumi", "docker"},
		License:     "Apache-2.0",
		Homepage:    "https://pulumi.io",
		Repository:  "https://github.com/pulumi/pulumi-docker",
		Resources: map[string]*tfbridge.ResourceInfo{
			"docker_container": {
				Tok: dockerResource(dockerMod, "Container"),
				Fields: map[string]*tfbridge.SchemaInfo{
					// Despite being a list, "command" represents a single command and should not be puralized.
					"command": {
						Name: "command",
					},
					// This property is named strangely in Terraform, despite allowing multiple entries it is not plural
					// and is not intended to be plural.
					"networks_advanced": {
						Name: "networks_advanced",
					},
				},
			},
			"docker_image":   {Tok: dockerResource(dockerMod, "RemoteImage")},
			"docker_network": {Tok: dockerResource(dockerMod, "Network")},
			"docker_volume":  {Tok: dockerResource(dockerMod, "Volume")},
			"docker_config":  {Tok: dockerResource(dockerSwarmMod, "Config")},
			"docker_secret":  {Tok: dockerResource(dockerSwarmMod, "Secret")},
			"docker_service": {Tok: dockerResource(dockerSwarmMod, "Service")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"docker_registry_image": {Tok: dockerDataSource(dockerMod, "getRegistryImage")},
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
			Overlay: &tfbridge.OverlayInfo{
				DestFiles: []string{
					"docker.ts",
					"image.ts",
				},
			},
		},
		Python: &tfbridge.PythonInfo{
			Requires: map[string]string{
				"pulumi": ">=0.16.4",
			},
		},
	}

	// For all resources with name properties, we will add an auto-name property.  Make sure to skip those that
	// already have a name mapping entry, since those may have custom overrides set above (e.g., for length).
	const dockerName = "name"
	for resname, res := range prov.Resources {
		if schema := p.ResourcesMap[resname]; schema != nil {
			// Only apply auto-name to input properties (Optional || Required) named `name`
			if tfs, has := schema.Schema[dockerName]; has && (tfs.Optional || tfs.Required) {
				if _, hasfield := res.Fields[dockerName]; !hasfield {
					if res.Fields == nil {
						res.Fields = make(map[string]*tfbridge.SchemaInfo)
					}
					res.Fields[dockerName] = tfbridge.AutoName(dockerName, 255)
				}
			}
		}
	}

	return prov
}
