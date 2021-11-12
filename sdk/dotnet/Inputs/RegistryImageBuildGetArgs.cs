// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Docker.Inputs
{

    public sealed class RegistryImageBuildGetArgs : Pulumi.ResourceArgs
    {
        [Input("authConfigs")]
        private InputList<Inputs.RegistryImageBuildAuthConfigGetArgs>? _authConfigs;

        /// <summary>
        /// The configuration for the authentication
        /// </summary>
        public InputList<Inputs.RegistryImageBuildAuthConfigGetArgs> AuthConfigs
        {
            get => _authConfigs ?? (_authConfigs = new InputList<Inputs.RegistryImageBuildAuthConfigGetArgs>());
            set => _authConfigs = value;
        }

        [Input("buildArgs")]
        private InputMap<string>? _buildArgs;

        /// <summary>
        /// Pairs for build-time variables in the form TODO
        /// </summary>
        public InputMap<string> BuildArgs
        {
            get => _buildArgs ?? (_buildArgs = new InputMap<string>());
            set => _buildArgs = value;
        }

        /// <summary>
        /// BuildID is an optional identifier that can be passed together with the build request. The
        /// </summary>
        [Input("buildId")]
        public Input<string>? BuildId { get; set; }

        [Input("cacheFroms")]
        private InputList<string>? _cacheFroms;

        /// <summary>
        /// Images to consider as cache sources
        /// </summary>
        public InputList<string> CacheFroms
        {
            get => _cacheFroms ?? (_cacheFroms = new InputList<string>());
            set => _cacheFroms = value;
        }

        /// <summary>
        /// Optional parent cgroup for the container
        /// </summary>
        [Input("cgroupParent")]
        public Input<string>? CgroupParent { get; set; }

        /// <summary>
        /// The absolute path to the context folder. You can use the helper function '${path.cwd}/context-dir'.
        /// </summary>
        [Input("context", required: true)]
        public Input<string> Context { get; set; } = null!;

        /// <summary>
        /// The length of a CPU period in microseconds
        /// </summary>
        [Input("cpuPeriod")]
        public Input<int>? CpuPeriod { get; set; }

        /// <summary>
        /// Microseconds of CPU time that the container can get in a CPU period
        /// </summary>
        [Input("cpuQuota")]
        public Input<int>? CpuQuota { get; set; }

        /// <summary>
        /// CPUs in which to allow execution (e.g., `0-3`, `0`, `1`)
        /// </summary>
        [Input("cpuSetCpus")]
        public Input<string>? CpuSetCpus { get; set; }

        /// <summary>
        /// MEMs in which to allow execution (`0-3`, `0`, `1`)
        /// </summary>
        [Input("cpuSetMems")]
        public Input<string>? CpuSetMems { get; set; }

        /// <summary>
        /// CPU shares (relative weight)
        /// </summary>
        [Input("cpuShares")]
        public Input<int>? CpuShares { get; set; }

        /// <summary>
        /// Dockerfile file. Defaults to `Dockerfile`
        /// </summary>
        [Input("dockerfile")]
        public Input<string>? Dockerfile { get; set; }

        [Input("extraHosts")]
        private InputList<string>? _extraHosts;

        /// <summary>
        /// A list of hostnames/IP mappings to add to the container’s /etc/hosts file. Specified in the form ["hostname:IP"]
        /// </summary>
        public InputList<string> ExtraHosts
        {
            get => _extraHosts ?? (_extraHosts = new InputList<string>());
            set => _extraHosts = value;
        }

        /// <summary>
        /// Always remove intermediate containers
        /// </summary>
        [Input("forceRemove")]
        public Input<bool>? ForceRemove { get; set; }

        /// <summary>
        /// Isolation represents the isolation technology of a container. The supported values are
        /// </summary>
        [Input("isolation")]
        public Input<string>? Isolation { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// User-defined key/value metadata
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Set memory limit for build
        /// </summary>
        [Input("memory")]
        public Input<int>? Memory { get; set; }

        /// <summary>
        /// Total memory (memory + swap), -1 to enable unlimited swap
        /// </summary>
        [Input("memorySwap")]
        public Input<int>? MemorySwap { get; set; }

        /// <summary>
        /// Set the networking mode for the RUN instructions during build
        /// </summary>
        [Input("networkMode")]
        public Input<string>? NetworkMode { get; set; }

        /// <summary>
        /// Do not use the cache when building the image
        /// </summary>
        [Input("noCache")]
        public Input<bool>? NoCache { get; set; }

        /// <summary>
        /// Set platform if server is multi-platform capable
        /// </summary>
        [Input("platform")]
        public Input<string>? Platform { get; set; }

        /// <summary>
        /// Attempt to pull the image even if an older image exists locally
        /// </summary>
        [Input("pullParent")]
        public Input<bool>? PullParent { get; set; }

        /// <summary>
        /// A Git repository URI or HTTP/HTTPS context URI
        /// </summary>
        [Input("remoteContext")]
        public Input<string>? RemoteContext { get; set; }

        /// <summary>
        /// Remove intermediate containers after a successful build (default behavior)
        /// </summary>
        [Input("remove")]
        public Input<bool>? Remove { get; set; }

        [Input("securityOpts")]
        private InputList<string>? _securityOpts;

        /// <summary>
        /// The security options
        /// </summary>
        public InputList<string> SecurityOpts
        {
            get => _securityOpts ?? (_securityOpts = new InputList<string>());
            set => _securityOpts = value;
        }

        /// <summary>
        /// Set an ID for the build session
        /// </summary>
        [Input("sessionId")]
        public Input<string>? SessionId { get; set; }

        /// <summary>
        /// Size of /dev/shm in bytes. The size must be greater than 0
        /// </summary>
        [Input("shmSize")]
        public Input<int>? ShmSize { get; set; }

        /// <summary>
        /// If true the new layers are squashed into a new image with a single new layer
        /// </summary>
        [Input("squash")]
        public Input<bool>? Squash { get; set; }

        /// <summary>
        /// Suppress the build output and print image ID on success
        /// </summary>
        [Input("suppressOutput")]
        public Input<bool>? SuppressOutput { get; set; }

        /// <summary>
        /// Set the target build stage to build
        /// </summary>
        [Input("target")]
        public Input<string>? Target { get; set; }

        [Input("ulimits")]
        private InputList<Inputs.RegistryImageBuildUlimitGetArgs>? _ulimits;

        /// <summary>
        /// Configuration for ulimits
        /// </summary>
        public InputList<Inputs.RegistryImageBuildUlimitGetArgs> Ulimits
        {
            get => _ulimits ?? (_ulimits = new InputList<Inputs.RegistryImageBuildUlimitGetArgs>());
            set => _ulimits = value;
        }

        /// <summary>
        /// Version of the unerlying builder to use
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public RegistryImageBuildGetArgs()
        {
        }
    }
}
