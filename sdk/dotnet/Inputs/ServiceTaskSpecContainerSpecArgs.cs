// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Docker.Inputs
{

    public sealed class ServiceTaskSpecContainerSpecArgs : global::Pulumi.ResourceArgs
    {
        [Input("args")]
        private InputList<string>? _args;
        public InputList<string> Args
        {
            get => _args ?? (_args = new InputList<string>());
            set => _args = value;
        }

        [Input("commands")]
        private InputList<string>? _commands;
        public InputList<string> Commands
        {
            get => _commands ?? (_commands = new InputList<string>());
            set => _commands = value;
        }

        [Input("configs")]
        private InputList<Inputs.ServiceTaskSpecContainerSpecConfigArgs>? _configs;
        public InputList<Inputs.ServiceTaskSpecContainerSpecConfigArgs> Configs
        {
            get => _configs ?? (_configs = new InputList<Inputs.ServiceTaskSpecContainerSpecConfigArgs>());
            set => _configs = value;
        }

        [Input("dir")]
        public Input<string>? Dir { get; set; }

        [Input("dnsConfig")]
        public Input<Inputs.ServiceTaskSpecContainerSpecDnsConfigArgs>? DnsConfig { get; set; }

        [Input("env")]
        private InputMap<string>? _env;
        public InputMap<string> Env
        {
            get => _env ?? (_env = new InputMap<string>());
            set => _env = value;
        }

        [Input("groups")]
        private InputList<string>? _groups;
        public InputList<string> Groups
        {
            get => _groups ?? (_groups = new InputList<string>());
            set => _groups = value;
        }

        [Input("healthcheck")]
        public Input<Inputs.ServiceTaskSpecContainerSpecHealthcheckArgs>? Healthcheck { get; set; }

        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        [Input("hosts")]
        private InputList<Inputs.ServiceTaskSpecContainerSpecHostArgs>? _hosts;
        public InputList<Inputs.ServiceTaskSpecContainerSpecHostArgs> Hosts
        {
            get => _hosts ?? (_hosts = new InputList<Inputs.ServiceTaskSpecContainerSpecHostArgs>());
            set => _hosts = value;
        }

        [Input("image", required: true)]
        public Input<string> Image { get; set; } = null!;

        [Input("isolation")]
        public Input<string>? Isolation { get; set; }

        [Input("labels")]
        private InputList<Inputs.ServiceTaskSpecContainerSpecLabelArgs>? _labels;
        public InputList<Inputs.ServiceTaskSpecContainerSpecLabelArgs> Labels
        {
            get => _labels ?? (_labels = new InputList<Inputs.ServiceTaskSpecContainerSpecLabelArgs>());
            set => _labels = value;
        }

        [Input("mounts")]
        private InputList<Inputs.ServiceTaskSpecContainerSpecMountArgs>? _mounts;
        public InputList<Inputs.ServiceTaskSpecContainerSpecMountArgs> Mounts
        {
            get => _mounts ?? (_mounts = new InputList<Inputs.ServiceTaskSpecContainerSpecMountArgs>());
            set => _mounts = value;
        }

        [Input("privileges")]
        public Input<Inputs.ServiceTaskSpecContainerSpecPrivilegesArgs>? Privileges { get; set; }

        [Input("readOnly")]
        public Input<bool>? ReadOnly { get; set; }

        [Input("secrets")]
        private InputList<Inputs.ServiceTaskSpecContainerSpecSecretArgs>? _secrets;
        public InputList<Inputs.ServiceTaskSpecContainerSpecSecretArgs> Secrets
        {
            get => _secrets ?? (_secrets = new InputList<Inputs.ServiceTaskSpecContainerSpecSecretArgs>());
            set => _secrets = value;
        }

        [Input("stopGracePeriod")]
        public Input<string>? StopGracePeriod { get; set; }

        [Input("stopSignal")]
        public Input<string>? StopSignal { get; set; }

        [Input("sysctl")]
        private InputMap<object>? _sysctl;
        public InputMap<object> Sysctl
        {
            get => _sysctl ?? (_sysctl = new InputMap<object>());
            set => _sysctl = value;
        }

        [Input("user")]
        public Input<string>? User { get; set; }

        public ServiceTaskSpecContainerSpecArgs()
        {
        }
        public static new ServiceTaskSpecContainerSpecArgs Empty => new ServiceTaskSpecContainerSpecArgs();
    }
}
