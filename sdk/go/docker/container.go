// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package docker

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages the lifecycle of a Docker container.
type Container struct {
	s *pulumi.ResourceState
}

// NewContainer registers a new resource with the given unique name, arguments, and options.
func NewContainer(ctx *pulumi.Context,
	name string, args *ContainerArgs, opts ...pulumi.ResourceOpt) (*Container, error) {
	if args == nil || args.Image == nil {
		return nil, errors.New("missing required argument 'Image'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["attach"] = nil
		inputs["capabilities"] = nil
		inputs["command"] = nil
		inputs["cpuSet"] = nil
		inputs["cpuShares"] = nil
		inputs["destroyGraceSeconds"] = nil
		inputs["devices"] = nil
		inputs["dns"] = nil
		inputs["dnsOpts"] = nil
		inputs["dnsSearches"] = nil
		inputs["domainname"] = nil
		inputs["entrypoints"] = nil
		inputs["envs"] = nil
		inputs["healthcheck"] = nil
		inputs["hosts"] = nil
		inputs["hostname"] = nil
		inputs["image"] = nil
		inputs["labels"] = nil
		inputs["links"] = nil
		inputs["logDriver"] = nil
		inputs["logOpts"] = nil
		inputs["logs"] = nil
		inputs["maxRetryCount"] = nil
		inputs["memory"] = nil
		inputs["memorySwap"] = nil
		inputs["mustRun"] = nil
		inputs["name"] = nil
		inputs["networkAliases"] = nil
		inputs["networkMode"] = nil
		inputs["networks"] = nil
		inputs["networks_advanced"] = nil
		inputs["pidMode"] = nil
		inputs["ports"] = nil
		inputs["privileged"] = nil
		inputs["publishAllPorts"] = nil
		inputs["restart"] = nil
		inputs["rm"] = nil
		inputs["start"] = nil
		inputs["ulimits"] = nil
		inputs["uploads"] = nil
		inputs["user"] = nil
		inputs["usernsMode"] = nil
		inputs["volumes"] = nil
	} else {
		inputs["attach"] = args.Attach
		inputs["capabilities"] = args.Capabilities
		inputs["command"] = args.Command
		inputs["cpuSet"] = args.CpuSet
		inputs["cpuShares"] = args.CpuShares
		inputs["destroyGraceSeconds"] = args.DestroyGraceSeconds
		inputs["devices"] = args.Devices
		inputs["dns"] = args.Dns
		inputs["dnsOpts"] = args.DnsOpts
		inputs["dnsSearches"] = args.DnsSearches
		inputs["domainname"] = args.Domainname
		inputs["entrypoints"] = args.Entrypoints
		inputs["envs"] = args.Envs
		inputs["healthcheck"] = args.Healthcheck
		inputs["hosts"] = args.Hosts
		inputs["hostname"] = args.Hostname
		inputs["image"] = args.Image
		inputs["labels"] = args.Labels
		inputs["links"] = args.Links
		inputs["logDriver"] = args.LogDriver
		inputs["logOpts"] = args.LogOpts
		inputs["logs"] = args.Logs
		inputs["maxRetryCount"] = args.MaxRetryCount
		inputs["memory"] = args.Memory
		inputs["memorySwap"] = args.MemorySwap
		inputs["mustRun"] = args.MustRun
		inputs["name"] = args.Name
		inputs["networkAliases"] = args.NetworkAliases
		inputs["networkMode"] = args.NetworkMode
		inputs["networks"] = args.Networks
		inputs["networks_advanced"] = args.Networks_advanced
		inputs["pidMode"] = args.PidMode
		inputs["ports"] = args.Ports
		inputs["privileged"] = args.Privileged
		inputs["publishAllPorts"] = args.PublishAllPorts
		inputs["restart"] = args.Restart
		inputs["rm"] = args.Rm
		inputs["start"] = args.Start
		inputs["ulimits"] = args.Ulimits
		inputs["uploads"] = args.Uploads
		inputs["user"] = args.User
		inputs["usernsMode"] = args.UsernsMode
		inputs["volumes"] = args.Volumes
	}
	inputs["bridge"] = nil
	inputs["containerLogs"] = nil
	inputs["exitCode"] = nil
	inputs["gateway"] = nil
	inputs["ipAddress"] = nil
	inputs["ipPrefixLength"] = nil
	inputs["networkDatas"] = nil
	s, err := ctx.RegisterResource("docker:index/container:Container", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Container{s: s}, nil
}

// GetContainer gets an existing Container resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContainer(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ContainerState, opts ...pulumi.ResourceOpt) (*Container, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["attach"] = state.Attach
		inputs["bridge"] = state.Bridge
		inputs["capabilities"] = state.Capabilities
		inputs["command"] = state.Command
		inputs["containerLogs"] = state.ContainerLogs
		inputs["cpuSet"] = state.CpuSet
		inputs["cpuShares"] = state.CpuShares
		inputs["destroyGraceSeconds"] = state.DestroyGraceSeconds
		inputs["devices"] = state.Devices
		inputs["dns"] = state.Dns
		inputs["dnsOpts"] = state.DnsOpts
		inputs["dnsSearches"] = state.DnsSearches
		inputs["domainname"] = state.Domainname
		inputs["entrypoints"] = state.Entrypoints
		inputs["envs"] = state.Envs
		inputs["exitCode"] = state.ExitCode
		inputs["gateway"] = state.Gateway
		inputs["healthcheck"] = state.Healthcheck
		inputs["hosts"] = state.Hosts
		inputs["hostname"] = state.Hostname
		inputs["image"] = state.Image
		inputs["ipAddress"] = state.IpAddress
		inputs["ipPrefixLength"] = state.IpPrefixLength
		inputs["labels"] = state.Labels
		inputs["links"] = state.Links
		inputs["logDriver"] = state.LogDriver
		inputs["logOpts"] = state.LogOpts
		inputs["logs"] = state.Logs
		inputs["maxRetryCount"] = state.MaxRetryCount
		inputs["memory"] = state.Memory
		inputs["memorySwap"] = state.MemorySwap
		inputs["mustRun"] = state.MustRun
		inputs["name"] = state.Name
		inputs["networkAliases"] = state.NetworkAliases
		inputs["networkDatas"] = state.NetworkDatas
		inputs["networkMode"] = state.NetworkMode
		inputs["networks"] = state.Networks
		inputs["networks_advanced"] = state.Networks_advanced
		inputs["pidMode"] = state.PidMode
		inputs["ports"] = state.Ports
		inputs["privileged"] = state.Privileged
		inputs["publishAllPorts"] = state.PublishAllPorts
		inputs["restart"] = state.Restart
		inputs["rm"] = state.Rm
		inputs["start"] = state.Start
		inputs["ulimits"] = state.Ulimits
		inputs["uploads"] = state.Uploads
		inputs["user"] = state.User
		inputs["usernsMode"] = state.UsernsMode
		inputs["volumes"] = state.Volumes
	}
	s, err := ctx.ReadResource("docker:index/container:Container", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Container{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Container) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Container) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// If true attach to the container after its creation and waits the end of his execution.
func (r *Container) Attach() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["attach"])
}

// The network bridge of the container as read from its NetworkSettings.
func (r *Container) Bridge() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["bridge"])
}

// See Capabilities below for details.
func (r *Container) Capabilities() *pulumi.Output {
	return r.s.State["capabilities"]
}

// The command to use to start the
// container. For example, to run `/usr/bin/myprogram -f baz.conf` set the
// command to be `["/usr/bin/myprogram", "-f", "baz.conf"]`.
func (r *Container) Command() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["command"])
}

// The logs of the container if its execution is done (`attach` must be disabled).
func (r *Container) ContainerLogs() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["containerLogs"])
}

// A comma-separated list or hyphen-separated range of CPUs a container can use, e.g. `0-1`.
func (r *Container) CpuSet() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["cpuSet"])
}

// CPU shares (relative weight) for the container.
func (r *Container) CpuShares() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["cpuShares"])
}

// If defined will attempt to stop the container before destroying. Container will be destroyed after `n` seconds or on successful stop.
func (r *Container) DestroyGraceSeconds() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["destroyGraceSeconds"])
}

// See Devices below for details.
func (r *Container) Devices() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["devices"])
}

// Set of DNS servers.
func (r *Container) Dns() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["dns"])
}

// Set of DNS options used by the DNS provider(s), see `resolv.conf` documentation for valid list of options.
func (r *Container) DnsOpts() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["dnsOpts"])
}

// Set of DNS search domains that are used when bare unqualified hostnames are used inside of the container.
func (r *Container) DnsSearches() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["dnsSearches"])
}

// Domain name of the container.
func (r *Container) Domainname() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["domainname"])
}

// The command to use as the
// Entrypoint for the container. The Entrypoint allows you to configure a
// container to run as an executable. For example, to run `/usr/bin/myprogram`
// when starting a container, set the entrypoint to be
// `["/usr/bin/myprogram"]`.
func (r *Container) Entrypoints() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["entrypoints"])
}

// Environment variables to set.
func (r *Container) Envs() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["envs"])
}

// The exit code of the container if its execution is done (`must_run` must be disabled).
func (r *Container) ExitCode() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["exitCode"])
}

// *Deprecated:* Use `network_data` instead. The network gateway of the container as read from its
// NetworkSettings.
func (r *Container) Gateway() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["gateway"])
}

// See Healthcheck below for details.
func (r *Container) Healthcheck() *pulumi.Output {
	return r.s.State["healthcheck"]
}

// Hostname to add.
func (r *Container) Hosts() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["hosts"])
}

// Hostname of the container.
func (r *Container) Hostname() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["hostname"])
}

// The ID of the image to back this container.
// The easiest way to get this value is to use the `docker_image` resource
// as is shown in the example above.
func (r *Container) Image() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["image"])
}

// *Deprecated:* Use `network_data` instead. The IP address of the container's first network it.
func (r *Container) IpAddress() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["ipAddress"])
}

// *Deprecated:* Use `network_data` instead. The IP prefix length of the container as read from its
// NetworkSettings.
func (r *Container) IpPrefixLength() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["ipPrefixLength"])
}

// Key/value pairs to set as labels on the
// container.
func (r *Container) Labels() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["labels"])
}

// Set of links for link based
// connectivity between containers that are running on the same host.
func (r *Container) Links() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["links"])
}

// The logging driver to use for the container.
// Defaults to "json-file".
func (r *Container) LogDriver() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["logDriver"])
}

// Key/value pairs to use as options for
// the logging driver.
func (r *Container) LogOpts() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["logOpts"])
}

// Save the container logs (`attach` must be enabled).
func (r *Container) Logs() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["logs"])
}

// The maximum amount of times to an attempt
// a restart when `restart` is set to "on-failure"
func (r *Container) MaxRetryCount() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["maxRetryCount"])
}

// The memory limit for the container in MBs.
func (r *Container) Memory() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["memory"])
}

// The total memory limit (memory + swap) for the
// container in MBs. This setting may compute to `-1` after `terraform apply` if the target host doesn't support memory swap, when that is the case docker will use a soft limitation.
func (r *Container) MemorySwap() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["memorySwap"])
}

// If true, then the Docker container will be
// kept running. If false, then as long as the container exists, Terraform
// assumes it is successful.
func (r *Container) MustRun() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["mustRun"])
}

func (r *Container) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// Network aliases of the container for user-defined networks only. *Deprecated:* use `networks_advanced` instead.
func (r *Container) NetworkAliases() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["networkAliases"])
}

// (Map of a block) The IP addresses of the container on each
// network. Key are the network names, values are the IP addresses.
func (r *Container) NetworkDatas() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["networkDatas"])
}

// Network mode of the container.
func (r *Container) NetworkMode() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["networkMode"])
}

// Id of the networks in which the
// container is. *Deprecated:* use `networks_advanced` instead.
func (r *Container) Networks() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["networks"])
}

// See Networks Advanced below for details. If this block has priority to the deprecated `network_alias` and `network` properties.
func (r *Container) Networks_advanced() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["networks_advanced"])
}

// The PID (Process) Namespace mode for the container. Either `container:<name|id>` or `host`.
func (r *Container) PidMode() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["pidMode"])
}

// See Ports below for details.
func (r *Container) Ports() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["ports"])
}

// Run container in privileged mode.
func (r *Container) Privileged() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["privileged"])
}

// Publish all ports of the container.
func (r *Container) PublishAllPorts() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["publishAllPorts"])
}

// The restart policy for the container. Must be
// one of "no", "on-failure", "always", "unless-stopped".
func (r *Container) Restart() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["restart"])
}

// If true, then the container will be automatically removed after his execution. Terraform
// won't check this container after creation.
func (r *Container) Rm() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["rm"])
}

// If true, then the Docker container will be
// started after creation. If false, then the container is only created.
func (r *Container) Start() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["start"])
}

// See Ulimits below for
// details.
func (r *Container) Ulimits() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["ulimits"])
}

// See File Upload below for details.
func (r *Container) Uploads() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["uploads"])
}

// User used for run the first process. Format is
// `user` or `user:group` which user and group can be passed literraly or
// by name.
func (r *Container) User() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["user"])
}

// Sets the usernamespace mode for the container when usernamespace remapping option is enabled.
func (r *Container) UsernsMode() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["usernsMode"])
}

// See Volumes below for details.
func (r *Container) Volumes() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["volumes"])
}

// Input properties used for looking up and filtering Container resources.
type ContainerState struct {
	// If true attach to the container after its creation and waits the end of his execution.
	Attach interface{}
	// The network bridge of the container as read from its NetworkSettings.
	Bridge interface{}
	// See Capabilities below for details.
	Capabilities interface{}
	// The command to use to start the
	// container. For example, to run `/usr/bin/myprogram -f baz.conf` set the
	// command to be `["/usr/bin/myprogram", "-f", "baz.conf"]`.
	Command interface{}
	// The logs of the container if its execution is done (`attach` must be disabled).
	ContainerLogs interface{}
	// A comma-separated list or hyphen-separated range of CPUs a container can use, e.g. `0-1`.
	CpuSet interface{}
	// CPU shares (relative weight) for the container.
	CpuShares interface{}
	// If defined will attempt to stop the container before destroying. Container will be destroyed after `n` seconds or on successful stop.
	DestroyGraceSeconds interface{}
	// See Devices below for details.
	Devices interface{}
	// Set of DNS servers.
	Dns interface{}
	// Set of DNS options used by the DNS provider(s), see `resolv.conf` documentation for valid list of options.
	DnsOpts interface{}
	// Set of DNS search domains that are used when bare unqualified hostnames are used inside of the container.
	DnsSearches interface{}
	// Domain name of the container.
	Domainname interface{}
	// The command to use as the
	// Entrypoint for the container. The Entrypoint allows you to configure a
	// container to run as an executable. For example, to run `/usr/bin/myprogram`
	// when starting a container, set the entrypoint to be
	// `["/usr/bin/myprogram"]`.
	Entrypoints interface{}
	// Environment variables to set.
	Envs interface{}
	// The exit code of the container if its execution is done (`must_run` must be disabled).
	ExitCode interface{}
	// *Deprecated:* Use `network_data` instead. The network gateway of the container as read from its
	// NetworkSettings.
	Gateway interface{}
	// See Healthcheck below for details.
	Healthcheck interface{}
	// Hostname to add.
	Hosts interface{}
	// Hostname of the container.
	Hostname interface{}
	// The ID of the image to back this container.
	// The easiest way to get this value is to use the `docker_image` resource
	// as is shown in the example above.
	Image interface{}
	// *Deprecated:* Use `network_data` instead. The IP address of the container's first network it.
	IpAddress interface{}
	// *Deprecated:* Use `network_data` instead. The IP prefix length of the container as read from its
	// NetworkSettings.
	IpPrefixLength interface{}
	// Key/value pairs to set as labels on the
	// container.
	Labels interface{}
	// Set of links for link based
	// connectivity between containers that are running on the same host.
	Links interface{}
	// The logging driver to use for the container.
	// Defaults to "json-file".
	LogDriver interface{}
	// Key/value pairs to use as options for
	// the logging driver.
	LogOpts interface{}
	// Save the container logs (`attach` must be enabled).
	Logs interface{}
	// The maximum amount of times to an attempt
	// a restart when `restart` is set to "on-failure"
	MaxRetryCount interface{}
	// The memory limit for the container in MBs.
	Memory interface{}
	// The total memory limit (memory + swap) for the
	// container in MBs. This setting may compute to `-1` after `terraform apply` if the target host doesn't support memory swap, when that is the case docker will use a soft limitation.
	MemorySwap interface{}
	// If true, then the Docker container will be
	// kept running. If false, then as long as the container exists, Terraform
	// assumes it is successful.
	MustRun interface{}
	Name interface{}
	// Network aliases of the container for user-defined networks only. *Deprecated:* use `networks_advanced` instead.
	NetworkAliases interface{}
	// (Map of a block) The IP addresses of the container on each
	// network. Key are the network names, values are the IP addresses.
	NetworkDatas interface{}
	// Network mode of the container.
	NetworkMode interface{}
	// Id of the networks in which the
	// container is. *Deprecated:* use `networks_advanced` instead.
	Networks interface{}
	// See Networks Advanced below for details. If this block has priority to the deprecated `network_alias` and `network` properties.
	Networks_advanced interface{}
	// The PID (Process) Namespace mode for the container. Either `container:<name|id>` or `host`.
	PidMode interface{}
	// See Ports below for details.
	Ports interface{}
	// Run container in privileged mode.
	Privileged interface{}
	// Publish all ports of the container.
	PublishAllPorts interface{}
	// The restart policy for the container. Must be
	// one of "no", "on-failure", "always", "unless-stopped".
	Restart interface{}
	// If true, then the container will be automatically removed after his execution. Terraform
	// won't check this container after creation.
	Rm interface{}
	// If true, then the Docker container will be
	// started after creation. If false, then the container is only created.
	Start interface{}
	// See Ulimits below for
	// details.
	Ulimits interface{}
	// See File Upload below for details.
	Uploads interface{}
	// User used for run the first process. Format is
	// `user` or `user:group` which user and group can be passed literraly or
	// by name.
	User interface{}
	// Sets the usernamespace mode for the container when usernamespace remapping option is enabled.
	UsernsMode interface{}
	// See Volumes below for details.
	Volumes interface{}
}

// The set of arguments for constructing a Container resource.
type ContainerArgs struct {
	// If true attach to the container after its creation and waits the end of his execution.
	Attach interface{}
	// See Capabilities below for details.
	Capabilities interface{}
	// The command to use to start the
	// container. For example, to run `/usr/bin/myprogram -f baz.conf` set the
	// command to be `["/usr/bin/myprogram", "-f", "baz.conf"]`.
	Command interface{}
	// A comma-separated list or hyphen-separated range of CPUs a container can use, e.g. `0-1`.
	CpuSet interface{}
	// CPU shares (relative weight) for the container.
	CpuShares interface{}
	// If defined will attempt to stop the container before destroying. Container will be destroyed after `n` seconds or on successful stop.
	DestroyGraceSeconds interface{}
	// See Devices below for details.
	Devices interface{}
	// Set of DNS servers.
	Dns interface{}
	// Set of DNS options used by the DNS provider(s), see `resolv.conf` documentation for valid list of options.
	DnsOpts interface{}
	// Set of DNS search domains that are used when bare unqualified hostnames are used inside of the container.
	DnsSearches interface{}
	// Domain name of the container.
	Domainname interface{}
	// The command to use as the
	// Entrypoint for the container. The Entrypoint allows you to configure a
	// container to run as an executable. For example, to run `/usr/bin/myprogram`
	// when starting a container, set the entrypoint to be
	// `["/usr/bin/myprogram"]`.
	Entrypoints interface{}
	// Environment variables to set.
	Envs interface{}
	// See Healthcheck below for details.
	Healthcheck interface{}
	// Hostname to add.
	Hosts interface{}
	// Hostname of the container.
	Hostname interface{}
	// The ID of the image to back this container.
	// The easiest way to get this value is to use the `docker_image` resource
	// as is shown in the example above.
	Image interface{}
	// Key/value pairs to set as labels on the
	// container.
	Labels interface{}
	// Set of links for link based
	// connectivity between containers that are running on the same host.
	Links interface{}
	// The logging driver to use for the container.
	// Defaults to "json-file".
	LogDriver interface{}
	// Key/value pairs to use as options for
	// the logging driver.
	LogOpts interface{}
	// Save the container logs (`attach` must be enabled).
	Logs interface{}
	// The maximum amount of times to an attempt
	// a restart when `restart` is set to "on-failure"
	MaxRetryCount interface{}
	// The memory limit for the container in MBs.
	Memory interface{}
	// The total memory limit (memory + swap) for the
	// container in MBs. This setting may compute to `-1` after `terraform apply` if the target host doesn't support memory swap, when that is the case docker will use a soft limitation.
	MemorySwap interface{}
	// If true, then the Docker container will be
	// kept running. If false, then as long as the container exists, Terraform
	// assumes it is successful.
	MustRun interface{}
	Name interface{}
	// Network aliases of the container for user-defined networks only. *Deprecated:* use `networks_advanced` instead.
	NetworkAliases interface{}
	// Network mode of the container.
	NetworkMode interface{}
	// Id of the networks in which the
	// container is. *Deprecated:* use `networks_advanced` instead.
	Networks interface{}
	// See Networks Advanced below for details. If this block has priority to the deprecated `network_alias` and `network` properties.
	Networks_advanced interface{}
	// The PID (Process) Namespace mode for the container. Either `container:<name|id>` or `host`.
	PidMode interface{}
	// See Ports below for details.
	Ports interface{}
	// Run container in privileged mode.
	Privileged interface{}
	// Publish all ports of the container.
	PublishAllPorts interface{}
	// The restart policy for the container. Must be
	// one of "no", "on-failure", "always", "unless-stopped".
	Restart interface{}
	// If true, then the container will be automatically removed after his execution. Terraform
	// won't check this container after creation.
	Rm interface{}
	// If true, then the Docker container will be
	// started after creation. If false, then the container is only created.
	Start interface{}
	// See Ulimits below for
	// details.
	Ulimits interface{}
	// See File Upload below for details.
	Uploads interface{}
	// User used for run the first process. Format is
	// `user` or `user:group` which user and group can be passed literraly or
	// by name.
	User interface{}
	// Sets the usernamespace mode for the container when usernamespace remapping option is enabled.
	UsernsMode interface{}
	// See Volumes below for details.
	Volumes interface{}
}
