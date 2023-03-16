package provider

import (
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

<<<<<<< HEAD
func TestHashIgnoresFile(t *testing.T) {
	baseResult, err := hashContext("./testdata/ignores/basedir", "./Dockerfile")
	require.NoError(t, err)

	result, err := hashContext("./testdata/ignores/basedir-with-ignored-files", "./Dockerfile")
	require.NoError(t, err)

	assert.Equal(t, result, baseResult)
}

func TestHashRenamingMatters(t *testing.T) {
	baseResult, err := hashContext("./testdata/renaming-matters/step1", "./Dockerfile")
	require.NoError(t, err)

	result, err := hashContext("./testdata/renaming-matters/step2", "./Dockerfile")
	require.NoError(t, err)

	assert.NotEqual(t, result, baseResult)
}

func TestHashFilemodeMatters(t *testing.T) {
	baseResult, err := hashContext("./testdata/filemode-matters/step1", "./Dockerfile")
	require.NoError(t, err)

	result, err := hashContext("./testdata/filemode-matters/step2-chmod-x", "./Dockerfile")
	require.NoError(t, err)

	assert.NotEqual(t, result, baseResult)
}

func TestHashDeepSymlinks(t *testing.T) {
	_, err := hashContext("./testdata/symlinks", "./Dockerfile")
	assert.NoError(t, err)

}

func TestGetRelDockerfilePath(t *testing.T) {

	t.Run("A Dockerfile name with no separators is relative to the build context", func(t *testing.T) {
		expected := "Dockerfile"
		input1, input2 := ".", "Dockerfile"

		actual, err := getRelDockerfilePath(input1, input2)
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("A Dockerfile name with separators will return its relative position to the build context", func(t *testing.T) {
		expected := "../Dockerfile"
		input1, input2 := "./special-context", "./Dockerfile"

		actual, err := getRelDockerfilePath(input1, input2)
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("A Dockerfile name with multiple separators will return its relative position to the build context", func(t *testing.T) {
		expected := "../other-folder/Dockerfile"
		input1, input2 := "./special-context", "./other-folder/Dockerfile"

		actual, err := getRelDockerfilePath(input1, input2)
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
}
