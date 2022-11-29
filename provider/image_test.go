package provider

import (
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetRegistry(t *testing.T) {

	expectedRegistry := Registry{
		Server:   "https://index.docker.io/v1/",
		Username: "pulumipus",
		Password: "supersecret",
	}

	expectedRegistryIncomplete := Registry{
		Server:   "https://index.docker.io/v1/",
		Username: "pulumipus",
	}

	expectedRegistryNil := Registry{}

	testRegistryValidInput := resource.PropertyValue{
		resource.NewPropertyMapFromMap(map[string]interface{}{
			"server":   "https://index.docker.io/v1/",
			"username": "pulumipus",
			"password": "supersecret",
		}),
	}

	testRegistryIncompleteInput := resource.PropertyValue{
		resource.NewPropertyMapFromMap(map[string]interface{}{
			"server":   "https://index.docker.io/v1/",
			"username": "pulumipus",
		}),
	}

	testRegistryNilInput := resource.PropertyValue{}

	testRegistryValid := setRegistry(testRegistryValidInput)
	testRegistryIncomplete := setRegistry(testRegistryIncompleteInput)
	testRegistryNil := setRegistry(testRegistryNilInput)

	assert.Equal(t, expectedRegistry, testRegistryValid)
	assert.NotEqual(t, expectedRegistry, testRegistryIncomplete)
	assert.Equal(t, expectedRegistryIncomplete, testRegistryIncomplete)
	assert.Equal(t, expectedRegistryNil, testRegistryNil)

}

func TestMarshalBuild(t *testing.T) {

	t.Run("Default Build on empty input", func(t *testing.T) {
		expected := Build{
			Context:    ".",
			Dockerfile: "Dockerfile",
			Args:       map[string]*string{},
			Env:        map[string]string{},
		}
		input := resource.PropertyValue{
			resource.NewPropertyMapFromMap(map[string]interface{}{}),
		}
		actual := marshalBuild(input)
		assert.Equal(t, expected, actual)
	})

	t.Run("String input Build", func(t *testing.T) {
		expected := Build{
			Context:    "/twilight/sparkle/bin",
			Dockerfile: "Dockerfile",
			Args:       map[string]*string{},
			Env:        map[string]string{},
		}
		input := resource.NewStringProperty("/twilight/sparkle/bin")
		actual := marshalBuild(input)
		assert.Equal(t, expected, actual)
	})

	t.Run("Custom Dockerfile with default context", func(t *testing.T) {
		expected := Build{
			Context:    ".",
			Dockerfile: "TheLastUnicorn",
			Args:       map[string]*string{},
			Env:        map[string]string{},
		}
		input := resource.PropertyValue{
			resource.NewPropertyMapFromMap(map[string]interface{}{
				"dockerfile": "TheLastUnicorn",
			}),
		}
		actual := marshalBuild(input)
		assert.Equal(t, expected, actual)
	})

	t.Run("Custom Dockerfile with custom context", func(t *testing.T) {
		expected := Build{
			Context:    "/twilight/sparkle/bin",
			Dockerfile: "TheLastUnicorn",
			Args:       map[string]*string{},
			Env:        map[string]string{},
		}
		input := resource.PropertyValue{
			resource.NewPropertyMapFromMap(map[string]interface{}{
				"dockerfile": "TheLastUnicorn",
				"context":    "/twilight/sparkle/bin",
			}),
		}
		actual := marshalBuild(input)
		assert.Equal(t, expected, actual)
	})

}
