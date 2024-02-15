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
public final class ExportOCI {
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
     * @return The local export path.
     * 
     */
    private @Nullable String dest;
    /**
     * @return Forcefully apply compression.
     * 
     */
    private @Nullable Boolean forceCompression;
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
     * @return Bundle the output into a tarball layout.
     * 
     */
    private @Nullable Boolean tar;

    private ExportOCI() {}
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
     * @return The local export path.
     * 
     */
    public Optional<String> dest() {
        return Optional.ofNullable(this.dest);
    }
    /**
     * @return Forcefully apply compression.
     * 
     */
    public Optional<Boolean> forceCompression() {
        return Optional.ofNullable(this.forceCompression);
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
     * @return Bundle the output into a tarball layout.
     * 
     */
    public Optional<Boolean> tar() {
        return Optional.ofNullable(this.tar);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ExportOCI defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Map<String,String> annotations;
        private @Nullable CompressionType compression;
        private @Nullable Integer compressionLevel;
        private @Nullable String dest;
        private @Nullable Boolean forceCompression;
        private @Nullable List<String> names;
        private @Nullable Boolean ociMediaTypes;
        private @Nullable Boolean tar;
        public Builder() {}
        public Builder(ExportOCI defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.annotations = defaults.annotations;
    	      this.compression = defaults.compression;
    	      this.compressionLevel = defaults.compressionLevel;
    	      this.dest = defaults.dest;
    	      this.forceCompression = defaults.forceCompression;
    	      this.names = defaults.names;
    	      this.ociMediaTypes = defaults.ociMediaTypes;
    	      this.tar = defaults.tar;
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
        public Builder dest(@Nullable String dest) {

            this.dest = dest;
            return this;
        }
        @CustomType.Setter
        public Builder forceCompression(@Nullable Boolean forceCompression) {

            this.forceCompression = forceCompression;
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
        public Builder tar(@Nullable Boolean tar) {

            this.tar = tar;
            return this;
        }
        public ExportOCI build() {
            final var _resultValue = new ExportOCI();
            _resultValue.annotations = annotations;
            _resultValue.compression = compression;
            _resultValue.compressionLevel = compressionLevel;
            _resultValue.dest = dest;
            _resultValue.forceCompression = forceCompression;
            _resultValue.names = names;
            _resultValue.ociMediaTypes = ociMediaTypes;
            _resultValue.tar = tar;
            return _resultValue;
        }
    }
}
