package provider

import (
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
		input := resource.NewObjectProperty(resource.PropertyMap{
			"dockerfile": resource.NewStringProperty("TheLastUnicorn"),
		})
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
		input := resource.NewObjectProperty(resource.PropertyMap{
			"dockerfile": resource.NewStringProperty("TheLastUnicorn"),
			"context":    resource.NewStringProperty("/twilight/sparkle/bin"),
		})

		actual := marshalBuild(input)
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
			Env: map[string]string{},
		}

		input := resource.NewObjectProperty(resource.PropertyMap{
			"args": resource.NewObjectProperty(resource.PropertyMap{
				"Swiftwind": resource.NewStringProperty("Alicorn"),
			}),
		})

		actual := marshalBuild(input)
		assert.Equal(t, expected, actual)
	})

	t.Run("Setting Env", func(t *testing.T) {

		expected := Build{
			Context:    ".",
			Dockerfile: "Dockerfile",
			Args:       map[string]*string{},
			Env: map[string]string{
				"Strawberry": "fruit",
			},
		}

		input := resource.NewObjectProperty(resource.PropertyMap{
			"env": resource.NewObjectProperty(resource.PropertyMap{
				"Strawberry": resource.NewStringProperty("fruit"),
			}),
		})

		actual := marshalBuild(input)
		assert.Equal(t, expected, actual)
	})

}

func TestSetArgs(t *testing.T) {
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
		actual := setArgs(input)
		assert.Equal(t, expected, actual)
	})
}

func TestSetEnvs(t *testing.T) {
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
		actual := setEnvs(input)
		assert.Equal(t, expected, actual)
	})
}
