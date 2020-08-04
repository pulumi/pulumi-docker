module github.com/pulumi/pulumi-docker/provider/v2

go 1.14

require (
	github.com/docker/go-units v0.3.3 // indirect
	github.com/hashicorp/terraform-plugin-sdk v1.7.0
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.5.4
	github.com/pulumi/pulumi/sdk/v2 v2.5.1-0.20200701223250-45d2fa95d60b
	github.com/terraform-providers/terraform-provider-docker v1.2.1-0.20200803235045-85aa923d5041
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
)
