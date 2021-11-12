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
    public sealed class ContainerHealthcheck
    {
        /// <summary>
        /// Time between running the check (ms|s|m|h). Defaults to `0s`.
        /// </summary>
        public readonly string? Interval;
        /// <summary>
        /// Consecutive failures needed to report unhealthy. Defaults to `0`.
        /// </summary>
        public readonly int? Retries;
        /// <summary>
        /// Start period for the container to initialize before counting retries towards unstable (ms|s|m|h). Defaults to `0s`.
        /// </summary>
        public readonly string? StartPeriod;
        /// <summary>
        /// Command to run to check health. For example, to run `curl -f localhost/health` set the command to be `["CMD", "curl", "-f", "localhost/health"]`.
        /// </summary>
        public readonly ImmutableArray<string> Tests;
        /// <summary>
        /// Maximum time to allow one check to run (ms|s|m|h). Defaults to `0s`.
        /// </summary>
        public readonly string? Timeout;

        [OutputConstructor]
        private ContainerHealthcheck(
            string? interval,

            int? retries,

            string? startPeriod,

            ImmutableArray<string> tests,

            string? timeout)
        {
            Interval = interval;
            Retries = retries;
            StartPeriod = startPeriod;
            Tests = tests;
            Timeout = timeout;
        }
    }
}
