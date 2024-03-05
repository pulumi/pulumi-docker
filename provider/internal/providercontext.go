//go:generate go run go.uber.org/mock/mockgen -typed -package internal -source providercontext.go -destination mockprovidercontext_test.go --self_package github.com/pulumi/pulumi-docker/provider/v4/internal
package internal

import (
	provider "github.com/pulumi/pulumi-go-provider"
)

// Workaround for https://github.com/pulumi/pulumi-go-provider/issues/159
type ProviderContext interface {
	provider.Context
}
