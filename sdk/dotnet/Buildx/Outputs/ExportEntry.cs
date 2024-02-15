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
    public sealed class ExportEntry
    {
        /// <summary>
        /// When `true` this entry will be excluded. Defaults to `false`.
        /// </summary>
        public readonly bool? Disabled;
        /// <summary>
        /// Export as a Docker image layout.
        /// </summary>
        public readonly Outputs.ExportDocker? Docker;
        /// <summary>
        /// Outputs the build result into a container image format.
        /// </summary>
        public readonly Outputs.ExportImage? Image;
        /// <summary>
        /// Export to a local directory as files and directories.
        /// </summary>
        public readonly Outputs.ExportLocal? Local;
        /// <summary>
        /// An output property populated for exporters that pushed image
        /// manifest(s) to a registry.
        /// </summary>
        public readonly ImmutableArray<Outputs.Manifest> Manifests;
        /// <summary>
        /// Identical to the Docker exporter but uses OCI media types by default.
        /// </summary>
        public readonly Outputs.ExportOCI? Oci;
        /// <summary>
        /// A raw string as you would provide it to the Docker CLI (e.g.,
        /// `type=docker`)
        /// </summary>
        public readonly string? Raw;
        /// <summary>
        /// Identical to the Image exporter, but pushes by default.
        /// </summary>
        public readonly Outputs.ExportRegistry? Registry;
        /// <summary>
        /// Export to a local directory as a tarball.
        /// </summary>
        public readonly Outputs.ExportTar? Tar;

        [OutputConstructor]
        private ExportEntry(
            bool? disabled,

            Outputs.ExportDocker? docker,

            Outputs.ExportImage? image,

            Outputs.ExportLocal? local,

            ImmutableArray<Outputs.Manifest> manifests,

            Outputs.ExportOCI? oci,

            string? raw,

            Outputs.ExportRegistry? registry,

            Outputs.ExportTar? tar)
        {
            Disabled = disabled;
            Docker = docker;
            Image = image;
            Local = local;
            Manifests = manifests;
            Oci = oci;
            Raw = raw;
            Registry = registry;
            Tar = tar;
        }
    }
}
