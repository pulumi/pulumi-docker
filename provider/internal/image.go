package internal

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	// These imports are needed to register the drivers with buildkit.
	_ "github.com/docker/buildx/driver/docker-container"
	_ "github.com/docker/buildx/driver/kubernetes"
	_ "github.com/docker/buildx/driver/remote"

	"github.com/distribution/reference"
	controllerapi "github.com/docker/buildx/controller/pb"
	"github.com/docker/buildx/util/buildflags"
	"github.com/docker/buildx/util/platformutil"
	"github.com/docker/docker/errdefs"
	"github.com/moby/buildkit/exporter/containerimage/exptypes"
	"github.com/muesli/reflow/dedent"
	"github.com/opencontainers/go-digest"

	provider "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/diag"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"

	"github.com/pulumi/pulumi-docker/provider/v4/internal/properties"
)

var (
	_ infer.Annotated                             = (infer.Annotated)((*Image)(nil))
	_ infer.Annotated                             = (infer.Annotated)((*ImageArgs)(nil))
	_ infer.Annotated                             = (infer.Annotated)((*ImageState)(nil))
	_ infer.CustomCheck[ImageArgs]                = (*Image)(nil)
	_ infer.CustomDelete[ImageState]              = (*Image)(nil)
	_ infer.CustomDiff[ImageArgs, ImageState]     = (*Image)(nil)
	_ infer.CustomRead[ImageArgs, ImageState]     = (*Image)(nil)
	_ infer.CustomResource[ImageArgs, ImageState] = (*Image)(nil)
	_ infer.CustomUpdate[ImageArgs, ImageState]   = (*Image)(nil)
)

// Image is a Docker image build using buildkit.
type Image struct{}

// Annotate provides a description of the Image resource.
func (i *Image) Annotate(a infer.Annotator) {
	a.Describe(&i, "A Docker image built using Buildkit")
}

// ImageArgs instantiates a new Image.
type ImageArgs struct {
	BuildArgs      map[string]string         `pulumi:"buildArgs,optional"`
	Builder        string                    `pulumi:"builder,optional"`
	BuildOnPreview bool                      `pulumi:"buildOnPreview,optional"`
	CacheFrom      []string                  `pulumi:"cacheFrom,optional"`
	CacheTo        []string                  `pulumi:"cacheTo,optional"`
	Context        string                    `pulumi:"context,optional"`
	Exports        []string                  `pulumi:"exports,optional"`
	File           string                    `pulumi:"file,optional"`
	Platforms      []string                  `pulumi:"platforms,optional"`
	Pull           bool                      `pulumi:"pull,optional"`
	Registries     []properties.RegistryAuth `pulumi:"registries,optional"`
	Tags           []string                  `pulumi:"tags"`
}

// Annotate describes inputs to the Image resource.
func (ia *ImageArgs) Annotate(a infer.Annotator) {
	a.Describe(&ia.BuildArgs, dedent.String(`
		An optional map of named build-time argument variables to set during
		the Docker build. This flag allows you to pass build-time variables that
		can be accessed like environment variables inside the RUN
		instruction.`,
	))
	a.Describe(&ia.Builder, dedent.String(`
		Build with a specific builder instance`,
	))
	a.Describe(&ia.BuildOnPreview, dedent.String(`
		When true, attempt to build the image during previews. Outputs are not
		pushed to registries, however caches are still populated.
	`))
	a.Describe(&ia.CacheFrom, dedent.String(`
		External cache sources (e.g., "user/app:cache", "type=local,src=path/to/dir")`,
	))
	a.Describe(&ia.CacheTo, dedent.String(`
		Cache export destinations (e.g., "user/app:cache", "type=local,dest=path/to/dir")`,
	))
	a.Describe(&ia.Context, dedent.String(`
		Path to use for build context. If omitted, an empty context is used.`,
	))
	a.Describe(&ia.Exports, dedent.String(`
		Name and optionally a tag (format: "name:tag"). If outputting to a
		registry, the name should include the fully qualified registry address.`,
	))
	a.Describe(&ia.File, dedent.String(`
		Name of the Dockerfile to use (defaults to "${context}/Dockerfile").`,
	))
	a.Describe(&ia.Platforms, dedent.String(`
		Set target platforms for the build. Defaults to the host's platform`,
	))
	a.Describe(&ia.Pull, dedent.String(`
		Always attempt to pull all referenced images`,
	))
	a.Describe(&ia.Tags, dedent.String(`
		Name and optionally a tag (format: "name:tag"). If outputting to a
		registry, the name should include the fully qualified registry address.`,
	))
	a.Describe(&ia.Registries, dedent.String(`
		Logins for registry outputs`,
	))

	a.SetDefault(&ia.File, "Dockerfile")
}

// ImageState is serialized to the program's state file.
type ImageState struct {
	ImageArgs

	ContextHash string                `pulumi:"contextHash,optional" provider:"internal"`
	Manifests   []properties.Manifest `pulumi:"manifests" provider:"output"`

	id string
}

// Annotate describes outputs of the Image resource.
func (is *ImageState) Annotate(a infer.Annotator) {
	is.ImageArgs.Annotate(a)
}

// Check validates ImageArgs, sets defaults, and ensures our client is
// authenticated.
func (*Image) Check(
	ctx provider.Context,
	_ string,
	_ resource.PropertyMap,
	news resource.PropertyMap,
) (ImageArgs, []provider.CheckFailure, error) {
	args, failures, err := infer.DefaultCheck[ImageArgs](news)
	if err != nil || len(failures) != 0 {
		return args, failures, err
	}

	// :(
	preview := news.ContainsUnknowns()

	if _, berr := args.toBuildOptions(preview); berr != nil {
		errs := berr.(interface{ Unwrap() []error }).Unwrap()
		for _, e := range errs {
			if cf, ok := e.(checkFailure); ok {
				failures = append(failures, cf.CheckFailure)
			}
		}
	}

	// Check is called before every operation except Read, so this ensures
	// we're authenticated in almost all cases.
	cfg := infer.GetConfig[Config](ctx)
	for _, reg := range args.Registries {
		// TODO(https://github.com/pulumi/pulumi-go-provider/pull/155): This is likely unresolved.
		if reg.Address == "" {
			continue
		}
		if err = cfg.client.Auth(ctx, reg); err != nil {
			failures = append(failures,
				provider.CheckFailure{Property: "registries", Reason: fmt.Sprintf("unable to authenticate: %s", err.Error())})
		}
	}

	return args, failures, err
}

type checkFailure struct {
	provider.CheckFailure
}

func (cf checkFailure) Error() string {
	return cf.Reason
}

func newCheckFailure(property string, err error) checkFailure {
	return checkFailure{provider.CheckFailure{Property: property, Reason: err.Error()}}
}

func (ia *ImageArgs) withoutUnknowns(preview bool) ImageArgs {
	sk := stringKeeper{preview}
	filtered := ImageArgs{
		BuildArgs:      mapKeeper{preview}.keep(ia.BuildArgs),
		Builder:        ia.Builder,
		BuildOnPreview: ia.BuildOnPreview,
		CacheFrom:      filter(sk, ia.CacheFrom...),
		CacheTo:        filter(sk, ia.CacheTo...),
		Context:        ia.Context,
		Exports:        filter(sk, ia.Exports...),
		File:           ia.File, //
		Platforms:      filter(sk, ia.Platforms...),
		Pull:           ia.Pull,
		Registries:     filter(registryKeeper{preview}, ia.Registries...),
		Tags:           filter(sk, ia.Tags...),
	}

	return filtered
}

func (ia *ImageArgs) buildable() bool {
	// We can build the given inputs if filtered out unknowns is a no-op.
	filtered := ia.withoutUnknowns(true)
	return reflect.DeepEqual(ia, &filtered)
}

func (ia ImageArgs) toBuilds(
	ctx provider.Context,
	preview bool,
) ([]controllerapi.BuildOptions, error) {
	opts, err := ia.toBuildOptions(preview)
	if err != nil {
		return nil, err
	}

	// Check if we need a workaround for https://github.com/docker/buildx/issues/1044
	if len(ia.CacheTo) == 0 || len(ia.Platforms) <= 1 {
		return []controllerapi.BuildOptions{opts}, nil
	}

	// Split the build into N pieces: one build with only local caching, and an
	// additional cache-only build for each platform.
	builds := []controllerapi.BuildOptions{}

	origCacheTo := opts.CacheTo

	// Build 1:
	// - No --cache-to.
	// - Extend --cache-from with platform-specific caches, while preserving existing ones.
	// - Preserve exports.
	opts.CacheTo = nil
	opts.CacheFrom = append(opts.CacheFrom, cachesFor(ctx, opts.CacheFrom, opts.Platforms...)...)
	builds = append(builds, opts)

	// Build 2..P for each platform:
	// - --output=type=cacheonly.
	// - No --cache-from (rely on local build cache).
	// - --cache-to
	for _, p := range opts.Platforms {
		// Only build for this platform.
		opts.Platforms = []string{p}
		// Only export caches.
		opts.Exports = []*controllerapi.ExportEntry{{Type: "cacheonly"}}
		// Rely on local build cache.
		opts.CacheFrom = nil
		// Cache to platform-aware tags.
		opts.CacheTo = cachesFor(ctx, origCacheTo, p)
		builds = append(builds, opts)
	}

	return builds, nil
}

// cachesFor is a workaround for https://github.com/docker/buildx/issues/1044
// which modifies the names of cache to/from entries to be platform-aware.
func cachesFor(
	ctx provider.Context,
	existing []*controllerapi.CacheOptionsEntry,
	platforms ...string,
) []*controllerapi.CacheOptionsEntry {
	caches := []*controllerapi.CacheOptionsEntry{}

	// Iterate over existing cache entries first to preserve precedence.
	for _, c := range existing {
		for _, p := range platforms {
			entry := &controllerapi.CacheOptionsEntry{
				Type:  c.Type,
				Attrs: make(map[string]string),
			}
			for k, v := range c.Attrs {
				entry.Attrs[k] = v
			}
			plat := strings.Replace(p, "/", "-", -1)

			switch c.Type {
			case "gha":
				entry.Attrs["scope"] += "-" + plat
			case "s3", "azblob":
				if entry.Attrs["name"] != "" {
					entry.Attrs["name"] += "-" + plat
				} else {
					entry.Attrs["name"] = plat
				}
			case "registry":
				// We don't want these synthetic caches to show up as tags on
				// registries, so instead we identify them as opaque blobs
				// whose names are derived from the base ref + platform.
				h := sha256.New()
				h.Write([]byte(entry.Attrs["ref"]))
				h.Write([]byte(p))
				dgst := digest.NewDigest(digest.SHA256, h)

				ref, err := reference.Parse(entry.Attrs["ref"])
				if err != nil {
					ctx.Log(
						diag.Warning,
						fmt.Errorf("Unable to parse cache entry: %w", err).Error(),
					)
					continue
				}

				if named, ok := ref.(reference.Named); ok {
					named = reference.TrimNamed(named)
					ref, _ = reference.WithDigest(named, dgst) // Can't error.
				} else {
					named, err := reference.WithName(ref.String())
					if err != nil {
						ctx.Log(diag.Warning, fmt.Errorf("Unable to build cache key: %w", err).Error())
						continue
					}
					ref, _ = reference.WithDigest(named, dgst) // Can't error.
				}

				entry.Attrs["ref"] = ref.String()
			default:
			}
			caches = append(caches, entry)
		}
	}
	return caches
}

func (ia *ImageArgs) toBuildOptions(preview bool) (controllerapi.BuildOptions, error) {
	var multierr error

	if len(ia.Tags) == 0 {
		multierr = errors.Join(multierr, newCheckFailure("tags", errors.New("at least one tag is required")))
	}

	// TODO(https://github.com/pulumi/pulumi-docker/issues/860): Empty build context
	if ia.Context != "" && !preview {
		if ia.File == "" {
			ia.File = filepath.Join(ia.Context, "Dockerfile")
		}
		if _, err := os.Stat(ia.File); err != nil {
			multierr = errors.Join(multierr, newCheckFailure("context", fmt.Errorf("%q: %w", ia.File, err)))
		}
	}

	// Discard any unknown inputs if this is a preview -- we don't want them to
	// cause validation errors.
	filtered := ia.withoutUnknowns(preview)

	exports, err := buildflags.ParseExports(filtered.Exports)
	if err != nil {
		multierr = errors.Join(multierr, newCheckFailure("exports", err))
	}
	if preview {
		// Don't perform registry pushes during previews.
		for _, e := range exports {
			if e.Type == "image" && e.Attrs["push"] == "true" {
				e.Attrs["push"] = "false"
			}
		}
	}

	_, err = platformutil.Parse(filtered.Platforms)
	if err != nil {
		multierr = errors.Join(multierr, newCheckFailure("platforms", err))
	}

	cacheFrom, err := buildflags.ParseCacheEntry(filtered.CacheFrom)
	if err != nil {
		multierr = errors.Join(multierr, newCheckFailure("cacheFrom", err))
	}

	cacheTo, err := buildflags.ParseCacheEntry(filtered.CacheTo)
	if err != nil {
		multierr = errors.Join(multierr, newCheckFailure("cacheTo", err))
	}

	for _, t := range filtered.Tags {
		if _, err := reference.Parse(t); err != nil {
			multierr = errors.Join(multierr, newCheckFailure("tags", err))
		}
	}

	opts := controllerapi.BuildOptions{
		BuildArgs:      filtered.BuildArgs,
		Builder:        filtered.Builder,
		CacheFrom:      cacheFrom,
		CacheTo:        cacheTo,
		ContextPath:    filtered.Context,
		DockerfileName: filtered.File,
		Exports:        exports,
		Platforms:      filtered.Platforms,
		Pull:           filtered.Pull,
		Tags:           filtered.Tags,
	}

	return opts, multierr
}

// Update builds the image using buildkit.
func (i *Image) Update(
	ctx provider.Context,
	name string,
	state ImageState,
	input ImageArgs,
	preview bool,
) (ImageState, error) {
	cfg := infer.GetConfig[Config](ctx)

	state.ImageArgs = input

	ok, err := cfg.client.BuildKitEnabled()
	if err != nil {
		return state, fmt.Errorf("checking buildkit compatibility: %w", err)
	}
	if !ok {
		return state, fmt.Errorf("buildkit is not supported on this host")
	}

	builds, err := input.toBuilds(ctx, preview)
	if err != nil {
		return state, fmt.Errorf("preparing: %w", err)
	}

	hash, err := HashContext(input.Context, input.File)
	if err != nil {
		return state, fmt.Errorf("hashing build context: %w", err)
	}
	state.ContextHash = hash

	if preview && !input.BuildOnPreview {
		return state, nil
	}
	if preview && !input.buildable() {
		ctx.Log(diag.Warning, "Skipping preview build because some inputs are unknown.")
		return state, nil
	}

	var id string

	for _, b := range builds {
		result, err := cfg.client.Build(ctx, b)
		if err != nil {
			return state, fmt.Errorf("building %q: %w", input.Tags, err)
		}
		if id != "" {
			continue
		}

		if digest, ok := result.ExporterResponse["containerimage.digest"]; ok {
			id = digest
		} else if digest, ok := result.ExporterResponse[exptypes.ExporterImageConfigDigestKey]; ok {
			id = digest
		} else if tags, ok := result.ExporterResponse["image.name"]; ok {
			id = tags
		} else {
			id = name
		}
	}

	// TODO: Handle case with no export.
	_, _, state, err = i.Read(ctx, id, input, state)

	return state, err
}

// Create initializes a new resource and performs an Update on it.
func (i *Image) Create(
	ctx provider.Context,
	name string,
	input ImageArgs,
	preview bool,
) (string, ImageState, error) {
	state := ImageState{
		Manifests: []properties.Manifest{},
	}
	state, err := i.Update(ctx, name, state, input, preview)
	return state.id, state, err
}

// Read attempts to read manifests from an image's exports. An image without
// exports will have no manifests.
func (*Image) Read(
	ctx provider.Context,
	id string,
	input ImageArgs,
	state ImageState,
) (
	string, // id
	ImageArgs, // normalized inputs
	ImageState, // normalized state
	error,
) {
	opts, err := input.toBuildOptions(false)
	if err != nil {
		return id, input, state, err
	}

	// Ensure we're authenticated.
	cfg := infer.GetConfig[Config](ctx)
	for _, reg := range input.Registries {
		if err = cfg.client.Auth(ctx, reg); err != nil {
			return id, input, state, err
		}
	}

	expectManifest := false
	manifests := []properties.Manifest{}
	for _, export := range opts.Exports {
		switch export.GetType() {
		case "image":
			if export.GetAttrs()["push"] != "true" {
				// No manifest to read if we didn't push.
				continue
			}

			for _, tag := range input.Tags {
				expectManifest = true
				infos, err := cfg.client.Inspect(ctx, tag)
				if err != nil {
					continue
				}
				for _, m := range infos {
					if m.Descriptor.Platform != nil && m.Descriptor.Platform.Architecture == "unknown" {
						// Ignore cache manifests.
						continue
					}
					if m.Ref == nil {
						// Shouldn't happen, but just in case.
						continue
					}
					manifests = append(manifests, properties.Manifest{
						Digest: m.Descriptor.Digest.String(),
						Platform: properties.Platform{
							OS:           m.Descriptor.Platform.OS,
							Architecture: m.Descriptor.Platform.Architecture,
						},
						Ref:  m.Ref.String(),
						Size: m.Descriptor.Size,
					})
				}
			}
		case "docker":
			// TODO: Inspect local image and munge it into a manifest.
		default:
			// Other export types (e.g. file) are not supported.
			continue
		}
	}

	// If we couldn't find the tags we expected then return an empty ID to
	// delete the resource.
	if expectManifest && len(manifests) == 0 {
		return "", input, state, nil
	}

	state.id = id
	state.Manifests = manifests

	return id, input, state, nil
}

// Delete deletes an Image. If the Image was already deleted out-of-band it is treated as a success.
//
// Any tags previously pushed to registries will not be deleted.
func (*Image) Delete(
	ctx provider.Context,
	id string,
	_ ImageState,
) error {
	cfg := infer.GetConfig[Config](ctx)

	deletions, err := cfg.client.Delete(ctx.(context.Context), id)
	if errdefs.IsNotFound(err) {
		return nil // Nothing to do.
	}

	for _, d := range deletions {
		if d.Deleted != "" {
			ctx.Log(diag.Info, d.Deleted)
		}
		if d.Untagged != "" {
			ctx.Log(diag.Info, d.Untagged)
		}
	}

	// TODO: Delete tags from registries?

	return err
}

// Diff re-implements most of the default diff behavior, with the exception of
// ignoring "password" changes on registry inputs.
func (*Image) Diff(_ provider.Context, id string, olds ImageState, news ImageArgs) (provider.DiffResponse, error) {
	diff := map[string]provider.PropertyDiff{}
	update := provider.PropertyDiff{Kind: provider.Update}

	if !reflect.DeepEqual(olds.BuildArgs, news.BuildArgs) {
		diff["buildArgs"] = update
	}
	if !reflect.DeepEqual(olds.CacheFrom, news.CacheFrom) {
		diff["cacheFrom"] = update
	}
	if !reflect.DeepEqual(olds.CacheTo, news.CacheTo) {
		diff["cacheTo"] = update
	}
	if olds.Context != news.Context {
		diff["context"] = update
	}
	if !reflect.DeepEqual(olds.Exports, news.Exports) {
		diff["exports"] = update
	}
	if olds.File != news.File {
		diff["file"] = update
	}
	if !reflect.DeepEqual(olds.Platforms, news.Platforms) {
		diff["platforms"] = update
	}
	if olds.Pull != news.Pull {
		diff["pull"] = update
	}
	if !reflect.DeepEqual(olds.Tags, news.Tags) {
		diff["tags"] = update
	}

	// Check if anything has changed in our build context.
	hash, err := HashContext(news.Context, news.File)
	if err != nil {
		return provider.DiffResponse{}, err
	}
	if hash != olds.ContextHash {
		diff["manifests"] = update
	}

	// Registries need special handling because we ignore "password" changes to not introduce unnecessary changes.
	if len(olds.Registries) != len(news.Registries) {
		diff["registries"] = update
	} else {
		for idx, oldr := range olds.Registries {
			newr := news.Registries[idx]
			if (oldr.Username == newr.Username) && (oldr.Address == newr.Address) {
				continue
			}
			diff["registries"] = update
			break
		}
	}

	return provider.DiffResponse{
		DeleteBeforeReplace: false,
		HasChanges:          len(diff) > 0,
		DetailedDiff:        diff,
	}, nil
}

// Cancel cleans up temporary on-disk credentials.
func (*Image) Cancel(ctx provider.Context) error {
	cfg := infer.GetConfig[Config](ctx)
	return cfg.client.Close(ctx)
}
