package provider

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
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
	version         string
	bridgedProvider rpc.ResourceProviderServer
	nativeProvider  rpc.ResourceProviderServer
	buildxProvider  rpc.ResourceProviderServer
}

// Track a list of native resource tokens
const (
	dockerImageTok = "docker:index/image:Image"
	buildxTok      = "docker:buildx/image:Image"
)

// gRPC methods for the hybrid provider

func (dp dockerHybridProvider) Attach(ctx context.Context, attach *rpc.PluginAttach) (*emptypb.Empty, error) {
	_, err := dp.bridgedProvider.Attach(ctx, attach)
	if err != nil {
		return nil, fmt.Errorf("attaching bridge provider: %w", err)
	}
	_, err = dp.nativeProvider.Attach(ctx, attach)
	if err != nil {
		return nil, fmt.Errorf("attaching native provider: %w", err)
	}
	_, err = dp.buildxProvider.Attach(ctx, attach)
	if err != nil {
		return nil, fmt.Errorf("attaching buildx provider; %w", err)
	}
	return nil, err
}

func (dp dockerHybridProvider) Call(context.Context, *rpc.CallRequest) (*rpc.CallResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Call is not yet implemented")
}

func (dp dockerHybridProvider) Cancel(context.Context, *empty.Empty) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func (dp dockerHybridProvider) GetSchema(_ context.Context, request *rpc.GetSchemaRequest) (
	*rpc.GetSchemaResponse, error,
) {
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
	// Delegate to the bridged provider, as native Provider does not implement it.
	return dp.bridgedProvider.DiffConfig(ctx, request)
}

func (dp dockerHybridProvider) Configure(
	ctx context.Context,
	request *rpc.ConfigureRequest,
) (*rpc.ConfigureResponse, error) {
	providers := map[string]rpc.ResourceProviderServer{
		"bridged": dp.bridgedProvider,
		"native":  dp.nativeProvider,
		"buildx":  dp.buildxProvider,
	}

	for pname, p := range providers {
		// Native provider returns empty response and error from Configure, just call it to propagate the information.
		r, err := p.Configure(ctx, request)
		if err != nil {
			return nil, fmt.Errorf("Docker %s provider returned an unexpected error from Configure: %w", pname, err)
		}

		if pname != "native" {
			continue
		}

		contract.Assertf(!r.AcceptOutputs,
			fmt.Sprintf("Unexpected AcceptOutputs=true from Docker %s provider Configure", pname))
		contract.Assertf(!r.AcceptResources,
			fmt.Sprintf("Unexpected AcceptResources=true from Docker %s provider Configure", pname))
		contract.Assertf(!r.AcceptSecrets,
			fmt.Sprintf("Unexpected AcceptSecrets=true from Docker %s provider Configure", pname))
		contract.Assertf(!r.SupportsPreview,
			fmt.Sprintf("Unexpected SupportsPreview=true from Docker %s provider Configure", pname))
	}

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
	return dp.providerFor(urn).Check(ctx, request)
}

func (dp dockerHybridProvider) Diff(ctx context.Context, request *rpc.DiffRequest) (*rpc.DiffResponse, error) {
	urn := resource.URN(request.GetUrn())
	return dp.providerFor(urn).Diff(ctx, request)
}

func (dp dockerHybridProvider) Create(ctx context.Context, request *rpc.CreateRequest) (*rpc.CreateResponse, error) {
	urn := resource.URN(request.GetUrn())
	return dp.providerFor(urn).Create(ctx, request)
}

func (dp dockerHybridProvider) Read(ctx context.Context, request *rpc.ReadRequest) (*rpc.ReadResponse, error) {
	urn := resource.URN(request.GetUrn())
	return dp.providerFor(urn).Read(ctx, request)
}

func (dp dockerHybridProvider) Update(ctx context.Context, request *rpc.UpdateRequest) (*rpc.UpdateResponse, error) {
	urn := resource.URN(request.GetUrn())
	return dp.providerFor(urn).Update(ctx, request)
}

func (dp dockerHybridProvider) Delete(ctx context.Context, request *rpc.DeleteRequest) (*empty.Empty, error) {
	urn := resource.URN(request.GetUrn())
	return dp.providerFor(urn).Delete(ctx, request)
}

func (dp dockerHybridProvider) Construct(context.Context, *rpc.ConstructRequest) (
	*rpc.ConstructResponse, error,
) {
	return nil, status.Error(codes.Unimplemented, "Construct is not yet implemented")
}

func (dp dockerHybridProvider) GetPluginInfo(context.Context, *empty.Empty) (*rpc.PluginInfo, error) {
	return &rpc.PluginInfo{
		Version: dp.version,
	}, nil
}

func (dp dockerHybridProvider) providerFor(urn resource.URN) rpc.ResourceProviderServer {
	tok := urn.Type().String()
	switch tok {
	case dockerImageTok:
		return dp.nativeProvider
	case buildxTok:
		return dp.buildxProvider
	default:
		return dp.bridgedProvider
	}
}
