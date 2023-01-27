package provider

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetRegistry(t *testing.T) {

	t.Run("Valid Registry", func(t *testing.T) {
		expected := Registry{
			Server:   "https://index.docker.io/v1/",
			Username: "pulumipus",
			Password: "supersecret",
		}
		input := resource.NewObjectProperty(resource.PropertyMap{
			"server":   resource.NewStringProperty("https://index.docker.io/v1/"),
			"username": resource.NewStringProperty("pulumipus"),
			"password": resource.NewStringProperty("supersecret"),
		})

		actual := marshalRegistry(input)
		assert.Equal(t, expected, actual)
	})
	t.Run("Incomplete Registry sets all available fields", func(t *testing.T) {
		expected := Registry{
			Server:   "https://index.docker.io/v1/",
			Username: "pulumipus",
		}
		input := resource.NewObjectProperty(resource.PropertyMap{
			"server":   resource.NewStringProperty("https://index.docker.io/v1/"),
			"username": resource.NewStringProperty("pulumipus"),
			"password": resource.NewStringProperty(""),
		})

		actual := marshalRegistry(input)
		assert.Equal(t, expected, actual)
	})

	t.Run("Registry can be nil", func(t *testing.T) {
		expected := Registry{}
		input := resource.PropertyValue{}
		actual := marshalRegistry(input)
		assert.Equal(t, expected, actual)
	})
}

func TestMarshalBuildAndApplyDefaults(t *testing.T) {

	t.Run("Default Build on empty input", func(t *testing.T) {
		expected := Build{
			Context:        ".",
			Dockerfile:     "Dockerfile",
			BuilderVersion: "2",
		}
		input := resource.NewObjectProperty(resource.PropertyMap{})
		actual, err := marshalBuildAndApplyDefaults(input)
		assert.Equal(t, expected, actual)
		assert.NoError(t, err)
	})

	t.Run("String input Build", func(t *testing.T) {
		expected := Build{
			Context:    "/twilight/sparkle/bin",
			Dockerfile: "Dockerfile",
		}
		input := resource.NewStringProperty("/twilight/sparkle/bin")
		actual, err := marshalBuildAndApplyDefaults(input)
		assert.Equal(t, expected, actual)
		assert.NoError(t, err)
	})

	t.Run("Custom Dockerfile with default context", func(t *testing.T) {
		expected := Build{
			Context:        ".",
			Dockerfile:     "TheLastUnicorn",
			BuilderVersion: "2",
		}
		input := resource.NewObjectProperty(resource.PropertyMap{
			"dockerfile": resource.NewStringProperty("TheLastUnicorn"),
		})
		actual, err := marshalBuildAndApplyDefaults(input)
		assert.Equal(t, expected, actual)
		assert.NoError(t, err)
	})

	t.Run("Custom Dockerfile with custom context", func(t *testing.T) {
		expected := Build{
			Context:        "/twilight/sparkle/bin",
			Dockerfile:     "TheLastUnicorn",
			BuilderVersion: "2",
		}
		input := resource.NewObjectProperty(resource.PropertyMap{
			"dockerfile": resource.NewStringProperty("TheLastUnicorn"),
			"context":    resource.NewStringProperty("/twilight/sparkle/bin"),
		})

		actual, err := marshalBuildAndApplyDefaults(input)
		assert.Equal(t, expected, actual)
		assert.NoError(t, err)
	})

	t.Run("Setting Args", func(t *testing.T) {
		argval := "Alicorn"
		expected := Build{
			Context:    ".",
			Dockerfile: "Dockerfile",
			Args: map[string]*string{
				"Swiftwind": &argval,
			},
			BuilderVersion: "2",
		}

		input := resource.NewObjectProperty(resource.PropertyMap{
			"args": resource.NewObjectProperty(resource.PropertyMap{
				"Swiftwind": resource.NewStringProperty("Alicorn"),
			}),
		})

		actual, err := marshalBuildAndApplyDefaults(input)
		assert.Equal(t, expected, actual)
		assert.NoError(t, err)
	})

	t.Run("Sets Extra Options", func(t *testing.T) {
		expected := Build{
			Context:        ".",
			Dockerfile:     "Dockerfile",
			ExtraOptions:   []string{"cat", "dog", "pot-bellied pig"},
			BuilderVersion: "2",
		}

		input := resource.NewObjectProperty(resource.PropertyMap{
			"extraOptions": resource.NewArrayProperty([]resource.PropertyValue{
				resource.NewStringProperty("cat"),
				resource.NewStringProperty("dog"),
				resource.NewStringProperty("pot-bellied pig"),
			}),
		})

		actual, err := marshalBuildAndApplyDefaults(input)
		assert.Equal(t, expected, actual)
		assert.NoError(t, err)
	})

	t.Run("Does Not Set Extra Options on Empty Input", func(t *testing.T) {
		expected := Build{
			Context:        ".",
			Dockerfile:     "Dockerfile",
			BuilderVersion: "2",
		}

		input := resource.NewObjectProperty(resource.PropertyMap{
			"extraOptions": resource.NewArrayProperty([]resource.PropertyValue{}),
		})

		actual, err := marshalBuildAndApplyDefaults(input)
		assert.Equal(t, expected, actual)
		assert.NoError(t, err)
	})
	t.Run("Sets Target", func(t *testing.T) {
		expected := Build{
			Context:        ".",
			Dockerfile:     "Dockerfile",
			Target:         "bullseye",
			BuilderVersion: "2",
		}

		input := resource.NewObjectProperty(resource.PropertyMap{
			"target": resource.NewStringProperty("bullseye"),
		})

		actual, err := marshalBuildAndApplyDefaults(input)
		assert.Equal(t, expected, actual)
		assert.NoError(t, err)
	})
	t.Run("Sets Builder to classic V1 builder", func(t *testing.T) {
		expected := Build{
			Context:        ".",
			Dockerfile:     "Dockerfile",
			BuilderVersion: "1",
		}

		input := resource.NewObjectProperty(resource.PropertyMap{
			"builderVersion": resource.NewStringProperty("BuilderV1"),
		})

		actual, err := marshalBuildAndApplyDefaults(input)
		assert.Equal(t, expected, actual)
		assert.NoError(t, err)
	})
	t.Run("Sets Builder to default on no input", func(t *testing.T) {
		expected := Build{
			Context:        ".",
			Dockerfile:     "Dockerfile",
			BuilderVersion: "2",
		}

		input := resource.NewObjectProperty(resource.PropertyMap{})

		actual, err := marshalBuildAndApplyDefaults(input)
		assert.Equal(t, expected, actual)
		assert.NoError(t, err)
	})
}

func TestMarshalArgs(t *testing.T) {
	t.Run("Set any args", func(t *testing.T) {
		a := "Alicorn"
		p := "Pegasus"
		tl := "Unicorn"
		expected := map[string]*string{
			"Swiftwind": &a,
			"Fledge":    &p,
			"The Last":  &tl,
		}
		input := resource.NewObjectProperty(resource.PropertyMap{
			"Swiftwind": resource.NewStringProperty("Alicorn"),
			"Fledge":    resource.NewStringProperty("Pegasus"),
			"The Last":  resource.NewStringProperty("Unicorn"),
		})
		actual := marshalArgs(input)
		assert.Equal(t, expected, actual)
	})

	t.Run("Returns nil when no args set", func(t *testing.T) {
		expected := map[string]*string(nil)
		input := resource.NewObjectProperty(resource.PropertyMap{})
		actual := marshalArgs(input)
		assert.Equal(t, expected, actual)
	})
}

func TestMarshalCachedImages(t *testing.T) {
	t.Run("Test Cached Images", func(t *testing.T) {
		expected := []string{"apple", "banana", "cherry"}
		imgInput := Image{
			Name:     "unicornsareawesome",
			SkipPush: false,
			Registry: Registry{
				Server:   "https://index.docker.io/v1/",
				Username: "pulumipus",
				Password: "supersecret",
			},
		}
		buildInput := resource.NewObjectProperty(resource.PropertyMap{
			"dockerfile": resource.NewStringProperty("TheLastUnicorn"),
			"context":    resource.NewStringProperty("/twilight/sparkle/bin"),

			"cacheFrom": resource.NewObjectProperty(resource.PropertyMap{
				"images": resource.NewArrayProperty([]resource.PropertyValue{
					resource.NewStringProperty("apple"),
					resource.NewStringProperty("banana"),
					resource.NewStringProperty("cherry"),
				}),
			}),
		})

		actual := marshalCachedImages(imgInput, buildInput)
		assert.Equal(t, expected, actual)

	})
	t.Run("Test Cached Images No Build Input Returns Nil", func(t *testing.T) {
		expected := []string(nil)
		imgInput := Image{
			Name:     "unicornsareawesome",
			SkipPush: false,
			Registry: Registry{
				Server:   "https://index.docker.io/v1/",
				Username: "pulumipus",
				Password: "supersecret",
			},
		}
		buildInput := resource.NewObjectProperty(resource.PropertyMap{})
		actual := marshalCachedImages(imgInput, buildInput)
		assert.Equal(t, expected, actual)
	})

	t.Run("Test Cached Images No cacheFrom Input Returns Nil", func(t *testing.T) {
		expected := []string(nil)
		imgInput := Image{
			Name:     "unicornsareawesome",
			SkipPush: false,
			Registry: Registry{
				Server:   "https://index.docker.io/v1/",
				Username: "pulumipus",
				Password: "supersecret",
			},
		}
		buildInput := resource.NewObjectProperty(resource.PropertyMap{
			"dockerfile": resource.NewStringProperty("TheLastUnicorn"),
			"context":    resource.NewStringProperty("/twilight/sparkle/bin"),
		})
		actual := marshalCachedImages(imgInput, buildInput)
		assert.Equal(t, expected, actual)
	})
}

func TestMarshalBuilder(t *testing.T) {
	t.Run("Test Builder Version Default", func(t *testing.T) {
		expected := types.BuilderBuildKit
		input := resource.NewPropertyValue(nil)
		actual, err := marshalBuilder(input)
		assert.Equal(t, expected, actual)
		assert.NoError(t, err)

	})
	t.Run("Test Builder BuildKit Version", func(t *testing.T) {
		expected := types.BuilderBuildKit
		input := resource.NewStringProperty("BuilderBuildKit")

		actual, err := marshalBuilder(input)
		assert.Equal(t, expected, actual)
		assert.NoError(t, err)

	})
	t.Run("Test Builder V1 Version", func(t *testing.T) {
		expected := types.BuilderV1
		input := resource.NewStringProperty("BuilderV1")

		actual, err := marshalBuilder(input)
		assert.Equal(t, expected, actual)
		assert.NoError(t, err)

	})
	t.Run("Test Invalid Builder Returns Error", func(t *testing.T) {
		expected := types.BuilderV1
		input := resource.NewStringProperty("BuilderV1")

		actual, err := marshalBuilder(input)
		assert.Equal(t, expected, actual)
		assert.NoError(t, err)
	})
}

func TestMarshalSkipPush(t *testing.T) {
	t.Run("Test SkipPush defaults to false", func(t *testing.T) {
		expected := false
		input := resource.NewPropertyValue(nil)
		actual := marshalSkipPush(input)
		assert.Equal(t, expected, actual)

	})
	t.Run("Test SkipPush returns true if set to true", func(t *testing.T) {
		expected := true
		input := resource.NewBoolProperty(true)

		actual := marshalSkipPush(input)
		assert.Equal(t, expected, actual)
	})
	t.Run("Test SkipPush returns false if set to false", func(t *testing.T) {
		expected := false
		input := resource.NewBoolProperty(false)

		actual := marshalSkipPush(input)
		assert.Equal(t, expected, actual)
	})
}

func TestGetRegistryAddrFromImage(t *testing.T) {

	t.Run("Returns registry name of correct spec format", func(t *testing.T) {
		expected := "pulumi.test.registry"
		input := "pulumi.test.registry/unicorns/swiftwind:latest"
		actual, err := getRegistryAddrFromImage(input)
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("Returns error for incorrect registry format", func(t *testing.T) {
		expected := ""
		input := "pulumi-test-registry/unicorns/swiftwind:latest"

		expectedError := fmt.Errorf(
			"error: repository name must be canonical. This provider requires " +
				"all image names to be fully qualified.\nFor example, if you are " +
				"attempting to push to Dockerhub, prefix your image name with " +
				"`docker.io`:\n\n`docker.io/repository/image:tag`",
		)
		actual, err := getRegistryAddrFromImage(input)
		assert.Equal(t, expected, actual)
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
}
