// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker.buildx;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.docker.Utilities;
import com.pulumi.docker.buildx.ImageArgs;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * A Docker image built using Buildkit
 * 
 */
@ResourceType(type="docker:buildx/image:Image")
public class Image extends com.pulumi.resources.CustomResource {
    @Export(name="architecture", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> architecture;

    public Output<Optional<String>> architecture() {
        return Codegen.optional(this.architecture);
    }
    /**
     * Contexts to use while building the image. If omitted, an empty context
     * is used. If more than one value is specified, they should be of the
     * form &#34;name=value&#34;.
     * 
     */
    @Export(name="context", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> context;

    /**
     * @return
     * Contexts to use while building the image. If omitted, an empty context
     * is used. If more than one value is specified, they should be of the
     * form &#34;name=value&#34;.
     * 
     */
    public Output<Optional<List<String>>> context() {
        return Codegen.optional(this.context);
    }
    /**
     * Name and optionally a tag (format: &#34;name:tag&#34;). If outputting to a
     * registry, the name should include the fully qualified registry address.
     * 
     */
    @Export(name="exports", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> exports;

    /**
     * @return
     * Name and optionally a tag (format: &#34;name:tag&#34;). If outputting to a
     * registry, the name should include the fully qualified registry address.
     * 
     */
    public Output<Optional<List<String>>> exports() {
        return Codegen.optional(this.exports);
    }
    /**
     * Name of the Dockerfile to use (default: &#34;$PATH/Dockerfile&#34;).
     * 
     */
    @Export(name="file", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> file;

    /**
     * @return
     * Name of the Dockerfile to use (default: &#34;$PATH/Dockerfile&#34;).
     * 
     */
    public Output<Optional<String>> file() {
        return Codegen.optional(this.file);
    }
    @Export(name="os", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> os;

    public Output<Optional<String>> os() {
        return Codegen.optional(this.os);
    }
    @Export(name="repoDigests", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> repoDigests;

    public Output<Optional<List<String>>> repoDigests() {
        return Codegen.optional(this.repoDigests);
    }
    @Export(name="repoTags", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> repoTags;

    public Output<Optional<List<String>>> repoTags() {
        return Codegen.optional(this.repoTags);
    }
    @Export(name="size", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> size;

    public Output<Optional<Integer>> size() {
        return Codegen.optional(this.size);
    }
    /**
     * Name and optionally a tag (format: &#34;name:tag&#34;). If outputting to a
     * registry, the name should include the fully qualified registry address.
     * 
     */
    @Export(name="tags", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> tags;

    /**
     * @return
     * Name and optionally a tag (format: &#34;name:tag&#34;). If outputting to a
     * registry, the name should include the fully qualified registry address.
     * 
     */
    public Output<List<String>> tags() {
        return this.tags;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Image(String name) {
        this(name, ImageArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Image(String name, ImageArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Image(String name, ImageArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("docker:buildx/image:Image", name, args == null ? ImageArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Image(String name, Output<String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("docker:buildx/image:Image", name, null, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static Image get(String name, Output<String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Image(name, id, options);
    }
}