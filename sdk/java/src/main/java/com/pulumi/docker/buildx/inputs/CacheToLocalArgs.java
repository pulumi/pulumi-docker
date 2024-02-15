// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker.buildx.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.core.internal.Codegen;
import com.pulumi.docker.buildx.enums.CacheMode;
import com.pulumi.docker.buildx.enums.CompressionType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class CacheToLocalArgs extends com.pulumi.resources.ResourceArgs {

    public static final CacheToLocalArgs Empty = new CacheToLocalArgs();

    /**
     * The compression type to use.
     * 
     */
    @Import(name="compression")
    private @Nullable Output<CompressionType> compression;

    /**
     * @return The compression type to use.
     * 
     */
    public Optional<Output<CompressionType>> compression() {
        return Optional.ofNullable(this.compression);
    }

    /**
     * Compression level from 0 to 22.
     * 
     */
    @Import(name="compressionLevel")
    private @Nullable Output<Integer> compressionLevel;

    /**
     * @return Compression level from 0 to 22.
     * 
     */
    public Optional<Output<Integer>> compressionLevel() {
        return Optional.ofNullable(this.compressionLevel);
    }

    /**
     * Path of the local directory to export the cache.
     * 
     */
    @Import(name="dest", required=true)
    private Output<String> dest;

    /**
     * @return Path of the local directory to export the cache.
     * 
     */
    public Output<String> dest() {
        return this.dest;
    }

    /**
     * Forcefully apply compression.
     * 
     */
    @Import(name="forceCompression")
    private @Nullable Output<Boolean> forceCompression;

    /**
     * @return Forcefully apply compression.
     * 
     */
    public Optional<Output<Boolean>> forceCompression() {
        return Optional.ofNullable(this.forceCompression);
    }

    /**
     * Ignore errors caused by failed cache exports.
     * 
     */
    @Import(name="ignoreError")
    private @Nullable Output<Boolean> ignoreError;

    /**
     * @return Ignore errors caused by failed cache exports.
     * 
     */
    public Optional<Output<Boolean>> ignoreError() {
        return Optional.ofNullable(this.ignoreError);
    }

    /**
     * The cache mode to use. Defaults to `min`.
     * 
     */
    @Import(name="mode")
    private @Nullable Output<CacheMode> mode;

    /**
     * @return The cache mode to use. Defaults to `min`.
     * 
     */
    public Optional<Output<CacheMode>> mode() {
        return Optional.ofNullable(this.mode);
    }

    private CacheToLocalArgs() {}

    private CacheToLocalArgs(CacheToLocalArgs $) {
        this.compression = $.compression;
        this.compressionLevel = $.compressionLevel;
        this.dest = $.dest;
        this.forceCompression = $.forceCompression;
        this.ignoreError = $.ignoreError;
        this.mode = $.mode;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CacheToLocalArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CacheToLocalArgs $;

        public Builder() {
            $ = new CacheToLocalArgs();
        }

        public Builder(CacheToLocalArgs defaults) {
            $ = new CacheToLocalArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param compression The compression type to use.
         * 
         * @return builder
         * 
         */
        public Builder compression(@Nullable Output<CompressionType> compression) {
            $.compression = compression;
            return this;
        }

        /**
         * @param compression The compression type to use.
         * 
         * @return builder
         * 
         */
        public Builder compression(CompressionType compression) {
            return compression(Output.of(compression));
        }

        /**
         * @param compressionLevel Compression level from 0 to 22.
         * 
         * @return builder
         * 
         */
        public Builder compressionLevel(@Nullable Output<Integer> compressionLevel) {
            $.compressionLevel = compressionLevel;
            return this;
        }

        /**
         * @param compressionLevel Compression level from 0 to 22.
         * 
         * @return builder
         * 
         */
        public Builder compressionLevel(Integer compressionLevel) {
            return compressionLevel(Output.of(compressionLevel));
        }

        /**
         * @param dest Path of the local directory to export the cache.
         * 
         * @return builder
         * 
         */
        public Builder dest(Output<String> dest) {
            $.dest = dest;
            return this;
        }

        /**
         * @param dest Path of the local directory to export the cache.
         * 
         * @return builder
         * 
         */
        public Builder dest(String dest) {
            return dest(Output.of(dest));
        }

        /**
         * @param forceCompression Forcefully apply compression.
         * 
         * @return builder
         * 
         */
        public Builder forceCompression(@Nullable Output<Boolean> forceCompression) {
            $.forceCompression = forceCompression;
            return this;
        }

        /**
         * @param forceCompression Forcefully apply compression.
         * 
         * @return builder
         * 
         */
        public Builder forceCompression(Boolean forceCompression) {
            return forceCompression(Output.of(forceCompression));
        }

        /**
         * @param ignoreError Ignore errors caused by failed cache exports.
         * 
         * @return builder
         * 
         */
        public Builder ignoreError(@Nullable Output<Boolean> ignoreError) {
            $.ignoreError = ignoreError;
            return this;
        }

        /**
         * @param ignoreError Ignore errors caused by failed cache exports.
         * 
         * @return builder
         * 
         */
        public Builder ignoreError(Boolean ignoreError) {
            return ignoreError(Output.of(ignoreError));
        }

        /**
         * @param mode The cache mode to use. Defaults to `min`.
         * 
         * @return builder
         * 
         */
        public Builder mode(@Nullable Output<CacheMode> mode) {
            $.mode = mode;
            return this;
        }

        /**
         * @param mode The cache mode to use. Defaults to `min`.
         * 
         * @return builder
         * 
         */
        public Builder mode(CacheMode mode) {
            return mode(Output.of(mode));
        }

        public CacheToLocalArgs build() {
            $.compression = Codegen.objectProp("compression", CompressionType.class).output().arg($.compression).def(CompressionType.Gzip).getNullable();
            $.compressionLevel = Codegen.integerProp("compressionLevel").output().arg($.compressionLevel).def(0).getNullable();
            if ($.dest == null) {
                throw new MissingRequiredPropertyException("CacheToLocalArgs", "dest");
            }
            $.forceCompression = Codegen.booleanProp("forceCompression").output().arg($.forceCompression).def(false).getNullable();
            $.ignoreError = Codegen.booleanProp("ignoreError").output().arg($.ignoreError).def(false).getNullable();
            $.mode = Codegen.objectProp("mode", CacheMode.class).output().arg($.mode).def(CacheMode.Min).getNullable();
            return $;
        }
    }

}
