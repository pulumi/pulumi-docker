// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Docker.Outputs
{

    [OutputType]
    public sealed class ServiceTaskSpecContainerSpec
    {
        /// <summary>
        /// Arguments to the command
        /// </summary>
        public readonly ImmutableArray<string> Args;
        /// <summary>
        /// The command/entrypoint to be run in the image. According to the [docker cli](https://github.com/docker/cli/blob/v20.10.7/cli/command/service/opts.go#L705) the override of the entrypoint is also passed to the `command` property and there is no `entrypoint` attribute in the `ContainerSpec` of the service.
        /// </summary>
        public readonly ImmutableArray<string> Commands;
        /// <summary>
        /// References to zero or more configs that will be exposed to the service
        /// </summary>
        public readonly ImmutableArray<Outputs.ServiceTaskSpecContainerSpecConfig> Configs;
        /// <summary>
        /// The working directory for commands to run in
        /// </summary>
        public readonly string? Dir;
        /// <summary>
        /// Specification for DNS related configurations in resolver configuration file (`resolv.conf`)
        /// </summary>
        public readonly Outputs.ServiceTaskSpecContainerSpecDnsConfig? DnsConfig;
        /// <summary>
        /// A list of environment variables in the form VAR="value"
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Env;
        /// <summary>
        /// A list of additional groups that the container process will run as
        /// </summary>
        public readonly ImmutableArray<string> Groups;
        /// <summary>
        /// A test to perform to check that the container is healthy
        /// </summary>
        public readonly Outputs.ServiceTaskSpecContainerSpecHealthcheck? Healthcheck;
        /// <summary>
        /// The hostname to use for the container, as a valid RFC 1123 hostname
        /// </summary>
        public readonly string? Hostname;
        /// <summary>
        /// A list of hostname/IP mappings to add to the container's hosts file
        /// </summary>
        public readonly ImmutableArray<Outputs.ServiceTaskSpecContainerSpecHost> Hosts;
        /// <summary>
        /// The image name to use for the containers of the service, like `nginx:1.17.6`. Also use the data-source or resource of `docker.RemoteImage` with the `repo_digest` or `docker.RegistryImage` with the `name` attribute for this, as shown in the examples.
        /// </summary>
        public readonly string Image;
        /// <summary>
        /// Isolation technology of the containers running the service. (Windows only). Defaults to `default`.
        /// </summary>
        public readonly string? Isolation;
        /// <summary>
        /// User-defined key/value metadata
        /// </summary>
        public readonly ImmutableArray<Outputs.ServiceTaskSpecContainerSpecLabel> Labels;
        /// <summary>
        /// Specification for mounts to be added to containers created as part of the service
        /// </summary>
        public readonly ImmutableArray<Outputs.ServiceTaskSpecContainerSpecMount> Mounts;
        /// <summary>
        /// Security options for the container
        /// </summary>
        public readonly Outputs.ServiceTaskSpecContainerSpecPrivileges? Privileges;
        /// <summary>
        /// Whether the mount should be read-only
        /// </summary>
        public readonly bool? ReadOnly;
        /// <summary>
        /// References to zero or more secrets that will be exposed to the service
        /// </summary>
        public readonly ImmutableArray<Outputs.ServiceTaskSpecContainerSpecSecret> Secrets;
        /// <summary>
        /// Amount of time to wait for the container to terminate before forcefully removing it (ms|s|m|h). If not specified or '0s' the destroy will not check if all tasks/containers of the service terminate.
        /// </summary>
        public readonly string? StopGracePeriod;
        /// <summary>
        /// Signal to stop the container
        /// </summary>
        public readonly string? StopSignal;
        /// <summary>
        /// Sysctls config (Linux only)
        /// </summary>
        public readonly ImmutableDictionary<string, object>? Sysctl;
        /// <summary>
        /// SELinux user label
        /// </summary>
        public readonly string? User;

        [OutputConstructor]
        private ServiceTaskSpecContainerSpec(
            ImmutableArray<string> args,

            ImmutableArray<string> commands,

            ImmutableArray<Outputs.ServiceTaskSpecContainerSpecConfig> configs,

            string? dir,

            Outputs.ServiceTaskSpecContainerSpecDnsConfig? dnsConfig,

            ImmutableDictionary<string, string>? env,

            ImmutableArray<string> groups,

            Outputs.ServiceTaskSpecContainerSpecHealthcheck? healthcheck,

            string? hostname,

            ImmutableArray<Outputs.ServiceTaskSpecContainerSpecHost> hosts,

            string image,

            string? isolation,

            ImmutableArray<Outputs.ServiceTaskSpecContainerSpecLabel> labels,

            ImmutableArray<Outputs.ServiceTaskSpecContainerSpecMount> mounts,

            Outputs.ServiceTaskSpecContainerSpecPrivileges? privileges,

            bool? readOnly,

            ImmutableArray<Outputs.ServiceTaskSpecContainerSpecSecret> secrets,

            string? stopGracePeriod,

            string? stopSignal,

            ImmutableDictionary<string, object>? sysctl,

            string? user)
        {
            Args = args;
            Commands = commands;
            Configs = configs;
            Dir = dir;
            DnsConfig = dnsConfig;
            Env = env;
            Groups = groups;
            Healthcheck = healthcheck;
            Hostname = hostname;
            Hosts = hosts;
            Image = image;
            Isolation = isolation;
            Labels = labels;
            Mounts = mounts;
            Privileges = privileges;
            ReadOnly = readOnly;
            Secrets = secrets;
            StopGracePeriod = stopGracePeriod;
            StopSignal = stopSignal;
            Sysctl = sysctl;
            User = user;
        }
    }
}
