// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Docker.Inputs
{

    /// <summary>
    /// Contains a list of images to reference when building using a cache
    /// </summary>
    public sealed class CacheFromArgs : global::Pulumi.ResourceArgs
    {
        [Input("images")]
        private InputList<string>? _images;

        /// <summary>
        /// Specifies cached images
        /// </summary>
        public InputList<string> Images
        {
            get => _images ?? (_images = new InputList<string>());
            set => _images = value;
        }

        public CacheFromArgs()
        {
        }
        public static new CacheFromArgs Empty => new CacheFromArgs();
    }
}
