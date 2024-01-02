// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class ContainerLabelArgs extends com.pulumi.resources.ResourceArgs {

    public static final ContainerLabelArgs Empty = new ContainerLabelArgs();

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

    private ContainerLabelArgs() {}

    private ContainerLabelArgs(ContainerLabelArgs $) {
        this.label = $.label;
        this.value = $.value;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ContainerLabelArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ContainerLabelArgs $;

        public Builder() {
            $ = new ContainerLabelArgs();
        }

        public Builder(ContainerLabelArgs defaults) {
            $ = new ContainerLabelArgs(Objects.requireNonNull(defaults));
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

        public ContainerLabelArgs build() {
            if ($.label == null) {
                throw new MissingRequiredPropertyException("ContainerLabelArgs", "label");
            }
            if ($.value == null) {
                throw new MissingRequiredPropertyException("ContainerLabelArgs", "value");
            }
            return $;
        }
    }

}
