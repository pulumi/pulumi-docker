package provider

import (
	"context"
	"os"
	"path/filepath"
	"runtime"
	"syscall"
	"testing"

	"github.com/distribution/reference"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
	rpc "github.com/pulumi/pulumi/sdk/v3/proto/go"

	"github.com/pulumi/pulumi-docker/provider/v5/pkg/version"
)

func TestProviderHostDefaultMatchesPlatform(t *testing.T) {
	oldVersion := version.Version
	version.Version = "v4.0.0"
	t.Cleanup(func() {
		version.Version = oldVersion
	})

	info := Provider()
	host := info.Config[lintHost]
	require.NotNil(t, host)
	require.NotNil(t, host.Default)
	assert.Equal(t, []string{"DOCKER_HOST"}, host.Default.EnvVars)
	assert.Equal(t, "npipe:////./pipe/docker_engine", defaultDockerProviderHost("windows"))
	assert.Equal(t, "unix:///var/run/docker.sock", defaultDockerProviderHost("linux"))

	value, err := host.Default.ComputeDefault(context.Background(), tfbridge.ComputeDefaultOptions{})
	require.NoError(t, err)
	assert.Equal(t, defaultDockerProviderHost(runtime.GOOS), value)
}

func TestDiffUpdates(t *testing.T) {
	t.Run("No diff happens on changed password", func(t *testing.T) {
		expected := map[string]*rpc.PropertyDiff{}
		input := map[resource.PropertyKey]resource.ValueDiff{
			lintRegistry: {
				Object: &resource.ObjectDiff{
					Updates: map[resource.PropertyKey]resource.ValueDiff{lintPassword: {
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
			lintRegistry: {
				Object: &resource.ObjectDiff{
					Updates: map[resource.PropertyKey]resource.ValueDiff{lintUsername: {
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
			lintRegistry: {
				Kind: rpc.PropertyDiff_UPDATE,
			},
		}
		input := map[resource.PropertyKey]resource.ValueDiff{
			lintRegistry: {
				Object: &resource.ObjectDiff{
					Updates: map[resource.PropertyKey]resource.ValueDiff{lintServer: {
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

	t.Run("Diff happens on unknown new registry", func(t *testing.T) {
		expected := map[string]*rpc.PropertyDiff{
			lintRegistry: {
				Kind: rpc.PropertyDiff_UPDATE,
			},
		}
		input := map[resource.PropertyKey]resource.ValueDiff{
			lintRegistry: {
				Old: resource.NewObjectProperty(resource.PropertyMap{
					lintServer:   resource.NewStringProperty("https://index.docker.io/v1/"),
					lintUsername: resource.NewStringProperty(lintPulumipus),
					lintPassword: resource.NewStringProperty("supersecret"),
				}),
				New: resource.NewComputedProperty(resource.Computed{Element: resource.NewStringProperty("X")}),
			},
		}
		actual := diffUpdates(input)
		assert.Equal(t, expected, actual)
	})

	t.Run("Diff happens on changed build context", func(t *testing.T) {
		expected := map[string]*rpc.PropertyDiff{
			lintBuild: {
				Kind: rpc.PropertyDiff_UPDATE,
			},
		}
		input := map[resource.PropertyKey]resource.ValueDiff{
			lintBuild: {
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
	_, err := hashContext(dir, filepath.Join(dir, defaultDockerfile))
	assert.NoError(t, err)
}

func TestIgnoreIrregularFiles(t *testing.T) {
	dir := t.TempDir()

	// Create a Dockerfile
	dockerfile := filepath.Join(dir, defaultDockerfile)
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
			lintHost:       lintTestHost,
			lintCAMaterial: lintTestCAMaterial,
		}
		input := map[string]string{
			lintHost:       lintTestHost,
			lintCAMaterial: lintTestCAMaterial,
		}
		actual := setConfiguration(input)
		assert.Equal(t, expected, actual)
	})
	t.Run("Sets provider config correctly from environment variables", func(t *testing.T) {
		expected := map[string]string{
			lintHost:       lintTestHost,
			lintCAMaterial: lintTestCAMaterial,
		}
		t.Setenv("DOCKER_HOST", lintTestHost)
		t.Setenv("DOCKER_CA_MATERIAL", lintTestCAMaterial)
		input := map[string]string{}
		actual := setConfiguration(input)
		assert.Equal(t, expected, actual)
	})
	t.Run("Sets provider config with preference to stack config variables", func(t *testing.T) {
		expected := map[string]string{
			lintHost:       lintTestHost,
			lintCAMaterial: lintTestCAMaterial,
		}
		input := map[string]string{
			lintHost:       lintTestHost,
			lintCAMaterial: lintTestCAMaterial,
		}

		t.Setenv("DOCKER_HOST", "thishostshouldbeignored")

		actual := setConfiguration(input)
		assert.Equal(t, expected, actual)
	})
	t.Run("Sets provider config by correctly merging stack config and env vars", func(t *testing.T) {
		expected := map[string]string{
			lintHost:        lintTestHost,
			lintCAMaterial:  lintTestCAMaterial,
			lintAuthConfigs: lintAuthConfigs,
		}
		input := map[string]string{
			lintCAMaterial:  lintTestCAMaterial,
			lintAuthConfigs: lintAuthConfigs,
		}

		t.Setenv("DOCKER_HOST", lintTestHost)

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
				lintImageName: resource.NewStringProperty("not-fully-qualified-image-name:latest"),
				lintBuild: resource.NewObjectProperty(
					resource.PropertyMap{
						lintDockerfile: resource.NewStringProperty("testdata/Dockerfile"),
					},
				),
			},
			wantErr: reference.ErrNameNotCanonical,
		},
		{
			name: "image name can be non-canonical if not pushing",
			news: resource.PropertyMap{
				lintImageName: resource.NewStringProperty("not-pushing:latest"),
				lintSkipPush:  resource.NewBoolProperty(true),
				lintBuild: resource.NewObjectProperty(
					resource.PropertyMap{
						lintDockerfile: resource.NewStringProperty("testdata/Dockerfile"),
					},
				),
			},
			wantErr: nil,
		},
		{
			name: "image name can be non-canonical if registry server is provided",
			news: resource.PropertyMap{
				lintImageName: resource.NewStringProperty("foo/bar:latest"),
				lintSkipPush:  resource.NewBoolProperty(true),
				lintBuild: resource.NewObjectProperty(
					resource.PropertyMap{
						lintDockerfile: resource.NewStringProperty("testdata/Dockerfile"),
					},
				),
				lintRegistry: resource.NewObjectProperty(
					resource.PropertyMap{
						lintServer: resource.NewStringProperty("docker.io"),
					},
				),
			},
			wantErr: nil,
		},
		{
			name: "image name must be canonical if using caching, even when not pushing",
			news: resource.PropertyMap{
				lintImageName: resource.NewStringProperty("not-pushing:latest"),
				lintSkipPush:  resource.NewBoolProperty(true),
				lintBuild: resource.NewObjectProperty(
					resource.PropertyMap{
						lintDockerfile: resource.NewStringProperty("testdata/Dockerfile"),
						lintCacheFrom: resource.NewObjectProperty(
							resource.PropertyMap{
								lintImages: resource.NewArrayProperty(
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
				lintImageName: resource.NewStringProperty("docker.io/foo/bar:latest"),
				lintBuild: resource.NewObjectProperty(
					resource.PropertyMap{
						lintDockerfile: resource.NewStringProperty("testdata/Dockerfile"),
						lintCacheFrom: resource.NewObjectProperty(
							resource.PropertyMap{
								lintImages: resource.NewArrayProperty(
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
				lintImageName: resource.NewStringProperty("foo/bar:latest"),
				lintBuild: resource.NewObjectProperty(
					resource.PropertyMap{
						lintDockerfile: resource.NewStringProperty("testdata/Dockerfile"),
						lintCacheFrom: resource.NewObjectProperty(
							resource.PropertyMap{
								lintImages: resource.NewArrayProperty(
									[]resource.PropertyValue{resource.NewStringProperty("not-fully-qualified-cache:latest")},
								),
							},
						),
					},
				),
				lintRegistry: resource.NewObjectProperty(
					resource.PropertyMap{
						lintServer: resource.NewStringProperty("docker.io"),
					},
				),
			},
			wantErr: nil,
		},
		{
			name: "validation is skipped if imageName is unknown",
			news: resource.PropertyMap{
				lintImageName: resource.NewComputedProperty(resource.Computed{Element: resource.NewStringProperty("foo")}),
				lintBuild: resource.NewObjectProperty(
					resource.PropertyMap{
						lintDockerfile: resource.NewStringProperty("testdata/Dockerfile"),
						lintCacheFrom: resource.NewObjectProperty(
							resource.PropertyMap{
								lintImages: resource.NewArrayProperty(
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
				lintImageName: resource.NewStringProperty("docker.io/foo/bar:latest"),
				lintBuild:     resource.NewComputedProperty(resource.Computed{Element: resource.NewStringProperty("a")}),
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
				lintBuildOnPreview: resource.NewComputedProperty(
					resource.Computed{},
				),
			},
			want: false,
		},
		{
			name: "buildOnPreview is false",
			inputs: resource.PropertyMap{
				lintBuildOnPreview: resource.NewBoolProperty(false),
			},
			want: false,
		},
		{
			name: "dockerfile is unknown",
			inputs: resource.PropertyMap{
				lintBuildOnPreview: resource.NewBoolProperty(true),
				lintBuild: resource.NewObjectProperty(resource.PropertyMap{
					lintDockerfile: resource.NewComputedProperty(resource.Computed{}),
					lintContext:    resource.NewStringProperty("."),
					lintArgs:       resource.NewObjectProperty(resource.PropertyMap{}),
				}),
			},
			want: false,
		},
		{
			name: "context is unknown",
			inputs: resource.PropertyMap{
				lintBuildOnPreview: resource.NewBoolProperty(true),
				lintBuild: resource.NewObjectProperty(resource.PropertyMap{
					lintDockerfile: resource.NewStringProperty(defaultDockerfile),
					lintContext:    resource.NewComputedProperty(resource.Computed{}),
					lintArgs:       resource.NewObjectProperty(resource.PropertyMap{}),
				}),
			},
			want: false,
		},
		{
			name: "args is unknown",
			inputs: resource.PropertyMap{
				lintBuildOnPreview: resource.NewBoolProperty(true),
				lintBuild: resource.NewObjectProperty(resource.PropertyMap{
					lintDockerfile: resource.NewStringProperty(defaultDockerfile),
					lintContext:    resource.NewStringProperty("."),
					lintArgs:       resource.NewComputedProperty(resource.Computed{}),
				}),
			},
			want: false,
		},
		{
			name: "args contains unknown",
			inputs: resource.PropertyMap{
				lintBuildOnPreview: resource.NewBoolProperty(true),
				lintBuild: resource.NewObjectProperty(resource.PropertyMap{
					lintDockerfile: resource.NewStringProperty(defaultDockerfile),
					lintContext:    resource.NewStringProperty("."),
					lintArgs: resource.NewObjectProperty(resource.PropertyMap{
						"unknown": resource.NewComputedProperty(resource.Computed{}),
					}),
				}),
			},
			want: false,
		},
		{
			name: "everything known",
			inputs: resource.PropertyMap{
				lintBuildOnPreview: resource.NewBoolProperty(true),
				lintBuild: resource.NewObjectProperty(resource.PropertyMap{
					lintDockerfile: resource.NewStringProperty(defaultDockerfile),
					lintContext:    resource.NewStringProperty("."),
					lintArgs:       resource.NewObjectProperty(resource.PropertyMap{}),
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
