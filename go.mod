module github.com/pulumi/pulumi-docker

go 1.13

require (
	github.com/Masterminds/semver v1.5.0
	github.com/apparentlymart/go-dump v0.0.0-20190214190832-042adf3cf4a0 // indirect
	github.com/aws/aws-sdk-go v1.25.3 // indirect
	github.com/docker/go-units v0.3.3 // indirect
	github.com/hashicorp/terraform-plugin-sdk v1.0.0
	github.com/hashicorp/vault/api v1.0.5-0.20190730042357-746c0b111519 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/pkg/errors v0.8.1
<<<<<<< HEAD
	github.com/pulumi/pulumi v1.12.2-0.20200313203005-b6e5d2737d4f
	github.com/pulumi/pulumi-terraform-bridge v1.8.2
=======
	github.com/pulumi/pulumi v1.12.2-0.20200312230059-ef6f0d4de4e7
	github.com/pulumi/pulumi-terraform-bridge v1.6.5
>>>>>>> e4cf0ed... just edit go.mod
	github.com/stretchr/testify v1.4.1-0.20191106224347-f1bd0923b832
	github.com/terraform-providers/terraform-provider-docker v1.2.1-0.20200210195100-e2a14e7e7cc6
	github.com/vmihailenco/msgpack v4.0.1+incompatible // indirect
	github.com/xanzy/ssh-agent v0.2.1 // indirect
	golang.org/x/net v0.0.0-20191009170851-d66e71096ffb // indirect
	golang.org/x/sys v0.0.0-20190804053845-51ab0e2deafa // indirect
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
)
