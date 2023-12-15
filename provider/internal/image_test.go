package internal

import (
	"path/filepath"
	"testing"

	_ "go.uber.org/mock/gomock"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	provider "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/integration"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

func TestLifecycle(t *testing.T) {
	tests := []struct {
		name string
		op   func(t *testing.T) integration.Operation
	}{
		{
			name: "tags is required",
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
			name: "non-zero tags is required",
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
			name: "file defaults to Dockerfile",
			op: func(t *testing.T) integration.Operation {
				return integration.Operation{
					Inputs: resource.PropertyMap{
						"tags": resource.NewArrayProperty(
							[]resource.PropertyValue{resource.NewStringProperty("foo")},
						),
					},
					Hook: func(_, output resource.PropertyMap) {
						file := output["file"]
						require.NotNil(t, file)
						require.True(t, file.IsString())
						assert.Equal(t, "Dockerfile", filepath.Base(file.StringValue()))
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
			lc.Run(t, newServer())
		})
	}
}
