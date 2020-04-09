module github.com/pulumi/pulumi-docker/examples/v2

go 1.13

require (
	github.com/pulumi/pulumi-docker/sdk v0.0.0-20200407192733-912752761bfd
	github.com/pulumi/pulumi/pkg v1.14.0
	github.com/pulumi/pulumi/sdk v1.14.0
	github.com/stretchr/testify v1.5.1
)

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
