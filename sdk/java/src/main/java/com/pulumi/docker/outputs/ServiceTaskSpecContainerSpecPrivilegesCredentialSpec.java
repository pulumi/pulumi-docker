// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ServiceTaskSpecContainerSpecPrivilegesCredentialSpec {
    /**
     * @return Load credential spec from this file
     * 
     */
    private @Nullable String file;
    /**
     * @return Load credential spec from this value in the Windows registry
     * 
     */
    private @Nullable String registry;

    private ServiceTaskSpecContainerSpecPrivilegesCredentialSpec() {}
    /**
     * @return Load credential spec from this file
     * 
     */
    public Optional<String> file() {
        return Optional.ofNullable(this.file);
    }
    /**
     * @return Load credential spec from this value in the Windows registry
     * 
     */
    public Optional<String> registry() {
        return Optional.ofNullable(this.registry);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ServiceTaskSpecContainerSpecPrivilegesCredentialSpec defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String file;
        private @Nullable String registry;
        public Builder() {}
        public Builder(ServiceTaskSpecContainerSpecPrivilegesCredentialSpec defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.file = defaults.file;
    	      this.registry = defaults.registry;
        }

        @CustomType.Setter
        public Builder file(@Nullable String file) {

            this.file = file;
            return this;
        }
        @CustomType.Setter
        public Builder registry(@Nullable String registry) {

            this.registry = registry;
            return this;
        }
        public ServiceTaskSpecContainerSpecPrivilegesCredentialSpec build() {
            final var _resultValue = new ServiceTaskSpecContainerSpecPrivilegesCredentialSpec();
            _resultValue.file = file;
            _resultValue.registry = registry;
            return _resultValue;
        }
    }
}
