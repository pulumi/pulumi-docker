package provider

import (
	"context"
	"os"
	"path/filepath"
	"syscall"
	"testing"

	"github.com/docker/distribution/reference"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
	rpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
)

func TestDiffUpdates(t *testing.T) {
	t.Run("No diff happens on changed password", func(t *testing.T) {
		expected := map[string]*rpc.PropertyDiff{}
		input := map[resource.PropertyKey]resource.ValueDiff{
			"registry": {
				Object: &resource.ObjectDiff{
					Updates: map[resource.PropertyKey]resource.ValueDiff{"password": {
						Old: resource.PropertyValue{
							V: "FancyToken",
						},
						New: resource.PropertyValue{
							V: "PedestrianPassword",
						},
						Array:  (*resource.ArrayDiff)(nil),
						Object: (*resource.ObjectDiff)(nil),
					}},
				},
			},
		}
		actual := diffUpdates(input)
		assert.Equal(t, expected, actual)
	})

	t.Run("No diff happens on changed username", func(t *testing.T) {
		expected := map[string]*rpc.PropertyDiff{}
		input := map[resource.PropertyKey]resource.ValueDiff{
			"registry": {
				Object: &resource.ObjectDiff{
					Updates: map[resource.PropertyKey]resource.ValueDiff{"username": {
						Old: resource.PropertyValue{
							V: "platypus",
						},
						New: resource.PropertyValue{
							V: "Schnabeltier",
						},
						Array:  (*resource.ArrayDiff)(nil),
						Object: (*resource.ObjectDiff)(nil),
					}},
				},
			},
		}
		actual := diffUpdates(input)
		assert.Equal(t, expected, actual)
	})

	t.Run("Diff happens on changed server name", func(t *testing.T) {
		expected := map[string]*rpc.PropertyDiff{
			"registry": {
				Kind: rpc.PropertyDiff_UPDATE,
			},
		}
		input := map[resource.PropertyKey]resource.ValueDiff{
			"registry": {
				Object: &resource.ObjectDiff{
					Updates: map[resource.PropertyKey]resource.ValueDiff{"server": {
						Old: resource.PropertyValue{
							V: "dockerhub",
						},
						New: resource.PropertyValue{
							V: "ShinyPrivateGHCR",
						},
						Array:  (*resource.ArrayDiff)(nil),
						Object: (*resource.ObjectDiff)(nil),
					}},
				},
			},
		}
		actual := diffUpdates(input)
		assert.Equal(t, expected, actual)
	})

	t.Run("DiffConfig happens on changed address name", func(t *testing.T) {
		expected := map[string]*rpc.PropertyDiff{
			"registryAuth": {
				Kind: rpc.PropertyDiff_UPDATE,
			},
		}
		input := map[resource.PropertyKey]resource.ValueDiff{
			"registryAuth": {
				Array: &resource.ArrayDiff{
					Updates: map[int]resource.ValueDiff{
						0: {
							Object: &resource.ObjectDiff{
								Updates: map[resource.PropertyKey]resource.ValueDiff{
									"address": {
										Old: resource.PropertyValue{
											V: "dockerhub",
										},
										New: resource.PropertyValue{
											V: "ShinyPrivateGHCR",
										},
									},
								},
							},
						},
					},
				},
			},
		}
		actual := diffUpdates(input)
		assert.Equal(t, expected, actual)
	})

	t.Run("No DiffConfig happens on changed password", func(t *testing.T) {
		expected := map[string]*rpc.PropertyDiff{}
		input := map[resource.PropertyKey]resource.ValueDiff{
			"registryAuth": {
				Array: &resource.ArrayDiff{
					Updates: map[int]resource.ValueDiff{
						0: {
							Object: &resource.ObjectDiff{
								Updates: map[resource.PropertyKey]resource.ValueDiff{
									"password": {
										Old: resource.PropertyValue{
											V: "platypus",
										},
										New: resource.PropertyValue{
											V: "Schnabeltier",
										},
									},
								},
							},
						},
					},
				},
			},
		}
		actual := diffUpdates(input)
		assert.Equal(t, expected, actual)
	})

	t.Run("No DiffConfig happens on no changes", func(t *testing.T) {
		expected := map[string]*rpc.PropertyDiff{}
		input := map[resource.PropertyKey]resource.ValueDiff{
			"registryAuth": {},
		}
		actual := diffUpdates(input)
		assert.Equal(t, expected, actual)
	})

	t.Run("Diff happens on unknown new registry", func(t *testing.T) {
		expected := map[string]*rpc.PropertyDiff{
			"registry": {
				Kind: rpc.PropertyDiff_UPDATE,
			},
		}
		input := map[resource.PropertyKey]resource.ValueDiff{
			"registry": {
				Old: resource.NewObjectProperty(resource.PropertyMap{
					"server":   resource.NewStringProperty("https://index.docker.io/v1/"),
					"username": resource.NewStringProperty("pulumipus"),
					"password": resource.NewStringProperty("supersecret"),
				}),
				New: resource.NewComputedProperty(resource.Computed{Element: resource.NewStringProperty("X")}),
			},
		}
		actual := diffUpdates(input)
		assert.Equal(t, expected, actual)
	})

	t.Run("Diff happens on changed build context", func(t *testing.T) {
		expected := map[string]*rpc.PropertyDiff{
			"build": {
				Kind: rpc.PropertyDiff_UPDATE,
			},
		}
		input := map[resource.PropertyKey]resource.ValueDiff{
			"build": {
				Object: &resource.ObjectDiff{
					Updates: map[resource.PropertyKey]resource.ValueDiff{"contextDigest": {
						Old: resource.PropertyValue{
							V: "12345",
						},
						New: resource.PropertyValue{
							V: "54321",
						},
						Array:  (*resource.ArrayDiff)(nil),
						Object: (*resource.ObjectDiff)(nil),
					}},
				},
			},
		}
		actual := diffUpdates(input)
		assert.Equal(t, expected, actual)
	})
}

func TestHashIgnoresFile(t *testing.T) {
	step1Dir := "./testdata/ignores/basedir"
	baseResult, err := hashContext(step1Dir, filepath.Join(step1Dir, defaultDockerfile))
	require.NoError(t, err)

	step2Dir := "./testdata/ignores/basedir-with-ignored-files"
	result, err := hashContext(step2Dir, filepath.Join(step2Dir, defaultDockerfile))
	require.NoError(t, err)

	assert.Equal(t, result, baseResult)
}

// Tests that we handle .dockerignore exclusions such as "!foo/*/bar".
//
// See:
// - https://github.com/moby/moby/issues/30018
// - https://github.com/moby/moby/issues/45608
//
// Buildkit handles these correctly (according to spec), Docker's classic builder does not.
func TestHashIgnoresWildcards(t *testing.T) {
	baselineDir := "testdata/ignores-wildcard/basedir"
	baselineResult, err := hashContext(baselineDir, filepath.Join(baselineDir, defaultDockerfile))
	require.NoError(t, err)

	modIgnoredDir := "testdata/ignores-wildcard/basedir-modified-ignored-file"
	modIgnoredResult, err := hashContext(modIgnoredDir, filepath.Join(modIgnoredDir, defaultDockerfile))
	require.NoError(t, err)

	modIncludedDir := "testdata/ignores-wildcard/basedir-modified-included-file"
	modIncludedResult, err := hashContext(modIncludedDir, filepath.Join(modIncludedDir, defaultDockerfile))
	require.NoError(t, err)

	assert.Equal(t, baselineResult, modIgnoredResult, "hash should not change when modifying ignored files")
	assert.NotEqual(t, baselineResult, modIncludedResult,
		"hash should change when modifying included (via wildcard ignore exclusion) files")
}

// Tests that we handle .dockerignore exclusions such as "!foo/*/bar", as above, when using a
// relative context path.
func TestHashIgnoresWildcardsRelative(t *testing.T) {
	err := os.Chdir("pkg")
	require.NoError(t, err)
	defer func() {
		err = os.Chdir("..")
		require.NoError(t, err)
	}()

	baselineDir := "../testdata/ignores-wildcard/basedir"
	baselineResult, err := hashContext(baselineDir, filepath.Join(baselineDir, defaultDockerfile))
	require.NoError(t, err)

	modIgnoredDir := "../testdata/ignores-wildcard/basedir-modified-ignored-file"
	modIgnoredResult, err := hashContext(modIgnoredDir, filepath.Join(modIgnoredDir, defaultDockerfile))
	require.NoError(t, err)

	modIncludedDir := "../testdata/ignores-wildcard/basedir-modified-included-file"
	modIncludedResult, err := hashContext(modIncludedDir, filepath.Join(modIncludedDir, defaultDockerfile))
	require.NoError(t, err)

	assert.Equal(t, baselineResult, modIgnoredResult, "hash should not change when modifying ignored files")
	assert.NotEqual(t, baselineResult, modIncludedResult,
		"hash should change when modifying included (via wildcard ignore exclusion) files")
}

func TestHashIgnoresDockerfileOutsideDirMove(t *testing.T) {
	appDir := "./testdata/dockerfile-location-irrelevant/app"
	baseResult, err := hashContext(appDir, "./testdata/dockerfile-location-irrelevant/step1.Dockerfile")
	require.NoError(t, err)

	result, err := hashContext(appDir, "./testdata/dockerfile-location-irrelevant/step2.Dockerfile")
	require.NoError(t, err)

	assert.Equal(t, result, baseResult)
}

func TestHashRenamingMatters(t *testing.T) {
	step1Dir := "./testdata/filemode-matters/step1"
	baseResult, err := hashContext(step1Dir, filepath.Join(step1Dir, defaultDockerfile))
	require.NoError(t, err)

	step2Dir := "./testdata/renaming-matters/step2"
	result, err := hashContext(step2Dir, filepath.Join(step2Dir, defaultDockerfile))
	require.NoError(t, err)

	assert.NotEqual(t, result, baseResult)
}

func TestHashFilemodeMatters(t *testing.T) {
	step1Dir := "./testdata/filemode-matters/step1"
	baseResult, err := hashContext(step1Dir, filepath.Join(step1Dir, defaultDockerfile))
	require.NoError(t, err)

	step2Dir := "./testdata/filemode-matters/step2-chmod-x"
	result, err := hashContext(step2Dir, filepath.Join(step2Dir, defaultDockerfile))
	require.NoError(t, err)

	assert.NotEqual(t, result, baseResult)
}

func TestHashDeepSymlinks(t *testing.T) {
	dir := "./testdata/symlinks"
	_, err := hashContext(dir, filepath.Join(dir, "Dockerfile"))
	assert.NoError(t, err)
}

func TestIgnoreIrregularFiles(t *testing.T) {
	dir := t.TempDir()

	// Create a Dockerfile
	dockerfile := filepath.Join(dir, "Dockerfile")
	err := os.WriteFile(dockerfile, []byte{}, 0o600)
	require.NoError(t, err)

	// Create a pipe which should be ignored. (We will time out trying to read
	// it if it's not.)
	pipe := filepath.Join(dir, "pipe")
	err = syscall.Mkfifo(pipe, 0o666)
	require.NoError(t, err)
	// Confirm it's irregular.
	fi, err := os.Stat(pipe)
	require.NoError(t, err)
	assert.False(t, fi.Mode().IsRegular())

	_, err = hashContext(dir, dockerfile)
	assert.NoError(t, err)
}

func TestHashUnignoredDirs(t *testing.T) {
	step1Dir := "./testdata/unignores/basedir"
	baseResult, err := hashContext(step1Dir, filepath.Join(step1Dir, defaultDockerfile))
	require.NoError(t, err)

	step2Dir := "./testdata/unignores/basedir-with-unignored-files"
	unignoreResult, err := hashContext(step2Dir, filepath.Join(step2Dir, defaultDockerfile))
	require.NoError(t, err)

	assert.Equal(t, baseResult, unignoreResult)
}

func TestSetConfiguration(t *testing.T) {
	t.Run("Sets provider config correctly when passed a valid input map", func(t *testing.T) {
		expected := map[string]string{
			"host":       "thisisatesthost",
			"caMaterial": "materialsareweird",
		}
		input := map[string]string{
			"host":       "thisisatesthost",
			"caMaterial": "materialsareweird",
		}
		actual := setConfiguration(input)
		assert.Equal(t, expected, actual)
	})
	t.Run("Sets provider config correctly from environment variables", func(t *testing.T) {
		expected := map[string]string{
			"host":       "thisisatesthost",
			"caMaterial": "materialsareweird",
		}
		t.Setenv("DOCKER_HOST", "thisisatesthost")
		t.Setenv("DOCKER_CA_MATERIAL", "materialsareweird")
		input := map[string]string{}
		actual := setConfiguration(input)
		assert.Equal(t, expected, actual)
	})
	t.Run("Sets provider config with preference to stack config variables", func(t *testing.T) {
		expected := map[string]string{
			"host":       "thisisatesthost",
			"caMaterial": "materialsareweird",
		}
		input := map[string]string{
			"host":       "thisisatesthost",
			"caMaterial": "materialsareweird",
		}

		t.Setenv("DOCKER_HOST", "thishostshouldbeignored")

		actual := setConfiguration(input)
		assert.Equal(t, expected, actual)
	})
	t.Run("Sets provider config by correctly merging stack config and env vars", func(t *testing.T) {
		expected := map[string]string{
			"host":        "thisisatesthost",
			"caMaterial":  "materialsareweird",
			"authConfigs": "authConfigs",
		}
		input := map[string]string{
			"caMaterial":  "materialsareweird",
			"authConfigs": "authConfigs",
		}

		t.Setenv("DOCKER_HOST", "thisisatesthost")

		actual := setConfiguration(input)
		assert.Equal(t, expected, actual)
	})
}

func TestCheck(t *testing.T) {
	tests := []struct {
		name string
		news resource.PropertyMap

		wantErr error
	}{
		{
			name: "can't push a non-canonical image name",
			news: resource.PropertyMap{
				"imageName": resource.NewStringProperty("not-fully-qualified-image-name:latest"),
				"build": resource.NewObjectProperty(
					resource.PropertyMap{
						"dockerfile": resource.NewStringProperty("testdata/Dockerfile"),
					},
				),
			},
			wantErr: reference.ErrNameNotCanonical,
		},
		{
			name: "image name can be non-canonical if not pushing",
			news: resource.PropertyMap{
				"imageName": resource.NewStringProperty("not-pushing:latest"),
				"skipPush":  resource.NewBoolProperty(true),
				"build": resource.NewObjectProperty(
					resource.PropertyMap{
						"dockerfile": resource.NewStringProperty("testdata/Dockerfile"),
					},
				),
			},
			wantErr: nil,
		},
		{
			name: "image name can be non-canonical if registry server is provided",
			news: resource.PropertyMap{
				"imageName": resource.NewStringProperty("foo/bar:latest"),
				"skipPush":  resource.NewBoolProperty(true),
				"build": resource.NewObjectProperty(
					resource.PropertyMap{
						"dockerfile": resource.NewStringProperty("testdata/Dockerfile"),
					},
				),
				"registry": resource.NewObjectProperty(
					resource.PropertyMap{
						"server": resource.NewStringProperty("docker.io"),
					},
				),
			},
			wantErr: nil,
		},
		{
			name: "image name must be canonical if using caching, even when not pushing",
			news: resource.PropertyMap{
				"imageName": resource.NewStringProperty("not-pushing:latest"),
				"skipPush":  resource.NewBoolProperty(true),
				"build": resource.NewObjectProperty(
					resource.PropertyMap{
						"dockerfile": resource.NewStringProperty("testdata/Dockerfile"),
						"cacheFrom": resource.NewObjectProperty(
							resource.PropertyMap{
								"images": resource.NewArrayProperty(
									[]resource.PropertyValue{resource.NewStringProperty("docker.io/pulumi/pulumi:latest")},
								),
							},
						),
					},
				),
			},
			wantErr: reference.ErrNameNotCanonical,
		},
		{
			name: "cacheFrom can infer host from imageName",
			news: resource.PropertyMap{
				"imageName": resource.NewStringProperty("docker.io/foo/bar:latest"),
				"build": resource.NewObjectProperty(
					resource.PropertyMap{
						"dockerfile": resource.NewStringProperty("testdata/Dockerfile"),
						"cacheFrom": resource.NewObjectProperty(
							resource.PropertyMap{
								"images": resource.NewArrayProperty(
									[]resource.PropertyValue{resource.NewStringProperty("foo/bar:latest")},
								),
							},
						),
					},
				),
			},
			wantErr: nil,
		},
		{
			name: "can use non-canonical cacheFrom with a registry server",
			news: resource.PropertyMap{
				"imageName": resource.NewStringProperty("foo/bar:latest"),
				"build": resource.NewObjectProperty(
					resource.PropertyMap{
						"dockerfile": resource.NewStringProperty("testdata/Dockerfile"),
						"cacheFrom": resource.NewObjectProperty(
							resource.PropertyMap{
								"images": resource.NewArrayProperty(
									[]resource.PropertyValue{resource.NewStringProperty("not-fully-qualified-cache:latest")},
								),
							},
						),
					},
				),
				"registry": resource.NewObjectProperty(
					resource.PropertyMap{
						"server": resource.NewStringProperty("docker.io"),
					},
				),
			},
			wantErr: nil,
		},
		{
			name: "validation is skipped if imageName is unknown",
			news: resource.PropertyMap{
				"imageName": resource.NewComputedProperty(resource.Computed{Element: resource.NewStringProperty("foo")}),
				"build": resource.NewObjectProperty(
					resource.PropertyMap{
						"dockerfile": resource.NewStringProperty("testdata/Dockerfile"),
						"cacheFrom": resource.NewObjectProperty(
							resource.PropertyMap{
								"images": resource.NewArrayProperty(
									[]resource.PropertyValue{resource.NewStringProperty("foo/bar:latest")},
								),
							},
						),
					},
				),
			},
			wantErr: nil,
		},
		{
			name: "build is unknown",
			news: resource.PropertyMap{
				"imageName": resource.NewStringProperty("docker.io/foo/bar:latest"),
				"build":     resource.NewComputedProperty(resource.Computed{Element: resource.NewStringProperty("a")}),
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := dockerNativeProvider{}

			news, err := plugin.MarshalProperties(tt.news, plugin.MarshalOptions{KeepUnknowns: true})
			require.NoError(t, err)

			req := &rpc.CheckRequest{
				Urn:  string("urn:pulumi:test::docker-provider::docker:index/image:Image::foo"),
				News: news,
			}

			_, err = p.Check(context.Background(), req)

			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}

func TestCanPreview(t *testing.T) {
	tests := []struct {
		name string

		inputs resource.PropertyMap
		want   bool
	}{
		{
			name: "buildOnPreview is unknown",
			inputs: resource.PropertyMap{
				"buildOnPreview": resource.NewComputedProperty(
					resource.Computed{},
				),
			},
			want: false,
		},
		{
			name: "buildOnPreview is false",
			inputs: resource.PropertyMap{
				"buildOnPreview": resource.NewBoolProperty(false),
			},
			want: false,
		},
		{
			name: "dockerfile is unknown",
			inputs: resource.PropertyMap{
				"buildOnPreview": resource.NewBoolProperty(true),
				"build": resource.NewObjectProperty(resource.PropertyMap{
					"dockerfile": resource.NewComputedProperty(resource.Computed{}),
					"context":    resource.NewStringProperty("."),
					"args":       resource.NewObjectProperty(resource.PropertyMap{}),
				}),
			},
			want: false,
		},
		{
			name: "context is unknown",
			inputs: resource.PropertyMap{
				"buildOnPreview": resource.NewBoolProperty(true),
				"build": resource.NewObjectProperty(resource.PropertyMap{
					"dockerfile": resource.NewStringProperty("Dockerfile"),
					"context":    resource.NewComputedProperty(resource.Computed{}),
					"args":       resource.NewObjectProperty(resource.PropertyMap{}),
				}),
			},
			want: false,
		},
		{
			name: "args is unknown",
			inputs: resource.PropertyMap{
				"buildOnPreview": resource.NewBoolProperty(true),
				"build": resource.NewObjectProperty(resource.PropertyMap{
					"dockerfile": resource.NewStringProperty("Dockerfile"),
					"context":    resource.NewStringProperty("."),
					"args":       resource.NewComputedProperty(resource.Computed{}),
				}),
			},
			want: false,
		},
		{
			name: "args contains unknown",
			inputs: resource.PropertyMap{
				"buildOnPreview": resource.NewBoolProperty(true),
				"build": resource.NewObjectProperty(resource.PropertyMap{
					"dockerfile": resource.NewStringProperty("Dockerfile"),
					"context":    resource.NewStringProperty("."),
					"args": resource.NewObjectProperty(resource.PropertyMap{
						"unknown": resource.NewComputedProperty(resource.Computed{}),
					}),
				}),
			},
			want: false,
		},
		{
			name: "everything known",
			inputs: resource.PropertyMap{
				"buildOnPreview": resource.NewBoolProperty(true),
				"build": resource.NewObjectProperty(resource.PropertyMap{
					"dockerfile": resource.NewStringProperty("Dockerfile"),
					"context":    resource.NewStringProperty("."),
					"args":       resource.NewObjectProperty(resource.PropertyMap{}),
				}),
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &dockerNativeProvider{}
			actual, err := p.canPreview(context.Background(), tt.inputs, resource.URN("a"))
			require.NoError(t, err)
			assert.Equal(t, tt.want, actual)
		})
	}
}
