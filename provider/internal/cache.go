package internal

import (
	"fmt"
	"strings"

	"github.com/pulumi/pulumi-go-provider/infer"
)

var (
	_ = (fmt.Stringer)((*CacheFromEntry)(nil))
	_ = (fmt.Stringer)((*CacheFromGitHubActions)(nil))
	_ = (fmt.Stringer)((*CacheFromRegistry)(nil))
	_ = (fmt.Stringer)((*CacheToInline)(nil))
	_ = (fmt.Stringer)((*CacheToEntry)(nil))
	_ = (fmt.Stringer)((*CacheToGitHubActions)(nil))
	_ = (fmt.Stringer)((*CacheToRegistry)(nil))
	_ = (fmt.Stringer)(CacheWithCompression{})
	_ = (fmt.Stringer)(CacheWithIgnoreError{})
	_ = (fmt.Stringer)(CacheWithMode{})
	_ = (fmt.Stringer)(CacheWithOCI{})
	_ = (infer.Annotated)((*CacheFromEntry)(nil))
	_ = (infer.Annotated)((*CacheFromGitHubActions)(nil))
	_ = (infer.Annotated)((*CacheFromRegistry)(nil))
	_ = (infer.Annotated)((*CacheToEntry)(nil))
	_ = (infer.Annotated)((*CacheWithCompression)(nil))
	_ = (infer.Annotated)((*CacheWithIgnoreError)(nil))
	_ = (infer.Annotated)((*CacheWithMode)(nil))
	_ = (infer.Annotated)((*CacheWithOCI)(nil))
	_ = (infer.Enum[CacheMode])((*CacheMode)(nil))
	_ = (infer.Enum[CompressionType])((*CompressionType)(nil))
)

type CacheFromLocal struct {
	Src    string `pulumi:"src"`
	Digest string `pulumi:"digest,optional"`
}

func (c *CacheFromLocal) String() string {
	if c == nil {
		return ""
	}
	parts := []string{"type=local"}
	if c.Src != "" {
		parts = append(parts, fmt.Sprintf("src=%s", c.Src))
	}
	if c.Digest != "" {
		parts = append(parts, fmt.Sprintf("digest=%s", c.Digest))
	}
	return strings.Join(parts, ",")
}

func (c *CacheFromLocal) Annotate(a infer.Annotator) {
	a.Describe(&c.Src, "Path of the local directory where cache gets imported from.")
	a.Describe(&c.Digest, "Digest of manifest to import.")
}

type CacheFromRegistry struct {
	Ref string `pulumi:"ref"`
}

func (c *CacheFromRegistry) Annotate(a infer.Annotator) {
	a.Describe(&c.Ref, "Fully qualified name of the cache image to import.")
}

func (c *CacheFromRegistry) String() string {
	if c == nil {
		return ""
	}
	return fmt.Sprintf("type=registry,ref=%s", c.Ref)
}

type CacheWithOCI struct {
	OCI           *bool `pulumi:"ociMediaTypes,optional"`
	ImageManifest *bool `pulumi:"imageManifest,optional"`
}

func (c *CacheWithOCI) Annotate(a infer.Annotator) {
	a.Describe(&c.OCI, dedent(`
		Whether to use OCI media types in exported manifests. Defaults to
		"true".
	`))
	a.Describe(&c.ImageManifest, dedent(`
		Export cache manifest as an OCI-compatible image manifest instead of a
		manifest list (requires OCI media types).

		Defaults to "false".
	`))

	a.SetDefault(&c.OCI, true)
	a.SetDefault(&c.ImageManifest, false)
}

func (c CacheWithOCI) String() string {
	if c.OCI == nil {
		return ""
	}
	parts := []string{fmt.Sprintf("oci-mediatypes=%t", *c.OCI)}
	if c.ImageManifest != nil {
		parts = append(parts, fmt.Sprintf("image-manifest=%t", *c.ImageManifest))
	}
	return strings.Join(parts, ",")
}

type CacheFromGitHubActions struct {
	URL   string `pulumi:"url,optional"`
	Token string `pulumi:"token,optional" provider:"secret"`
	Scope string `pulumi:"scope,optional"`
}

func (c *CacheFromGitHubActions) Annotate(a infer.Annotator) {
	a.SetDefault(&c.URL, "", "ACTIONS_RUNTIME_URL")
	a.SetDefault(&c.Token, "", "ACTIONS_RUNTIME_TOKEN")
	a.SetDefault(&c.Scope, "", "buildkit")

	a.Describe(&c.URL, dedent(`
		The cache server URL to use for artifacts.

		Defaults to "$ACTIONS_RUNTIME_URL", although a separate action like
		"crazy-max/ghaction-github-runtime" is recommended to expose this
		environment variable to your jobs.
	`))
	a.Describe(&c.Token, dedent(`
		The GitHub Actions token to use. This is not a personal access tokens
		and is typically generated automatically as part of each job.

		Defaults to "$ACTIONS_RUNTIME_TOKEN", although a separate action like
		"crazy-max/ghaction-github-runtime" is recommended to expose this
		environment variable to your jobs.

	`))
	a.Describe(&c.Scope, dedent(`
		The scope to use for cache keys. Defaults to "buildkit".

		This should be set if building and caching multiple images in one
		workflow, otherwise caches will overwrite each other.
	`))
}

func (c *CacheFromGitHubActions) String() string {
	if c == nil {
		return ""
	}
	parts := []string{"type=gha"}
	if c.Scope != "" {
		parts = append(parts, fmt.Sprintf("scope=%s", c.Scope))
	}
	if c.Token != "" {
		parts = append(parts, fmt.Sprintf("token=%s", c.Token))
	}
	if c.URL != "" {
		parts = append(parts, fmt.Sprintf("url=%s", c.URL))
	}
	return strings.Join(parts, ",")
}

type CacheFromAzureBlob struct {
	Name            string `pulumi:"name"`
	AccountURL      string `pulumi:"accountUrl,optional"`
	SecretAccessKey string `pulumi:"secretAccessKey,optional" provider:"secret"`
}

func (c *CacheFromAzureBlob) String() string {
	if c == nil {
		return ""
	}
	parts := []string{"type=azblob"}
	if c.Name != "" {
		parts = append(parts, fmt.Sprintf("name=%s", c.Name))
	}
	if c.AccountURL != "" {
		parts = append(parts, fmt.Sprintf("account_url=%s", c.AccountURL))
	}
	if c.SecretAccessKey != "" {
		parts = append(parts, fmt.Sprintf("secret_access_key=%s", c.SecretAccessKey))
	}
	return strings.Join(parts, ",")
}

func (c *CacheFromAzureBlob) Annotate(a infer.Annotator) {
	a.Describe(&c.Name, "The name of the cache image.")
	a.Describe(&c.AccountURL, "Base URL of the storage account.")
	a.Describe(&c.SecretAccessKey, "Blob storage account key.")
}

type CacheToAzureBlob struct {
	CacheWithMode
	CacheWithIgnoreError

	CacheFromAzureBlob
}

func (c *CacheToAzureBlob) String() string {
	if c == nil {
		return ""
	}
	return join(&c.CacheFromAzureBlob, c.CacheWithMode, c.CacheWithIgnoreError)
}

type CacheFromS3 struct {
	Region          string `pulumi:"region"`
	Bucket          string `pulumi:"bucket"`
	Name            string `pulumi:"name,optional"`
	EndpointURL     string `pulumi:"endpointUrl,optional"`
	BlobsPrefix     string `pulumi:"blobsPrefix,optional"`
	ManifestsPrefix string `pulumi:"manifestsPrefix,optional"`
	UsePathStyle    *bool  `pulumi:"usePathStyle,optional"`
	AccessKeyID     string `pulumi:"accessKeyId,optional"`
	SecretAccessKey string `pulumi:"secretAccessKey,optional" provider:"secret"`
	SessionToken    string `pulumi:"sessionToken,optional" provider:"secret"`
}

func (c *CacheFromS3) Annotate(a infer.Annotator) {
	a.SetDefault(&c.Region, "", "AWS_REGION")
	a.SetDefault(&c.AccessKeyID, "", "AWS_ACCESS_KEY_ID")
	a.SetDefault(&c.SecretAccessKey, "", "AWS_SECRET_ACCESS_KEY")
	a.SetDefault(&c.SessionToken, "", "AWS_SESSION_TOKEN")

	a.Describe(&c.Bucket, dedent(`
		Name of the S3 bucket.
	`))
	a.Describe(&c.Region, dedent(`
		The geographic location of the bucket. Defaults to "$AWS_REGION".
	`))
	a.Describe(&c.AccessKeyID, dedent(`
		Defaults to "$AWS_ACCESS_KEY_ID".
	`))
	a.Describe(&c.SecretAccessKey, dedent(`
		Defaults to "$AWS_SECRET_ACCESS_KEY".
	`))
	a.Describe(&c.SessionToken, dedent(`
		Defaults to "$AWS_SESSION_TOKEN".
	`))
	a.Describe(&c.BlobsPrefix, dedent(`
		Prefix to prepend to blob filenames.
	`))
	a.Describe(&c.EndpointURL, dedent(`
		Endpoint of the S3 bucket.
	`))
	a.Describe(&c.ManifestsPrefix, dedent(`
		Prefix to prepend on manifest filenames.
	`))
	a.Describe(&c.Name, dedent(`
		Name of the cache image.
	`))
	a.Describe(&c.UsePathStyle, dedent(`
		Uses "bucket" in the URL instead of hostname when "true".
	`))
}

func (c *CacheFromS3) String() string {
	if c == nil {
		return ""
	}
	parts := []string{"type=s3"}
	if c.Bucket != "" {
		parts = append(parts, fmt.Sprintf("bucket=%s", c.Bucket))
	}
	if c.Name != "" {
		parts = append(parts, fmt.Sprintf("name=%s", c.Name))
	}
	if c.EndpointURL != "" {
		parts = append(parts, fmt.Sprintf("endpoint_url=%s", c.EndpointURL))
	}
	if c.BlobsPrefix != "" {
		parts = append(parts, fmt.Sprintf("blobs_prefix=%s", c.BlobsPrefix))
	}
	if c.ManifestsPrefix != "" {
		parts = append(parts, fmt.Sprintf("manifests_prefix=%s", c.ManifestsPrefix))
	}
	if c.UsePathStyle != nil {
		parts = append(parts, fmt.Sprintf("use_path_type=%t", *c.UsePathStyle))
	}
	if c.AccessKeyID != "" {
		parts = append(parts, fmt.Sprintf("access_key_id=%s", c.AccessKeyID))
	}
	if c.SecretAccessKey != "" {
		parts = append(parts, fmt.Sprintf("secret_access_key=%s", c.SecretAccessKey))
	}
	if c.SessionToken != "" {
		parts = append(parts, fmt.Sprintf("session_token=%s", c.SessionToken))
	}

	return strings.Join(parts, ",")
}

type CacheWithMode struct {
	Mode CacheMode `pulumi:"mode,optional"`
}

func (c *CacheWithMode) Annotate(a infer.Annotator) {
	a.SetDefault(&c.Mode, CacheModeMin)
	a.Describe(&c.Mode, dedent(`
		The cache mode to use. Defaults to "min".
	`))
}

func (c CacheWithMode) String() string {
	if c.Mode == "" {
		return ""
	}
	return fmt.Sprintf("mode=%s", c.Mode)
}

type CacheWithIgnoreError struct {
	IgnoreError *bool `pulumi:"ignoreError,optional"`
}

func (c *CacheWithIgnoreError) Annotate(a infer.Annotator) {
	a.SetDefault(&c.IgnoreError, false)
	a.Describe(&c.IgnoreError, "Ignore errors caused by failed cache exports.")
}

func (c CacheWithIgnoreError) String() string {
	if c.IgnoreError == nil {
		return ""
	}
	return fmt.Sprintf("ignore-error=%t", *c.IgnoreError)
}

type CacheToS3 struct {
	CacheWithMode
	CacheWithIgnoreError

	CacheFromS3
}

func (c *CacheToS3) String() string {
	if c == nil {
		return ""
	}
	return join(&c.CacheFromS3, c.CacheWithMode, c.CacheWithIgnoreError)
}

type Raw string

func (c Raw) String() string {
	return string(c)
}

type CacheFromEntry struct {
	Local    *CacheFromLocal         `pulumi:"local,optional"`
	Registry *CacheFromRegistry      `pulumi:"registry,optional"`
	GHA      *CacheFromGitHubActions `pulumi:"gha,optional"`
	AZBlob   *CacheFromAzureBlob     `pulumi:"azblob,optional"`
	S3       *CacheFromS3            `pulumi:"s3,optional"`
	Raw      Raw                     `pulumi:"raw,optional"`

	Disabled bool `pulumi:"disabled,optional"`
}

func (c *CacheFromEntry) Annotate(a infer.Annotator) {
	a.Describe(&c.Local, dedent(`
		A simple backend which caches images on your local filesystem.
	`))
	a.Describe(&c.Registry, dedent(`
		Upload build caches to remote registries.
	`))
	a.Describe(&c.GHA, dedent(`
		Recommended for use with GitHub Actions workflows.

		An action like "crazy-max/ghaction-github-runtime" is recommended to
		expose appropriate credentials to your GitHub workflow.
	`))
	a.Describe(&c.AZBlob, dedent(`
		Upload build caches to Azure's blob storage service.
	`))
	a.Describe(&c.S3, dedent(`
		Upload build caches to AWS S3 or an S3-compatible services such as
		MinIO.
	`))
	a.Describe(&c.Raw, dedent(`
		A raw string as you would provide it to the Docker CLI (e.g.,
		"type=inline").
	`))

	a.Describe(&c.Disabled, dedent(`
		When "true" this entry will be excluded. Defaults to "false".
	`))
}

func (c CacheFromEntry) String() string {
	if c.Disabled {
		return ""
	}
	return join(c.Local, c.Registry, c.GHA, c.AZBlob, c.S3, c.Raw)
}

type CacheToInline struct{}

func (c *CacheToInline) String() string {
	if c == nil {
		return ""
	}
	return "type=inline"
}

type CacheToLocal struct {
	CacheWithCompression
	CacheWithIgnoreError
	CacheWithMode

	Dest string `pulumi:"dest"`
}

func (c *CacheToLocal) Annotate(a infer.Annotator) {
	a.Describe(&c.Dest, dedent(`
		Path of the local directory to export the cache.
	`))
}

func (c *CacheToLocal) String() string {
	if c == nil {
		return ""
	}
	return join(
		Raw(fmt.Sprintf("type=local,dest=%s", c.Dest)),
		c.CacheWithCompression,
		c.CacheWithIgnoreError,
	)
}

type CacheToRegistry struct {
	CacheWithMode
	CacheWithIgnoreError
	CacheWithOCI
	CacheWithCompression

	CacheFromRegistry
}

func (c *CacheToRegistry) String() string {
	if c == nil {
		return ""
	}
	return join(
		&c.CacheFromRegistry,
		c.CacheWithMode,
		c.CacheWithIgnoreError,
		c.CacheWithOCI,
		c.CacheWithCompression,
	)
}

type CacheWithCompression struct {
	Compression      CompressionType `pulumi:"compression,optional"`
	CompressionLevel int             `pulumi:"compressionLevel,optional"`
	ForceCompression *bool           `pulumi:"forceCompression,optional"`
}

func (c *CacheWithCompression) Annotate(a infer.Annotator) {
	a.SetDefault(&c.Compression, CompressionTypeGzip)
	a.SetDefault(&c.CompressionLevel, 0)
	a.SetDefault(&c.ForceCompression, false)

	a.Describe(&c.Compression, "The compression type to use.")
	a.Describe(&c.CompressionLevel, "Compression level from 0 to 22.")
	a.Describe(&c.ForceCompression, "Forcefully apply compression.")
}

func (c CacheWithCompression) String() string {
	if c.CompressionLevel == 0 {
		return ""
	}
	parts := []string{}
	if c.Compression != "" {
		parts = append(parts, fmt.Sprintf("compression=%s", c.Compression))
	}
	if c.CompressionLevel > 0 {
		cl := c.CompressionLevel
		if cl > 22 {
			cl = 22
		}
		parts = append(parts, fmt.Sprintf("compression-level=%d", cl))
	}
	if c.ForceCompression != nil {
		parts = append(parts, fmt.Sprintf("force-compression=%t", *c.ForceCompression))
	}
	return strings.Join(parts, ",")
}

type CacheToGitHubActions struct {
	CacheWithMode
	CacheWithIgnoreError

	CacheFromGitHubActions
}

func (c *CacheToGitHubActions) String() string {
	if c == nil {
		return ""
	}
	return join(&c.CacheFromGitHubActions, c.CacheWithMode, c.CacheWithIgnoreError)
}

type CacheToEntry struct {
	Inline   *CacheToInline        `pulumi:"inline,optional"`
	Local    *CacheToLocal         `pulumi:"local,optional"`
	Registry *CacheToRegistry      `pulumi:"registry,optional"`
	GHA      *CacheToGitHubActions `pulumi:"gha,optional"`
	AZBlob   *CacheToAzureBlob     `pulumi:"azblob,optional"`
	S3       *CacheToS3            `pulumi:"s3,optional"`
	Raw      Raw                   `pulumi:"raw,optional"`

	Disabled bool `pulumi:"disabled,optional"`
}

func (c *CacheToEntry) Annotate(a infer.Annotator) {
	a.Describe(&c.Inline, dedent(`
		The inline cache storage backend is the simplest implementation to get
		started with, but it does not handle multi-stage builds. Consider the
		"registry" cache backend instead.
	`))
	a.Describe(&c.Local, dedent(`
		A simple backend which caches imagines on your local filesystem.
	`))
	a.Describe(&c.Registry, dedent(`
		Push caches to remote registries. Incompatible with the "docker" build
		driver.
	`))
	a.Describe(&c.GHA, dedent(`
		Recommended for use with GitHub Actions workflows.

		An action like "crazy-max/ghaction-github-runtime" is recommended to
		expose appropriate credentials to your GitHub workflow.
	`))
	a.Describe(&c.AZBlob, dedent(`
		Push cache to Azure's blob storage service.
	`))
	a.Describe(&c.S3, dedent(`
		Push cache to AWS S3 or S3-compatible services such as MinIO.
	`))
	a.Describe(&c.Raw, dedent(`
		A raw string as you would provide it to the Docker CLI (e.g.,
		"type=inline")`,
	))

	a.Describe(&c.Disabled, dedent(`
		When "true" this entry will be excluded. Defaults to "false".
	`))
}

func (c CacheToEntry) String() string {
	if c.Disabled {
		return ""
	}
	return join(c.Inline, c.Local, c.Registry, c.GHA, c.AZBlob, c.S3, c.Raw)
}

type CacheMode string

const (
	CacheModeMin CacheMode = "min"
	CacheModeMax CacheMode = "max"
)

func (CacheMode) Values() []infer.EnumValue[CacheMode] {
	return []infer.EnumValue[CacheMode]{
		{
			Value:       CacheModeMin,
			Description: "Only layers that are exported into the resulting image are cached.",
		},
		{
			Value:       CacheModeMax,
			Description: "All layers are cached, even those of intermediate steps.",
		},
	}
}

type CompressionType string

const (
	CompressionTypeGzip    CompressionType = "gzip"
	CompressionTypeEstargz CompressionType = "estargz"
	CompressionTypeZstd    CompressionType = "zstd"
)

func (CompressionType) Values() []infer.EnumValue[CompressionType] {
	return []infer.EnumValue[CompressionType]{
		{Value: CompressionTypeGzip, Description: "Use `gzip` for compression."},
		{Value: CompressionTypeEstargz, Description: "Use `estargz` for compression."},
		{Value: CompressionTypeZstd, Description: "Use `zstd` for compression."},
	}
}

type joiner struct{ sep string }

func (j joiner) join(ss ...fmt.Stringer) string {
	parts := []string{}
	for _, s := range ss {
		p := s.String()
		if p == "" {
			continue
		}
		parts = append(parts, p)
	}
	return strings.Join(parts, j.sep)
}

func join(ss ...fmt.Stringer) string {
	return joiner{","}.join(ss...)
}
