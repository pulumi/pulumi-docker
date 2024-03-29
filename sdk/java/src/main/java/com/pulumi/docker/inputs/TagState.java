// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class TagState extends com.pulumi.resources.ResourceArgs {

    public static final TagState Empty = new TagState();

    /**
     * Name of the source image.
     * 
     */
    @Import(name="sourceImage")
    private @Nullable Output<String> sourceImage;

    /**
     * @return Name of the source image.
     * 
     */
    public Optional<Output<String>> sourceImage() {
        return Optional.ofNullable(this.sourceImage);
    }

    /**
     * ImageID of the source image in the format of `sha256:&lt;&lt;ID&gt;&gt;`
     * 
     */
    @Import(name="sourceImageId")
    private @Nullable Output<String> sourceImageId;

    /**
     * @return ImageID of the source image in the format of `sha256:&lt;&lt;ID&gt;&gt;`
     * 
     */
    public Optional<Output<String>> sourceImageId() {
        return Optional.ofNullable(this.sourceImageId);
    }

    /**
     * Name of the target image.
     * 
     */
    @Import(name="targetImage")
    private @Nullable Output<String> targetImage;

    /**
     * @return Name of the target image.
     * 
     */
    public Optional<Output<String>> targetImage() {
        return Optional.ofNullable(this.targetImage);
    }

    private TagState() {}

    private TagState(TagState $) {
        this.sourceImage = $.sourceImage;
        this.sourceImageId = $.sourceImageId;
        this.targetImage = $.targetImage;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TagState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TagState $;

        public Builder() {
            $ = new TagState();
        }

        public Builder(TagState defaults) {
            $ = new TagState(Objects.requireNonNull(defaults));
        }

        /**
         * @param sourceImage Name of the source image.
         * 
         * @return builder
         * 
         */
        public Builder sourceImage(@Nullable Output<String> sourceImage) {
            $.sourceImage = sourceImage;
            return this;
        }

        /**
         * @param sourceImage Name of the source image.
         * 
         * @return builder
         * 
         */
        public Builder sourceImage(String sourceImage) {
            return sourceImage(Output.of(sourceImage));
        }

        /**
         * @param sourceImageId ImageID of the source image in the format of `sha256:&lt;&lt;ID&gt;&gt;`
         * 
         * @return builder
         * 
         */
        public Builder sourceImageId(@Nullable Output<String> sourceImageId) {
            $.sourceImageId = sourceImageId;
            return this;
        }

        /**
         * @param sourceImageId ImageID of the source image in the format of `sha256:&lt;&lt;ID&gt;&gt;`
         * 
         * @return builder
         * 
         */
        public Builder sourceImageId(String sourceImageId) {
            return sourceImageId(Output.of(sourceImageId));
        }

        /**
         * @param targetImage Name of the target image.
         * 
         * @return builder
         * 
         */
        public Builder targetImage(@Nullable Output<String> targetImage) {
            $.targetImage = targetImage;
            return this;
        }

        /**
         * @param targetImage Name of the target image.
         * 
         * @return builder
         * 
         */
        public Builder targetImage(String targetImage) {
            return targetImage(Output.of(targetImage));
        }

        public TagState build() {
            return $;
        }
    }

}
