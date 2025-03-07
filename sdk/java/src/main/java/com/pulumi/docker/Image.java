// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker;

import com.pulumi.core.Alias;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.docker.ImageArgs;
import com.pulumi.docker.Utilities;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * `Image` builds a Docker image and pushes it Docker and OCI compatible registries.
 * This resource enables running Docker builds as part of a Pulumi deployment.
 * 
 * Note: We recommend you migrate your images to the more modern [Docker
 * Build](https://www.pulumi.com/registry/packages/docker-build/) provider to get
 * the best possible support, features, and performance.
 * 
 * Note: This resource does not delete tags, locally or remotely, when destroyed.
 * 
 * ## Image name
 * 
 * The Image resource uses `imageName` to refer to a fully qualified Docker image name, by the format `repository:tag`.
 * Note that this does not include any digest information and thus will not cause any updates when passed to dependencies,
 * even when using `latest` tag. To trigger such updates, e.g. when referencing pushed images in container orchestration
 * and management resources, please use the `repoDigest` Output instead, which is of the format
 * `repository{@literal @}&lt;algorithm&gt;:&lt;hash&gt;` and unique per build/push.
 * As of Docker v4.4, `repoDigest` is now available for local Images.
 * 
 * ## Cross-platform builds
 * 
 * The Image resource supports cross-platform builds when the [Docker engine has cross-platform support enabled via emulators](https://docs.docker.com/build/building/multi-platform/#building-multi-platform-images).
 * The Image resource currently supports providing only a single operating system and architecture in the `platform` field, e.g.: `linux/amd64`.
 * To enable this support, you may need to install the emulators in the environment running your Pulumi program.
 * 
 * If you are using Linux, you may be using Docker Engine or Docker Desktop for Linux, depending on how you have installed Docker. The [FAQ for Docker Desktop for Linux](https://docs.docker.com/desktop/faqs/linuxfaqs/#context) describes the differences and how to select which Docker context is in use.
 * 
 * * For local development using Docker Desktop, this is enabled by default.
 * * For systems using Docker Engine, install the QEMU binaries and register them with using the docker image from [github.com/tonistiigi/binfmt](https://github.com/tonistiigi/binfmt):
 * * In a GitHub Actions workflow, the [docker/setup-qemu-action](https://github.com/docker/setup-qemu-action) can be used instead by adding this step to your workflow file. Example workflow usage:
 * 
 * ## Example Usage
 * ### A Docker image build
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.docker.Image;
 * import com.pulumi.docker.ImageArgs;
 * import com.pulumi.docker.inputs.DockerBuildArgs;
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
 *         var demoImage = new Image("demoImage", ImageArgs.builder()
 *             .build(DockerBuildArgs.builder()
 *                 .context(".")
 *                 .dockerfile("Dockerfile")
 *                 .platform("linux/amd64")
 *                 .build())
 *             .imageName("username/image:tag1")
 *             .skipPush(true)
 *             .build());
 * 
 *         ctx.export("imageName", demoImage.imageName());
 *     }
 * }
 * }
 * </pre>
 * ### A Docker image build and push
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.docker.Image;
 * import com.pulumi.docker.ImageArgs;
 * import com.pulumi.docker.inputs.DockerBuildArgs;
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
 *         var demoPushImage = new Image("demoPushImage", ImageArgs.builder()
 *             .build(DockerBuildArgs.builder()
 *                 .context(".")
 *                 .dockerfile("Dockerfile")
 *                 .build())
 *             .imageName("docker.io/username/push-image:tag1")
 *             .build());
 * 
 *         ctx.export("imageName", demoPushImage.imageName());
 *         ctx.export("repoDigest", demoPushImage.repoDigest());
 *     }
 * }
 * }
 * </pre>
 * ### Docker image build using caching with AWS Elastic Container Registry
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.ecr.Repository;
 * import com.pulumi.aws.ecr.RepositoryArgs;
 * import com.pulumi.aws.ecr.EcrFunctions;
 * import com.pulumi.aws.ecr.inputs.GetAuthorizationTokenArgs;
 * import com.pulumi.docker.Image;
 * import com.pulumi.docker.ImageArgs;
 * import com.pulumi.docker.inputs.DockerBuildArgs;
 * import com.pulumi.docker.inputs.CacheFromArgs;
 * import com.pulumi.docker.inputs.RegistryArgs;
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
 *         var ecrRepository = new Repository("ecrRepository", RepositoryArgs.builder()
 *             .name("docker-repository")
 *             .build());
 * 
 *         final var authToken = EcrFunctions.getAuthorizationToken(GetAuthorizationTokenArgs.builder()
 *             .registryId(ecrRepository.registryId())
 *             .build());
 * 
 *         var myAppImage = new Image("myAppImage", ImageArgs.builder()
 *             .build(DockerBuildArgs.builder()
 *                 .args(Map.of("BUILDKIT_INLINE_CACHE", "1"))
 *                 .cacheFrom(CacheFromArgs.builder()
 *                     .images(ecrRepository.repositoryUrl().applyValue(_repositoryUrl -> String.format("%s:latest", _repositoryUrl)))
 *                     .build())
 *                 .context("app/")
 *                 .dockerfile("app/Dockerfile")
 *                 .build())
 *             .imageName(ecrRepository.repositoryUrl().applyValue(_repositoryUrl -> String.format("%s:latest", _repositoryUrl)))
 *             .registry(RegistryArgs.builder()
 *                 .password(Output.ofSecret(authToken.applyValue(_authToken -> _authToken.password())))
 *                 .server(ecrRepository.repositoryUrl())
 *                 .username(authToken.applyValue(_authToken -> _authToken.userName()))
 *                 .build())
 *             .build());
 * 
 *         ctx.export("imageName", myAppImage.imageName());
 *     }
 * }
 * }
 * </pre>
 * 
 */
@ResourceType(type="docker:index/image:Image")
public class Image extends com.pulumi.resources.CustomResource {
    /**
     * The fully qualified image name that was pushed to the registry.
     * 
     */
    @Export(name="baseImageName", refs={String.class}, tree="[0]")
    private Output<String> baseImageName;

    /**
     * @return The fully qualified image name that was pushed to the registry.
     * 
     */
    public Output<String> baseImageName() {
        return this.baseImageName;
    }
    /**
     * The path to the build context to use.
     * 
     */
    @Export(name="context", refs={String.class}, tree="[0]")
    private Output<String> context;

    /**
     * @return The path to the build context to use.
     * 
     */
    public Output<String> context() {
        return this.context;
    }
    /**
     * The location of the Dockerfile relative to the docker build context.
     * 
     */
    @Export(name="dockerfile", refs={String.class}, tree="[0]")
    private Output<String> dockerfile;

    /**
     * @return The location of the Dockerfile relative to the docker build context.
     * 
     */
    public Output<String> dockerfile() {
        return this.dockerfile;
    }
    /**
     * The fully qualified image name
     * 
     */
    @Export(name="imageName", refs={String.class}, tree="[0]")
    private Output<String> imageName;

    /**
     * @return The fully qualified image name
     * 
     */
    public Output<String> imageName() {
        return this.imageName;
    }
    /**
     * The image&#39;s architecture and OS
     * 
     */
    @Export(name="platform", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> platform;

    /**
     * @return The image&#39;s architecture and OS
     * 
     */
    public Output<Optional<String>> platform() {
        return Codegen.optional(this.platform);
    }
    /**
     * The name of the registry server hosting the image.
     * 
     */
    @Export(name="registryServer", refs={String.class}, tree="[0]")
    private Output<String> registryServer;

    /**
     * @return The name of the registry server hosting the image.
     * 
     */
    public Output<String> registryServer() {
        return this.registryServer;
    }
    /**
     * **For pushed images:**
     * The manifest digest of an image pushed to a registry, of the format repository{@literal @}&lt;algorithm&gt;:&lt;hash&gt;, e.g. `username/demo-image{@literal @}sha256:a6ae6dd8d39c5bb02320e41abf00cd4cb35905fec540e37d306c878be8d38bd3`.
     * This reference is unique per image build and push.
     * Only available for images pushed to a registry.
     * Use when passing a reference to a pushed image to container management resources.
     * 
     * **Local-only images**For local images, this field is the image ID of the built local image, of the format &lt;algorithm&gt;:&lt;hash&gt;, e.g `sha256:826a130323165bb0ccb0374ae774f885c067a951b51a6ee133577f4e5dbc4119`
     * 
     */
    @Export(name="repoDigest", refs={String.class}, tree="[0]")
    private Output<String> repoDigest;

    /**
     * @return **For pushed images:**
     * The manifest digest of an image pushed to a registry, of the format repository{@literal @}&lt;algorithm&gt;:&lt;hash&gt;, e.g. `username/demo-image{@literal @}sha256:a6ae6dd8d39c5bb02320e41abf00cd4cb35905fec540e37d306c878be8d38bd3`.
     * This reference is unique per image build and push.
     * Only available for images pushed to a registry.
     * Use when passing a reference to a pushed image to container management resources.
     * 
     * **Local-only images**For local images, this field is the image ID of the built local image, of the format &lt;algorithm&gt;:&lt;hash&gt;, e.g `sha256:826a130323165bb0ccb0374ae774f885c067a951b51a6ee133577f4e5dbc4119`
     * 
     */
    public Output<String> repoDigest() {
        return this.repoDigest;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Image(java.lang.String name) {
        this(name, ImageArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Image(java.lang.String name, ImageArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Image(java.lang.String name, ImageArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("docker:index/image:Image", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Image(java.lang.String name, Output<java.lang.String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("docker:index/image:Image", name, null, makeResourceOptions(options, id), false);
    }

    private static ImageArgs makeArgs(ImageArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ImageArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .aliases(List.of(
                Output.of(Alias.builder().type("docker:image:Image").build())
            ))
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
    public static Image get(java.lang.String name, Output<java.lang.String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Image(name, id, options);
    }
}
