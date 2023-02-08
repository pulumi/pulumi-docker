// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.docker.ImageArgs;
import com.pulumi.docker.Utilities;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Builds a Docker Image and pushes to a Docker registry.
 * 
 * ## Example Usage
 * ### A Docker image build
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.docker.Image;
 * import com.pulumi.docker.ImageArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var demoImage = new Image(&#34;demoImage&#34;, ImageArgs.builder()        
 *             .build(%!v(PANIC=Format method: interface conversion: model.Expression is *model.TemplateExpression, not *model.LiteralValueExpression))
 *             .imageName(&#34;username/image:tag1&#34;)
 *             .skipPush(true)
 *             .build());
 * 
 *         ctx.export(&#34;imageName&#34;, demoImage.imageName());
 *     }
 * }
 * ```
 * 
 * {{% //examples %}}
 * 
 */
@ResourceType(type="docker:index/image:Image")
public class Image extends com.pulumi.resources.CustomResource {
    /**
     * The fully qualified image name that was pushed to the registry.
     * 
     */
    @Export(name="baseImageName", type=String.class, parameters={})
    private Output</* @Nullable */ String> baseImageName;

    /**
     * @return The fully qualified image name that was pushed to the registry.
     * 
     */
    public Output<Optional<String>> baseImageName() {
        return Codegen.optional(this.baseImageName);
    }
    /**
     * The fully qualified image name
     * 
     */
    @Export(name="imageName", type=String.class, parameters={})
    private Output</* @Nullable */ String> imageName;

    /**
     * @return The fully qualified image name
     * 
     */
    public Output<Optional<String>> imageName() {
        return Codegen.optional(this.imageName);
    }
    /**
     * The name of the registry server hosting the image.
     * 
     */
    @Export(name="registryServer", type=String.class, parameters={})
    private Output</* @Nullable */ String> registryServer;

    /**
     * @return The name of the registry server hosting the image.
     * 
     */
    public Output<Optional<String>> registryServer() {
        return Codegen.optional(this.registryServer);
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
        super("docker:index/image:Image", name, args == null ? ImageArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Image(String name, Output<String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("docker:index/image:Image", name, null, makeResourceOptions(options, id));
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
