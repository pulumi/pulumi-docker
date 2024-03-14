package internal

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/docker/docker/api/types/registry"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestAuth(t *testing.T) {
	h, err := newHost(nil)
	require.NoError(t, err)

	user := "pulumibot"
	if u := os.Getenv("DOCKER_HUB_USER"); u != "" {
		user = u
	}
	password := os.Getenv("DOCKER_HUB_PASSWORD")
	address := "docker.io"

	cli, err := wrap(h, RegistryAuth{
		Address:  address,
		Username: user,
		Password: password,
	})
	require.NoError(t, err)

	_, err = cli.Client().RegistryLogin(context.Background(), registry.AuthConfig{ServerAddress: address})
	assert.NoError(t, err)
}

func TestCustomHost(t *testing.T) {
	socket := "unix:///foo/bar.sock"

	t.Run("env", func(t *testing.T) {
		t.Setenv("DOCKER_HOST", socket)

		h, err := newHost(nil)
		require.NoError(t, err)
		cli, err := wrap(h)
		require.NoError(t, err)

		assert.Equal(t, socket, cli.Client().DaemonHost())
		assert.Equal(t, socket, cli.DockerEndpoint().Host)
	})

	t.Run("config", func(t *testing.T) {
		h, err := newHost(&Config{Host: socket})
		require.NoError(t, err)
		cli, err := wrap(h)
		require.NoError(t, err)

		assert.Equal(t, socket, cli.Client().DaemonHost())
		assert.Equal(t, socket, cli.DockerEndpoint().Host)
	})
}

func TestBuild(t *testing.T) {
	// Workaround for https://github.com/pulumi/pulumi-go-provider/issues/159
	ctrl, ctx := gomock.WithContext(context.Background(), t)
	pctx := NewMockProviderContext(ctrl)
	pctx.EXPECT().Log(gomock.Any(), gomock.Any()).AnyTimes()
	pctx.EXPECT().LogStatus(gomock.Any(), gomock.Any()).AnyTimes()
	pctx.EXPECT().Done().Return(ctx.Done()).AnyTimes()
	pctx.EXPECT().Value(gomock.Any()).DoAndReturn(func(key any) any { return ctx.Value(key) }).AnyTimes()
	pctx.EXPECT().Err().Return(ctx.Err()).AnyTimes()
	pctx.EXPECT().Deadline().Return(ctx.Deadline()).AnyTimes()

	tmpdir := t.TempDir()

	context := func(path string) BuildContext {
		return BuildContext{Context: Context{Location: path}}
	}
	exampleContext := context("../../examples/buildx/app")

	tests := []struct {
		name string
		skip bool
		args ImageArgs

		auths []RegistryAuth
	}{
		{
			name: "multiPlatform",
			args: ImageArgs{
				Context: exampleContext,
				Dockerfile: Dockerfile{
					Location: "../../examples/buildx/app/Dockerfile.multiPlatform",
				},
				Platforms: []Platform{"plan9/amd64", "plan9/arm64"},
			},
		},
		{
			name: "registryPush",
			skip: os.Getenv("DOCKER_HUB_PASSWORD") == "",
			args: ImageArgs{
				Context: exampleContext,
				Tags:    []string{"docker.io/pulumibot/buildkit-e2e:unit"},
				Push:    true,
			},
			auths: []RegistryAuth{{
				Address:  "docker.io",
				Username: "pulumibot",
				Password: os.Getenv("DOCKER_HUB_PASSWORD"),
			}},
		},
		{
			name: "cached",
			args: ImageArgs{
				Context:   exampleContext,
				Tags:      []string{"cached"},
				CacheTo:   []CacheTo{{Local: &CacheToLocal{Dest: filepath.Join(tmpdir, "cache"), CacheWithMode: CacheWithMode{Mode: "max"}}}},
				CacheFrom: []CacheFrom{{Local: &CacheFromLocal{Src: filepath.Join(tmpdir, "cache")}}},
			},
		},
		{
			name: "buildArgs",
			args: ImageArgs{
				Context: exampleContext,
				Dockerfile: Dockerfile{
					Location: "../../examples/buildx/app/Dockerfile.buildArgs",
				},
				BuildArgs: map[string]string{
					"SET_ME_TO_TRUE": "true",
				},
			},
		},
		{
			name: "extraHosts",
			args: ImageArgs{
				Context: exampleContext,
				Dockerfile: Dockerfile{
					Location: "../../examples/buildx/app/Dockerfile.extraHosts",
				},
				AddHosts: []string{
					"metadata.google.internal:169.254.169.254",
				},
			},
		},
		{
			name: "sshMount",
			skip: os.Getenv("SSH_AUTH_SOCK") == "",
			args: ImageArgs{
				Context: exampleContext,
				Dockerfile: Dockerfile{
					Location: "../../examples/buildx/app/Dockerfile.sshMount",
				},
				SSH: []SSH{{ID: "default"}},
			},
		},
		{
			name: "secrets",
			args: ImageArgs{
				Context: exampleContext,
				Dockerfile: Dockerfile{
					Location: "../../examples/buildx/app/Dockerfile.secrets",
				},
				Secrets: map[string]string{
					"password": "hunter2",
				},
				NoCache: true,
			},
		},
		{
			name: "labels",
			args: ImageArgs{
				Context: exampleContext,
				Labels: map[string]string{
					"description": "foo",
				},
			},
		},
		{
			name: "target",
			args: ImageArgs{
				Context: exampleContext,
				Dockerfile: Dockerfile{
					Location: "../../examples/buildx/app/Dockerfile.target",
				},
				Target: "build-me",
			},
		},
		{
			name: "namedContext",
			args: ImageArgs{
				Context: BuildContext{
					Context: Context{
						Location: "../../examples/buildx/app",
					},
					Named: NamedContexts{
						"golang:latest": Context{
							Location: "docker-image://golang@sha256:b8e62cf593cdaff36efd90aa3a37de268e6781a2e68c6610940c48f7cdf36984",
						},
					},
				},
				Dockerfile: Dockerfile{
					Location: "../../examples/buildx/app/Dockerfile.namedContexts",
				},
			},
		},
		{
			name: "remoteContext",
			args: ImageArgs{
				Context: BuildContext{
					Context: Context{
						Location: "https://raw.githubusercontent.com/pulumi/pulumi-docker/api-types/provider/testdata/Dockerfile",
					},
				},
			},
		},
		{
			name: "remoteContextWithInline",
			args: ImageArgs{
				Context: BuildContext{
					Context: Context{
						Location: "https://github.com/docker-library/hello-world.git",
					},
				},
				Dockerfile: Dockerfile{
					Inline: dedent(`
					FROM busybox
					COPY hello.c ./
					`),
				},
			},
		},
		{
			name: "inline",
			args: ImageArgs{
				Context: exampleContext,
				Dockerfile: Dockerfile{
					Inline: dedent(`
					FROM alpine
					RUN echo 👍
					`),
				},
			},
		},
		{
			name: "dockerLoad",
			args: ImageArgs{
				Context: exampleContext,
				Load:    true,
			},
		},
	}

	// Add an exec: true version for all of our test cases.
	for _, tt := range tests {
		tt := tt
		tt.name = "exec-" + tt.name
		tt.args.Exec = true
		for _, c := range tt.args.CacheTo {
			if c.Local != nil {
				c.Local.Dest += "-exec"
			}
		}
		for _, c := range tt.args.CacheFrom {
			if c.Local != nil {
				c.Local.Src += "-exec"
			}
		}
		tests = append(tests, tt)
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.skip {
				t.Skip()
			}
			h, err := newHost(nil)
			require.NoError(t, err)
			cli, err := wrap(h, tt.auths...)
			require.NoError(t, err)

			build, err := tt.args.toBuild(pctx, false)
			require.NoError(t, err)

			_, err = cli.Build(pctx, build)
			assert.NoError(t, err, cli.err.String())
		})
	}
}

func TestBuildkitEnabled(t *testing.T) {
	h, err := newHost(nil)
	require.NoError(t, err)
	cli, err := wrap(h)
	require.NoError(t, err)

	ok, err := cli.BuildKitEnabled()
	assert.NoError(t, err)
	assert.True(t, ok)
}

func TestInspect(t *testing.T) {
	h, err := newHost(nil)
	require.NoError(t, err)
	cli, err := wrap(h)
	require.NoError(t, err)

	v2, err := cli.Inspect(context.Background(), "pulumibot/myapp:buildx")
	assert.NoError(t, err)
	assert.Equal(t, 2, v2[0].OCIManifest.SchemaVersion)

	v1, err := cli.Inspect(context.Background(), "pulumi/pulumi")
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
