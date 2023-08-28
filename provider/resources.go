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

package provider

import (
	"fmt"
	"path/filepath"
	"unicode"

	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/pulumi/pulumi-docker/provider/v4/pkg/version"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/x"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
	"github.com/terraform-providers/terraform-provider-docker/shim"
)

const (
	// dockerPkg is the root package name for the Docker package.
	dockerPkg = "docker"

	// dockerMod is the root module for unparented Docker resources.
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

// dockerDataSource manufactures a data source toeken given a module and resource name.
func dockerDataSource(mod string, res string) tokens.ModuleMember {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return dockerMember(mod+"/"+fn, res)
}

func Provider() tfbridge.ProviderInfo {
	p := shimv2.NewProvider(shim.NewProvider())
	prov := tfbridge.ProviderInfo{
		P:                p,
		Name:             "docker",
		Description:      "A Pulumi package for interacting with Docker in Pulumi programs",
		Keywords:         []string{"pulumi", "docker"},
		License:          "Apache-2.0",
		Homepage:         "https://pulumi.io",
		Repository:       "https://github.com/pulumi/pulumi-docker",
		UpstreamRepoPath: "./upstream",
		Config: map[string]*tfbridge.SchemaInfo{
			"host": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"DOCKER_HOST"},
				},
			},
			"registry_auth": {
				Name:        "registryAuth", // not plural
				MaxItemsOne: tfbridge.False(),
			},
		},
		Resources: map[string]*tfbridge.ResourceInfo{
			"docker_config": {
				Tok: dockerResource(dockerMod, "ServiceConfig"),
			},
			"docker_container": {
				Tok:                 dockerResource(dockerMod, "Container"),
				DeleteBeforeReplace: true,
				Fields: map[string]*tfbridge.SchemaInfo{
					// Despite being a list, "command" represents a single command and should not be puralized.
					"command": { // ["echo", "1"]
						Name: "command",
					},
					// This property is named strangely in Terraform, despite allowing multiple entries it is not plural
					// and is not intended to be plural.
					"networks_advanced": {
						Name: "networksAdvanced",
					},
				},
			},
			"docker_image": {
				Tok: dockerResource(dockerMod, "RemoteImage"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"name": {Name: "name"},
				},
			},
			"docker_network":        {Tok: dockerResource(dockerMod, "Network")},
			"docker_secret":         {Tok: dockerResource(dockerMod, "Secret")},
			"docker_service":        {Tok: dockerResource(dockerMod, "Service")},
			"docker_volume":         {Tok: dockerResource(dockerMod, "Volume")},
			"docker_registry_image": {Tok: dockerResource(dockerMod, "RegistryImage")},
			"docker_plugin":         {Tok: dockerResource(dockerMod, "Plugin")},
			"docker_tag":            {Tok: dockerResource(dockerMod, "Tag")},
		},
		ExtraTypes: map[string]schema.ComplexTypeSpec{
			dockerResource(dockerMod, "Registry").String(): {
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Type:        "object",
					Description: "Describes a Docker container registry",
					Properties: map[string]schema.PropertySpec{
						"server": {
							Description: "The URL of the Docker registry server",
							TypeSpec:    schema.TypeSpec{Type: "string"},
						},
						"username": {
							Description: "The username to authenticate to the registry. " +
								"Does not cause image rebuild when changed.",
							TypeSpec: schema.TypeSpec{Type: "string"},
						},

						"password": {
							Description: "The password to authenticate to the registry. " +
								"Does not cause image rebuild when changed.",
							TypeSpec: schema.TypeSpec{Type: "string"},
							Secret:   true,
						},
					},
				},
			},
			dockerResource(dockerMod, "DockerBuild").String(): {
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Type:        "object",
					Description: "The Docker build context",
					Properties: map[string]schema.PropertySpec{
						"context": {
							Description: "The path to the build context to use.",
							TypeSpec:    schema.TypeSpec{Type: "string"},
						},
						"dockerfile": {
							Description: "The path to the Dockerfile to use.",
							TypeSpec:    schema.TypeSpec{Type: "string"},
						},
						"cacheFrom": {
							Description: "A list of image names to use as build cache. " +
								"Images provided must have a cache manifest. " +
								"Must provide authentication to cache registry.",

							TypeSpec: schema.TypeSpec{
								Ref: "#/types/docker:index/cacheFrom:CacheFrom",
							},
						},
						"args": {
							Description: "An optional map of named build-time argument variables to set " +
								"during the Docker build. This flag allows you to pass build-time variables" +
								"that can be accessed like environment variables inside the RUN instruction.",
							TypeSpec: schema.TypeSpec{
								Type: "object",
								AdditionalProperties: &schema.TypeSpec{
									Type: "string",
								},
							},
						},
						"target": {
							Description: "The target of the Dockerfile to build",
							TypeSpec:    schema.TypeSpec{Type: "string"},
						},
						"builderVersion": {
							Description: "The version of the Docker builder.",
							TypeSpec: schema.TypeSpec{
								Ref: "#/types/docker:index/builderVersion:BuilderVersion",
							},
							Default: "BuilderBuildKit",
						},
						"platform": {
							Description: "The architecture of the platform you want to build this image for, " +
								"e.g. `linux/arm64`.",
							TypeSpec: schema.TypeSpec{Type: "string"},
						},
					},
				},
			},
			dockerResource(dockerMod, "CacheFrom").String(): {
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Type:        "object",
					Description: "Contains a list of images to reference when building using a cache",
					Properties: map[string]schema.PropertySpec{
						"images": {
							Description: "Specifies cached images",
							TypeSpec: schema.TypeSpec{
								Type: "array",
								Items: &schema.TypeSpec{
									Type: "string",
								},
							},
						},
					},
				},
			},
			dockerResource(dockerMod, "BuilderVersion").String(): {
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Description: "The version of the Docker builder.",
					Type:        "string",
				},
				Enum: []schema.EnumValueSpec{
					{Name: "BuilderV1", Value: "BuilderV1",
						Description: "The first generation builder for Docker Daemon",
					},
					{Name: "BuilderBuildKit", Value: "BuilderBuildKit",
						Description: "The builder based on moby/buildkit project",
					},
				},
			},
		},
		ExtraResources: map[string]schema.ResourceSpec{
			dockerResource(dockerMod, "Image").String(): {
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Type:        "object",
					Description: docImage,
					Required: []string{
						"dockerfile", "context", "baseImageName", "registryServer", "imageName",
					},
					Properties: map[string]schema.PropertySpec{
						"imageName": {
							Description: "The fully qualified image name",
							TypeSpec:    schema.TypeSpec{Type: "string"},
						},
						"registryServer": {
							Description: "The name of the registry server hosting the image.",
							TypeSpec:    schema.TypeSpec{Type: "string"},
						},
						"baseImageName": {
							Description: "The fully qualified image name that was pushed to the registry.",
							TypeSpec:    schema.TypeSpec{Type: "string"},
						},
						"dockerfile": {
							Description: "The location of the Dockerfile relative to the docker build context.",
							TypeSpec:    schema.TypeSpec{Type: "string"},
						},
						"context": {
							Description: "The path to the build context to use.",
							TypeSpec:    schema.TypeSpec{Type: "string"},
						},
						"repoDigest": {
							Description: "The digest of the manifest pushed to the registry, e.g.: repository@<algorithm>:<hash>",
							TypeSpec:    schema.TypeSpec{Type: "string"},
						},
					},
				},
				IsComponent: false,
				InputProperties: map[string]schema.PropertySpec{
					"imageName": {
						Description: "The image name, of the format repository[:tag]. For the manifest SHA of a pushed docker image, please use `repoDigest`.",
						TypeSpec:    schema.TypeSpec{Type: "string"},
					},
					"registry": {
						Description: "The registry to push the image to",
						TypeSpec: schema.TypeSpec{
							Ref: "#/types/docker:index/registry:Registry",
						},
					},
					"build": {
						Description: "The Docker build context",
						TypeSpec: schema.TypeSpec{
							Ref: "#/types/docker:index/dockerBuild:DockerBuild",
						},
					},
					"skipPush": {
						Description: "A flag to skip a registry push.",
						TypeSpec: schema.TypeSpec{
							Type: "boolean",
						},
						Default: false,
					},
				},
				RequiredInputs: []string{"imageName"},
				Aliases: []schema.AliasSpec{
					{
						Type: pulumi.StringRef("docker:image:Image"),
					},
				},
			},
		},

		DataSources: map[string]*tfbridge.DataSourceInfo{
			"docker_network":        {Tok: dockerDataSource(dockerMod, "getNetwork")},
			"docker_registry_image": {Tok: dockerDataSource(dockerMod, "getRegistryImage")},
			"docker_plugin":         {Tok: dockerDataSource(dockerMod, "getPlugin")},
			"docker_image":          {Tok: dockerDataSource(dockerMod, "getRemoteImage")},
			"docker_logs":           {Tok: dockerDataSource(dockerMod, "getLogs")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
				"semver":         "^5.4.0",
			},
			DevDependencies: map[string]string{
				"@types/node":   "^10.0.0",
				"@types/semver": "^5.4.0",
			},
		},
		Python: &tfbridge.PythonInfo{
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumi/pulumi-%[1]s/sdk/", dockerPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				dockerPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
			Namespaces: map[string]string{
				dockerPkg: "Docker",
			},
		},
	}
	err := x.ComputeDefaults(&prov, x.TokensSingleModule("docker_", dockerMod,
		x.MakeStandardToken(dockerPkg)))
	contract.AssertNoErrorf(err, "failed to map tokens")
	prov.SetAutonaming(255, "-")

	return prov
}
