// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.docker.outputs.ServiceTaskSpecContainerSpecPrivilegesCredentialSpec;
import com.pulumi.docker.outputs.ServiceTaskSpecContainerSpecPrivilegesSeLinuxContext;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ServiceTaskSpecContainerSpecPrivileges {
    /**
     * @return CredentialSpec for managed service account (Windows only)
     * 
     */
    private @Nullable ServiceTaskSpecContainerSpecPrivilegesCredentialSpec credentialSpec;
    /**
     * @return SELinux labels of the container
     * 
     */
    private @Nullable ServiceTaskSpecContainerSpecPrivilegesSeLinuxContext seLinuxContext;

    private ServiceTaskSpecContainerSpecPrivileges() {}
    /**
     * @return CredentialSpec for managed service account (Windows only)
     * 
     */
    public Optional<ServiceTaskSpecContainerSpecPrivilegesCredentialSpec> credentialSpec() {
        return Optional.ofNullable(this.credentialSpec);
    }
    /**
     * @return SELinux labels of the container
     * 
     */
    public Optional<ServiceTaskSpecContainerSpecPrivilegesSeLinuxContext> seLinuxContext() {
        return Optional.ofNullable(this.seLinuxContext);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ServiceTaskSpecContainerSpecPrivileges defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable ServiceTaskSpecContainerSpecPrivilegesCredentialSpec credentialSpec;
        private @Nullable ServiceTaskSpecContainerSpecPrivilegesSeLinuxContext seLinuxContext;
        public Builder() {}
        public Builder(ServiceTaskSpecContainerSpecPrivileges defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.credentialSpec = defaults.credentialSpec;
    	      this.seLinuxContext = defaults.seLinuxContext;
        }

        @CustomType.Setter
        public Builder credentialSpec(@Nullable ServiceTaskSpecContainerSpecPrivilegesCredentialSpec credentialSpec) {

            this.credentialSpec = credentialSpec;
            return this;
        }
        @CustomType.Setter
        public Builder seLinuxContext(@Nullable ServiceTaskSpecContainerSpecPrivilegesSeLinuxContext seLinuxContext) {

            this.seLinuxContext = seLinuxContext;
            return this;
        }
        public ServiceTaskSpecContainerSpecPrivileges build() {
            final var _resultValue = new ServiceTaskSpecContainerSpecPrivileges();
            _resultValue.credentialSpec = credentialSpec;
            _resultValue.seLinuxContext = seLinuxContext;
            return _resultValue;
        }
    }
}
