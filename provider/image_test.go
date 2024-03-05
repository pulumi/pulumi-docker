package provider

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/stretchr/testify/assert"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
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

	t.Run("Unknown Registry Server", func(t *testing.T) {
		expected := Registry{
			Username: "pulumipus",
			Password: "supersecret",
		}
		input := resource.NewObjectProperty(resource.PropertyMap{
			"server":   resource.MakeComputed(resource.NewStringProperty("X")),
			"username": resource.NewStringProperty("pulumipus"),
			"password": resource.NewStringProperty("supersecret"),
		})

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

	t.Run("No default Dockerfile and Context for Unknown input",
		func(t *testing.T) {
			expected := Build{
				BuilderVersion: "2",
			}
			input := resource.NewObjectProperty(resource.PropertyMap{
				"dockerfile": resource.MakeComputed(resource.NewStringProperty("dockerfile-from-elsewhere")),
				"context":    resource.MakeComputed(resource.NewStringProperty("context-is-computed-at-up-time")),
			})
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

	t.Run("Sets Args", func(t *testing.T) {
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

	t.Run("Handles Unknown Args", func(t *testing.T) {
		argval := "rainbow-mane"
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
				"Swiftwind":  resource.NewStringProperty("rainbow-mane"),
				"Fluttershy": resource.MakeComputed(resource.NewStringProperty("pink-hair")),
			}),
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

	t.Run("Handles Unknown Target", func(t *testing.T) {
		expected := Build{
			Context:        ".",
			Dockerfile:     "Dockerfile",
			BuilderVersion: "2",
		}

		input := resource.NewObjectProperty(resource.PropertyMap{
			"target": resource.MakeComputed(resource.NewStringProperty("moving-target")),
		})

		actual, err := marshalBuildAndApplyDefaults(input)
		assert.Equal(t, expected, actual)
		assert.NoError(t, err)
	})

	t.Run("Sets Platform", func(t *testing.T) {
		expected := Build{
			Context:        ".",
			Dockerfile:     "Dockerfile",
			Platform:       "linux/leg32",
			BuilderVersion: "2",
		}

		input := resource.NewObjectProperty(resource.PropertyMap{
			"platform": resource.NewStringProperty("linux/leg32"),
		})

		actual, err := marshalBuildAndApplyDefaults(input)
		assert.Equal(t, expected, actual)
		assert.NoError(t, err)
	})

	t.Run("Handles Unknown Platform", func(t *testing.T) {
		expected := Build{
			Context:        ".",
			Dockerfile:     "Dockerfile",
			BuilderVersion: "2",
		}

		input := resource.NewObjectProperty(resource.PropertyMap{
			"platform": resource.MakeComputed(resource.NewStringProperty("wheres-my-train")),
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
		a := "alicorn"
		p := "Pegasus"
		tl := "Unicorn"
		expected := map[string]*string{
			"Swiftwind": &a,
			"Fledge":    &p,
			"The Last":  &tl,
		}
		input := resource.NewObjectProperty(resource.PropertyMap{
			"Swiftwind": resource.NewStringProperty("alicorn"),
			"Fledge":    resource.NewStringProperty("Pegasus"),
			"The Last":  resource.NewStringProperty("Unicorn"),
		})
		actual := marshalArgs(input)
		assert.Equal(t, expected, actual)
	})
	t.Run("Does not set Computed args", func(t *testing.T) {
		a := "unicorn-with-wings"

		expected := map[string]*string{
			"Swiftwind": &a,
		}
		input := resource.NewObjectProperty(resource.PropertyMap{
			"Swiftwind": resource.NewStringProperty("unicorn-with-wings"),
			"Fledge":    resource.MakeComputed(resource.NewStringProperty("pegasus")),
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
		actual, err := marshalCachedImages(buildInput)
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
	t.Run("Test Cached Images No Build Input Returns Nil", func(t *testing.T) {
		expected := []string(nil)
		buildInput := resource.NewObjectProperty(resource.PropertyMap{})
		actual, err := marshalCachedImages(buildInput)
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("Test Cached Images No cacheFrom Input Returns Nil", func(t *testing.T) {
		expected := []string(nil)
		buildInput := resource.NewObjectProperty(resource.PropertyMap{
			"dockerfile": resource.NewStringProperty("TheLastUnicorn"),
			"context":    resource.NewStringProperty("/twilight/sparkle/bin"),
		})
		actual, err := marshalCachedImages(buildInput)
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
	t.Run("Test Cached Images Non-array Images Returns Nil and Error", func(t *testing.T) {
		expected := []string(nil)
		buildInput := resource.NewObjectProperty(resource.PropertyMap{
			"dockerfile": resource.NewStringProperty("TheLastUnicorn"),
			"context":    resource.NewStringProperty("/twilight/sparkle/bin"),
			"cacheFrom": resource.NewObjectProperty(resource.PropertyMap{
				"images": resource.NewStringProperty("Shadowfax"),
			}),
		})
		actual, err := marshalCachedImages(buildInput)
		expectedError := fmt.Errorf("the `images` field must be a list of strings")
		if assert.Error(t, err) {
			assert.Equal(t, expectedError, err)
		}
		assert.Equal(t, expected, actual)
		assert.Nil(t, actual)
	})
	t.Run("Test Cached Images No images Input Returns Nil and error", func(t *testing.T) {
		buildInput := resource.NewObjectProperty(resource.PropertyMap{
			"dockerfile": resource.NewStringProperty("TheLastUnicorn"),
			"context":    resource.NewStringProperty("/twilight/sparkle/bin"),
			"cacheFrom":  resource.NewObjectProperty(resource.PropertyMap{}),
		})
		actual, err := marshalCachedImages(buildInput)
		expectedError := fmt.Errorf("cacheFrom requires an `images` field")
		if assert.Error(t, err) {
			assert.Equal(t, expectedError, err)
		}
		assert.Nil(t, actual)
	})

	t.Run("Test Cached Images Passes On Unknowns", func(t *testing.T) {
		expected := []string(nil)
		buildInput := resource.NewObjectProperty(resource.PropertyMap{
			"cacheFrom": resource.NewObjectProperty(resource.PropertyMap{
				"images": resource.NewArrayProperty([]resource.PropertyValue{
					resource.MakeComputed(resource.NewStringProperty("looking-for-my-image")),
				}),
			}),
		})
		actual, err := marshalCachedImages(buildInput)
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
	t.Run("Test Cached Images For Preview Passes On Unknowns And Keeps Knowns", func(t *testing.T) {
		expected := []string{"apple", "banana", "cherry"}
		buildInput := resource.NewObjectProperty(resource.PropertyMap{
			"cacheFrom": resource.NewObjectProperty(resource.PropertyMap{
				"images": resource.NewArrayProperty([]resource.PropertyValue{
					resource.NewNullProperty(),
					resource.MakeComputed(resource.NewStringProperty("looking-for-my-image")),
					resource.NewStringProperty("apple"),
					resource.NewStringProperty("banana"),
					resource.NewStringProperty("cherry"),
				}),
			}),
		})
		actual, err := marshalCachedImages(buildInput)
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
	t.Run("Test Cached Images Passes On Unknown Images List", func(t *testing.T) {
		expected := []string(nil)
		buildInput := resource.NewObjectProperty(resource.PropertyMap{
			"cacheFrom": resource.NewObjectProperty(resource.PropertyMap{
				"images": resource.NewComputedProperty(
					resource.Computed{
						Element: resource.NewArrayProperty([]resource.PropertyValue{
							resource.MakeComputed(resource.NewStringProperty("looking-for-my-image")),
						}),
					},
				),
			}),
		})
		actual, err := marshalCachedImages(buildInput)
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
	t.Run("Test Cached Images Passes On Unknown cacheFrom", func(t *testing.T) {
		expected := []string(nil)
		buildInput := resource.NewObjectProperty(resource.PropertyMap{
			"cacheFrom": resource.NewComputedProperty(
				resource.Computed{Element: resource.NewObjectProperty(
					resource.NewPropertyMapFromMap(map[string]interface{}{}),
				)},
			),
		})
		actual, err := marshalCachedImages(buildInput)
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
}

// TODO: do we want to allow Builder to be Unknown? there's very little use case here
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

		expectedError := "\"pulumi-test-registry/unicorns/swiftwind:latest\": repository name must be canonical.\n" +
			"This resource requires all image names to be fully qualified.\n" +
			"For example, if you are attempting to push to Dockerhub, prefix your image name with `docker.io`:\n\n" +
			"`docker.io/repository/image:tag`"

		actual, err := getRegistryAddrFromImage(input)
		assert.Equal(t, expected, actual)
		assert.Error(t, err)
		assert.ErrorContains(t, err, expectedError)
	})
}

func TestConfigureDockerClient(t *testing.T) {
	t.Run("Given a host passed via pulumi config, a client should have that host", func(t *testing.T) {
		expected := "testhost://something.sock"
		input := map[string]string{
			"host": "testhost://something.sock",
		}

		actual, err := configureDockerClient(input, false)
		assert.NoError(t, err)
		assert.Equal(t, expected, actual.DaemonHost())
	})

	t.Run("For TLS, must pass certMaterial, keyMaterial, and caMaterial", func(t *testing.T) {
		input := map[string]string{
			"caMaterial": "raw-cert-string",
		}
		actual, err := configureDockerClient(input, false)
		expectedError := fmt.Errorf("certMaterial, keyMaterial, and caMaterial must all be specified")
		if assert.Error(t, err) {
			assert.Equal(t, expectedError, err)
		}
		assert.Nil(t, actual)
	})
	t.Run("Errors if only caMaterial is specified", func(t *testing.T) {
		input := map[string]string{
			"caMaterial": "raw-ca-string",
		}
		actual, err := configureDockerClient(input, false)
		expectedError := fmt.Errorf("certMaterial, keyMaterial, and caMaterial must all be specified")
		if assert.Error(t, err) {
			assert.Equal(t, expectedError, err)
		}
		assert.Nil(t, actual)
	})
	t.Run("Errors if only keyMaterial is specified", func(t *testing.T) {
		input := map[string]string{
			"keyMaterial": "raw-key-string",
		}
		actual, err := configureDockerClient(input, false)
		expectedError := fmt.Errorf("certMaterial, keyMaterial, and caMaterial must all be specified")
		if assert.Error(t, err) {
			assert.Equal(t, expectedError, err)
		}
		assert.Nil(t, actual)
	})

	t.Run("Errors if not all of certMaterial, keyMaterial, and caMaterial are specified", func(t *testing.T) {
		input := map[string]string{
			"caMaterial":   "raw-ca-string",
			"certMaterial": "raw-cert-string",
		}
		actual, err := configureDockerClient(input, false)
		expectedError := fmt.Errorf("certMaterial, keyMaterial, and caMaterial must all be specified")
		if assert.Error(t, err) {
			assert.Equal(t, expectedError, err)
		}
		assert.Nil(t, actual)
	})

	t.Run("Fails if both a certPath and raw certificates are passed", func(t *testing.T) {
		input := map[string]string{
			"certPath":     "path/to/certs",
			"caMaterial":   "raw-ca-string",
			"keyMaterial":  "raw-key-string",
			"certMaterial": "raw-cert-string",
		}
		actual, err := configureDockerClient(input, false)
		expectedError := fmt.Errorf("when using raw certificates, certPath must not be specified")
		if assert.Error(t, err) {
			assert.Equal(t, expectedError, err)
		}
		assert.Nil(t, actual)
	})

	t.Run("When passed a valid ssh scheme for the host, a client with a helper daemon host will be returned",
		func(t *testing.T) {
			input := map[string]string{
				"host": "ssh://test@128.199.8.23",
			}
			actual, _ := configureDockerClient(input, false)
			// The connection helper returns http://docker.example.com as the client's daemon host.
			assert.Equal(t, actual.DaemonHost(), "http://docker.example.com")
		})
	t.Run("When passed an invalid ssh scheme for the host, no client is returned",
		func(t *testing.T) {
			input := map[string]string{
				"host": "ssh://this/is?invalid",
			}
			actual, err := configureDockerClient(input, false)
			assert.Nil(t, actual)
			assert.ErrorContains(t, err, "ssh host connection is not valid")
		})

	t.Run("When passed a valid non-ssh scheme for the host, a client without daemon host will be returned",
		func(t *testing.T) {
			input := map[string]string{
				"host": "unix:///var/run/docker.sock",
			}
			actual, _ := configureDockerClient(input, false)
			assert.Equal(t, actual.DaemonHost(), input["host"])
		})
	t.Run("When host is empty, returns default host ", func(t *testing.T) {
		input := map[string]string{
			"host": "",
		}
		actual, _ := configureDockerClient(input, false)

		os := runtime.GOOS
		switch os {
		case "windows":
			assert.Equal(t, actual.DaemonHost(), "npipe:////./pipe/docker_engine")
		default:
			assert.Equal(t, actual.DaemonHost(), "unix:///var/run/docker.sock")
		}
	})
}
