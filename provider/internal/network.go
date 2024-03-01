package internal

import "github.com/pulumi/pulumi-go-provider/infer"

var _ = (infer.Enum[NetworkMode])((*NetworkMode)(nil))

type NetworkMode string

const (
	NetworkModeDefault NetworkMode = "default"
	NetworkModeHost    NetworkMode = "host"
	NetworkModeNone    NetworkMode = "none"
)

func (NetworkMode) Values() []infer.EnumValue[NetworkMode] {
	return []infer.EnumValue[NetworkMode]{
		{
			Value:       NetworkModeDefault,
			Description: "The default sandbox network mode.",
		},
		{
			Value:       NetworkModeHost,
			Description: "Host network mode.",
		},
		{
			Value:       NetworkModeNone,
			Description: "Disable network access.",
		},
	}
}
