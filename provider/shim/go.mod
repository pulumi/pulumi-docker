module github.com/terraform-providers/terraform-provider-docker/shim

go 1.15

require (
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.10.1
	github.com/terraform-providers/terraform-provider-docker v0.0.0
)

replace github.com/terraform-providers/terraform-provider-docker => github.com/kreuzwerker/terraform-provider-docker v0.0.0-20220124093807-a8517ea3024f
