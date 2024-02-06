// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Docker.Outputs
{

    [OutputType]
    public sealed class ServiceTaskSpecContainerSpecSecret
    {
        /// <summary>
        /// Represents the file GID. Defaults to `0`
        /// </summary>
        public readonly string? FileGid;
        /// <summary>
        /// Represents represents the FileMode of the file. Defaults to `0o444`
        /// </summary>
        public readonly int? FileMode;
        /// <summary>
        /// Represents the final filename in the filesystem
        /// </summary>
        public readonly string FileName;
        /// <summary>
        /// Represents the file UID. Defaults to `0`
        /// </summary>
        public readonly string? FileUid;
        /// <summary>
        /// ID of the specific secret that we're referencing
        /// </summary>
        public readonly string SecretId;
        /// <summary>
        /// Name of the secret that this references, but this is just provided for lookup/display purposes. The config in the reference will be identified by its ID
        /// </summary>
        public readonly string? SecretName;

        [OutputConstructor]
        private ServiceTaskSpecContainerSpecSecret(
            string? fileGid,

            int? fileMode,

            string fileName,

            string? fileUid,

            string secretId,

            string? secretName)
        {
            FileGid = fileGid;
            FileMode = fileMode;
            FileName = fileName;
            FileUid = fileUid;
            SecretId = secretId;
            SecretName = secretName;
        }
    }
}
