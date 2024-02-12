// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker.buildx;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.core.internal.Codegen;
import com.pulumi.docker.buildx.enums.Platform;
import com.pulumi.docker.buildx.inputs.CacheFromEntryArgs;
import com.pulumi.docker.buildx.inputs.CacheToEntryArgs;
import com.pulumi.docker.buildx.inputs.ExportEntryArgs;
import com.pulumi.docker.buildx.inputs.RegistryAuthArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ImageArgs extends com.pulumi.resources.ResourceArgs {

    public static final ImageArgs Empty = new ImageArgs();

    /**
     * An optional map of named build-time argument variables to set during
     * the Docker build. This flag allows you to pass build-time variables that
     * can be accessed like environment variables inside the RUN
     * instruction.
     * 
     */
    @Import(name="buildArgs")
    private @Nullable Output<Map<String,String>> buildArgs;

    /**
     * @return
     * An optional map of named build-time argument variables to set during
     * the Docker build. This flag allows you to pass build-time variables that
     * can be accessed like environment variables inside the RUN
     * instruction.
     * 
     */
    public Optional<Output<Map<String,String>>> buildArgs() {
        return Optional.ofNullable(this.buildArgs);
    }

    /**
     * When true, attempt to build the image during previews. Outputs are not
     * pushed to registries, however caches are still populated.
     * 
     */
    @Import(name="buildOnPreview")
    private @Nullable Output<Boolean> buildOnPreview;

    /**
     * @return
     * When true, attempt to build the image during previews. Outputs are not
     * pushed to registries, however caches are still populated.
     * 
     */
    public Optional<Output<Boolean>> buildOnPreview() {
        return Optional.ofNullable(this.buildOnPreview);
    }

    /**
     * Build with a specific builder instance
     * 
     */
    @Import(name="builder")
    private @Nullable Output<String> builder;

    /**
     * @return
     * Build with a specific builder instance
     * 
     */
    public Optional<Output<String>> builder_() {
        return Optional.ofNullable(this.builder);
    }

    /**
     * External cache sources (e.g., &#34;user/app:cache&#34;, &#34;type=local,src=path/to/dir&#34;)
     * 
     */
    @Import(name="cacheFrom")
    private @Nullable Output<List<CacheFromEntryArgs>> cacheFrom;

    /**
     * @return
     * External cache sources (e.g., &#34;user/app:cache&#34;, &#34;type=local,src=path/to/dir&#34;)
     * 
     */
    public Optional<Output<List<CacheFromEntryArgs>>> cacheFrom() {
        return Optional.ofNullable(this.cacheFrom);
    }

    /**
     * Cache export destinations (e.g., &#34;user/app:cache&#34;, &#34;type=local,dest=path/to/dir&#34;)
     * 
     */
    @Import(name="cacheTo")
    private @Nullable Output<List<CacheToEntryArgs>> cacheTo;

    /**
     * @return
     * Cache export destinations (e.g., &#34;user/app:cache&#34;, &#34;type=local,dest=path/to/dir&#34;)
     * 
     */
    public Optional<Output<List<CacheToEntryArgs>>> cacheTo() {
        return Optional.ofNullable(this.cacheTo);
    }

    /**
     * Path to use for build context. If omitted, an empty context is used.
     * 
     */
    @Import(name="context")
    private @Nullable Output<String> context;

    /**
     * @return
     * Path to use for build context. If omitted, an empty context is used.
     * 
     */
    public Optional<Output<String>> context() {
        return Optional.ofNullable(this.context);
    }

    /**
     * Name and optionally a tag (format: &#34;name:tag&#34;). If outputting to a
     * registry, the name should include the fully qualified registry address.
     * 
     */
    @Import(name="exports")
    private @Nullable Output<List<ExportEntryArgs>> exports;

    /**
     * @return
     * Name and optionally a tag (format: &#34;name:tag&#34;). If outputting to a
     * registry, the name should include the fully qualified registry address.
     * 
     */
    public Optional<Output<List<ExportEntryArgs>>> exports() {
        return Optional.ofNullable(this.exports);
    }

    /**
     * Name of the Dockerfile to use (defaults to &#34;${context}/Dockerfile&#34;).
     * 
     */
    @Import(name="file")
    private @Nullable Output<String> file;

    /**
     * @return
     * Name of the Dockerfile to use (defaults to &#34;${context}/Dockerfile&#34;).
     * 
     */
    public Optional<Output<String>> file() {
        return Optional.ofNullable(this.file);
    }

    /**
     * Set target platforms for the build. Defaults to the host&#39;s platform
     * 
     */
    @Import(name="platforms")
    private @Nullable Output<List<Platform>> platforms;

    /**
     * @return
     * Set target platforms for the build. Defaults to the host&#39;s platform
     * 
     */
    public Optional<Output<List<Platform>>> platforms() {
        return Optional.ofNullable(this.platforms);
    }

    /**
     * Always attempt to pull referenced images.
     * 
     */
    @Import(name="pull")
    private @Nullable Output<Boolean> pull;

    /**
     * @return
     * Always attempt to pull referenced images.
     * 
     */
    public Optional<Output<Boolean>> pull() {
        return Optional.ofNullable(this.pull);
    }

    /**
     * Logins for registry outputs
     * 
     */
    @Import(name="registries")
    private @Nullable Output<List<RegistryAuthArgs>> registries;

    /**
     * @return
     * Logins for registry outputs
     * 
     */
    public Optional<Output<List<RegistryAuthArgs>>> registries() {
        return Optional.ofNullable(this.registries);
    }

    /**
     * Name and optionally a tag (format: &#34;name:tag&#34;). If outputting to a
     * registry, the name should include the fully qualified registry address.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<List<String>> tags;

    /**
     * @return
     * Name and optionally a tag (format: &#34;name:tag&#34;). If outputting to a
     * registry, the name should include the fully qualified registry address.
     * 
     */
    public Optional<Output<List<String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    @Import(name="target")
    private @Nullable Output<String> target;

    public Optional<Output<String>> target() {
        return Optional.ofNullable(this.target);
    }

    private ImageArgs() {}

    private ImageArgs(ImageArgs $) {
        this.buildArgs = $.buildArgs;
        this.buildOnPreview = $.buildOnPreview;
        this.builder = $.builder;
        this.cacheFrom = $.cacheFrom;
        this.cacheTo = $.cacheTo;
        this.context = $.context;
        this.exports = $.exports;
        this.file = $.file;
        this.platforms = $.platforms;
        this.pull = $.pull;
        this.registries = $.registries;
        this.tags = $.tags;
        this.target = $.target;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ImageArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ImageArgs $;

        public Builder() {
            $ = new ImageArgs();
        }

        public Builder(ImageArgs defaults) {
            $ = new ImageArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param buildArgs
         * An optional map of named build-time argument variables to set during
         * the Docker build. This flag allows you to pass build-time variables that
         * can be accessed like environment variables inside the RUN
         * instruction.
         * 
         * @return builder
         * 
         */
        public Builder buildArgs(@Nullable Output<Map<String,String>> buildArgs) {
            $.buildArgs = buildArgs;
            return this;
        }

        /**
         * @param buildArgs
         * An optional map of named build-time argument variables to set during
         * the Docker build. This flag allows you to pass build-time variables that
         * can be accessed like environment variables inside the RUN
         * instruction.
         * 
         * @return builder
         * 
         */
        public Builder buildArgs(Map<String,String> buildArgs) {
            return buildArgs(Output.of(buildArgs));
        }

        /**
         * @param buildOnPreview
         * When true, attempt to build the image during previews. Outputs are not
         * pushed to registries, however caches are still populated.
         * 
         * @return builder
         * 
         */
        public Builder buildOnPreview(@Nullable Output<Boolean> buildOnPreview) {
            $.buildOnPreview = buildOnPreview;
            return this;
        }

        /**
         * @param buildOnPreview
         * When true, attempt to build the image during previews. Outputs are not
         * pushed to registries, however caches are still populated.
         * 
         * @return builder
         * 
         */
        public Builder buildOnPreview(Boolean buildOnPreview) {
            return buildOnPreview(Output.of(buildOnPreview));
        }

        /**
         * @param builder
         * Build with a specific builder instance
         * 
         * @return builder
         * 
         */
        public Builder builder_(@Nullable Output<String> builder) {
            $.builder = builder;
            return this;
        }

        /**
         * @param builder
         * Build with a specific builder instance
         * 
         * @return builder
         * 
         */
        public Builder builder_(String builder) {
            return builder_(Output.of(builder));
        }

        /**
         * @param cacheFrom
         * External cache sources (e.g., &#34;user/app:cache&#34;, &#34;type=local,src=path/to/dir&#34;)
         * 
         * @return builder
         * 
         */
        public Builder cacheFrom(@Nullable Output<List<CacheFromEntryArgs>> cacheFrom) {
            $.cacheFrom = cacheFrom;
            return this;
        }

        /**
         * @param cacheFrom
         * External cache sources (e.g., &#34;user/app:cache&#34;, &#34;type=local,src=path/to/dir&#34;)
         * 
         * @return builder
         * 
         */
        public Builder cacheFrom(List<CacheFromEntryArgs> cacheFrom) {
            return cacheFrom(Output.of(cacheFrom));
        }

        /**
         * @param cacheFrom
         * External cache sources (e.g., &#34;user/app:cache&#34;, &#34;type=local,src=path/to/dir&#34;)
         * 
         * @return builder
         * 
         */
        public Builder cacheFrom(CacheFromEntryArgs... cacheFrom) {
            return cacheFrom(List.of(cacheFrom));
        }

        /**
         * @param cacheTo
         * Cache export destinations (e.g., &#34;user/app:cache&#34;, &#34;type=local,dest=path/to/dir&#34;)
         * 
         * @return builder
         * 
         */
        public Builder cacheTo(@Nullable Output<List<CacheToEntryArgs>> cacheTo) {
            $.cacheTo = cacheTo;
            return this;
        }

        /**
         * @param cacheTo
         * Cache export destinations (e.g., &#34;user/app:cache&#34;, &#34;type=local,dest=path/to/dir&#34;)
         * 
         * @return builder
         * 
         */
        public Builder cacheTo(List<CacheToEntryArgs> cacheTo) {
            return cacheTo(Output.of(cacheTo));
        }

        /**
         * @param cacheTo
         * Cache export destinations (e.g., &#34;user/app:cache&#34;, &#34;type=local,dest=path/to/dir&#34;)
         * 
         * @return builder
         * 
         */
        public Builder cacheTo(CacheToEntryArgs... cacheTo) {
            return cacheTo(List.of(cacheTo));
        }

        /**
         * @param context
         * Path to use for build context. If omitted, an empty context is used.
         * 
         * @return builder
         * 
         */
        public Builder context(@Nullable Output<String> context) {
            $.context = context;
            return this;
        }

        /**
         * @param context
         * Path to use for build context. If omitted, an empty context is used.
         * 
         * @return builder
         * 
         */
        public Builder context(String context) {
            return context(Output.of(context));
        }

        /**
         * @param exports
         * Name and optionally a tag (format: &#34;name:tag&#34;). If outputting to a
         * registry, the name should include the fully qualified registry address.
         * 
         * @return builder
         * 
         */
        public Builder exports(@Nullable Output<List<ExportEntryArgs>> exports) {
            $.exports = exports;
            return this;
        }

        /**
         * @param exports
         * Name and optionally a tag (format: &#34;name:tag&#34;). If outputting to a
         * registry, the name should include the fully qualified registry address.
         * 
         * @return builder
         * 
         */
        public Builder exports(List<ExportEntryArgs> exports) {
            return exports(Output.of(exports));
        }

        /**
         * @param exports
         * Name and optionally a tag (format: &#34;name:tag&#34;). If outputting to a
         * registry, the name should include the fully qualified registry address.
         * 
         * @return builder
         * 
         */
        public Builder exports(ExportEntryArgs... exports) {
            return exports(List.of(exports));
        }

        /**
         * @param file
         * Name of the Dockerfile to use (defaults to &#34;${context}/Dockerfile&#34;).
         * 
         * @return builder
         * 
         */
        public Builder file(@Nullable Output<String> file) {
            $.file = file;
            return this;
        }

        /**
         * @param file
         * Name of the Dockerfile to use (defaults to &#34;${context}/Dockerfile&#34;).
         * 
         * @return builder
         * 
         */
        public Builder file(String file) {
            return file(Output.of(file));
        }

        /**
         * @param platforms
         * Set target platforms for the build. Defaults to the host&#39;s platform
         * 
         * @return builder
         * 
         */
        public Builder platforms(@Nullable Output<List<Platform>> platforms) {
            $.platforms = platforms;
            return this;
        }

        /**
         * @param platforms
         * Set target platforms for the build. Defaults to the host&#39;s platform
         * 
         * @return builder
         * 
         */
        public Builder platforms(List<Platform> platforms) {
            return platforms(Output.of(platforms));
        }

        /**
         * @param platforms
         * Set target platforms for the build. Defaults to the host&#39;s platform
         * 
         * @return builder
         * 
         */
        public Builder platforms(Platform... platforms) {
            return platforms(List.of(platforms));
        }

        /**
         * @param pull
         * Always attempt to pull referenced images.
         * 
         * @return builder
         * 
         */
        public Builder pull(@Nullable Output<Boolean> pull) {
            $.pull = pull;
            return this;
        }

        /**
         * @param pull
         * Always attempt to pull referenced images.
         * 
         * @return builder
         * 
         */
        public Builder pull(Boolean pull) {
            return pull(Output.of(pull));
        }

        /**
         * @param registries
         * Logins for registry outputs
         * 
         * @return builder
         * 
         */
        public Builder registries(@Nullable Output<List<RegistryAuthArgs>> registries) {
            $.registries = registries;
            return this;
        }

        /**
         * @param registries
         * Logins for registry outputs
         * 
         * @return builder
         * 
         */
        public Builder registries(List<RegistryAuthArgs> registries) {
            return registries(Output.of(registries));
        }

        /**
         * @param registries
         * Logins for registry outputs
         * 
         * @return builder
         * 
         */
        public Builder registries(RegistryAuthArgs... registries) {
            return registries(List.of(registries));
        }

        /**
         * @param tags
         * Name and optionally a tag (format: &#34;name:tag&#34;). If outputting to a
         * registry, the name should include the fully qualified registry address.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<List<String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags
         * Name and optionally a tag (format: &#34;name:tag&#34;). If outputting to a
         * registry, the name should include the fully qualified registry address.
         * 
         * @return builder
         * 
         */
        public Builder tags(List<String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tags
         * Name and optionally a tag (format: &#34;name:tag&#34;). If outputting to a
         * registry, the name should include the fully qualified registry address.
         * 
         * @return builder
         * 
         */
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }

        public Builder target(@Nullable Output<String> target) {
            $.target = target;
            return this;
        }

        public Builder target(String target) {
            return target(Output.of(target));
        }

        public ImageArgs build() {
            $.file = Codegen.stringProp("file").output().arg($.file).def("Dockerfile").getNullable();
            return $;
        }
    }

}
