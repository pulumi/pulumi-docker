// Copyright 2016-2019, Pulumi Corporation.

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Linq;

namespace Pulumi.Docker
{
    internal static class UnwrapExtensions
    {
        public static Output<T?> ToOutputNullable<T>(this Input<T>? input) where T : class
            => input == null ? Output.Create((T?)null) : input.Apply(v => (T?)v);

        public static Output<Union<T0, T1>?> ToOutputNullable<T0, T1>(this InputUnion<T0, T1>? items)
            => items == null ? Output.Create((Union<T0, T1>?)null) : items.Apply(v => (Union<T0, T1>?)v);

        public static Output<U?> ApplyNullable<T, U>(this Output<T?> output, Func<T, Output<U>> func) where T : class where U : class
            => output.Apply(v => v != null ? func(v).Apply(v => (U?)v) : Output.Create((U?)null));

        public static Output<U?> ApplyNullable<T, U>(this Output<T?> output, Func<T, Output<U>> func) where T : struct where U : class
            => output.Apply(v => v.HasValue ? func(v.Value).Apply(v => (U?)v) : Output.Create((U?)null));

        public static Output<Union<U, T1>> ApplyT0<T0, T1, U>(this Output<Union<T0, T1>> union, Func<T0, Output<U>> func)
            => union.Apply(v => v.IsT0 ? func(v.AsT0).Apply(v => Union<U, T1>.FromT0(v)) : Output.Create(Union<U, T1>.FromT1(v.AsT1)));

        public static Output<Union<T0, U>> ApplyT1<T0, T1, U>(this Output<Union<T0, T1>> union, Func<T1, Output<U>> func)
            => union.Apply(v => v.IsT1 ? func(v.AsT1).Apply(v => Union<T0, U>.FromT1(v)) : Output.Create(Union<T0, U>.FromT0(v.AsT0)));
    }

    internal static class DockerExtensions
    { 
        public static Output<CacheFromUnwrap?> Unwrap(this InputUnion<bool, CacheFrom>? build)
        {
            return build.ToOutputNullable().ApplyNullable(v =>
            {
                if (v.IsT0)
                    return Output.Create(new CacheFromUnwrap { Stages = new ImmutableArray<string>() });

                return v.AsT1.Stages.ToOutput().Apply(b => new CacheFromUnwrap { Stages = b });
            });
        }

        public static Output<Union<string, DockerBuildUnwrap>> Unwrap(this InputUnion<string, DockerBuild> build)
        {
            return build
                .ToOutput()
                .ApplyT1(v => 
                    Output.Tuple(v.Args.ToOutput(), v.CacheFrom.Unwrap(), v.Context.ToOutput(),
                                 v.Dockerfile.ToOutput(), v.Env.ToOutput(),
                                 v.ExtraOptions.ToOutput(), v.Target.ToOutput())
                          .Apply(vs => new DockerBuildUnwrap
                          {
                              Args = vs.Item1,
                              CacheFrom = vs.Item2,
                              Context = vs.Item3,
                              Dockerfile = vs.Item4,
                              Env = vs.Item5,
                              ExtraOptions = vs.Item6,
                              Target = vs.Item7,
                          }));
        }

        public static Output<ImageRegistryUnwrap?> Unwrap(this Input<ImageRegistry>? registry)
        {
            return registry
                .ToOutputNullable()
                .ApplyNullable(v => Output.Tuple(v.Password, v.Server, v.Username)
                                          .Apply(v => new ImageRegistryUnwrap { Password = v.Item1, Server = v.Item2, Username = v.Item3 }));
        }
    }
}
