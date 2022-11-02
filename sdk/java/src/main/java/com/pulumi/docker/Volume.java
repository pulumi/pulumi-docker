// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.docker.Utilities;
import com.pulumi.docker.VolumeArgs;
import com.pulumi.docker.inputs.VolumeState;
import com.pulumi.docker.outputs.VolumeLabel;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * &lt;!-- Bug: Type and Name are switched --&gt;
 * Creates and destroys a volume in Docker. This can be used alongside docker.Container to prepare volumes that can be shared across containers.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.docker.Volume;
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
 *         var sharedVolume = new Volume(&#34;sharedVolume&#34;);
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * ### Example Assuming you created a `volume` as follows #!/bin/bash docker volume create prints the long ID 524b0457aa2a87dd2b75c74c3e4e53f406974249e63ab3ed9bf21e5644f9dc7d you provide the definition for the resource as follows terraform resource &#34;docker_volume&#34; &#34;foo&#34; {
 * 
 *  name = &#34;524b0457aa2a87dd2b75c74c3e4e53f406974249e63ab3ed9bf21e5644f9dc7d&#34; } then the import command is as follows #!/bin/bash
 * 
 * ```sh
 *  $ pulumi import docker:index/volume:Volume foo 524b0457aa2a87dd2b75c74c3e4e53f406974249e63ab3ed9bf21e5644f9dc7d
 * ```
 * 
 */
@ResourceType(type="docker:index/volume:Volume")
public class Volume extends com.pulumi.resources.CustomResource {
    /**
     * Driver type for the volume. Defaults to `local`.
     * 
     */
    @Export(name="driver", type=String.class, parameters={})
    private Output<String> driver;

    /**
     * @return Driver type for the volume. Defaults to `local`.
     * 
     */
    public Output<String> driver() {
        return this.driver;
    }
    /**
     * Options specific to the driver.
     * 
     */
    @Export(name="driverOpts", type=Map.class, parameters={String.class, Object.class})
    private Output</* @Nullable */ Map<String,Object>> driverOpts;

    /**
     * @return Options specific to the driver.
     * 
     */
    public Output<Optional<Map<String,Object>>> driverOpts() {
        return Codegen.optional(this.driverOpts);
    }
    /**
     * User-defined key/value metadata
     * 
     */
    @Export(name="labels", type=List.class, parameters={VolumeLabel.class})
    private Output</* @Nullable */ List<VolumeLabel>> labels;

    /**
     * @return User-defined key/value metadata
     * 
     */
    public Output<Optional<List<VolumeLabel>>> labels() {
        return Codegen.optional(this.labels);
    }
    /**
     * The mountpoint of the volume.
     * 
     */
    @Export(name="mountpoint", type=String.class, parameters={})
    private Output<String> mountpoint;

    /**
     * @return The mountpoint of the volume.
     * 
     */
    public Output<String> mountpoint() {
        return this.mountpoint;
    }
    /**
     * The name of the Docker volume (will be generated if not provided).
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return The name of the Docker volume (will be generated if not provided).
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Volume(String name) {
        this(name, VolumeArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Volume(String name, @Nullable VolumeArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Volume(String name, @Nullable VolumeArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("docker:index/volume:Volume", name, args == null ? VolumeArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Volume(String name, Output<String> id, @Nullable VolumeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("docker:index/volume:Volume", name, state, makeResourceOptions(options, id));
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
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static Volume get(String name, Output<String> id, @Nullable VolumeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Volume(name, id, state, options);
    }
}
