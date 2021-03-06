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
        /// Arguments to the command.
        /// </summary>
        public readonly ImmutableArray<string> Args;
        /// <summary>
        /// The command to be run in the image.
        /// </summary>
        public readonly ImmutableArray<string> Commands;
        /// <summary>
        /// See Configs below for details.
        /// </summary>
        public readonly ImmutableArray<Outputs.ServiceTaskSpecContainerSpecConfig> Configs;
        /// <summary>
        /// The working directory for commands to run in.
        /// </summary>
        public readonly string? Dir;
        /// <summary>
        /// See DNS Config below for details.
        /// </summary>
        public readonly Outputs.ServiceTaskSpecContainerSpecDnsConfig? DnsConfig;
        /// <summary>
        /// A list of environment variables in the form VAR=value.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Env;
        /// <summary>
        /// A list of additional groups that the container process will run as.
        /// </summary>
        public readonly ImmutableArray<string> Groups;
        /// <summary>
        /// See Healthcheck below for details.
        /// </summary>
        public readonly Outputs.ServiceTaskSpecContainerSpecHealthcheck? Healthcheck;
        /// <summary>
        /// The hostname to use for the container, as a valid RFC 1123 hostname.
        /// </summary>
        public readonly string? Hostname;
        public readonly ImmutableArray<Outputs.ServiceTaskSpecContainerSpecHost> Hosts;
        /// <summary>
        /// The image used to create the Docker service.
        /// </summary>
        public readonly string Image;
        /// <summary>
        /// Isolation technology of the containers running the service. (Windows only). Valid values are: `default|process|hyperv`
        /// </summary>
        public readonly string? Isolation;
        /// <summary>
        /// See Labels below for details.
        /// </summary>
        public readonly ImmutableArray<Outputs.ServiceTaskSpecContainerSpecLabel> Labels;
        /// <summary>
        /// See Mounts below for details.
        /// </summary>
        public readonly ImmutableArray<Outputs.ServiceTaskSpecContainerSpecMount> Mounts;
        /// <summary>
        /// See Privileges below for details.
        /// </summary>
        public readonly Outputs.ServiceTaskSpecContainerSpecPrivileges? Privileges;
        /// <summary>
        /// Mount the container's root filesystem as read only.
        /// </summary>
        public readonly bool? ReadOnly;
        /// <summary>
        /// See Secrets below for details.
        /// </summary>
        public readonly ImmutableArray<Outputs.ServiceTaskSpecContainerSpecSecret> Secrets;
        /// <summary>
        /// Amount of time to wait for the container to terminate before forcefully removing it `(ms|s|m|h)`.
        /// </summary>
        public readonly string? StopGracePeriod;
        /// <summary>
        /// Signal to stop the container.
        /// </summary>
        public readonly string? StopSignal;
        /// <summary>
        /// The user inside the container.
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
            User = user;
        }
    }
}
