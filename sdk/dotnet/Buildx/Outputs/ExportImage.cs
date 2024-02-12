// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Docker.Buildx.Outputs
{

    [OutputType]
    public sealed class ExportImage
    {
        public readonly ImmutableDictionary<string, string>? Annotations;
        /// <summary>
        /// The compression type to use.
        /// </summary>
        public readonly Pulumi.Docker.Buildx.CompressionType? Compression;
        /// <summary>
        /// Compression level from 0 to 22.
        /// </summary>
        public readonly int? CompressionLevel;
        public readonly string? DanglingNamePrefix;
        /// <summary>
        /// Forcefully apply compression.
        /// </summary>
        public readonly bool? ForceCompression;
        public readonly bool? Insecure;
        public readonly bool? NameCanonical;
        /// <summary>
        /// Specify images names to export. This is overridden if tags are already specified.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        /// <summary>
        /// Use OCI media types in exporter manifests.
        /// </summary>
        public readonly bool? OciMediaTypes;
        /// <summary>
        /// Push after creating the image.
        /// </summary>
        public readonly bool? Push;
        public readonly bool? PushByDigest;
        /// <summary>
        /// 
        /// Store resulting images to the worker's image store, and ensure all its
        /// blobs are in the content store. Ignored if the worker doesn't have
        /// image store (when using OCI workers, for example).
        /// </summary>
        public readonly bool? Store;
        public readonly bool? Unpack;

        [OutputConstructor]
        private ExportImage(
            ImmutableDictionary<string, string>? annotations,

            Pulumi.Docker.Buildx.CompressionType? compression,

            int? compressionLevel,

            string? danglingNamePrefix,

            bool? forceCompression,

            bool? insecure,

            bool? nameCanonical,

            ImmutableArray<string> names,

            bool? ociMediaTypes,

            bool? push,

            bool? pushByDigest,

            bool? store,

            bool? unpack)
        {
            Annotations = annotations;
            Compression = compression;
            CompressionLevel = compressionLevel;
            DanglingNamePrefix = danglingNamePrefix;
            ForceCompression = forceCompression;
            Insecure = insecure;
            NameCanonical = nameCanonical;
            Names = names;
            OciMediaTypes = ociMediaTypes;
            Push = push;
            PushByDigest = pushByDigest;
            Store = store;
            Unpack = unpack;
        }
    }
}
