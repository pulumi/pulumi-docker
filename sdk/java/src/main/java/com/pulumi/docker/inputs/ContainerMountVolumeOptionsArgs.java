// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.docker.inputs.ContainerMountVolumeOptionsLabelArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ContainerMountVolumeOptionsArgs extends com.pulumi.resources.ResourceArgs {

    public static final ContainerMountVolumeOptionsArgs Empty = new ContainerMountVolumeOptionsArgs();

    @Import(name="driverName")
    private @Nullable Output<String> driverName;

    public Optional<Output<String>> driverName() {
        return Optional.ofNullable(this.driverName);
    }

    @Import(name="driverOptions")
    private @Nullable Output<Map<String,String>> driverOptions;

    public Optional<Output<Map<String,String>>> driverOptions() {
        return Optional.ofNullable(this.driverOptions);
    }

    /**
     * User-defined key/value metadata
     * 
     */
    @Import(name="labels")
    private @Nullable Output<List<ContainerMountVolumeOptionsLabelArgs>> labels;

    /**
     * @return User-defined key/value metadata
     * 
     */
    public Optional<Output<List<ContainerMountVolumeOptionsLabelArgs>>> labels() {
        return Optional.ofNullable(this.labels);
    }

    @Import(name="noCopy")
    private @Nullable Output<Boolean> noCopy;

    public Optional<Output<Boolean>> noCopy() {
        return Optional.ofNullable(this.noCopy);
    }

    private ContainerMountVolumeOptionsArgs() {}

    private ContainerMountVolumeOptionsArgs(ContainerMountVolumeOptionsArgs $) {
        this.driverName = $.driverName;
        this.driverOptions = $.driverOptions;
        this.labels = $.labels;
        this.noCopy = $.noCopy;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ContainerMountVolumeOptionsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ContainerMountVolumeOptionsArgs $;

        public Builder() {
            $ = new ContainerMountVolumeOptionsArgs();
        }

        public Builder(ContainerMountVolumeOptionsArgs defaults) {
            $ = new ContainerMountVolumeOptionsArgs(Objects.requireNonNull(defaults));
        }

        public Builder driverName(@Nullable Output<String> driverName) {
            $.driverName = driverName;
            return this;
        }

        public Builder driverName(String driverName) {
            return driverName(Output.of(driverName));
        }

        public Builder driverOptions(@Nullable Output<Map<String,String>> driverOptions) {
            $.driverOptions = driverOptions;
            return this;
        }

        public Builder driverOptions(Map<String,String> driverOptions) {
            return driverOptions(Output.of(driverOptions));
        }

        /**
         * @param labels User-defined key/value metadata
         * 
         * @return builder
         * 
         */
        public Builder labels(@Nullable Output<List<ContainerMountVolumeOptionsLabelArgs>> labels) {
            $.labels = labels;
            return this;
        }

        /**
         * @param labels User-defined key/value metadata
         * 
         * @return builder
         * 
         */
        public Builder labels(List<ContainerMountVolumeOptionsLabelArgs> labels) {
            return labels(Output.of(labels));
        }

        /**
         * @param labels User-defined key/value metadata
         * 
         * @return builder
         * 
         */
        public Builder labels(ContainerMountVolumeOptionsLabelArgs... labels) {
            return labels(List.of(labels));
        }

        public Builder noCopy(@Nullable Output<Boolean> noCopy) {
            $.noCopy = noCopy;
            return this;
        }

        public Builder noCopy(Boolean noCopy) {
            return noCopy(Output.of(noCopy));
        }

        public ContainerMountVolumeOptionsArgs build() {
            return $;
        }
    }

}
