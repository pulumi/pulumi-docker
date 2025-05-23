// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.docker.inputs.RemoteImageBuildArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RemoteImageState extends com.pulumi.resources.ResourceArgs {

    public static final RemoteImageState Empty = new RemoteImageState();

    @Import(name="build")
    private @Nullable Output<RemoteImageBuildArgs> build;

    public Optional<Output<RemoteImageBuildArgs>> build() {
        return Optional.ofNullable(this.build);
    }

    /**
     * If true, then the image is removed forcibly when the resource is destroyed.
     * 
     */
    @Import(name="forceRemove")
    private @Nullable Output<Boolean> forceRemove;

    /**
     * @return If true, then the image is removed forcibly when the resource is destroyed.
     * 
     */
    public Optional<Output<Boolean>> forceRemove() {
        return Optional.ofNullable(this.forceRemove);
    }

    /**
     * The ID of the image (as seen when executing `docker inspect` on the image). Can be used to reference the image via its ID in other resources.
     * 
     */
    @Import(name="imageId")
    private @Nullable Output<String> imageId;

    /**
     * @return The ID of the image (as seen when executing `docker inspect` on the image). Can be used to reference the image via its ID in other resources.
     * 
     */
    public Optional<Output<String>> imageId() {
        return Optional.ofNullable(this.imageId);
    }

    /**
     * If true, then the Docker image won&#39;t be deleted on destroy operation. If this is false, it will delete the image from the docker local storage on destroy operation.
     * 
     */
    @Import(name="keepLocally")
    private @Nullable Output<Boolean> keepLocally;

    /**
     * @return If true, then the Docker image won&#39;t be deleted on destroy operation. If this is false, it will delete the image from the docker local storage on destroy operation.
     * 
     */
    public Optional<Output<Boolean>> keepLocally() {
        return Optional.ofNullable(this.keepLocally);
    }

    /**
     * The name of the Docker image, including any tags or SHA256 repo digests.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the Docker image, including any tags or SHA256 repo digests.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The platform to use when pulling the image. Defaults to the platform of the current machine.
     * 
     */
    @Import(name="platform")
    private @Nullable Output<String> platform;

    /**
     * @return The platform to use when pulling the image. Defaults to the platform of the current machine.
     * 
     */
    public Optional<Output<String>> platform() {
        return Optional.ofNullable(this.platform);
    }

    /**
     * List of values which cause an image pull when changed. This is used to store the image digest from the registry when using the docker*registry*image.
     * 
     */
    @Import(name="pullTriggers")
    private @Nullable Output<List<String>> pullTriggers;

    /**
     * @return List of values which cause an image pull when changed. This is used to store the image digest from the registry when using the docker*registry*image.
     * 
     */
    public Optional<Output<List<String>>> pullTriggers() {
        return Optional.ofNullable(this.pullTriggers);
    }

    /**
     * The image sha256 digest in the form of `repo[:tag]{@literal @}sha256:&lt;hash&gt;`. This may not be populated when building an image, because it is read from the local Docker client and so may be available only when the image was either pulled from the repo or pushed to the repo (perhaps using `docker.RegistryImage`) in a previous run.
     * 
     */
    @Import(name="repoDigest")
    private @Nullable Output<String> repoDigest;

    /**
     * @return The image sha256 digest in the form of `repo[:tag]{@literal @}sha256:&lt;hash&gt;`. This may not be populated when building an image, because it is read from the local Docker client and so may be available only when the image was either pulled from the repo or pushed to the repo (perhaps using `docker.RegistryImage`) in a previous run.
     * 
     */
    public Optional<Output<String>> repoDigest() {
        return Optional.ofNullable(this.repoDigest);
    }

    /**
     * A map of arbitrary strings that, when changed, will force the `docker.RemoteImage` resource to be replaced. This can be used to rebuild an image when contents of source code folders change
     * 
     */
    @Import(name="triggers")
    private @Nullable Output<Map<String,String>> triggers;

    /**
     * @return A map of arbitrary strings that, when changed, will force the `docker.RemoteImage` resource to be replaced. This can be used to rebuild an image when contents of source code folders change
     * 
     */
    public Optional<Output<Map<String,String>>> triggers() {
        return Optional.ofNullable(this.triggers);
    }

    private RemoteImageState() {}

    private RemoteImageState(RemoteImageState $) {
        this.build = $.build;
        this.forceRemove = $.forceRemove;
        this.imageId = $.imageId;
        this.keepLocally = $.keepLocally;
        this.name = $.name;
        this.platform = $.platform;
        this.pullTriggers = $.pullTriggers;
        this.repoDigest = $.repoDigest;
        this.triggers = $.triggers;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RemoteImageState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RemoteImageState $;

        public Builder() {
            $ = new RemoteImageState();
        }

        public Builder(RemoteImageState defaults) {
            $ = new RemoteImageState(Objects.requireNonNull(defaults));
        }

        public Builder build(@Nullable Output<RemoteImageBuildArgs> build) {
            $.build = build;
            return this;
        }

        public Builder build(RemoteImageBuildArgs build) {
            return build(Output.of(build));
        }

        /**
         * @param forceRemove If true, then the image is removed forcibly when the resource is destroyed.
         * 
         * @return builder
         * 
         */
        public Builder forceRemove(@Nullable Output<Boolean> forceRemove) {
            $.forceRemove = forceRemove;
            return this;
        }

        /**
         * @param forceRemove If true, then the image is removed forcibly when the resource is destroyed.
         * 
         * @return builder
         * 
         */
        public Builder forceRemove(Boolean forceRemove) {
            return forceRemove(Output.of(forceRemove));
        }

        /**
         * @param imageId The ID of the image (as seen when executing `docker inspect` on the image). Can be used to reference the image via its ID in other resources.
         * 
         * @return builder
         * 
         */
        public Builder imageId(@Nullable Output<String> imageId) {
            $.imageId = imageId;
            return this;
        }

        /**
         * @param imageId The ID of the image (as seen when executing `docker inspect` on the image). Can be used to reference the image via its ID in other resources.
         * 
         * @return builder
         * 
         */
        public Builder imageId(String imageId) {
            return imageId(Output.of(imageId));
        }

        /**
         * @param keepLocally If true, then the Docker image won&#39;t be deleted on destroy operation. If this is false, it will delete the image from the docker local storage on destroy operation.
         * 
         * @return builder
         * 
         */
        public Builder keepLocally(@Nullable Output<Boolean> keepLocally) {
            $.keepLocally = keepLocally;
            return this;
        }

        /**
         * @param keepLocally If true, then the Docker image won&#39;t be deleted on destroy operation. If this is false, it will delete the image from the docker local storage on destroy operation.
         * 
         * @return builder
         * 
         */
        public Builder keepLocally(Boolean keepLocally) {
            return keepLocally(Output.of(keepLocally));
        }

        /**
         * @param name The name of the Docker image, including any tags or SHA256 repo digests.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the Docker image, including any tags or SHA256 repo digests.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param platform The platform to use when pulling the image. Defaults to the platform of the current machine.
         * 
         * @return builder
         * 
         */
        public Builder platform(@Nullable Output<String> platform) {
            $.platform = platform;
            return this;
        }

        /**
         * @param platform The platform to use when pulling the image. Defaults to the platform of the current machine.
         * 
         * @return builder
         * 
         */
        public Builder platform(String platform) {
            return platform(Output.of(platform));
        }

        /**
         * @param pullTriggers List of values which cause an image pull when changed. This is used to store the image digest from the registry when using the docker*registry*image.
         * 
         * @return builder
         * 
         */
        public Builder pullTriggers(@Nullable Output<List<String>> pullTriggers) {
            $.pullTriggers = pullTriggers;
            return this;
        }

        /**
         * @param pullTriggers List of values which cause an image pull when changed. This is used to store the image digest from the registry when using the docker*registry*image.
         * 
         * @return builder
         * 
         */
        public Builder pullTriggers(List<String> pullTriggers) {
            return pullTriggers(Output.of(pullTriggers));
        }

        /**
         * @param pullTriggers List of values which cause an image pull when changed. This is used to store the image digest from the registry when using the docker*registry*image.
         * 
         * @return builder
         * 
         */
        public Builder pullTriggers(String... pullTriggers) {
            return pullTriggers(List.of(pullTriggers));
        }

        /**
         * @param repoDigest The image sha256 digest in the form of `repo[:tag]{@literal @}sha256:&lt;hash&gt;`. This may not be populated when building an image, because it is read from the local Docker client and so may be available only when the image was either pulled from the repo or pushed to the repo (perhaps using `docker.RegistryImage`) in a previous run.
         * 
         * @return builder
         * 
         */
        public Builder repoDigest(@Nullable Output<String> repoDigest) {
            $.repoDigest = repoDigest;
            return this;
        }

        /**
         * @param repoDigest The image sha256 digest in the form of `repo[:tag]{@literal @}sha256:&lt;hash&gt;`. This may not be populated when building an image, because it is read from the local Docker client and so may be available only when the image was either pulled from the repo or pushed to the repo (perhaps using `docker.RegistryImage`) in a previous run.
         * 
         * @return builder
         * 
         */
        public Builder repoDigest(String repoDigest) {
            return repoDigest(Output.of(repoDigest));
        }

        /**
         * @param triggers A map of arbitrary strings that, when changed, will force the `docker.RemoteImage` resource to be replaced. This can be used to rebuild an image when contents of source code folders change
         * 
         * @return builder
         * 
         */
        public Builder triggers(@Nullable Output<Map<String,String>> triggers) {
            $.triggers = triggers;
            return this;
        }

        /**
         * @param triggers A map of arbitrary strings that, when changed, will force the `docker.RemoteImage` resource to be replaced. This can be used to rebuild an image when contents of source code folders change
         * 
         * @return builder
         * 
         */
        public Builder triggers(Map<String,String> triggers) {
            return triggers(Output.of(triggers));
        }

        public RemoteImageState build() {
            return $;
        }
    }

}
