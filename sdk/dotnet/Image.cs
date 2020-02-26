// Copyright 2016-2019, Pulumi Corporation.

using Pulumi;
using Pulumi.Serialization;
using System;
using System.Collections.Generic;

namespace Pulumi.Docker
{
    /// <summary>
    /// Arguments for constructing an Image resource.
    /// </summary>
    public class ImageArgs : ResourceArgs
    {
        /// <summary>
        /// The qualified image name that will be pushed to the remote registry.  Must be a supported
        /// image name for the target registry user.  This name can include a tag at the end.  If
        /// provided all pushed image resources will contain that tag as well.
        /// <para />
        /// Either <see cref="ImageName"/> or <see cref="LocalImageName"/> can have a tag.  However, 
        /// if both have a tag, then those tags must match.
        /// </summary>
        [Input("imageName", required: true)]
        public Input<string> ImageName { get; set; } = null!;

        /// <summary>
        /// The Docker build context, as a folder path or a detailed DockerBuild object.
        /// </summary>
        [Input("build", required: true)]
        public InputUnion<string, DockerBuild> Build { get; set; } = null!;

        /// <summary>
        /// The docker image name to build locally before tagging with image name.  If not provided, it
        /// will be given the value of to <see cref="ImageName"/>.  This name can include a tag at the end.
        /// If provided all pushed image resources will contain that tag as well.
        /// <para />
        /// Either <see cref="ImageName"/> or <see cref="LocalImageName"/> can have a tag. However,
        /// if both have a tag, then those tags must match.
        /// </summary>
        [Input("localImageName")]
        public Input<string>? LocalImageName { get; set; }

        /// <summary>
        /// Credentials for the docker registry to push to.
        /// </summary>
        [Input("registry")]
        public Input<ImageRegistry>? Registry { get; set; }

        /// <summary>
        /// Skip push flag.
        /// </summary>
        [Input("skipPush")]
        public Input<bool>? SkipPush { get; set; }
    }

    public class ImageRegistry : ResourceArgs
    {
        /// <summary>
        /// Docker registry server URL to push to.  Some common values include:
        /// <list type="bullet">
        /// <item>DockerHub: docker.io or https://index.docker.io/v1</item>
        /// <item>Azure Container Registry: (name).azurecr.io</item>
        /// <item>AWS Elastic Container Registry: (account).dkr.ecr.us-east-2.amazonaws.com</item>
        /// <item>Google Container Registry: (name).gcr.io</item>
        /// </list>
        /// </summary>
        [Input("server")]
        public Input<string> Server { get; set; } = null!;

        /// <summary>
        /// Username for login to the target Docker registry.
        /// </summary>
        [Input("username")]
        public Input<string> Username { get; set; } = null!;

        /// <summary>
        /// Password for login to the target Docker registry.
        /// </summary>
        [Input("password")]
        public Input<string> Password { get; set; } = null!;
    }

    /// <summary>
    /// A copy of <see cref="ImageRegistry"/> but without using <see cref="Input{T}"/> in types.
    /// </summary>
    internal class ImageRegistryUnwrap
    {
        public string Server { get; set; } = null!;

        public string Username { get; set; } = null!;

        public string Password { get; set; } = null!;
    }

    /// <summary>
    /// A <see cref="Image"/> resource represents a Docker image built locally which is published and made
    /// available via a remote Docker registry.  This can be used to ensure that a Docker source
    /// directory from a local deployment environment is built and pushed to a cloud-hosted Docker
    /// registry as part of a Pulumi deployment, so that it can be referenced as an image input from
    /// other cloud services that reference Docker images - including Kubernetes Pods, AWS ECS Tasks, and
    /// Azure Container Instances.
    /// </summary>
    public class Image : ComponentResource
    {
        /// <summary>
        /// The base image name that was built and pushed.  This does not include the id annotation, so
        /// is not pinned to the specific build performed by this <see cref="Image"/>.
        /// </summary>
        [Output("baseImageName")]
        public Output<string> BaseImageName { get; private set; }

        /// <summary>
        /// The unique pinned image name on the remote repository.
        /// </summary>
        [Output("imageName")]
        public Output<string> ImageName { get; private set; }

        /// <summary>
        /// The server the image is located at.
        /// </summary>
        [Output("registryServer")]
        public Output<string?> RegistryServer { get; private set; }

        public Image(string name, ImageArgs args, ComponentResourceOptions? options = null)
            : base("docker:image:Image", name, options)
        {
            this.ImageName = Output.Tuple(args.ImageName.ToOutput(), args.LocalImageName.ToOutputNullable(), 
                                          args.SkipPush.ToOutput()).Apply(imageArgs =>
            {
                var imageName = imageArgs.Item1;

                Docker.RunCommandThatMustSucceed("docker", new[] { "ps"}, this).ConfigureAwait(false);

                // If there is no localImageName set it equal to imageName.  Note: this means
                // that if imageName contains a tag, localImageName will contain the same tag.
                var localImageName = imageArgs.Item2 ?? imageName;

                var skipPush = imageArgs.Item3;

                // Now break both the localImageName and the imageName into the untagged part and the
                // optional tag.  If both have tags, they must match.  If one or the other has a tag, we
                // just use that as the tag to use.  This allows users to flexibly provide a tag on one
                // option or the other and still have it work out.
                var (localImageNameWithoutTag, localImageNameTag) = Utils.GetImageNameAndTag(localImageName);
                var (imageNameWithoutTag, imageNameTag) = Utils.GetImageNameAndTag(imageName);

                if (localImageNameTag != null && imageNameTag != null && localImageNameTag != imageNameTag)
                    throw new InvalidOperationException(
                        $"[{nameof(args.LocalImageName)}] and [{nameof(args.ImageName)}] had mismatched tags. {localImageNameTag} != {imageNameTag}");

                var tag = localImageNameTag ?? imageNameTag;

                // BuildAndPushImageAsync expects only the baseImageName to have a tag.  So build that
                // name appropriately if we were given a tag.
                var baseImageName = tag != null ? $"{localImageNameWithoutTag}:{tag}" : localImageName;

                // BuildAndPushImageAsync does not want the repositoryUrl to have a tag.  This is just
                // the base url where the images will be pushed to.  All tagging will be taken care of
                // inside that api.
                var repositoryUrl = imageNameWithoutTag;
                return Docker.BuildAndPushImageAsync(
                    baseImageName,
                    args.Build,
                    repositoryUrl,
                    this,
                    skipPush,
                    args.Registry);
            });

            this.RegistryServer = args.Registry != null ? args.Registry.Apply(r => r.Server).Apply(v => (string?)v) : Output.Create((string?)null);
            this.BaseImageName = args.ImageName;

            this.RegisterOutputs(
                new Dictionary<string, object?>
                {
                    { "baseImageName", this.BaseImageName },
                    { "imageName", this.ImageName },
                    { "registryServer", this.RegistryServer },
                });
        }
    }
}
