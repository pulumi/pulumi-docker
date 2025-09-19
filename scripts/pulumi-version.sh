# version=$(cd provider && go list -f "{{slice .Version 1}}" -m github.com/pulumi/pulumi/pkg/v3)
version="3.190.0"
echo "PULUMI_VERSION=$version"
export PULUMI_VERSION=$version
