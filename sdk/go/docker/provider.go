// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package docker

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The provider type for the docker package. By default, resources use package-wide configuration
// settings, however an explicit `Provider` instance may be created and passed during resource
// construction to achieve fine-grained programmatic control over provider settings. See the
// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
type Provider struct {
	pulumi.ProviderResourceState

	// PEM-encoded content of Docker host CA certificate
	CaMaterial pulumi.StringPtrOutput `pulumi:"caMaterial"`
	// PEM-encoded content of Docker client certificate
	CertMaterial pulumi.StringPtrOutput `pulumi:"certMaterial"`
	// Path to directory with Docker TLS config
	CertPath pulumi.StringPtrOutput `pulumi:"certPath"`
	// The Docker daemon address
	Host pulumi.StringPtrOutput `pulumi:"host"`
	// PEM-encoded content of Docker client private key
	KeyMaterial pulumi.StringPtrOutput `pulumi:"keyMaterial"`
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		args = &ProviderArgs{}
	}

	if isZero(args.Host) {
		args.Host = pulumi.StringPtr(getEnvOrDefault("unix:///var/run/docker.sock", nil, "DOCKER_HOST").(string))
	}
	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:docker", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	// PEM-encoded content of Docker host CA certificate
	CaMaterial *string `pulumi:"caMaterial"`
	// PEM-encoded content of Docker client certificate
	CertMaterial *string `pulumi:"certMaterial"`
	// Path to directory with Docker TLS config
	CertPath *string `pulumi:"certPath"`
	// The Docker daemon address
	Host *string `pulumi:"host"`
	// PEM-encoded content of Docker client private key
	KeyMaterial  *string                `pulumi:"keyMaterial"`
	RegistryAuth []ProviderRegistryAuth `pulumi:"registryAuth"`
	// Additional SSH option flags to be appended when using `ssh://` protocol
	SshOpts []string `pulumi:"sshOpts"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	// PEM-encoded content of Docker host CA certificate
	CaMaterial pulumi.StringPtrInput
	// PEM-encoded content of Docker client certificate
	CertMaterial pulumi.StringPtrInput
	// Path to directory with Docker TLS config
	CertPath pulumi.StringPtrInput
	// The Docker daemon address
	Host pulumi.StringPtrInput
	// PEM-encoded content of Docker client private key
	KeyMaterial  pulumi.StringPtrInput
	RegistryAuth ProviderRegistryAuthArrayInput
	// Additional SSH option flags to be appended when using `ssh://` protocol
	SshOpts pulumi.StringArrayInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}

type ProviderInput interface {
	pulumi.Input

	ToProviderOutput() ProviderOutput
	ToProviderOutputWithContext(ctx context.Context) ProviderOutput
}

func (*Provider) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (i *Provider) ToProviderOutput() ProviderOutput {
	return i.ToProviderOutputWithContext(context.Background())
}

func (i *Provider) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderOutput)
}

type ProviderOutput struct{ *pulumi.OutputState }

func (ProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (o ProviderOutput) ToProviderOutput() ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return o
}

// PEM-encoded content of Docker host CA certificate
func (o ProviderOutput) CaMaterial() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.CaMaterial }).(pulumi.StringPtrOutput)
}

// PEM-encoded content of Docker client certificate
func (o ProviderOutput) CertMaterial() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.CertMaterial }).(pulumi.StringPtrOutput)
}

// Path to directory with Docker TLS config
func (o ProviderOutput) CertPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.CertPath }).(pulumi.StringPtrOutput)
}

// The Docker daemon address
func (o ProviderOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.Host }).(pulumi.StringPtrOutput)
}

// PEM-encoded content of Docker client private key
func (o ProviderOutput) KeyMaterial() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.KeyMaterial }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderInput)(nil)).Elem(), &Provider{})
	pulumi.RegisterOutputType(ProviderOutput{})
}
