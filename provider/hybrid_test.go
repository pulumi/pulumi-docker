package provider

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
	rpc "github.com/pulumi/pulumi/sdk/v3/proto/go"

	"github.com/pulumi/pulumi-docker/provider/v4/pkg/version"
)

type mockResourceProviderServer struct {
	mock.Mock
	rpc.UnimplementedResourceProviderServer
}

type diffConfig struct {
	olds resource.PropertyMap
	news resource.PropertyMap
}

func (m *mockResourceProviderServer) DiffConfig(ctx context.Context, req *rpc.DiffRequest) (*rpc.DiffResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*rpc.DiffResponse), args.Error(1)
}

func TestDiffConfig(t *testing.T) {
	tests := []struct {
		name     string
		input    diffConfig
		expected diffConfig
	}{
		{
			name: "empty config",
			input: diffConfig{
				olds: resource.PropertyMap{},
				news: resource.PropertyMap{},
			},
			expected: diffConfig{
				olds: resource.PropertyMap{},
				news: resource.PropertyMap{},
			},
		},
		{
			name: "unwraps nested json",
			input: diffConfig{
				olds: resource.PropertyMap{
					"registryAuth": resource.NewStringProperty(
						"[{\"address\":\"dockerhub\", \"username\":\"bob\", \"password\":\"supersecret\"}]"),
				},
				news: resource.PropertyMap{
					"registryAuth": resource.NewStringProperty(
						"[{\"address\":\"ShinyPrivateGHCR\", \"username\":\"alice\", \"password\":\"moresecret\"}]"),
				},
			},
			expected: diffConfig{
				olds: resource.PropertyMap{"registryAuth": resource.NewArrayProperty([]resource.PropertyValue{
					{
						V: resource.PropertyMap{
							"address":  resource.NewStringProperty("dockerhub"),
							"username": resource.NewStringProperty("bob"),
							"password": resource.NewStringProperty("supersecret"),
						},
					},
				})},
				news: resource.PropertyMap{"registryAuth": resource.NewArrayProperty([]resource.PropertyValue{
					{
						V: resource.PropertyMap{
							"address":  resource.NewStringProperty("ShinyPrivateGHCR"),
							"username": resource.NewStringProperty("alice"),
							"password": resource.NewStringProperty("moresecret"),
						},
					},
				})},
			},
		},
		{
			name: "unwraps nested json with secrets",
			input: diffConfig{
				olds: resource.PropertyMap{
					"registryAuth": resource.NewStringProperty(
						"[{\"address\":\"dockerhub\", \"username\":\"bob\", \"password\":\"supersecret\"}]"),
				},
				news: resource.PropertyMap{
					"registryAuth": resource.MakeSecret(resource.NewStringProperty(
						"[{\"address\":\"ShinyPrivateGHCR\", \"username\":\"alice\", \"password\":\"moresecret\"}]")),
				},
			},
			expected: diffConfig{
				olds: resource.PropertyMap{"registryAuth": resource.NewArrayProperty([]resource.PropertyValue{
					{
						V: resource.PropertyMap{
							"address":  resource.NewStringProperty("dockerhub"),
							"username": resource.NewStringProperty("bob"),
							"password": resource.NewStringProperty("supersecret"),
						},
					},
				})},
				news: resource.PropertyMap{"registryAuth": resource.NewArrayProperty([]resource.PropertyValue{
					{
						V: resource.PropertyMap{
							"address":  resource.NewStringProperty("ShinyPrivateGHCR"),
							"username": resource.NewStringProperty("alice"),
							"password": resource.NewStringProperty("moresecret"),
						},
					},
				})},
			},
		},
		{
			name: "unwraps nested json with nested secrets",
			input: diffConfig{
				olds: resource.PropertyMap{
					"registryAuth": resource.NewStringProperty(
						"[{\"address\":\"dockerhub\",\"username\":\"bob\"," +
							"\"password\":{\"4dabf18193072939515e22adb298388d\": " +
							"\"1b47061264138c4ac30d75fd1eb44270\",\"value\": \"supersecret\"}}]"),
				},
				news: resource.PropertyMap{
					"registryAuth": resource.NewStringProperty(
						"[{\"address\":\"ShinyPrivateGHCR\",\"username\":\"alice\"," +
							"\"password\":{\"4dabf18193072939515e22adb298388d\": " +
							"\"1b47061264138c4ac30d75fd1eb44270\",\"value\": \"moresecret\"}}]"),
				},
			},
			expected: diffConfig{
				olds: resource.PropertyMap{"registryAuth": resource.NewArrayProperty([]resource.PropertyValue{
					{
						V: resource.PropertyMap{
							"address":  resource.NewStringProperty("dockerhub"),
							"username": resource.NewStringProperty("bob"),
							"password": resource.NewStringProperty("supersecret"),
						},
					},
				})},
				news: resource.PropertyMap{"registryAuth": resource.NewArrayProperty([]resource.PropertyValue{
					{
						V: resource.PropertyMap{
							"address":  resource.NewStringProperty("ShinyPrivateGHCR"),
							"username": resource.NewStringProperty("alice"),
							"password": resource.NewStringProperty("moresecret"),
						},
					},
				})},
			},
		},
		{
			name: "does not modify non-json config",
			input: diffConfig{
				olds: resource.PropertyMap{"registryAuth": resource.NewArrayProperty([]resource.PropertyValue{
					{
						V: resource.PropertyMap{
							"address":  resource.NewStringProperty("dockerhub"),
							"username": resource.NewStringProperty("bob"),
							"password": resource.NewStringProperty("supersecret"),
						},
					},
				})},
				news: resource.PropertyMap{"registryAuth": resource.NewArrayProperty([]resource.PropertyValue{
					{
						V: resource.PropertyMap{
							"address":  resource.NewStringProperty("ShinyPrivateGHCR"),
							"username": resource.NewStringProperty("alice"),
							"password": resource.NewStringProperty("moresecret"),
						},
					},
				})},
			},
			expected: diffConfig{
				olds: resource.PropertyMap{"registryAuth": resource.NewArrayProperty([]resource.PropertyValue{
					{
						V: resource.PropertyMap{
							"address":  resource.NewStringProperty("dockerhub"),
							"username": resource.NewStringProperty("bob"),
							"password": resource.NewStringProperty("supersecret"),
						},
					},
				})},
				news: resource.PropertyMap{"registryAuth": resource.NewArrayProperty([]resource.PropertyValue{
					{
						V: resource.PropertyMap{
							"address":  resource.NewStringProperty("ShinyPrivateGHCR"),
							"username": resource.NewStringProperty("alice"),
							"password": resource.NewStringProperty("moresecret"),
						},
					},
				})},
			},
		},
		{
			name: "does not modify unknown config options",
			input: diffConfig{
				olds: resource.PropertyMap{"unknownOption": resource.NewStringProperty("old")},
				news: resource.PropertyMap{"unknownOption": resource.NewStringProperty("new")},
			},
			expected: diffConfig{
				olds: resource.PropertyMap{"unknownOption": resource.NewStringProperty("old")},
				news: resource.PropertyMap{"unknownOption": resource.NewStringProperty("new")},
			},
		},
		{
			name: "keeps computed values in non-json config",
			input: diffConfig{
				olds: resource.PropertyMap{"registryAuth": resource.NewArrayProperty([]resource.PropertyValue{
					{
						V: resource.PropertyMap{
							"address":  resource.NewStringProperty("dockerhub"),
							"username": resource.NewStringProperty("bob"),
							"password": resource.NewStringProperty("supersecret"),
						},
					},
				})},
				news: resource.PropertyMap{"registryAuth": resource.NewArrayProperty([]resource.PropertyValue{
					{
						V: resource.PropertyMap{
							"address":  resource.NewStringProperty("ShinyPrivateGHCR"),
							"username": resource.NewStringProperty("alice"),
							"password": resource.NewComputedProperty(resource.Computed{
								Element: resource.NewStringProperty(""),
							}),
						},
					},
				})},
			},
			expected: diffConfig{
				olds: resource.PropertyMap{"registryAuth": resource.NewArrayProperty([]resource.PropertyValue{
					{
						V: resource.PropertyMap{
							"address":  resource.NewStringProperty("dockerhub"),
							"username": resource.NewStringProperty("bob"),
							"password": resource.NewStringProperty("supersecret"),
						},
					},
				})},
				news: resource.PropertyMap{"registryAuth": resource.NewArrayProperty([]resource.PropertyValue{
					{
						V: resource.PropertyMap{
							"address":  resource.NewStringProperty("ShinyPrivateGHCR"),
							"username": resource.NewStringProperty("alice"),
							"password": resource.NewComputedProperty(resource.Computed{
								Element: resource.NewStringProperty(""),
							}),
						},
					},
				})},
			},
		},
		{
			name: "keeps computed values in json config",
			input: diffConfig{
				olds: resource.PropertyMap{
					"registryAuth": resource.NewStringProperty(
						"[{\"address\":\"dockerhub\",\"username\":\"bob\",\"password\":\"supersecret\"}]"),
				},
				news: resource.PropertyMap{
					"registryAuth": resource.NewStringProperty(
						"[{\"address\":\"ShinyPrivateGHCR\",\"username\":\"alice\"," +
							"\"password\":\"04da6b54-80e4-46f7-96ec-b56ff0331ba9\"}]"),
				},
			},
			expected: diffConfig{
				olds: resource.PropertyMap{"registryAuth": resource.NewArrayProperty([]resource.PropertyValue{
					{
						V: resource.PropertyMap{
							"address":  resource.NewStringProperty("dockerhub"),
							"username": resource.NewStringProperty("bob"),
							"password": resource.NewStringProperty("supersecret"),
						},
					},
				})},
				news: resource.PropertyMap{"registryAuth": resource.NewArrayProperty([]resource.PropertyValue{
					{
						V: resource.PropertyMap{
							"address":  resource.NewStringProperty("ShinyPrivateGHCR"),
							"username": resource.NewStringProperty("alice"),
							"password": resource.NewComputedProperty(resource.Computed{
								Element: resource.NewStringProperty(""),
							}),
						},
					},
				})},
			},
		},
	}

	version.Version = "v4.0.0"
	prov := Provider()
	configEncoding := tfbridge.NewConfigEncoding(prov.P.Schema(), prov.Config)
	mockNativeProvider := new(mockResourceProviderServer)

	dp := dockerHybridProvider{
		name:           "test",
		nativeProvider: mockNativeProvider,
		configEncoding: configEncoding,
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset the mock calls before each test
			mockNativeProvider.Calls = nil
			mockNativeProvider.On("DiffConfig", mock.Anything, mock.Anything).Return(&rpc.DiffResponse{}, nil)

			inputOlds, err := plugin.MarshalProperties(tt.input.olds, plugin.MarshalOptions{
				Label:        "test",
				KeepUnknowns: true,
				SkipNulls:    true,
				KeepSecrets:  true,
			})
			require.NoError(t, err)
			inputNews, err := plugin.MarshalProperties(tt.input.news, plugin.MarshalOptions{
				Label:        "test",
				KeepUnknowns: true,
				SkipNulls:    true,
				KeepSecrets:  true,
			})
			require.NoError(t, err)

			_, err = dp.DiffConfig(context.TODO(), &rpc.DiffRequest{
				Urn:  "testURN",
				Olds: inputOlds,
				News: inputNews,
			})
			require.NoError(t, err)
			// Get the actual arguments that were passed
			actualCall := mockNativeProvider.Calls[0]
			actualRequest := actualCall.Arguments[1].(*rpc.DiffRequest)

			actualOlds, err := plugin.UnmarshalProperties(actualRequest.GetOlds(), plugin.MarshalOptions{
				Label:        "test",
				KeepUnknowns: true,
				SkipNulls:    true,
				KeepSecrets:  true,
			})
			require.NoError(t, err)

			actualNews, err := plugin.UnmarshalProperties(actualRequest.GetNews(), plugin.MarshalOptions{
				Label:        "test",
				KeepUnknowns: true,
				SkipNulls:    true,
				KeepSecrets:  true,
			})
			require.NoError(t, err)

			assert.Equal(t, tt.expected.olds, actualOlds)
			assert.Equal(t, tt.expected.news, actualNews)
		})
	}
}
