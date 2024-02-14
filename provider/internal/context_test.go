package internal

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var _dockerfile = "Dockerfile"

func TestHashIgnoresFile(t *testing.T) {
	step1Dir := "./testdata/ignores/basedir"
	baseResult, err := HashContext(step1Dir, filepath.Join(step1Dir, _dockerfile))
	require.NoError(t, err)

	step2Dir := "./testdata/ignores/basedir-with-ignored-files"
	result, err := HashContext(step2Dir, filepath.Join(step2Dir, _dockerfile))
	require.NoError(t, err)

	assert.Equal(t, result, baseResult)
}

// Tests that we handle .dockerignore exclusions such as "!foo/*/bar".
//
// See:
// - https://github.com/moby/moby/issues/30018
// - https://github.com/moby/moby/issues/45608
//
// Buildkit handles these correctly (according to spec), Docker's classic builder does not.
func TestHashIgnoresWildcards(t *testing.T) {
	baselineDir := "testdata/ignores-wildcard/basedir"
	baselineResult, err := HashContext(baselineDir, filepath.Join(baselineDir, _dockerfile))
	require.NoError(t, err)

	modIgnoredDir := "testdata/ignores-wildcard/basedir-modified-ignored-file"
	modIgnoredResult, err := HashContext(modIgnoredDir, filepath.Join(modIgnoredDir, _dockerfile))
	require.NoError(t, err)

	modIncludedDir := "testdata/ignores-wildcard/basedir-modified-included-file"
	modIncludedResult, err := HashContext(modIncludedDir, filepath.Join(modIncludedDir, _dockerfile))
	require.NoError(t, err)

	assert.Equal(t, baselineResult, modIgnoredResult, "hash should not change when modifying ignored files")
	assert.NotEqual(t, baselineResult, modIncludedResult,
		"hash should change when modifying included (via wildcard ignore exclusion) files")
}

func BenchmarkHashContext(b *testing.B) {
	dir := "testdata/ignores-wildcard/basedir-modified-ignored-file"
	for n := 0; n < b.N; n++ {
		_, err := HashContext(dir, filepath.Join(dir, _dockerfile))
		require.NoError(b, err)

	}
}

// Tests that we handle .dockerignore exclusions such as "!foo/*/bar", as above, when using a
// relative context path.
func TestHashIgnoresWildcardsRelative(t *testing.T) {
	err := os.Chdir("properties")
	require.NoError(t, err)
	defer func() {
		err = os.Chdir("..")
		require.NoError(t, err)
	}()

	baselineDir := "../testdata/ignores-wildcard/basedir"
	baselineResult, err := HashContext(baselineDir, filepath.Join(baselineDir, _dockerfile))
	require.NoError(t, err)

	modIgnoredDir := "../testdata/ignores-wildcard/basedir-modified-ignored-file"
	modIgnoredResult, err := HashContext(modIgnoredDir, filepath.Join(modIgnoredDir, _dockerfile))
	require.NoError(t, err)

	modIncludedDir := "../testdata/ignores-wildcard/basedir-modified-included-file"
	modIncludedResult, err := HashContext(modIncludedDir, filepath.Join(modIncludedDir, _dockerfile))
	require.NoError(t, err)

	assert.Equal(t, baselineResult, modIgnoredResult, "hash should not change when modifying ignored files")
	assert.NotEqual(t, baselineResult, modIncludedResult,
		"hash should change when modifying included (via wildcard ignore exclusion) files")
}

func TestHashIgnoresDockerfileOutsideDirMove(t *testing.T) {
	appDir := "./testdata/dockerfile-location-irrelevant/app"
	baseResult, err := HashContext(appDir, "./testdata/dockerfile-location-irrelevant/step1.Dockerfile")
	require.NoError(t, err)

	result, err := HashContext(appDir, "./testdata/dockerfile-location-irrelevant/step2.Dockerfile")
	require.NoError(t, err)

	assert.Equal(t, result, baseResult)
}

func TestHashRenamingMatters(t *testing.T) {
	step1Dir := "./testdata/filemode-matters/step1"
	baseResult, err := HashContext(step1Dir, filepath.Join(step1Dir, _dockerfile))
	require.NoError(t, err)

	step2Dir := "./testdata/renaming-matters/step2"
	result, err := HashContext(step2Dir, filepath.Join(step2Dir, _dockerfile))
	require.NoError(t, err)

	assert.NotEqual(t, result, baseResult)
}

func TestHashFilemodeMatters(t *testing.T) {
	step1Dir := "./testdata/filemode-matters/step1"
	baseResult, err := HashContext(step1Dir, filepath.Join(step1Dir, _dockerfile))
	require.NoError(t, err)

	step2Dir := "./testdata/filemode-matters/step2-chmod-x"
	result, err := HashContext(step2Dir, filepath.Join(step2Dir, _dockerfile))
	require.NoError(t, err)

	assert.NotEqual(t, result, baseResult)
}

func TestHashDeepSymlinks(t *testing.T) {
	dir := "./testdata/symlinks"
	_, err := HashContext(dir, filepath.Join(dir, "Dockerfile"))
	assert.NoError(t, err)
}

func TestIgnoreIrregularFiles(t *testing.T) {
	dir := t.TempDir()

	// Create a Dockerfile
	dockerfile := filepath.Join(dir, "Dockerfile")
	err := os.WriteFile(dockerfile, []byte{}, 0o600)
	require.NoError(t, err)

	// Create a pipe which should be ignored. (We will time out trying to read
	// it if it's not.)
	pipe := filepath.Join(dir, "pipe")
	err = syscall.Mkfifo(pipe, 0o666)
	require.NoError(t, err)
	// Confirm it's irregular.
	fi, err := os.Stat(pipe)
	require.NoError(t, err)
	assert.False(t, fi.Mode().IsRegular())

	_, err = HashContext(dir, dockerfile)
	assert.NoError(t, err)
}

func TestHashUnignoredDirs(t *testing.T) {
	step1Dir := "./testdata/unignores/basedir"
	baseResult, err := HashContext(step1Dir, filepath.Join(step1Dir, _dockerfile))
	require.NoError(t, err)

	step2Dir := "./testdata/unignores/basedir-with-unignored-files"
	unignoreResult, err := HashContext(step2Dir, filepath.Join(step2Dir, _dockerfile))
	require.NoError(t, err)

	assert.Equal(t, baseResult, unignoreResult)
}

func TestDockerIgnore(t *testing.T) {
	tests := []struct {
		name string

		dockerfile string
		context    string
		fs         map[string]string

		want    []string
		wantErr error
	}{
		{
			name:       "Dockerfile with root dockerignore",
			dockerfile: "./foo/Dockerfile",
			fs: map[string]string{
				".dockerignore": "rootignore",
			},
			want: []string{"rootignore"},
		},
		{
			name:       "Dockerfile with root dockerignore and custom dockerignore",
			dockerfile: "./foo/Dockerfile",
			fs: map[string]string{
				"foo/Dockerfile.dockerignore": "customignore",
				".dockerignore":               "rootignore",
			},
			want: []string{"customignore"},
		},
		{
			name:       "Dockerfile with root dockerignore and relative context",
			dockerfile: "./foo/Dockerfile",
			context:    "../",
			fs: map[string]string{
				"../.dockerignore": "rootignore",
			},
			want: []string{"rootignore"},
		},
		{
			name:       "Dockerfile without root dockerignore",
			dockerfile: "./foo/Dockerfile",
			want:       nil,
		},
		{
			name:       "Dockerfile with invalid root dockerignore",
			dockerfile: "./foo/Dockerfile",
			fs: map[string]string{
				".dockerignore": strings.Repeat("*", bufio.MaxScanTokenSize),
			},
			wantErr: bufio.ErrTooLong,
		},
		{
			name:       "custom.Dockerfile without custom dockerignore and without root dockerignore",
			dockerfile: "./foo/custom.Dockerfile",
			want:       nil,
		},
		{
			name:       "custom.Dockerfile with custom dockerignore and without root dockerignore",
			dockerfile: "./foo/custom.Dockerfile",
			fs: map[string]string{
				"foo/custom.Dockerfile.dockerignore": "customignore",
			},
			want: []string{"customignore"},
		},
		{
			name:       "custom.Dockerfile with custom dockerignore and with root dockerignore",
			dockerfile: "foo/custom.Dockerfile",
			fs: map[string]string{
				"foo/custom.Dockerfile.dockerignore": "customignore",
				".dockerignore":                      "rootignore",
			},
			want: []string{"customignore"},
		},
		{
			name:       "custom.Dockerfile without custom dockerignore and with root dockerignore",
			dockerfile: "foo/custom.Dockerfile",
			fs: map[string]string{
				".dockerignore": "rootignore",
			},
			want: []string{"rootignore"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := afero.NewMemMapFs()
			for fname, fdata := range tt.fs {
				f, err := fs.Create(fname)
				require.NoError(t, err)
				_, err = f.Write([]byte(fdata))
				require.NoError(t, err)
			}
			actual, err := GetIgnorePatterns(fs, tt.dockerfile, tt.context)

			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.want, actual)
		})
	}
}
