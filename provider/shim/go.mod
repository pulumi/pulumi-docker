module github.com/terraform-providers/terraform-provider-docker/shim

go 1.15

require (
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.7.0
	github.com/terraform-providers/terraform-provider-docker v0.0.0
)

replace github.com/terraform-providers/terraform-provider-docker => github.com/kreuzwerker/terraform-provider-docker v0.0.0-20210811081153-7966d3f57fa6
