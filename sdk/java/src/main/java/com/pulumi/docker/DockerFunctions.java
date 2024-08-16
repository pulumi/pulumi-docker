// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker;

import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import com.pulumi.docker.Utilities;
import com.pulumi.docker.inputs.GetLogsArgs;
import com.pulumi.docker.inputs.GetLogsPlainArgs;
import com.pulumi.docker.inputs.GetNetworkArgs;
import com.pulumi.docker.inputs.GetNetworkPlainArgs;
import com.pulumi.docker.inputs.GetPluginArgs;
import com.pulumi.docker.inputs.GetPluginPlainArgs;
import com.pulumi.docker.inputs.GetRegistryImageArgs;
import com.pulumi.docker.inputs.GetRegistryImagePlainArgs;
import com.pulumi.docker.inputs.GetRemoteImageArgs;
import com.pulumi.docker.inputs.GetRemoteImagePlainArgs;
import com.pulumi.docker.outputs.GetLogsResult;
import com.pulumi.docker.outputs.GetNetworkResult;
import com.pulumi.docker.outputs.GetPluginResult;
import com.pulumi.docker.outputs.GetRegistryImageResult;
import com.pulumi.docker.outputs.GetRemoteImageResult;
import java.util.concurrent.CompletableFuture;

public final class DockerFunctions {
    /**
     * `docker.getLogs` provides logs from specific container
     * 
     */
    public static Output<GetLogsResult> getLogs(GetLogsArgs args) {
        return getLogs(args, InvokeOptions.Empty);
    }
    /**
     * `docker.getLogs` provides logs from specific container
     * 
     */
    public static CompletableFuture<GetLogsResult> getLogsPlain(GetLogsPlainArgs args) {
        return getLogsPlain(args, InvokeOptions.Empty);
    }
    /**
     * `docker.getLogs` provides logs from specific container
     * 
     */
    public static Output<GetLogsResult> getLogs(GetLogsArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("docker:index/getLogs:getLogs", TypeShape.of(GetLogsResult.class), args, Utilities.withVersion(options));
    }
    /**
     * `docker.getLogs` provides logs from specific container
     * 
     */
    public static CompletableFuture<GetLogsResult> getLogsPlain(GetLogsPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("docker:index/getLogs:getLogs", TypeShape.of(GetLogsResult.class), args, Utilities.withVersion(options));
    }
    /**
     * `docker.Network` provides details about a specific Docker Network.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.docker.DockerFunctions;
     * import com.pulumi.docker.inputs.GetNetworkArgs;
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
     *         final var main = DockerFunctions.getNetwork(GetNetworkArgs.builder()
     *             .name("main")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetNetworkResult> getNetwork(GetNetworkArgs args) {
        return getNetwork(args, InvokeOptions.Empty);
    }
    /**
     * `docker.Network` provides details about a specific Docker Network.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.docker.DockerFunctions;
     * import com.pulumi.docker.inputs.GetNetworkArgs;
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
     *         final var main = DockerFunctions.getNetwork(GetNetworkArgs.builder()
     *             .name("main")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetNetworkResult> getNetworkPlain(GetNetworkPlainArgs args) {
        return getNetworkPlain(args, InvokeOptions.Empty);
    }
    /**
     * `docker.Network` provides details about a specific Docker Network.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.docker.DockerFunctions;
     * import com.pulumi.docker.inputs.GetNetworkArgs;
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
     *         final var main = DockerFunctions.getNetwork(GetNetworkArgs.builder()
     *             .name("main")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetNetworkResult> getNetwork(GetNetworkArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("docker:index/getNetwork:getNetwork", TypeShape.of(GetNetworkResult.class), args, Utilities.withVersion(options));
    }
    /**
     * `docker.Network` provides details about a specific Docker Network.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.docker.DockerFunctions;
     * import com.pulumi.docker.inputs.GetNetworkArgs;
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
     *         final var main = DockerFunctions.getNetwork(GetNetworkArgs.builder()
     *             .name("main")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetNetworkResult> getNetworkPlain(GetNetworkPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("docker:index/getNetwork:getNetwork", TypeShape.of(GetNetworkResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Reads the local Docker plugin. The plugin must be installed locally.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.docker.DockerFunctions;
     * import com.pulumi.docker.inputs.GetPluginArgs;
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
     *         //## With alias
     *         final var byAlias = DockerFunctions.getPlugin(GetPluginArgs.builder()
     *             .alias("sample-volume-plugin:latest")
     *             .build());
     * 
     *         //## With ID
     *         final var byId = DockerFunctions.getPlugin(GetPluginArgs.builder()
     *             .id("e9a9db917b3bfd6706b5d3a66d4bceb9f")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetPluginResult> getPlugin() {
        return getPlugin(GetPluginArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Reads the local Docker plugin. The plugin must be installed locally.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.docker.DockerFunctions;
     * import com.pulumi.docker.inputs.GetPluginArgs;
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
     *         //## With alias
     *         final var byAlias = DockerFunctions.getPlugin(GetPluginArgs.builder()
     *             .alias("sample-volume-plugin:latest")
     *             .build());
     * 
     *         //## With ID
     *         final var byId = DockerFunctions.getPlugin(GetPluginArgs.builder()
     *             .id("e9a9db917b3bfd6706b5d3a66d4bceb9f")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetPluginResult> getPluginPlain() {
        return getPluginPlain(GetPluginPlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Reads the local Docker plugin. The plugin must be installed locally.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.docker.DockerFunctions;
     * import com.pulumi.docker.inputs.GetPluginArgs;
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
     *         //## With alias
     *         final var byAlias = DockerFunctions.getPlugin(GetPluginArgs.builder()
     *             .alias("sample-volume-plugin:latest")
     *             .build());
     * 
     *         //## With ID
     *         final var byId = DockerFunctions.getPlugin(GetPluginArgs.builder()
     *             .id("e9a9db917b3bfd6706b5d3a66d4bceb9f")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetPluginResult> getPlugin(GetPluginArgs args) {
        return getPlugin(args, InvokeOptions.Empty);
    }
    /**
     * Reads the local Docker plugin. The plugin must be installed locally.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.docker.DockerFunctions;
     * import com.pulumi.docker.inputs.GetPluginArgs;
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
     *         //## With alias
     *         final var byAlias = DockerFunctions.getPlugin(GetPluginArgs.builder()
     *             .alias("sample-volume-plugin:latest")
     *             .build());
     * 
     *         //## With ID
     *         final var byId = DockerFunctions.getPlugin(GetPluginArgs.builder()
     *             .id("e9a9db917b3bfd6706b5d3a66d4bceb9f")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetPluginResult> getPluginPlain(GetPluginPlainArgs args) {
        return getPluginPlain(args, InvokeOptions.Empty);
    }
    /**
     * Reads the local Docker plugin. The plugin must be installed locally.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.docker.DockerFunctions;
     * import com.pulumi.docker.inputs.GetPluginArgs;
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
     *         //## With alias
     *         final var byAlias = DockerFunctions.getPlugin(GetPluginArgs.builder()
     *             .alias("sample-volume-plugin:latest")
     *             .build());
     * 
     *         //## With ID
     *         final var byId = DockerFunctions.getPlugin(GetPluginArgs.builder()
     *             .id("e9a9db917b3bfd6706b5d3a66d4bceb9f")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetPluginResult> getPlugin(GetPluginArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("docker:index/getPlugin:getPlugin", TypeShape.of(GetPluginResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Reads the local Docker plugin. The plugin must be installed locally.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.docker.DockerFunctions;
     * import com.pulumi.docker.inputs.GetPluginArgs;
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
     *         //## With alias
     *         final var byAlias = DockerFunctions.getPlugin(GetPluginArgs.builder()
     *             .alias("sample-volume-plugin:latest")
     *             .build());
     * 
     *         //## With ID
     *         final var byId = DockerFunctions.getPlugin(GetPluginArgs.builder()
     *             .id("e9a9db917b3bfd6706b5d3a66d4bceb9f")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetPluginResult> getPluginPlain(GetPluginPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("docker:index/getPlugin:getPlugin", TypeShape.of(GetPluginResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Reads the image metadata from a Docker Registry. Used in conjunction with the docker.RemoteImage resource to keep an image up to date on the latest available version of the tag.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.docker.DockerFunctions;
     * import com.pulumi.docker.inputs.GetRegistryImageArgs;
     * import com.pulumi.docker.RemoteImage;
     * import com.pulumi.docker.RemoteImageArgs;
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
     *         final var ubuntu = DockerFunctions.getRegistryImage(GetRegistryImageArgs.builder()
     *             .name("ubuntu:precise")
     *             .build());
     * 
     *         var ubuntuRemoteImage = new RemoteImage("ubuntuRemoteImage", RemoteImageArgs.builder()
     *             .name(ubuntu.applyValue(getRegistryImageResult -> getRegistryImageResult.name()))
     *             .pullTriggers(ubuntu.applyValue(getRegistryImageResult -> getRegistryImageResult.sha256Digest()))
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetRegistryImageResult> getRegistryImage(GetRegistryImageArgs args) {
        return getRegistryImage(args, InvokeOptions.Empty);
    }
    /**
     * Reads the image metadata from a Docker Registry. Used in conjunction with the docker.RemoteImage resource to keep an image up to date on the latest available version of the tag.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.docker.DockerFunctions;
     * import com.pulumi.docker.inputs.GetRegistryImageArgs;
     * import com.pulumi.docker.RemoteImage;
     * import com.pulumi.docker.RemoteImageArgs;
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
     *         final var ubuntu = DockerFunctions.getRegistryImage(GetRegistryImageArgs.builder()
     *             .name("ubuntu:precise")
     *             .build());
     * 
     *         var ubuntuRemoteImage = new RemoteImage("ubuntuRemoteImage", RemoteImageArgs.builder()
     *             .name(ubuntu.applyValue(getRegistryImageResult -> getRegistryImageResult.name()))
     *             .pullTriggers(ubuntu.applyValue(getRegistryImageResult -> getRegistryImageResult.sha256Digest()))
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetRegistryImageResult> getRegistryImagePlain(GetRegistryImagePlainArgs args) {
        return getRegistryImagePlain(args, InvokeOptions.Empty);
    }
    /**
     * Reads the image metadata from a Docker Registry. Used in conjunction with the docker.RemoteImage resource to keep an image up to date on the latest available version of the tag.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.docker.DockerFunctions;
     * import com.pulumi.docker.inputs.GetRegistryImageArgs;
     * import com.pulumi.docker.RemoteImage;
     * import com.pulumi.docker.RemoteImageArgs;
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
     *         final var ubuntu = DockerFunctions.getRegistryImage(GetRegistryImageArgs.builder()
     *             .name("ubuntu:precise")
     *             .build());
     * 
     *         var ubuntuRemoteImage = new RemoteImage("ubuntuRemoteImage", RemoteImageArgs.builder()
     *             .name(ubuntu.applyValue(getRegistryImageResult -> getRegistryImageResult.name()))
     *             .pullTriggers(ubuntu.applyValue(getRegistryImageResult -> getRegistryImageResult.sha256Digest()))
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetRegistryImageResult> getRegistryImage(GetRegistryImageArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("docker:index/getRegistryImage:getRegistryImage", TypeShape.of(GetRegistryImageResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Reads the image metadata from a Docker Registry. Used in conjunction with the docker.RemoteImage resource to keep an image up to date on the latest available version of the tag.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.docker.DockerFunctions;
     * import com.pulumi.docker.inputs.GetRegistryImageArgs;
     * import com.pulumi.docker.RemoteImage;
     * import com.pulumi.docker.RemoteImageArgs;
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
     *         final var ubuntu = DockerFunctions.getRegistryImage(GetRegistryImageArgs.builder()
     *             .name("ubuntu:precise")
     *             .build());
     * 
     *         var ubuntuRemoteImage = new RemoteImage("ubuntuRemoteImage", RemoteImageArgs.builder()
     *             .name(ubuntu.applyValue(getRegistryImageResult -> getRegistryImageResult.name()))
     *             .pullTriggers(ubuntu.applyValue(getRegistryImageResult -> getRegistryImageResult.sha256Digest()))
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetRegistryImageResult> getRegistryImagePlain(GetRegistryImagePlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("docker:index/getRegistryImage:getRegistryImage", TypeShape.of(GetRegistryImageResult.class), args, Utilities.withVersion(options));
    }
    /**
     * `docker.RemoteImage` provides details about a specific Docker Image which needs to be present on the Docker Host
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.docker.DockerFunctions;
     * import com.pulumi.docker.inputs.GetRemoteImageArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App }{{@code
     *     public static void main(String[] args) }{{@code
     *         Pulumi.run(App::stack);
     *     }}{@code
     * 
     *     public static void stack(Context ctx) }{{@code
     *         // uses the 'latest' tag
     *         final var latest = DockerFunctions.getRemoteImage(GetRemoteImageArgs.builder()
     *             .name("nginx")
     *             .build());
     * 
     *         // uses a specific tag
     *         final var specific = DockerFunctions.getRemoteImage(GetRemoteImageArgs.builder()
     *             .name("nginx:1.17.6")
     *             .build());
     * 
     *         // use the image digest
     *         final var digest = DockerFunctions.getRemoteImage(GetRemoteImageArgs.builder()
     *             .name("nginx}{@literal @}{@code sha256:36b74457bccb56fbf8b05f79c85569501b721d4db813b684391d63e02287c0b2")
     *             .build());
     * 
     *         // uses the tag and the image digest
     *         final var tagAndDigest = DockerFunctions.getRemoteImage(GetRemoteImageArgs.builder()
     *             .name("nginx:1.19.1}{@literal @}{@code sha256:36b74457bccb56fbf8b05f79c85569501b721d4db813b684391d63e02287c0b2")
     *             .build());
     * 
     *     }}{@code
     * }}{@code
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetRemoteImageResult> getRemoteImage(GetRemoteImageArgs args) {
        return getRemoteImage(args, InvokeOptions.Empty);
    }
    /**
     * `docker.RemoteImage` provides details about a specific Docker Image which needs to be present on the Docker Host
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.docker.DockerFunctions;
     * import com.pulumi.docker.inputs.GetRemoteImageArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App }{{@code
     *     public static void main(String[] args) }{{@code
     *         Pulumi.run(App::stack);
     *     }}{@code
     * 
     *     public static void stack(Context ctx) }{{@code
     *         // uses the 'latest' tag
     *         final var latest = DockerFunctions.getRemoteImage(GetRemoteImageArgs.builder()
     *             .name("nginx")
     *             .build());
     * 
     *         // uses a specific tag
     *         final var specific = DockerFunctions.getRemoteImage(GetRemoteImageArgs.builder()
     *             .name("nginx:1.17.6")
     *             .build());
     * 
     *         // use the image digest
     *         final var digest = DockerFunctions.getRemoteImage(GetRemoteImageArgs.builder()
     *             .name("nginx}{@literal @}{@code sha256:36b74457bccb56fbf8b05f79c85569501b721d4db813b684391d63e02287c0b2")
     *             .build());
     * 
     *         // uses the tag and the image digest
     *         final var tagAndDigest = DockerFunctions.getRemoteImage(GetRemoteImageArgs.builder()
     *             .name("nginx:1.19.1}{@literal @}{@code sha256:36b74457bccb56fbf8b05f79c85569501b721d4db813b684391d63e02287c0b2")
     *             .build());
     * 
     *     }}{@code
     * }}{@code
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetRemoteImageResult> getRemoteImagePlain(GetRemoteImagePlainArgs args) {
        return getRemoteImagePlain(args, InvokeOptions.Empty);
    }
    /**
     * `docker.RemoteImage` provides details about a specific Docker Image which needs to be present on the Docker Host
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.docker.DockerFunctions;
     * import com.pulumi.docker.inputs.GetRemoteImageArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App }{{@code
     *     public static void main(String[] args) }{{@code
     *         Pulumi.run(App::stack);
     *     }}{@code
     * 
     *     public static void stack(Context ctx) }{{@code
     *         // uses the 'latest' tag
     *         final var latest = DockerFunctions.getRemoteImage(GetRemoteImageArgs.builder()
     *             .name("nginx")
     *             .build());
     * 
     *         // uses a specific tag
     *         final var specific = DockerFunctions.getRemoteImage(GetRemoteImageArgs.builder()
     *             .name("nginx:1.17.6")
     *             .build());
     * 
     *         // use the image digest
     *         final var digest = DockerFunctions.getRemoteImage(GetRemoteImageArgs.builder()
     *             .name("nginx}{@literal @}{@code sha256:36b74457bccb56fbf8b05f79c85569501b721d4db813b684391d63e02287c0b2")
     *             .build());
     * 
     *         // uses the tag and the image digest
     *         final var tagAndDigest = DockerFunctions.getRemoteImage(GetRemoteImageArgs.builder()
     *             .name("nginx:1.19.1}{@literal @}{@code sha256:36b74457bccb56fbf8b05f79c85569501b721d4db813b684391d63e02287c0b2")
     *             .build());
     * 
     *     }}{@code
     * }}{@code
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetRemoteImageResult> getRemoteImage(GetRemoteImageArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("docker:index/getRemoteImage:getRemoteImage", TypeShape.of(GetRemoteImageResult.class), args, Utilities.withVersion(options));
    }
    /**
     * `docker.RemoteImage` provides details about a specific Docker Image which needs to be present on the Docker Host
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.docker.DockerFunctions;
     * import com.pulumi.docker.inputs.GetRemoteImageArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App }{{@code
     *     public static void main(String[] args) }{{@code
     *         Pulumi.run(App::stack);
     *     }}{@code
     * 
     *     public static void stack(Context ctx) }{{@code
     *         // uses the 'latest' tag
     *         final var latest = DockerFunctions.getRemoteImage(GetRemoteImageArgs.builder()
     *             .name("nginx")
     *             .build());
     * 
     *         // uses a specific tag
     *         final var specific = DockerFunctions.getRemoteImage(GetRemoteImageArgs.builder()
     *             .name("nginx:1.17.6")
     *             .build());
     * 
     *         // use the image digest
     *         final var digest = DockerFunctions.getRemoteImage(GetRemoteImageArgs.builder()
     *             .name("nginx}{@literal @}{@code sha256:36b74457bccb56fbf8b05f79c85569501b721d4db813b684391d63e02287c0b2")
     *             .build());
     * 
     *         // uses the tag and the image digest
     *         final var tagAndDigest = DockerFunctions.getRemoteImage(GetRemoteImageArgs.builder()
     *             .name("nginx:1.19.1}{@literal @}{@code sha256:36b74457bccb56fbf8b05f79c85569501b721d4db813b684391d63e02287c0b2")
     *             .build());
     * 
     *     }}{@code
     * }}{@code
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetRemoteImageResult> getRemoteImagePlain(GetRemoteImagePlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("docker:index/getRemoteImage:getRemoteImage", TypeShape.of(GetRemoteImageResult.class), args, Utilities.withVersion(options));
    }
}
