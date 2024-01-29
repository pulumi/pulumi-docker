// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class ServiceTaskSpecContainerSpecMountVolumeOptionsLabelArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceTaskSpecContainerSpecMountVolumeOptionsLabelArgs Empty = new ServiceTaskSpecContainerSpecMountVolumeOptionsLabelArgs();

    @Import(name="label", required=true)
    private Output<String> label;

    public Output<String> label() {
        return this.label;
    }

    @Import(name="value", required=true)
    private Output<String> value;

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

        public Builder label(Output<String> label) {
            $.label = label;
            return this;
        }

        public Builder label(String label) {
            return label(Output.of(label));
        }

        public Builder value(Output<String> value) {
            $.value = value;
            return this;
        }

        public Builder value(String value) {
            return value(Output.of(value));
        }

        public ServiceTaskSpecContainerSpecMountVolumeOptionsLabelArgs build() {
            if ($.label == null) {
                throw new MissingRequiredPropertyException("ServiceTaskSpecContainerSpecMountVolumeOptionsLabelArgs", "label");
            }
            if ($.value == null) {
                throw new MissingRequiredPropertyException("ServiceTaskSpecContainerSpecMountVolumeOptionsLabelArgs", "value");
            }
            return $;
        }
    }

}
