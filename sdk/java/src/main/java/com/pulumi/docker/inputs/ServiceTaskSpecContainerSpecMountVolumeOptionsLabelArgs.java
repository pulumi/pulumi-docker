// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class ServiceTaskSpecContainerSpecMountVolumeOptionsLabelArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceTaskSpecContainerSpecMountVolumeOptionsLabelArgs Empty = new ServiceTaskSpecContainerSpecMountVolumeOptionsLabelArgs();

    /**
     * Name of the label
     * 
     */
    @Import(name="label", required=true)
    private Output<String> label;

    /**
     * @return Name of the label
     * 
     */
    public Output<String> label() {
        return this.label;
    }

    /**
     * Value of the label
     * 
     */
    @Import(name="value", required=true)
    private Output<String> value;

    /**
     * @return Value of the label
     * 
     */
    public Output<String> value() {
        return this.value;
    }

    private ServiceTaskSpecContainerSpecMountVolumeOptionsLabelArgs() {}

    private ServiceTaskSpecContainerSpecMountVolumeOptionsLabelArgs(ServiceTaskSpecContainerSpecMountVolumeOptionsLabelArgs $) {
        this.label = $.label;
        this.value = $.value;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceTaskSpecContainerSpecMountVolumeOptionsLabelArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceTaskSpecContainerSpecMountVolumeOptionsLabelArgs $;

        public Builder() {
            $ = new ServiceTaskSpecContainerSpecMountVolumeOptionsLabelArgs();
        }

        public Builder(ServiceTaskSpecContainerSpecMountVolumeOptionsLabelArgs defaults) {
            $ = new ServiceTaskSpecContainerSpecMountVolumeOptionsLabelArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param label Name of the label
         * 
         * @return builder
         * 
         */
        public Builder label(Output<String> label) {
            $.label = label;
            return this;
        }

        /**
         * @param label Name of the label
         * 
         * @return builder
         * 
         */
        public Builder label(String label) {
            return label(Output.of(label));
        }

        /**
         * @param value Value of the label
         * 
         * @return builder
         * 
         */
        public Builder value(Output<String> value) {
            $.value = value;
            return this;
        }

        /**
         * @param value Value of the label
         * 
         * @return builder
         * 
         */
        public Builder value(String value) {
            return value(Output.of(value));
        }

        public ServiceTaskSpecContainerSpecMountVolumeOptionsLabelArgs build() {
            $.label = Objects.requireNonNull($.label, "expected parameter 'label' to be non-null");
            $.value = Objects.requireNonNull($.value, "expected parameter 'value' to be non-null");
            return $;
        }
    }

}
