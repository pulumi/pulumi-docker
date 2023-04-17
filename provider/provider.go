package provider

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	pbempty "github.com/golang/protobuf/ptypes/empty"
	"github.com/moby/buildkit/frontend/dockerfile/dockerignore"
	"github.com/moby/patternmatcher"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/pkg/v3/resource/provider"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
	rpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type dockerNativeProvider struct {
	rpc.UnimplementedResourceProviderServer

	host        *provider.HostClient
	name        string
	version     string
	schemaBytes []byte
	config      map[string]string
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
	for key, val := range req.GetVariables() {
		p.config[strings.TrimPrefix(key, "docker:config:")] = val
	}
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
		KeepUnknowns: false,
		SkipNulls:    true,
	})
	if err != nil {
		return nil, err
	}

	// Set defaults
	build, err := marshalBuildAndApplyDefaults(inputs["build"])
	if err != nil {
		return nil, err
	}

	// Verify Dockerfile at given location
	if _, statErr := os.Stat(build.Dockerfile); statErr != nil {
		if filepath.IsAbs(build.Dockerfile) {
			return nil, fmt.Errorf("could not open dockerfile at absolute path %s: %v", build.Dockerfile, statErr)
		}
		relPath := filepath.Join(build.Context, build.Dockerfile)
		_, err = os.Stat(relPath)

		// In the case of a pulumi project that looks as follows:
		// infra/
		//   app/
		//     # some content for the Docker build
		//     Dockerfile
		//   Pulumi.yaml
		//
		//
		// the user inputs:
		//    context: "./app"
		//    dockerfile: "./Dockerfile" # this is in error because it is in "./app/Dockerfile"
		//
		// we want an error message that tells the user: try "./app/Dockerfile"
		if err != nil {
			// no clue case
			return nil, fmt.Errorf("could not open dockerfile at relative path %s: %v", build.Dockerfile, statErr)
		}

		// we could open the relative path
		return nil, fmt.Errorf("could not open dockerfile at relative path %s. "+
			"Try setting `dockerfile` to %q", build.Dockerfile, relPath)

	}
	// Get the relative path to Dockerfile from docker context
	relDockerfile, err := getRelDockerfilePath(build.Context, build.Dockerfile)
	if err != nil {
		return nil, err
	}

	// Hash docker build context digest
	contextDigest, err := hashContext(build.Context, relDockerfile)
	if err != nil {
		return nil, err
	}

	// OS defaults to Linux in all cases
	os := "linux"
	arch := runtime.GOARCH
	hostPlatform := os + "/" + arch
	msg := fmt.Sprintf(
		"Building your image for %s architecture.\n"+
			"To ensure you are building for the correct platform, consider "+
			"explicitly setting the `platform` field on ImageBuildOptions.", hostPlatform)

	// build options: add implicit resource contextDigest and set default host platform
	if inputs["build"].IsNull() {
		inputs["build"] = resource.NewObjectProperty(resource.PropertyMap{
			"contextDigest": resource.NewStringProperty(contextDigest),
			"platform":      resource.NewStringProperty(hostPlatform),
		})
		err = p.host.Log(ctx, "info", urn, msg)
		if err != nil {
			return nil, err
		}
	} else {
		inputs["build"].ObjectValue()["contextDigest"] = resource.NewStringProperty(contextDigest)
		if inputs["build"].ObjectValue()["platform"].IsNull() {
			inputs["build"].ObjectValue()["platform"] = resource.NewStringProperty(hostPlatform)
			err = p.host.Log(ctx, "info", urn, msg)
			if err != nil {
				return nil, err
			}
		}

	}
	if _, err = marshalCachedImages(inputs["build"]); err != nil {
		return nil, err
	}

	inputStruct, err := plugin.MarshalProperties(inputs, plugin.MarshalOptions{
		KeepUnknowns: true,
		SkipNulls:    true,
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
	})
	if err != nil {
		return nil, err
	}

	// Extract old inputs from the `__inputs` field of the old state.
	oldInputs := parseCheckpointObject(oldState)

	news, err := plugin.UnmarshalProperties(req.GetNews(), plugin.MarshalOptions{
		KeepUnknowns: true,
		SkipNulls:    true,
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

	detailedUpdates := diffUpdates(d.Updates)

	// merge detailedUpdates into diff
	for k, v := range detailedUpdates {
		diff[k] = v
	}

	// if diff is empty, it means we skipped any changes to username and password
	if len(diff) == 0 {
		return &rpc.DiffResponse{
			Changes: rpc.DiffResponse_DIFF_NONE,
		}, nil
	}
	return &rpc.DiffResponse{
		Changes:         rpc.DiffResponse_DIFF_SOME,
		DetailedDiff:    diff,
		HasDetailedDiff: true,
	}, nil
}

func diffUpdates(updates map[resource.PropertyKey]resource.ValueDiff) map[string]*rpc.PropertyDiff {
	updateDiff := map[string]*rpc.PropertyDiff{}
	for key, valueDiff := range updates {
		if string(key) != "registry" {
			updateDiff[string(key)] = &rpc.PropertyDiff{
				Kind: rpc.PropertyDiff_UPDATE,
			}
		} else {
			// only register a diff on "server" field, but not on "username" or "password",
			// as they can change frequently and should not trigger a rebuild.
			serverDiff := valueDiff.Object.Updates["server"]
			// if serverDiff is not empty, we register a property diff update
			if serverDiff != (resource.ValueDiff{}) {
				updateDiff[string(key)] = &rpc.PropertyDiff{
					Kind: rpc.PropertyDiff_UPDATE,
				}
			}
		}
	}
	return updateDiff
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
	inputs := req.GetInputs()
	properties := req.GetProperties()

	// Return properties as passed, since we do no reconciliation,
	return &rpc.ReadResponse{Id: id, Inputs: inputs, Properties: properties}, nil

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

func (accumulator *contextHashAccumulator) hashPath(file string, fileMode fs.FileMode) error {
	hash := sha256.New()

	if fileMode.Type() == fs.ModeSymlink {
		// For symlinks, we hash the symlink _path_ instead of the file content.
		// This will allow us to:
		// a) ignore changes at the symlink target
		// b) detect if the symlink _itself_ changes
		// c) avoid a panic on io.Copy if the symlink target is a directory
		symLinkPath, err := filepath.EvalSymlinks(filepath.Join(accumulator.dockerContextPath, file))
		if err != nil {
			return fmt.Errorf("could not evaluate symlink at %s: %w", file, err)
		}
		// Hashed content is the clean, os-agnostic file path:
		_, err = io.Copy(hash, strings.NewReader(path.Clean(symLinkPath)))
		if err != nil {
			return fmt.Errorf("could not copy symlink path %s to hash: %w", file, err)
		}
	} else {
		// For regular files, we can hash their content.
		// TODO: consider only hashing file metadata to improve performance
		f, err := os.Open(filepath.Join(accumulator.dockerContextPath, file))
		if err != nil {
			return fmt.Errorf("could not open file %s: %w", file, err)
		}
		defer f.Close()
		_, err = io.Copy(hash, f)
		if err != nil {
			return fmt.Errorf("could not copy file %s to hash: %w", file, err)
		}
	}

	// Capture all information in the accumulator buffer and add a separator
	accumulator.input.Write([]byte(filepath.Clean(file))) // use os-agnostic filepath
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
	dockerIgnorePath := filepath.Join(dockerContextPath, mapDockerignore(filepath.Base(dockerfile)))
	ignorePatterns, err := getIgnore(dockerIgnorePath)
	if err != nil {
		return "", err
	}
	ignoreMatcher, err := patternmatcher.New(ignorePatterns)
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
		ignore, err := ignoreMatcher.MatchesOrParentMatches(path)
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

func getIgnore(dockerIgnorePath string) ([]string, error) {
	var ignorePatterns []string
	dockerIgnore, err := os.ReadFile(dockerIgnorePath)
	if err != nil {
		if os.IsNotExist(err) {
			// .dockerignore files are optional
			return ignorePatterns, nil
		}
		return ignorePatterns, fmt.Errorf("unable to read %s file: %w", dockerIgnorePath, err)
	}
	ignorePatterns, err = dockerignore.ReadAll(bytes.NewReader(dockerIgnore))
	if err != nil {
		return ignorePatterns, fmt.Errorf("unable to parse %s file: %w", ".dockerignore", err)
	}
	return ignorePatterns, nil
}

func getRelDockerfilePath(buildContext, dockerfile string) (string, error) {
	// if the Pulumi program specifies an absolute path or a path relative to the program's local directory,
	// we need to get the Dockerfile's relative path to the context directory for the hash function
	if strings.Contains(dockerfile, string(filepath.Separator)) {
		absDockerfile, err := filepath.Abs(dockerfile)
		if err != nil {
			return "", fmt.Errorf("absDockerfile error: %s", err)
		}
		absBuildpath, err := filepath.Abs(buildContext)
		if err != nil {
			return "", fmt.Errorf("absBuildPath error: %s", err)
		}
		dockerfile, err = filepath.Rel(absBuildpath, absDockerfile)
		if err != nil {
			return "", fmt.Errorf("relDockerfile error: %s", err)
		}

	}
	return dockerfile, nil
}
