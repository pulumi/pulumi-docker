// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
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
    /// ### Example Assuming you created a `config` as follows #!/bin/bash printf '{"a":"b"}' | docker config create foo - prints the id
    /// 
    /// 08c26c477474478d971139f750984775a7f019dbe8a2e7f09d66a187c009e66d you provide the definition for the resource as follows terraform resource "docker_config" "foo" {
    /// 
    ///  name = "foo"
    /// 
    ///  data = base64encode("{\"a\"\"b\"}") } then the import command is as follows #!/bin/bash
    /// 
    /// ```sh
    ///  $ pulumi import docker:index/serviceConfig:ServiceConfig foo 08c26c477474478d971139f750984775a7f019dbe8a2e7f09d66a187c009e66d
    /// ```
    /// </summary>
    [DockerResourceType("docker:index/serviceConfig:ServiceConfig")]
    public partial class ServiceConfig : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Base64-url-safe-encoded config data
        /// </summary>
        [Output("data")]
        public Output<string> Data { get; private set; } = null!;

        /// <summary>
        /// User-defined name of the config
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceConfig(string name, ServiceConfigArgs args, CustomResourceOptions? options = null)
            : base("docker:index/serviceConfig:ServiceConfig", name, args ?? new ServiceConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceConfig(string name, Input<string> id, ServiceConfigState? state = null, CustomResourceOptions? options = null)
            : base("docker:index/serviceConfig:ServiceConfig", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServiceConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceConfig Get(string name, Input<string> id, ServiceConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new ServiceConfig(name, id, state, options);
        }
    }

    public sealed class ServiceConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Base64-url-safe-encoded config data
        /// </summary>
        [Input("data", required: true)]
        public Input<string> Data { get; set; } = null!;

        /// <summary>
        /// User-defined name of the config
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public ServiceConfigArgs()
        {
        }
        public static new ServiceConfigArgs Empty => new ServiceConfigArgs();
    }

    public sealed class ServiceConfigState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Base64-url-safe-encoded config data
        /// </summary>
        [Input("data")]
        public Input<string>? Data { get; set; }

        /// <summary>
        /// User-defined name of the config
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public ServiceConfigState()
        {
        }
        public static new ServiceConfigState Empty => new ServiceConfigState();
    }
}
