// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.docker.inputs.ServiceTaskSpecContainerSpecMountVolumeOptionsLabelArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceTaskSpecContainerSpecMountVolumeOptionsArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceTaskSpecContainerSpecMountVolumeOptionsArgs Empty = new ServiceTaskSpecContainerSpecMountVolumeOptionsArgs();

    /**
     * Name of the driver to use to create the volume
     * 
     */
    @Import(name="driverName")
    private @Nullable Output<String> driverName;

    /**
     * @return Name of the driver to use to create the volume
     * 
     */
    public Optional<Output<String>> driverName() {
        return Optional.ofNullable(this.driverName);
    }

    /**
     * key/value map of driver specific options
     * 
     */
    @Import(name="driverOptions")
    private @Nullable Output<Map<String,String>> driverOptions;

    /**
     * @return key/value map of driver specific options
     * 
     */
    public Optional<Output<Map<String,String>>> driverOptions() {
        return Optional.ofNullable(this.driverOptions);
    }

    /**
     * User-defined key/value metadata
     * 
     */
    @Import(name="labels")
    private @Nullable Output<List<ServiceTaskSpecContainerSpecMountVolumeOptionsLabelArgs>> labels;

    /**
     * @return User-defined key/value metadata
     * 
     */
    public Optional<Output<List<ServiceTaskSpecContainerSpecMountVolumeOptionsLabelArgs>>> labels() {
        return Optional.ofNullable(this.labels);
    }

    /**
     * Populate volume with data from the target
     * 
     */
    @Import(name="noCopy")
    private @Nullable Output<Boolean> noCopy;

    /**
     * @return Populate volume with data from the target
     * 
     */
    public Optional<Output<Boolean>> noCopy() {
        return Optional.ofNullable(this.noCopy);
    }

    private ServiceTaskSpecContainerSpecMountVolumeOptionsArgs() {}

    private ServiceTaskSpecContainerSpecMountVolumeOptionsArgs(ServiceTaskSpecContainerSpecMountVolumeOptionsArgs $) {
        this.driverName = $.driverName;
        this.driverOptions = $.driverOptions;
        this.labels = $.labels;
        this.noCopy = $.noCopy;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceTaskSpecContainerSpecMountVolumeOptionsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceTaskSpecContainerSpecMountVolumeOptionsArgs $;

        public Builder() {
            $ = new ServiceTaskSpecContainerSpecMountVolumeOptionsArgs();
        }

        public Builder(ServiceTaskSpecContainerSpecMountVolumeOptionsArgs defaults) {
            $ = new ServiceTaskSpecContainerSpecMountVolumeOptionsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param driverName Name of the driver to use to create the volume
         * 
         * @return builder
         * 
         */
        public Builder driverName(@Nullable Output<String> driverName) {
            $.driverName = driverName;
            return this;
        }

        /**
         * @param driverName Name of the driver to use to create the volume
         * 
         * @return builder
         * 
         */
        public Builder driverName(String driverName) {
            return driverName(Output.of(driverName));
        }

        /**
         * @param driverOptions key/value map of driver specific options
         * 
         * @return builder
         * 
         */
        public Builder driverOptions(@Nullable Output<Map<String,String>> driverOptions) {
            $.driverOptions = driverOptions;
            return this;
        }

        /**
         * @param driverOptions key/value map of driver specific options
         * 
         * @return builder
         * 
         */
        public Builder driverOptions(Map<String,String> driverOptions) {
            return driverOptions(Output.of(driverOptions));
        }

        /**
         * @param labels User-defined key/value metadata
         * 
         * @return builder
         * 
         */
        public Builder labels(@Nullable Output<List<ServiceTaskSpecContainerSpecMountVolumeOptionsLabelArgs>> labels) {
            $.labels = labels;
            return this;
        }

        /**
         * @param labels User-defined key/value metadata
         * 
         * @return builder
         * 
         */
        public Builder labels(List<ServiceTaskSpecContainerSpecMountVolumeOptionsLabelArgs> labels) {
            return labels(Output.of(labels));
        }

        /**
         * @param labels User-defined key/value metadata
         * 
         * @return builder
         * 
         */
        public Builder labels(ServiceTaskSpecContainerSpecMountVolumeOptionsLabelArgs... labels) {
            return labels(List.of(labels));
        }

        /**
         * @param noCopy Populate volume with data from the target
         * 
         * @return builder
         * 
         */
        public Builder noCopy(@Nullable Output<Boolean> noCopy) {
            $.noCopy = noCopy;
            return this;
        }

        /**
         * @param noCopy Populate volume with data from the target
         * 
         * @return builder
         * 
         */
        public Builder noCopy(Boolean noCopy) {
            return noCopy(Output.of(noCopy));
        }

        public ServiceTaskSpecContainerSpecMountVolumeOptionsArgs build() {
            return $;
        }
    }

}
