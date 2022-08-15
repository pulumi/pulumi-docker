// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetRegistryImageArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetRegistryImageArgs Empty = new GetRegistryImageArgs();

    /**
     * If `true`, the verification of TLS certificates of the server/registry is disabled. Defaults to `false`
     * 
     */
    @Import(name="insecureSkipVerify")
    private @Nullable Output<Boolean> insecureSkipVerify;

    /**
     * @return If `true`, the verification of TLS certificates of the server/registry is disabled. Defaults to `false`
     * 
     */
    public Optional<Output<Boolean>> insecureSkipVerify() {
        return Optional.ofNullable(this.insecureSkipVerify);
    }

    /**
     * The name of the Docker image, including any tags. e.g. `alpine:latest`
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return The name of the Docker image, including any tags. e.g. `alpine:latest`
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    private GetRegistryImageArgs() {}

    private GetRegistryImageArgs(GetRegistryImageArgs $) {
        this.insecureSkipVerify = $.insecureSkipVerify;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetRegistryImageArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetRegistryImageArgs $;

        public Builder() {
            $ = new GetRegistryImageArgs();
        }

        public Builder(GetRegistryImageArgs defaults) {
            $ = new GetRegistryImageArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param insecureSkipVerify If `true`, the verification of TLS certificates of the server/registry is disabled. Defaults to `false`
         * 
         * @return builder
         * 
         */
        public Builder insecureSkipVerify(@Nullable Output<Boolean> insecureSkipVerify) {
            $.insecureSkipVerify = insecureSkipVerify;
            return this;
        }

        /**
         * @param insecureSkipVerify If `true`, the verification of TLS certificates of the server/registry is disabled. Defaults to `false`
         * 
         * @return builder
         * 
         */
        public Builder insecureSkipVerify(Boolean insecureSkipVerify) {
            return insecureSkipVerify(Output.of(insecureSkipVerify));
        }

        /**
         * @param name The name of the Docker image, including any tags. e.g. `alpine:latest`
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the Docker image, including any tags. e.g. `alpine:latest`
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public GetRegistryImageArgs build() {
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            return $;
        }
    }

}
