module github.com/terraform-providers/terraform-provider-docker/shim

go 1.15

require (
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.6.1
	github.com/terraform-providers/terraform-provider-docker v0.0.0-20201111135144-b8d28d67e632
)

replace github.com/terraform-providers/terraform-provider-docker => github.com/kreuzwerker/terraform-provider-docker v0.0.0-20210526070551-22c490336c7b
