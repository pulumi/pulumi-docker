// Copyright 2016-2020, Pulumi Corporation.

package docker

//     var nestedType = reflect.TypeOf((*Nested)(nil)).Elem()
//
//     type NestedInput interface {
//         pulumi.Input
//
//         ToNestedOutput() NestedOutput
//         ToNestedOutputWithContext(context.Context) NestedOutput
//     }
//
//     type Nested struct {
//         Foo int `pulumi:"foo"`
//         Bar string `pulumi:"bar"`
//     }
//
//     type NestedInputValue struct {
//         Foo pulumi.IntInput `pulumi:"foo"`
//         Bar pulumi.StringInput `pulumi:"bar"`
//     }
//
//     func (NestedInputValue) ElementType() reflect.Type {
//         return nestedType
//     }
//
//     func (v NestedInputValue) ToNestedOutput() NestedOutput {
//         return pulumi.ToOutput(v).(NestedOutput)
//     }
//
//     func (v NestedInputValue) ToNestedOutputWithContext(ctx context.Context) NestedOutput {
//         return pulumi.ToOutputWithContext(ctx, v).(NestedOutput)
//     }
//
//     type NestedOutput struct { *pulumi.OutputState }
//
//     func (NestedOutput) ElementType() reflect.Type {
//         return nestedType
//     }
//
//     func (o NestedOutput) ToNestedOutput() NestedOutput {
//         return o
//     }
//
//     func (o NestedOutput) ToNestedOutputWithContext(ctx context.Context) NestedOutput {
//         return o
//     }
