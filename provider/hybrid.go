package provider

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
	rpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Hybrid provider struct
// In order to add schematized, pulumi-native resources to a TF bridged provider,
// we declare both a native provider and a bridged provider, and multiplex them
// into a hybrid provider. The hybrid provider gRPC methods determine which resource
// gets handled by which provider via if/else logic.

type dockerHybridProvider struct {
	rpc.UnimplementedResourceProviderServer

	schemaBytes     []byte
	version         string
	bridgedProvider rpc.ResourceProviderServer
	nativeProvider  rpc.ResourceProviderServer
}

// Track a list of native resource tokens
const dockerImageTok = "docker:index/image:Image"

// gRPC methods for the hybrid provider

func (dp dockerHybridProvider) Attach(ctx context.Context, attach *rpc.PluginAttach) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (dp dockerHybridProvider) Call(ctx context.Context, request *rpc.CallRequest) (*rpc.CallResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Call is not yet implemented")
}

func (dp dockerHybridProvider) Cancel(ctx context.Context, e *empty.Empty) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func (dp dockerHybridProvider) GetSchema(ctx context.Context, request *rpc.GetSchemaRequest) (
	*rpc.GetSchemaResponse, error) {
	if v := request.GetVersion(); v != 0 {
		return nil, fmt.Errorf("unsupported schema version %d", v)
	}
	return &rpc.GetSchemaResponse{Schema: string(dp.schemaBytes)}, nil
}

func (dp dockerHybridProvider) CheckConfig(ctx context.Context, request *rpc.CheckRequest) (*rpc.CheckResponse, error) {
	return &rpc.CheckResponse{Inputs: request.GetNews()}, nil
}

func (dp dockerHybridProvider) DiffConfig(ctx context.Context, request *rpc.DiffRequest) (*rpc.DiffResponse, error) {
	return &rpc.DiffResponse{
		Changes:             0,
		Replaces:            []string{},
		Stables:             []string{},
		DeleteBeforeReplace: false,
	}, nil
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

	// Mostly delegate Configure handling to the bridged provider.
	return dp.bridgedProvider.Configure(ctx, request)
}

func (dp dockerHybridProvider) Invoke(ctx context.Context, request *rpc.InvokeRequest) (*rpc.InvokeResponse, error) {
	logging.V(9).Infof("Invoking on bridge provider for: %q", request.Tok)
	return dp.bridgedProvider.Invoke(ctx, request)
}

func (dp dockerHybridProvider) StreamInvoke(
	request *rpc.InvokeRequest, server rpc.ResourceProvider_StreamInvokeServer) error {
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

func (dp dockerHybridProvider) Construct(ctx context.Context, request *rpc.ConstructRequest) (
	*rpc.ConstructResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Construct is not yet implemented")
}

func (dp dockerHybridProvider) GetPluginInfo(ctx context.Context, empty *empty.Empty) (*rpc.PluginInfo, error) {
	return &rpc.PluginInfo{
		Version: dp.version,
	}, nil
}
