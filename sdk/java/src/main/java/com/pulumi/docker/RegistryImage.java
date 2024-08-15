// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.docker.RegistryImageArgs;
import com.pulumi.docker.Utilities;
import com.pulumi.docker.inputs.RegistryImageState;
import java.lang.Boolean;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * &lt;!-- Bug: Type and Name are switched --&gt;
 * Manages the lifecycle of docker image in a registry. You can upload images to a registry (= `docker push`) and also delete them again
 * 
 */
@ResourceType(type="docker:index/registryImage:RegistryImage")
public class RegistryImage extends com.pulumi.resources.CustomResource {
    /**
     * If `true`, the verification of TLS certificates of the server/registry is disabled. Defaults to `false`
     * 
     */
    @Export(name="insecureSkipVerify", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> insecureSkipVerify;

    /**
     * @return If `true`, the verification of TLS certificates of the server/registry is disabled. Defaults to `false`
     * 
     */
    public Output<Optional<Boolean>> insecureSkipVerify() {
        return Codegen.optional(this.insecureSkipVerify);
    }
    /**
     * If true, then the Docker image won&#39;t be deleted on destroy operation. If this is false, it will delete the image from the docker registry on destroy operation. Defaults to `false`
     * 
     */
    @Export(name="keepRemotely", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> keepRemotely;

    /**
     * @return If true, then the Docker image won&#39;t be deleted on destroy operation. If this is false, it will delete the image from the docker registry on destroy operation. Defaults to `false`
     * 
     */
    public Output<Optional<Boolean>> keepRemotely() {
        return Codegen.optional(this.keepRemotely);
    }
    /**
     * The name of the Docker image.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the Docker image.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The sha256 digest of the image.
     * 
     */
    @Export(name="sha256Digest", refs={String.class}, tree="[0]")
    private Output<String> sha256Digest;

    /**
     * @return The sha256 digest of the image.
     * 
     */
    public Output<String> sha256Digest() {
        return this.sha256Digest;
    }
    /**
     * A map of arbitrary strings that, when changed, will force the `docker.RegistryImage` resource to be replaced. This can be used to repush a local image
     * 
     */
    @Export(name="triggers", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> triggers;

    /**
     * @return A map of arbitrary strings that, when changed, will force the `docker.RegistryImage` resource to be replaced. This can be used to repush a local image
     * 
     */
    public Output<Optional<Map<String,String>>> triggers() {
        return Codegen.optional(this.triggers);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RegistryImage(java.lang.String name) {
        this(name, RegistryImageArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RegistryImage(java.lang.String name, @Nullable RegistryImageArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RegistryImage(java.lang.String name, @Nullable RegistryImageArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("docker:index/registryImage:RegistryImage", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private RegistryImage(java.lang.String name, Output<java.lang.String> id, @Nullable RegistryImageState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("docker:index/registryImage:RegistryImage", name, state, makeResourceOptions(options, id), false);
    }

    private static RegistryImageArgs makeArgs(@Nullable RegistryImageArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? RegistryImageArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
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
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static RegistryImage get(java.lang.String name, Output<java.lang.String> id, @Nullable RegistryImageState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RegistryImage(name, id, state, options);
    }
}
