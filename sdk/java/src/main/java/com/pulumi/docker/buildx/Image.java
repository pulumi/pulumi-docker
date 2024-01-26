// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker.buildx;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.docker.Utilities;
import com.pulumi.docker.buildx.ImageArgs;
import com.pulumi.docker.buildx.outputs.Manifest;
import com.pulumi.docker.buildx.outputs.RegistryAuth;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * A Docker image built using Buildkit
 * 
 */
@ResourceType(type="docker:buildx/image:Image")
public class Image extends com.pulumi.resources.CustomResource {
    /**
     * An optional map of named build-time argument variables to set during
     * the Docker build. This flag allows you to pass build-time variables that
     * can be accessed like environment variables inside the RUN
     * instruction.
     * 
     */
    @Export(name="buildArgs", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> buildArgs;

    /**
     * @return
     * An optional map of named build-time argument variables to set during
     * the Docker build. This flag allows you to pass build-time variables that
     * can be accessed like environment variables inside the RUN
     * instruction.
     * 
     */
    public Output<Optional<Map<String,String>>> buildArgs() {
        return Codegen.optional(this.buildArgs);
    }
    /**
     * Build with a specific builder instance
     * 
     */
    @Export(name="builder", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> builder;

    /**
     * @return
     * Build with a specific builder instance
     * 
     */
    public Output<Optional<String>> builder_() {
        return Codegen.optional(this.builder);
    }
    /**
     * External cache sources (e.g., &#34;user/app:cache&#34;, &#34;type=local,src=path/to/dir&#34;)
     * 
     */
    @Export(name="cacheFrom", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> cacheFrom;

    /**
     * @return
     * External cache sources (e.g., &#34;user/app:cache&#34;, &#34;type=local,src=path/to/dir&#34;)
     * 
     */
    public Output<Optional<List<String>>> cacheFrom() {
        return Codegen.optional(this.cacheFrom);
    }
    /**
     * Cache export destinations (e.g., &#34;user/app:cache&#34;, &#34;type=local,dest=path/to/dir&#34;)
     * 
     */
    @Export(name="cacheTo", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> cacheTo;

    /**
     * @return
     * Cache export destinations (e.g., &#34;user/app:cache&#34;, &#34;type=local,dest=path/to/dir&#34;)
     * 
     */
    public Output<Optional<List<String>>> cacheTo() {
        return Codegen.optional(this.cacheTo);
    }
    /**
     * Path to use for build context. If omitted, an empty context is used.
     * 
     */
    @Export(name="context", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> context;

    /**
     * @return
     * Path to use for build context. If omitted, an empty context is used.
     * 
     */
    public Output<Optional<String>> context() {
        return Codegen.optional(this.context);
    }
    @Export(name="contextHash", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> contextHash;

    public Output<Optional<String>> contextHash() {
        return Codegen.optional(this.contextHash);
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
     * Name of the Dockerfile to use (defaults to &#34;${context}/Dockerfile&#34;).
     * 
     */
    @Export(name="file", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> file;

    /**
     * @return
     * Name of the Dockerfile to use (defaults to &#34;${context}/Dockerfile&#34;).
     * 
     */
    public Output<Optional<String>> file() {
        return Codegen.optional(this.file);
    }
    @Export(name="manifests", refs={List.class,Manifest.class}, tree="[0,1]")
    private Output<List<Manifest>> manifests;

    public Output<List<Manifest>> manifests() {
        return this.manifests;
    }
    /**
     * Set target platforms for the build. Defaults to the host&#39;s platform
     * 
     */
    @Export(name="platforms", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> platforms;

    /**
     * @return
     * Set target platforms for the build. Defaults to the host&#39;s platform
     * 
     */
    public Output<Optional<List<String>>> platforms() {
        return Codegen.optional(this.platforms);
    }
    /**
     * Always attempt to pull all referenced images
     * 
     */
    @Export(name="pull", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> pull;

    /**
     * @return
     * Always attempt to pull all referenced images
     * 
     */
    public Output<Optional<Boolean>> pull() {
        return Codegen.optional(this.pull);
    }
    /**
     * Logins for registry outputs
     * 
     */
    @Export(name="registries", refs={List.class,RegistryAuth.class}, tree="[0,1]")
    private Output</* @Nullable */ List<RegistryAuth>> registries;

    /**
     * @return
     * Logins for registry outputs
     * 
     */
    public Output<Optional<List<RegistryAuth>>> registries() {
        return Codegen.optional(this.registries);
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
