// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Docker.Inputs
{

    public sealed class ServiceTaskSpecPlacementArgs : global::Pulumi.ResourceArgs
    {
        [Input("constraints")]
        private InputList<string>? _constraints;

        /// <summary>
        /// An array of constraints. e.g.: `node.role==manager`
        /// </summary>
        public InputList<string> Constraints
        {
            get => _constraints ?? (_constraints = new InputList<string>());
            set => _constraints = value;
        }

        /// <summary>
        /// Maximum number of replicas for per node (default value is `0`, which is unlimited)
        /// </summary>
        [Input("maxReplicas")]
        public Input<int>? MaxReplicas { get; set; }

        [Input("platforms")]
        private InputList<Inputs.ServiceTaskSpecPlacementPlatformArgs>? _platforms;

        /// <summary>
        /// Platforms stores all the platforms that the service's image can run on
        /// </summary>
        public InputList<Inputs.ServiceTaskSpecPlacementPlatformArgs> Platforms
        {
            get => _platforms ?? (_platforms = new InputList<Inputs.ServiceTaskSpecPlacementPlatformArgs>());
            set => _platforms = value;
        }

        [Input("prefs")]
        private InputList<string>? _prefs;

        /// <summary>
        /// Preferences provide a way to make the scheduler aware of factors such as topology. They are provided in order from highest to lowest precedence, e.g.: `spread=node.role.manager`
        /// </summary>
        public InputList<string> Prefs
        {
            get => _prefs ?? (_prefs = new InputList<string>());
            set => _prefs = value;
        }

        public ServiceTaskSpecPlacementArgs()
        {
        }
        public static new ServiceTaskSpecPlacementArgs Empty => new ServiceTaskSpecPlacementArgs();
    }
}
