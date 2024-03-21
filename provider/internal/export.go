package internal

import (
	"fmt"
	"slices"
	"strings"

	controllerapi "github.com/docker/buildx/controller/pb"
	"github.com/docker/buildx/util/buildflags"

	"github.com/pulumi/pulumi-go-provider/infer"
)

var (
	_ = (fmt.Stringer)((*Export)(nil))
	_ = (fmt.Stringer)((*ExportDocker)(nil))
	_ = (fmt.Stringer)((*ExportImage)(nil))
	_ = (fmt.Stringer)((*ExportLocal)(nil))
	_ = (fmt.Stringer)((*ExportOCI)(nil))
	_ = (fmt.Stringer)((*ExportRegistry)(nil))
	_ = (fmt.Stringer)((*ExportTar)(nil))
	_ = (fmt.Stringer)(ExportWithAnnotations{})
	_ = (fmt.Stringer)(ExportWithCompression{})
	_ = (fmt.Stringer)(ExportWithNames{})
	_ = (fmt.Stringer)(ExportWithOCI{})
	_ = (infer.Annotated)((*Export)(nil))
	_ = (infer.Annotated)((*ExportDocker)(nil))
	_ = (infer.Annotated)((*ExportImage)(nil))
	_ = (infer.Annotated)((*ExportLocal)(nil))
	_ = (infer.Annotated)((*ExportOCI)(nil))
	_ = (infer.Annotated)((*ExportRegistry)(nil))
	_ = (infer.Annotated)((*ExportTar)(nil))
)

type Export struct {
	Tar       *ExportTar       `pulumi:"tar,optional"`
	Local     *ExportLocal     `pulumi:"local,optional"`
	Registry  *ExportRegistry  `pulumi:"registry,optional"`
	Image     *ExportImage     `pulumi:"image,optional"`
	OCI       *ExportOCI       `pulumi:"oci,optional"`
	Docker    *ExportDocker    `pulumi:"docker,optional"`
	CacheOnly *ExportCacheOnly `pulumi:"cacheonly,optional"`
	Raw       Raw              `pulumi:"raw,optional"`

	Disabled bool `pulumi:"disabled,optional"`
}

func (e *Export) Annotate(a infer.Annotator) {
	a.Describe(&e.Tar, dedent(`
		Export to a local directory as a tarball.`,
	))
	a.Describe(&e.Local, dedent(`
		Export to a local directory as files and directories.`,
	))
	a.Describe(&e.Registry, dedent(`
		Identical to the Image exporter, but pushes by default.`,
	))
	a.Describe(&e.Image, dedent(`
		Outputs the build result into a container image format.`,
	))
	a.Describe(&e.OCI, dedent(`
		Identical to the Docker exporter but uses OCI media types by default.`,
	))
	a.Describe(&e.Docker, dedent(`
		Export as a Docker image layout.`,
	))
	a.Describe(&e.Raw, dedent(`
		A raw string as you would provide it to the Docker CLI (e.g.,
		"type=docker")`,
	))
	a.Describe(&e.CacheOnly, dedent(`
		A no-op export. Helpful for silencing the 'no exports' warning if you
		just want to populate caches.
	`))

	a.Describe(&e.Disabled, dedent(`
		When "true" this entry will be excluded. Defaults to "false".
	`))
}

func (e Export) String() string {
	if e.Disabled {
		return ""
	}
	return join(e.Tar, e.Local, e.Registry, e.Image, e.OCI, e.Docker, e.CacheOnly, e.Raw)
}

func (e Export) pushed() bool {
	if e.Raw != "" {
		exp, err := buildflags.ParseExports([]string{e.Raw.String()})
		if err != nil {
			return false
		}
		return exp[0].Attrs["push"] == "true"
	}
	if e.Registry != nil {
		return e.Registry.Push == nil || *e.Registry.Push
	}
	if e.Image != nil {
		return e.Image.Push != nil && *e.Image.Push
	}
	return false
}

type ExportCacheOnly struct{}

func (e *ExportCacheOnly) String() string {
	if e == nil {
		return ""
	}
	return "type=cacheonly"
}

type ExportDocker struct {
	ExportWithOCI
	ExportWithCompression
	ExportWithAnnotations
	ExportWithNames

	Dest string `pulumi:"dest,optional"`
	Tar  *bool  `pulumi:"tar,optional"`
}

func (e *ExportDocker) Annotate(a infer.Annotator) {
	a.SetDefault(&e.Tar, true)

	a.Describe(&e.Dest, "The local export path.")
	a.Describe(&e.Tar, "Bundle the output into a tarball layout.")
}

func (e *ExportDocker) String() string {
	if e == nil {
		return ""
	}
	parts := []string{}
	if e.Dest != "" {
		parts = append(parts, fmt.Sprintf("dest=%s", e.Dest))
	}
	if e.Tar != nil {
		parts = append(parts, fmt.Sprintf("tar=%t", *e.Tar))
	}

	return join(
		Raw("type=docker"),
		Raw(strings.Join(parts, ",")),
		e.ExportWithOCI,
		e.ExportWithCompression,
		e.ExportWithAnnotations,
		e.ExportWithNames,
	)
}

type ExportOCI struct {
	ExportDocker
}

func (e *ExportOCI) Annotate(a infer.Annotator) {
	a.SetDefault(&e.OCI, true)
	a.Describe(&e.OCI, "Use OCI media types in exporter manifests.")
}

func (e *ExportOCI) String() string {
	if e == nil {
		return ""
	}
	return strings.Replace(e.ExportDocker.String(), "type=docker", "type=oci", 1)
}

type ExportImage struct {
	ExportWithOCI
	ExportWithCompression
	ExportWithNames
	ExportWithAnnotations

	Push               *bool  `pulumi:"push,optional"`
	PushByDigest       *bool  `pulumi:"pushByDigest,optional"`
	Insecure           *bool  `pulumi:"insecure,optional"`
	DanglingNamePrefix string `pulumi:"danglingNamePrefix,optional"`
	NameCanonical      *bool  `pulumi:"nameCanonical,optional"`
	Unpack             *bool  `pulumi:"unpack,optional"`
	Store              *bool  `pulumi:"store,optional"`
}

func (e *ExportImage) Annotate(a infer.Annotator) {
	a.SetDefault(&e.Store, true)

	a.Describe(&e.Store, dedent(`
		Store resulting images to the worker's image store and ensure all of
		its blobs are in the content store.

		Defaults to "true".

		Ignored if the worker doesn't have image store (when using OCI workers,
		for example).
	`))
	a.Describe(&e.Push, "Push after creating the image.")
	a.Describe(&e.DanglingNamePrefix, dedent(`
		Name image with "prefix@<digest>", used for anonymous images.
	`))
	a.Describe(&e.NameCanonical, dedent(`
		Add additional canonical name ("name@<digest>").
	`))
	a.Describe(&e.Insecure, dedent(`
		Allow pushing to an insecure registry.
	`))
	a.Describe(&e.PushByDigest, dedent(`
		Push image without name.
	`))
	a.Describe(&e.Unpack, dedent(`
		Unpack image after creation (for use with containerd). Defaults to
		"false".
	`))
}

func (e *ExportImage) String() string {
	if e == nil {
		return ""
	}
	parts := []string{}
	if e.Push != nil {
		parts = append(parts, fmt.Sprintf("push=%t", *e.Push))
	}
	if e.PushByDigest != nil {
		parts = append(parts, fmt.Sprintf("push-by-digest=%t", *e.PushByDigest))
	}
	if e.Insecure != nil {
		parts = append(parts, fmt.Sprintf("insecure=%t", *e.Insecure))
	}
	if e.DanglingNamePrefix != "" {
		parts = append(parts, fmt.Sprintf("dangling-name-prefix=%s", e.DanglingNamePrefix))
	}
	if e.NameCanonical != nil {
		parts = append(parts, fmt.Sprintf("name-canonical=%t", *e.NameCanonical))
	}
	if e.Unpack != nil {
		parts = append(parts, fmt.Sprintf("unpack=%t", *e.Unpack))
	}
	if e.Store != nil {
		parts = append(parts, fmt.Sprintf("store=%t", *e.Store))
	}
	return join(
		Raw("type=image"),
		Raw(strings.Join(parts, ",")),
		e.ExportWithOCI,
		e.ExportWithCompression,
		e.ExportWithNames,
		e.ExportWithAnnotations,
	)
}

type ExportRegistry struct {
	ExportImage
}

func (e *ExportRegistry) Annotate(a infer.Annotator) {
	a.SetDefault(&e.Push, true)
}

func (e *ExportRegistry) String() string {
	if e == nil {
		return ""
	}
	return strings.Replace(e.ExportImage.String(), "type=image", "type=registry", 1)
}

type ExportLocal struct {
	Dest string `pulumi:"dest"`
}

func (e *ExportLocal) String() string {
	if e == nil {
		return ""
	}
	return fmt.Sprintf("type=local,dest=%s", e.Dest)
}

func (e *ExportLocal) Annotate(a infer.Annotator) {
	a.Describe(&e.Dest, "Output path.")
}

type ExportTar struct {
	ExportLocal
}

func (e *ExportTar) String() string {
	if e == nil {
		return ""
	}
	return fmt.Sprintf("type=tar,dest=%s", e.Dest)
}

type ExportWithOCI struct {
	OCI *bool `pulumi:"ociMediaTypes,optional"`
}

func (c *ExportWithOCI) Annotate(a infer.Annotator) {
	a.SetDefault(&c.OCI, false)
	a.Describe(&c.OCI, "Use OCI media types in exporter manifests.")
}

func (c ExportWithOCI) String() string {
	if c.OCI == nil {
		return ""
	}
	return fmt.Sprintf("oci-mediatypes=%t", *c.OCI)
}

type ExportWithCompression struct {
	Compression      CompressionType `pulumi:"compression,optional"`
	CompressionLevel int             `pulumi:"compressionLevel,optional"`
	ForceCompression *bool           `pulumi:"forceCompression,optional"`
}

func (e *ExportWithCompression) Annotate(a infer.Annotator) {
	a.SetDefault(&e.Compression, CompressionTypeGzip)
	a.SetDefault(&e.CompressionLevel, 0)
	a.SetDefault(&e.ForceCompression, false)

	a.Describe(&e.Compression, "The compression type to use.")
	a.Describe(&e.CompressionLevel, "Compression level from 0 to 22.")
	a.Describe(&e.ForceCompression, "Forcefully apply compression.")
}

func (e ExportWithCompression) String() string {
	if e.CompressionLevel == 0 {
		return ""
	}
	parts := []string{}
	if e.Compression != "" {
		parts = append(parts, fmt.Sprintf("compression=%s", e.Compression))
	}
	if e.CompressionLevel > 0 {
		cl := e.CompressionLevel
		if cl > 22 {
			cl = 22
		}
		parts = append(parts, fmt.Sprintf("compression-level=%d", cl))
	}
	if e.ForceCompression != nil {
		parts = append(parts, fmt.Sprintf("force-compression=%t", *e.ForceCompression))
	}
	return strings.Join(parts, ",")
}

type ExportWithNames struct {
	Names []string `pulumi:"names,optional"`
}

func (e ExportWithNames) String() string {
	parts := []string{}
	for _, n := range e.Names {
		parts = append(parts, fmt.Sprintf("name=%s", n))
	}
	return strings.Join(parts, ",")
}

func (e *ExportWithNames) Annotate(a infer.Annotator) {
	a.Describe(&e.Names, "Specify images names to export. This is overridden if tags are already specified.")
}

type ExportWithAnnotations struct {
	Annotations map[string]string `pulumi:"annotations,optional"`
}

func (e ExportWithAnnotations) String() string {
	parts := []string{}
	for k, v := range e.Annotations {
		parts = append(parts, fmt.Sprintf("annotation.%s=%s", k, v))
	}
	slices.Sort(parts)
	return strings.Join(parts, ",")
}

func (e *ExportWithAnnotations) Annotate(a infer.Annotator) {
	a.Describe(&e.Annotations, dedent(`
		Attach an arbitrary key/value annotation to the image.
	`))
}

func isRegistryPush(export *controllerapi.ExportEntry) bool {
	return export.Type == "image" && export.Attrs["push"] == "true"
}
