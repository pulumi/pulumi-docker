// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

// PEM-encoded content of Docker host CA certificate
func GetCaMaterial(ctx *pulumi.Context) string {
	return config.Get(ctx, "docker:caMaterial")
}

// PEM-encoded content of Docker client certificate
func GetCertMaterial(ctx *pulumi.Context) string {
	return config.Get(ctx, "docker:certMaterial")
}

// Path to directory with Docker TLS config
func GetCertPath(ctx *pulumi.Context) string {
	return config.Get(ctx, "docker:certPath")
}

// The Docker daemon address
func GetHost(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "docker:host")
	if err == nil {
		return v
	}
	return getEnvOrDefault("unix:///var/run/docker.sock", nil, "DOCKER_HOST").(string)
}

// PEM-encoded content of Docker client private key
func GetKeyMaterial(ctx *pulumi.Context) string {
	return config.Get(ctx, "docker:keyMaterial")
}
func GetRegistryAuth(ctx *pulumi.Context) string {
	return config.Get(ctx, "docker:registryAuth")
}

// Additional SSH option flags to be appended when using `ssh://` protocol
func GetSshOpts(ctx *pulumi.Context) string {
	return config.Get(ctx, "docker:sshOpts")
}
