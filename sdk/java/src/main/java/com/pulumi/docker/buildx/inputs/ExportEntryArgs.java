// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker.buildx.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.docker.buildx.inputs.ExportDockerArgs;
import com.pulumi.docker.buildx.inputs.ExportImageArgs;
import com.pulumi.docker.buildx.inputs.ExportLocalArgs;
import com.pulumi.docker.buildx.inputs.ExportOCIArgs;
import com.pulumi.docker.buildx.inputs.ExportRegistryArgs;
import com.pulumi.docker.buildx.inputs.ExportTarArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ExportEntryArgs extends com.pulumi.resources.ResourceArgs {

    public static final ExportEntryArgs Empty = new ExportEntryArgs();

    /**
     * When &#34;true&#34; this entry will be excluded. Defaults to &#34;false&#34;.
     * 
     */
    @Import(name="disabled")
    private @Nullable Output<Boolean> disabled;

    /**
     * @return
     * When &#34;true&#34; this entry will be excluded. Defaults to &#34;false&#34;.
     * 
     */
    public Optional<Output<Boolean>> disabled() {
        return Optional.ofNullable(this.disabled);
    }

    /**
     * Export as a Docker image layout.
     * 
     */
    @Import(name="docker")
    private @Nullable Output<ExportDockerArgs> docker;

    /**
     * @return
     * Export as a Docker image layout.
     * 
     */
    public Optional<Output<ExportDockerArgs>> docker() {
        return Optional.ofNullable(this.docker);
    }

    /**
     * Outputs the build result into a container image format.
     * 
     */
    @Import(name="image")
    private @Nullable Output<ExportImageArgs> image;

    /**
     * @return
     * Outputs the build result into a container image format.
     * 
     */
    public Optional<Output<ExportImageArgs>> image() {
        return Optional.ofNullable(this.image);
    }

    /**
     * Export to a local directory as files and directories.
     * 
     */
    @Import(name="local")
    private @Nullable Output<ExportLocalArgs> local;

    /**
     * @return
     * Export to a local directory as files and directories.
     * 
     */
    public Optional<Output<ExportLocalArgs>> local() {
        return Optional.ofNullable(this.local);
    }

    /**
     * Identical to the Docker exporter but uses OCI media types by default.
     * 
     */
    @Import(name="oci")
    private @Nullable Output<ExportOCIArgs> oci;

    /**
     * @return
     * Identical to the Docker exporter but uses OCI media types by default.
     * 
     */
    public Optional<Output<ExportOCIArgs>> oci() {
        return Optional.ofNullable(this.oci);
    }

    /**
     * A raw string as you would provide it to the Docker CLI (e.g.,
     * &#34;type=docker&#34;)
     * 
     */
    @Import(name="raw")
    private @Nullable Output<String> raw;

    /**
     * @return
     * A raw string as you would provide it to the Docker CLI (e.g.,
     * &#34;type=docker&#34;)
     * 
     */
    public Optional<Output<String>> raw() {
        return Optional.ofNullable(this.raw);
    }

    /**
     * Identical to the Image exporter, but pushes by default.
     * 
     */
    @Import(name="registry")
    private @Nullable Output<ExportRegistryArgs> registry;

    /**
     * @return
     * Identical to the Image exporter, but pushes by default.
     * 
     */
    public Optional<Output<ExportRegistryArgs>> registry() {
        return Optional.ofNullable(this.registry);
    }

    /**
     * Export to a local directory as a tarball.
     * 
     */
    @Import(name="tar")
    private @Nullable Output<ExportTarArgs> tar;

    /**
     * @return
     * Export to a local directory as a tarball.
     * 
     */
    public Optional<Output<ExportTarArgs>> tar() {
        return Optional.ofNullable(this.tar);
    }

    private ExportEntryArgs() {}

    private ExportEntryArgs(ExportEntryArgs $) {
        this.disabled = $.disabled;
        this.docker = $.docker;
        this.image = $.image;
        this.local = $.local;
        this.oci = $.oci;
        this.raw = $.raw;
        this.registry = $.registry;
        this.tar = $.tar;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ExportEntryArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ExportEntryArgs $;

        public Builder() {
            $ = new ExportEntryArgs();
        }

        public Builder(ExportEntryArgs defaults) {
            $ = new ExportEntryArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param disabled
         * When &#34;true&#34; this entry will be excluded. Defaults to &#34;false&#34;.
         * 
         * @return builder
         * 
         */
        public Builder disabled(@Nullable Output<Boolean> disabled) {
            $.disabled = disabled;
            return this;
        }

        /**
         * @param disabled
         * When &#34;true&#34; this entry will be excluded. Defaults to &#34;false&#34;.
         * 
         * @return builder
         * 
         */
        public Builder disabled(Boolean disabled) {
            return disabled(Output.of(disabled));
        }

        /**
         * @param docker
         * Export as a Docker image layout.
         * 
         * @return builder
         * 
         */
        public Builder docker(@Nullable Output<ExportDockerArgs> docker) {
            $.docker = docker;
            return this;
        }

        /**
         * @param docker
         * Export as a Docker image layout.
         * 
         * @return builder
         * 
         */
        public Builder docker(ExportDockerArgs docker) {
            return docker(Output.of(docker));
        }

        /**
         * @param image
         * Outputs the build result into a container image format.
         * 
         * @return builder
         * 
         */
        public Builder image(@Nullable Output<ExportImageArgs> image) {
            $.image = image;
            return this;
        }

        /**
         * @param image
         * Outputs the build result into a container image format.
         * 
         * @return builder
         * 
         */
        public Builder image(ExportImageArgs image) {
            return image(Output.of(image));
        }

        /**
         * @param local
         * Export to a local directory as files and directories.
         * 
         * @return builder
         * 
         */
        public Builder local(@Nullable Output<ExportLocalArgs> local) {
            $.local = local;
            return this;
        }

        /**
         * @param local
         * Export to a local directory as files and directories.
         * 
         * @return builder
         * 
         */
        public Builder local(ExportLocalArgs local) {
            return local(Output.of(local));
        }

        /**
         * @param oci
         * Identical to the Docker exporter but uses OCI media types by default.
         * 
         * @return builder
         * 
         */
        public Builder oci(@Nullable Output<ExportOCIArgs> oci) {
            $.oci = oci;
            return this;
        }

        /**
         * @param oci
         * Identical to the Docker exporter but uses OCI media types by default.
         * 
         * @return builder
         * 
         */
        public Builder oci(ExportOCIArgs oci) {
            return oci(Output.of(oci));
        }

        /**
         * @param raw
         * A raw string as you would provide it to the Docker CLI (e.g.,
         * &#34;type=docker&#34;)
         * 
         * @return builder
         * 
         */
        public Builder raw(@Nullable Output<String> raw) {
            $.raw = raw;
            return this;
        }

        /**
         * @param raw
         * A raw string as you would provide it to the Docker CLI (e.g.,
         * &#34;type=docker&#34;)
         * 
         * @return builder
         * 
         */
        public Builder raw(String raw) {
            return raw(Output.of(raw));
        }

        /**
         * @param registry
         * Identical to the Image exporter, but pushes by default.
         * 
         * @return builder
         * 
         */
        public Builder registry(@Nullable Output<ExportRegistryArgs> registry) {
            $.registry = registry;
            return this;
        }

        /**
         * @param registry
         * Identical to the Image exporter, but pushes by default.
         * 
         * @return builder
         * 
         */
        public Builder registry(ExportRegistryArgs registry) {
            return registry(Output.of(registry));
        }

        /**
         * @param tar
         * Export to a local directory as a tarball.
         * 
         * @return builder
         * 
         */
        public Builder tar(@Nullable Output<ExportTarArgs> tar) {
            $.tar = tar;
            return this;
        }

        /**
         * @param tar
         * Export to a local directory as a tarball.
         * 
         * @return builder
         * 
         */
        public Builder tar(ExportTarArgs tar) {
            return tar(Output.of(tar));
        }

        public ExportEntryArgs build() {
            return $;
        }
    }

}
