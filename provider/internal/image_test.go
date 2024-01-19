package internal

import (
	"context"
	"errors"
	"testing"

	_ "github.com/docker/buildx/driver/docker-container"

	controllerapi "github.com/docker/buildx/controller/pb"
	manifesttypes "github.com/docker/cli/cli/manifest/types"
	"github.com/docker/distribution/reference"
	"github.com/docker/docker/api/types"
	"github.com/moby/buildkit/client"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	provider "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/integration"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"

	"github.com/pulumi/pulumi-docker/provider/v4/internal/mock"
)

func TestLifecycle(t *testing.T) {
	// realClient := func(t *testing.T) Client { return nil }
	noClient := func(t *testing.T) Client {
		ctrl := gomock.NewController(t)
		return mock.NewMockClient(ctrl)
	}

	tests := []struct {
		name string

		op     func(t *testing.T) integration.Operation
		client func(t *testing.T) Client
	}{
		{
			name: "happy path builds",
			client: func(t *testing.T) Client {
				ctrl := gomock.NewController(t)
				c := mock.NewMockClient(ctrl)
				c.EXPECT().Auth(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
				gomock.InOrder(
					c.EXPECT().BuildKitEnabled().Return(true, nil), // Preview.
					c.EXPECT().BuildKitEnabled().Return(true, nil), // Create.
					c.EXPECT().Build(gomock.Any(), gomock.AssignableToTypeOf(controllerapi.BuildOptions{})).DoAndReturn(
						func(_ context.Context, opts controllerapi.BuildOptions) (*client.SolveResponse, error) {
							assert.Equal(t, "../testdata/Dockerfile", opts.DockerfileName)
							return &client.SolveResponse{ExporterResponse: map[string]string{"containerimage.digest": "SHA256:digest"}}, nil
						},
					),
					c.EXPECT().Inspect(gomock.Any(), "docker.io/blampe/buildkit-e2e").Return(
						[]manifesttypes.ImageManifest{}, nil,
					),
					c.EXPECT().Inspect(gomock.Any(), "docker.io/blampe/buildkit-e2e:main"),
					c.EXPECT().Delete(gomock.Any(), "SHA256:digest").Return(
						[]types.ImageDeleteResponseItem{{Deleted: "deleted"}, {Untagged: "untagged"}}, nil),
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
							[]resource.PropertyValue{resource.NewStringProperty("type=registry")},
						),
						"registries": resource.NewArrayProperty(
							[]resource.PropertyValue{
								resource.NewObjectProperty(resource.PropertyMap{
									"address":  resource.NewStringProperty("fakeaddress"),
									"username": resource.NewStringProperty("fakeuser"),
									"password": resource.NewStringProperty("fakepass"),
								}),
							},
						),
					},
				}
			},
		},
		{
			name:   "tags is required",
			client: noClient,
			op: func(t *testing.T) integration.Operation {
				return integration.Operation{
					Inputs:        resource.PropertyMap{},
					ExpectFailure: true,
					CheckFailures: []provider.CheckFailure{{
						Property: "tags",
						Reason:   "Missing required field 'tags' on 'internal.ImageArgs'",
					}},
				}
			},
		},
		{
			name:   "non-zero tags is required",
			client: noClient,
			op: func(t *testing.T) integration.Operation {
				return integration.Operation{
					Inputs: resource.PropertyMap{
						"tags":    resource.NewArrayProperty([]resource.PropertyValue{}),
						"context": resource.NewStringProperty("../testdata"),
					},
					ExpectFailure: true,
					CheckFailures: []provider.CheckFailure{
						{
							Property: "tags",
							Reason:   "at least one tag is required",
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
							[]resource.PropertyValue{resource.NewStringProperty("type=")},
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
				c := mock.NewMockClient(ctrl)
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
				c := mock.NewMockClient(ctrl)
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
				c := mock.NewMockClient(ctrl)
				gomock.InOrder(
					c.EXPECT().BuildKitEnabled().Return(true, nil), // Preview.
					c.EXPECT().BuildKitEnabled().Return(true, nil), // Create.
					c.EXPECT().Build(gomock.Any(), gomock.AssignableToTypeOf(controllerapi.BuildOptions{})).DoAndReturn(
						func(_ context.Context, opts controllerapi.BuildOptions) (*client.SolveResponse, error) {
							assert.Equal(t, "Dockerfile", opts.DockerfileName)
							return &client.SolveResponse{ExporterResponse: map[string]string{"image.name": "test:latest"}}, nil
						},
					),
					c.EXPECT().Delete(gomock.Any(), "test:latest").Return(nil, nil),
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
						assert.Equal(t, "Dockerfile", file.StringValue())
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
		client := mock.NewMockClient(ctrl)
		client.EXPECT().Delete(gomock.Any(), imageID).Return(nil, errNotFound{})

		s := newServer(client)
		err := s.Configure(provider.ConfigureRequest{})
		require.NoError(t, err)

		err = s.Delete(provider.DeleteRequest{
			ID:  imageID,
			Urn: resource.NewURN("test", "provider", "a", "docker:buildx/image:Image", "test"),
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
	client := mock.NewMockClient(ctrl)
	client.EXPECT().Inspect(gomock.Any(), tag).Return([]manifesttypes.ImageManifest{
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
		ID:  "tag",
		Urn: resource.NewURN("test", "provider", "a", "docker:buildx/image:Image", "test"),
		Inputs: resource.PropertyMap{
			"exports": resource.NewArrayProperty([]resource.PropertyValue{
				resource.NewStringProperty("type=registry"),
				resource.NewStringProperty("type=unrecognized"),
			}),
			"tags": resource.NewArrayProperty([]resource.PropertyValue{
				resource.NewStringProperty(tag),
			}),
		},
	})
	require.NoError(t, err)
	assert.Len(t, state.Properties["manifests"].ArrayValue(), 1)
}

func TestBuildOptionParsing(t *testing.T) {
	args := ImageArgs{
		Tags:      []string{"a/bad:tag:format"},
		Exports:   []string{"badexport,-"},
		Platforms: []string{","},
		CacheFrom: []string{"=badcachefrom"},
		CacheTo:   []string{"=badcacheto"},
	}

	_, err := args.toBuildOptions()
	assert.ErrorContains(t, err, "invalid value badexport")
	assert.ErrorContains(t, err, "platform specifier component must match")
	assert.ErrorContains(t, err, "badcachefrom")
	assert.ErrorContains(t, err, "badcacheto")
	assert.ErrorContains(t, err, "invalid reference format")
}
