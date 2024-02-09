package internal

import (
	"fmt"
	"testing"

	"github.com/docker/buildx/util/buildflags"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
)

func TestCacheString(t *testing.T) {
	tests := []struct {
		name  string
		given fmt.Stringer
		want  string
	}{
		{
			name: "s3",
			given: CacheToEntry{S3: &CacheToS3{
				CacheFromS3: CacheFromS3{
					Region:          "us-west-2",
					Bucket:          "bucket-foo",
					Name:            "myname",
					EndpointURL:     "https://some.endpoint",
					BlobsPrefix:     "blob-prefix",
					ManifestsPrefix: "manifest-prefix",
					UsePathStyle:    pulumi.BoolRef(true),
					AccessKeyID:     "access-key-id",
					SecretAccessKey: "secret-key",
					SessionToken:    "session",
				},
			}},
			//nolint:lll
			want: "type=s3,bucket=bucket-foo,name=myname,endpoint_url=https://some.endpoint,blobs_prefix=blob-prefix,manifests_prefix=manifest-prefix,use_path_type=true,access_key_id=access-key-id,secret_access_key=secret-key,session_token=session",
		},
		{
			name:  "gha",
			given: CacheToEntry{GHA: &CacheToGitHubActions{}},
			want:  "type=gha",
		},
		{
			name:  "from-local",
			given: CacheFromEntry{Local: &CacheFromLocal{Src: "/foo/bar"}},
			want:  "type=local,src=/foo/bar",
		},
		{
			name:  "to-local",
			given: CacheToEntry{Local: &CacheToLocal{Dest: "/foo/bar"}},
			want:  "type=local,dest=/foo/bar",
		},
		{
			name:  "inline",
			given: CacheFromEntry{Inline: &CacheInline{}},
			want:  "type=inline",
		},
		{
			name:  "raw",
			given: CacheToEntry{Raw: Raw("type=gha")},
			want:  "type=gha",
		},
		{
			name: "compression",
			given: CacheToEntry{Local: &CacheToLocal{
				Dest: "/foo",
				CacheWithCompression: CacheWithCompression{
					Compression:      "gz2",
					CompressionLevel: 100,
					ForceCompression: pulumi.BoolRef(true),
				},
			}},
			want: "type=local,dest=/foo,compression=gz2,compression-level=22,force-compression=true",
		},
		{
			name: "ignore-error",
			given: CacheToEntry{
				AZBlob: &CacheToAzureBlob{CacheWithIgnoreError: CacheWithIgnoreError{pulumi.BoolRef(true)}},
			},
			want: "type=azblob,ignore-error=true",
		},
		{
			name: "oci",
			given: CacheToEntry{
				Registry: &CacheToRegistry{
					CacheFromRegistry: CacheFromRegistry{Ref: "docker.io/foo/bar:baz"},
					CacheWithOCI:      CacheWithOCI{OCI: pulumi.BoolRef(true), ImageManifest: pulumi.BoolRef(true)},
				},
			},
			want: "type=registry,ref=docker.io/foo/bar:baz,oci-mediatypes=true,image-manifest=true",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.given.String()
			assert.Equal(t, tt.want, actual)

			// Our output should be parsable by Docker.
			_, err := buildflags.ParseCacheEntry([]string{actual})
			assert.NoError(t, err)
		})
	}
}
