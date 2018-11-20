// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package swarm

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages the secrets of a Docker service in a swarm.
type Secret struct {
	s *pulumi.ResourceState
}

// NewSecret registers a new resource with the given unique name, arguments, and options.
func NewSecret(ctx *pulumi.Context,
	name string, args *SecretArgs, opts ...pulumi.ResourceOpt) (*Secret, error) {
	if args == nil || args.Data == nil {
		return nil, errors.New("missing required argument 'Data'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["data"] = nil
		inputs["labels"] = nil
		inputs["name"] = nil
	} else {
		inputs["data"] = args.Data
		inputs["labels"] = args.Labels
		inputs["name"] = args.Name
	}
	s, err := ctx.RegisterResource("docker:swarm/secret:Secret", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Secret{s: s}, nil
}

// GetSecret gets an existing Secret resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecret(ctx *pulumi.Context,
	name string, id pulumi.ID, state *SecretState, opts ...pulumi.ResourceOpt) (*Secret, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["data"] = state.Data
		inputs["labels"] = state.Labels
		inputs["name"] = state.Name
	}
	s, err := ctx.ReadResource("docker:swarm/secret:Secret", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Secret{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Secret) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Secret) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The base64 encoded data of the secret.
func (r *Secret) Data() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["data"])
}

// User-defined key/value metadata.
func (r *Secret) Labels() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["labels"])
}

// The name of the Docker secret.
func (r *Secret) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// Input properties used for looking up and filtering Secret resources.
type SecretState struct {
	// The base64 encoded data of the secret.
	Data interface{}
	// User-defined key/value metadata.
	Labels interface{}
	// The name of the Docker secret.
	Name interface{}
}

// The set of arguments for constructing a Secret resource.
type SecretArgs struct {
	// The base64 encoded data of the secret.
	Data interface{}
	// User-defined key/value metadata.
	Labels interface{}
	// The name of the Docker secret.
	Name interface{}
}
