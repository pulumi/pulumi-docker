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
    public sealed class RemoteImageBuildAuthConfig
    {
        /// <summary>
        /// the auth token
        /// </summary>
        public readonly string? Auth;
        /// <summary>
        /// the user emal
        /// </summary>
        public readonly string? Email;
        /// <summary>
        /// hostname of the registry
        /// </summary>
        public readonly string HostName;
        /// <summary>
        /// the identity token
        /// </summary>
        public readonly string? IdentityToken;
        /// <summary>
        /// the registry password
        /// </summary>
        public readonly string? Password;
        /// <summary>
        /// the registry token
        /// </summary>
        public readonly string? RegistryToken;
        /// <summary>
        /// the server address
        /// </summary>
        public readonly string? ServerAddress;
        /// <summary>
        /// the registry user name
        /// </summary>
        public readonly string? UserName;

        [OutputConstructor]
        private RemoteImageBuildAuthConfig(
            string? auth,

            string? email,

            string hostName,

            string? identityToken,

            string? password,

            string? registryToken,

            string? serverAddress,

            string? userName)
        {
            Auth = auth;
            Email = email;
            HostName = hostName;
            IdentityToken = identityToken;
            Password = password;
            RegistryToken = registryToken;
            ServerAddress = serverAddress;
            UserName = userName;
        }
    }
}
