// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ServiceTaskSpecContainerSpecPrivilegesSeLinuxContext {
    /**
     * @return Disable SELinux
     * 
     */
    private @Nullable Boolean disable;
    /**
     * @return SELinux level label
     * 
     */
    private @Nullable String level;
    /**
     * @return SELinux role label
     * 
     */
    private @Nullable String role;
    /**
     * @return SELinux type label
     * 
     */
    private @Nullable String type;
    /**
     * @return SELinux user label
     * 
     */
    private @Nullable String user;

    private ServiceTaskSpecContainerSpecPrivilegesSeLinuxContext() {}
    /**
     * @return Disable SELinux
     * 
     */
    public Optional<Boolean> disable() {
        return Optional.ofNullable(this.disable);
    }
    /**
     * @return SELinux level label
     * 
     */
    public Optional<String> level() {
        return Optional.ofNullable(this.level);
    }
    /**
     * @return SELinux role label
     * 
     */
    public Optional<String> role() {
        return Optional.ofNullable(this.role);
    }
    /**
     * @return SELinux type label
     * 
     */
    public Optional<String> type() {
        return Optional.ofNullable(this.type);
    }
    /**
     * @return SELinux user label
     * 
     */
    public Optional<String> user() {
        return Optional.ofNullable(this.user);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ServiceTaskSpecContainerSpecPrivilegesSeLinuxContext defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean disable;
        private @Nullable String level;
        private @Nullable String role;
        private @Nullable String type;
        private @Nullable String user;
        public Builder() {}
        public Builder(ServiceTaskSpecContainerSpecPrivilegesSeLinuxContext defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.disable = defaults.disable;
    	      this.level = defaults.level;
    	      this.role = defaults.role;
    	      this.type = defaults.type;
    	      this.user = defaults.user;
        }

        @CustomType.Setter
        public Builder disable(@Nullable Boolean disable) {

            this.disable = disable;
            return this;
        }
        @CustomType.Setter
        public Builder level(@Nullable String level) {

            this.level = level;
            return this;
        }
        @CustomType.Setter
        public Builder role(@Nullable String role) {

            this.role = role;
            return this;
        }
        @CustomType.Setter
        public Builder type(@Nullable String type) {

            this.type = type;
            return this;
        }
        @CustomType.Setter
        public Builder user(@Nullable String user) {

            this.user = user;
            return this;
        }
        public ServiceTaskSpecContainerSpecPrivilegesSeLinuxContext build() {
            final var _resultValue = new ServiceTaskSpecContainerSpecPrivilegesSeLinuxContext();
            _resultValue.disable = disable;
            _resultValue.level = level;
            _resultValue.role = role;
            _resultValue.type = type;
            _resultValue.user = user;
            return _resultValue;
        }
    }
}
