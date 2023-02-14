package provider

import (
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	rpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
	"github.com/stretchr/testify/assert"
	"testing"
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
