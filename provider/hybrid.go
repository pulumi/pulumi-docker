package provider

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
	rpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Hybrid provider struct
type dockerHybridProvider struct {
	schemaBytes     []byte
	version         string
	bridgedProvider rpc.ResourceProviderServer
	nativeProvider  rpc.ResourceProviderServer
}

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

func (dp dockerHybridProvider) GetSchema(ctx context.Context, request *rpc.GetSchemaRequest) (*rpc.GetSchemaResponse, error) {
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

func (dp dockerHybridProvider) Configure(ctx context.Context, request *rpc.ConfigureRequest) (*rpc.ConfigureResponse, error) {
	var myResp rpc.ConfigureResponse
	for _, prov := range []rpc.ResourceProviderServer{dp.bridgedProvider, dp.nativeProvider} {
		resp, err := prov.Configure(ctx, request)
		if err != nil {
			return nil, err
		}
		myResp = *resp
	}
	return &myResp, nil
}

// TODO: this is for functions AKA data sources, and our provider doesn't have any metadatea becuase we're just implementing from scratch
func (dp dockerHybridProvider) Invoke(ctx context.Context, request *rpc.InvokeRequest) (*rpc.InvokeResponse, error) {
	// TODO: remove below snippet, as we're not implementing data sources here atm, or implement a default once the way in which we're passing in any ExtraDataSources is better
	//if _, ok := dp.metadata.Functions[request.Tok]; ok {
	//	return dp.nativeProvider.Invoke(ctx, request)
	//}
	logging.V(9).Infof("Invoking on bridge provider for: %q", request.Tok)
	return dp.bridgedProvider.Invoke(ctx, request)
}

func (dp dockerHybridProvider) StreamInvoke(request *rpc.InvokeRequest, server rpc.ResourceProvider_StreamInvokeServer) error {
	return status.Error(codes.Unimplemented, "StreamInvoke is not yet implemented")
}

func (dp dockerHybridProvider) Check(ctx context.Context, request *rpc.CheckRequest) (*rpc.CheckResponse, error) {
	urn := resource.URN(request.GetUrn())

	tok := urn.Type().String()
	fmt.Println(tok)
	// TODO: implement this for actual!!!
	if tok == dockerImageToken {
		return dp.nativeProvider.Check(ctx, request)
	}
	return dp.bridgedProvider.Check(ctx, request)
}

func (dp dockerHybridProvider) Diff(ctx context.Context, request *rpc.DiffRequest) (*rpc.DiffResponse, error) {
	urn := resource.URN(request.GetUrn())
	tok := urn.Type().String()
	if tok == dockerImageToken {
		return dp.nativeProvider.Diff(ctx, request)
	}
	return dp.bridgedProvider.Diff(ctx, request)
}

func (dp dockerHybridProvider) Create(ctx context.Context, request *rpc.CreateRequest) (*rpc.CreateResponse, error) {
	urn := resource.URN(request.GetUrn())
	tok := urn.Type().String()
	if tok == dockerImageToken {
		return dp.nativeProvider.Create(ctx, request)
	}
	return dp.bridgedProvider.Create(ctx, request)
}

func (dp dockerHybridProvider) Read(ctx context.Context, request *rpc.ReadRequest) (*rpc.ReadResponse, error) {
	urn := resource.URN(request.GetUrn())
	tok := urn.Type().String()
	if tok == dockerImageToken {
		return dp.nativeProvider.Read(ctx, request)
	}
	return dp.bridgedProvider.Read(ctx, request)
}

func (dp dockerHybridProvider) Update(ctx context.Context, request *rpc.UpdateRequest) (*rpc.UpdateResponse, error) {
	urn := resource.URN(request.GetUrn())
	tok := urn.Type().String()
	if tok == dockerImageToken {
		return dp.nativeProvider.Update(ctx, request)
	}
	return dp.bridgedProvider.Update(ctx, request)
}

func (dp dockerHybridProvider) Delete(ctx context.Context, request *rpc.DeleteRequest) (*empty.Empty, error) {
	urn := resource.URN(request.GetUrn())
	tok := urn.Type().String()
	if tok == dockerImageToken {
		return dp.nativeProvider.Delete(ctx, request)
	}
	return dp.bridgedProvider.Delete(ctx, request)
}

func (dp dockerHybridProvider) Construct(ctx context.Context, request *rpc.ConstructRequest) (*rpc.ConstructResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Construct is not yet implemented")
}

func (dp dockerHybridProvider) GetPluginInfo(ctx context.Context, empty *empty.Empty) (*rpc.PluginInfo, error) {
	return &rpc.PluginInfo{
		Version: dp.version,
	}, nil
}
