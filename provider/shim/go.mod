module github.com/terraform-providers/terraform-provider-docker/shim

go 1.19

require (
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.24.0
	github.com/terraform-providers/terraform-provider-docker v0.0.0
)

replace github.com/terraform-providers/terraform-provider-docker => github.com/kreuzwerker/terraform-provider-docker v0.0.0-20221102105757-3489fbbafc7b

require (
	github.com/Azure/go-ansiterm v0.0.0-20210617225240-d185dfc1b5a1 // indirect
	github.com/Microsoft/go-winio v0.5.1 // indirect
	github.com/Microsoft/hcsshim v0.9.2 // indirect
	github.com/agext/levenshtein v1.2.3 // indirect
	github.com/apparentlymart/go-textseg/v13 v13.0.0 // indirect
	github.com/containerd/cgroups v1.0.3 // indirect
	github.com/containerd/containerd v1.6.3-0.20220401172941-5ff8fce1fcc6 // indirect
	github.com/containerd/typeurl v1.0.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/docker/cli v20.10.21+incompatible // indirect
	github.com/docker/distribution v2.8.1+incompatible // indirect
	github.com/docker/docker v20.10.21+incompatible // indirect
	github.com/docker/docker-credential-helpers v0.6.4 // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.5.0 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/go-logr/logr v1.2.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-checkpoint v0.5.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-cty v1.4.1-0.20200414143053-d3edf31b6320 // indirect
	github.com/hashicorp/go-hclog v1.2.1 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-plugin v1.4.4 // indirect
	github.com/hashicorp/go-uuid v1.0.3 // indirect
	github.com/hashicorp/go-version v1.6.0 // indirect
	github.com/hashicorp/hc-install v0.4.0 // indirect
	github.com/hashicorp/hcl/v2 v2.14.1 // indirect
	github.com/hashicorp/logutils v1.0.0 // indirect
	github.com/hashicorp/terraform-exec v0.17.3 // indirect
	github.com/hashicorp/terraform-json v0.14.0 // indirect
	github.com/hashicorp/terraform-plugin-go v0.14.0 // indirect
	github.com/hashicorp/terraform-plugin-log v0.7.0 // indirect
	github.com/hashicorp/terraform-registry-address v0.0.0-20220623143253-7d51757b572c // indirect
	github.com/hashicorp/terraform-svchost v0.0.0-20200729002733-f050f53b9734 // indirect
	github.com/hashicorp/yamux v0.0.0-20181012175058-2f1d1f20f75d // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/mitchellh/go-wordwrap v1.0.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/moby/buildkit v0.10.5 // indirect
	github.com/moby/sys/mount v0.3.0 // indirect
	github.com/moby/sys/mountinfo v0.6.0 // indirect
	github.com/moby/sys/symlink v0.2.0 // indirect
	github.com/moby/term v0.0.0-20210619224110-3f7ff695adc6 // indirect
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/oklog/run v1.0.0 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.3-0.20211202183452-c5a74bcca799 // indirect
	github.com/opencontainers/runc v1.1.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/sirupsen/logrus v1.9.0 // indirect
	github.com/vmihailenco/msgpack v4.0.4+incompatible // indirect
	github.com/vmihailenco/msgpack/v4 v4.3.12 // indirect
	github.com/vmihailenco/tagparser v0.1.1 // indirect
	github.com/zclconf/go-cty v1.11.0 // indirect
	go.opencensus.io v0.23.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.29.0 // indirect
	go.opentelemetry.io/otel v1.4.1 // indirect
	go.opentelemetry.io/otel/trace v1.4.1 // indirect
	golang.org/x/crypto v0.0.0-20220722155217-630584e8d5aa // indirect
	golang.org/x/net v0.0.0-20220722155237-a158d28d115b // indirect
	golang.org/x/sys v0.0.0-20220915200043-7b5979e65e41 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/time v0.0.0-20210723032227-1f47c861a9ac // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20220519153652-3a47de7e79bd // indirect
	google.golang.org/grpc v1.48.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)
