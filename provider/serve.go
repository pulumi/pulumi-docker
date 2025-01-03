package provider

import (
	"context"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi/pkg/v3/resource/provider"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/cmdutil"
	rpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
)

// Serve launches the gRPC server for the resource provider.
func Serve(providerName, version string, schemaBytes []byte) {
	// Start gRPC service.
	err := provider.Main(providerName, func(host *provider.HostClient) (rpc.ResourceProviderServer, error) {
		return makeProvider(host, providerName, version, schemaBytes)
	})
	if err != nil {
		cmdutil.ExitError(err.Error())
	}
}

func makeProvider(host *provider.HostClient, name, version string, schemaBytes []byte) (
	rpc.ResourceProviderServer, error) {
	nativeProvider := &dockerNativeProvider{
		host:        host,
		name:        name,
		version:     version,
		schemaBytes: schemaBytes,
		config:      map[string]string{},
	}

	prov := Provider()
	bridgedProvider := tfbridge.NewProvider(context.Background(), host, name, version, prov.P, prov, schemaBytes)
	configEncoding := tfbridge.NewConfigEncoding(prov.P.Schema(), prov.Config)
	return &dockerHybridProvider{
		schemaBytes:     schemaBytes,
		version:         version,
		name:            name,
		bridgedProvider: bridgedProvider,
		nativeProvider:  nativeProvider,
		configEncoding:  configEncoding,
	}, nil
}