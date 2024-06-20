// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package docker

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-docker/sdk/v4/go/docker/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// <!-- Bug: Type and Name are switched -->
// Manages the lifecycle of a Docker plugin.
//
// ## Import
//
// #!/bin/bash
//
// ```sh
// $ pulumi import docker:index/plugin:Plugin sample-volume-plugin "$(docker plugin inspect -f {{.ID}} tiborvass/sample-volume-plugin:latest)"
// ```
type Plugin struct {
	pulumi.CustomResourceState

	// Docker Plugin alias
	Alias pulumi.StringOutput `pulumi:"alias"`
	// HTTP client timeout to enable the plugin
	EnableTimeout pulumi.IntPtrOutput `pulumi:"enableTimeout"`
	// If `true` the plugin is enabled. Defaults to `true`
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The environment variables in the form of `KEY=VALUE`, e.g. `DEBUG=0`
	Envs pulumi.StringArrayOutput `pulumi:"envs"`
	// If true, then the plugin is destroyed forcibly
	ForceDestroy pulumi.BoolPtrOutput `pulumi:"forceDestroy"`
	// If true, then the plugin is disabled forcibly
	ForceDisable pulumi.BoolPtrOutput `pulumi:"forceDisable"`
	// If true, grant all permissions necessary to run the plugin
	GrantAllPermissions pulumi.BoolPtrOutput `pulumi:"grantAllPermissions"`
	// Grant specific permissions only
	GrantPermissions PluginGrantPermissionArrayOutput `pulumi:"grantPermissions"`
	// Docker Plugin name
	Name pulumi.StringOutput `pulumi:"name"`
	// Docker Plugin Reference
	PluginReference pulumi.StringOutput `pulumi:"pluginReference"`
}

// NewPlugin registers a new resource with the given unique name, arguments, and options.
func NewPlugin(ctx *pulumi.Context,
	name string, args *PluginArgs, opts ...pulumi.ResourceOption) (*Plugin, error) {
	if args == nil {
		args = &PluginArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Plugin
	err := ctx.RegisterResource("docker:index/plugin:Plugin", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPlugin gets an existing Plugin resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPlugin(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PluginState, opts ...pulumi.ResourceOption) (*Plugin, error) {
	var resource Plugin
	err := ctx.ReadResource("docker:index/plugin:Plugin", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Plugin resources.
type pluginState struct {
	// Docker Plugin alias
	Alias *string `pulumi:"alias"`
	// HTTP client timeout to enable the plugin
	EnableTimeout *int `pulumi:"enableTimeout"`
	// If `true` the plugin is enabled. Defaults to `true`
	Enabled *bool `pulumi:"enabled"`
	// The environment variables in the form of `KEY=VALUE`, e.g. `DEBUG=0`
	Envs []string `pulumi:"envs"`
	// If true, then the plugin is destroyed forcibly
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// If true, then the plugin is disabled forcibly
	ForceDisable *bool `pulumi:"forceDisable"`
	// If true, grant all permissions necessary to run the plugin
	GrantAllPermissions *bool `pulumi:"grantAllPermissions"`
	// Grant specific permissions only
	GrantPermissions []PluginGrantPermission `pulumi:"grantPermissions"`
	// Docker Plugin name
	Name *string `pulumi:"name"`
	// Docker Plugin Reference
	PluginReference *string `pulumi:"pluginReference"`
}

type PluginState struct {
	// Docker Plugin alias
	Alias pulumi.StringPtrInput
	// HTTP client timeout to enable the plugin
	EnableTimeout pulumi.IntPtrInput
	// If `true` the plugin is enabled. Defaults to `true`
	Enabled pulumi.BoolPtrInput
	// The environment variables in the form of `KEY=VALUE`, e.g. `DEBUG=0`
	Envs pulumi.StringArrayInput
	// If true, then the plugin is destroyed forcibly
	ForceDestroy pulumi.BoolPtrInput
	// If true, then the plugin is disabled forcibly
	ForceDisable pulumi.BoolPtrInput
	// If true, grant all permissions necessary to run the plugin
	GrantAllPermissions pulumi.BoolPtrInput
	// Grant specific permissions only
	GrantPermissions PluginGrantPermissionArrayInput
	// Docker Plugin name
	Name pulumi.StringPtrInput
	// Docker Plugin Reference
	PluginReference pulumi.StringPtrInput
}

func (PluginState) ElementType() reflect.Type {
	return reflect.TypeOf((*pluginState)(nil)).Elem()
}

type pluginArgs struct {
	// Docker Plugin alias
	Alias *string `pulumi:"alias"`
	// HTTP client timeout to enable the plugin
	EnableTimeout *int `pulumi:"enableTimeout"`
	// If `true` the plugin is enabled. Defaults to `true`
	Enabled *bool `pulumi:"enabled"`
	// The environment variables in the form of `KEY=VALUE`, e.g. `DEBUG=0`
	Envs []string `pulumi:"envs"`
	// If true, then the plugin is destroyed forcibly
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// If true, then the plugin is disabled forcibly
	ForceDisable *bool `pulumi:"forceDisable"`
	// If true, grant all permissions necessary to run the plugin
	GrantAllPermissions *bool `pulumi:"grantAllPermissions"`
	// Grant specific permissions only
	GrantPermissions []PluginGrantPermission `pulumi:"grantPermissions"`
	// Docker Plugin name
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a Plugin resource.
type PluginArgs struct {
	// Docker Plugin alias
	Alias pulumi.StringPtrInput
	// HTTP client timeout to enable the plugin
	EnableTimeout pulumi.IntPtrInput
	// If `true` the plugin is enabled. Defaults to `true`
	Enabled pulumi.BoolPtrInput
	// The environment variables in the form of `KEY=VALUE`, e.g. `DEBUG=0`
	Envs pulumi.StringArrayInput
	// If true, then the plugin is destroyed forcibly
	ForceDestroy pulumi.BoolPtrInput
	// If true, then the plugin is disabled forcibly
	ForceDisable pulumi.BoolPtrInput
	// If true, grant all permissions necessary to run the plugin
	GrantAllPermissions pulumi.BoolPtrInput
	// Grant specific permissions only
	GrantPermissions PluginGrantPermissionArrayInput
	// Docker Plugin name
	Name pulumi.StringPtrInput
}

func (PluginArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pluginArgs)(nil)).Elem()
}

type PluginInput interface {
	pulumi.Input

	ToPluginOutput() PluginOutput
	ToPluginOutputWithContext(ctx context.Context) PluginOutput
}

func (*Plugin) ElementType() reflect.Type {
	return reflect.TypeOf((**Plugin)(nil)).Elem()
}

func (i *Plugin) ToPluginOutput() PluginOutput {
	return i.ToPluginOutputWithContext(context.Background())
}

func (i *Plugin) ToPluginOutputWithContext(ctx context.Context) PluginOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PluginOutput)
}

// PluginArrayInput is an input type that accepts PluginArray and PluginArrayOutput values.
// You can construct a concrete instance of `PluginArrayInput` via:
//
//	PluginArray{ PluginArgs{...} }
type PluginArrayInput interface {
	pulumi.Input

	ToPluginArrayOutput() PluginArrayOutput
	ToPluginArrayOutputWithContext(context.Context) PluginArrayOutput
}

type PluginArray []PluginInput

func (PluginArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Plugin)(nil)).Elem()
}

func (i PluginArray) ToPluginArrayOutput() PluginArrayOutput {
	return i.ToPluginArrayOutputWithContext(context.Background())
}

func (i PluginArray) ToPluginArrayOutputWithContext(ctx context.Context) PluginArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PluginArrayOutput)
}

// PluginMapInput is an input type that accepts PluginMap and PluginMapOutput values.
// You can construct a concrete instance of `PluginMapInput` via:
//
//	PluginMap{ "key": PluginArgs{...} }
type PluginMapInput interface {
	pulumi.Input

	ToPluginMapOutput() PluginMapOutput
	ToPluginMapOutputWithContext(context.Context) PluginMapOutput
}

type PluginMap map[string]PluginInput

func (PluginMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Plugin)(nil)).Elem()
}

func (i PluginMap) ToPluginMapOutput() PluginMapOutput {
	return i.ToPluginMapOutputWithContext(context.Background())
}

func (i PluginMap) ToPluginMapOutputWithContext(ctx context.Context) PluginMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PluginMapOutput)
}

type PluginOutput struct{ *pulumi.OutputState }

func (PluginOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Plugin)(nil)).Elem()
}

func (o PluginOutput) ToPluginOutput() PluginOutput {
	return o
}

func (o PluginOutput) ToPluginOutputWithContext(ctx context.Context) PluginOutput {
	return o
}

// Docker Plugin alias
func (o PluginOutput) Alias() pulumi.StringOutput {
	return o.ApplyT(func(v *Plugin) pulumi.StringOutput { return v.Alias }).(pulumi.StringOutput)
}

// HTTP client timeout to enable the plugin
func (o PluginOutput) EnableTimeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Plugin) pulumi.IntPtrOutput { return v.EnableTimeout }).(pulumi.IntPtrOutput)
}

// If `true` the plugin is enabled. Defaults to `true`
func (o PluginOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Plugin) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// The environment variables in the form of `KEY=VALUE`, e.g. `DEBUG=0`
func (o PluginOutput) Envs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Plugin) pulumi.StringArrayOutput { return v.Envs }).(pulumi.StringArrayOutput)
}

// If true, then the plugin is destroyed forcibly
func (o PluginOutput) ForceDestroy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Plugin) pulumi.BoolPtrOutput { return v.ForceDestroy }).(pulumi.BoolPtrOutput)
}

// If true, then the plugin is disabled forcibly
func (o PluginOutput) ForceDisable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Plugin) pulumi.BoolPtrOutput { return v.ForceDisable }).(pulumi.BoolPtrOutput)
}

// If true, grant all permissions necessary to run the plugin
func (o PluginOutput) GrantAllPermissions() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Plugin) pulumi.BoolPtrOutput { return v.GrantAllPermissions }).(pulumi.BoolPtrOutput)
}

// Grant specific permissions only
func (o PluginOutput) GrantPermissions() PluginGrantPermissionArrayOutput {
	return o.ApplyT(func(v *Plugin) PluginGrantPermissionArrayOutput { return v.GrantPermissions }).(PluginGrantPermissionArrayOutput)
}

// Docker Plugin name
func (o PluginOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Plugin) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Docker Plugin Reference
func (o PluginOutput) PluginReference() pulumi.StringOutput {
	return o.ApplyT(func(v *Plugin) pulumi.StringOutput { return v.PluginReference }).(pulumi.StringOutput)
}

type PluginArrayOutput struct{ *pulumi.OutputState }

func (PluginArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Plugin)(nil)).Elem()
}

func (o PluginArrayOutput) ToPluginArrayOutput() PluginArrayOutput {
	return o
}

func (o PluginArrayOutput) ToPluginArrayOutputWithContext(ctx context.Context) PluginArrayOutput {
	return o
}

func (o PluginArrayOutput) Index(i pulumi.IntInput) PluginOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Plugin {
		return vs[0].([]*Plugin)[vs[1].(int)]
	}).(PluginOutput)
}

type PluginMapOutput struct{ *pulumi.OutputState }

func (PluginMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Plugin)(nil)).Elem()
}

func (o PluginMapOutput) ToPluginMapOutput() PluginMapOutput {
	return o
}

func (o PluginMapOutput) ToPluginMapOutputWithContext(ctx context.Context) PluginMapOutput {
	return o
}

func (o PluginMapOutput) MapIndex(k pulumi.StringInput) PluginOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Plugin {
		return vs[0].(map[string]*Plugin)[vs[1].(string)]
	}).(PluginOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PluginInput)(nil)).Elem(), &Plugin{})
	pulumi.RegisterInputType(reflect.TypeOf((*PluginArrayInput)(nil)).Elem(), PluginArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PluginMapInput)(nil)).Elem(), PluginMap{})
	pulumi.RegisterOutputType(PluginOutput{})
	pulumi.RegisterOutputType(PluginArrayOutput{})
	pulumi.RegisterOutputType(PluginMapOutput{})
}
