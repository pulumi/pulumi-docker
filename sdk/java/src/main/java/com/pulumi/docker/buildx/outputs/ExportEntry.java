// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker.buildx.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.docker.buildx.outputs.ExportDocker;
import com.pulumi.docker.buildx.outputs.ExportImage;
import com.pulumi.docker.buildx.outputs.ExportLocal;
import com.pulumi.docker.buildx.outputs.ExportOCI;
import com.pulumi.docker.buildx.outputs.ExportRegistry;
import com.pulumi.docker.buildx.outputs.ExportTar;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ExportEntry {
    /**
     * @return
     * When &#34;true&#34; this entry will be excluded. Defaults to &#34;false&#34;.
     * 
     */
    private @Nullable Boolean disabled;
    /**
     * @return
     * Export as a Docker image layout.
     * 
     */
    private @Nullable ExportDocker docker;
    /**
     * @return
     * Outputs the build result into a container image format.
     * 
     */
    private @Nullable ExportImage image;
    /**
     * @return
     * Export to a local directory as files and directories.
     * 
     */
    private @Nullable ExportLocal local;
    /**
     * @return
     * Identical to the Docker exporter but uses OCI media types by default.
     * 
     */
    private @Nullable ExportOCI oci;
    /**
     * @return
     * A raw string as you would provide it to the Docker CLI (e.g.,
     * &#34;type=docker&#34;)
     * 
     */
    private @Nullable String raw;
    /**
     * @return
     * Identical to the Image exporter, but pushes by default.
     * 
     */
    private @Nullable ExportRegistry registry;
    /**
     * @return
     * Export to a local directory as a tarball.
     * 
     */
    private @Nullable ExportTar tar;

    private ExportEntry() {}
    /**
     * @return
     * When &#34;true&#34; this entry will be excluded. Defaults to &#34;false&#34;.
     * 
     */
    public Optional<Boolean> disabled() {
        return Optional.ofNullable(this.disabled);
    }
    /**
     * @return
     * Export as a Docker image layout.
     * 
     */
    public Optional<ExportDocker> docker() {
        return Optional.ofNullable(this.docker);
    }
    /**
     * @return
     * Outputs the build result into a container image format.
     * 
     */
    public Optional<ExportImage> image() {
        return Optional.ofNullable(this.image);
    }
    /**
     * @return
     * Export to a local directory as files and directories.
     * 
     */
    public Optional<ExportLocal> local() {
        return Optional.ofNullable(this.local);
    }
    /**
     * @return
     * Identical to the Docker exporter but uses OCI media types by default.
     * 
     */
    public Optional<ExportOCI> oci() {
        return Optional.ofNullable(this.oci);
    }
    /**
     * @return
     * A raw string as you would provide it to the Docker CLI (e.g.,
     * &#34;type=docker&#34;)
     * 
     */
    public Optional<String> raw() {
        return Optional.ofNullable(this.raw);
    }
    /**
     * @return
     * Identical to the Image exporter, but pushes by default.
     * 
     */
    public Optional<ExportRegistry> registry() {
        return Optional.ofNullable(this.registry);
    }
    /**
     * @return
     * Export to a local directory as a tarball.
     * 
     */
    public Optional<ExportTar> tar() {
        return Optional.ofNullable(this.tar);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ExportEntry defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean disabled;
        private @Nullable ExportDocker docker;
        private @Nullable ExportImage image;
        private @Nullable ExportLocal local;
        private @Nullable ExportOCI oci;
        private @Nullable String raw;
        private @Nullable ExportRegistry registry;
        private @Nullable ExportTar tar;
        public Builder() {}
        public Builder(ExportEntry defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.disabled = defaults.disabled;
    	      this.docker = defaults.docker;
    	      this.image = defaults.image;
    	      this.local = defaults.local;
    	      this.oci = defaults.oci;
    	      this.raw = defaults.raw;
    	      this.registry = defaults.registry;
    	      this.tar = defaults.tar;
        }

        @CustomType.Setter
        public Builder disabled(@Nullable Boolean disabled) {

            this.disabled = disabled;
            return this;
        }
        @CustomType.Setter
        public Builder docker(@Nullable ExportDocker docker) {

            this.docker = docker;
            return this;
        }
        @CustomType.Setter
        public Builder image(@Nullable ExportImage image) {

            this.image = image;
            return this;
        }
        @CustomType.Setter
        public Builder local(@Nullable ExportLocal local) {

            this.local = local;
            return this;
        }
        @CustomType.Setter
        public Builder oci(@Nullable ExportOCI oci) {

            this.oci = oci;
            return this;
        }
        @CustomType.Setter
        public Builder raw(@Nullable String raw) {

            this.raw = raw;
            return this;
        }
        @CustomType.Setter
        public Builder registry(@Nullable ExportRegistry registry) {

            this.registry = registry;
            return this;
        }
        @CustomType.Setter
        public Builder tar(@Nullable ExportTar tar) {

            this.tar = tar;
            return this;
        }
        public ExportEntry build() {
            final var _resultValue = new ExportEntry();
            _resultValue.disabled = disabled;
            _resultValue.docker = docker;
            _resultValue.image = image;
            _resultValue.local = local;
            _resultValue.oci = oci;
            _resultValue.raw = raw;
            _resultValue.registry = registry;
            _resultValue.tar = tar;
            return _resultValue;
        }
    }
}
