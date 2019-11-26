// Copyright 2016-2019, Pulumi Corporation.

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Linq;

namespace Pulumi.Docker
{
    internal static class UnwrapExtensions
    {
        public static Input<ImmutableDictionary<TKey, TValue>?> Unwrap<TKey, TValue>(this Input<Dictionary<TKey, Input<TValue>>>? dict) where TKey : notnull
        {
            if (dict == null)
                return Output.Create((ImmutableDictionary<TKey, TValue>?)null);

            return dict.Apply(d =>
            {
                var keyValues = d
                    .Select(kv => kv.Value.Apply(v => new KeyValuePair<TKey, TValue>(kv.Key, v)))
                    .Select(v => (Input<KeyValuePair<TKey, TValue>>)v)
                    .ToArray();
                return Output.All(keyValues).Apply(vs => vs.ToImmutableDictionary());
            }).Apply(v => (ImmutableDictionary<TKey, TValue>?)v);
        }

        public static Input<ImmutableArray<T>?> Unwrap<T>(this Input<Input<T>[]>? items)
        {
            if (items == null)
                return Output.Create((ImmutableArray<T>?)null);

            return items.Apply(v => Output.All(v)).Apply(v => (ImmutableArray<T>?)v);
        }

        public static Input<T?> Unwrap<T>(this Input<T>? input) where T : class
        {
            return input == null ? Output.Create((T?)null) : input.Apply(v => (T?)v);
        }

        public static Input<CacheFromUnwrap?> Unwrap(this InputUnion<bool, CacheFrom>? build)
        {
            if (build == null)
                return Output.Create((CacheFromUnwrap?)null);

            return build
                .Apply(v => v.IsT0 ? new ImmutableArray<string>() : v.AsT1.Stages.Unwrap())
                .Apply(v => v != null ? new CacheFromUnwrap { Stages = v } : null);
        }

        public static Input<Union<string, DockerBuildUnwrap>> Unwrap(this InputUnion<string, DockerBuild> build)
        {
            return build.Apply(b =>
            {
                if (b.IsT0)
                    return Output.Create((Union<string, DockerBuildUnwrap>)b.AsT0);

                var v = b.AsT1;
                return Output.Tuple(v.Args.Unwrap(), v.CacheFrom.Unwrap(), v.Context.Unwrap(), v.Dockerfile.Unwrap(), v.ExtraOptions.Unwrap()).Apply(vs =>
                {
                    return (Union<string, DockerBuildUnwrap>)new DockerBuildUnwrap
                    {
                        Args = vs.Item1,
                        CacheFrom = vs.Item2,
                        Context = vs.Item3,
                        Dockerfile = vs.Item4,
                        Env = v.Env,
                        ExtraOptions = vs.Item5,
                    };
                });
            });
        }

        public static Input<ImageRegistryUnwrap?> Unwrap(this Input<ImageRegistry>? registry)
        {
            if (registry == null)
                return Output.Create((ImageRegistryUnwrap?)null);

            return registry
                .Apply(v => Output.Tuple(v.Password, v.Server, v.Username))
                .Apply(v => (ImageRegistryUnwrap?)new ImageRegistryUnwrap { Password = v.Item1, Server = v.Item2, Username = v.Item3 });
        }
    }
}
