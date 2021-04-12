module github.com/pulumi/pulumi-docker/provider/v2

go 1.16

require (
	github.com/hashicorp/terraform-plugin-sdk v1.7.0
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.23.0
	github.com/pulumi/pulumi/sdk/v2 v2.24.1
	github.com/terraform-providers/terraform-provider-docker v0.0.0-20201111135144-b8d28d67e632
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/hashicorp/go-getter v1.5.0 => github.com/hashicorp/go-getter v1.4.0
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
	github.com/terraform-providers/terraform-provider-docker => github.com/kreuzwerker/terraform-provider-docker v0.0.0-20210122201038-8da701db26ba
	golang.org/x/sys => golang.org/x/sys v0.0.0-20190916202348-b4ddaad3f8a3
)
