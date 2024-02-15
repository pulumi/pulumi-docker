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
    public sealed class ExportRegistry
    {
        /// <summary>
        /// Attach an arbitrary key/value annotation to the image.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Annotations;
        /// <summary>
        /// The compression type to use.
        /// </summary>
        public readonly Pulumi.Docker.Buildx.CompressionType? Compression;
        /// <summary>
        /// Compression level from 0 to 22.
        /// </summary>
        public readonly int? CompressionLevel;
        /// <summary>
        /// Name image with `prefix@&lt;digest&gt;`, used for anonymous images.
        /// </summary>
        public readonly string? DanglingNamePrefix;
        /// <summary>
        /// Forcefully apply compression.
        /// </summary>
        public readonly bool? ForceCompression;
        /// <summary>
        /// Allow pushing to an insecure registry.
        /// </summary>
        public readonly bool? Insecure;
        /// <summary>
        /// Add additional canonical name (`name@&lt;digest&gt;`).
        /// </summary>
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
        /// <summary>
        /// Push image without name.
        /// </summary>
        public readonly bool? PushByDigest;
        /// <summary>
        /// Store resulting images to the worker's image store and ensure all of
        /// its blobs are in the content store.
        /// 
        /// Defaults to `true`.
        /// 
        /// Ignored if the worker doesn't have image store (when using OCI workers,
        /// for example).
        /// </summary>
        public readonly bool? Store;
        /// <summary>
        /// Unpack image after creation (for use with containerd). Defaults to
        /// `false`.
        /// </summary>
        public readonly bool? Unpack;

        [OutputConstructor]
        private ExportRegistry(
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
