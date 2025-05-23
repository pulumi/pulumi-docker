// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.docker.inputs.RegistryImageAuthConfigArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RegistryImageState extends com.pulumi.resources.ResourceArgs {

    public static final RegistryImageState Empty = new RegistryImageState();

    /**
     * Authentication configuration for the Docker registry. It is only used for this resource.
     * 
     */
    @Import(name="authConfig")
    private @Nullable Output<RegistryImageAuthConfigArgs> authConfig;

    /**
     * @return Authentication configuration for the Docker registry. It is only used for this resource.
     * 
     */
    public Optional<Output<RegistryImageAuthConfigArgs>> authConfig() {
        return Optional.ofNullable(this.authConfig);
    }

    /**
     * If `true`, the verification of TLS certificates of the server/registry is disabled. Defaults to `false`
     * 
     */
    @Import(name="insecureSkipVerify")
    private @Nullable Output<Boolean> insecureSkipVerify;

    /**
     * @return If `true`, the verification of TLS certificates of the server/registry is disabled. Defaults to `false`
     * 
     */
    public Optional<Output<Boolean>> insecureSkipVerify() {
        return Optional.ofNullable(this.insecureSkipVerify);
    }

    /**
     * If true, then the Docker image won&#39;t be deleted on destroy operation. If this is false, it will delete the image from the docker registry on destroy operation. Defaults to `false`
     * 
     */
    @Import(name="keepRemotely")
    private @Nullable Output<Boolean> keepRemotely;

    /**
     * @return If true, then the Docker image won&#39;t be deleted on destroy operation. If this is false, it will delete the image from the docker registry on destroy operation. Defaults to `false`
     * 
     */
    public Optional<Output<Boolean>> keepRemotely() {
        return Optional.ofNullable(this.keepRemotely);
    }

    /**
     * The name of the Docker image.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the Docker image.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The sha256 digest of the image.
     * 
     */
    @Import(name="sha256Digest")
    private @Nullable Output<String> sha256Digest;

    /**
     * @return The sha256 digest of the image.
     * 
     */
    public Optional<Output<String>> sha256Digest() {
        return Optional.ofNullable(this.sha256Digest);
    }

    /**
     * A map of arbitrary strings that, when changed, will force the `docker.RegistryImage` resource to be replaced. This can be used to repush a local image
     * 
     */
    @Import(name="triggers")
    private @Nullable Output<Map<String,String>> triggers;

    /**
     * @return A map of arbitrary strings that, when changed, will force the `docker.RegistryImage` resource to be replaced. This can be used to repush a local image
     * 
     */
    public Optional<Output<Map<String,String>>> triggers() {
        return Optional.ofNullable(this.triggers);
    }

    private RegistryImageState() {}

    private RegistryImageState(RegistryImageState $) {
        this.authConfig = $.authConfig;
        this.insecureSkipVerify = $.insecureSkipVerify;
        this.keepRemotely = $.keepRemotely;
        this.name = $.name;
        this.sha256Digest = $.sha256Digest;
        this.triggers = $.triggers;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RegistryImageState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RegistryImageState $;

        public Builder() {
            $ = new RegistryImageState();
        }

        public Builder(RegistryImageState defaults) {
            $ = new RegistryImageState(Objects.requireNonNull(defaults));
        }

        /**
         * @param authConfig Authentication configuration for the Docker registry. It is only used for this resource.
         * 
         * @return builder
         * 
         */
        public Builder authConfig(@Nullable Output<RegistryImageAuthConfigArgs> authConfig) {
            $.authConfig = authConfig;
            return this;
        }

        /**
         * @param authConfig Authentication configuration for the Docker registry. It is only used for this resource.
         * 
         * @return builder
         * 
         */
        public Builder authConfig(RegistryImageAuthConfigArgs authConfig) {
            return authConfig(Output.of(authConfig));
        }

        /**
         * @param insecureSkipVerify If `true`, the verification of TLS certificates of the server/registry is disabled. Defaults to `false`
         * 
         * @return builder
         * 
         */
        public Builder insecureSkipVerify(@Nullable Output<Boolean> insecureSkipVerify) {
            $.insecureSkipVerify = insecureSkipVerify;
            return this;
        }

        /**
         * @param insecureSkipVerify If `true`, the verification of TLS certificates of the server/registry is disabled. Defaults to `false`
         * 
         * @return builder
         * 
         */
        public Builder insecureSkipVerify(Boolean insecureSkipVerify) {
            return insecureSkipVerify(Output.of(insecureSkipVerify));
        }

        /**
         * @param keepRemotely If true, then the Docker image won&#39;t be deleted on destroy operation. If this is false, it will delete the image from the docker registry on destroy operation. Defaults to `false`
         * 
         * @return builder
         * 
         */
        public Builder keepRemotely(@Nullable Output<Boolean> keepRemotely) {
            $.keepRemotely = keepRemotely;
            return this;
        }

        /**
         * @param keepRemotely If true, then the Docker image won&#39;t be deleted on destroy operation. If this is false, it will delete the image from the docker registry on destroy operation. Defaults to `false`
         * 
         * @return builder
         * 
         */
        public Builder keepRemotely(Boolean keepRemotely) {
            return keepRemotely(Output.of(keepRemotely));
        }

        /**
         * @param name The name of the Docker image.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the Docker image.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param sha256Digest The sha256 digest of the image.
         * 
         * @return builder
         * 
         */
        public Builder sha256Digest(@Nullable Output<String> sha256Digest) {
            $.sha256Digest = sha256Digest;
            return this;
        }

        /**
         * @param sha256Digest The sha256 digest of the image.
         * 
         * @return builder
         * 
         */
        public Builder sha256Digest(String sha256Digest) {
            return sha256Digest(Output.of(sha256Digest));
        }

        /**
         * @param triggers A map of arbitrary strings that, when changed, will force the `docker.RegistryImage` resource to be replaced. This can be used to repush a local image
         * 
         * @return builder
         * 
         */
        public Builder triggers(@Nullable Output<Map<String,String>> triggers) {
            $.triggers = triggers;
            return this;
        }

        /**
         * @param triggers A map of arbitrary strings that, when changed, will force the `docker.RegistryImage` resource to be replaced. This can be used to repush a local image
         * 
         * @return builder
         * 
         */
        public Builder triggers(Map<String,String> triggers) {
            return triggers(Output.of(triggers));
        }

        public RegistryImageState build() {
            return $;
        }
    }

}
