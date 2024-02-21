package internal

import (
	"context"
	"os"
	"testing"

	"github.com/docker/buildx/controller/pb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/pulumi/pulumi/sdk/v3/go/common/diag"
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
	name := "test-resource"

	t.Cleanup(func() {
		_ = d.cli.ConfigFile().GetCredentialsStore(host).Erase(host)
	})

	err = d.Auth(context.Background(), name, RegistryAuth{
		Address:  host,
		Username: user,
		Password: password,
	})
	assert.NoError(t, err)

	// Perform a second auth; it should be cached.
	err = d.Auth(context.Background(), name, RegistryAuth{
		Address:  host,
		Username: user,
		Password: password,
	})
	assert.NoError(t, err)
}

func TestBuild(t *testing.T) {
	d, err := newDockerClient()
	require.NoError(t, err)

	// Workaround for https://github.com/pulumi/pulumi-go-provider/issues/159
	ctrl, ctx := gomock.WithContext(context.Background(), t)
	pctx := NewMockProviderContext(ctrl)
	pctx.EXPECT().LogStatus(diag.Info, gomock.Any()).AnyTimes()
	pctx.EXPECT().Done().Return(ctx.Done()).AnyTimes()
	pctx.EXPECT().Value(gomock.Any()).DoAndReturn(func(key any) any { return ctx.Value(key) }).AnyTimes()
	pctx.EXPECT().Err().Return(ctx.Err()).AnyTimes()
	pctx.EXPECT().Deadline().Return(ctx.Deadline()).AnyTimes()

	_, err = d.Build(pctx, "resource-name", build{opts: pb.BuildOptions{
		ContextPath:    "../testdata/",
		DockerfileName: "../testdata/Dockerfile",
	}})
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

	v2, err := d.Inspect(context.Background(), "test", "pulumibot/myapp:buildx")
	assert.NoError(t, err)
	assert.Equal(t, 2, v2[0].OCIManifest.SchemaVersion)

	v1, err := d.Inspect(context.Background(), "test", "pulumi/pulumi")
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
