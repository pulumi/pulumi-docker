// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker.buildx.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.docker.buildx.inputs.CacheFromAzureBlobArgs;
import com.pulumi.docker.buildx.inputs.CacheFromGitHubActionsArgs;
import com.pulumi.docker.buildx.inputs.CacheFromLocalArgs;
import com.pulumi.docker.buildx.inputs.CacheFromRegistryArgs;
import com.pulumi.docker.buildx.inputs.CacheFromS3Args;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class CacheFromEntryArgs extends com.pulumi.resources.ResourceArgs {

    public static final CacheFromEntryArgs Empty = new CacheFromEntryArgs();

    /**
     * Push cache to Azure&#39;s blob storage service.
     * 
     */
    @Import(name="azblob")
    private @Nullable Output<CacheFromAzureBlobArgs> azblob;

    /**
     * @return
     * Push cache to Azure&#39;s blob storage service.
     * 
     */
    public Optional<Output<CacheFromAzureBlobArgs>> azblob() {
        return Optional.ofNullable(this.azblob);
    }

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
     * Recommended for use with GitHub Actions workflows.
     * 
     * An action like &#34;crazy-max/ghaction-github-runtime&#34; is recommended to
     * expose appropriate credentials to your GitHub workflow.
     * 
     */
    @Import(name="gha")
    private @Nullable Output<CacheFromGitHubActionsArgs> gha;

    /**
     * @return
     * Recommended for use with GitHub Actions workflows.
     * 
     * An action like &#34;crazy-max/ghaction-github-runtime&#34; is recommended to
     * expose appropriate credentials to your GitHub workflow.
     * 
     */
    public Optional<Output<CacheFromGitHubActionsArgs>> gha() {
        return Optional.ofNullable(this.gha);
    }

    /**
     * A simple backend which caches imagines on your local filesystem.
     * 
     */
    @Import(name="local")
    private @Nullable Output<CacheFromLocalArgs> local;

    /**
     * @return
     * A simple backend which caches imagines on your local filesystem.
     * 
     */
    public Optional<Output<CacheFromLocalArgs>> local() {
        return Optional.ofNullable(this.local);
    }

    /**
     * A raw string as you would provide it to the Docker CLI (e.g.,
     * &#34;type=inline&#34;)
     * 
     */
    @Import(name="raw")
    private @Nullable Output<String> raw;

    /**
     * @return
     * A raw string as you would provide it to the Docker CLI (e.g.,
     * &#34;type=inline&#34;)
     * 
     */
    public Optional<Output<String>> raw() {
        return Optional.ofNullable(this.raw);
    }

    /**
     * Push caches to remote registries. Incompatible with the &#34;docker&#34; build
     * driver.
     * 
     */
    @Import(name="registry")
    private @Nullable Output<CacheFromRegistryArgs> registry;

    /**
     * @return
     * Push caches to remote registries. Incompatible with the &#34;docker&#34; build
     * driver.
     * 
     */
    public Optional<Output<CacheFromRegistryArgs>> registry() {
        return Optional.ofNullable(this.registry);
    }

    /**
     * Push cache to AWS S3 or S3-compatible services such as MinIO.
     * 
     */
    @Import(name="s3")
    private @Nullable Output<CacheFromS3Args> s3;

    /**
     * @return
     * Push cache to AWS S3 or S3-compatible services such as MinIO.
     * 
     */
    public Optional<Output<CacheFromS3Args>> s3() {
        return Optional.ofNullable(this.s3);
    }

    private CacheFromEntryArgs() {}

    private CacheFromEntryArgs(CacheFromEntryArgs $) {
        this.azblob = $.azblob;
        this.disabled = $.disabled;
        this.gha = $.gha;
        this.local = $.local;
        this.raw = $.raw;
        this.registry = $.registry;
        this.s3 = $.s3;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CacheFromEntryArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CacheFromEntryArgs $;

        public Builder() {
            $ = new CacheFromEntryArgs();
        }

        public Builder(CacheFromEntryArgs defaults) {
            $ = new CacheFromEntryArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param azblob
         * Push cache to Azure&#39;s blob storage service.
         * 
         * @return builder
         * 
         */
        public Builder azblob(@Nullable Output<CacheFromAzureBlobArgs> azblob) {
            $.azblob = azblob;
            return this;
        }

        /**
         * @param azblob
         * Push cache to Azure&#39;s blob storage service.
         * 
         * @return builder
         * 
         */
        public Builder azblob(CacheFromAzureBlobArgs azblob) {
            return azblob(Output.of(azblob));
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
         * @param gha
         * Recommended for use with GitHub Actions workflows.
         * 
         * An action like &#34;crazy-max/ghaction-github-runtime&#34; is recommended to
         * expose appropriate credentials to your GitHub workflow.
         * 
         * @return builder
         * 
         */
        public Builder gha(@Nullable Output<CacheFromGitHubActionsArgs> gha) {
            $.gha = gha;
            return this;
        }

        /**
         * @param gha
         * Recommended for use with GitHub Actions workflows.
         * 
         * An action like &#34;crazy-max/ghaction-github-runtime&#34; is recommended to
         * expose appropriate credentials to your GitHub workflow.
         * 
         * @return builder
         * 
         */
        public Builder gha(CacheFromGitHubActionsArgs gha) {
            return gha(Output.of(gha));
        }

        /**
         * @param local
         * A simple backend which caches imagines on your local filesystem.
         * 
         * @return builder
         * 
         */
        public Builder local(@Nullable Output<CacheFromLocalArgs> local) {
            $.local = local;
            return this;
        }

        /**
         * @param local
         * A simple backend which caches imagines on your local filesystem.
         * 
         * @return builder
         * 
         */
        public Builder local(CacheFromLocalArgs local) {
            return local(Output.of(local));
        }

        /**
         * @param raw
         * A raw string as you would provide it to the Docker CLI (e.g.,
         * &#34;type=inline&#34;)
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
         * &#34;type=inline&#34;)
         * 
         * @return builder
         * 
         */
        public Builder raw(String raw) {
            return raw(Output.of(raw));
        }

        /**
         * @param registry
         * Push caches to remote registries. Incompatible with the &#34;docker&#34; build
         * driver.
         * 
         * @return builder
         * 
         */
        public Builder registry(@Nullable Output<CacheFromRegistryArgs> registry) {
            $.registry = registry;
            return this;
        }

        /**
         * @param registry
         * Push caches to remote registries. Incompatible with the &#34;docker&#34; build
         * driver.
         * 
         * @return builder
         * 
         */
        public Builder registry(CacheFromRegistryArgs registry) {
            return registry(Output.of(registry));
        }

        /**
         * @param s3
         * Push cache to AWS S3 or S3-compatible services such as MinIO.
         * 
         * @return builder
         * 
         */
        public Builder s3(@Nullable Output<CacheFromS3Args> s3) {
            $.s3 = s3;
            return this;
        }

        /**
         * @param s3
         * Push cache to AWS S3 or S3-compatible services such as MinIO.
         * 
         * @return builder
         * 
         */
        public Builder s3(CacheFromS3Args s3) {
            return s3(Output.of(s3));
        }

        public CacheFromEntryArgs build() {
            return $;
        }
    }

}
