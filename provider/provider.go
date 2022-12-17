package provider

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	pbempty "github.com/golang/protobuf/ptypes/empty"
	"github.com/moby/buildkit/frontend/dockerfile/dockerignore"
	"github.com/moby/moby/pkg/fileutils"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/pkg/v3/resource/provider"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
	rpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type dockerNativeProvider struct {
	host        *provider.HostClient
	name        string
	version     string
	schemaBytes []byte
	//loginLock   sync.Mutex
}

// docker native methods

// Attach sends the engine address to an already running plugin.
func (p *dockerNativeProvider) Attach(context context.Context, req *rpc.PluginAttach) (*emptypb.Empty, error) {
	host, err := provider.NewHostClient(req.GetAddress())
	if err != nil {
		return nil, err
	}
	p.host = host
	return &pbempty.Empty{}, nil
}

// Call dynamically executes a method in the provider associated with a component resource.
func (p *dockerNativeProvider) Call(ctx context.Context, req *rpc.CallRequest) (*rpc.CallResponse, error) {
	return nil, status.Error(codes.Unimplemented, "call is not yet implemented")
}

// Construct creates a new component resource.
func (p *dockerNativeProvider) Construct(ctx context.Context, req *rpc.ConstructRequest) (
	*rpc.ConstructResponse, error) {
	return nil, status.Error(codes.Unimplemented, "construct is not yet implemented")
}

// CheckConfig validates the configuration for this provider.
func (p *dockerNativeProvider) CheckConfig(ctx context.Context, req *rpc.CheckRequest) (*rpc.CheckResponse, error) {
	return &rpc.CheckResponse{Inputs: req.GetNews()}, nil
}

// DiffConfig diffs the configuration for this provider.
func (p *dockerNativeProvider) DiffConfig(ctx context.Context, req *rpc.DiffRequest) (*rpc.DiffResponse, error) {
	return &rpc.DiffResponse{}, nil
}

// Configure configures the resource provider with "globals" that control its behavior.
func (p *dockerNativeProvider) Configure(_ context.Context, req *rpc.ConfigureRequest) (*rpc.ConfigureResponse, error) {
	return &rpc.ConfigureResponse{
		AcceptSecrets: true,
	}, nil
}

// Invoke dynamically executes a built-in function in the provider.
func (p *dockerNativeProvider) Invoke(_ context.Context, req *rpc.InvokeRequest) (*rpc.InvokeResponse, error) {
	tok := req.GetTok()
	return nil, fmt.Errorf("unknown Invoke token '%s'", tok)
}

// StreamInvoke dynamically executes a built-in function in the provider. The result is streamed
// back as a series of messages.
func (p *dockerNativeProvider) StreamInvoke(
	req *rpc.InvokeRequest, server rpc.ResourceProvider_StreamInvokeServer) error {
	tok := req.GetTok()
	return fmt.Errorf("unknown StreamInvoke token '%s'", tok)
}

// Check validates that the given property bag is valid for a resource of the given type and returns
// the inputs that should be passed to successive calls to Diff, Create, or Update for this
// resource. As a rule, the provider inputs returned by a call to Check should preserve the original
// representation of the properties as present in the program inputs. Though this rule is not
// required for correctness, violations thereof can negatively impact the end-user experience, as
// the provider inputs are using for detecting and rendering diffs.
func (p *dockerNativeProvider) Check(ctx context.Context, req *rpc.CheckRequest) (*rpc.CheckResponse, error) {
	urn := resource.URN(req.GetUrn())
	label := fmt.Sprintf("%s.Create(%s)", p.name, urn)
	logging.V(9).Infof("%s executing", label)

	inputs, err := plugin.UnmarshalProperties(req.GetNews(), plugin.MarshalOptions{
		KeepUnknowns: true,
		SkipNulls:    true,
		KeepSecrets:  true,
	})
	if err != nil {
		return nil, err
	}

	// Set defaults
	build, err := marshalBuildAndApplyDefaults(inputs["build"])
	if err != nil {
		return nil, err
	}

	dockerContext := build.Context
	dockerfile := build.Dockerfile

	// Set docker build context
	contextDigest, err := hashContext(dockerContext, dockerfile)
	if err != nil {
		return nil, err
	}

	// add implicit resource to provider
	inputs["build"].ObjectValue()["contextDigest"] = resource.NewStringProperty(contextDigest)

	inputStruct, err := plugin.MarshalProperties(inputs, plugin.MarshalOptions{
		KeepUnknowns: true,
		SkipNulls:    true,
		KeepSecrets:  true,
	})

	if err != nil {
		return nil, err
	}

	return &rpc.CheckResponse{Inputs: inputStruct, Failures: nil}, nil
}

// Diff checks what impacts a hypothetical update will have on the resource's properties.
func (p *dockerNativeProvider) Diff(ctx context.Context, req *rpc.DiffRequest) (*rpc.DiffResponse, error) {
	urn := resource.URN(req.GetUrn())
	label := fmt.Sprintf("%s.Diff(%s)", p.name, urn)
	logging.V(9).Infof("%s executing", label)

	oldState, err := plugin.UnmarshalProperties(req.GetOlds(), plugin.MarshalOptions{
		KeepUnknowns: true,
		SkipNulls:    true,
		KeepSecrets:  true,
	})
	if err != nil {
		return nil, err
	}

	// Extract old inputs from the `__inputs` field of the old state.
	oldInputs := parseCheckpointObject(oldState)

	news, err := plugin.UnmarshalProperties(req.GetNews(), plugin.MarshalOptions{
		KeepUnknowns: true,
		SkipNulls:    true,
		KeepSecrets:  true,
	})
	if err != nil {
		return nil, err
	}

	d := oldInputs.Diff(news)

	if d == nil {
		return &rpc.DiffResponse{
			Changes: rpc.DiffResponse_DIFF_NONE,
		}, nil
	}

	diff := map[string]*rpc.PropertyDiff{}
	for key := range d.Adds {
		diff[string(key)] = &rpc.PropertyDiff{Kind: rpc.PropertyDiff_ADD}
	}
	for key := range d.Deletes {
		diff[string(key)] = &rpc.PropertyDiff{Kind: rpc.PropertyDiff_DELETE}
	}
	for key := range d.Updates {
		diff[string(key)] = &rpc.PropertyDiff{Kind: rpc.PropertyDiff_UPDATE}
	}
	return &rpc.DiffResponse{
		Changes:         rpc.DiffResponse_DIFF_SOME,
		DetailedDiff:    diff,
		HasDetailedDiff: true,
	}, nil
}

// Create allocates a new instance of the provided resource and returns its unique ID afterwards.
func (p *dockerNativeProvider) Create(ctx context.Context, req *rpc.CreateRequest) (*rpc.CreateResponse, error) {
	urn := resource.URN(req.GetUrn())
	label := fmt.Sprintf("%s.Create(%s)", p.name, urn)
	logging.V(9).Infof("%s executing", label)

	// Deserialize RPC inputs.
	inputs, err := plugin.UnmarshalProperties(req.GetProperties(), plugin.MarshalOptions{
		Label:        fmt.Sprintf("%s.inputs", label),
		KeepUnknowns: true,
		RejectAssets: true,
		KeepSecrets:  true,
		SkipNulls:    true,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "malformed resource inputs")
	}

	id, outputProperties, err := p.dockerBuild(ctx, urn, req.GetProperties())
	if err != nil {
		return nil, err
	}

	outputs, err := plugin.UnmarshalProperties(outputProperties, plugin.MarshalOptions{
		Label:        fmt.Sprintf("%s.outputs", label),
		KeepUnknowns: true,
		RejectAssets: true,
		KeepSecrets:  true,
		SkipNulls:    true,
	})
	if err != nil {
		return nil, err
	}

	// Store both outputs and inputs into the state.
	checkpoint, err := plugin.MarshalProperties(
		checkpointObject(inputs, outputs.Mappable()),
		plugin.MarshalOptions{
			Label:        fmt.Sprintf("%s.checkpoint", label),
			KeepSecrets:  true,
			KeepUnknowns: true,
			SkipNulls:    true,
		},
	)
	if err != nil {
		return nil, err
	}

	return &rpc.CreateResponse{
		Id:         id,
		Properties: checkpoint,
	}, nil
}

// Read the current live state associated with a resource.
func (p *dockerNativeProvider) Read(ctx context.Context, req *rpc.ReadRequest) (*rpc.ReadResponse, error) {
	urn := resource.URN(req.GetUrn())
	label := fmt.Sprintf("%s.Read(%s)", p.name, urn)
	logging.V(9).Infof("%s executing", label)
	id := req.GetId()

	// For now, we will return the minimum implementation for Read,
	// until we find a use case for reading an image from a registry and comparing it to existing state
	return &rpc.ReadResponse{Id: id}, nil

}

// Update updates an existing resource with new values.
func (p *dockerNativeProvider) Update(ctx context.Context, req *rpc.UpdateRequest) (*rpc.UpdateResponse, error) {
	urn := resource.URN(req.GetUrn())
	label := fmt.Sprintf("%s.Update(%s)", p.name, urn)
	logging.V(9).Infof("%s executing", label)

	// Read the inputs to persist them into state.
	newInputs, err := plugin.UnmarshalProperties(req.GetNews(), plugin.MarshalOptions{
		Label:        fmt.Sprintf("%s.newInputs", label),
		KeepUnknowns: true,
		RejectAssets: true,
		KeepSecrets:  true,
		SkipNulls:    true,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "diff failed because malformed resource inputs")
	}
	// When the docker image is updated, we build and push again.
	_, outputProperties, err := p.dockerBuild(ctx, urn, req.GetNews())
	if err != nil {
		return nil, err
	}
	outputs, err := plugin.UnmarshalProperties(outputProperties, plugin.MarshalOptions{
		Label:        fmt.Sprintf("%s.outputs", label),
		KeepUnknowns: true,
		RejectAssets: true,
		KeepSecrets:  true,
		SkipNulls:    true,
	})
	if err != nil {
		return nil, err
	}
	// Store both outputs and inputs into the state and return RPC checkpoint.
	checkpoint, err := plugin.MarshalProperties(
		checkpointObject(newInputs, outputs.Mappable()),
		plugin.MarshalOptions{
			Label:        fmt.Sprintf("%s.checkpoint", label),
			KeepSecrets:  true,
			KeepUnknowns: true,
			SkipNulls:    true,
		},
	)

	if err != nil {
		return nil, err
	}
	return &rpc.UpdateResponse{
		Properties: checkpoint,
	}, nil
}

// Delete tears down an existing resource with the given ID.  If it fails, the resource is assumed
// to still exist.
func (p *dockerNativeProvider) Delete(ctx context.Context, req *rpc.DeleteRequest) (*pbempty.Empty, error) {
	urn := resource.URN(req.GetUrn())
	label := fmt.Sprintf("%s.Update(%s)", p.name, urn)
	logging.V(9).Infof("%s executing", label)
	return &pbempty.Empty{}, nil
}

// GetPluginInfo returns generic information about this plugin, like its version.
func (p *dockerNativeProvider) GetPluginInfo(context.Context, *pbempty.Empty) (*rpc.PluginInfo, error) {
	return &rpc.PluginInfo{
		Version: p.version,
	}, nil
}

// GetSchema returns the JSON-serialized schema for the provider.
func (p *dockerNativeProvider) GetSchema(ctx context.Context, req *rpc.GetSchemaRequest) (
	*rpc.GetSchemaResponse, error) {
	if v := req.GetVersion(); v != 0 {
		return nil, fmt.Errorf("unsupported schema version %d", v)
	}
	return &rpc.GetSchemaResponse{Schema: string(p.schemaBytes)}, nil
}

// Cancel signals the provider to gracefully shut down and abort any ongoing resource operations.
// Operations aborted in this way will return an error (e.g., `Update` and `Create` will either a
// creation error or an initialization error). Since Cancel is advisory and non-blocking, it is up
// to the host to decide how long to wait after Cancel is called before (e.g.)
// hard-closing any gRPC connection.
func (p *dockerNativeProvider) Cancel(context.Context, *pbempty.Empty) (*pbempty.Empty, error) {
	// TODO
	return &pbempty.Empty{}, nil
}

// checkpointObject puts inputs in the `__inputs` field of the state.
func checkpointObject(inputs resource.PropertyMap, outputs map[string]interface{}) resource.PropertyMap {
	object := resource.NewPropertyMapFromMap(outputs)
	object["__inputs"] = resource.MakeSecret(resource.NewObjectProperty(inputs))
	return object
}

// parseCheckpointObject returns inputs that are saved in the `__inputs` field of the state.
func parseCheckpointObject(obj resource.PropertyMap) resource.PropertyMap {
	if inputs, ok := obj["__inputs"]; ok {
		return inputs.SecretValue().Element.ObjectValue()
	}

	return nil
}

type contextHashAccumulator struct {
	dockerContextPath string
	input             bytes.Buffer // This will hold the file info and content bytes to pass to a hash object
}

func (accumulator *contextHashAccumulator) hashPath(path string, fileMode fs.FileMode) error {

	hash := sha256.New()

	if fileMode.Type() == fs.ModeSymlink {
		// For symlinks, we hash the symlink _path_ instead of the file content.
		// This will allow us to:
		// a) ignore changes at the symlink target
		// b) detect if the symlink _itself_ changes
		// c) avoid a panic on io.Copy if the symlink target is a directory
		symLinkPath, err := filepath.EvalSymlinks(path)
		if err != nil {
			return fmt.Errorf("could not evaluate symlink at %s: %w", path, err)
		}
		symLinkReader := strings.NewReader(symLinkPath)
		_, err = io.Copy(hash, symLinkReader)
		if err != nil {
			return fmt.Errorf("could not copy symlink path %s to hash: %w", path, err)
		}
	} else {
		// For regular files, we can hash their content.
		// TODO: consider only hashing file metadata to improve performance
		f, err := os.Open(filepath.Join(accumulator.dockerContextPath, path))
		if err != nil {
			return fmt.Errorf("could not open file %s: %w", path, err)
		}
		defer f.Close()
		_, err = io.Copy(hash, f)
		if err != nil {
			return fmt.Errorf("could not copy file %s to hash: %w", path, err)
		}
	}

	// Capture all information in the accumulator buffer and add a separator
	accumulator.input.Write([]byte(path))
	accumulator.input.Write([]byte(fileMode.String()))
	accumulator.input.Write(hash.Sum(nil))
	accumulator.input.WriteByte(0)
	return nil
}

func (accumulator *contextHashAccumulator) hexSumContext() string {
	h := sha256.New()
	_, err := accumulator.input.WriteTo(h)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(h.Sum(nil))
}

func hashContext(dockerContextPath string, dockerfile string) (string, error) {
	// exclude all files listed in dockerignore
	dockerIgnorePath := dockerfile + ".dockerignore"
	dockerIgnore, err := os.ReadFile(dockerIgnorePath)
	if err != nil {
		if os.IsNotExist(err) {
			dockerIgnorePath = filepath.Join(dockerContextPath, ".dockerignore")
			dockerIgnore, err = os.ReadFile(dockerIgnorePath)
			if err != nil && !os.IsNotExist(err) {
				return "", fmt.Errorf("unable to read %s file: %w", dockerIgnorePath, err)
			}
		} else {
			return "", fmt.Errorf("unable to read %s file: %w", dockerIgnorePath, err)
		}
	}
	ignorePatterns, err := dockerignore.ReadAll(bytes.NewReader(dockerIgnore))
	if err != nil {
		return "", fmt.Errorf("unable to parse %s file: %w", dockerIgnorePath, err)
	}
	ignoreMatcher, err := fileutils.NewPatternMatcher(ignorePatterns)
	if err != nil {
		return "", fmt.Errorf("unable to load rules from %s: %w", dockerIgnorePath, err)
	}

	accumulator := contextHashAccumulator{dockerContextPath: dockerContextPath}
	err = accumulator.hashPath(dockerfile, 0)
	if err != nil {
		return "", fmt.Errorf("error hashing dockerfile %q: %w", dockerfile, err)
	}
	// for each file in the Docker build context, create a hash of its content
	err = filepath.WalkDir(dockerContextPath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		path, err = filepath.Rel(dockerContextPath, path)

		if err != nil {
			return err
		}
		if path == "." {
			return nil
		}
		ignore, err := ignoreMatcher.Matches(path)
		if err != nil {
			return fmt.Errorf("%s rule failed: %w", dockerIgnorePath, err)
		}
		if ignore {
			if d.IsDir() {
				return filepath.SkipDir
			}
			return nil

		} else if d.IsDir() {
			return nil
		}
		fileInfo, err := d.Info()
		if err != nil {
			return fmt.Errorf("determining mode for %q: %w", path, err)
		}

		err = accumulator.hashPath(path, fileInfo.Mode())

		if err != nil {
			return fmt.Errorf("error while hashing %q: %w", path, err)
		}
		return nil
	})
	if err != nil {
		return "", fmt.Errorf("unable to hash build context: %w", err)
	}
	// create a hash of the entire input of the hash accumulator
	return accumulator.hexSumContext(), nil
}
