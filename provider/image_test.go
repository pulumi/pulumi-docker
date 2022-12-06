package provider

import (
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/mapper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegistryMapper(t *testing.T) {
	mapRegistry := func(pm resource.PropertyMap) (r Registry) {
		err := mapper.Map(pm.Mappable(), &r)
		require.NoError(t, err)
		return
	}

	t.Run("Valid Registry", func(t *testing.T) {
		expected := Registry{
			Server:   "https://index.docker.io/v1/",
			Username: "pulumipus",
			Password: "supersecret",
		}
		input := resource.PropertyMap{
			"server":   resource.NewStringProperty("https://index.docker.io/v1/"),
			"username": resource.NewStringProperty("pulumipus"),
			"password": resource.NewStringProperty("supersecret"),
		}
		actual := mapRegistry(input)
		assert.Equal(t, expected, actual)
	})

	t.Run("Incomplete Registry sets all available fields", func(t *testing.T) {
		expected := Registry{
			Server:   "https://index.docker.io/v1/",
			Username: "pulumipus",
		}
		input := resource.PropertyMap{
			"server":   resource.NewStringProperty("https://index.docker.io/v1/"),
			"username": resource.NewStringProperty("pulumipus"),
		}
		actual := mapRegistry(input)
		assert.Equal(t, expected, actual)
	})

	t.Run("Registry can be nil", func(t *testing.T) {
		expected := Registry{}
		input := resource.PropertyMap{
			"imageName": resource.NewStringProperty("foo"),
		}
		img := new(Image)
		err := img.unmarshalPropertyMap(input)
		require.NoError(t, err)
		actual := img.Registry
		assert.Equal(t, expected, actual)
	})
}

func TestMarshalBuildAndApplyDefaults(t *testing.T) {
	marshalBuildAndApplyDefaults := func(rawBuild resource.PropertyValue) (Build, error) {
		img := new(Image)
		input := resource.PropertyMap{
			"imageName": resource.NewStringProperty("foo"),
			"build":     rawBuild,
		}
		err := img.unmarshalPropertyMap(input)
		if err != nil {
			return Build{}, err
		}
		return *img.Build(), nil
	}

	t.Run("Default Build on empty input", func(t *testing.T) {
		expected := Build{
			Context:    ".",
			Dockerfile: "Dockerfile",
		}
		input := resource.NewObjectProperty(resource.PropertyMap{})
		actual, err := marshalBuildAndApplyDefaults(input)
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("String input Build", func(t *testing.T) {
		expected := Build{
			Context:    "/twilight/sparkle/bin",
			Dockerfile: "Dockerfile",
		}
		input := resource.NewStringProperty("/twilight/sparkle/bin")
		actual, err := marshalBuildAndApplyDefaults(input)
		assert.NoError(t, err)
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
		actual, err := marshalBuildAndApplyDefaults(input)
		assert.NoError(t, err)
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

		actual, err := marshalBuildAndApplyDefaults(input)
		assert.NoError(t, err)
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

		actual, err := marshalBuildAndApplyDefaults(input)
		assert.NoError(t, err)
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

		actual, err := marshalBuildAndApplyDefaults(input)
		assert.NoError(t, err)
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

		actual, err := marshalBuildAndApplyDefaults(input)
		assert.NoError(t, err)
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

		actual, err := marshalBuildAndApplyDefaults(input)
		assert.NoError(t, err)
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

		actual, err := marshalBuildAndApplyDefaults(input)
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
}

func TestMarshalArgs(t *testing.T) {
	marshalArgs := func(args resource.PropertyValue) map[string]*string {
		img := new(Image)
		input := resource.PropertyMap{
			"imageName": resource.NewStringProperty("foo"),
			"build": resource.NewObjectProperty(resource.PropertyMap{
				"args": args,
			}),
		}
		err := img.unmarshalPropertyMap(input)
		require.NoError(t, err)
		return img.Build().Args
	}

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

	t.Run("Returns empty when no args set", func(t *testing.T) {
		expected := map[string]*string{}
		input := resource.NewObjectProperty(resource.PropertyMap{})
		actual := marshalArgs(input)
		assert.Equal(t, expected, actual)
	})
}

func TestMarshalEnvs(t *testing.T) {
	marshalEnvs := func(args resource.PropertyValue) map[string]string {
		img := new(Image)
		input := resource.PropertyMap{
			"imageName": resource.NewStringProperty("foo"),
			"build": resource.NewObjectProperty(resource.PropertyMap{
				"env": args,
			}),
		}
		err := img.unmarshalPropertyMap(input)
		require.NoError(t, err)
		return img.Build().Env
	}

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

	t.Run("Returns empty when no environment variables set", func(t *testing.T) {
		expected := map[string]string{}
		input := resource.NewObjectProperty(resource.PropertyMap{})
		actual := marshalEnvs(input)
		assert.Equal(t, expected, actual)
	})
}

func TestMarshalCachedImages(t *testing.T) {
	marshalCachedImages := func(image Image, build resource.PropertyValue) []string {
		img := new(Image)
		*img = image
		input := resource.PropertyMap{
			"imageName": resource.NewStringProperty(img.ImageName),
			"build":     build,
		}
		err := img.unmarshalPropertyMap(input)
		require.NoError(t, err)
		return img.CacheImages()
	}

	t.Run("Test Cached Images", func(t *testing.T) {
		expected := []string{"apple", "banana", "cherry"}
		imgInput := Image{
			ImageName: "unicornsareawesome",
			SkipPush:  false,
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
			ImageName: "unicornsareawesome",
			SkipPush:  false,
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
			ImageName: "unicornsareawesome",
			SkipPush:  false,
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
