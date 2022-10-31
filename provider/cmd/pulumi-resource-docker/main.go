// Copyright 2016-2018, Pulumi Corporation.
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

//go:generate go run ./generate.go

package main

import (
	"context"
	_ "embed"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	docker "github.com/pulumi/pulumi-docker/provider/v3"
	"github.com/pulumi/pulumi-docker/provider/v3/pkg/version"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	rpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
)

//go:embed schema-embed.json
var pulumiSchema []byte


// Track a list of native resource tokens
nativeResourceTokens := {"index/image.Image"}

// Hybrid provider struct
type hybridDockerProvider struct {
	schemaBytes []byte
	version string
	bridgedProvider rpc.ResourceProviderServer
	nativeProvider rpc.ResourceProviderServer
}



func main() {
	// TODO: intercept the native resources here
	// Instead of calling tfbridge.Main, as is customary for bridged providers,
	// implement all calls separately so that we can apply native vs bridged logic at every RPC call
	//
	//tfbridge.Main("docker", version.Version, docker.Provider(), pulumiSchema)

}

func (dp hybridDockerProvider) Call(ctx context.Context, request *rpc.CallRequest) (*rpc.CallResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Call is not yet implemented")
}

func (dp hybridDockerProvider) Cancel(ctx context.Context, e *empty.Empty) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func (dp hybridDockerProvider) GetSchema(ctx context.Context, request *rpc.GetSchemaRequest) (*rpc.GetSchemaResponse, error) {
	decompressed, err := gen.DecompressSchema(dp.schemaBytes) // TODO: WTF is gen
	if err != nil {
		return nil, errors.New("failure loading schema")
	}
	return &rpc.GetSchemaResponse{Schema: string(decompressed)}, nil
}

func (dp hybridDockerProvider) CheckConfig(ctx context.Context, request *rpc.CheckRequest) (*rpc.CheckResponse, error) {
	return &rpc.CheckResponse{Inputs: request.GetNews()}, nil
}

func (dp hybridDockerProvider) DiffConfig(ctx context.Context, request *rpc.DiffRequest) (*rpc.DiffResponse, error) {
	return &rpc.DiffResponse{
		Changes:             0,
		Replaces:            []string{},
		Stables:             []string{},
		DeleteBeforeReplace: false,
	}, nil
}

func (dp hybridDockerProvider) Configure(ctx context.Context, request *rpc.ConfigureRequest) (*rpc.ConfigureResponse, error) {
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

// TODO: this is for functionsd AKA data sources, and our provider doesn't have any metadatea becuase we're just implementing from scratch
func (dp hybridDockerProvider) Invoke(ctx context.Context, request *rpc.InvokeRequest) (*rpc.InvokeResponse, error) {
	// TODO: remove below snippet, as we're not implementing data sources here atm, or implement a default once the way in which we're passing in any ExtraDataSources is better
	//if _, ok := dp.metadata.Functions[request.Tok]; ok {
	//	return dp.nativeProvider.Invoke(ctx, request)
	//}
	logging.V(9).Infof("Invoking on bridge provider for: %q", request.Tok)
	return dp.bridgedProvider.Invoke(ctx, request)
}

func (dp hybridDockerProvider) StreamInvoke(request *rpc.InvokeRequest, server rpc.ResourceProvider_StreamInvokeServer) error {
	return status.Error(codes.Unimplemented, "StreamInvoke is not yet implemented")
}

func (dp hybridDockerProvider) Check(ctx context.Context, request *rpc.CheckRequest) (*rpc.CheckResponse, error) {
	urn := resource.URN(request.GetUrn())
	tok := urn.Type().String()
	// TODO: implement this for actual!!!
	if _, isNative := dp.metadata.Resources[tok]; isNative {
		return dp.nativeProvider.Check(ctx, request)
	}
	return dp.bridgedProvider.Check(ctx, request)
}

func (dp hybridDockerProvider) Diff(ctx context.Context, request *rpc.DiffRequest) (*rpc.DiffResponse, error) {
	urn := resource.URN(request.GetUrn())
	tok := urn.Type().String()
	if _, isNative := dp.metadata.Resources[tok]; isNative {
		return dp.nativeProvider.Diff(ctx, request)
	}
	return dp.bridgedProvider.Diff(ctx, request)
}

func (dp hybridDockerProvider) Create(ctx context.Context, request *rpc.CreateRequest) (*rpc.CreateResponse, error) {
	urn := resource.URN(request.GetUrn())
	tok := urn.Type().String()
	if _, isNative := dp.metadata.Resources[tok]; isNative {
		return dp.nativeProvider.Create(ctx, request)
	}
	return dp.bridgedProvider.Create(ctx, request)
}

func (dp hybridDockerProvider) Read(ctx context.Context, request *rpc.ReadRequest) (*rpc.ReadResponse, error) {
	urn := resource.URN(request.GetUrn())
	tok := urn.Type().String()
	if _, isNative := dp.metadata.Resources[tok]; isNative {
		return dp.nativeProvider.Read(ctx, request)
	}
	return dp.bridgedProvider.Read(ctx, request)
}

func (dp hybridDockerProvider) Update(ctx context.Context, request *rpc.UpdateRequest) (*rpc.UpdateResponse, error) {
	urn := resource.URN(request.GetUrn())
	tok := urn.Type().String()
	if _, isNative := dp.metadata.Resources[tok]; isNative {
		return dp.nativeProvider.Update(ctx, request)
	}
	return dp.bridgedProvider.Update(ctx, request)
}

func (dp hybridDockerProvider) Delete(ctx context.Context, request *rpc.DeleteRequest) (*empty.Empty, error) {
	urn := resource.URN(request.GetUrn())
	tok := urn.Type().String()
	if _, isNative := dp.metadata.Resources[tok]; isNative {
		return dp.nativeProvider.Delete(ctx, request)
	}
	return dp.bridgedProvider.Delete(ctx, request)
}

func (dp hybridDockerProvider) Construct(ctx context.Context, request *rpc.ConstructRequest) (*rpc.ConstructResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Construct is not yet implemented")
}

func (dp hybridDockerProvider) GetPluginInfo(ctx context.Context, empty *empty.Empty) (*rpc.PluginInfo, error) {
	return &rpc.PluginInfo{
		Version: dp.version,
	}, nil
}
