// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Docker.Buildx.Inputs
{

    public sealed class ExportRegistryArgs : global::Pulumi.ResourceArgs
    {
        [Input("annotations")]
        private InputMap<string>? _annotations;
        public InputMap<string> Annotations
        {
            get => _annotations ?? (_annotations = new InputMap<string>());
            set => _annotations = value;
        }

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

        [Input("danglingNamePrefix")]
        public Input<string>? DanglingNamePrefix { get; set; }

        /// <summary>
        /// Forcefully apply compression.
        /// </summary>
        [Input("forceCompression")]
        public Input<bool>? ForceCompression { get; set; }

        [Input("insecure")]
        public Input<bool>? Insecure { get; set; }

        [Input("nameCanonical")]
        public Input<bool>? NameCanonical { get; set; }

        [Input("names")]
        private InputList<string>? _names;

        /// <summary>
        /// Specify images names to export. This is overridden if tags are already specified.
        /// </summary>
        public InputList<string> Names
        {
            get => _names ?? (_names = new InputList<string>());
            set => _names = value;
        }

        /// <summary>
        /// Use OCI media types in exporter manifests.
        /// </summary>
        [Input("ociMediaTypes")]
        public Input<bool>? OciMediaTypes { get; set; }

        /// <summary>
        /// Push after creating the image.
        /// </summary>
        [Input("push")]
        public Input<bool>? Push { get; set; }

        [Input("pushByDigest")]
        public Input<bool>? PushByDigest { get; set; }

        /// <summary>
        /// 
        /// Store resulting images to the worker's image store, and ensure all its
        /// blobs are in the content store. Ignored if the worker doesn't have
        /// image store (when using OCI workers, for example).
        /// </summary>
        [Input("store")]
        public Input<bool>? Store { get; set; }

        [Input("unpack")]
        public Input<bool>? Unpack { get; set; }

        public ExportRegistryArgs()
        {
            Compression = Pulumi.Docker.Buildx.CompressionType.Gzip;
            CompressionLevel = 0;
            ForceCompression = false;
            OciMediaTypes = false;
            Push = true;
            Store = true;
        }
        public static new ExportRegistryArgs Empty => new ExportRegistryArgs();
    }
}
