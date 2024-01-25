// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Docker.Inputs
{

    public sealed class ServiceTaskSpecContainerSpecSecretGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("fileGid")]
        public Input<string>? FileGid { get; set; }

        [Input("fileMode")]
        public Input<int>? FileMode { get; set; }

        [Input("fileName", required: true)]
        public Input<string> FileName { get; set; } = null!;

        [Input("fileUid")]
        public Input<string>? FileUid { get; set; }

        [Input("secretId", required: true)]
        public Input<string> SecretId { get; set; } = null!;

        [Input("secretName")]
        public Input<string>? SecretName { get; set; }

        public ServiceTaskSpecContainerSpecSecretGetArgs()
        {
        }
        public static new ServiceTaskSpecContainerSpecSecretGetArgs Empty => new ServiceTaskSpecContainerSpecSecretGetArgs();
    }
}
