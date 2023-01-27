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
        /// An optional map of named build-time argument variables to set during the Docker build. This flag allows you to pass built-time variablesthat can be accessed like environment variables inside the RUN instruction.
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
        /// A list of images to use as build cache
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

        [Input("extraOptions")]
        private InputList<string>? _extraOptions;

        /// <summary>
        /// A bag of extra options to pass on to the docker SDK.
        /// </summary>
        public InputList<string> ExtraOptions
        {
            get => _extraOptions ?? (_extraOptions = new InputList<string>());
            set => _extraOptions = value;
        }

        /// <summary>
        /// The target of the Dockerfile to build
        /// </summary>
        [Input("target")]
        public Input<string>? Target { get; set; }

        public DockerBuildArgs()
        {
            BuilderVersion = Pulumi.Docker.BuilderVersion.BuilderBuildKit;
            Context = ".";
            Dockerfile = "Dockerfile";
        }
        public static new DockerBuildArgs Empty => new DockerBuildArgs();
    }
}
