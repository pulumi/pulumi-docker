// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Docker.Inputs
{

    public sealed class ServiceTaskSpecContainerSpecGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("args")]
        private InputList<string>? _args;

        /// <summary>
        /// Arguments to the command
        /// </summary>
        public InputList<string> Args
        {
            get => _args ?? (_args = new InputList<string>());
            set => _args = value;
        }

        [Input("capAdds")]
        private InputList<string>? _capAdds;

        /// <summary>
        /// List of Linux capabilities to add to the container
        /// </summary>
        public InputList<string> CapAdds
        {
            get => _capAdds ?? (_capAdds = new InputList<string>());
            set => _capAdds = value;
        }

        [Input("capDrops")]
        private InputList<string>? _capDrops;

        /// <summary>
        /// List of Linux capabilities to drop from the container
        /// </summary>
        public InputList<string> CapDrops
        {
            get => _capDrops ?? (_capDrops = new InputList<string>());
            set => _capDrops = value;
        }

        [Input("commands")]
        private InputList<string>? _commands;

        /// <summary>
        /// The command/entrypoint to be run in the image. According to the [docker cli](https://github.com/docker/cli/blob/v20.10.7/cli/command/service/opts.go#L705) the override of the entrypoint is also passed to the `command` property and there is no `entrypoint` attribute in the `ContainerSpec` of the service.
        /// </summary>
        public InputList<string> Commands
        {
            get => _commands ?? (_commands = new InputList<string>());
            set => _commands = value;
        }

        [Input("configs")]
        private InputList<Inputs.ServiceTaskSpecContainerSpecConfigGetArgs>? _configs;

        /// <summary>
        /// References to zero or more configs that will be exposed to the service
        /// </summary>
        public InputList<Inputs.ServiceTaskSpecContainerSpecConfigGetArgs> Configs
        {
            get => _configs ?? (_configs = new InputList<Inputs.ServiceTaskSpecContainerSpecConfigGetArgs>());
            set => _configs = value;
        }

        /// <summary>
        /// The working directory for commands to run in
        /// </summary>
        [Input("dir")]
        public Input<string>? Dir { get; set; }

        /// <summary>
        /// Specification for DNS related configurations in resolver configuration file (`resolv.conf`)
        /// </summary>
        [Input("dnsConfig")]
        public Input<Inputs.ServiceTaskSpecContainerSpecDnsConfigGetArgs>? DnsConfig { get; set; }

        [Input("env")]
        private InputMap<string>? _env;

        /// <summary>
        /// A list of environment variables in the form VAR="value"
        /// </summary>
        public InputMap<string> Env
        {
            get => _env ?? (_env = new InputMap<string>());
            set => _env = value;
        }

        [Input("groups")]
        private InputList<string>? _groups;

        /// <summary>
        /// A list of additional groups that the container process will run as
        /// </summary>
        public InputList<string> Groups
        {
            get => _groups ?? (_groups = new InputList<string>());
            set => _groups = value;
        }

        /// <summary>
        /// A test to perform to check that the container is healthy
        /// </summary>
        [Input("healthcheck")]
        public Input<Inputs.ServiceTaskSpecContainerSpecHealthcheckGetArgs>? Healthcheck { get; set; }

        /// <summary>
        /// The hostname to use for the container, as a valid RFC 1123 hostname
        /// </summary>
        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        [Input("hosts")]
        private InputList<Inputs.ServiceTaskSpecContainerSpecHostGetArgs>? _hosts;

        /// <summary>
        /// A list of hostname/IP mappings to add to the container's hosts file
        /// </summary>
        public InputList<Inputs.ServiceTaskSpecContainerSpecHostGetArgs> Hosts
        {
            get => _hosts ?? (_hosts = new InputList<Inputs.ServiceTaskSpecContainerSpecHostGetArgs>());
            set => _hosts = value;
        }

        /// <summary>
        /// The image name to use for the containers of the service, like `nginx:1.17.6`. Also use the data-source or resource of `docker.RemoteImage` with the `repo_digest` or `docker.RegistryImage` with the `name` attribute for this, as shown in the examples.
        /// </summary>
        [Input("image", required: true)]
        public Input<string> Image { get; set; } = null!;

        /// <summary>
        /// Isolation technology of the containers running the service. (Windows only). Defaults to `default`.
        /// </summary>
        [Input("isolation")]
        public Input<string>? Isolation { get; set; }

        [Input("labels")]
        private InputList<Inputs.ServiceTaskSpecContainerSpecLabelGetArgs>? _labels;

        /// <summary>
        /// User-defined key/value metadata
        /// </summary>
        public InputList<Inputs.ServiceTaskSpecContainerSpecLabelGetArgs> Labels
        {
            get => _labels ?? (_labels = new InputList<Inputs.ServiceTaskSpecContainerSpecLabelGetArgs>());
            set => _labels = value;
        }

        [Input("mounts")]
        private InputList<Inputs.ServiceTaskSpecContainerSpecMountGetArgs>? _mounts;

        /// <summary>
        /// Specification for mounts to be added to containers created as part of the service
        /// </summary>
        public InputList<Inputs.ServiceTaskSpecContainerSpecMountGetArgs> Mounts
        {
            get => _mounts ?? (_mounts = new InputList<Inputs.ServiceTaskSpecContainerSpecMountGetArgs>());
            set => _mounts = value;
        }

        /// <summary>
        /// Security options for the container
        /// </summary>
        [Input("privileges")]
        public Input<Inputs.ServiceTaskSpecContainerSpecPrivilegesGetArgs>? Privileges { get; set; }

        /// <summary>
        /// Mount the container's root filesystem as read only
        /// </summary>
        [Input("readOnly")]
        public Input<bool>? ReadOnly { get; set; }

        [Input("secrets")]
        private InputList<Inputs.ServiceTaskSpecContainerSpecSecretGetArgs>? _secrets;

        /// <summary>
        /// References to zero or more secrets that will be exposed to the service
        /// </summary>
        public InputList<Inputs.ServiceTaskSpecContainerSpecSecretGetArgs> Secrets
        {
            get => _secrets ?? (_secrets = new InputList<Inputs.ServiceTaskSpecContainerSpecSecretGetArgs>());
            set => _secrets = value;
        }

        /// <summary>
        /// Amount of time to wait for the container to terminate before forcefully removing it (ms|s|m|h). If not specified or '0s' the destroy will not check if all tasks/containers of the service terminate.
        /// </summary>
        [Input("stopGracePeriod")]
        public Input<string>? StopGracePeriod { get; set; }

        /// <summary>
        /// Signal to stop the container
        /// </summary>
        [Input("stopSignal")]
        public Input<string>? StopSignal { get; set; }

        [Input("sysctl")]
        private InputMap<string>? _sysctl;

        /// <summary>
        /// Sysctls config (Linux only)
        /// </summary>
        public InputMap<string> Sysctl
        {
            get => _sysctl ?? (_sysctl = new InputMap<string>());
            set => _sysctl = value;
        }

        /// <summary>
        /// The user inside the container
        /// </summary>
        [Input("user")]
        public Input<string>? User { get; set; }

        public ServiceTaskSpecContainerSpecGetArgs()
        {
        }
        public static new ServiceTaskSpecContainerSpecGetArgs Empty => new ServiceTaskSpecContainerSpecGetArgs();
    }
}
