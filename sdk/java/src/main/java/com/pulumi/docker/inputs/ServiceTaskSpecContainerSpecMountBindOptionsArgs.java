// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceTaskSpecContainerSpecMountBindOptionsArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceTaskSpecContainerSpecMountBindOptionsArgs Empty = new ServiceTaskSpecContainerSpecMountBindOptionsArgs();

    /**
     * Bind propagation refers to whether or not mounts created within a given bind-mount or named volume can be propagated to replicas of that mount. See the [docs](https://docs.docker.com/storage/bind-mounts/#configure-bind-propagation) for details. Defaults to `rprivate`
     * 
     */
    @Import(name="propagation")
    private @Nullable Output<String> propagation;

    /**
     * @return Bind propagation refers to whether or not mounts created within a given bind-mount or named volume can be propagated to replicas of that mount. See the [docs](https://docs.docker.com/storage/bind-mounts/#configure-bind-propagation) for details. Defaults to `rprivate`
     * 
     */
    public Optional<Output<String>> propagation() {
        return Optional.ofNullable(this.propagation);
    }

    private ServiceTaskSpecContainerSpecMountBindOptionsArgs() {}

    private ServiceTaskSpecContainerSpecMountBindOptionsArgs(ServiceTaskSpecContainerSpecMountBindOptionsArgs $) {
        this.propagation = $.propagation;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceTaskSpecContainerSpecMountBindOptionsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceTaskSpecContainerSpecMountBindOptionsArgs $;

        public Builder() {
            $ = new ServiceTaskSpecContainerSpecMountBindOptionsArgs();
        }

        public Builder(ServiceTaskSpecContainerSpecMountBindOptionsArgs defaults) {
            $ = new ServiceTaskSpecContainerSpecMountBindOptionsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param propagation Bind propagation refers to whether or not mounts created within a given bind-mount or named volume can be propagated to replicas of that mount. See the [docs](https://docs.docker.com/storage/bind-mounts/#configure-bind-propagation) for details. Defaults to `rprivate`
         * 
         * @return builder
         * 
         */
        public Builder propagation(@Nullable Output<String> propagation) {
            $.propagation = propagation;
            return this;
        }

        /**
         * @param propagation Bind propagation refers to whether or not mounts created within a given bind-mount or named volume can be propagated to replicas of that mount. See the [docs](https://docs.docker.com/storage/bind-mounts/#configure-bind-propagation) for details. Defaults to `rprivate`
         * 
         * @return builder
         * 
         */
        public Builder propagation(String propagation) {
            return propagation(Output.of(propagation));
        }

        public ServiceTaskSpecContainerSpecMountBindOptionsArgs build() {
            return $;
        }
    }

}
