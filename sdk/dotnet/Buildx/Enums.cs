// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.Docker.Buildx
{
    [EnumType]
    public readonly struct CacheMode : IEquatable<CacheMode>
    {
        private readonly string _value;

        private CacheMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Only layers that are exported into the resulting image are cached.
        /// </summary>
        public static CacheMode Min { get; } = new CacheMode("min");
        /// <summary>
        /// All layers are cached, even those of intermediate steps.
        /// </summary>
        public static CacheMode Max { get; } = new CacheMode("max");

        public static bool operator ==(CacheMode left, CacheMode right) => left.Equals(right);
        public static bool operator !=(CacheMode left, CacheMode right) => !left.Equals(right);

        public static explicit operator string(CacheMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is CacheMode other && Equals(other);
        public bool Equals(CacheMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct CompressionType : IEquatable<CompressionType>
    {
        private readonly string _value;

        private CompressionType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Use `gzip` for compression.
        /// </summary>
        public static CompressionType Gzip { get; } = new CompressionType("gzip");
        /// <summary>
        /// Use `estargz` for compression.
        /// </summary>
        public static CompressionType Estargz { get; } = new CompressionType("estargz");
        /// <summary>
        /// Use `zstd` for compression.
        /// </summary>
        public static CompressionType Zstd { get; } = new CompressionType("zstd");

        public static bool operator ==(CompressionType left, CompressionType right) => left.Equals(right);
        public static bool operator !=(CompressionType left, CompressionType right) => !left.Equals(right);

        public static explicit operator string(CompressionType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is CompressionType other && Equals(other);
        public bool Equals(CompressionType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct NetworkMode : IEquatable<NetworkMode>
    {
        private readonly string _value;

        private NetworkMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// The default sandbox network mode.
        /// </summary>
        public static NetworkMode @Default { get; } = new NetworkMode("default");
        /// <summary>
        /// Host network mode.
        /// </summary>
        public static NetworkMode Host { get; } = new NetworkMode("host");
        /// <summary>
        /// Disable network access.
        /// </summary>
        public static NetworkMode None { get; } = new NetworkMode("none");

        public static bool operator ==(NetworkMode left, NetworkMode right) => left.Equals(right);
        public static bool operator !=(NetworkMode left, NetworkMode right) => !left.Equals(right);

        public static explicit operator string(NetworkMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is NetworkMode other && Equals(other);
        public bool Equals(NetworkMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct Platform : IEquatable<Platform>
    {
        private readonly string _value;

        private Platform(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static Platform Darwin_386 { get; } = new Platform("darwin/386");
        public static Platform Darwin_amd64 { get; } = new Platform("darwin/amd64");
        public static Platform Darwin_arm { get; } = new Platform("darwin/arm");
        public static Platform Darwin_arm64 { get; } = new Platform("darwin/arm64");
        public static Platform Dragonfly_amd64 { get; } = new Platform("dragonfly/amd64");
        public static Platform Freebsd_386 { get; } = new Platform("freebsd/386");
        public static Platform Freebsd_amd64 { get; } = new Platform("freebsd/amd64");
        public static Platform Freebsd_arm { get; } = new Platform("freebsd/arm");
        public static Platform Linux_386 { get; } = new Platform("linux/386");
        public static Platform Linux_amd64 { get; } = new Platform("linux/amd64");
        public static Platform Linux_arm { get; } = new Platform("linux/arm");
        public static Platform Linux_arm64 { get; } = new Platform("linux/arm64");
        public static Platform Linux_mips64 { get; } = new Platform("linux/mips64");
        public static Platform Linux_mips64le { get; } = new Platform("linux/mips64le");
        public static Platform Linux_ppc64le { get; } = new Platform("linux/ppc64le");
        public static Platform Linux_riscv64 { get; } = new Platform("linux/riscv64");
        public static Platform Linux_s390x { get; } = new Platform("linux/s390x");
        public static Platform Netbsd_386 { get; } = new Platform("netbsd/386");
        public static Platform Netbsd_amd64 { get; } = new Platform("netbsd/amd64");
        public static Platform Netbsd_arm { get; } = new Platform("netbsd/arm");
        public static Platform Openbsd_386 { get; } = new Platform("openbsd/386");
        public static Platform Openbsd_amd64 { get; } = new Platform("openbsd/amd64");
        public static Platform Openbsd_arm { get; } = new Platform("openbsd/arm");
        public static Platform Plan9_386 { get; } = new Platform("plan9/386");
        public static Platform Plan9_amd64 { get; } = new Platform("plan9/amd64");
        public static Platform Solaris_amd64 { get; } = new Platform("solaris/amd64");
        public static Platform Windows_386 { get; } = new Platform("windows/386");
        public static Platform Windows_amd64 { get; } = new Platform("windows/amd64");

        public static bool operator ==(Platform left, Platform right) => left.Equals(right);
        public static bool operator !=(Platform left, Platform right) => !left.Equals(right);

        public static explicit operator string(Platform value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is Platform other && Equals(other);
        public bool Equals(Platform other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
