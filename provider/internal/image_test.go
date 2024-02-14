package internal

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	_ "github.com/docker/buildx/driver/docker-container"

	"github.com/docker/buildx/util/buildflags"
	manifesttypes "github.com/docker/cli/cli/manifest/types"
	"github.com/docker/distribution/reference"
	"github.com/docker/docker/api/types/image"
	"github.com/moby/buildkit/client"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	provider "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/integration"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/mapper"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/pulumi/pulumi-docker/provider/v4/internal/properties"
)

var _fakeURN = resource.NewURN("test", "provider", "a", "docker:buildx/image:Image", "test")

func TestLifecycle(t *testing.T) {
	// realClient := func(t *testing.T) Client { return nil }
	noClient := func(t *testing.T) Client {
		ctrl := gomock.NewController(t)
		return NewMockClient(ctrl)
	}

	ref, err := reference.ParseNamed("docker.io/pulumibot/myapp")
	require.NoError(t, err)

	tests := []struct {
		name string

		op     func(t *testing.T) integration.Operation
		client func(t *testing.T) Client
	}{
		{
			name: "happy path builds",
			client: func(t *testing.T) Client {
				ctrl := gomock.NewController(t)
				c := NewMockClient(ctrl)
				c.EXPECT().Auth(gomock.Any(), "test", gomock.Any()).Return(nil).AnyTimes()
				gomock.InOrder(
					c.EXPECT().BuildKitEnabled().Return(true, nil), // Preview.
					c.EXPECT().BuildKitEnabled().Return(true, nil), // Create.
					c.EXPECT().Build(gomock.Any(), "test", gomock.AssignableToTypeOf(build{})).DoAndReturn(
						func(_ provider.Context, name string, b Build) (map[string]*client.SolveResponse, error) {
							assert.Equal(t, "../testdata/Dockerfile", b.BuildOptions().DockerfileName)
							return map[string]*client.SolveResponse{
								b.Targets()[0]: {ExporterResponse: map[string]string{"containerimage.digest": "SHA256:digest"}},
							}, nil
						},
					),
					c.EXPECT().Inspect(gomock.Any(), "test", "docker.io/blampe/buildkit-e2e").Return(
						[]manifesttypes.ImageManifest{
							{
								Ref:        &manifesttypes.SerializableNamed{Named: ref},
								Descriptor: v1.Descriptor{Platform: &v1.Platform{}},
							},
						}, nil,
					),
					c.EXPECT().Inspect(gomock.Any(), "test", "docker.io/blampe/buildkit-e2e:main"),
					c.EXPECT().Delete(gomock.Any(), "test").Return(
						[]image.DeleteResponse{{Deleted: "deleted"}, {Untagged: "untagged"}}, nil),
				)
				return c
			},
			op: func(t *testing.T) integration.Operation {
				return integration.Operation{
					Inputs: resource.PropertyMap{
						"tags": resource.NewArrayProperty(
							[]resource.PropertyValue{
								resource.NewStringProperty("docker.io/blampe/buildkit-e2e"),
								resource.NewStringProperty("docker.io/blampe/buildkit-e2e:main"),
							},
						),
						"platforms": resource.NewArrayProperty(
							[]resource.PropertyValue{
								resource.NewStringProperty("linux/arm64"),
								resource.NewStringProperty("linux/amd64"),
							},
						),
						"context": resource.NewStringProperty("../testdata"),
						"file":    resource.NewStringProperty("../testdata/Dockerfile"),
						"exports": resource.NewArrayProperty(
							[]resource.PropertyValue{
								resource.NewObjectProperty(resource.PropertyMap{
									"raw": resource.NewStringProperty("type=registry"),
								},
								),
							},
						),
						"registries": resource.NewArrayProperty(
							[]resource.PropertyValue{
								resource.NewObjectProperty(resource.PropertyMap{
									"address":  resource.NewStringProperty("fakeaddress"),
									"username": resource.NewStringProperty("fakeuser"),
									"password": resource.MakeSecret(resource.NewStringProperty("password")),
								}),
							},
						),
					},
				}
			},
		},
		{
			name:   "tags are required when pushing",
			client: noClient,
			op: func(t *testing.T) integration.Operation {
				return integration.Operation{
					Inputs: resource.PropertyMap{
						"tags":    resource.NewArrayProperty([]resource.PropertyValue{}),
						"context": resource.NewStringProperty("../testdata"),
						"exports": resource.NewArrayProperty(
							[]resource.PropertyValue{
								resource.NewObjectProperty(resource.PropertyMap{
									"raw": resource.NewStringProperty("type=registry"),
								}),
							},
						),
					},
					ExpectFailure: true,
					CheckFailures: []provider.CheckFailure{
						{
							Property: "tags",
							Reason:   "at least one tag or export name is needed when pushing to a registry",
						},
					},
				}
			},
		},
		{
			name:   "invalid exports",
			client: noClient,
			op: func(t *testing.T) integration.Operation {
				return integration.Operation{
					Inputs: resource.PropertyMap{
						"tags": resource.NewArrayProperty(
							[]resource.PropertyValue{resource.NewStringProperty("invalid-exports")},
						),
						"exports": resource.NewArrayProperty(
							[]resource.PropertyValue{
								resource.NewObjectProperty(resource.PropertyMap{
									"raw": resource.NewStringProperty("type="),
								}),
							},
						),
					},
					ExpectFailure: true,
					CheckFailures: []provider.CheckFailure{{
						Property: "exports",
						Reason:   "type is required for output",
					}},
				}
			},
		},
		{
			name: "requires buildkit",
			client: func(t *testing.T) Client {
				ctrl := gomock.NewController(t)
				c := NewMockClient(ctrl)
				gomock.InOrder(
					c.EXPECT().BuildKitEnabled().Return(false, nil), // Preview.
				)
				return c
			},
			op: func(t *testing.T) integration.Operation {
				return integration.Operation{
					Inputs: resource.PropertyMap{
						"tags": resource.NewArrayProperty(
							[]resource.PropertyValue{resource.NewStringProperty("foo")},
						),
					},
					ExpectFailure: true,
				}
			},
		},
		{
			name: "error reading DOCKER_BUILDKIT",
			client: func(t *testing.T) Client {
				ctrl := gomock.NewController(t)
				c := NewMockClient(ctrl)
				gomock.InOrder(
					c.EXPECT().
						BuildKitEnabled().
						Return(false, errors.New("invalid DOCKER_BUILDKIT")), // Preview.
				)
				return c
			},
			op: func(t *testing.T) integration.Operation {
				return integration.Operation{
					Inputs: resource.PropertyMap{
						"tags": resource.NewArrayProperty(
							[]resource.PropertyValue{resource.NewStringProperty("foo")},
						),
					},
					ExpectFailure: true,
				}
			},
		},
		{
			name: "file defaults to Dockerfile",
			client: func(t *testing.T) Client {
				ctrl := gomock.NewController(t)
				c := NewMockClient(ctrl)
				gomock.InOrder(
					c.EXPECT().BuildKitEnabled().Return(true, nil), // Preview.
					c.EXPECT().BuildKitEnabled().Return(true, nil), // Create.
					c.EXPECT().Build(gomock.Any(), "test", gomock.AssignableToTypeOf(build{})).DoAndReturn(
						func(_ provider.Context, name string, b Build) (map[string]*client.SolveResponse, error) {
							assert.Equal(t, "../testdata/Dockerfile", b.BuildOptions().DockerfileName)
							return map[string]*client.SolveResponse{
								b.Targets()[0]: {ExporterResponse: map[string]string{"image.name": "test:latest"}},
							}, nil
						},
					),
					c.EXPECT().Delete(gomock.Any(), "test").Return(nil, nil),
				)
				return c
			},
			op: func(t *testing.T) integration.Operation {
				return integration.Operation{
					Inputs: resource.PropertyMap{
						"tags": resource.NewArrayProperty(
							[]resource.PropertyValue{
								resource.NewStringProperty("default-dockerfile"),
							},
						),
						"context": resource.NewStringProperty("../testdata"),
					},
					Hook: func(_, output resource.PropertyMap) {
						file := output["file"]
						require.NotNil(t, file)
						require.True(t, file.IsString())
						assert.Equal(t, "../testdata/Dockerfile", file.StringValue())
					},
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lc := integration.LifeCycleTest{
				Resource: "docker:buildx/image:Image",
				Create:   tt.op(t),
			}
			s := newServer(tt.client(t))

			err := s.Configure(provider.ConfigureRequest{})
			require.NoError(t, err)

			lc.Run(t, s)
		})
	}
}

type errNotFound struct{}

func (errNotFound) NotFound()     {}
func (errNotFound) Error() string { return "not found " }

func TestDelete(t *testing.T) {
	t.Run("image was already deleted", func(t *testing.T) {
		imageID := "doesnt-exist"

		ctrl := gomock.NewController(t)
		client := NewMockClient(ctrl)
		client.EXPECT().Delete(gomock.Any(), imageID).Return(nil, errNotFound{})

		s := newServer(client)
		err := s.Configure(provider.ConfigureRequest{})
		require.NoError(t, err)

		err = s.Delete(provider.DeleteRequest{
			ID:  imageID,
			Urn: _fakeURN,
			Properties: resource.PropertyMap{
				"tags": resource.NewArrayProperty([]resource.PropertyValue{
					resource.NewStringProperty("tag"),
				}),
				"manifests": resource.NewArrayProperty([]resource.PropertyValue{}),
			},
		})
		assert.NoError(t, err)
	})
}

func TestRead(t *testing.T) {
	tag := "docker.io/pulumi/pulumi"
	ref, err := reference.ParseNamed(tag)
	require.NoError(t, err)

	ctrl := gomock.NewController(t)
	client := NewMockClient(ctrl)
	client.EXPECT().Inspect(gomock.Any(), "my-image", tag).Return([]manifesttypes.ImageManifest{
		{
			Descriptor: v1.Descriptor{Platform: &v1.Platform{Architecture: "arm64"}},
			Ref:        &manifesttypes.SerializableNamed{Named: ref},
		},
		{
			Descriptor: v1.Descriptor{Platform: &v1.Platform{Architecture: "unknown"}},
			Ref:        &manifesttypes.SerializableNamed{Named: ref},
		},
		{
			Descriptor: v1.Descriptor{},
		},
	}, nil)

	s := newServer(client)
	err = s.Configure(provider.ConfigureRequest{})
	require.NoError(t, err)

	state, err := s.Read(provider.ReadRequest{
		ID:  "my-image",
		Urn: _fakeURN,
		Inputs: resource.PropertyMap{
			"exports": resource.NewArrayProperty([]resource.PropertyValue{
				resource.NewObjectProperty(resource.PropertyMap{
					"raw": resource.NewStringProperty("type=registry"),
				}),
			}),
			"tags": resource.NewArrayProperty([]resource.PropertyValue{
				resource.NewStringProperty(tag),
			}),
		},
	})
	require.NoError(t, err)
	assert.Len(t, state.Properties["manifests"].ArrayValue(), 1)
}

func TestDiff(t *testing.T) {
	emptyDir := t.TempDir()

	baseArgs := ImageArgs{
		Context: emptyDir,
		File:    "../testdata/Dockerfile",
		Tags:    []string{},
	}
	baseState := ImageState{
		ContextHash: "f04bea490d45e7ae69d542846511e7c90eb683deaa1e0df19e9fca4d227265c2",
		Manifests:   []properties.Manifest{},
		ImageArgs:   baseArgs,
	}

	tests := []struct {
		name string
		olds func(*testing.T, ImageState) ImageState
		news func(*testing.T, ImageArgs) ImageArgs

		wantChanges bool
	}{
		{
			name:        "no diff if build context is unchanged",
			olds:        func(*testing.T, ImageState) ImageState { return baseState },
			news:        func(*testing.T, ImageArgs) ImageArgs { return baseArgs },
			wantChanges: false,
		},
		{
			name: "diff if build context changes",
			olds: func(*testing.T, ImageState) ImageState { return baseState },
			news: func(t *testing.T, a ImageArgs) ImageArgs {
				tmp := filepath.Join(a.Context, "tmp")
				err := os.WriteFile(tmp, []byte{}, 0o600)
				require.NoError(t, err)
				t.Cleanup(func() { _ = os.Remove(tmp) })
				return a
			},
			wantChanges: true,
		},
		{
			name: "no diff if registry password changes",
			olds: func(_ *testing.T, s ImageState) ImageState {
				s.Registries = []properties.RegistryAuth{{
					Address:  "foo",
					Username: "foo",
					Password: "foo",
				}}
				return s
			},
			news: func(_ *testing.T, a ImageArgs) ImageArgs {
				a.Registries = []properties.RegistryAuth{{
					Address:  "foo",
					Username: "foo",
					Password: "DIFFERENT PASSWORD",
				}}
				return a
			},
			wantChanges: false,
		},
		{
			name: "diff if registry added",
			olds: func(*testing.T, ImageState) ImageState { return baseState },
			news: func(_ *testing.T, a ImageArgs) ImageArgs {
				a.Registries = []properties.RegistryAuth{{}}
				return a
			},
			wantChanges: true,
		},
		{
			name: "diff if registry user changes",
			olds: func(_ *testing.T, s ImageState) ImageState {
				s.Registries = []properties.RegistryAuth{{
					Address:  "foo",
					Username: "foo",
					Password: "foo",
				}}
				return s
			},
			news: func(_ *testing.T, a ImageArgs) ImageArgs {
				a.Registries = []properties.RegistryAuth{{
					Address:  "DIFFERENT USER",
					Username: "foo",
					Password: "foo",
				}}
				return a
			},
			wantChanges: true,
		},
		{
			name: "diff if buildArgs changes",
			olds: func(*testing.T, ImageState) ImageState { return baseState },
			news: func(_ *testing.T, a ImageArgs) ImageArgs {
				a.BuildArgs = map[string]string{
					"foo": "bar",
				}
				return a
			},
			wantChanges: true,
		},
		{
			name: "diff if cacheFrom changes",
			olds: func(*testing.T, ImageState) ImageState { return baseState },
			news: func(_ *testing.T, a ImageArgs) ImageArgs {
				a.CacheFrom = []CacheFromEntry{{Raw: "a"}}
				return a
			},
			wantChanges: true,
		},
		{
			name: "diff if cacheTo changes",
			olds: func(*testing.T, ImageState) ImageState { return baseState },
			news: func(_ *testing.T, a ImageArgs) ImageArgs {
				a.CacheTo = []CacheToEntry{{Raw: "a"}}
				return a
			},
			wantChanges: true,
		},
		{
			name: "diff if context changes",
			olds: func(*testing.T, ImageState) ImageState { return baseState },
			news: func(_ *testing.T, a ImageArgs) ImageArgs {
				a.Context = "testdata/ignores"
				return a
			},
			wantChanges: true,
		},
		{
			name: "diff if file changes",
			olds: func(*testing.T, ImageState) ImageState { return baseState },
			news: func(_ *testing.T, a ImageArgs) ImageArgs {
				a.File = "testdata/ignores/basedir/Dockerfile"
				return a
			},
			wantChanges: true,
		},
		{
			name: "diff if platforms change",
			olds: func(*testing.T, ImageState) ImageState { return baseState },
			news: func(_ *testing.T, a ImageArgs) ImageArgs {
				a.Platforms = []Platform{"linux/amd64"}
				return a
			},
			wantChanges: true,
		},
		{
			name: "diff if pull changes",
			olds: func(*testing.T, ImageState) ImageState { return baseState },
			news: func(_ *testing.T, a ImageArgs) ImageArgs {
				a.Pull = true
				return a
			},
			wantChanges: true,
		},
		{
			name: "diff if builder changes",
			olds: func(*testing.T, ImageState) ImageState { return baseState },
			news: func(_ *testing.T, a ImageArgs) ImageArgs {
				a.Builder = "foo"
				return a
			},
			wantChanges: true,
		},
		{
			name: "diff if tags change",
			olds: func(*testing.T, ImageState) ImageState { return baseState },
			news: func(_ *testing.T, a ImageArgs) ImageArgs {
				a.Tags = []string{"foo"}
				return a
			},
			wantChanges: true,
		},
		{
			name: "diff if exports change",
			olds: func(*testing.T, ImageState) ImageState { return baseState },
			news: func(_ *testing.T, a ImageArgs) ImageArgs {
				a.Exports = []ExportEntry{{Raw: "foo"}}
				return a
			},
			wantChanges: true,
		},
	}

	s := newServer(nil)

	encode := func(t *testing.T, x any) resource.PropertyMap {
		raw, err := mapper.New(&mapper.Opts{IgnoreMissing: true}).Encode(x)
		require.NoError(t, err)
		return resource.NewPropertyMapFromMap(raw)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := s.Diff(provider.DiffRequest{
				Urn:  _fakeURN,
				Olds: encode(t, tt.olds(t, baseState)),
				News: encode(t, tt.news(t, baseArgs)),
			})
			assert.NoError(t, err)
			assert.Equal(t, tt.wantChanges, resp.HasChanges, resp.DetailedDiff)
		})
	}
}

func TestBuildOptions(t *testing.T) {
	t.Run("invalid inputs", func(t *testing.T) {
		args := ImageArgs{
			Tags:      []string{"a/bad:tag:format"},
			Exports:   []ExportEntry{{Raw: "badexport,-"}},
			Context:   "./testdata",
			Platforms: []Platform{","},
			CacheFrom: []CacheFromEntry{{Raw: "=badcachefrom"}},
			CacheTo:   []CacheToEntry{{Raw: "=badcacheto"}},
		}

		_, err := args.toBuildOptions(false)
		assert.ErrorContains(t, err, "invalid value badexport")
		assert.ErrorContains(t, err, "platform specifier component must match")
		assert.ErrorContains(t, err, "badcachefrom")
		assert.ErrorContains(t, err, "badcacheto")
		assert.ErrorContains(t, err, "invalid reference format")
		assert.ErrorContains(t, err, "testdata/Dockerfile: no such file or directory")
	})

	t.Run("buildOnPreview", func(t *testing.T) {
		args := ImageArgs{
			Tags:    []string{"my-tag"},
			Exports: []ExportEntry{{Registry: &ExportRegistry{ExportImage{Push: pulumi.BoolRef(true)}}}},
		}
		actual, err := args.toBuildOptions(true)
		assert.NoError(t, err)
		assert.Equal(t, "image", actual.Exports[0].Type)
		assert.Equal(t, "false", actual.Exports[0].Attrs["push"])

		actual, err = args.toBuildOptions(false)
		assert.NoError(t, err)
		assert.Equal(t, "image", actual.Exports[0].Type)
		assert.Equal(t, "true", actual.Exports[0].Attrs["push"])
	})

	t.Run("unknowns", func(t *testing.T) {
		// pulumi-go-provider gives us zero-values when a property is unknown.
		// We can't distinguish this from user-provided zero-values, but we
		// should:
		// - not fail previews due to these zero values,
		// - not attempt builds with invalid zero values,
		// - not allow invalid zero values in non-preview operations.
		unknowns := ImageArgs{
			BuildArgs: map[string]string{
				"known": "value",
				"":      "",
			},
			Builder:   "",
			CacheFrom: []CacheFromEntry{{GHA: &CacheFromGitHubActions{}}, {Raw: ""}},
			CacheTo:   []CacheToEntry{{GHA: &CacheToGitHubActions{}}, {Raw: ""}},
			Context:   "",
			Exports:   []ExportEntry{{Raw: ""}},
			File:      "",
			Platforms: []Platform{"linux/amd64", ""},
			Registries: []properties.RegistryAuth{
				{
					Address:  "",
					Password: "",
					Username: "",
				},
			},
			Tags: []string{"known", ""},
		}

		_, err := unknowns.toBuildOptions(true)
		assert.NoError(t, err)
		assert.False(t, unknowns.buildable())

		_, err = unknowns.toBuildOptions(false)
		assert.Error(t, err)
	})

	t.Run("multiple exports aren't allowed yet", func(t *testing.T) {
		args := ImageArgs{
			Exports: []ExportEntry{{Raw: "type=local"}, {Raw: "type=tar"}},
		}
		_, err := args.toBuildOptions(false)
		assert.ErrorContains(t, err, "multiple exports are currently unsupported")
	})

	t.Run("cache and export entries are union-ish", func(t *testing.T) {
		args := ImageArgs{
			Exports:   []ExportEntry{{Tar: &ExportTar{}, Local: &ExportLocal{}}},
			CacheTo:   []CacheToEntry{{Raw: "type=tar", Local: &CacheToLocal{Dest: "/foo"}}},
			CacheFrom: []CacheFromEntry{{Raw: "type=tar", Registry: &CacheFromRegistry{}}},
		}
		_, err := args.toBuildOptions(false)
		assert.ErrorContains(t, err, "exports should only specify one export type")
		assert.ErrorContains(t, err, "cacheFrom should only specify one cache type")
		assert.ErrorContains(t, err, "cacheTo should only specify one cache type")
	})
}

func TestBuildable(t *testing.T) {
	tests := []struct {
		name string
		args ImageArgs

		want bool
	}{
		{
			name: "unknown tags",
			args: ImageArgs{Tags: []string{""}},
			want: false,
		},
		{
			name: "unknown exports",
			args: ImageArgs{
				Tags:    []string{"known"},
				Exports: []ExportEntry{{Raw: ""}},
			},
			want: false,
		},
		{
			name: "unknown registry",
			args: ImageArgs{
				Tags:    []string{"known"},
				Exports: []ExportEntry{{Docker: &ExportDocker{}}},
				Registries: []properties.RegistryAuth{
					{
						Address:  "docker.io",
						Username: "foo",
						Password: "",
					},
				},
			},
			want: false,
		},
		{
			name: "known tags",
			args: ImageArgs{
				Tags: []string{"known"},
			},
			want: true,
		},
		{
			name: "known exports",
			args: ImageArgs{
				Tags:    []string{"known"},
				Exports: []ExportEntry{{Registry: &ExportRegistry{}}},
			},
			want: true,
		},
		{
			name: "known registry",
			args: ImageArgs{
				Tags:    []string{"known"},
				Exports: []ExportEntry{{Registry: &ExportRegistry{}}},
				Registries: []properties.RegistryAuth{
					{
						Address:  "docker.io",
						Username: "foo",
						Password: "bar",
					},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.args.buildable()
			assert.Equal(t, tt.want, actual)
		})
	}
}

func TestToBuilds(t *testing.T) {
	t.Run("single-platform caching", func(t *testing.T) {
		ia := ImageArgs{
			Tags:      []string{"foo", "bar"},
			Platforms: []Platform{"linux/amd64"},
			CacheTo: []CacheToEntry{
				{GHA: &CacheToGitHubActions{CacheWithMode: CacheWithMode{CacheModeMax}}},
				{
					Registry: &CacheToRegistry{
						CacheFromRegistry: CacheFromRegistry{Ref: "docker.io/foo/bar"},
					},
				},
				{
					Registry: &CacheToRegistry{
						CacheFromRegistry: CacheFromRegistry{Ref: "docker.io/foo/bar:baz"},
					},
				},
			},
			CacheFrom: []CacheFromEntry{
				{S3: &CacheFromS3{Name: "bar"}},
				{Registry: &CacheFromRegistry{Ref: "docker.io/foo/bar"}},
				{Registry: &CacheFromRegistry{Ref: "docker.io/foo/bar:baz"}},
			},
		}
		builds, err := ia.toBuilds(nil, false)
		assert.NoError(t, err)
		assert.Len(t, builds, 1)
	})

	t.Run("multi-platform caching", func(t *testing.T) {
		t.Setenv("ACTIONS_CACHE_URL", "fake-url")
		t.Setenv("ACTIONS_RUNTIME_TOKEN", "fake-token")

		ia := ImageArgs{
			Tags:      []string{"foo", "bar"},
			Platforms: []Platform{"linux/amd64", "linux/arm64"},
			CacheTo: []CacheToEntry{
				{GHA: &CacheToGitHubActions{CacheWithMode: CacheWithMode{CacheModeMax}}},
				{
					Registry: &CacheToRegistry{
						CacheFromRegistry: CacheFromRegistry{Ref: "docker.io/foo/bar"},
					},
				},
				{
					Registry: &CacheToRegistry{
						CacheFromRegistry: CacheFromRegistry{Ref: "docker.io/foo/bar:baz"},
					},
				},
			},
			CacheFrom: []CacheFromEntry{
				{S3: &CacheFromS3{Name: "bar"}},
				{Registry: &CacheFromRegistry{Ref: "docker.io/foo/bar"}},
				{Registry: &CacheFromRegistry{Ref: "docker.io/foo/bar:baz"}},
			},
		}

		builds, err := ia.toBuilds(nil, false)
		assert.NoError(t, err)

		assert.Len(t, builds, 3)

		// Build 0
		b0 := builds[0].BuildOptions()
		assert.Nil(t, b0.CacheTo)
		assert.Len(t, b0.CacheFrom, len(ia.CacheFrom)*(1+len(ia.Platforms)))
		assert.Len(t, b0.Platforms, len(ia.Platforms))

		// Build 1
		b1 := builds[1].BuildOptions()
		assert.Len(t, b1.Platforms, 1)
		assert.Equal(t, "linux/amd64", b1.Platforms[0])
		assert.Len(t, b1.Exports, 1)
		assert.Equal(t, "cacheonly", b1.Exports[0].Type)
		assert.Len(t, b1.CacheTo, len(ia.CacheTo))

		// Build 2
		b2 := builds[2].BuildOptions()
		assert.Len(t, b2.Platforms, 1)
		assert.Equal(t, "linux/arm64", b2.Platforms[0])
		assert.Len(t, b2.Exports, 1)
		assert.Equal(t, "cacheonly", b2.Exports[0].Type)
		assert.Len(t, b2.CacheTo, len(ia.CacheTo))
	})
}

func TestToCaches(t *testing.T) {
	tests := []struct {
		name      string
		platforms []string
		caches    []string

		want []string
	}{
		{
			name:      "single-platform",
			platforms: []string{"linux/amd64"},
			caches: []string{
				"type=registry,ref=docker.io/foo/bar:baz",
				"type=inline",
				"type=local,src=/foo",
				"type=s3",
			},
			want: []string{
				"type=registry,ref=docker.io/foo/bar:baz",
				"type=inline",
				"type=local,src=/foo",
				"type=s3",
			},
		},
		{
			name:      "multi-platform",
			platforms: []string{"linux/amd64", "linux/arm64", "linux/amd64"},
			caches: []string{
				"type=registry,ref=docker.io/foo/bar",
				"type=registry,ref=docker.io/foo/bar:baz",
				"type=inline",
				"type=local,src=/foo",
				"type=s3",
				"type=gha",
			},
			want: []string{
				"type=registry,ref=docker.io/foo/bar:linux-amd64",
				"type=registry,ref=docker.io/foo/bar:linux-arm64",
				"type=registry,ref=docker.io/foo/bar:baz-linux-amd64",
				"type=registry,ref=docker.io/foo/bar:baz-linux-arm64",
				"type=inline",
				"type=local,src=/foo-linux-amd64",
				"type=local,src=/foo-linux-arm64",
				"type=s3,name=linux-amd64",
				"type=s3,name=linux-arm64",
				"type=gha,scope=linux-amd64",
				"type=gha,scope=linux-arm64",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			caches, err := buildflags.ParseCacheEntry(tt.caches)
			require.NoError(t, err)
			want, err := buildflags.ParseCacheEntry(tt.want)
			require.NoError(t, err)

			actual := cachesFor(nil, caches, tt.platforms...)
			assert.Equal(t, want, actual)
		})
	}
}
