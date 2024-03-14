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
    /// An index (or manifest list) referencing one or more existing images.
    /// 
    /// Useful for crafting a multi-platform image from several
    /// platform-specific images.
    /// 
    /// This creates an OCI image index or a Docker manifest list depending on
    /// the media types of the source images.
    /// 
    /// ## Example Usage
    /// ### Multi-platform registry caching
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Docker = Pulumi.Docker;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var amd64 = new Docker.Buildx.Image("amd64", new()
    ///     {
    ///         CacheFrom = new[]
    ///         {
    ///             new Docker.Buildx.Inputs.CacheFromArgs
    ///             {
    ///                 Registry = new Docker.Buildx.Inputs.CacheFromRegistryArgs
    ///                 {
    ///                     Ref = "docker.io/pulumi/pulumi:cache-amd64",
    ///                 },
    ///             },
    ///         },
    ///         CacheTo = new[]
    ///         {
    ///             new Docker.Buildx.Inputs.CacheToArgs
    ///             {
    ///                 Registry = new Docker.Buildx.Inputs.CacheToRegistryArgs
    ///                 {
    ///                     Mode = Docker.Buildx.Image.CacheMode.Max,
    ///                     Ref = "docker.io/pulumi/pulumi:cache-amd64",
    ///                 },
    ///             },
    ///         },
    ///         Context = new Docker.Buildx.Inputs.BuildContextArgs
    ///         {
    ///             Location = "app",
    ///         },
    ///         Platforms = new[]
    ///         {
    ///             Docker.Buildx.Image.Platform.Linux_amd64,
    ///         },
    ///         Tags = new[]
    ///         {
    ///             "docker.io/pulumi/pulumi:3.107.0-amd64",
    ///         },
    ///     });
    /// 
    ///     var arm64 = new Docker.Buildx.Image("arm64", new()
    ///     {
    ///         CacheFrom = new[]
    ///         {
    ///             new Docker.Buildx.Inputs.CacheFromArgs
    ///             {
    ///                 Registry = new Docker.Buildx.Inputs.CacheFromRegistryArgs
    ///                 {
    ///                     Ref = "docker.io/pulumi/pulumi:cache-arm64",
    ///                 },
    ///             },
    ///         },
    ///         CacheTo = new[]
    ///         {
    ///             new Docker.Buildx.Inputs.CacheToArgs
    ///             {
    ///                 Registry = new Docker.Buildx.Inputs.CacheToRegistryArgs
    ///                 {
    ///                     Mode = Docker.Buildx.Image.CacheMode.Max,
    ///                     Ref = "docker.io/pulumi/pulumi:cache-arm64",
    ///                 },
    ///             },
    ///         },
    ///         Context = new Docker.Buildx.Inputs.BuildContextArgs
    ///         {
    ///             Location = "app",
    ///         },
    ///         Platforms = new[]
    ///         {
    ///             Docker.Buildx.Image.Platform.Linux_arm64,
    ///         },
    ///         Tags = new[]
    ///         {
    ///             "docker.io/pulumi/pulumi:3.107.0-arm64",
    ///         },
    ///     });
    /// 
    ///     var index = new Docker.Buildx.Index("index", new()
    ///     {
    ///         Sources = new[]
    ///         {
    ///             amd64.Ref,
    ///             arm64.Ref,
    ///         },
    ///         Tag = "docker.io/pulumi/pulumi:3.107.0",
    ///     });
    /// 
    ///     return new Dictionary&lt;string, object?&gt;
    ///     {
    ///         ["ref"] = index.Ref,
    ///     };
    /// });
    /// 
    /// ```
    /// </summary>
    [DockerResourceType("docker:buildx/image:Index")]
    public partial class Index : global::Pulumi.CustomResource
    {
        /// <summary>
        /// If true, push the index to the target registry.
        /// 
        /// Defaults to `true`.
        /// </summary>
        [Output("push")]
        public Output<bool?> Push { get; private set; } = null!;

        /// <summary>
        /// The pushed tag with digest.
        /// 
        /// Identical to the tag if the index was not pushed.
        /// </summary>
        [Output("ref")]
        public Output<string> Ref { get; private set; } = null!;

        /// <summary>
        /// Authentication for the registry where the tagged index will be pushed.
        /// 
        /// Credentials can also be included with the provider's configuration.
        /// </summary>
        [Output("registry")]
        public Output<Outputs.RegistryAuth?> Registry { get; private set; } = null!;

        /// <summary>
        /// Existing images to include in the index.
        /// </summary>
        [Output("sources")]
        public Output<ImmutableArray<string>> Sources { get; private set; } = null!;

        /// <summary>
        /// The tag to apply to the index.
        /// </summary>
        [Output("tag")]
        public Output<string> Tag { get; private set; } = null!;


        /// <summary>
        /// Create a Index resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Index(string name, IndexArgs args, CustomResourceOptions? options = null)
            : base("docker:buildx/image:Index", name, args ?? new IndexArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Index(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("docker:buildx/image:Index", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Index resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Index Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Index(name, id, options);
        }
    }

    public sealed class IndexArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If true, push the index to the target registry.
        /// 
        /// Defaults to `true`.
        /// </summary>
        [Input("push")]
        public Input<bool>? Push { get; set; }

        /// <summary>
        /// Authentication for the registry where the tagged index will be pushed.
        /// 
        /// Credentials can also be included with the provider's configuration.
        /// </summary>
        [Input("registry")]
        public Input<Inputs.RegistryAuthArgs>? Registry { get; set; }

        [Input("sources", required: true)]
        private InputList<string>? _sources;

        /// <summary>
        /// Existing images to include in the index.
        /// </summary>
        public InputList<string> Sources
        {
            get => _sources ?? (_sources = new InputList<string>());
            set => _sources = value;
        }

        /// <summary>
        /// The tag to apply to the index.
        /// </summary>
        [Input("tag", required: true)]
        public Input<string> Tag { get; set; } = null!;

        public IndexArgs()
        {
            Push = true;
        }
        public static new IndexArgs Empty => new IndexArgs();
    }
}
