// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Docker
{
    /// <summary>
    /// ## Import
    /// 
    /// #!/bin/bash
    /// 
    /// Docker secret cannot be imported as the secret data, once set, is never exposed again.
    /// </summary>
    [DockerResourceType("docker:index/secret:Secret")]
    public partial class Secret : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Base64-url-safe-encoded secret data
        /// </summary>
        [Output("data")]
        public Output<string> Data { get; private set; } = null!;

        /// <summary>
        /// User-defined key/value metadata
        /// </summary>
        [Output("labels")]
        public Output<ImmutableArray<Outputs.SecretLabel>> Labels { get; private set; } = null!;

        /// <summary>
        /// User-defined name of the secret
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a Secret resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Secret(string name, SecretArgs args, CustomResourceOptions? options = null)
            : base("docker:index/secret:Secret", name, args ?? new SecretArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Secret(string name, Input<string> id, SecretState? state = null, CustomResourceOptions? options = null)
            : base("docker:index/secret:Secret", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "data",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Secret resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Secret Get(string name, Input<string> id, SecretState? state = null, CustomResourceOptions? options = null)
        {
            return new Secret(name, id, state, options);
        }
    }

    public sealed class SecretArgs : global::Pulumi.ResourceArgs
    {
        [Input("data", required: true)]
        private Input<string>? _data;

        /// <summary>
        /// Base64-url-safe-encoded secret data
        /// </summary>
        public Input<string>? Data
        {
            get => _data;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _data = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("labels")]
        private InputList<Inputs.SecretLabelArgs>? _labels;

        /// <summary>
        /// User-defined key/value metadata
        /// </summary>
        public InputList<Inputs.SecretLabelArgs> Labels
        {
            get => _labels ?? (_labels = new InputList<Inputs.SecretLabelArgs>());
            set => _labels = value;
        }

        /// <summary>
        /// User-defined name of the secret
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public SecretArgs()
        {
        }
        public static new SecretArgs Empty => new SecretArgs();
    }

    public sealed class SecretState : global::Pulumi.ResourceArgs
    {
        [Input("data")]
        private Input<string>? _data;

        /// <summary>
        /// Base64-url-safe-encoded secret data
        /// </summary>
        public Input<string>? Data
        {
            get => _data;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _data = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("labels")]
        private InputList<Inputs.SecretLabelGetArgs>? _labels;

        /// <summary>
        /// User-defined key/value metadata
        /// </summary>
        public InputList<Inputs.SecretLabelGetArgs> Labels
        {
            get => _labels ?? (_labels = new InputList<Inputs.SecretLabelGetArgs>());
            set => _labels = value;
        }

        /// <summary>
        /// User-defined name of the secret
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public SecretState()
        {
        }
        public static new SecretState Empty => new SecretState();
    }
}
