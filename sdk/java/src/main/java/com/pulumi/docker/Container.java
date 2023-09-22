// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.docker.ContainerArgs;
import com.pulumi.docker.Utilities;
import com.pulumi.docker.inputs.ContainerState;
import com.pulumi.docker.outputs.ContainerCapabilities;
import com.pulumi.docker.outputs.ContainerDevice;
import com.pulumi.docker.outputs.ContainerHealthcheck;
import com.pulumi.docker.outputs.ContainerHost;
import com.pulumi.docker.outputs.ContainerLabel;
import com.pulumi.docker.outputs.ContainerMount;
import com.pulumi.docker.outputs.ContainerNetworkData;
import com.pulumi.docker.outputs.ContainerNetworksAdvanced;
import com.pulumi.docker.outputs.ContainerPort;
import com.pulumi.docker.outputs.ContainerUlimit;
import com.pulumi.docker.outputs.ContainerUpload;
import com.pulumi.docker.outputs.ContainerVolume;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * &lt;!-- Bug: Type and Name are switched --&gt;
 * Manages the lifecycle of a Docker container.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.docker.RemoteImage;
 * import com.pulumi.docker.RemoteImageArgs;
 * import com.pulumi.docker.Container;
 * import com.pulumi.docker.ContainerArgs;
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
 *         var ubuntuRemoteImage = new RemoteImage(&#34;ubuntuRemoteImage&#34;, RemoteImageArgs.builder()        
 *             .name(&#34;ubuntu:precise&#34;)
 *             .build());
 * 
 *         var ubuntuContainer = new Container(&#34;ubuntuContainer&#34;, ContainerArgs.builder()        
 *             .image(ubuntuRemoteImage.imageId())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * ### Example Assuming you created a `container` as follows #!/bin/bash docker run --name foo -p8080:80 -d nginx
 * 
 * prints the container ID
 * 
 * 9a550c0f0163d39d77222d3efd58701b625d47676c25c686c95b5b92d1cba6fd you provide the definition for the resource as follows terraform resource &#34;docker_container&#34; &#34;foo&#34; {
 * 
 *  name
 * 
 * = &#34;foo&#34;
 * 
 *  image = &#34;nginx&#34;
 * 
 *  ports {
 * 
 *  internal = &#34;80&#34;
 * 
 *  external = &#34;8080&#34;
 * 
 *  } } then the import command is as follows #!/bin/bash
 * 
 * ```sh
 *  $ pulumi import docker:index/container:Container foo 9a550c0f0163d39d77222d3efd58701b625d47676c25c686c95b5b92d1cba6fd
 * ```
 * 
 */
@ResourceType(type="docker:index/container:Container")
public class Container extends com.pulumi.resources.CustomResource {
    /**
     * If `true` attach to the container after its creation and waits the end of its execution. Defaults to `false`.
     * 
     */
    @Export(name="attach", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> attach;

    /**
     * @return If `true` attach to the container after its creation and waits the end of its execution. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> attach() {
        return Codegen.optional(this.attach);
    }
    /**
     * The network bridge of the container as read from its NetworkSettings.
     * 
     */
    @Export(name="bridge", refs={String.class}, tree="[0]")
    private Output<String> bridge;

    /**
     * @return The network bridge of the container as read from its NetworkSettings.
     * 
     */
    public Output<String> bridge() {
        return this.bridge;
    }
    /**
     * Add or drop certrain linux capabilities.
     * 
     */
    @Export(name="capabilities", refs={ContainerCapabilities.class}, tree="[0]")
    private Output</* @Nullable */ ContainerCapabilities> capabilities;

    /**
     * @return Add or drop certrain linux capabilities.
     * 
     */
    public Output<Optional<ContainerCapabilities>> capabilities() {
        return Codegen.optional(this.capabilities);
    }
    /**
     * Cgroup namespace mode to use for the container. Possible values are: `private`, `host`.
     * 
     */
    @Export(name="cgroupnsMode", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> cgroupnsMode;

    /**
     * @return Cgroup namespace mode to use for the container. Possible values are: `private`, `host`.
     * 
     */
    public Output<Optional<String>> cgroupnsMode() {
        return Codegen.optional(this.cgroupnsMode);
    }
    /**
     * The command to use to start the container. For example, to run `/usr/bin/myprogram -f baz.conf` set the command to be `[&#34;/usr/bin/myprogram&#34;,&#34;-f&#34;,&#34;baz.con&#34;]`.
     * 
     */
    @Export(name="command", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> command;

    /**
     * @return The command to use to start the container. For example, to run `/usr/bin/myprogram -f baz.conf` set the command to be `[&#34;/usr/bin/myprogram&#34;,&#34;-f&#34;,&#34;baz.con&#34;]`.
     * 
     */
    public Output<List<String>> command() {
        return this.command;
    }
    /**
     * The logs of the container if its execution is done (`attach` must be disabled).
     * 
     */
    @Export(name="containerLogs", refs={String.class}, tree="[0]")
    private Output<String> containerLogs;

    /**
     * @return The logs of the container if its execution is done (`attach` must be disabled).
     * 
     */
    public Output<String> containerLogs() {
        return this.containerLogs;
    }
    /**
     * The total number of milliseconds to wait for the container to reach status &#39;running&#39;
     * 
     */
    @Export(name="containerReadRefreshTimeoutMilliseconds", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> containerReadRefreshTimeoutMilliseconds;

    /**
     * @return The total number of milliseconds to wait for the container to reach status &#39;running&#39;
     * 
     */
    public Output<Optional<Integer>> containerReadRefreshTimeoutMilliseconds() {
        return Codegen.optional(this.containerReadRefreshTimeoutMilliseconds);
    }
    /**
     * A comma-separated list or hyphen-separated range of CPUs a container can use, e.g. `0-1`.
     * 
     */
    @Export(name="cpuSet", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> cpuSet;

    /**
     * @return A comma-separated list or hyphen-separated range of CPUs a container can use, e.g. `0-1`.
     * 
     */
    public Output<Optional<String>> cpuSet() {
        return Codegen.optional(this.cpuSet);
    }
    /**
     * CPU shares (relative weight) for the container.
     * 
     */
    @Export(name="cpuShares", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> cpuShares;

    /**
     * @return CPU shares (relative weight) for the container.
     * 
     */
    public Output<Optional<Integer>> cpuShares() {
        return Codegen.optional(this.cpuShares);
    }
    /**
     * If defined will attempt to stop the container before destroying. Container will be destroyed after `n` seconds or on successful stop.
     * 
     */
    @Export(name="destroyGraceSeconds", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> destroyGraceSeconds;

    /**
     * @return If defined will attempt to stop the container before destroying. Container will be destroyed after `n` seconds or on successful stop.
     * 
     */
    public Output<Optional<Integer>> destroyGraceSeconds() {
        return Codegen.optional(this.destroyGraceSeconds);
    }
    /**
     * Bind devices to the container.
     * 
     */
    @Export(name="devices", refs={List.class,ContainerDevice.class}, tree="[0,1]")
    private Output</* @Nullable */ List<ContainerDevice>> devices;

    /**
     * @return Bind devices to the container.
     * 
     */
    public Output<Optional<List<ContainerDevice>>> devices() {
        return Codegen.optional(this.devices);
    }
    /**
     * DNS servers to use.
     * 
     */
    @Export(name="dns", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> dns;

    /**
     * @return DNS servers to use.
     * 
     */
    public Output<Optional<List<String>>> dns() {
        return Codegen.optional(this.dns);
    }
    /**
     * DNS options used by the DNS provider(s), see `resolv.conf` documentation for valid list of options.
     * 
     */
    @Export(name="dnsOpts", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> dnsOpts;

    /**
     * @return DNS options used by the DNS provider(s), see `resolv.conf` documentation for valid list of options.
     * 
     */
    public Output<Optional<List<String>>> dnsOpts() {
        return Codegen.optional(this.dnsOpts);
    }
    /**
     * DNS search domains that are used when bare unqualified hostnames are used inside of the container.
     * 
     */
    @Export(name="dnsSearches", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> dnsSearches;

    /**
     * @return DNS search domains that are used when bare unqualified hostnames are used inside of the container.
     * 
     */
    public Output<Optional<List<String>>> dnsSearches() {
        return Codegen.optional(this.dnsSearches);
    }
    /**
     * Domain name of the container.
     * 
     */
    @Export(name="domainname", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> domainname;

    /**
     * @return Domain name of the container.
     * 
     */
    public Output<Optional<String>> domainname() {
        return Codegen.optional(this.domainname);
    }
    /**
     * The command to use as the Entrypoint for the container. The Entrypoint allows you to configure a container to run as an executable. For example, to run `/usr/bin/myprogram` when starting a container, set the entrypoint to be `&#34;/usr/bin/myprogra&#34;]`.
     * 
     */
    @Export(name="entrypoints", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> entrypoints;

    /**
     * @return The command to use as the Entrypoint for the container. The Entrypoint allows you to configure a container to run as an executable. For example, to run `/usr/bin/myprogram` when starting a container, set the entrypoint to be `&#34;/usr/bin/myprogra&#34;]`.
     * 
     */
    public Output<List<String>> entrypoints() {
        return this.entrypoints;
    }
    /**
     * Environment variables to set in the form of `KEY=VALUE`, e.g. `DEBUG=0`
     * 
     */
    @Export(name="envs", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> envs;

    /**
     * @return Environment variables to set in the form of `KEY=VALUE`, e.g. `DEBUG=0`
     * 
     */
    public Output<List<String>> envs() {
        return this.envs;
    }
    /**
     * The exit code of the container if its execution is done (`must_run` must be disabled).
     * 
     */
    @Export(name="exitCode", refs={Integer.class}, tree="[0]")
    private Output<Integer> exitCode;

    /**
     * @return The exit code of the container if its execution is done (`must_run` must be disabled).
     * 
     */
    public Output<Integer> exitCode() {
        return this.exitCode;
    }
    /**
     * GPU devices to add to the container. Currently, only the value `all` is supported. Passing any other value will result in unexpected behavior.
     * 
     */
    @Export(name="gpus", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> gpus;

    /**
     * @return GPU devices to add to the container. Currently, only the value `all` is supported. Passing any other value will result in unexpected behavior.
     * 
     */
    public Output<Optional<String>> gpus() {
        return Codegen.optional(this.gpus);
    }
    /**
     * Additional groups for the container user
     * 
     */
    @Export(name="groupAdds", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> groupAdds;

    /**
     * @return Additional groups for the container user
     * 
     */
    public Output<Optional<List<String>>> groupAdds() {
        return Codegen.optional(this.groupAdds);
    }
    /**
     * A test to perform to check that the container is healthy
     * 
     */
    @Export(name="healthcheck", refs={ContainerHealthcheck.class}, tree="[0]")
    private Output<ContainerHealthcheck> healthcheck;

    /**
     * @return A test to perform to check that the container is healthy
     * 
     */
    public Output<ContainerHealthcheck> healthcheck() {
        return this.healthcheck;
    }
    /**
     * Hostname of the container.
     * 
     */
    @Export(name="hostname", refs={String.class}, tree="[0]")
    private Output<String> hostname;

    /**
     * @return Hostname of the container.
     * 
     */
    public Output<String> hostname() {
        return this.hostname;
    }
    /**
     * Additional hosts to add to the container.
     * 
     */
    @Export(name="hosts", refs={List.class,ContainerHost.class}, tree="[0,1]")
    private Output</* @Nullable */ List<ContainerHost>> hosts;

    /**
     * @return Additional hosts to add to the container.
     * 
     */
    public Output<Optional<List<ContainerHost>>> hosts() {
        return Codegen.optional(this.hosts);
    }
    /**
     * The ID of the image to back this container. The easiest way to get this value is to use the `docker.RemoteImage` resource as is shown in the example.
     * 
     */
    @Export(name="image", refs={String.class}, tree="[0]")
    private Output<String> image;

    /**
     * @return The ID of the image to back this container. The easiest way to get this value is to use the `docker.RemoteImage` resource as is shown in the example.
     * 
     */
    public Output<String> image() {
        return this.image;
    }
    /**
     * Configured whether an init process should be injected for this container. If unset this will default to the `dockerd` defaults.
     * 
     */
    @Export(name="init", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> init;

    /**
     * @return Configured whether an init process should be injected for this container. If unset this will default to the `dockerd` defaults.
     * 
     */
    public Output<Boolean> init() {
        return this.init;
    }
    /**
     * IPC sharing mode for the container. Possible values are: `none`, `private`, `shareable`, `container:&lt;name|id&gt;` or `host`.
     * 
     */
    @Export(name="ipcMode", refs={String.class}, tree="[0]")
    private Output<String> ipcMode;

    /**
     * @return IPC sharing mode for the container. Possible values are: `none`, `private`, `shareable`, `container:&lt;name|id&gt;` or `host`.
     * 
     */
    public Output<String> ipcMode() {
        return this.ipcMode;
    }
    /**
     * User-defined key/value metadata
     * 
     */
    @Export(name="labels", refs={List.class,ContainerLabel.class}, tree="[0,1]")
    private Output<List<ContainerLabel>> labels;

    /**
     * @return User-defined key/value metadata
     * 
     */
    public Output<List<ContainerLabel>> labels() {
        return this.labels;
    }
    /**
     * The logging driver to use for the container.
     * 
     */
    @Export(name="logDriver", refs={String.class}, tree="[0]")
    private Output<String> logDriver;

    /**
     * @return The logging driver to use for the container.
     * 
     */
    public Output<String> logDriver() {
        return this.logDriver;
    }
    /**
     * Key/value pairs to use as options for the logging driver.
     * 
     */
    @Export(name="logOpts", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> logOpts;

    /**
     * @return Key/value pairs to use as options for the logging driver.
     * 
     */
    public Output<Optional<Map<String,Object>>> logOpts() {
        return Codegen.optional(this.logOpts);
    }
    /**
     * Save the container logs (`attach` must be enabled). Defaults to `false`.
     * 
     */
    @Export(name="logs", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> logs;

    /**
     * @return Save the container logs (`attach` must be enabled). Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> logs() {
        return Codegen.optional(this.logs);
    }
    /**
     * The maximum amount of times to an attempt a restart when `restart` is set to &#39;on-failure&#39;.
     * 
     */
    @Export(name="maxRetryCount", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> maxRetryCount;

    /**
     * @return The maximum amount of times to an attempt a restart when `restart` is set to &#39;on-failure&#39;.
     * 
     */
    public Output<Optional<Integer>> maxRetryCount() {
        return Codegen.optional(this.maxRetryCount);
    }
    /**
     * The memory limit for the container in MBs.
     * 
     */
    @Export(name="memory", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> memory;

    /**
     * @return The memory limit for the container in MBs.
     * 
     */
    public Output<Optional<Integer>> memory() {
        return Codegen.optional(this.memory);
    }
    /**
     * The total memory limit (memory + swap) for the container in MBs. This setting may compute to `-1` after `pulumi up` if the target host doesn&#39;t support memory swap, when that is the case docker will use a soft limitation.
     * 
     */
    @Export(name="memorySwap", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> memorySwap;

    /**
     * @return The total memory limit (memory + swap) for the container in MBs. This setting may compute to `-1` after `pulumi up` if the target host doesn&#39;t support memory swap, when that is the case docker will use a soft limitation.
     * 
     */
    public Output<Optional<Integer>> memorySwap() {
        return Codegen.optional(this.memorySwap);
    }
    /**
     * Specification for mounts to be added to containers created as part of the service.
     * 
     */
    @Export(name="mounts", refs={List.class,ContainerMount.class}, tree="[0,1]")
    private Output</* @Nullable */ List<ContainerMount>> mounts;

    /**
     * @return Specification for mounts to be added to containers created as part of the service.
     * 
     */
    public Output<Optional<List<ContainerMount>>> mounts() {
        return Codegen.optional(this.mounts);
    }
    /**
     * If `true`, then the Docker container will be kept running. If `false`, then as long as the container exists, Terraform
     * assumes it is successful. Defaults to `true`.
     * 
     */
    @Export(name="mustRun", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> mustRun;

    /**
     * @return If `true`, then the Docker container will be kept running. If `false`, then as long as the container exists, Terraform
     * assumes it is successful. Defaults to `true`.
     * 
     */
    public Output<Optional<Boolean>> mustRun() {
        return Codegen.optional(this.mustRun);
    }
    /**
     * The name of the container.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the container.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The data of the networks the container is connected to.
     * 
     */
    @Export(name="networkDatas", refs={List.class,ContainerNetworkData.class}, tree="[0,1]")
    private Output<List<ContainerNetworkData>> networkDatas;

    /**
     * @return The data of the networks the container is connected to.
     * 
     */
    public Output<List<ContainerNetworkData>> networkDatas() {
        return this.networkDatas;
    }
    /**
     * Network mode of the container.
     * 
     */
    @Export(name="networkMode", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> networkMode;

    /**
     * @return Network mode of the container.
     * 
     */
    public Output<Optional<String>> networkMode() {
        return Codegen.optional(this.networkMode);
    }
    /**
     * The networks the container is attached to
     * 
     */
    @Export(name="networksAdvanced", refs={List.class,ContainerNetworksAdvanced.class}, tree="[0,1]")
    private Output</* @Nullable */ List<ContainerNetworksAdvanced>> networksAdvanced;

    /**
     * @return The networks the container is attached to
     * 
     */
    public Output<Optional<List<ContainerNetworksAdvanced>>> networksAdvanced() {
        return Codegen.optional(this.networksAdvanced);
    }
    /**
     * he PID (Process) Namespace mode for the container. Either `container:&lt;name|id&gt;` or `host`.
     * 
     */
    @Export(name="pidMode", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> pidMode;

    /**
     * @return he PID (Process) Namespace mode for the container. Either `container:&lt;name|id&gt;` or `host`.
     * 
     */
    public Output<Optional<String>> pidMode() {
        return Codegen.optional(this.pidMode);
    }
    /**
     * Publish a container&#39;s port(s) to the host.
     * 
     */
    @Export(name="ports", refs={List.class,ContainerPort.class}, tree="[0,1]")
    private Output</* @Nullable */ List<ContainerPort>> ports;

    /**
     * @return Publish a container&#39;s port(s) to the host.
     * 
     */
    public Output<Optional<List<ContainerPort>>> ports() {
        return Codegen.optional(this.ports);
    }
    /**
     * If `true`, the container runs in privileged mode.
     * 
     */
    @Export(name="privileged", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> privileged;

    /**
     * @return If `true`, the container runs in privileged mode.
     * 
     */
    public Output<Optional<Boolean>> privileged() {
        return Codegen.optional(this.privileged);
    }
    /**
     * Publish all ports of the container.
     * 
     */
    @Export(name="publishAllPorts", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> publishAllPorts;

    /**
     * @return Publish all ports of the container.
     * 
     */
    public Output<Optional<Boolean>> publishAllPorts() {
        return Codegen.optional(this.publishAllPorts);
    }
    /**
     * If `true`, the container will be started as readonly. Defaults to `false`.
     * 
     */
    @Export(name="readOnly", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> readOnly;

    /**
     * @return If `true`, the container will be started as readonly. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> readOnly() {
        return Codegen.optional(this.readOnly);
    }
    /**
     * If `true`, it will remove anonymous volumes associated with the container. Defaults to `true`.
     * 
     */
    @Export(name="removeVolumes", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> removeVolumes;

    /**
     * @return If `true`, it will remove anonymous volumes associated with the container. Defaults to `true`.
     * 
     */
    public Output<Optional<Boolean>> removeVolumes() {
        return Codegen.optional(this.removeVolumes);
    }
    /**
     * The restart policy for the container. Must be one of &#39;no&#39;, &#39;on-failure&#39;, &#39;always&#39;, &#39;unless-stopped&#39;. Defaults to `no`.
     * 
     */
    @Export(name="restart", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> restart;

    /**
     * @return The restart policy for the container. Must be one of &#39;no&#39;, &#39;on-failure&#39;, &#39;always&#39;, &#39;unless-stopped&#39;. Defaults to `no`.
     * 
     */
    public Output<Optional<String>> restart() {
        return Codegen.optional(this.restart);
    }
    /**
     * If `true`, then the container will be automatically removed when it exits. Defaults to `false`.
     * 
     */
    @Export(name="rm", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> rm;

    /**
     * @return If `true`, then the container will be automatically removed when it exits. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> rm() {
        return Codegen.optional(this.rm);
    }
    /**
     * Runtime to use for the container.
     * 
     */
    @Export(name="runtime", refs={String.class}, tree="[0]")
    private Output<String> runtime;

    /**
     * @return Runtime to use for the container.
     * 
     */
    public Output<String> runtime() {
        return this.runtime;
    }
    /**
     * List of string values to customize labels for MLS systems, such as SELinux. See https://docs.docker.com/engine/reference/run/#security-configuration.
     * 
     */
    @Export(name="securityOpts", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> securityOpts;

    /**
     * @return List of string values to customize labels for MLS systems, such as SELinux. See https://docs.docker.com/engine/reference/run/#security-configuration.
     * 
     */
    public Output<List<String>> securityOpts() {
        return this.securityOpts;
    }
    /**
     * Size of `/dev/shm` in MBs.
     * 
     */
    @Export(name="shmSize", refs={Integer.class}, tree="[0]")
    private Output<Integer> shmSize;

    /**
     * @return Size of `/dev/shm` in MBs.
     * 
     */
    public Output<Integer> shmSize() {
        return this.shmSize;
    }
    /**
     * If `true`, then the Docker container will be started after creation. If `false`, then the container is only created. Defaults to `true`.
     * 
     */
    @Export(name="start", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> start;

    /**
     * @return If `true`, then the Docker container will be started after creation. If `false`, then the container is only created. Defaults to `true`.
     * 
     */
    public Output<Optional<Boolean>> start() {
        return Codegen.optional(this.start);
    }
    /**
     * If `true`, keep STDIN open even if not attached (`docker run -i`). Defaults to `false`.
     * 
     */
    @Export(name="stdinOpen", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> stdinOpen;

    /**
     * @return If `true`, keep STDIN open even if not attached (`docker run -i`). Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> stdinOpen() {
        return Codegen.optional(this.stdinOpen);
    }
    /**
     * Signal to stop a container (default `SIGTERM`).
     * 
     */
    @Export(name="stopSignal", refs={String.class}, tree="[0]")
    private Output<String> stopSignal;

    /**
     * @return Signal to stop a container (default `SIGTERM`).
     * 
     */
    public Output<String> stopSignal() {
        return this.stopSignal;
    }
    /**
     * Timeout (in seconds) to stop a container.
     * 
     */
    @Export(name="stopTimeout", refs={Integer.class}, tree="[0]")
    private Output<Integer> stopTimeout;

    /**
     * @return Timeout (in seconds) to stop a container.
     * 
     */
    public Output<Integer> stopTimeout() {
        return this.stopTimeout;
    }
    /**
     * Key/value pairs for the storage driver options, e.g. `size`: `120G`
     * 
     */
    @Export(name="storageOpts", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> storageOpts;

    /**
     * @return Key/value pairs for the storage driver options, e.g. `size`: `120G`
     * 
     */
    public Output<Optional<Map<String,Object>>> storageOpts() {
        return Codegen.optional(this.storageOpts);
    }
    /**
     * A map of kernel parameters (sysctls) to set in the container.
     * 
     */
    @Export(name="sysctls", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> sysctls;

    /**
     * @return A map of kernel parameters (sysctls) to set in the container.
     * 
     */
    public Output<Optional<Map<String,Object>>> sysctls() {
        return Codegen.optional(this.sysctls);
    }
    /**
     * A map of container directories which should be replaced by `tmpfs mounts`, and their corresponding mount options.
     * 
     */
    @Export(name="tmpfs", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> tmpfs;

    /**
     * @return A map of container directories which should be replaced by `tmpfs mounts`, and their corresponding mount options.
     * 
     */
    public Output<Optional<Map<String,Object>>> tmpfs() {
        return Codegen.optional(this.tmpfs);
    }
    /**
     * If `true`, allocate a pseudo-tty (`docker run -t`). Defaults to `false`.
     * 
     */
    @Export(name="tty", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> tty;

    /**
     * @return If `true`, allocate a pseudo-tty (`docker run -t`). Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> tty() {
        return Codegen.optional(this.tty);
    }
    /**
     * Ulimit options to add.
     * 
     */
    @Export(name="ulimits", refs={List.class,ContainerUlimit.class}, tree="[0,1]")
    private Output</* @Nullable */ List<ContainerUlimit>> ulimits;

    /**
     * @return Ulimit options to add.
     * 
     */
    public Output<Optional<List<ContainerUlimit>>> ulimits() {
        return Codegen.optional(this.ulimits);
    }
    /**
     * Specifies files to upload to the container before starting it. Only one of `content` or `content_base64` can be set and at least one of them has to be set.
     * 
     */
    @Export(name="uploads", refs={List.class,ContainerUpload.class}, tree="[0,1]")
    private Output</* @Nullable */ List<ContainerUpload>> uploads;

    /**
     * @return Specifies files to upload to the container before starting it. Only one of `content` or `content_base64` can be set and at least one of them has to be set.
     * 
     */
    public Output<Optional<List<ContainerUpload>>> uploads() {
        return Codegen.optional(this.uploads);
    }
    /**
     * User used for run the first process. Format is `user` or `user:group` which user and group can be passed literraly or by name.
     * 
     */
    @Export(name="user", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> user;

    /**
     * @return User used for run the first process. Format is `user` or `user:group` which user and group can be passed literraly or by name.
     * 
     */
    public Output<Optional<String>> user() {
        return Codegen.optional(this.user);
    }
    /**
     * Sets the usernamespace mode for the container when usernamespace remapping option is enabled.
     * 
     */
    @Export(name="usernsMode", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> usernsMode;

    /**
     * @return Sets the usernamespace mode for the container when usernamespace remapping option is enabled.
     * 
     */
    public Output<Optional<String>> usernsMode() {
        return Codegen.optional(this.usernsMode);
    }
    /**
     * Spec for mounting volumes in the container.
     * 
     */
    @Export(name="volumes", refs={List.class,ContainerVolume.class}, tree="[0,1]")
    private Output</* @Nullable */ List<ContainerVolume>> volumes;

    /**
     * @return Spec for mounting volumes in the container.
     * 
     */
    public Output<Optional<List<ContainerVolume>>> volumes() {
        return Codegen.optional(this.volumes);
    }
    /**
     * If `true`, then the Docker container is waited for being healthy state after creation. If `false`, then the container health state is not checked. Defaults to `false`.
     * 
     */
    @Export(name="wait", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> wait;

    /**
     * @return If `true`, then the Docker container is waited for being healthy state after creation. If `false`, then the container health state is not checked. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> wait_() {
        return Codegen.optional(this.wait);
    }
    /**
     * The timeout in seconds to wait the container to be healthy after creation. Defaults to `60`.
     * 
     */
    @Export(name="waitTimeout", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> waitTimeout;

    /**
     * @return The timeout in seconds to wait the container to be healthy after creation. Defaults to `60`.
     * 
     */
    public Output<Optional<Integer>> waitTimeout() {
        return Codegen.optional(this.waitTimeout);
    }
    /**
     * The working directory for commands to run in.
     * 
     */
    @Export(name="workingDir", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> workingDir;

    /**
     * @return The working directory for commands to run in.
     * 
     */
    public Output<Optional<String>> workingDir() {
        return Codegen.optional(this.workingDir);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Container(String name) {
        this(name, ContainerArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Container(String name, ContainerArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Container(String name, ContainerArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("docker:index/container:Container", name, args == null ? ContainerArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Container(String name, Output<String> id, @Nullable ContainerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("docker:index/container:Container", name, state, makeResourceOptions(options, id));
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
    public static Container get(String name, Output<String> id, @Nullable ContainerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Container(name, id, state, options);
    }
}
