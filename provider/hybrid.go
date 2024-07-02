package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

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
	version         string
	bridgedProvider rpc.ResourceProviderServer
	nativeProvider  rpc.ResourceProviderServer
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
	// Delegate to the bridged provider, as native Provider does not implement it.

	urn := resource.URN(request.GetUrn())
	label := fmt.Sprintf("DiffConfig(%s)", urn)

	olds, err := plugin.UnmarshalProperties(request.GetOlds(), plugin.MarshalOptions{
		Label:        fmt.Sprintf("%s.olds", label),
		KeepUnknowns: true,
		SkipNulls:    true,
	})
	if err != nil {
		return nil, err
	}
	news, err := plugin.UnmarshalProperties(request.GetNews(), plugin.MarshalOptions{
		Label:        fmt.Sprintf("%s.news", label),
		KeepUnknowns: true,
		SkipNulls:    true,
	})
	if err != nil {
		return nil, fmt.Errorf("DiffConfig failed because of malformed resource inputs: %w", err)
	}

	// forcing this to true for the prototype. This needs to be configurable
	var ignoreRegistryAuthChanges bool = true
	if v := news["ignoreRegistryAuthChanges"]; v.HasValue() && v.IsBool() {
		ignoreRegistryAuthChanges = v.BoolValue()
	}

	resp, err := dp.bridgedProvider.DiffConfig(ctx, request)
	if err != nil {
		return nil, err
	}

	if hasDiff, err := hasVolatileRegistryAuthDiff(olds, news); err == nil && hasDiff && ignoreRegistryAuthChanges {
		for idx, prop := range resp.Diffs {
			if prop == "registryAuth" {
				resp.Diffs = append(resp.Diffs[:idx], resp.Diffs[idx+1:]...)
				break
			}
		}
		delete(resp.DetailedDiff, "registryAuth")

		// Set diff to none if there are no other diffs
		if len(resp.DetailedDiff) == 0 && len(resp.Diffs) == 0 {
			resp.Changes = rpc.DiffResponse_DIFF_NONE
		}
	} else if err != nil {
		return nil, fmt.Errorf("DiffConfig failed because of malformed registryAuth inputs: %w", err)
	}

	return resp, err
}

// hasVolatileRegistryAuthDiff checks if the registryAuth field has a diff that is volatile and should be ignored.
func hasVolatileRegistryAuthDiff(olds, news resource.PropertyMap) (bool, error) {
	if !olds.HasValue("registryAuth") || !news.HasValue("registryAuth") {
		// if old or new registryAuth is missing then there is a diff, no need to clean
		return false, nil
	} else if !olds["registryAuth"].IsString() || !news["registryAuth"].IsString() {
		// if old or new registryAuth is not a string then there is no diff to be cleaned
		return false, nil
	} else if olds["registryAuth"].StringValue() == news["registryAuth"].StringValue() {
		// if the strings are equal then there is no diff to be cleaned
		return false, nil
	}

	var oldAuthData []map[string]interface{}
	var newAuthData []map[string]interface{}
	
	err := json.Unmarshal([]byte(olds["registryAuth"].StringValue()), &oldAuthData)
	if err != nil {
		return false, err
	}
	err = json.Unmarshal([]byte(news["registryAuth"].StringValue()), &newAuthData)
	if err != nil {
		return false, err
	}

	if len(oldAuthData) != len(newAuthData) {
		// an auth data was added/removed, no need to clean because there's an actual diff
		return false, nil
	}

	newAuthMap, err := toAuthMapWithoutVolatileFields(newAuthData)
	if err != nil {
		return false, err
	}
	oldAuthMap, err := toAuthMapWithoutVolatileFields(oldAuthData)
	if err != nil {
		return false, err
	}

	if reflect.DeepEqual(newAuthMap, oldAuthMap) {
		// if the maps are equal without the volatile then they need to be cleaned
		return true, nil
	}

	// fall through, there is an actual diff. No need to clean
	return false, nil
}

var volatileAuthFields = map[string]bool{
	"password": true,
	"username": true,
}

// toAuthMapWithoutVolatileFields converts a slice of auth data maps to a map of auth data maps without volatile fields.
func toAuthMapWithoutVolatileFields(authDataSlice []map[string]interface{}) (map[string]map[string]interface{}, error) {
	authMap := make(map[string]map[string]interface{})
	for _, authData := range authDataSlice {
		address, found := authData["address"]
		if !found {
			return nil, fmt.Errorf("required address field not found in new registryAuth data")
		}
		addressStr, ok := address.(string)
		if !ok {
			return nil, fmt.Errorf("address field in new registryAuth data is not a string")
		}

		authMap[addressStr] = copyMapIgnoringKeys(authData, volatileAuthFields)
	}
	
	return authMap, nil
}

// copyMapIgnoringKeys creates a shallow copy of the given map without the specified keys that should be ignored.
func copyMapIgnoringKeys(original map[string]interface{}, ignoreKeys map[string]bool) map[string]interface{} {
    copy := make(map[string]interface{}, len(original))
    for key, value := range original {
		if ignore, found := ignoreKeys[key]; !found || !ignore {
			copy[key] = value
		}
    }
    return copy
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
