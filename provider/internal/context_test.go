package internal

import (
	"os"
	"path/filepath"
	"syscall"
	"testing"

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

// Tests that we handle .dockerignore exclusions such as "!foo/*/bar", as above, when using a
// relative context path.
func TestHashIgnoresWildcardsRelative(t *testing.T) {
	err := os.Chdir("mock")
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
