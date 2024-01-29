//go:generate go run go.uber.org/mock/mockgen -typed -package mock -source providercontext.go -destination mock/providercontext.go
package internal

import (
	provider "github.com/pulumi/pulumi-go-provider"
)

// Workaround for https://github.com/pulumi/pulumi-go-provider/issues/159
type ProviderContext interface {
	provider.Context
}
