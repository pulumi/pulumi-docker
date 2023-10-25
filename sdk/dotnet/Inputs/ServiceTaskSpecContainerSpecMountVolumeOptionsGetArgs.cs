// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Docker.Inputs
{

    public sealed class ServiceTaskSpecContainerSpecMountVolumeOptionsGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the driver to use to create the volume
        /// </summary>
        [Input("driverName")]
        public Input<string>? DriverName { get; set; }

        [Input("driverOptions")]
        private InputMap<string>? _driverOptions;

        /// <summary>
        /// key/value map of driver specific options
        /// </summary>
        public InputMap<string> DriverOptions
        {
            get => _driverOptions ?? (_driverOptions = new InputMap<string>());
            set => _driverOptions = value;
        }

        [Input("labels")]
        private InputList<Inputs.ServiceTaskSpecContainerSpecMountVolumeOptionsLabelGetArgs>? _labels;

        /// <summary>
        /// User-defined key/value metadata
        /// </summary>
        public InputList<Inputs.ServiceTaskSpecContainerSpecMountVolumeOptionsLabelGetArgs> Labels
        {
            get => _labels ?? (_labels = new InputList<Inputs.ServiceTaskSpecContainerSpecMountVolumeOptionsLabelGetArgs>());
            set => _labels = value;
        }

        /// <summary>
        /// Populate volume with data from the target
        /// </summary>
        [Input("noCopy")]
        public Input<bool>? NoCopy { get; set; }

        public ServiceTaskSpecContainerSpecMountVolumeOptionsGetArgs()
        {
        }
        public static new ServiceTaskSpecContainerSpecMountVolumeOptionsGetArgs Empty => new ServiceTaskSpecContainerSpecMountVolumeOptionsGetArgs();
    }
}
