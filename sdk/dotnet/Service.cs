// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Docker
{
    /// <summary>
    /// ## Import
    /// 
    /// Docker service can be imported using the long id, e.g. for a service with the short id `55ba873dd`
    /// 
    /// ```sh
    ///  $ pulumi import docker:index/service:Service foo $(docker service inspect -f {{.ID}} 55b)
    /// ```
    /// </summary>
    public partial class Service : Pulumi.CustomResource
    {
        /// <summary>
        /// See Auth below for details.
        /// </summary>
        [Output("auth")]
        public Output<Outputs.ServiceAuth?> Auth { get; private set; } = null!;

        /// <summary>
        /// See Converge Config below for details.
        /// </summary>
        [Output("convergeConfig")]
        public Output<Outputs.ServiceConvergeConfig?> ConvergeConfig { get; private set; } = null!;

        /// <summary>
        /// See EndpointSpec below for details.
        /// </summary>
        [Output("endpointSpec")]
        public Output<Outputs.ServiceEndpointSpec> EndpointSpec { get; private set; } = null!;

        /// <summary>
        /// See Labels below for details.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableArray<Outputs.ServiceLabel>> Labels { get; private set; } = null!;

        /// <summary>
        /// See Mode below for details.
        /// </summary>
        [Output("mode")]
        public Output<Outputs.ServiceMode> Mode { get; private set; } = null!;

        /// <summary>
        /// The name of the Docker service.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// See RollbackConfig below for details.
        /// </summary>
        [Output("rollbackConfig")]
        public Output<Outputs.ServiceRollbackConfig?> RollbackConfig { get; private set; } = null!;

        /// <summary>
        /// See TaskSpec below for details.
        /// </summary>
        [Output("taskSpec")]
        public Output<Outputs.ServiceTaskSpec> TaskSpec { get; private set; } = null!;

        /// <summary>
        /// See UpdateConfig below for details.
        /// </summary>
        [Output("updateConfig")]
        public Output<Outputs.ServiceUpdateConfig?> UpdateConfig { get; private set; } = null!;


        /// <summary>
        /// Create a Service resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Service(string name, ServiceArgs args, CustomResourceOptions? options = null)
            : base("docker:index/service:Service", name, args ?? new ServiceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Service(string name, Input<string> id, ServiceState? state = null, CustomResourceOptions? options = null)
            : base("docker:index/service:Service", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Service resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Service Get(string name, Input<string> id, ServiceState? state = null, CustomResourceOptions? options = null)
        {
            return new Service(name, id, state, options);
        }
    }

    public sealed class ServiceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// See Auth below for details.
        /// </summary>
        [Input("auth")]
        public Input<Inputs.ServiceAuthArgs>? Auth { get; set; }

        /// <summary>
        /// See Converge Config below for details.
        /// </summary>
        [Input("convergeConfig")]
        public Input<Inputs.ServiceConvergeConfigArgs>? ConvergeConfig { get; set; }

        /// <summary>
        /// See EndpointSpec below for details.
        /// </summary>
        [Input("endpointSpec")]
        public Input<Inputs.ServiceEndpointSpecArgs>? EndpointSpec { get; set; }

        [Input("labels")]
        private InputList<Inputs.ServiceLabelArgs>? _labels;

        /// <summary>
        /// See Labels below for details.
        /// </summary>
        public InputList<Inputs.ServiceLabelArgs> Labels
        {
            get => _labels ?? (_labels = new InputList<Inputs.ServiceLabelArgs>());
            set => _labels = value;
        }

        /// <summary>
        /// See Mode below for details.
        /// </summary>
        [Input("mode")]
        public Input<Inputs.ServiceModeArgs>? Mode { get; set; }

        /// <summary>
        /// The name of the Docker service.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// See RollbackConfig below for details.
        /// </summary>
        [Input("rollbackConfig")]
        public Input<Inputs.ServiceRollbackConfigArgs>? RollbackConfig { get; set; }

        /// <summary>
        /// See TaskSpec below for details.
        /// </summary>
        [Input("taskSpec", required: true)]
        public Input<Inputs.ServiceTaskSpecArgs> TaskSpec { get; set; } = null!;

        /// <summary>
        /// See UpdateConfig below for details.
        /// </summary>
        [Input("updateConfig")]
        public Input<Inputs.ServiceUpdateConfigArgs>? UpdateConfig { get; set; }

        public ServiceArgs()
        {
        }
    }

    public sealed class ServiceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// See Auth below for details.
        /// </summary>
        [Input("auth")]
        public Input<Inputs.ServiceAuthGetArgs>? Auth { get; set; }

        /// <summary>
        /// See Converge Config below for details.
        /// </summary>
        [Input("convergeConfig")]
        public Input<Inputs.ServiceConvergeConfigGetArgs>? ConvergeConfig { get; set; }

        /// <summary>
        /// See EndpointSpec below for details.
        /// </summary>
        [Input("endpointSpec")]
        public Input<Inputs.ServiceEndpointSpecGetArgs>? EndpointSpec { get; set; }

        [Input("labels")]
        private InputList<Inputs.ServiceLabelGetArgs>? _labels;

        /// <summary>
        /// See Labels below for details.
        /// </summary>
        public InputList<Inputs.ServiceLabelGetArgs> Labels
        {
            get => _labels ?? (_labels = new InputList<Inputs.ServiceLabelGetArgs>());
            set => _labels = value;
        }

        /// <summary>
        /// See Mode below for details.
        /// </summary>
        [Input("mode")]
        public Input<Inputs.ServiceModeGetArgs>? Mode { get; set; }

        /// <summary>
        /// The name of the Docker service.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// See RollbackConfig below for details.
        /// </summary>
        [Input("rollbackConfig")]
        public Input<Inputs.ServiceRollbackConfigGetArgs>? RollbackConfig { get; set; }

        /// <summary>
        /// See TaskSpec below for details.
        /// </summary>
        [Input("taskSpec")]
        public Input<Inputs.ServiceTaskSpecGetArgs>? TaskSpec { get; set; }

        /// <summary>
        /// See UpdateConfig below for details.
        /// </summary>
        [Input("updateConfig")]
        public Input<Inputs.ServiceUpdateConfigGetArgs>? UpdateConfig { get; set; }

        public ServiceState()
        {
        }
    }
}
