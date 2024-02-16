package internal

import (
	"testing"

	"github.com/docker/buildx/util/buildflags"
	"github.com/stretchr/testify/assert"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func TestExportString(t *testing.T) {
	tests := []struct {
		name  string
		given ExportEntry
		want  string
	}{
		{
			name:  "tar",
			given: ExportEntry{Tar: &ExportTar{ExportLocal: ExportLocal{Dest: "/foo"}}},
			want:  "type=tar,dest=/foo",
		},
		{
			name:  "local",
			given: ExportEntry{Local: &ExportLocal{Dest: "/bar"}},
			want:  "type=local,dest=/bar",
		},
		{
			name: "registry-with-compression",
			given: ExportEntry{Registry: &ExportRegistry{
				ExportImage: ExportImage{
					ExportWithCompression: ExportWithCompression{
						Compression:      "gz2",
						CompressionLevel: 100,
						ForceCompression: pulumi.BoolRef(true),
					},
				},
			}},
			want: "type=registry,compression=gz2,compression-level=22,force-compression=true",
		},
		{
			name: "registry-without-push",
			given: ExportEntry{Registry: &ExportRegistry{
				ExportImage: ExportImage{
					Push: pulumi.BoolRef(false),
				},
			}},
			want: "type=registry,push=false",
		},
		{
			name: "image",
			given: ExportEntry{
				Image: &ExportImage{
					Push:               pulumi.BoolRef(true),
					PushByDigest:       pulumi.BoolRef(true),
					Insecure:           pulumi.BoolRef(true),
					DanglingNamePrefix: "prefix",
					Unpack:             pulumi.BoolRef(true),
					Store:              pulumi.BoolRef(false),
				},
			},
			//nolint:lll
			want: "type=image,push=true,push-by-digest=true,insecure=true,dangling-name-prefix=prefix,unpack=true,store=false",
		},
		{
			name: "oci-with-names",
			given: ExportEntry{OCI: &ExportOCI{
				ExportDocker: ExportDocker{
					ExportWithNames: ExportWithNames{
						Names: []string{"foo", "bar"},
					},
				},
			}},
			want: "type=oci,name=foo,name=bar",
		},
		{
			name: "docker-with-annotations",
			given: ExportEntry{Docker: &ExportDocker{
				ExportWithAnnotations: ExportWithAnnotations{
					Annotations: map[string]string{
						"foo": "bar",
						"boo": "baz",
					},
				},
			}},
			want: "type=docker,annotation.boo=baz,annotation.foo=bar",
		},
		{
			name:  "raw",
			given: ExportEntry{Raw: Raw("type=docker")},
			want:  "type=docker",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.given.String()
			assert.Equal(t, tt.want, tt.given.String())

			// Our output should be parsable by Docker.
			_, err := buildflags.ParseExports([]string{actual})
			assert.NoError(t, err)
		})
	}
}