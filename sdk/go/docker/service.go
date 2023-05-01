// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package docker

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// ### Example Assuming you created a `service` as follows #!/bin/bash docker service create --name foo -p 8080:80 nginx prints th ID 4pcphbxkfn2rffhbhe6czytgi you provide the definition for the resource as follows terraform resource "docker_service" "foo" {
//
//	name = "foo"
//
//	task_spec {
//
//	container_spec {
//
//	image = "nginx"
//
//	}
//
//	}
//
//	endpoint_spec {
//
//	ports {
//
//	target_port
//
// = "80"
//
//	published_port = "8080"
//
//	}
//
//	} } then the import command is as follows #!/bin/bash
//
// ```sh
//
//	$ pulumi import docker:index/service:Service foo 4pcphbxkfn2rffhbhe6czytgi
//
// ```
type Service struct {
	pulumi.CustomResourceState

	// Configuration for the authentication for pulling the images of the service
	Auth ServiceAuthPtrOutput `pulumi:"auth"`
	// A configuration to ensure that a service converges aka reaches the desired that of all task up and running
	ConvergeConfig ServiceConvergeConfigPtrOutput `pulumi:"convergeConfig"`
	// Properties that can be configured to access and load balance a service
	EndpointSpec ServiceEndpointSpecOutput `pulumi:"endpointSpec"`
	// User-defined key/value metadata
	Labels ServiceLabelArrayOutput `pulumi:"labels"`
	// Scheduling mode for the service
	Mode ServiceModeOutput `pulumi:"mode"`
	// Name of the service
	Name pulumi.StringOutput `pulumi:"name"`
	// Specification for the rollback strategy of the service
	RollbackConfig ServiceRollbackConfigPtrOutput `pulumi:"rollbackConfig"`
	// User modifiable task configuration
	TaskSpec ServiceTaskSpecOutput `pulumi:"taskSpec"`
	// Specification for the update strategy of the service
	UpdateConfig ServiceUpdateConfigPtrOutput `pulumi:"updateConfig"`
}

// NewService registers a new resource with the given unique name, arguments, and options.
func NewService(ctx *pulumi.Context,
	name string, args *ServiceArgs, opts ...pulumi.ResourceOption) (*Service, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.TaskSpec == nil {
		return nil, errors.New("invalid value for required argument 'TaskSpec'")
	}
	var resource Service
	err := ctx.RegisterResource("docker:index/service:Service", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetService gets an existing Service resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceState, opts ...pulumi.ResourceOption) (*Service, error) {
	var resource Service
	err := ctx.ReadResource("docker:index/service:Service", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Service resources.
type serviceState struct {
	// Configuration for the authentication for pulling the images of the service
	Auth *ServiceAuth `pulumi:"auth"`
	// A configuration to ensure that a service converges aka reaches the desired that of all task up and running
	ConvergeConfig *ServiceConvergeConfig `pulumi:"convergeConfig"`
	// Properties that can be configured to access and load balance a service
	EndpointSpec *ServiceEndpointSpec `pulumi:"endpointSpec"`
	// User-defined key/value metadata
	Labels []ServiceLabel `pulumi:"labels"`
	// Scheduling mode for the service
	Mode *ServiceMode `pulumi:"mode"`
	// Name of the service
	Name *string `pulumi:"name"`
	// Specification for the rollback strategy of the service
	RollbackConfig *ServiceRollbackConfig `pulumi:"rollbackConfig"`
	// User modifiable task configuration
	TaskSpec *ServiceTaskSpec `pulumi:"taskSpec"`
	// Specification for the update strategy of the service
	UpdateConfig *ServiceUpdateConfig `pulumi:"updateConfig"`
}

type ServiceState struct {
	// Configuration for the authentication for pulling the images of the service
	Auth ServiceAuthPtrInput
	// A configuration to ensure that a service converges aka reaches the desired that of all task up and running
	ConvergeConfig ServiceConvergeConfigPtrInput
	// Properties that can be configured to access and load balance a service
	EndpointSpec ServiceEndpointSpecPtrInput
	// User-defined key/value metadata
	Labels ServiceLabelArrayInput
	// Scheduling mode for the service
	Mode ServiceModePtrInput
	// Name of the service
	Name pulumi.StringPtrInput
	// Specification for the rollback strategy of the service
	RollbackConfig ServiceRollbackConfigPtrInput
	// User modifiable task configuration
	TaskSpec ServiceTaskSpecPtrInput
	// Specification for the update strategy of the service
	UpdateConfig ServiceUpdateConfigPtrInput
}

func (ServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceState)(nil)).Elem()
}

type serviceArgs struct {
	// Configuration for the authentication for pulling the images of the service
	Auth *ServiceAuth `pulumi:"auth"`
	// A configuration to ensure that a service converges aka reaches the desired that of all task up and running
	ConvergeConfig *ServiceConvergeConfig `pulumi:"convergeConfig"`
	// Properties that can be configured to access and load balance a service
	EndpointSpec *ServiceEndpointSpec `pulumi:"endpointSpec"`
	// User-defined key/value metadata
	Labels []ServiceLabel `pulumi:"labels"`
	// Scheduling mode for the service
	Mode *ServiceMode `pulumi:"mode"`
	// Name of the service
	Name string `pulumi:"name"`
	// Specification for the rollback strategy of the service
	RollbackConfig *ServiceRollbackConfig `pulumi:"rollbackConfig"`
	// User modifiable task configuration
	TaskSpec ServiceTaskSpec `pulumi:"taskSpec"`
	// Specification for the update strategy of the service
	UpdateConfig *ServiceUpdateConfig `pulumi:"updateConfig"`
}

// The set of arguments for constructing a Service resource.
type ServiceArgs struct {
	// Configuration for the authentication for pulling the images of the service
	Auth ServiceAuthPtrInput
	// A configuration to ensure that a service converges aka reaches the desired that of all task up and running
	ConvergeConfig ServiceConvergeConfigPtrInput
	// Properties that can be configured to access and load balance a service
	EndpointSpec ServiceEndpointSpecPtrInput
	// User-defined key/value metadata
	Labels ServiceLabelArrayInput
	// Scheduling mode for the service
	Mode ServiceModePtrInput
	// Name of the service
	Name pulumi.StringInput
	// Specification for the rollback strategy of the service
	RollbackConfig ServiceRollbackConfigPtrInput
	// User modifiable task configuration
	TaskSpec ServiceTaskSpecInput
	// Specification for the update strategy of the service
	UpdateConfig ServiceUpdateConfigPtrInput
}

func (ServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceArgs)(nil)).Elem()
}

type ServiceInput interface {
	pulumi.Input

	ToServiceOutput() ServiceOutput
	ToServiceOutputWithContext(ctx context.Context) ServiceOutput
}

func (*Service) ElementType() reflect.Type {
	return reflect.TypeOf((**Service)(nil)).Elem()
}

func (i *Service) ToServiceOutput() ServiceOutput {
	return i.ToServiceOutputWithContext(context.Background())
}

func (i *Service) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceOutput)
}

// ServiceArrayInput is an input type that accepts ServiceArray and ServiceArrayOutput values.
// You can construct a concrete instance of `ServiceArrayInput` via:
//
//	ServiceArray{ ServiceArgs{...} }
type ServiceArrayInput interface {
	pulumi.Input

	ToServiceArrayOutput() ServiceArrayOutput
	ToServiceArrayOutputWithContext(context.Context) ServiceArrayOutput
}

type ServiceArray []ServiceInput

func (ServiceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Service)(nil)).Elem()
}

func (i ServiceArray) ToServiceArrayOutput() ServiceArrayOutput {
	return i.ToServiceArrayOutputWithContext(context.Background())
}

func (i ServiceArray) ToServiceArrayOutputWithContext(ctx context.Context) ServiceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceArrayOutput)
}

// ServiceMapInput is an input type that accepts ServiceMap and ServiceMapOutput values.
// You can construct a concrete instance of `ServiceMapInput` via:
//
//	ServiceMap{ "key": ServiceArgs{...} }
type ServiceMapInput interface {
	pulumi.Input

	ToServiceMapOutput() ServiceMapOutput
	ToServiceMapOutputWithContext(context.Context) ServiceMapOutput
}

type ServiceMap map[string]ServiceInput

func (ServiceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Service)(nil)).Elem()
}

func (i ServiceMap) ToServiceMapOutput() ServiceMapOutput {
	return i.ToServiceMapOutputWithContext(context.Background())
}

func (i ServiceMap) ToServiceMapOutputWithContext(ctx context.Context) ServiceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceMapOutput)
}

type ServiceOutput struct{ *pulumi.OutputState }

func (ServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Service)(nil)).Elem()
}

func (o ServiceOutput) ToServiceOutput() ServiceOutput {
	return o
}

func (o ServiceOutput) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return o
}

// Configuration for the authentication for pulling the images of the service
func (o ServiceOutput) Auth() ServiceAuthPtrOutput {
	return o.ApplyT(func(v *Service) ServiceAuthPtrOutput { return v.Auth }).(ServiceAuthPtrOutput)
}

// A configuration to ensure that a service converges aka reaches the desired that of all task up and running
func (o ServiceOutput) ConvergeConfig() ServiceConvergeConfigPtrOutput {
	return o.ApplyT(func(v *Service) ServiceConvergeConfigPtrOutput { return v.ConvergeConfig }).(ServiceConvergeConfigPtrOutput)
}

// Properties that can be configured to access and load balance a service
func (o ServiceOutput) EndpointSpec() ServiceEndpointSpecOutput {
	return o.ApplyT(func(v *Service) ServiceEndpointSpecOutput { return v.EndpointSpec }).(ServiceEndpointSpecOutput)
}

// User-defined key/value metadata
func (o ServiceOutput) Labels() ServiceLabelArrayOutput {
	return o.ApplyT(func(v *Service) ServiceLabelArrayOutput { return v.Labels }).(ServiceLabelArrayOutput)
}

// Scheduling mode for the service
func (o ServiceOutput) Mode() ServiceModeOutput {
	return o.ApplyT(func(v *Service) ServiceModeOutput { return v.Mode }).(ServiceModeOutput)
}

// Name of the service
func (o ServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specification for the rollback strategy of the service
func (o ServiceOutput) RollbackConfig() ServiceRollbackConfigPtrOutput {
	return o.ApplyT(func(v *Service) ServiceRollbackConfigPtrOutput { return v.RollbackConfig }).(ServiceRollbackConfigPtrOutput)
}

// User modifiable task configuration
func (o ServiceOutput) TaskSpec() ServiceTaskSpecOutput {
	return o.ApplyT(func(v *Service) ServiceTaskSpecOutput { return v.TaskSpec }).(ServiceTaskSpecOutput)
}

// Specification for the update strategy of the service
func (o ServiceOutput) UpdateConfig() ServiceUpdateConfigPtrOutput {
	return o.ApplyT(func(v *Service) ServiceUpdateConfigPtrOutput { return v.UpdateConfig }).(ServiceUpdateConfigPtrOutput)
}

type ServiceArrayOutput struct{ *pulumi.OutputState }

func (ServiceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Service)(nil)).Elem()
}

func (o ServiceArrayOutput) ToServiceArrayOutput() ServiceArrayOutput {
	return o
}

func (o ServiceArrayOutput) ToServiceArrayOutputWithContext(ctx context.Context) ServiceArrayOutput {
	return o
}

func (o ServiceArrayOutput) Index(i pulumi.IntInput) ServiceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Service {
		return vs[0].([]*Service)[vs[1].(int)]
	}).(ServiceOutput)
}

type ServiceMapOutput struct{ *pulumi.OutputState }

func (ServiceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Service)(nil)).Elem()
}

func (o ServiceMapOutput) ToServiceMapOutput() ServiceMapOutput {
	return o
}

func (o ServiceMapOutput) ToServiceMapOutputWithContext(ctx context.Context) ServiceMapOutput {
	return o
}

func (o ServiceMapOutput) MapIndex(k pulumi.StringInput) ServiceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Service {
		return vs[0].(map[string]*Service)[vs[1].(string)]
	}).(ServiceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceInput)(nil)).Elem(), &Service{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceArrayInput)(nil)).Elem(), ServiceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceMapInput)(nil)).Elem(), ServiceMap{})
	pulumi.RegisterOutputType(ServiceOutput{})
	pulumi.RegisterOutputType(ServiceArrayOutput{})
	pulumi.RegisterOutputType(ServiceMapOutput{})
}
