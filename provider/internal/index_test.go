package internal

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	provider "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/integration"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/mapper"
)

func TestIndexLifecycle(t *testing.T) {
	realClient := func(t *testing.T) Client { return nil }

	tests := []struct {
		name string
		skip bool

		op     func(t *testing.T) integration.Operation
		client func(t *testing.T) Client
	}{
		{
			name:   "not pushed",
			client: realClient,
			op: func(t *testing.T) integration.Operation {
				return integration.Operation{
					Inputs: resource.PropertyMap{
						"tag": resource.NewStringProperty("docker.io/pulumibot/buildkit-e2e:manifest-unit"),
						"sources": resource.NewArrayProperty([]resource.PropertyValue{
							resource.NewStringProperty("docker.io/pulumibot/buildkit-e2e:arm64"),
							resource.NewStringProperty("docker.io/pulumibot/buildkit-e2e:amd64"),
						}),
						"push": resource.NewBoolProperty(false),
					},
				}
			},
		},
		{
			name:   "pushed",
			skip:   os.Getenv("DOCKER_HUB_PASSWORD") == "",
			client: realClient,
			op: func(t *testing.T) integration.Operation {
				return integration.Operation{
					Inputs: resource.PropertyMap{
						"tag": resource.NewStringProperty("docker.io/pulumibot/buildkit-e2e:manifest"),
						"sources": resource.NewArrayProperty([]resource.PropertyValue{
							resource.NewStringProperty("docker.io/pulumibot/buildkit-e2e:arm64"),
							resource.NewStringProperty("docker.io/pulumibot/buildkit-e2e:amd64"),
						}),
						"push": resource.NewBoolProperty(true),
						"registry": resource.NewObjectProperty(resource.PropertyMap{
							"address":  resource.NewStringProperty("docker.io"),
							"username": resource.NewStringProperty("pulumibot"),
							"password": resource.NewSecretProperty(&resource.Secret{
								Element: resource.NewStringProperty(os.Getenv("DOCKER_HUB_PASSWORD")),
							}),
						}),
					},
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.skip {
				t.Skip("missing environment variables")
			}
			lc := integration.LifeCycleTest{
				Resource: "docker:buildx/image:Index",
				Create:   tt.op(t),
			}
			s := newServer(tt.client(t))

			err := s.Configure(provider.ConfigureRequest{})
			require.NoError(t, err)

			lc.Run(t, s)
		})
	}
}

func TestIndexDiff(t *testing.T) {
	urn := resource.NewURN("test", "provider", "a", "docker:buildx/image:Index", "test")
	baseArgs := IndexArgs{Sources: []string{"docker.io/nginx:latest"}}
	baseState := IndexState{IndexArgs: baseArgs}

	tests := []struct {
		name string
		olds func(*testing.T, IndexState) IndexState
		news func(*testing.T, IndexArgs) IndexArgs

		wantChanges bool
	}{
		{
			name:        "no diff if no changes",
			olds:        func(*testing.T, IndexState) IndexState { return baseState },
			news:        func(*testing.T, IndexArgs) IndexArgs { return baseArgs },
			wantChanges: false,
		},
		{
			name: "diff if tag changes",
			olds: func(*testing.T, IndexState) IndexState { return baseState },
			news: func(t *testing.T, a IndexArgs) IndexArgs {
				a.Tag = "new-tag"
				return a
			},
			wantChanges: true,
		},
		{
			name: "no diff if registry password changes",
			olds: func(_ *testing.T, s IndexState) IndexState {
				s.Registry = RegistryAuth{
					Address:  "foo",
					Username: "foo",
					Password: "foo",
				}
				return s
			},
			news: func(_ *testing.T, a IndexArgs) IndexArgs {
				a.Registry = RegistryAuth{
					Address:  "foo",
					Username: "foo",
					Password: "DIFFERENT PASSWORD",
				}
				return a
			},
			wantChanges: false,
		},
		{
			name: "diff if registry added",
			olds: func(*testing.T, IndexState) IndexState { return baseState },
			news: func(_ *testing.T, a IndexArgs) IndexArgs {
				a.Registry = RegistryAuth{Address: "foo.com", Username: "foo", Password: "foo"}
				return a
			},
			wantChanges: true,
		},
		{
			name: "diff if registry user changes",
			olds: func(_ *testing.T, s IndexState) IndexState {
				s.Registry = RegistryAuth{
					Address:  "foo",
					Username: "foo",
					Password: "foo",
				}
				return s
			},
			news: func(_ *testing.T, a IndexArgs) IndexArgs {
				a.Registry = RegistryAuth{
					Address:  "DIFFERENT USER",
					Username: "foo",
					Password: "foo",
				}
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
				Urn:  urn,
				Olds: encode(t, tt.olds(t, baseState)),
				News: encode(t, tt.news(t, baseArgs)),
			})
			assert.NoError(t, err)
			assert.Equal(t, tt.wantChanges, resp.HasChanges, resp.DetailedDiff)
		})
	}
}
