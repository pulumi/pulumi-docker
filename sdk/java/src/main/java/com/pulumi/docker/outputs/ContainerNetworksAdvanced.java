// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ContainerNetworksAdvanced {
    /**
     * @return The network aliases of the container in the specific network.
     * 
     */
    private final @Nullable List<String> aliases;
    /**
     * @return The IPV4 address of the container in the specific network.
     * 
     */
    private final @Nullable String ipv4Address;
    /**
     * @return The IPV6 address of the container in the specific network.
     * 
     */
    private final @Nullable String ipv6Address;
    /**
     * @return The name of the network.
     * 
     */
    private final String name;

    @CustomType.Constructor
    private ContainerNetworksAdvanced(
        @CustomType.Parameter("aliases") @Nullable List<String> aliases,
        @CustomType.Parameter("ipv4Address") @Nullable String ipv4Address,
        @CustomType.Parameter("ipv6Address") @Nullable String ipv6Address,
        @CustomType.Parameter("name") String name) {
        this.aliases = aliases;
        this.ipv4Address = ipv4Address;
        this.ipv6Address = ipv6Address;
        this.name = name;
    }

    /**
     * @return The network aliases of the container in the specific network.
     * 
     */
    public List<String> aliases() {
        return this.aliases == null ? List.of() : this.aliases;
    }
    /**
     * @return The IPV4 address of the container in the specific network.
     * 
     */
    public Optional<String> ipv4Address() {
        return Optional.ofNullable(this.ipv4Address);
    }
    /**
     * @return The IPV6 address of the container in the specific network.
     * 
     */
    public Optional<String> ipv6Address() {
        return Optional.ofNullable(this.ipv6Address);
    }
    /**
     * @return The name of the network.
     * 
     */
    public String name() {
        return this.name;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ContainerNetworksAdvanced defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable List<String> aliases;
        private @Nullable String ipv4Address;
        private @Nullable String ipv6Address;
        private String name;

        public Builder() {
    	      // Empty
        }

        public Builder(ContainerNetworksAdvanced defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.aliases = defaults.aliases;
    	      this.ipv4Address = defaults.ipv4Address;
    	      this.ipv6Address = defaults.ipv6Address;
    	      this.name = defaults.name;
        }

        public Builder aliases(@Nullable List<String> aliases) {
            this.aliases = aliases;
            return this;
        }
        public Builder aliases(String... aliases) {
            return aliases(List.of(aliases));
        }
        public Builder ipv4Address(@Nullable String ipv4Address) {
            this.ipv4Address = ipv4Address;
            return this;
        }
        public Builder ipv6Address(@Nullable String ipv6Address) {
            this.ipv6Address = ipv6Address;
            return this;
        }
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }        public ContainerNetworksAdvanced build() {
            return new ContainerNetworksAdvanced(aliases, ipv4Address, ipv6Address, name);
        }
    }
}
