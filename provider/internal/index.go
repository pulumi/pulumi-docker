package internal

import (
	"fmt"
	"reflect"
	"strings"

	// For examples/docs.
	_ "embed"

	provider "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

var (
	_ infer.Annotated                             = (infer.Annotated)((*Index)(nil))
	_ infer.Annotated                             = (infer.Annotated)((*IndexArgs)(nil))
	_ infer.Annotated                             = (infer.Annotated)((*IndexState)(nil))
	_ infer.CustomCheck[IndexArgs]                = (*Index)(nil)
	_ infer.CustomResource[IndexArgs, IndexState] = (*Index)(nil)
	_ infer.CustomDelete[IndexState]              = (*Index)(nil)
	_ infer.CustomDiff[IndexArgs, IndexState]     = (*Index)(nil)
	_ infer.CustomRead[IndexArgs, IndexState]     = (*Index)(nil)
	_ infer.CustomUpdate[IndexArgs, IndexState]   = (*Index)(nil)
)

//go:embed doc/index.md
var _indexExamples string

type Index struct{}

type IndexArgs struct {
	Tag      string       `pulumi:"tag"`
	Sources  []string     `pulumi:"sources"`
	Push     bool         `pulumi:"push,optional"`
	Registry RegistryAuth `pulumi:"registry,optional"`
}

type IndexState struct {
	IndexArgs

	Ref string `pulumi:"ref" provider:"output"`
}

func (i *Index) Annotate(a infer.Annotator) {
	a.Describe(&i, dedent(`
		An index (or manifest list) referencing one or more existing images.
		
		Useful for crafting a multi-platform image from several
		platform-specific images.

		This creates an OCI image index or a Docker manifest list depending on
		the media types of the source images.
	`)+
		"\n\n"+_indexExamples,
	)
}

func (i *IndexArgs) Annotate(a infer.Annotator) {
	a.Describe(&i.Registry, dedent(`
		Authentication for the registry where the tagged index will be pushed.

		Credentials can also be included with the provider's configuration.
	`))
	a.Describe(&i.Sources, dedent(`
		Existing images to include in the index.
	`))
	a.Describe(&i.Tag, dedent(`
		The tag to apply to the index.
	`))
	a.Describe(&i.Push, dedent(`
		If true, push the index to the target registry.

		Defaults to "true".
	`))

	a.SetDefault(&i.Push, true)
}

func (i *IndexState) Annotate(a infer.Annotator) {
	a.Describe(&i.Ref, dedent(`
		The pushed tag with digest.

		Identical to the tag if the index was not pushed.
	`))
}

func (i *Index) Create(
	ctx provider.Context,
	name string,
	input IndexArgs,
	preview bool,
) (string, IndexState, error) {
	state, err := i.Update(ctx, name, IndexState{}, input, preview)
	return name, state, err
}

func (i *Index) Update(
	ctx provider.Context,
	name string,
	state IndexState,
	input IndexArgs,
	preview bool,
) (IndexState, error) {
	state.IndexArgs = input
	state.Ref = input.Tag

	cli, err := i.client(ctx, state, input)
	if err != nil {
		return state, err
	}

	if preview {
		return state, nil
	}

	err = cli.ManifestCreate(ctx, input.Push, input.Tag, input.Sources...)
	if err != nil {
		return state, fmt.Errorf("creating: %w", err)
	}

	_, _, state, err = i.Read(ctx, name, input, state)
	if err != nil {
		return state, fmt.Errorf("reading: %w", err)
	}
	return state, nil
}

func (i *Index) Read(
	ctx provider.Context,
	name string,
	input IndexArgs,
	state IndexState,
) (string, IndexArgs, IndexState, error) {
	state.IndexArgs = input
	state.Ref = input.Tag

	if !input.Push {
		return name, input, state, nil // Nothing to read.
	}

	cli, err := i.client(ctx, state, input)
	if err != nil {
		return name, input, state, err
	}

	digest, err := cli.ManifestInspect(ctx, input.Tag)
	if err != nil && strings.Contains(err.Error(), "No such manifest:") && input.Push {
		// A remote tag was expected but isn't there -- delete the resource.
		return "", input, state, err
	}
	if err != nil && strings.Contains(err.Error(), "No such manifest:") && !input.Push {
		// Nothing was pushed, so just use the tag without digest..
		return name, input, state, err
	}
	if err != nil {
		return name, input, state, err
	}

	if ref, ok := addDigest(input.Tag, digest); ok {
		state.Ref = ref
	}

	return name, input, state, nil
}

func (i *Index) Check(
	_ provider.Context,
	_ string,
	_ resource.PropertyMap,
	news resource.PropertyMap,
) (IndexArgs, []provider.CheckFailure, error) {
	args, failures, err := infer.DefaultCheck[IndexArgs](news)
	if err != nil {
		return args, failures, err
	}

	if _, err := normalizeReference(args.Tag); args.Tag != "" && err != nil {
		failures = append(
			failures,
			provider.CheckFailure{
				Property: "target",
				Reason:   err.Error(),
			},
		)
	}

	for idx, s := range args.Sources {
		if _, err := normalizeReference(s); s != "" && err != nil {
			failures = append(
				failures,
				provider.CheckFailure{
					Property: fmt.Sprintf("refs[%d]", idx),
					Reason:   err.Error(),
				},
			)
		}
	}

	return args, failures, nil
}

func (i *Index) Delete(ctx provider.Context, _ string, state IndexState) error {
	if !state.Push {
		return nil // Nothing to delete.
	}

	cli, err := i.client(ctx, state, state.IndexArgs)
	if err != nil {
		return err
	}

	err = cli.ManifestDelete(ctx, state.Ref)
	// TODO:
	if err != nil && strings.Contains(err.Error(), "No such manifest:") {
		return nil
	}
	return err
}

func (i *Index) Diff(
	_ provider.Context,
	_ string,
	olds IndexState,
	news IndexArgs,
) (provider.DiffResponse, error) {
	diff := map[string]provider.PropertyDiff{}
	update := provider.PropertyDiff{Kind: provider.Update}
	replace := provider.PropertyDiff{Kind: provider.UpdateReplace}

	if olds.Tag != news.Tag {
		diff["tag"] = update
	}
	if !reflect.DeepEqual(olds.Sources, news.Sources) {
		diff["sources"] = update
	}
	if olds.Registry.Address != news.Registry.Address {
		diff["registry.address"] = update
		if olds.Registry.Address != "" {
			diff["registry.address"] = replace
		}
	}
	if olds.Registry.Username != news.Registry.Username {
		diff["registry.username"] = update
	}
	// Intentionally ignore changes to registry.password

	return provider.DiffResponse{
		DeleteBeforeReplace: false,
		HasChanges:          len(diff) > 0,
		DetailedDiff:        diff,
	}, nil
}

func (i *Index) client(
	ctx provider.Context,
	state IndexState,
	args IndexArgs,
) (Client, error) {
	cfg := infer.GetConfig[Config](ctx)

	if cli, ok := ctx.Value(_mockClientKey).(Client); ok {
		return cli, nil
	}

	auths := cfg.RegistryAuth
	auths = append(auths, state.Registry)
	auths = append(auths, args.Registry)

	return wrap(cfg.host, auths...)
}
