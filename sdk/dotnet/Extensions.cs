// Copyright 2016-2019, Pulumi Corporation.

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Linq;

namespace Pulumi.Docker
{
    internal static class UnwrapExtensions
    {
        public static Input<ImmutableArray<T>?> Unwrap<T>(this InputList<T>? items)
            => items == null ? Output.Create((ImmutableArray<T>?)null) : items.Apply(v => (ImmutableArray<T>?)v);

        public static Input<TUnwrapped?> Unwrap<T, TUnwrapped>(this Input<T>? input, Func<T, Input<TUnwrapped?>> convert) where TUnwrapped : class
            => input == null ? Output.Create((TUnwrapped?)null) : input.Apply(convert);

        public static Input<TUnwrapped?> Unwrap<T, TUnwrapped>(this Input<T>? input, Func<T, Input<TUnwrapped?>> convert) where TUnwrapped : struct
            => input == null ? Output.Create((TUnwrapped?)null) : input.Apply(convert);

        public static Input<T?> Unwrap<T>(this Input<T>? input) where T : class
            => input.Unwrap<T, T>(v => v);

        public static Input<Union<T0Unwrapped, T1Unwrapped>?> Unwrap<T0, T1, T0Unwrapped, T1Unwrapped>
            (this InputUnion<T0, T1>? input, Func<T0, Output<T0Unwrapped>> unwrapT0, Func<T1, Output<T1Unwrapped>> unwrapT1)
        {
            if (input == null)
                Output.Create((Union<T0Unwrapped, T1Unwrapped>?)null);

            return input.Apply(value =>
            {
                return value.IsT0
                    ? unwrapT0(value.AsT0).Apply(Union<T0Unwrapped, T1Unwrapped>.FromT0)
                    : unwrapT1(value.AsT1).Apply(Union<T0Unwrapped, T1Unwrapped>.FromT1);
            }).Apply(v => (Union<T0Unwrapped, T1Unwrapped>?)v);
        }
    }

    internal static class DockerExtensions
    { 
        public static Input<CacheFromUnwrap?> Unwrap(this InputUnion<bool, CacheFrom>? build)
        {
            return build.Unwrap(v => v.IsT0 ? new ImmutableArray<string>() : v.AsT1.Stages.Unwrap())
                .Apply(v => v != null ? new CacheFromUnwrap { Stages = v } : null);
        }

        public static Input<Union<string, DockerBuildUnwrap>> Unwrap(this InputUnion<string, DockerBuild> build)
        {
            var a = build.Unwrap(
                x => Output.Create(x),
                v => Output.Tuple(v.Args.Unwrap(), v.CacheFrom.Unwrap(), v.Context.Unwrap(), v.Dockerfile.Unwrap(), v.Env.Unwrap(), v.ExtraOptions.Unwrap())
                           .Apply(vs => new DockerBuildUnwrap
                           {
                               Args = vs.Item1,
                               CacheFrom = vs.Item2,
                               Context = vs.Item3,
                               Dockerfile = vs.Item4,
                               Env = vs.Item5,
                               ExtraOptions = vs.Item6,
                           }));
            return a.Apply(v => v ?? throw new InvalidOperationException($"[{nameof(build)}] is never never, so unwrapped value is expected to be never null either"));
        }

        public static Input<ImageRegistryUnwrap?> Unwrap(this Input<ImageRegistry>? registry)
            => registry
                .Unwrap<ImageRegistry, ImageRegistryUnwrap>(v => 
                    Output.Tuple(v.Password, v.Server, v.Username)
                          .Apply(v => (ImageRegistryUnwrap?)new ImageRegistryUnwrap { Password = v.Item1, Server = v.Item2, Username = v.Item3 }));
    }
}
