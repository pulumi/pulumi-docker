// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker.buildx.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.docker.buildx.enums.CompressionType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ExportImage {
    /**
     * @return Attach an arbitrary key/value annotation to the image.
     * 
     */
    private @Nullable Map<String,String> annotations;
    /**
     * @return The compression type to use.
     * 
     */
    private @Nullable CompressionType compression;
    /**
     * @return Compression level from 0 to 22.
     * 
     */
    private @Nullable Integer compressionLevel;
    /**
     * @return Name image with `prefix@&lt;digest&gt;`, used for anonymous images.
     * 
     */
    private @Nullable String danglingNamePrefix;
    /**
     * @return Forcefully apply compression.
     * 
     */
    private @Nullable Boolean forceCompression;
    /**
     * @return Allow pushing to an insecure registry.
     * 
     */
    private @Nullable Boolean insecure;
    /**
     * @return Add additional canonical name (`name@&lt;digest&gt;`).
     * 
     */
    private @Nullable Boolean nameCanonical;
    /**
     * @return Specify images names to export. This is overridden if tags are already specified.
     * 
     */
    private @Nullable List<String> names;
    /**
     * @return Use OCI media types in exporter manifests.
     * 
     */
    private @Nullable Boolean ociMediaTypes;
    /**
     * @return Push after creating the image.
     * 
     */
    private @Nullable Boolean push;
    /**
     * @return Push image without name.
     * 
     */
    private @Nullable Boolean pushByDigest;
    /**
     * @return Store resulting images to the worker&#39;s image store and ensure all of
     * its blobs are in the content store.
     * 
     * Defaults to `true`.
     * 
     * Ignored if the worker doesn&#39;t have image store (when using OCI workers,
     * for example).
     * 
     */
    private @Nullable Boolean store;
    /**
     * @return Unpack image after creation (for use with containerd). Defaults to
     * `false`.
     * 
     */
    private @Nullable Boolean unpack;

    private ExportImage() {}
    /**
     * @return Attach an arbitrary key/value annotation to the image.
     * 
     */
    public Map<String,String> annotations() {
        return this.annotations == null ? Map.of() : this.annotations;
    }
    /**
     * @return The compression type to use.
     * 
     */
    public Optional<CompressionType> compression() {
        return Optional.ofNullable(this.compression);
    }
    /**
     * @return Compression level from 0 to 22.
     * 
     */
    public Optional<Integer> compressionLevel() {
        return Optional.ofNullable(this.compressionLevel);
    }
    /**
     * @return Name image with `prefix@&lt;digest&gt;`, used for anonymous images.
     * 
     */
    public Optional<String> danglingNamePrefix() {
        return Optional.ofNullable(this.danglingNamePrefix);
    }
    /**
     * @return Forcefully apply compression.
     * 
     */
    public Optional<Boolean> forceCompression() {
        return Optional.ofNullable(this.forceCompression);
    }
    /**
     * @return Allow pushing to an insecure registry.
     * 
     */
    public Optional<Boolean> insecure() {
        return Optional.ofNullable(this.insecure);
    }
    /**
     * @return Add additional canonical name (`name@&lt;digest&gt;`).
     * 
     */
    public Optional<Boolean> nameCanonical() {
        return Optional.ofNullable(this.nameCanonical);
    }
    /**
     * @return Specify images names to export. This is overridden if tags are already specified.
     * 
     */
    public List<String> names() {
        return this.names == null ? List.of() : this.names;
    }
    /**
     * @return Use OCI media types in exporter manifests.
     * 
     */
    public Optional<Boolean> ociMediaTypes() {
        return Optional.ofNullable(this.ociMediaTypes);
    }
    /**
     * @return Push after creating the image.
     * 
     */
    public Optional<Boolean> push() {
        return Optional.ofNullable(this.push);
    }
    /**
     * @return Push image without name.
     * 
     */
    public Optional<Boolean> pushByDigest() {
        return Optional.ofNullable(this.pushByDigest);
    }
    /**
     * @return Store resulting images to the worker&#39;s image store and ensure all of
     * its blobs are in the content store.
     * 
     * Defaults to `true`.
     * 
     * Ignored if the worker doesn&#39;t have image store (when using OCI workers,
     * for example).
     * 
     */
    public Optional<Boolean> store() {
        return Optional.ofNullable(this.store);
    }
    /**
     * @return Unpack image after creation (for use with containerd). Defaults to
     * `false`.
     * 
     */
    public Optional<Boolean> unpack() {
        return Optional.ofNullable(this.unpack);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ExportImage defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Map<String,String> annotations;
        private @Nullable CompressionType compression;
        private @Nullable Integer compressionLevel;
        private @Nullable String danglingNamePrefix;
        private @Nullable Boolean forceCompression;
        private @Nullable Boolean insecure;
        private @Nullable Boolean nameCanonical;
        private @Nullable List<String> names;
        private @Nullable Boolean ociMediaTypes;
        private @Nullable Boolean push;
        private @Nullable Boolean pushByDigest;
        private @Nullable Boolean store;
        private @Nullable Boolean unpack;
        public Builder() {}
        public Builder(ExportImage defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.annotations = defaults.annotations;
    	      this.compression = defaults.compression;
    	      this.compressionLevel = defaults.compressionLevel;
    	      this.danglingNamePrefix = defaults.danglingNamePrefix;
    	      this.forceCompression = defaults.forceCompression;
    	      this.insecure = defaults.insecure;
    	      this.nameCanonical = defaults.nameCanonical;
    	      this.names = defaults.names;
    	      this.ociMediaTypes = defaults.ociMediaTypes;
    	      this.push = defaults.push;
    	      this.pushByDigest = defaults.pushByDigest;
    	      this.store = defaults.store;
    	      this.unpack = defaults.unpack;
        }

        @CustomType.Setter
        public Builder annotations(@Nullable Map<String,String> annotations) {

            this.annotations = annotations;
            return this;
        }
        @CustomType.Setter
        public Builder compression(@Nullable CompressionType compression) {

            this.compression = compression;
            return this;
        }
        @CustomType.Setter
        public Builder compressionLevel(@Nullable Integer compressionLevel) {

            this.compressionLevel = compressionLevel;
            return this;
        }
        @CustomType.Setter
        public Builder danglingNamePrefix(@Nullable String danglingNamePrefix) {

            this.danglingNamePrefix = danglingNamePrefix;
            return this;
        }
        @CustomType.Setter
        public Builder forceCompression(@Nullable Boolean forceCompression) {

            this.forceCompression = forceCompression;
            return this;
        }
        @CustomType.Setter
        public Builder insecure(@Nullable Boolean insecure) {

            this.insecure = insecure;
            return this;
        }
        @CustomType.Setter
        public Builder nameCanonical(@Nullable Boolean nameCanonical) {

            this.nameCanonical = nameCanonical;
            return this;
        }
        @CustomType.Setter
        public Builder names(@Nullable List<String> names) {

            this.names = names;
            return this;
        }
        public Builder names(String... names) {
            return names(List.of(names));
        }
        @CustomType.Setter
        public Builder ociMediaTypes(@Nullable Boolean ociMediaTypes) {

            this.ociMediaTypes = ociMediaTypes;
            return this;
        }
        @CustomType.Setter
        public Builder push(@Nullable Boolean push) {

            this.push = push;
            return this;
        }
        @CustomType.Setter
        public Builder pushByDigest(@Nullable Boolean pushByDigest) {

            this.pushByDigest = pushByDigest;
            return this;
        }
        @CustomType.Setter
        public Builder store(@Nullable Boolean store) {

            this.store = store;
            return this;
        }
        @CustomType.Setter
        public Builder unpack(@Nullable Boolean unpack) {

            this.unpack = unpack;
            return this;
        }
        public ExportImage build() {
            final var _resultValue = new ExportImage();
            _resultValue.annotations = annotations;
            _resultValue.compression = compression;
            _resultValue.compressionLevel = compressionLevel;
            _resultValue.danglingNamePrefix = danglingNamePrefix;
            _resultValue.forceCompression = forceCompression;
            _resultValue.insecure = insecure;
            _resultValue.nameCanonical = nameCanonical;
            _resultValue.names = names;
            _resultValue.ociMediaTypes = ociMediaTypes;
            _resultValue.push = push;
            _resultValue.pushByDigest = pushByDigest;
            _resultValue.store = store;
            _resultValue.unpack = unpack;
            return _resultValue;
        }
    }
}
