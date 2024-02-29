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
	"encoding/json"
	"fmt"
	"sort"

	"google.golang.org/protobuf/types/known/structpb"

	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
)

type ConfigEncoding struct {
	schema schema.ConfigSpec
}

// New constructs a new config encoder for the provided spec.
func New(s schema.ConfigSpec) *ConfigEncoding {
	return &ConfigEncoding{schema: s}
}

func (*ConfigEncoding) tryUnwrapSecret(encoded any) (any, bool) {
	m, ok := encoded.(map[string]any)
	if !ok {
		return nil, false
	}
	sig, ok := m["4dabf18193072939515e22adb298388d"]
	if !ok {
		return nil, false
	}
	ss, ok := sig.(string)
	if !ok {
		return nil, false
	}
	if ss != "1b47061264138c4ac30d75fd1eb44270" {
		return nil, false
	}
	value, ok := m["value"]
	return value, ok
}

func (enc *ConfigEncoding) convertStringToPropertyValue(s string, prop schema.PropertySpec) (
	resource.PropertyValue, error,
) {
	// If the schema expects a string, we can just return this as-is.
	if prop.Type == "string" {
		return resource.NewStringProperty(s), nil
	}

	// Otherwise, we will attempt to deserialize the input string as JSON and convert the result into a Pulumi
	// property. If the input string is empty, we will return an appropriate zero value.
	if s == "" {
		return enc.zeroValue(prop.Type), nil
	}

	var jsonValue interface{}
	if err := json.Unmarshal([]byte(s), &jsonValue); err != nil {
		return resource.PropertyValue{}, err
	}

	opts := enc.unmarshalOpts()

	// Instead of using resource.NewPropertyValue, specialize it to detect nested json-encoded secrets.
	var replv func(encoded any) (resource.PropertyValue, bool)
	replv = func(encoded any) (resource.PropertyValue, bool) {
		encodedSecret, isSecret := enc.tryUnwrapSecret(encoded)
		if !isSecret {
			return resource.NewNullProperty(), false
		}

		v := resource.NewPropertyValueRepl(encodedSecret, nil, replv)
		if opts.KeepSecrets {
			v = resource.MakeSecret(v)
		}

		return v, true
	}

	return resource.NewPropertyValueRepl(jsonValue, nil, replv), nil
}

func (*ConfigEncoding) zeroValue(typ string) resource.PropertyValue {
	switch typ {
	case "boolean":
		return resource.NewPropertyValue(false)
	case "integer", "number":
		return resource.NewPropertyValue(0)
	case "array":
		return resource.NewPropertyValue([]interface{}{})
	default:
		return resource.NewPropertyValue(map[string]interface{}{})
	}
}

func (enc *ConfigEncoding) unmarshalOpts() plugin.MarshalOptions {
	return plugin.MarshalOptions{
		Label:        "config",
		KeepUnknowns: true,
		SkipNulls:    true,
		RejectAssets: true,
	}
}

// Like plugin.UnmarshalPropertyValue but overrides string parsing with convertStringToPropertyValue.
func (enc *ConfigEncoding) unmarshalPropertyValue(key resource.PropertyKey,
	v *structpb.Value,
) (*resource.PropertyValue, error) {
	opts := enc.unmarshalOpts()

	pv, err := plugin.UnmarshalPropertyValue(key, v, enc.unmarshalOpts())
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling property %q: %w", key, err)
	}

	prop, ok := enc.schema.Variables[string(key)]

	// Only apply JSON-encoded recognition for known fields.
	if !ok {
		return pv, nil
	}

	var jsonString string
	var jsonStringDetected, jsonStringSecret bool

	if pv.IsString() {
		jsonString = pv.StringValue()
		jsonStringDetected = true
	}

	if opts.KeepSecrets && pv.IsSecret() && pv.SecretValue().Element.IsString() {
		jsonString = pv.SecretValue().Element.StringValue()
		jsonStringDetected = true
		jsonStringSecret = true
	}

	if jsonStringDetected {
		v, err := enc.convertStringToPropertyValue(jsonString, prop)
		if err != nil {
			return nil, fmt.Errorf("error unmarshalling property %q: %w", key, err)
		}
		if jsonStringSecret {
			s := resource.MakeSecret(v)
			return &s, nil
		}
		return &v, nil
	}

	// Computed sentinels are coming in as always having an empty string, but the encoding coerses them to a zero
	// value of the appropriate type.
	if pv.IsComputed() {
		el := pv.V.(resource.Computed).Element
		if el.IsString() && el.StringValue() == "" {
			res := resource.MakeComputed(enc.zeroValue(prop.Type))
			return &res, nil
		}
	}

	return pv, nil
}

// Inline from plugin.UnmarshalProperties substituting plugin.UnmarshalPropertyValue.
func (enc *ConfigEncoding) UnmarshalProperties(props *structpb.Struct) (resource.PropertyMap, error) {
	opts := enc.unmarshalOpts()

	result := make(resource.PropertyMap)

	// First sort the keys so we enumerate them in order (in case errors happen, we want determinism).
	var keys []string
	if props != nil {
		for k := range props.Fields {
			keys = append(keys, k)
		}
		sort.Strings(keys)
	}

	// And now unmarshal every field it into the map.
	for _, key := range keys {
		pk := resource.PropertyKey(key)
		v, err := enc.unmarshalPropertyValue(pk, props.Fields[key])
		if err != nil {
			return nil, err
		} else if v != nil {
			if opts.SkipNulls && v.IsNull() {
				continue
			}
			if opts.SkipInternalKeys && resource.IsInternalPropertyKey(pk) {
				continue
			}
			result[pk] = *v
		}
	}

	return result, nil
}
