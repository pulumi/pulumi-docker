// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker.buildx.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.docker.buildx.outputs.CacheFromAzureBlob;
import com.pulumi.docker.buildx.outputs.CacheFromGitHubActions;
import com.pulumi.docker.buildx.outputs.CacheFromLocal;
import com.pulumi.docker.buildx.outputs.CacheFromRegistry;
import com.pulumi.docker.buildx.outputs.CacheFromS3;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class CacheFromEntry {
    /**
     * @return
     * Push cache to Azure&#39;s blob storage service.
     * 
     */
    private @Nullable CacheFromAzureBlob azblob;
    /**
     * @return
     * When &#34;true&#34; this entry will be excluded. Defaults to &#34;false&#34;.
     * 
     */
    private @Nullable Boolean disabled;
    /**
     * @return
     * Recommended for use with GitHub Actions workflows.
     * 
     * An action like &#34;crazy-max/ghaction-github-runtime&#34; is recommended to
     * expose appropriate credentials to your GitHub workflow.
     * 
     */
    private @Nullable CacheFromGitHubActions gha;
    /**
     * @return
     * A simple backend which caches imagines on your local filesystem.
     * 
     */
    private @Nullable CacheFromLocal local;
    /**
     * @return
     * A raw string as you would provide it to the Docker CLI (e.g.,
     * &#34;type=inline&#34;)
     * 
     */
    private @Nullable String raw;
    /**
     * @return
     * Push caches to remote registries. Incompatible with the &#34;docker&#34; build
     * driver.
     * 
     */
    private @Nullable CacheFromRegistry registry;
    /**
     * @return
     * Push cache to AWS S3 or S3-compatible services such as MinIO.
     * 
     */
    private @Nullable CacheFromS3 s3;

    private CacheFromEntry() {}
    /**
     * @return
     * Push cache to Azure&#39;s blob storage service.
     * 
     */
    public Optional<CacheFromAzureBlob> azblob() {
        return Optional.ofNullable(this.azblob);
    }
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
     * Recommended for use with GitHub Actions workflows.
     * 
     * An action like &#34;crazy-max/ghaction-github-runtime&#34; is recommended to
     * expose appropriate credentials to your GitHub workflow.
     * 
     */
    public Optional<CacheFromGitHubActions> gha() {
        return Optional.ofNullable(this.gha);
    }
    /**
     * @return
     * A simple backend which caches imagines on your local filesystem.
     * 
     */
    public Optional<CacheFromLocal> local() {
        return Optional.ofNullable(this.local);
    }
    /**
     * @return
     * A raw string as you would provide it to the Docker CLI (e.g.,
     * &#34;type=inline&#34;)
     * 
     */
    public Optional<String> raw() {
        return Optional.ofNullable(this.raw);
    }
    /**
     * @return
     * Push caches to remote registries. Incompatible with the &#34;docker&#34; build
     * driver.
     * 
     */
    public Optional<CacheFromRegistry> registry() {
        return Optional.ofNullable(this.registry);
    }
    /**
     * @return
     * Push cache to AWS S3 or S3-compatible services such as MinIO.
     * 
     */
    public Optional<CacheFromS3> s3() {
        return Optional.ofNullable(this.s3);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(CacheFromEntry defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable CacheFromAzureBlob azblob;
        private @Nullable Boolean disabled;
        private @Nullable CacheFromGitHubActions gha;
        private @Nullable CacheFromLocal local;
        private @Nullable String raw;
        private @Nullable CacheFromRegistry registry;
        private @Nullable CacheFromS3 s3;
        public Builder() {}
        public Builder(CacheFromEntry defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.azblob = defaults.azblob;
    	      this.disabled = defaults.disabled;
    	      this.gha = defaults.gha;
    	      this.local = defaults.local;
    	      this.raw = defaults.raw;
    	      this.registry = defaults.registry;
    	      this.s3 = defaults.s3;
        }

        @CustomType.Setter
        public Builder azblob(@Nullable CacheFromAzureBlob azblob) {

            this.azblob = azblob;
            return this;
        }
        @CustomType.Setter
        public Builder disabled(@Nullable Boolean disabled) {

            this.disabled = disabled;
            return this;
        }
        @CustomType.Setter
        public Builder gha(@Nullable CacheFromGitHubActions gha) {

            this.gha = gha;
            return this;
        }
        @CustomType.Setter
        public Builder local(@Nullable CacheFromLocal local) {

            this.local = local;
            return this;
        }
        @CustomType.Setter
        public Builder raw(@Nullable String raw) {

            this.raw = raw;
            return this;
        }
        @CustomType.Setter
        public Builder registry(@Nullable CacheFromRegistry registry) {

            this.registry = registry;
            return this;
        }
        @CustomType.Setter
        public Builder s3(@Nullable CacheFromS3 s3) {

            this.s3 = s3;
            return this;
        }
        public CacheFromEntry build() {
            final var _resultValue = new CacheFromEntry();
            _resultValue.azblob = azblob;
            _resultValue.disabled = disabled;
            _resultValue.gha = gha;
            _resultValue.local = local;
            _resultValue.raw = raw;
            _resultValue.registry = registry;
            _resultValue.s3 = s3;
            return _resultValue;
        }
    }
}
