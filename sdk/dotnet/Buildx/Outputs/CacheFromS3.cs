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
    public sealed class CacheFromS3
    {
        /// <summary>
        /// Defaults to `$AWS_ACCESS_KEY_ID`.
        /// </summary>
        public readonly string? AccessKeyId;
        /// <summary>
        /// Prefix to prepend to blob filenames.
        /// </summary>
        public readonly string? BlobsPrefix;
        /// <summary>
        /// Name of the S3 bucket.
        /// </summary>
        public readonly string Bucket;
        /// <summary>
        /// Endpoint of the S3 bucket.
        /// </summary>
        public readonly string? EndpointUrl;
        /// <summary>
        /// Prefix to prepend on manifest filenames.
        /// </summary>
        public readonly string? ManifestsPrefix;
        /// <summary>
        /// Name of the cache image.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The geographic location of the bucket. Defaults to `$AWS_REGION`.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// Defaults to `$AWS_SECRET_ACCESS_KEY`.
        /// </summary>
        public readonly string? SecretAccessKey;
        /// <summary>
        /// Defaults to `$AWS_SESSION_TOKEN`.
        /// </summary>
        public readonly string? SessionToken;
        /// <summary>
        /// Uses `bucket` in the URL instead of hostname when `true`.
        /// </summary>
        public readonly bool? UsePathStyle;

        [OutputConstructor]
        private CacheFromS3(
            string? accessKeyId,

            string? blobsPrefix,

            string bucket,

            string? endpointUrl,

            string? manifestsPrefix,

            string? name,

            string region,

            string? secretAccessKey,

            string? sessionToken,

            bool? usePathStyle)
        {
            AccessKeyId = accessKeyId;
            BlobsPrefix = blobsPrefix;
            Bucket = bucket;
            EndpointUrl = endpointUrl;
            ManifestsPrefix = manifestsPrefix;
            Name = name;
            Region = region;
            SecretAccessKey = secretAccessKey;
            SessionToken = sessionToken;
            UsePathStyle = usePathStyle;
        }
    }
}
