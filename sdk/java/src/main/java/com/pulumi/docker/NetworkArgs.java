// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.docker.inputs.NetworkIpamConfigArgs;
import com.pulumi.docker.inputs.NetworkLabelArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class NetworkArgs extends com.pulumi.resources.ResourceArgs {

    public static final NetworkArgs Empty = new NetworkArgs();

    /**
     * Enable manual container attachment to the network.
     * 
     */
    @Import(name="attachable")
    private @Nullable Output<Boolean> attachable;

    /**
     * @return Enable manual container attachment to the network.
     * 
     */
    public Optional<Output<Boolean>> attachable() {
        return Optional.ofNullable(this.attachable);
    }

    /**
     * Requests daemon to check for networks with same name.
     * 
     */
    @Import(name="checkDuplicate")
    private @Nullable Output<Boolean> checkDuplicate;

    /**
     * @return Requests daemon to check for networks with same name.
     * 
     */
    public Optional<Output<Boolean>> checkDuplicate() {
        return Optional.ofNullable(this.checkDuplicate);
    }

    /**
     * The driver of the Docker network. Possible values are `bridge`, `host`, `overlay`, `macvlan`. See [network docs](https://docs.docker.com/network/#network-drivers) for more details.
     * 
     */
    @Import(name="driver")
    private @Nullable Output<String> driver;

    /**
     * @return The driver of the Docker network. Possible values are `bridge`, `host`, `overlay`, `macvlan`. See [network docs](https://docs.docker.com/network/#network-drivers) for more details.
     * 
     */
    public Optional<Output<String>> driver() {
        return Optional.ofNullable(this.driver);
    }

    /**
     * Create swarm routing-mesh network. Defaults to `false`.
     * 
     */
    @Import(name="ingress")
    private @Nullable Output<Boolean> ingress;

    /**
     * @return Create swarm routing-mesh network. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> ingress() {
        return Optional.ofNullable(this.ingress);
    }

    /**
     * Whether the network is internal.
     * 
     */
    @Import(name="internal")
    private @Nullable Output<Boolean> internal;

    /**
     * @return Whether the network is internal.
     * 
     */
    public Optional<Output<Boolean>> internal() {
        return Optional.ofNullable(this.internal);
    }

    /**
     * The IPAM configuration options
     * 
     */
    @Import(name="ipamConfigs")
    private @Nullable Output<List<NetworkIpamConfigArgs>> ipamConfigs;

    /**
     * @return The IPAM configuration options
     * 
     */
    public Optional<Output<List<NetworkIpamConfigArgs>>> ipamConfigs() {
        return Optional.ofNullable(this.ipamConfigs);
    }

    /**
     * Driver used by the custom IP scheme of the network. Defaults to `default`
     * 
     */
    @Import(name="ipamDriver")
    private @Nullable Output<String> ipamDriver;

    /**
     * @return Driver used by the custom IP scheme of the network. Defaults to `default`
     * 
     */
    public Optional<Output<String>> ipamDriver() {
        return Optional.ofNullable(this.ipamDriver);
    }

    /**
     * Provide explicit options to the IPAM driver. Valid options vary with `ipam_driver` and refer to that driver&#39;s documentation for more details.
     * 
     */
    @Import(name="ipamOptions")
    private @Nullable Output<Map<String,String>> ipamOptions;

    /**
     * @return Provide explicit options to the IPAM driver. Valid options vary with `ipam_driver` and refer to that driver&#39;s documentation for more details.
     * 
     */
    public Optional<Output<Map<String,String>>> ipamOptions() {
        return Optional.ofNullable(this.ipamOptions);
    }

    /**
     * Enable IPv6 networking. Defaults to `false`.
     * 
     */
    @Import(name="ipv6")
    private @Nullable Output<Boolean> ipv6;

    /**
     * @return Enable IPv6 networking. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> ipv6() {
        return Optional.ofNullable(this.ipv6);
    }

    /**
     * User-defined key/value metadata
     * 
     */
    @Import(name="labels")
    private @Nullable Output<List<NetworkLabelArgs>> labels;

    /**
     * @return User-defined key/value metadata
     * 
     */
    public Optional<Output<List<NetworkLabelArgs>>> labels() {
        return Optional.ofNullable(this.labels);
    }

    /**
     * The name of the Docker network.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the Docker network.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Only available with bridge networks. See [bridge options docs](https://docs.docker.com/engine/reference/commandline/network_create/#bridge-driver-options) for more details.
     * 
     */
    @Import(name="options")
    private @Nullable Output<Map<String,String>> options;

    /**
     * @return Only available with bridge networks. See [bridge options docs](https://docs.docker.com/engine/reference/commandline/network_create/#bridge-driver-options) for more details.
     * 
     */
    public Optional<Output<Map<String,String>>> options() {
        return Optional.ofNullable(this.options);
    }

    private NetworkArgs() {}

    private NetworkArgs(NetworkArgs $) {
        this.attachable = $.attachable;
        this.checkDuplicate = $.checkDuplicate;
        this.driver = $.driver;
        this.ingress = $.ingress;
        this.internal = $.internal;
        this.ipamConfigs = $.ipamConfigs;
        this.ipamDriver = $.ipamDriver;
        this.ipamOptions = $.ipamOptions;
        this.ipv6 = $.ipv6;
        this.labels = $.labels;
        this.name = $.name;
        this.options = $.options;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(NetworkArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private NetworkArgs $;

        public Builder() {
            $ = new NetworkArgs();
        }

        public Builder(NetworkArgs defaults) {
            $ = new NetworkArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param attachable Enable manual container attachment to the network.
         * 
         * @return builder
         * 
         */
        public Builder attachable(@Nullable Output<Boolean> attachable) {
            $.attachable = attachable;
            return this;
        }

        /**
         * @param attachable Enable manual container attachment to the network.
         * 
         * @return builder
         * 
         */
        public Builder attachable(Boolean attachable) {
            return attachable(Output.of(attachable));
        }

        /**
         * @param checkDuplicate Requests daemon to check for networks with same name.
         * 
         * @return builder
         * 
         */
        public Builder checkDuplicate(@Nullable Output<Boolean> checkDuplicate) {
            $.checkDuplicate = checkDuplicate;
            return this;
        }

        /**
         * @param checkDuplicate Requests daemon to check for networks with same name.
         * 
         * @return builder
         * 
         */
        public Builder checkDuplicate(Boolean checkDuplicate) {
            return checkDuplicate(Output.of(checkDuplicate));
        }

        /**
         * @param driver The driver of the Docker network. Possible values are `bridge`, `host`, `overlay`, `macvlan`. See [network docs](https://docs.docker.com/network/#network-drivers) for more details.
         * 
         * @return builder
         * 
         */
        public Builder driver(@Nullable Output<String> driver) {
            $.driver = driver;
            return this;
        }

        /**
         * @param driver The driver of the Docker network. Possible values are `bridge`, `host`, `overlay`, `macvlan`. See [network docs](https://docs.docker.com/network/#network-drivers) for more details.
         * 
         * @return builder
         * 
         */
        public Builder driver(String driver) {
            return driver(Output.of(driver));
        }

        /**
         * @param ingress Create swarm routing-mesh network. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder ingress(@Nullable Output<Boolean> ingress) {
            $.ingress = ingress;
            return this;
        }

        /**
         * @param ingress Create swarm routing-mesh network. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder ingress(Boolean ingress) {
            return ingress(Output.of(ingress));
        }

        /**
         * @param internal Whether the network is internal.
         * 
         * @return builder
         * 
         */
        public Builder internal(@Nullable Output<Boolean> internal) {
            $.internal = internal;
            return this;
        }

        /**
         * @param internal Whether the network is internal.
         * 
         * @return builder
         * 
         */
        public Builder internal(Boolean internal) {
            return internal(Output.of(internal));
        }

        /**
         * @param ipamConfigs The IPAM configuration options
         * 
         * @return builder
         * 
         */
        public Builder ipamConfigs(@Nullable Output<List<NetworkIpamConfigArgs>> ipamConfigs) {
            $.ipamConfigs = ipamConfigs;
            return this;
        }

        /**
         * @param ipamConfigs The IPAM configuration options
         * 
         * @return builder
         * 
         */
        public Builder ipamConfigs(List<NetworkIpamConfigArgs> ipamConfigs) {
            return ipamConfigs(Output.of(ipamConfigs));
        }

        /**
         * @param ipamConfigs The IPAM configuration options
         * 
         * @return builder
         * 
         */
        public Builder ipamConfigs(NetworkIpamConfigArgs... ipamConfigs) {
            return ipamConfigs(List.of(ipamConfigs));
        }

        /**
         * @param ipamDriver Driver used by the custom IP scheme of the network. Defaults to `default`
         * 
         * @return builder
         * 
         */
        public Builder ipamDriver(@Nullable Output<String> ipamDriver) {
            $.ipamDriver = ipamDriver;
            return this;
        }

        /**
         * @param ipamDriver Driver used by the custom IP scheme of the network. Defaults to `default`
         * 
         * @return builder
         * 
         */
        public Builder ipamDriver(String ipamDriver) {
            return ipamDriver(Output.of(ipamDriver));
        }

        /**
         * @param ipamOptions Provide explicit options to the IPAM driver. Valid options vary with `ipam_driver` and refer to that driver&#39;s documentation for more details.
         * 
         * @return builder
         * 
         */
        public Builder ipamOptions(@Nullable Output<Map<String,String>> ipamOptions) {
            $.ipamOptions = ipamOptions;
            return this;
        }

        /**
         * @param ipamOptions Provide explicit options to the IPAM driver. Valid options vary with `ipam_driver` and refer to that driver&#39;s documentation for more details.
         * 
         * @return builder
         * 
         */
        public Builder ipamOptions(Map<String,String> ipamOptions) {
            return ipamOptions(Output.of(ipamOptions));
        }

        /**
         * @param ipv6 Enable IPv6 networking. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder ipv6(@Nullable Output<Boolean> ipv6) {
            $.ipv6 = ipv6;
            return this;
        }

        /**
         * @param ipv6 Enable IPv6 networking. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder ipv6(Boolean ipv6) {
            return ipv6(Output.of(ipv6));
        }

        /**
         * @param labels User-defined key/value metadata
         * 
         * @return builder
         * 
         */
        public Builder labels(@Nullable Output<List<NetworkLabelArgs>> labels) {
            $.labels = labels;
            return this;
        }

        /**
         * @param labels User-defined key/value metadata
         * 
         * @return builder
         * 
         */
        public Builder labels(List<NetworkLabelArgs> labels) {
            return labels(Output.of(labels));
        }

        /**
         * @param labels User-defined key/value metadata
         * 
         * @return builder
         * 
         */
        public Builder labels(NetworkLabelArgs... labels) {
            return labels(List.of(labels));
        }

        /**
         * @param name The name of the Docker network.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the Docker network.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param options Only available with bridge networks. See [bridge options docs](https://docs.docker.com/engine/reference/commandline/network_create/#bridge-driver-options) for more details.
         * 
         * @return builder
         * 
         */
        public Builder options(@Nullable Output<Map<String,String>> options) {
            $.options = options;
            return this;
        }

        /**
         * @param options Only available with bridge networks. See [bridge options docs](https://docs.docker.com/engine/reference/commandline/network_create/#bridge-driver-options) for more details.
         * 
         * @return builder
         * 
         */
        public Builder options(Map<String,String> options) {
            return options(Output.of(options));
        }

        public NetworkArgs build() {
            return $;
        }
    }

}
