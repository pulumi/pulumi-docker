package provider

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
	rpc "github.com/pulumi/pulumi/sdk/v3/proto/go"

	"github.com/pulumi/pulumi-docker/provider/v4/pkg/version"
)

type diffConfig struct {
	oldInputs     resource.PropertyMap
	olds          resource.PropertyMap
	news          resource.PropertyMap
	ignoreChanges []string
}

type diffExpected struct {
	Replaces     []string
	Stables      []string
	Changes      rpc.DiffResponse_DiffChanges
	Diffs        []string
	DetailedDiff map[string]*propDiff
}

type propDiff struct {
	Kind      rpc.PropertyDiff_Kind
	InputDiff bool
}

func TestDiffConfig(t *testing.T) {
	tests := []struct {
		name     string
		input    diffConfig
		expected diffExpected
	}{
		{
			name: "empty config",
			input: diffConfig{
				oldInputs: resource.PropertyMap{},
				olds:      resource.PropertyMap{},
				news:      resource.PropertyMap{},
			},
			expected: diffExpected{
				Changes:      rpc.DiffResponse_DIFF_NONE,
				Replaces:     []string{},
				Diffs:        []string{},
				Stables:      nil,
				DetailedDiff: nil,
			},
		},
		{
			name: "ignores credential diff in nested json",
			input: diffConfig{
				oldInputs: resource.PropertyMap{
					"registryAuth": resource.NewStringProperty(
						"[{\"address\":\"dockerhub\", \"username\":\"bob\", \"password\":\"supersecret\"}]"),
				},
				olds: resource.PropertyMap{
					"registryAuth": resource.NewStringProperty(
						"[{\"address\":\"dockerhub\", \"username\":\"bob\", \"password\":\"supersecret\"}]"),
				},
				news: resource.PropertyMap{
					"registryAuth": resource.NewStringProperty(
						"[{\"address\":\"dockerhub\", \"username\":\"alice\", \"password\":\"moresecret\"}]"),
				},
			},
			expected: diffExpected{
				Changes:      rpc.DiffResponse_DIFF_NONE,
				Replaces:     []string{},
				Diffs:        []string{},
				Stables:      nil,
				DetailedDiff: nil,
			},
		},
		{
			name: "ignores credential diff",
			input: diffConfig{
				oldInputs: resource.PropertyMap{"registryAuth": resource.NewArrayProperty([]resource.PropertyValue{
					{
						V: resource.PropertyMap{
							"address":  resource.NewStringProperty("dockerhub"),
							"username": resource.NewStringProperty("bob"),
							"password": resource.NewStringProperty("supersecret"),
						},
					},
				})},
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
							"address":  resource.NewStringProperty("dockerhub"),
							"username": resource.NewStringProperty("alice"),
							"password": resource.NewStringProperty("moresecret"),
						},
					},
				})},
			},
			expected: diffExpected{
				Changes:      rpc.DiffResponse_DIFF_NONE,
				Replaces:     []string{},
				Diffs:        []string{},
				Stables:      nil,
				DetailedDiff: nil,
			},
		},
		{
			name: "keeps address diff in nested json",
			input: diffConfig{
				oldInputs: resource.PropertyMap{
					"registryAuth": resource.NewStringProperty(
						"[{\"address\":\"dockerhub\", \"username\":\"bob\", \"password\":\"supersecret\"}]"),
				},
				olds: resource.PropertyMap{
					"registryAuth": resource.NewStringProperty(
						"[{\"address\":\"dockerhub\", \"username\":\"bob\", \"password\":\"supersecret\"}]"),
				},
				news: resource.PropertyMap{
					"registryAuth": resource.NewStringProperty(
						"[{\"address\":\"ShinyPrivateGHCR\", \"username\":\"alice\", \"password\":\"moresecret\"}]"),
				},
			},
			expected: diffExpected{
				Changes:  rpc.DiffResponse_DIFF_SOME,
				Replaces: nil,
				Diffs:    []string{"registryAuth[0].address"},
				Stables:  nil,
				DetailedDiff: map[string]*propDiff{
					"registryAuth[0].address": {
						Kind:      rpc.PropertyDiff_UPDATE,
						InputDiff: true,
					},
				},
			},
		},
		{
			name: "keeps address diff",
			input: diffConfig{
				oldInputs: resource.PropertyMap{"registryAuth": resource.NewArrayProperty([]resource.PropertyValue{
					{
						V: resource.PropertyMap{
							"address":  resource.NewStringProperty("dockerhub"),
							"username": resource.NewStringProperty("bob"),
							"password": resource.NewStringProperty("supersecret"),
						},
					},
				})},
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
			expected: diffExpected{
				Changes:  rpc.DiffResponse_DIFF_SOME,
				Replaces: nil,
				Diffs:    []string{"registryAuth[0].address"},
				Stables:  nil,
				DetailedDiff: map[string]*propDiff{
					"registryAuth[0].address": {
						Kind:      rpc.PropertyDiff_UPDATE,
						InputDiff: true,
					},
				},
			},
		},
		{
			name: "handles secrets in nested json",
			input: diffConfig{
				oldInputs: resource.PropertyMap{
					"registryAuth": resource.NewStringProperty(
						"[{\"address\":\"dockerhub\", \"username\":\"bob\", \"password\":\"supersecret\"}]"),
				},
				olds: resource.PropertyMap{
					"registryAuth": resource.NewStringProperty(
						"[{\"address\":\"dockerhub\", \"username\":\"bob\", \"password\":\"supersecret\"}]"),
				},
				news: resource.PropertyMap{
					"registryAuth": resource.MakeSecret(resource.NewStringProperty(
						"[{\"address\":\"ShinyPrivateGHCR\", \"username\":\"alice\", \"password\":\"moresecret\"}]")),
				},
			},
			expected: diffExpected{
				Changes:  rpc.DiffResponse_DIFF_SOME,
				Replaces: nil,
				Diffs:    []string{"registryAuth[0].address"},
				Stables:  nil,
				DetailedDiff: map[string]*propDiff{
					"registryAuth[0].address": {
						Kind:      rpc.PropertyDiff_UPDATE,
						InputDiff: true,
					},
				},
			},
		},
		{
			name: "handles nested secrets in nested json",
			input: diffConfig{
				oldInputs: resource.PropertyMap{
					"registryAuth": resource.NewStringProperty(
						"[{\"address\":\"dockerhub\", \"username\":\"bob\", \"password\":\"supersecret\"}]"),
				},
				olds: resource.PropertyMap{
					"registryAuth": resource.NewStringProperty(
						"[{\"address\":\"dockerhub\", \"username\":\"bob\", \"password\":\"supersecret\"}]"),
				},
				news: resource.PropertyMap{
					"registryAuth": resource.NewStringProperty(
						"[{\"address\":\"ShinyPrivateGHCR\",\"username\":\"alice\"," +
							"\"password\":{\"4dabf18193072939515e22adb298388d\": " +
							"\"1b47061264138c4ac30d75fd1eb44270\",\"value\": \"moresecret\"}}]"),
				},
			},
			expected: diffExpected{
				Changes:  rpc.DiffResponse_DIFF_SOME,
				Replaces: nil,
				Diffs:    []string{"registryAuth[0].address"},
				Stables:  nil,
				DetailedDiff: map[string]*propDiff{
					"registryAuth[0].address": {
						Kind:      rpc.PropertyDiff_UPDATE,
						InputDiff: true,
					},
				},
			},
		},
		{
			name: "supports ignore changes",
			input: diffConfig{
				oldInputs: resource.PropertyMap{"registryAuth": resource.NewArrayProperty([]resource.PropertyValue{
					{
						V: resource.PropertyMap{
							"address":  resource.NewStringProperty("dockerhub"),
							"username": resource.NewStringProperty("bob"),
							"password": resource.NewStringProperty("supersecret"),
						},
					},
				})},
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
				ignoreChanges: []string{"registryAuth"},
			},
			expected: diffExpected{
				Changes:      rpc.DiffResponse_DIFF_NONE,
				Replaces:     []string{},
				Diffs:        []string{},
				Stables:      nil,
				DetailedDiff: nil,
			},
		},
		{
			name: "handles unknowns",
			input: diffConfig{
				oldInputs: resource.PropertyMap{"registryAuth": resource.NewArrayProperty([]resource.PropertyValue{
					{
						V: resource.PropertyMap{
							"address":  resource.NewStringProperty("dockerhub"),
							"username": resource.NewStringProperty("bob"),
							"password": resource.NewStringProperty("supersecret"),
						},
					},
				})},
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
							"address":  resource.NewStringProperty("dockerhub"),
							"username": resource.NewStringProperty("alice"),
							"password": resource.NewComputedProperty(resource.Computed{
								Element: resource.NewStringProperty(""),
							}),
						},
					},
				})},
			},
			expected: diffExpected{
				Changes:      rpc.DiffResponse_DIFF_NONE,
				Replaces:     []string{},
				Diffs:        []string{},
				Stables:      nil,
				DetailedDiff: nil,
			},
		},
		{
			name: "handles unknowns in nested json",
			input: diffConfig{
				oldInputs: resource.PropertyMap{
					"registryAuth": resource.NewStringProperty(
						"[{\"address\":\"dockerhub\", \"username\":\"bob\", \"password\":\"supersecret\"}]"),
				},
				olds: resource.PropertyMap{
					"registryAuth": resource.NewStringProperty(
						"[{\"address\":\"dockerhub\", \"username\":\"bob\", \"password\":\"supersecret\"}]"),
				},
				news: resource.PropertyMap{
					"registryAuth": resource.NewStringProperty(
						"[{\"address\":\"dockerhub\", \"username\":\"alice\"," +
							"\"password\":\"04da6b54-80e4-46f7-96ec-b56ff0331ba9\"}]"),
				},
			},
			expected: diffExpected{
				Changes:      rpc.DiffResponse_DIFF_NONE,
				Replaces:     []string{},
				Diffs:        []string{},
				Stables:      nil,
				DetailedDiff: nil,
			},
		},
	}

	version.Version = "v4.0.0"
	prov := Provider()
	configEncoding := tfbridge.NewConfigEncoding(prov.P.Schema(), prov.Config)

	dp := dockerHybridProvider{
		name:            "test",
		bridgedProvider: tfbridge.NewProvider(context.TODO(), nil, "docker", version.Version, prov.P, prov, []byte("{}")),
		configEncoding:  configEncoding,
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
			oldInputs, err := plugin.MarshalProperties(tt.input.oldInputs, plugin.MarshalOptions{
				Label:        "test",
				KeepUnknowns: true,
				SkipNulls:    true,
				KeepSecrets:  true,
			})
			require.NoError(t, err)

			res, err := dp.DiffConfig(context.TODO(), &rpc.DiffRequest{
				Urn:           "urn:pulumi:test::test::pulumi:providers:docker::prov",
				Olds:          inputOlds,
				News:          inputNews,
				OldInputs:     oldInputs,
				IgnoreChanges: tt.input.ignoreChanges,
			})
			require.NoError(t, err)

			assert.Equal(t, tt.expected.Changes, res.Changes, "Changes")
			assert.Equal(t, tt.expected.Diffs, res.Diffs, "Diffs")
			assert.Equal(t, tt.expected.Replaces, res.Replaces, "Replaces")
			assert.Equal(t, tt.expected.Stables, res.Stables, "Stables")

			if tt.expected.DetailedDiff == nil {
				assert.Nil(t, res.DetailedDiff, "DetailedDiff")
			} else {
				actualDetailedDiff := make(map[string]*propDiff, len(res.DetailedDiff))
				for k, v := range res.DetailedDiff {
					actualDetailedDiff[k] = &propDiff{
						Kind:      v.GetKind(),
						InputDiff: v.GetInputDiff(),
					}
				}
				assert.Equal(t, tt.expected.DetailedDiff, actualDetailedDiff, "DetailedDiff")
			}
		})
	}
}
