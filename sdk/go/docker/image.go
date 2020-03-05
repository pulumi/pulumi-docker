// Copyright 2016-2020, Pulumi Corporation.

package docker

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// ImageRegistry contains credentials for the docker registry.
type ImageRegistry struct {
	// Docker registry server URL to push to.  Some common values include:
	// DockerHub: `docker.io` or `https://index.docker.io/v1`
	// Azure Container Registry: `<name>.azurecr.io`
	// AWS Elastic Container Registry: `<account>.dkr.ecr.us-east-2.amazonaws.com`
	// Google Container Registry: `<name>.gcr.io`
	Server pulumi.StringInput `pulumi:"server"`

	// Username for login to the target Docker registry.
	Username pulumi.StringInput `pulumi:"username"`

	// Password for login to the target Docker registry.
	Password pulumi.StringInput `pulumi:"password"`
}

type imageRegistry struct {
	Server   string
	Username string
	Password string
}
