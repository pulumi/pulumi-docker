package provider

import (
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	rpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDiffUpdates(t *testing.T) {

	t.Run("No diff happens on changed password", func(t *testing.T) {
		expected := map[string]*rpc.PropertyDiff{}
		input := map[resource.PropertyKey]resource.ValueDiff{
			"registry": {
				Object: &resource.ObjectDiff{
					Updates: map[resource.PropertyKey]resource.ValueDiff{"password": {
						Old: resource.PropertyValue{
							V: "FancyToken",
						},
						New: resource.PropertyValue{
							V: "PedestrianPassword",
						},
						Array:  (*resource.ArrayDiff)(nil),
						Object: (*resource.ObjectDiff)(nil),
					}},
				},
			},
		}
		actual := diffUpdates(input)
		assert.Equal(t, expected, actual)
	})

	t.Run("No diff happens on changed username", func(t *testing.T) {
		expected := map[string]*rpc.PropertyDiff{}
		input := map[resource.PropertyKey]resource.ValueDiff{
			"registry": {
				Object: &resource.ObjectDiff{
					Updates: map[resource.PropertyKey]resource.ValueDiff{"username": {
						Old: resource.PropertyValue{
							V: "platypus",
						},
						New: resource.PropertyValue{
							V: "Schnabeltier",
						},
						Array:  (*resource.ArrayDiff)(nil),
						Object: (*resource.ObjectDiff)(nil),
					}},
				},
			},
		}
		actual := diffUpdates(input)
		assert.Equal(t, expected, actual)
	})

	t.Run("Diff happens on changed server name", func(t *testing.T) {
		expected := map[string]*rpc.PropertyDiff{
			"registry": {
				Kind: rpc.PropertyDiff_UPDATE,
			},
		}
		input := map[resource.PropertyKey]resource.ValueDiff{
			"registry": {
				Object: &resource.ObjectDiff{
					Updates: map[resource.PropertyKey]resource.ValueDiff{"server": {
						Old: resource.PropertyValue{
							V: "dockerhub",
						},
						New: resource.PropertyValue{
							V: "ShinyPrivateGHCR",
						},
						Array:  (*resource.ArrayDiff)(nil),
						Object: (*resource.ObjectDiff)(nil),
					}},
				},
			},
		}
		actual := diffUpdates(input)
		assert.Equal(t, expected, actual)
	})

	t.Run("Diff happens on changed build context", func(t *testing.T) {
		expected := map[string]*rpc.PropertyDiff{
			"build": {
				Kind: rpc.PropertyDiff_UPDATE,
			},
		}
		input := map[resource.PropertyKey]resource.ValueDiff{
			"build": {
				Object: &resource.ObjectDiff{
					Updates: map[resource.PropertyKey]resource.ValueDiff{"contextDigest": {
						Old: resource.PropertyValue{
							V: "12345",
						},
						New: resource.PropertyValue{
							V: "54321",
						},
						Array:  (*resource.ArrayDiff)(nil),
						Object: (*resource.ObjectDiff)(nil),
					}},
				},
			},
		}
		actual := diffUpdates(input)
		assert.Equal(t, expected, actual)
	})
}

func TestHashIgnoresFile(t *testing.T) {
	step1Dir := "./testdata/ignores/basedir"
	baseResult, err := hashContext(step1Dir, filepath.Join(step1Dir, defaultDockerfile))
	require.NoError(t, err)

	step2Dir := "./testdata/ignores/basedir-with-ignored-files"
	result, err := hashContext(step2Dir, filepath.Join(step2Dir, defaultDockerfile))
	require.NoError(t, err)

	assert.Equal(t, result, baseResult)
}

func TestHashIgnoresDockerfileOutsideDirMove(t *testing.T) {
	appDir := "./testdata/dockerfile-location-irrelevant/app"
	baseResult, err := hashContext(appDir, "./testdata/dockerfile-location-irrelevant/step1.Dockerfile")
	require.NoError(t, err)

	result, err := hashContext(appDir, "./testdata/dockerfile-location-irrelevant/step2.Dockerfile")
	require.NoError(t, err)

	assert.Equal(t, result, baseResult)
}

func TestHashRenamingMatters(t *testing.T) {
	step1Dir := "./testdata/filemode-matters/step1"
	baseResult, err := hashContext(step1Dir, filepath.Join(step1Dir, defaultDockerfile))
	require.NoError(t, err)

	step2Dir := "./testdata/renaming-matters/step2"
	result, err := hashContext(step2Dir, filepath.Join(step2Dir, defaultDockerfile))
	require.NoError(t, err)

	assert.NotEqual(t, result, baseResult)
}

func TestHashFilemodeMatters(t *testing.T) {
	step1Dir := "./testdata/filemode-matters/step1"
	baseResult, err := hashContext(step1Dir, filepath.Join(step1Dir, defaultDockerfile))
	require.NoError(t, err)

	step2Dir := "./testdata/filemode-matters/step2-chmod-x"
	result, err := hashContext(step2Dir, filepath.Join(step2Dir, defaultDockerfile))
	require.NoError(t, err)

	assert.NotEqual(t, result, baseResult)
}

func TestHashDeepSymlinks(t *testing.T) {
	dir := "./testdata/symlinks"
	_, err := hashContext(dir, filepath.Join(dir, "Dockerfile"))
	assert.NoError(t, err)

}

func TestHashUnignoredDirs(t *testing.T) {
	step1Dir := "./testdata/unignores/basedir"
	baseResult, err := hashContext(step1Dir, filepath.Join(step1Dir, defaultDockerfile))
	require.NoError(t, err)

	step2Dir := "./testdata/unignores/basedir-with-unignored-dirs"
	unignoreResult, err := hashContext(step2Dir, filepath.Join(step2Dir, defaultDockerfile))
	require.NoError(t, err)

	assert.Equal(t, baseResult, unignoreResult)
}
