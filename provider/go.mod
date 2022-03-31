module github.com/pulumi/pulumi-docker/provider/v3

go 1.16

require (
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.20.0
	github.com/pulumi/pulumi/sdk/v3 v3.27.0
	github.com/terraform-providers/terraform-provider-docker/shim v0.0.0
)

replace (
	github.com/hashicorp/go-getter v1.5.0 => github.com/hashicorp/go-getter v1.4.0
	github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20211019194827-62530c6537a4
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
	github.com/terraform-providers/terraform-provider-docker => github.com/kreuzwerker/terraform-provider-docker v0.0.0-20220124093807-a8517ea3024f
	github.com/terraform-providers/terraform-provider-docker/shim => ./shim
)
