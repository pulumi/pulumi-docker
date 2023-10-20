// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Docker.Inputs
{

    /// <summary>
    /// The Docker build context
    /// </summary>
    public sealed class DockerBuildArgs : global::Pulumi.ResourceArgs
    {
        [Input("args")]
        private InputMap<string>? _args;

        /// <summary>
        /// An optional map of named build-time argument variables to set during the Docker build. This flag allows you to pass build-time variablesthat can be accessed like environment variables inside the RUN instruction.
        /// </summary>
        public InputMap<string> Args
        {
            get => _args ?? (_args = new InputMap<string>());
            set => _args = value;
        }

        /// <summary>
        /// The version of the Docker builder.
        /// </summary>
        [Input("builderVersion")]
        public Input<Pulumi.Docker.BuilderVersion>? BuilderVersion { get; set; }

        /// <summary>
        /// A list of image names to use as build cache. Images provided must have a cache manifest. Must provide authentication to cache registry.
        /// </summary>
        [Input("cacheFrom")]
        public Input<Inputs.CacheFromArgs>? CacheFrom { get; set; }

        /// <summary>
        /// The path to the build context to use.
        /// </summary>
        [Input("context")]
        public Input<string>? Context { get; set; }

        /// <summary>
        /// The path to the Dockerfile to use.
        /// </summary>
        [Input("dockerfile")]
        public Input<string>? Dockerfile { get; set; }

        /// <summary>
        /// The architecture of the platform you want to build this image for, e.g. `linux/arm64`.
        /// </summary>
        [Input("platform")]
        public Input<string>? Platform { get; set; }

        /// <summary>
        /// The target of the Dockerfile to build
        /// </summary>
        [Input("target")]
        public Input<string>? Target { get; set; }

        public DockerBuildArgs()
        {
        }
        public static new DockerBuildArgs Empty => new DockerBuildArgs();
    }
}
