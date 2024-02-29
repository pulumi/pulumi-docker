package provider

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	pulumirpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
)

// TestBuildxConfigure exercises our legacy JSON config -> go-provider bridge.
func TestBuildxConfigure(t *testing.T) {
	p, err := makeProvider(nil, "docker", "4.6.0", nil)

	require.NoError(t, err)

	sshOpts, err := json.Marshal([]string{"StrictHostKeyChecking=no"})
	require.NoError(t, err)

	m := resource.PropertyMap{
		"sshOpts": resource.NewStringProperty(string(sshOpts)),
	}

	args, err := structpb.NewStruct(m.Mappable())
	require.NoError(t, err)

	_, err = p.Configure(context.Background(), &pulumirpc.ConfigureRequest{
		Variables: map[string]string{
			"sshOpts": string(sshOpts),
		},
		Args: args,
	})

	assert.NoError(t, err)
}
