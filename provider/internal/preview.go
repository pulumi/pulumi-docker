package internal

import (
	"github.com/pulumi/pulumi-docker/provider/v4/internal/properties"
)

// keeper decides whether an element should be included for a preview
// operation, optionally returning a mutated copy of that element.
type keeper[T any] interface {
	keep(T) bool
}

// filter applies a keeper to each element, returning a new slice.
func filter[T any](k keeper[T], elems ...T) []T {
	var result []T
	for _, e := range elems {
		if !k.keep(e) {
			continue
		}
		result = append(result, e)
	}
	return result
}

// stringKeeper preserves any non-empty string values for preview.
type stringKeeper struct{ preview bool }

func (k stringKeeper) keep(s string) bool {
	if !k.preview {
		return true
	}
	return s != ""
}

// registryKeeper preserves any registries with known values for address and
// password. This is imprecise and doesn't permit alternative auth strategies
// like registry tokens, email, etc.
type registryKeeper struct{ preview bool }

//nolint:unused // False positive due to generics.
func (k registryKeeper) keep(r properties.RegistryAuth) bool {
	if !k.preview {
		return true
	}
	return r.Password != "" && r.Address != ""
}

// mapKeeper preserves map elements with known keys and values.
type mapKeeper struct{ preview bool }

func (k mapKeeper) keep(m map[string]string) map[string]string {
	if !k.preview || len(m) == 0 {
		return m
	}
	kk := stringKeeper(k)
	filtered := make(map[string]string)
	for key, val := range m {
		if !kk.keep(key) {
			continue
		}
		if !kk.keep(val) {
			continue
		}
		filtered[key] = val
	}
	return filtered
}
