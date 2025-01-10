package provider

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	structpb "github.com/golang/protobuf/ptypes/struct"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
	rpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
)

// Hybrid provider struct
// In order to add schematized, pulumi-native resources to a TF bridged provider,
// we declare both a native provider and a bridged provider, and multiplex them
// into a hybrid provider. The hybrid provider gRPC methods determine which resource
// gets handled by which provider via if/else logic.

type dockerHybridProvider struct {
	rpc.UnimplementedResourceProviderServer

	schemaBytes     []byte
	name            string
	version         string
	bridgedProvider rpc.ResourceProviderServer
	nativeProvider  rpc.ResourceProviderServer
	configEncoding  *tfbridge.ConfigEncoding
}

// Track a list of native resource tokens
const dockerImageTok = "docker:index/image:Image"

// gRPC methods for the hybrid provider

func (dp dockerHybridProvider) Attach(ctx context.Context, attach *rpc.PluginAttach) (*emptypb.Empty, error) {
	_, err := dp.bridgedProvider.Attach(ctx, attach)
	if err != nil {
		return nil, err
	}
	return dp.nativeProvider.Attach(ctx, attach)
}

func (dp dockerHybridProvider) Call(context.Context, *rpc.CallRequest) (*rpc.CallResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Call is not yet implemented")
}

func (dp dockerHybridProvider) Cancel(context.Context, *empty.Empty) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func (dp dockerHybridProvider) GetSchema(_ context.Context, request *rpc.GetSchemaRequest) (
	*rpc.GetSchemaResponse, error) {
	if v := request.GetVersion(); v != 0 {
		return nil, fmt.Errorf("unsupported schema version %d", v)
	}
	return &rpc.GetSchemaResponse{Schema: string(dp.schemaBytes)}, nil
}

func (dp dockerHybridProvider) CheckConfig(ctx context.Context, request *rpc.CheckRequest) (*rpc.CheckResponse, error) {
	// Delegate to the bridged provider, as native Provider does not implement it.
	return dp.bridgedProvider.CheckConfig(ctx, request)
}

func (dp dockerHybridProvider) DiffConfig(ctx context.Context, request *rpc.DiffRequest) (*rpc.DiffResponse, error) {
	urn := resource.URN(request.GetUrn())
	label := fmt.Sprintf("%s.DiffConfig(%s)", dp.name, urn)
	logging.V(9).Infof("%s executing", label)

	var err error
	request.Olds, err = dp.unwrapJSONConfig(label, request.GetOlds())
	if err != nil {
		return nil, fmt.Errorf("error unwrapping old config: %w", err)
	}
	request.News, err = dp.unwrapJSONConfig(label, request.GetNews())
	if err != nil {
		return nil, fmt.Errorf("error unwrapping new config: %w", err)
	}
	request.OldInputs, err = dp.unwrapJSONConfig(label, request.GetOldInputs())
	if err != nil {
		return nil, fmt.Errorf("error unwrapping old inputs: %w", err)
	}

	ignoreChanges := []string{"registryAuth[*].password", "registryAuth[*].username"}
	ignoreChanges = append(ignoreChanges, request.IgnoreChanges...)
	request.IgnoreChanges = ignoreChanges

	res, err := dp.bridgedProvider.DiffConfig(ctx, request)
	if err != nil {
		return nil, err
	}

	// if the diff is empty, it means it only contained changes to username and password which we ignored
	if res != nil && len(res.Diffs) == 0 && len(res.Replaces) == 0 && len(res.DetailedDiff) == 0 {
		res.Changes = rpc.DiffResponse_DIFF_NONE
	}
	return res, err
}

// unwrapJSONConfig handles nested provider configuration data that can be in two formats:
// 1. A JSON-encoded string containing the nested configuration (used by default providers and explicit providers
// for TypeScript, Python, .NET, Java)
// 2. A regular gRPC struct (used by explicit providers for Go and YAML)
//
// For JSON-encoded strings, it decodes the nested config. For gRPC structs, it returns the config unchanged.
// Under the hood, this is implemented by unmarshalling the grpc struct, unfolding the properties,
// and then marshalling them back to a grpc struct.
//
// This dual format support is needed because different language runtimes serialize their
// provider configs differently when sending them over gRPC.
//
// Note that this function does not preserve secrets, as this provider does not accept secrets. The provider relies on
// the engine to handle secrets.
func (dp dockerHybridProvider) unwrapJSONConfig(label string, config *structpb.Struct) (*structpb.Struct, error) {
	unmarshalled, err := plugin.UnmarshalProperties(config, plugin.MarshalOptions{
		Label:        label,
		KeepUnknowns: true,
		SkipNulls:    true,
		// the provider does not accept secrets, so we should never receive them here. There's e2e tests ensuring that
		// secrets in provider config are handled correctly. If this assumption changes, those tests will catch it.
		KeepSecrets: false,
	})
	if err != nil {
		return nil, err
	}

	unwrappedConfig, err := dp.configEncoding.UnfoldProperties(unmarshalled)
	if err != nil {
		return nil, err
	}

	return plugin.MarshalProperties(unwrappedConfig, plugin.MarshalOptions{
		Label:        label,
		KeepUnknowns: true,
		SkipNulls:    true,
		KeepSecrets:  false,
	})
}

func (dp dockerHybridProvider) Configure(
	ctx context.Context,
	request *rpc.ConfigureRequest,
) (*rpc.ConfigureResponse, error) {
	// Native provider returns empty response and error from Configure, just call it to propagate the information.
	r, err := dp.nativeProvider.Configure(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("Docker native provider returned an unexpected error from Configure: %w", err)
	}

	contract.Assertf(!r.AcceptOutputs, "Unexpected AcceptOutputs=true from Docker native provider Configure")
	contract.Assertf(!r.AcceptResources, "Unexpected AcceptResources=true from Docker native provider Configure")
	contract.Assertf(!r.AcceptSecrets, "Unexpected AcceptSecrets=true from Docker native provider Configure")
	contract.Assertf(!r.SupportsPreview, "Unexpected SupportsPreview=true from Docker native provider Configure")

	// For the most part delegate Configure handling to the bridged provider.
	resp, err := dp.bridgedProvider.Configure(ctx, request)
	if err != nil {
		return nil, err
	}

	resp.SupportsPreview = true
	return resp, err
}

func (dp dockerHybridProvider) Invoke(ctx context.Context, request *rpc.InvokeRequest) (*rpc.InvokeResponse, error) {
	logging.V(9).Infof("Invoking on bridge provider for: %q", request.Tok)
	return dp.bridgedProvider.Invoke(ctx, request)
}

func (dp dockerHybridProvider) StreamInvoke(*rpc.InvokeRequest, rpc.ResourceProvider_StreamInvokeServer) error {
	return status.Error(codes.Unimplemented, "StreamInvoke is not yet implemented")
}

func (dp dockerHybridProvider) Check(ctx context.Context, request *rpc.CheckRequest) (*rpc.CheckResponse, error) {
	urn := resource.URN(request.GetUrn())
	tok := urn.Type().String()
	if tok == dockerImageTok {
		return dp.nativeProvider.Check(ctx, request)
	}
	return dp.bridgedProvider.Check(ctx, request)
}

func (dp dockerHybridProvider) Diff(ctx context.Context, request *rpc.DiffRequest) (*rpc.DiffResponse, error) {
	urn := resource.URN(request.GetUrn())
	tok := urn.Type().String()
	if tok == dockerImageTok {
		return dp.nativeProvider.Diff(ctx, request)
	}
	return dp.bridgedProvider.Diff(ctx, request)
}

func (dp dockerHybridProvider) Create(ctx context.Context, request *rpc.CreateRequest) (*rpc.CreateResponse, error) {
	urn := resource.URN(request.GetUrn())
	tok := urn.Type().String()
	if tok == dockerImageTok {
		return dp.nativeProvider.Create(ctx, request)
	}
	return dp.bridgedProvider.Create(ctx, request)
}

func (dp dockerHybridProvider) Read(ctx context.Context, request *rpc.ReadRequest) (*rpc.ReadResponse, error) {
	urn := resource.URN(request.GetUrn())
	tok := urn.Type().String()
	if tok == dockerImageTok {
		return dp.nativeProvider.Read(ctx, request)
	}
	return dp.bridgedProvider.Read(ctx, request)
}

func (dp dockerHybridProvider) Update(ctx context.Context, request *rpc.UpdateRequest) (*rpc.UpdateResponse, error) {
	urn := resource.URN(request.GetUrn())
	tok := urn.Type().String()
	if tok == dockerImageTok {
		return dp.nativeProvider.Update(ctx, request)
	}
	return dp.bridgedProvider.Update(ctx, request)
}

func (dp dockerHybridProvider) Delete(ctx context.Context, request *rpc.DeleteRequest) (*empty.Empty, error) {
	urn := resource.URN(request.GetUrn())
	tok := urn.Type().String()
	if tok == dockerImageTok {
		return dp.nativeProvider.Delete(ctx, request)
	}
	return dp.bridgedProvider.Delete(ctx, request)
}

func (dp dockerHybridProvider) Construct(context.Context, *rpc.ConstructRequest) (
	*rpc.ConstructResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Construct is not yet implemented")
}

func (dp dockerHybridProvider) GetPluginInfo(context.Context, *empty.Empty) (*rpc.PluginInfo, error) {
	return &rpc.PluginInfo{
		Version: dp.version,
	}, nil
}
