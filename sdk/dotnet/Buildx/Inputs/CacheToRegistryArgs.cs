// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Docker.Buildx.Inputs
{

    public sealed class CacheToRegistryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The compression type to use.
        /// </summary>
        [Input("compression")]
        public Input<Pulumi.Docker.Buildx.CompressionType>? Compression { get; set; }

        /// <summary>
        /// Compression level from 0 to 22.
        /// </summary>
        [Input("compressionLevel")]
        public Input<int>? CompressionLevel { get; set; }

        /// <summary>
        /// Forcefully apply compression.
        /// </summary>
        [Input("forceCompression")]
        public Input<bool>? ForceCompression { get; set; }

        /// <summary>
        /// Ignore errors caused by failed cache exports.
        /// </summary>
        [Input("ignoreError")]
        public Input<bool>? IgnoreError { get; set; }

        /// <summary>
        /// Export cache manifest as an OCI-compatible image manifest instead of a
        /// manifest list. Requires `ociMediaTypes` to also be `true`.
        /// 
        /// Some registries like AWS ECR will not work with caching if this is
        /// `false`.
        /// 
        /// Defaults to `false` to match Docker's default behavior.
        /// </summary>
        [Input("imageManifest")]
        public Input<bool>? ImageManifest { get; set; }

        /// <summary>
        /// The cache mode to use. Defaults to `min`.
        /// </summary>
        [Input("mode")]
        public Input<Pulumi.Docker.Buildx.CacheMode>? Mode { get; set; }

        /// <summary>
        /// Whether to use OCI media types in exported manifests. Defaults to
        /// `true`.
        /// </summary>
        [Input("ociMediaTypes")]
        public Input<bool>? OciMediaTypes { get; set; }

        /// <summary>
        /// Fully qualified name of the cache image to import.
        /// </summary>
        [Input("ref", required: true)]
        public Input<string> Ref { get; set; } = null!;

        public CacheToRegistryArgs()
        {
            Compression = Pulumi.Docker.Buildx.CompressionType.Gzip;
            CompressionLevel = 0;
            ForceCompression = false;
            IgnoreError = false;
            ImageManifest = false;
            Mode = Pulumi.Docker.Buildx.CacheMode.Min;
            OciMediaTypes = true;
        }
        public static new CacheToRegistryArgs Empty => new CacheToRegistryArgs();
    }
}
