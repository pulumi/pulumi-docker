// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker.buildx.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.core.internal.Codegen;
import com.pulumi.docker.buildx.enums.CompressionType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ExportRegistryArgs extends com.pulumi.resources.ResourceArgs {

    public static final ExportRegistryArgs Empty = new ExportRegistryArgs();

    @Import(name="annotations")
    private @Nullable Output<Map<String,String>> annotations;

    public Optional<Output<Map<String,String>>> annotations() {
        return Optional.ofNullable(this.annotations);
    }

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

    @Import(name="danglingNamePrefix")
    private @Nullable Output<String> danglingNamePrefix;

    public Optional<Output<String>> danglingNamePrefix() {
        return Optional.ofNullable(this.danglingNamePrefix);
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

    @Import(name="insecure")
    private @Nullable Output<Boolean> insecure;

    public Optional<Output<Boolean>> insecure() {
        return Optional.ofNullable(this.insecure);
    }

    @Import(name="nameCanonical")
    private @Nullable Output<Boolean> nameCanonical;

    public Optional<Output<Boolean>> nameCanonical() {
        return Optional.ofNullable(this.nameCanonical);
    }

    /**
     * Specify images names to export. This is overridden if tags are already specified.
     * 
     */
    @Import(name="names")
    private @Nullable Output<List<String>> names;

    /**
     * @return Specify images names to export. This is overridden if tags are already specified.
     * 
     */
    public Optional<Output<List<String>>> names() {
        return Optional.ofNullable(this.names);
    }

    /**
     * Use OCI media types in exporter manifests.
     * 
     */
    @Import(name="ociMediaTypes")
    private @Nullable Output<Boolean> ociMediaTypes;

    /**
     * @return Use OCI media types in exporter manifests.
     * 
     */
    public Optional<Output<Boolean>> ociMediaTypes() {
        return Optional.ofNullable(this.ociMediaTypes);
    }

    /**
     * Push after creating the image.
     * 
     */
    @Import(name="push")
    private @Nullable Output<Boolean> push;

    /**
     * @return Push after creating the image.
     * 
     */
    public Optional<Output<Boolean>> push() {
        return Optional.ofNullable(this.push);
    }

    @Import(name="pushByDigest")
    private @Nullable Output<Boolean> pushByDigest;

    public Optional<Output<Boolean>> pushByDigest() {
        return Optional.ofNullable(this.pushByDigest);
    }

    /**
     * Store resulting images to the worker&#39;s image store, and ensure all its
     * blobs are in the content store. Ignored if the worker doesn&#39;t have
     * image store (when using OCI workers, for example).
     * 
     */
    @Import(name="store")
    private @Nullable Output<Boolean> store;

    /**
     * @return
     * Store resulting images to the worker&#39;s image store, and ensure all its
     * blobs are in the content store. Ignored if the worker doesn&#39;t have
     * image store (when using OCI workers, for example).
     * 
     */
    public Optional<Output<Boolean>> store() {
        return Optional.ofNullable(this.store);
    }

    @Import(name="unpack")
    private @Nullable Output<Boolean> unpack;

    public Optional<Output<Boolean>> unpack() {
        return Optional.ofNullable(this.unpack);
    }

    private ExportRegistryArgs() {}

    private ExportRegistryArgs(ExportRegistryArgs $) {
        this.annotations = $.annotations;
        this.compression = $.compression;
        this.compressionLevel = $.compressionLevel;
        this.danglingNamePrefix = $.danglingNamePrefix;
        this.forceCompression = $.forceCompression;
        this.insecure = $.insecure;
        this.nameCanonical = $.nameCanonical;
        this.names = $.names;
        this.ociMediaTypes = $.ociMediaTypes;
        this.push = $.push;
        this.pushByDigest = $.pushByDigest;
        this.store = $.store;
        this.unpack = $.unpack;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ExportRegistryArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ExportRegistryArgs $;

        public Builder() {
            $ = new ExportRegistryArgs();
        }

        public Builder(ExportRegistryArgs defaults) {
            $ = new ExportRegistryArgs(Objects.requireNonNull(defaults));
        }

        public Builder annotations(@Nullable Output<Map<String,String>> annotations) {
            $.annotations = annotations;
            return this;
        }

        public Builder annotations(Map<String,String> annotations) {
            return annotations(Output.of(annotations));
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

        public Builder danglingNamePrefix(@Nullable Output<String> danglingNamePrefix) {
            $.danglingNamePrefix = danglingNamePrefix;
            return this;
        }

        public Builder danglingNamePrefix(String danglingNamePrefix) {
            return danglingNamePrefix(Output.of(danglingNamePrefix));
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

        public Builder insecure(@Nullable Output<Boolean> insecure) {
            $.insecure = insecure;
            return this;
        }

        public Builder insecure(Boolean insecure) {
            return insecure(Output.of(insecure));
        }

        public Builder nameCanonical(@Nullable Output<Boolean> nameCanonical) {
            $.nameCanonical = nameCanonical;
            return this;
        }

        public Builder nameCanonical(Boolean nameCanonical) {
            return nameCanonical(Output.of(nameCanonical));
        }

        /**
         * @param names Specify images names to export. This is overridden if tags are already specified.
         * 
         * @return builder
         * 
         */
        public Builder names(@Nullable Output<List<String>> names) {
            $.names = names;
            return this;
        }

        /**
         * @param names Specify images names to export. This is overridden if tags are already specified.
         * 
         * @return builder
         * 
         */
        public Builder names(List<String> names) {
            return names(Output.of(names));
        }

        /**
         * @param names Specify images names to export. This is overridden if tags are already specified.
         * 
         * @return builder
         * 
         */
        public Builder names(String... names) {
            return names(List.of(names));
        }

        /**
         * @param ociMediaTypes Use OCI media types in exporter manifests.
         * 
         * @return builder
         * 
         */
        public Builder ociMediaTypes(@Nullable Output<Boolean> ociMediaTypes) {
            $.ociMediaTypes = ociMediaTypes;
            return this;
        }

        /**
         * @param ociMediaTypes Use OCI media types in exporter manifests.
         * 
         * @return builder
         * 
         */
        public Builder ociMediaTypes(Boolean ociMediaTypes) {
            return ociMediaTypes(Output.of(ociMediaTypes));
        }

        /**
         * @param push Push after creating the image.
         * 
         * @return builder
         * 
         */
        public Builder push(@Nullable Output<Boolean> push) {
            $.push = push;
            return this;
        }

        /**
         * @param push Push after creating the image.
         * 
         * @return builder
         * 
         */
        public Builder push(Boolean push) {
            return push(Output.of(push));
        }

        public Builder pushByDigest(@Nullable Output<Boolean> pushByDigest) {
            $.pushByDigest = pushByDigest;
            return this;
        }

        public Builder pushByDigest(Boolean pushByDigest) {
            return pushByDigest(Output.of(pushByDigest));
        }

        /**
         * @param store
         * Store resulting images to the worker&#39;s image store, and ensure all its
         * blobs are in the content store. Ignored if the worker doesn&#39;t have
         * image store (when using OCI workers, for example).
         * 
         * @return builder
         * 
         */
        public Builder store(@Nullable Output<Boolean> store) {
            $.store = store;
            return this;
        }

        /**
         * @param store
         * Store resulting images to the worker&#39;s image store, and ensure all its
         * blobs are in the content store. Ignored if the worker doesn&#39;t have
         * image store (when using OCI workers, for example).
         * 
         * @return builder
         * 
         */
        public Builder store(Boolean store) {
            return store(Output.of(store));
        }

        public Builder unpack(@Nullable Output<Boolean> unpack) {
            $.unpack = unpack;
            return this;
        }

        public Builder unpack(Boolean unpack) {
            return unpack(Output.of(unpack));
        }

        public ExportRegistryArgs build() {
            $.compression = Codegen.objectProp("compression", CompressionType.class).output().arg($.compression).def(CompressionType.Gzip).getNullable();
            $.compressionLevel = Codegen.integerProp("compressionLevel").output().arg($.compressionLevel).def(0).getNullable();
            $.forceCompression = Codegen.booleanProp("forceCompression").output().arg($.forceCompression).def(false).getNullable();
            $.ociMediaTypes = Codegen.booleanProp("ociMediaTypes").output().arg($.ociMediaTypes).def(false).getNullable();
            $.push = Codegen.booleanProp("push").output().arg($.push).def(true).getNullable();
            $.store = Codegen.booleanProp("store").output().arg($.store).def(true).getNullable();
            return $;
        }
    }

}
