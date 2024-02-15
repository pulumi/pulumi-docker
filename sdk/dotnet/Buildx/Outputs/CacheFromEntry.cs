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
    public sealed class CacheFromEntry
    {
        /// <summary>
        /// Upload build caches to Azure's blob storage service.
        /// </summary>
        public readonly Outputs.CacheFromAzureBlob? Azblob;
        /// <summary>
        /// When `true` this entry will be excluded. Defaults to `false`.
        /// </summary>
        public readonly bool? Disabled;
        /// <summary>
        /// Recommended for use with GitHub Actions workflows.
        /// 
        /// An action like `crazy-max/ghaction-github-runtime` is recommended to
        /// expose appropriate credentials to your GitHub workflow.
        /// </summary>
        public readonly Outputs.CacheFromGitHubActions? Gha;
        /// <summary>
        /// A simple backend which caches images on your local filesystem.
        /// </summary>
        public readonly Outputs.CacheFromLocal? Local;
        /// <summary>
        /// A raw string as you would provide it to the Docker CLI (e.g.,
        /// `type=inline`).
        /// </summary>
        public readonly string? Raw;
        /// <summary>
        /// Upload build caches to remote registries.
        /// </summary>
        public readonly Outputs.CacheFromRegistry? Registry;
        /// <summary>
        /// Upload build caches to AWS S3 or an S3-compatible services such as
        /// MinIO.
        /// </summary>
        public readonly Outputs.CacheFromS3? S3;

        [OutputConstructor]
        private CacheFromEntry(
            Outputs.CacheFromAzureBlob? azblob,

            bool? disabled,

            Outputs.CacheFromGitHubActions? gha,

            Outputs.CacheFromLocal? local,

            string? raw,

            Outputs.CacheFromRegistry? registry,

            Outputs.CacheFromS3? s3)
        {
            Azblob = azblob;
            Disabled = disabled;
            Gha = gha;
            Local = local;
            Raw = raw;
            Registry = registry;
            S3 = s3;
        }
    }
}
