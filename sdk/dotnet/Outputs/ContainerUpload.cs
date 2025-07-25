// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Docker.Outputs
{

    [OutputType]
    public sealed class ContainerUpload
    {
        /// <summary>
        /// Literal string value to use as the object content, which will be uploaded as UTF-8-encoded text. Conflicts with `content_base64` &amp; `source`
        /// </summary>
        public readonly string? Content;
        /// <summary>
        /// Base64-encoded data that will be decoded and uploaded as raw bytes for the object content. This allows safely uploading non-UTF8 binary data, but is recommended only for larger binary content such as the result of the `base64encode` interpolation function. See here for the reason. Conflicts with `content` &amp; `source`
        /// </summary>
        public readonly string? ContentBase64;
        /// <summary>
        /// If `true`, the file will be uploaded with user executable permission. Defaults to `false`.
        /// </summary>
        public readonly bool? Executable;
        /// <summary>
        /// Path to the file in the container where is upload goes to
        /// </summary>
        public readonly string File;
        /// <summary>
        /// The permission mode for the file in the container. Has precedence over `executable`.
        /// </summary>
        public readonly string? Permissions;
        /// <summary>
        /// A filename that references a file which will be uploaded as the object content. This allows for large file uploads that do not get stored in state. Conflicts with `content` &amp; `content_base64`
        /// </summary>
        public readonly string? Source;
        /// <summary>
        /// If using `source`, this will force an update if the file content has updated but the filename has not.
        /// </summary>
        public readonly string? SourceHash;

        [OutputConstructor]
        private ContainerUpload(
            string? content,

            string? contentBase64,

            bool? executable,

            string file,

            string? permissions,

            string? source,

            string? sourceHash)
        {
            Content = content;
            ContentBase64 = contentBase64;
            Executable = executable;
            File = file;
            Permissions = permissions;
            Source = source;
            SourceHash = sourceHash;
        }
    }
}
