// Copyright 2023, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package deprecated

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	structpb "google.golang.org/protobuf/types/known/structpb"

	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
)

func TestConfigEncoding(t *testing.T) {
	t.Parallel()

	type testCase struct {
		ty schema.TypeSpec
		v  *structpb.Value
		pv resource.PropertyValue
	}

	knownKey := "mykey"

	makeEnc := func(typ schema.TypeSpec) *ConfigEncoding {
		return New(
			schema.ConfigSpec{
				Variables: map[string]schema.PropertySpec{
					knownKey: {
						TypeSpec: typ,
					},
				},
			},
		)
	}

	makeValue := func(x any) *structpb.Value {
		vv, err := structpb.NewValue(x)
		assert.NoErrorf(t, err, "structpb.NewValue failed")
		return vv
	}

	checkUnmarshal := func(t *testing.T, tc testCase) {
		enc := makeEnc(tc.ty)
		pv, err := enc.unmarshalPropertyValue(resource.PropertyKey(knownKey), tc.v)
		assert.NoError(t, err)
		assert.NotNil(t, pv)
		assert.Equal(t, tc.pv, *pv)
	}

	turnaroundTestCases := []testCase{
		{
			schema.TypeSpec{Type: "boolean"},
			makeValue(`true`),
			resource.NewBoolProperty(true),
		},
		{
			schema.TypeSpec{Type: "boolean"},
			makeValue(`false`),
			resource.NewBoolProperty(false),
		},
		{
			schema.TypeSpec{Type: "integer"},
			makeValue(`0`),
			resource.NewNumberProperty(0),
		},
		{
			schema.TypeSpec{Type: "integer"},
			makeValue(`42`),
			resource.NewNumberProperty(42),
		},
		{
			schema.TypeSpec{Type: "number"},
			makeValue(`0`),
			resource.NewNumberProperty(0.0),
		},
		{
			schema.TypeSpec{Type: "number"},
			makeValue(`42.5`),
			resource.NewNumberProperty(42.5),
		},
		{
			schema.TypeSpec{Type: "string"},
			structpb.NewStringValue(""),
			resource.NewStringProperty(""),
		},
		{
			schema.TypeSpec{Type: "string"},
			structpb.NewStringValue("hello"),
			resource.NewStringProperty("hello"),
		},
		{
			schema.TypeSpec{Type: "array"},
			makeValue(`[]`),
			resource.NewArrayProperty([]resource.PropertyValue{}),
		},
		{
			schema.TypeSpec{Type: "array"},
			makeValue(`["hello","there"]`),
			resource.NewArrayProperty([]resource.PropertyValue{
				resource.NewStringProperty("hello"),
				resource.NewStringProperty("there"),
			}),
		},
		{
			schema.TypeSpec{Type: "object"},
			makeValue(`{}`),
			resource.NewObjectProperty(resource.PropertyMap{}),
		},
		{
			schema.TypeSpec{Type: "object"},
			makeValue(`{"key":"value"}`),
			resource.NewObjectProperty(resource.PropertyMap{
				"key": resource.NewStringProperty("value"),
			}),
		},
	}

	t.Run("turnaround", func(t *testing.T) {
		for i, tc := range turnaroundTestCases {
			tc := tc

			t.Run(fmt.Sprintf("UnmarshalPropertyValue/%d", i), func(t *testing.T) {
				t.Parallel()
				checkUnmarshal(t, tc)
			})
		}
	})

	t.Run("zero_values", func(t *testing.T) {
		// Historically the encoding was able to convert empty strings into type-appropriate zero values.
		cases := []testCase{
			{
				schema.TypeSpec{Type: "boolean"},
				makeValue(""),
				resource.NewBoolProperty(false),
			},
			{
				schema.TypeSpec{Type: "number"},
				makeValue(""),
				resource.NewNumberProperty(0.),
			},
			{
				schema.TypeSpec{Type: "integer"},
				makeValue(""),
				resource.NewNumberProperty(0),
			},
			{
				schema.TypeSpec{Type: "string"},
				makeValue(""),
				resource.NewStringProperty(""),
			},
			{
				schema.TypeSpec{Type: "object"},
				makeValue(""),
				resource.NewObjectProperty(make(resource.PropertyMap)),
			},
			{
				schema.TypeSpec{Type: "array"},
				makeValue(""),
				resource.NewArrayProperty([]resource.PropertyValue{}),
			},
		}
		for _, tc := range cases {
			tc := tc

			t.Run(fmt.Sprintf("%v", tc.ty), func(t *testing.T) {
				t.Parallel()
				checkUnmarshal(t, tc)
			})
		}
	})

	t.Run("computed", func(t *testing.T) {
		unk := makeValue(plugin.UnknownStringValue)

		for i, tc := range turnaroundTestCases {
			tc := tc

			t.Run(fmt.Sprintf("UnmarshalPropertyValue/%d", i), func(t *testing.T) {
				t.Parallel()
				// Unknown sentinel unmarshals to a Computed with a type-appropriate zero value.
				checkUnmarshal(t, testCase{
					tc.ty,
					unk,
					resource.MakeComputed(makeEnc(tc.ty).zeroValue(tc.ty.Type)),
				})
			})
		}
	})

	t.Run("secret", func(t *testing.T) {
		// Unmarshalling happens with KeepSecrets=false, replacing them with the underlying values. This case
		// does not need to be tested.
		//
		// Marhalling however supports sending secrets back to the engine, intending to mark values as secret
		// that happen on paths that are declared as secret in the schema. Due to the limitation of the
		// JSON-in-proto-encoding, secrets are communicated imprecisely as an approximation: if any nested
		// element of a property is secret, the entire property is marshalled as secret.

		var secretCases []testCase

		pbSecret := func(v *structpb.Value) *structpb.Value {
			return structpb.NewStructValue(&structpb.Struct{Fields: map[string]*structpb.Value{
				"4dabf18193072939515e22adb298388d": makeValue("1b47061264138c4ac30d75fd1eb44270"),
				"value":                            v,
			}})
		}

		for _, tc := range turnaroundTestCases {
			secretCases = append(secretCases, testCase{
				tc.ty,
				pbSecret(tc.v),
				resource.MakeSecret(tc.pv),
			})
		}

		for i, tc := range secretCases {
			tc := tc

			t.Run(fmt.Sprintf("secret/UnmarshalPropertyValue/%d", i), func(t *testing.T) {
				t.Parallel()

				// Unmarshallin will remove secrts, so the expected value needs to be modified.
				tc.pv = tc.pv.SecretValue().Element
				checkUnmarshal(t, tc)
			})
		}

		t.Run("tolerate secrets in Configure", func(t *testing.T) {
			// This is a bit of a histirocal quirk: the engine may send secrets to Configure before
			// receiving the response from Configure indicating that the provider does not want to receive
			// secrets. These are simply ignored. The engine does not currently send secrets to CheckConfig.
			// The engine does take care of making sure the secrets are stored as such in the statefile.
			//
			// Check here that unmarshalilng such values removes the secrets.
			checkUnmarshal(t, testCase{
				schema.TypeSpec{Type: "object"},
				pbSecret(makeValue(`{"key":"val"}`)),
				resource.NewObjectProperty(resource.PropertyMap{
					"key": resource.NewStringProperty("val"),
				}),
			})
		})
	})

	regressUnmarshalTestCases := []testCase{
		{
			schema.TypeSpec{Type: "array"},
			makeValue(`
			[
			  {
			    "address": "somewhere.org",
			    "password": {
			      "4dabf18193072939515e22adb298388d": "1b47061264138c4ac30d75fd1eb44270",
			      "value": "some-password"
			    },
			    "username": "some-user"
			  }
			]`),
			resource.NewArrayProperty([]resource.PropertyValue{
				resource.NewObjectProperty(resource.PropertyMap{
					"address":  resource.NewStringProperty("somewhere.org"),
					"password": resource.NewStringProperty("some-password"),
					"username": resource.NewStringProperty("some-user"),
				}),
			}),
		},
	}

	t.Run("regress-unmarshal", func(t *testing.T) {
		for i, tc := range regressUnmarshalTestCases {
			tc := tc
			t.Run(fmt.Sprintf("UnmarshalPropertyValue/%d", i), func(t *testing.T) {
				t.Parallel()
				checkUnmarshal(t, tc)
			})
		}
	})
}
