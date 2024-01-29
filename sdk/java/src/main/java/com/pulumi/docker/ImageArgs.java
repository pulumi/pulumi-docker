// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.core.internal.Codegen;
import com.pulumi.docker.inputs.DockerBuildArgs;
import com.pulumi.docker.inputs.RegistryArgs;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ImageArgs extends com.pulumi.resources.ResourceArgs {

    public static final ImageArgs Empty = new ImageArgs();

    /**
     * The Docker build context
     * 
     */
    @Import(name="build")
    private @Nullable Output<DockerBuildArgs> build;

    /**
     * @return The Docker build context
     * 
     */
    public Optional<Output<DockerBuildArgs>> build() {
        return Optional.ofNullable(this.build);
    }

    /**
     * A flag to build an image on preview
     * 
     */
    @Import(name="buildOnPreview")
    private @Nullable Output<Boolean> buildOnPreview;

    /**
     * @return A flag to build an image on preview
     * 
     */
    public Optional<Output<Boolean>> buildOnPreview() {
        return Optional.ofNullable(this.buildOnPreview);
    }

    /**
     * The image name, of the format repository[:tag], e.g. `docker.io/username/demo-image:v1`.
     * This reference is not unique to each build and push.For the unique manifest SHA of a pushed docker image, or the local image ID, please use `repoDigest`.
     * 
     */
    @Import(name="imageName", required=true)
    private Output<String> imageName;

    /**
     * @return The image name, of the format repository[:tag], e.g. `docker.io/username/demo-image:v1`.
     * This reference is not unique to each build and push.For the unique manifest SHA of a pushed docker image, or the local image ID, please use `repoDigest`.
     * 
     */
    public Output<String> imageName() {
        return this.imageName;
    }

    /**
     * The registry to push the image to
     * 
     */
    @Import(name="registry")
    private @Nullable Output<RegistryArgs> registry;

    /**
     * @return The registry to push the image to
     * 
     */
    public Optional<Output<RegistryArgs>> registry() {
        return Optional.ofNullable(this.registry);
    }

    /**
     * A flag to skip a registry push.
     * 
     */
    @Import(name="skipPush")
    private @Nullable Output<Boolean> skipPush;

    /**
     * @return A flag to skip a registry push.
     * 
     */
    public Optional<Output<Boolean>> skipPush() {
        return Optional.ofNullable(this.skipPush);
    }

    private ImageArgs() {}

    private ImageArgs(ImageArgs $) {
        this.build = $.build;
        this.buildOnPreview = $.buildOnPreview;
        this.imageName = $.imageName;
        this.registry = $.registry;
        this.skipPush = $.skipPush;
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
         * @param build The Docker build context
         * 
         * @return builder
         * 
         */
        public Builder build(@Nullable Output<DockerBuildArgs> build) {
            $.build = build;
            return this;
        }

        /**
         * @param build The Docker build context
         * 
         * @return builder
         * 
         */
        public Builder build(DockerBuildArgs build) {
            return build(Output.of(build));
        }

        /**
         * @param buildOnPreview A flag to build an image on preview
         * 
         * @return builder
         * 
         */
        public Builder buildOnPreview(@Nullable Output<Boolean> buildOnPreview) {
            $.buildOnPreview = buildOnPreview;
            return this;
        }

        /**
         * @param buildOnPreview A flag to build an image on preview
         * 
         * @return builder
         * 
         */
        public Builder buildOnPreview(Boolean buildOnPreview) {
            return buildOnPreview(Output.of(buildOnPreview));
        }

        /**
         * @param imageName The image name, of the format repository[:tag], e.g. `docker.io/username/demo-image:v1`.
         * This reference is not unique to each build and push.For the unique manifest SHA of a pushed docker image, or the local image ID, please use `repoDigest`.
         * 
         * @return builder
         * 
         */
        public Builder imageName(Output<String> imageName) {
            $.imageName = imageName;
            return this;
        }

        /**
         * @param imageName The image name, of the format repository[:tag], e.g. `docker.io/username/demo-image:v1`.
         * This reference is not unique to each build and push.For the unique manifest SHA of a pushed docker image, or the local image ID, please use `repoDigest`.
         * 
         * @return builder
         * 
         */
        public Builder imageName(String imageName) {
            return imageName(Output.of(imageName));
        }

        /**
         * @param registry The registry to push the image to
         * 
         * @return builder
         * 
         */
        public Builder registry(@Nullable Output<RegistryArgs> registry) {
            $.registry = registry;
            return this;
        }

        /**
         * @param registry The registry to push the image to
         * 
         * @return builder
         * 
         */
        public Builder registry(RegistryArgs registry) {
            return registry(Output.of(registry));
        }

        /**
         * @param skipPush A flag to skip a registry push.
         * 
         * @return builder
         * 
         */
        public Builder skipPush(@Nullable Output<Boolean> skipPush) {
            $.skipPush = skipPush;
            return this;
        }

        /**
         * @param skipPush A flag to skip a registry push.
         * 
         * @return builder
         * 
         */
        public Builder skipPush(Boolean skipPush) {
            return skipPush(Output.of(skipPush));
        }

        public ImageArgs build() {
            $.buildOnPreview = Codegen.booleanProp("buildOnPreview").output().arg($.buildOnPreview).def(false).getNullable();
            if ($.imageName == null) {
                throw new MissingRequiredPropertyException("ImageArgs", "imageName");
            }
            $.skipPush = Codegen.booleanProp("skipPush").output().arg($.skipPush).def(false).getNullable();
            return $;
        }
    }

}
