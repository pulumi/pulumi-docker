package internal

import (
	"context"
	"errors"
	"testing"

	_ "github.com/docker/buildx/driver/docker-container"

	controllerapi "github.com/docker/buildx/controller/pb"
	"github.com/docker/docker/api/types"
	"github.com/moby/buildkit/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	provider "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/integration"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"

	"github.com/pulumi/pulumi-docker/provider/v4/internal/mock"
)

func TestLifecycle(t *testing.T) {
	realClient := func(t *testing.T) Client { return nil }
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
			name:   "happy path builds",
			client: realClient,
			op: func(t *testing.T) integration.Operation {
				return integration.Operation{
					Inputs: resource.PropertyMap{
						"tags": resource.NewArrayProperty(
							[]resource.PropertyValue{resource.NewStringProperty("buildkit-e2e")},
						),
						"context": resource.NewArrayProperty(
							[]resource.PropertyValue{resource.NewStringProperty(
								"../testdata/ignores/basedir/",
							)},
						),
						"file": resource.NewStringProperty(
							"../testdata/ignores/basedir/Dockerfile",
						),
						"exports": resource.NewArrayProperty(
							[]resource.PropertyValue{resource.NewStringProperty("type=docker")},
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
						"context": resource.NewArrayProperty([]resource.PropertyValue{}),
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
					c.EXPECT().
						Build(gomock.Any(), gomock.AssignableToTypeOf(controllerapi.BuildOptions{})).
						DoAndReturn(
							func(_ context.Context, opts controllerapi.BuildOptions) (*client.SolveResponse, error) {
								assert.Equal(t, "Dockerfile", opts.DockerfileName)
								return &client.SolveResponse{
									ExporterResponse: map[string]string{
										"image.name": "test:latest",
									},
								}, nil
							},
						),
					c.EXPECT().Inspect(gomock.Any(), "test:latest").Return(
						types.ImageInspect{}, nil,
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
						"context": resource.NewArrayProperty(
							[]resource.PropertyValue{resource.NewStringProperty(
								"../testdata/ignores/basedir/",
							)},
						),
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
			},
		})
		assert.NoError(t, err)
	})
}
