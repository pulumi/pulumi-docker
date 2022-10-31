package provider

import (
	"bufio"
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
	"github.com/docker/docker/pkg/fileutils"
	"github.com/golang/protobuf/ptypes/empty"
	pbempty "github.com/golang/protobuf/ptypes/empty"
	structpb "github.com/golang/protobuf/ptypes/struct"
	"github.com/moby/buildkit/frontend/dockerfile/dockerignore"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi/pkg/v3/resource/provider"
	"github.com/pulumi/pulumi/sdk/v3/go/common/diag"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/cmdutil"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
	rpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
)

// Track a list of native resource tokens
var dockerImageToken = "index/image.Image"

// Hybrid provider struct
type dockerHybridProvider struct {
	schemaBytes     []byte
	version         string
	bridgedProvider rpc.ResourceProviderServer
	nativeProvider  rpc.ResourceProviderServer
}

type dockerNativeProvider struct {
	host        *provider.HostClient
	name        string
	version     string
	schemaBytes []byte
	loginLock   sync.Mutex
}

// TODO: implement this somewhere that makes sense Serve launches the gRPC server for the resource provider.
func Serve(providerName, version string, schemaBytes []byte) {
	// Start gRPC service.
	err := provider.Main(providerName, func(host *provider.HostClient) (rpc.ResourceProviderServer, error) {
		return makeProvider(host, providerName, version, schemaBytes)
	})
	if err != nil {
		cmdutil.ExitError(err.Error())
	}
}

func makeProvider(host *provider.HostClient, name, version string, schemaBytes []byte) (rpc.ResourceProviderServer, error) {
	nativeProvider := &dockerNativeProvider{
		host:        host,
		name:        name,
		version:     version,
		schemaBytes: schemaBytes,
	}

	prov := Provider()
	bridgedProvider := tfbridge.NewProvider(context.Background(), host, name, version, prov.P, prov, schemaBytes)
	return &dockerHybridProvider{
		schemaBytes:     schemaBytes,
		version:         version,
		bridgedProvider: bridgedProvider,
		nativeProvider:  nativeProvider,
	}, nil
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
func (p *dockerNativeProvider) Construct(ctx context.Context, req *rpc.ConstructRequest) (*rpc.ConstructResponse, error) {
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
	return &rpc.ConfigureResponse{}, nil
}

// Invoke dynamically executes a built-in function in the provider.
func (p *dockerNativeProvider) Invoke(_ context.Context, req *rpc.InvokeRequest) (*rpc.InvokeResponse, error) {
	tok := req.GetTok()
	return nil, fmt.Errorf("unknown Invoke token '%s'", tok)
}

// StreamInvoke dynamically executes a built-in function in the provider. The result is streamed
// back as a series of messages.
func (p *dockerNativeProvider) StreamInvoke(req *rpc.InvokeRequest, server rpc.ResourceProvider_StreamInvokeServer) error {
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

	return &rpc.CheckResponse{Inputs: req.News, Failures: nil}, nil
}

// Diff checks what impacts a hypothetical update will have on the resource's properties.
func (p *dockerNativeProvider) Diff(ctx context.Context, req *rpc.DiffRequest) (*rpc.DiffResponse, error) {
	urn := resource.URN(req.GetUrn())
	label := fmt.Sprintf("%s.Diff(%s)", p.name, urn)
	logging.V(9).Infof("%s executing", label)

	olds, err := plugin.UnmarshalProperties(req.GetOlds(), plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true})
	if err != nil {
		return nil, err
	}

	news, err := plugin.UnmarshalProperties(req.GetNews(), plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true})
	if err != nil {
		return nil, err
	}

	d := olds.Diff(news)
	changes := rpc.DiffResponse_DIFF_NONE

	// Replace the below condition with logic specific to your provider
	if d.Changed("length") {
		changes = rpc.DiffResponse_DIFF_SOME
	}

	return &rpc.DiffResponse{
		Changes:  changes,
		Replaces: []string{"length"},
	}, nil
}

// Create allocates a new instance of the provided resource and returns its unique ID afterwards.
func (p *dockerNativeProvider) Create(ctx context.Context, req *rpc.CreateRequest) (*rpc.CreateResponse, error) {
	urn := resource.URN(req.GetUrn())
	label := fmt.Sprintf("%s.Create(%s)", p.name, urn)
	logging.V(9).Infof("%s executing", label)

	// TODO: see if we need the input validation
	//inputs, err := plugin.UnmarshalProperties(req.GetProperties(), plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true})
	//if err != nil {
	//	return nil, err
	//}
	//
	//// Replace the below random number implementation with logic specific to your provider
	//if !inputs["length"].IsNumber() {
	//	return nil, fmt.Errorf("expected input property 'length' of type 'number' but got '%s", inputs["length"].TypeString())
	//}

	outputProperties, err := p.dockerBuildGo(ctx, urn, req.GetProperties())
	if err != nil {
		return nil, err
	}
	return &rpc.CreateResponse{
		Id:         "ignored",
		Properties: outputProperties,
	}, nil

	//outputProperties, err := plugin.MarshalProperties(
	//	resource.NewPropertyMapFromMap(outputs),
	//	plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true},
	//)
	//if err != nil {
	//	return nil, err
	//}
	//return &rpc.CreateResponse{
	//	Id:         result,
	//	Properties: outputProperties,
	//}, nil
}

func (p *dockerNativeProvider) dockerBuildGo(ctx context.Context,
	urn resource.URN,
	props *structpb.Struct) (*structpb.Struct, error) {
	fmt.Println("resource URN", urn)

	fmt.Println("getting inputs...")

	inputs, err := plugin.UnmarshalProperties(props, plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true})
	if err != nil {
		return nil, err
	}

	dockerfile := inputs["dockerfile"].StringValue()
	buildContext := inputs["context"].StringValue()
	registryURL := inputs["registryURL"].StringValue()
	imageName := inputs["name"].StringValue()
	//tag := inputs["tag"].StringValue()
	registry := inputs["registry"].ObjectValue()
	username := registry["username"].StringValue()
	password := registry["password"].StringValue()

	fmt.Println("USING THE DOCKER CLIENT NOW")

	cli, err := client.NewClientWithOpts(client.FromEnv)

	if err != nil {
		panic(err)
	}
	//fmt.Println("Get list of  existing images")
	//imgs, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println("List image labels")
	//for _, img := range imgs {
	//	fmt.Println("whole image: ", img.RepoTags) // RepoTags gives a list of tags that correspond to the REPOSITORY and TAGS headers of the "docker images" output
	//}

	fmt.Println("now let's BUILD an image")

	// make the build context
	tar, err := archive.TarWithOptions(buildContext, &archive.TarOptions{})
	if err != nil {
		panic(err)
	}

	// make the build options
	opts := types.ImageBuildOptions{
		Dockerfile: dockerfile,
		Tags:       []string{"gsaenger/" + imageName}, // TODO: definitely do not hardcode this - must be the registry namespace, for dockerhub this is the username
		Remove:     true,
	}

	imgBuildResp, err := cli.ImageBuild(context.Background(), tar, opts)
	if err != nil {
		panic(err)
	}
	// Print build logs to terminal
	scanner := bufio.NewScanner(imgBuildResp.Body)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	fmt.Println("see if we can push to the registry")

	// Quick and dirty auth; we can also preconfigure the client itself I believe

	var authConfig = types.AuthConfig{
		Username:      username,
		Password:      password,
		ServerAddress: registryURL,
	}

	authConfigBytes, _ := json.Marshal(authConfig)
	authConfigEncoded := base64.URLEncoding.EncodeToString(authConfigBytes)

	pushOpts := types.ImagePushOptions{RegistryAuth: authConfigEncoded}

	pushOutput, err := cli.ImagePush(context.Background(), "gsaenger/"+imageName, pushOpts)

	if err != nil {
		panic(err)
	}

	defer pushOutput.Close()

	// Print push logs to terminal
	pushScanner := bufio.NewScanner(pushOutput)
	for pushScanner.Scan() {
		fmt.Println(pushScanner.Text())
	}

	outputs := map[string]interface{}{
		"dockerfile":  dockerfile,
		"context":     buildContext,
		"imageName":   imageName,
		"registryURL": registryURL,
	}
	return plugin.MarshalProperties(
		resource.NewPropertyMapFromMap(outputs),
		plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true},
	)
}

func (p *dockerNativeProvider) dockerBuild(
	ctx context.Context,
	urn resource.URN,
	props *structpb.Struct,
) (*structpb.Struct, error) {
	inputs, err := plugin.UnmarshalProperties(props, plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true})
	if err != nil {
		return nil, err
	}
	applyDefaults(inputs)
	name := inputs["name"].StringValue()
	baseName := strings.Split(name, ":")[0]
	context := inputs["context"].StringValue()
	dockerfile := inputs["dockerfile"].StringValue()
	target := inputs["target"].StringValue()
	registry := inputs["registry"].ObjectValue()
	username := registry["username"]
	password := registry["password"]

	contextDigest, err := hashContext(context, dockerfile)
	if err != nil {
		return nil, err
	}

	if !username.IsNull() && !password.IsNull() {
		cmd := exec.Command(
			"docker", "login",
			"-u", username.StringValue(), "--password-stdin",
			registry["server"].StringValue(),
		)
		cmd.Stdin = strings.NewReader(password.StringValue())
		// On macOS, it seems simultaneous invocations of `docker login` can
		// fail. See #6. Use a lock to prevent multiple `dockerBuild` requests
		// from calling `docker login` simultaneously.
		p.loginLock.Lock()
		err := runCommand(ctx, p.host, urn, cmd)
		p.loginLock.Unlock()
		if err != nil {
			return nil, fmt.Errorf("docker login failed: %w", err)
		}
	}

	var platforms []string
	for _, v := range inputs["platforms"].ArrayValue() {
		platforms = append(platforms, v.StringValue())
	}

	args := []string{
		"buildx", "build",
		"--platform", strings.Join(platforms, ","),
		"--cache-from", name,
		"--cache-to", "type=inline",
		"-f", filepath.Join(context, dockerfile),
		"--target", target,
		"-t", name, "--push",
	}
	if !inputs["args"].IsNull() {
		for _, v := range inputs["args"].ArrayValue() {
			name := v.ObjectValue()["name"].StringValue()
			value := v.ObjectValue()["value"].StringValue()
			args = append(args, "--build-arg", fmt.Sprintf("%s=%s", name, value))
		}
	}
	args = append(args, context)
	cmd := exec.Command("docker", args...)
	if err := runCommand(ctx, p.host, urn, cmd); err != nil {
		return nil, fmt.Errorf("docker build failed: %w", err)
	}

	cmd = exec.Command("docker", "inspect", name, "-f", `{{join .RepoDigests "\n"}}`)
	repoDigests, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("docker inspect failed: %s: %s", err, string(repoDigests))
	}
	var repoDigest string
	for _, line := range strings.Split(string(repoDigests), "\n") {
		repo := strings.Split(line, "@")[0]
		if repo == baseName {
			repoDigest = line
			break
		}
	}
	if repoDigest == "" {
		return nil, fmt.Errorf("failed to find repo digest in docker inspect output: %s", repoDigests)
	}

	outputs := map[string]interface{}{
		"dockerfile":     dockerfile,
		"context":        context,
		"target":         target,
		"name":           name,
		"platforms":      platforms,
		"contextDigest":  contextDigest,
		"repoDigest":     repoDigest,
		"registryServer": registry["server"].StringValue(),
	}
	return plugin.MarshalProperties(
		resource.NewPropertyMapFromMap(outputs),
		plugin.MarshalOptions{KeepUnknowns: true, SkipNulls: true},
	)
}

func applyDefaults(inputs resource.PropertyMap) {
	if inputs["platforms"].IsNull() {
		inputs["platforms"] = resource.NewArrayProperty(
			[]resource.PropertyValue{resource.NewStringProperty("linux/amd64")},
		)
	}
}

func runCommand(
	ctx context.Context,
	host *provider.HostClient,
	urn resource.URN,
	cmd *exec.Cmd,
) error {
	cmd.Stdout = &logWriter{
		ctx:      ctx,
		host:     host,
		urn:      urn,
		severity: diag.Info,
	}
	cmd.Stderr = &logWriter{
		ctx:      ctx,
		host:     host,
		urn:      urn,
		severity: diag.Info,
	}
	return cmd.Run()
}

type logWriter struct {
	ctx      context.Context
	host     *provider.HostClient
	urn      resource.URN
	severity diag.Severity
}

func (w *logWriter) Write(p []byte) (n int, err error) {
	return len(p), w.host.Log(w.ctx, w.severity, w.urn, string(p))
}

type contextHash struct {
	contextPath string
	input       bytes.Buffer
}

func newContextHash(contextPath string) *contextHash {
	return &contextHash{contextPath: contextPath}
}

func (ch *contextHash) hashPath(path string, fileMode fs.FileMode) error {
	f, err := os.Open(filepath.Join(ch.contextPath, path))
	if err != nil {
		return fmt.Errorf("open %s: %w", path, err)
	}
	defer f.Close()
	h := sha256.New()
	_, err = io.Copy(h, f)
	if err != nil {
		return fmt.Errorf("read %s: %w", path, err)
	}
	ch.input.Write([]byte(path))
	ch.input.Write([]byte(fileMode.String()))
	ch.input.Write(h.Sum(nil))
	ch.input.WriteByte(0)
	return nil
}

func (ch *contextHash) hexSum() string {
	h := sha256.New()
	ch.input.WriteTo(h)
	return hex.EncodeToString(h.Sum(nil))
}

func hashContext(contextPath string, dockerfile string) (string, error) {
	dockerIgnorePath := dockerfile + ".dockerignore"
	dockerIgnore, err := os.ReadFile(dockerIgnorePath)
	if err != nil {
		if os.IsNotExist(err) {
			dockerIgnorePath = filepath.Join(contextPath, ".dockerignore")
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
	ch := newContextHash(contextPath)
	err = ch.hashPath(dockerfile, 0)
	if err != nil {
		return "", fmt.Errorf("hashing dockerfile %q: %w", dockerfile, err)
	}
	err = filepath.WalkDir(contextPath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		path, err = filepath.Rel(contextPath, path)
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
			} else {
				return nil
			}
		} else if d.IsDir() {
			return nil
		}
		info, err := d.Info()
		if err != nil {
			return fmt.Errorf("determining mode for %q: %w", path, err)
		}
		err = ch.hashPath(path, info.Mode())
		if err != nil {
			return fmt.Errorf("hashing %q: %w", path, err)
		}
		return nil
	})
	if err != nil {
		return "", fmt.Errorf("unable to hash build context: %w", err)
	}
	return ch.hexSum(), nil
}

// TODO: these are the remaining methods

// Read the current live state associated with a resource.
func (p *dockerNativeProvider) Read(ctx context.Context, req *rpc.ReadRequest) (*rpc.ReadResponse, error) {
	urn := resource.URN(req.GetUrn())
	label := fmt.Sprintf("%s.Read(%s)", p.name, urn)
	logging.V(9).Infof("%s executing", label)
	msg := fmt.Sprintf("Read is not yet implemented for %s", urn.Type())
	return nil, status.Error(codes.Unimplemented, msg)
}

// Update updates an existing resource with new values.
func (p *dockerNativeProvider) Update(ctx context.Context, req *rpc.UpdateRequest) (*rpc.UpdateResponse, error) {
	urn := resource.URN(req.GetUrn())
	label := fmt.Sprintf("%s.Update(%s)", p.name, urn)
	logging.V(9).Infof("%s executing", label)
	// Our example Random resource will never be updated - if there is a diff, it will be a replacement.
	msg := fmt.Sprintf("Update is not yet implemented for %s", urn.Type())
	return nil, status.Error(codes.Unimplemented, msg)
}

// Delete tears down an existing resource with the given ID.  If it fails, the resource is assumed
// to still exist.
func (p *dockerNativeProvider) Delete(ctx context.Context, req *rpc.DeleteRequest) (*pbempty.Empty, error) {
	urn := resource.URN(req.GetUrn())
	label := fmt.Sprintf("%s.Update(%s)", p.name, urn)
	logging.V(9).Infof("%s executing", label)
	// Implement Delete logic specific to your provider.
	// Note that for our Random resource, we don't have to do anything on Delete.
	return &pbempty.Empty{}, nil
}

// GetPluginInfo returns generic information about this plugin, like its version.
func (p *dockerNativeProvider) GetPluginInfo(context.Context, *pbempty.Empty) (*rpc.PluginInfo, error) {
	return &rpc.PluginInfo{
		Version: p.version,
	}, nil
}

// GetSchema returns the JSON-serialized schema for the provider.
func (p *dockerNativeProvider) GetSchema(ctx context.Context, req *rpc.GetSchemaRequest) (*rpc.GetSchemaResponse, error) {
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

// TODO: this is for functionsd AKA data sources, and our provider doesn't have any metadatea becuase we're just implementing from scratch
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
