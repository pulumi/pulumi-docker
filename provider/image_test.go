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

	assert.Equal(t, testRegistryValid, expectedRegistry)
	assert.NotEqual(t, testRegistryIncomplete, expectedRegistry)
	assert.Equal(t, testRegistryIncomplete, expectedRegistryIncomplete)
	assert.Equal(t, testRegistryNil, expectedRegistryNil)

}
