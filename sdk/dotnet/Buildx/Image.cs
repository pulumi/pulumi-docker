// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Docker.Buildx
{
    /// <summary>
    /// A Docker image built using Buildkit
    /// </summary>
    [DockerResourceType("docker:buildx/image:Image")]
    public partial class Image : global::Pulumi.CustomResource
    {
        [Output("architecture")]
        public Output<string?> Architecture { get; private set; } = null!;

        /// <summary>
        /// 
        /// Contexts to use while building the image. If omitted, an empty context
        /// is used. If more than one value is specified, they should be of the
        /// form "name=value".
        /// </summary>
        [Output("context")]
        public Output<ImmutableArray<string>> Context { get; private set; } = null!;

        /// <summary>
        /// 
        /// Name and optionally a tag (format: "name:tag"). If outputting to a
        /// registry, the name should include the fully qualified registry address.
        /// </summary>
        [Output("exports")]
        public Output<ImmutableArray<string>> Exports { get; private set; } = null!;

        /// <summary>
        /// 
        /// Name of the Dockerfile to use (default: "$PATH/Dockerfile").
        /// </summary>
        [Output("file")]
        public Output<string?> File { get; private set; } = null!;

        [Output("os")]
        public Output<string?> Os { get; private set; } = null!;

        [Output("repoDigests")]
        public Output<ImmutableArray<string>> RepoDigests { get; private set; } = null!;

        [Output("repoTags")]
        public Output<ImmutableArray<string>> RepoTags { get; private set; } = null!;

        [Output("size")]
        public Output<int?> Size { get; private set; } = null!;

        /// <summary>
        /// 
        /// Name and optionally a tag (format: "name:tag"). If outputting to a
        /// registry, the name should include the fully qualified registry address.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Image resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Image(string name, ImageArgs args, CustomResourceOptions? options = null)
            : base("docker:buildx/image:Image", name, args ?? new ImageArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Image(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("docker:buildx/image:Image", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Image resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Image Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Image(name, id, options);
        }
    }

    public sealed class ImageArgs : global::Pulumi.ResourceArgs
    {
        [Input("context")]
        private InputList<string>? _context;

        /// <summary>
        /// 
        /// Contexts to use while building the image. If omitted, an empty context
        /// is used. If more than one value is specified, they should be of the
        /// form "name=value".
        /// </summary>
        public InputList<string> Context
        {
            get => _context ?? (_context = new InputList<string>());
            set => _context = value;
        }

        [Input("exports")]
        private InputList<string>? _exports;

        /// <summary>
        /// 
        /// Name and optionally a tag (format: "name:tag"). If outputting to a
        /// registry, the name should include the fully qualified registry address.
        /// </summary>
        public InputList<string> Exports
        {
            get => _exports ?? (_exports = new InputList<string>());
            set => _exports = value;
        }

        /// <summary>
        /// 
        /// Name of the Dockerfile to use (default: "$PATH/Dockerfile").
        /// </summary>
        [Input("file")]
        public Input<string>? File { get; set; }

        [Input("tags", required: true)]
        private InputList<string>? _tags;

        /// <summary>
        /// 
        /// Name and optionally a tag (format: "name:tag"). If outputting to a
        /// registry, the name should include the fully qualified registry address.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        public ImageArgs()
        {
            File = "Dockerfile";
        }
        public static new ImageArgs Empty => new ImageArgs();
    }
}
