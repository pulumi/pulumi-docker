package internal

import (
	"context"
	"os"
	"testing"

	"github.com/docker/buildx/controller/pb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pulumi/pulumi-docker/provider/v4/internal/properties"
)

func TestAuth(t *testing.T) {
	d, err := newDockerClient()
	require.NoError(t, err)

	user := "pulumibot"
	if u := os.Getenv("DOCKER_HUB_USER"); u != "" {
		user = u
	}
	password := os.Getenv("DOCKER_HUB_PASSWORD")
	host := "pulumi.com" // Fake host -- we don't actually hit it.

	t.Cleanup(func() {
		_ = d.cli.ConfigFile().GetCredentialsStore(host).Erase(host)
	})

	err = d.Auth(context.Background(), properties.RegistryAuth{
		Address:  host,
		Username: user,
		Password: password,
	})
	assert.NoError(t, err)
}

func TestBuild(t *testing.T) {
	d, err := newDockerClient()
	require.NoError(t, err)

	_, err = d.Build(context.Background(), pb.BuildOptions{
		ContextPath:    "../testdata/",
		DockerfileName: "../testdata/Dockerfile",
	})
	assert.NoError(t, err)
}

func TestBuildkitEnabled(t *testing.T) {
	d, err := newDockerClient()
	require.NoError(t, err)
	ok, err := d.BuildKitEnabled()
	assert.NoError(t, err)
	assert.True(t, ok)
}

func TestInspect(t *testing.T) {
	d, err := newDockerClient()
	require.NoError(t, err)

	v2, err := d.Inspect(context.Background(), "blampe/myapp:buildx")
	assert.NoError(t, err)
	assert.Equal(t, 2, v2[0].OCIManifest.SchemaVersion)

	v1, err := d.Inspect(context.Background(), "pulumi/pulumi")
	assert.NoError(t, err)
	assert.Nil(t, v1[0].OCIManifest)
}

func TestNormalizatReference(t *testing.T) {
	tests := []struct {
		ref     string
		want    string
		wantErr string
	}{
		{
			ref:  "foo",
			want: "docker.io/library/foo:latest",
		},
		{
			ref:  "pulumi/pulumi:v3.100.0",
			want: "docker.io/pulumi/pulumi:v3.100.0",
		},
		{
			ref:     "invalid:ref:format",
			wantErr: "invalid reference format",
		},
	}

	for _, tt := range tests {
		t.Run(tt.ref, func(t *testing.T) {
			ref, err := normalizeReference(tt.ref)
			if err != nil {
				assert.ErrorContains(t, err, tt.wantErr)
			} else {
				assert.Equal(t, ref.String(), tt.want)
			}
		})
	}
}
