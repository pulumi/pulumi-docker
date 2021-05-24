module github.com/pulumi/pulumi-docker/provider/v3

go 1.16

require (
	github.com/hashicorp/terraform-plugin-sdk v1.7.0
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.0.0
	github.com/pulumi/pulumi/pkg/v3 v3.0.0
	github.com/pulumi/pulumi/sdk/v3 v3.0.0
	github.com/terraform-providers/terraform-provider-docker/shim v0.0.0
)

replace (
	github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20201218231525-9cca98608a5e
	github.com/hashicorp/go-getter v1.5.0 => github.com/hashicorp/go-getter v1.4.0
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
	github.com/terraform-providers/terraform-provider-docker => github.com/kreuzwerker/terraform-provider-docker v0.0.0-20210523100431-e73a0578496b
	github.com/terraform-providers/terraform-provider-docker/shim => ./shim
)
