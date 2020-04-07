module github.com/pulumi/pulumi-docker/provider

go 1.13

require (
	github.com/apparentlymart/go-dump v0.0.0-20190214190832-042adf3cf4a0 // indirect
	github.com/docker/go-units v0.3.3 // indirect
	github.com/hashicorp/terraform-plugin-sdk v1.0.0
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/pulumi/pulumi-terraform-bridge v1.8.4
	github.com/pulumi/pulumi/pkg v1.13.1 // indirect
	github.com/pulumi/pulumi/sdk v1.13.1
	github.com/terraform-providers/terraform-provider-docker v1.2.1-0.20200210195100-e2a14e7e7cc6
	github.com/vmihailenco/msgpack v4.0.1+incompatible // indirect
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
)
