package provider

import (
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

		actual := setRegistry(input)
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
		})

		actual := setRegistry(input)
		assert.Equal(t, expected, actual)
	})

	t.Run("Registry can be nil", func(t *testing.T) {
		expected := Registry{}
		input := resource.PropertyValue{}
		actual := setRegistry(input)
		assert.Equal(t, expected, actual)
	})
}

func TestMarshalBuildAndApplyDefaults(t *testing.T) {

	t.Run("Default Build on empty input", func(t *testing.T) {
		expected := Build{
			Context:    ".",
			Dockerfile: "Dockerfile",
		}
		input := resource.NewObjectProperty(resource.PropertyMap{})
		actual, _ := marshalBuildAndApplyDefaults(input)
		assert.Equal(t, expected, actual)
	})

	t.Run("String input Build", func(t *testing.T) {
		expected := Build{
			Context:    "/twilight/sparkle/bin",
			Dockerfile: "Dockerfile",
		}
		input := resource.NewStringProperty("/twilight/sparkle/bin")
		actual, _ := marshalBuildAndApplyDefaults(input)
		assert.Equal(t, expected, actual)
	})

	t.Run("Custom Dockerfile with default context", func(t *testing.T) {
		expected := Build{
			Context:    ".",
			Dockerfile: "TheLastUnicorn",
		}
		input := resource.NewObjectProperty(resource.PropertyMap{
			"dockerfile": resource.NewStringProperty("TheLastUnicorn"),
		})
		actual, _ := marshalBuildAndApplyDefaults(input)
		assert.Equal(t, expected, actual)
	})

	t.Run("Custom Dockerfile with custom context", func(t *testing.T) {
		expected := Build{
			Context:    "/twilight/sparkle/bin",
			Dockerfile: "TheLastUnicorn",
		}
		input := resource.NewObjectProperty(resource.PropertyMap{
			"dockerfile": resource.NewStringProperty("TheLastUnicorn"),
			"context":    resource.NewStringProperty("/twilight/sparkle/bin"),
		})

		actual, _ := marshalBuildAndApplyDefaults(input)
		assert.Equal(t, expected, actual)
	})

	t.Run("Setting Args", func(t *testing.T) {
		argval := "Alicorn"
		expected := Build{
			Context:    ".",
			Dockerfile: "Dockerfile",
			Args: map[string]*string{
				"Swiftwind": &argval,
			},
		}

		input := resource.NewObjectProperty(resource.PropertyMap{
			"args": resource.NewObjectProperty(resource.PropertyMap{
				"Swiftwind": resource.NewStringProperty("Alicorn"),
			}),
		})

		actual, _ := marshalBuildAndApplyDefaults(input)
		assert.Equal(t, expected, actual)
	})

	t.Run("Setting Env", func(t *testing.T) {

		expected := Build{
			Context:    ".",
			Dockerfile: "Dockerfile",
			Env: map[string]string{
				"Strawberry": "fruit",
			},
		}

		input := resource.NewObjectProperty(resource.PropertyMap{
			"env": resource.NewObjectProperty(resource.PropertyMap{
				"Strawberry": resource.NewStringProperty("fruit"),
			}),
		})

		actual, _ := marshalBuildAndApplyDefaults(input)
		assert.Equal(t, expected, actual)
	})

	t.Run("Sets Extra Options", func(t *testing.T) {
		expected := Build{
			Context:      ".",
			Dockerfile:   "Dockerfile",
			ExtraOptions: []string{"cat", "dog", "pot-bellied pig"},
		}

		input := resource.NewObjectProperty(resource.PropertyMap{
			"extraOptions": resource.NewArrayProperty([]resource.PropertyValue{
				resource.NewStringProperty("cat"),
				resource.NewStringProperty("dog"),
				resource.NewStringProperty("pot-bellied pig"),
			}),
		})

		actual, _ := marshalBuildAndApplyDefaults(input)
		assert.Equal(t, expected, actual)
	})

	t.Run("Does Not Set Extra Options on Empty Input", func(t *testing.T) {
		expected := Build{
			Context:    ".",
			Dockerfile: "Dockerfile",
		}

		input := resource.NewObjectProperty(resource.PropertyMap{
			"extraOptions": resource.NewArrayProperty([]resource.PropertyValue{}),
		})

		actual, _ := marshalBuildAndApplyDefaults(input)
		assert.Equal(t, expected, actual)
	})
	t.Run("Sets Target", func(t *testing.T) {
		expected := Build{
			Context:    ".",
			Dockerfile: "Dockerfile",
			Target:     "bullseye",
		}

		input := resource.NewObjectProperty(resource.PropertyMap{
			"target": resource.NewStringProperty("bullseye"),
		})

		actual, _ := marshalBuildAndApplyDefaults(input)
		assert.Equal(t, expected, actual)
	})
	t.Run("Sets Target", func(t *testing.T) {
		expected := Build{
			Context:    ".",
			Dockerfile: "Dockerfile",
			Target:     "bullseye",
		}

		input := resource.NewObjectProperty(resource.PropertyMap{
			"target": resource.NewStringProperty("bullseye"),
		})

		actual, _ := marshalBuildAndApplyDefaults(input)
		assert.Equal(t, expected, actual)
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

func TestMarshalEnvs(t *testing.T) {
	t.Run("Set any environment variables", func(t *testing.T) {
		expected := map[string]string{
			"Strawberry": "fruit",
			"Carrot":     "veggie",
			"Docker":     "a bit of a mess tbh",
		}
		input := resource.NewObjectProperty(resource.PropertyMap{
			"Strawberry": resource.NewStringProperty("fruit"),
			"Carrot":     resource.NewStringProperty("veggie"),
			"Docker":     resource.NewStringProperty("a bit of a mess tbh"),
		})
		actual := marshalEnvs(input)
		assert.Equal(t, expected, actual)
	})
	t.Run("Returns nil when no environment variables set", func(t *testing.T) {
		expected := map[string]string(nil)
		input := resource.NewObjectProperty(resource.PropertyMap{})
		actual := marshalEnvs(input)
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
				"stages": resource.NewArrayProperty([]resource.PropertyValue{
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
		actual, _ := marshalBuilder(input)
		assert.Equal(t, expected, actual)

	})
	t.Run("Test Builder BuildKit Version", func(t *testing.T) {
		expected := types.BuilderBuildKit
		input := resource.NewStringProperty("BuilderBuildKit")

		actual, _ := marshalBuilder(input)
		assert.Equal(t, expected, actual)

	})
	t.Run("Test Builder V1 Version", func(t *testing.T) {
		expected := types.BuilderV1
		input := resource.NewStringProperty("BuilderV1")

		actual, _ := marshalBuilder(input)
		assert.Equal(t, expected, actual)

	})
	t.Run("Test Invalid Builder Returns Error", func(t *testing.T) {
		expected := types.BuilderV1
		input := resource.NewStringProperty("BuilderV1")

		actual, _ := marshalBuilder(input)
		assert.Equal(t, expected, actual)

	})
}
