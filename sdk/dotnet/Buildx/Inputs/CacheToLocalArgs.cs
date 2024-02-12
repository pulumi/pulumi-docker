// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Docker.Buildx.Inputs
{

    public sealed class CacheToLocalArgs : global::Pulumi.ResourceArgs
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

        [Input("dest", required: true)]
        public Input<string> Dest { get; set; } = null!;

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

        [Input("mode")]
        public Input<Pulumi.Docker.Buildx.CacheMode>? Mode { get; set; }

        public CacheToLocalArgs()
        {
            Compression = Pulumi.Docker.Buildx.CompressionType.Gzip;
            CompressionLevel = 0;
            ForceCompression = false;
            IgnoreError = false;
            Mode = Pulumi.Docker.Buildx.CacheMode.Min;
        }
        public static new CacheToLocalArgs Empty => new CacheToLocalArgs();
    }
}
