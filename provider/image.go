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
	"github.com/docker/docker/builder/dockerignore"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
	"github.com/docker/docker/pkg/fileutils"
	structpb "github.com/golang/protobuf/ptypes/struct"
	"github.com/pulumi/pulumi/pkg/v3/resource/provider"
	"github.com/pulumi/pulumi/sdk/v3/go/common/diag"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

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
