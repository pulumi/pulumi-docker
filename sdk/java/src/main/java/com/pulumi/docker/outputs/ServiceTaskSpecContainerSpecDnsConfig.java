// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class ServiceTaskSpecContainerSpecDnsConfig {
    /**
     * @return The IP addresses of the name servers
     * 
     */
    private List<String> nameservers;
    /**
     * @return A list of internal resolver variables to be modified (e.g., `debug`, `ndots:3`, etc.)
     * 
     */
    private @Nullable List<String> options;
    /**
     * @return A search list for host-name lookup
     * 
     */
    private @Nullable List<String> searches;

    private ServiceTaskSpecContainerSpecDnsConfig() {}
    /**
     * @return The IP addresses of the name servers
     * 
     */
    public List<String> nameservers() {
        return this.nameservers;
    }
    /**
     * @return A list of internal resolver variables to be modified (e.g., `debug`, `ndots:3`, etc.)
     * 
     */
    public List<String> options() {
        return this.options == null ? List.of() : this.options;
    }
    /**
     * @return A search list for host-name lookup
     * 
     */
    public List<String> searches() {
        return this.searches == null ? List.of() : this.searches;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ServiceTaskSpecContainerSpecDnsConfig defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<String> nameservers;
        private @Nullable List<String> options;
        private @Nullable List<String> searches;
        public Builder() {}
        public Builder(ServiceTaskSpecContainerSpecDnsConfig defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.nameservers = defaults.nameservers;
    	      this.options = defaults.options;
    	      this.searches = defaults.searches;
        }

        @CustomType.Setter
        public Builder nameservers(List<String> nameservers) {
            if (nameservers == null) {
              throw new MissingRequiredPropertyException("ServiceTaskSpecContainerSpecDnsConfig", "nameservers");
            }
            this.nameservers = nameservers;
            return this;
        }
        public Builder nameservers(String... nameservers) {
            return nameservers(List.of(nameservers));
        }
        @CustomType.Setter
        public Builder options(@Nullable List<String> options) {

            this.options = options;
            return this;
        }
        public Builder options(String... options) {
            return options(List.of(options));
        }
        @CustomType.Setter
        public Builder searches(@Nullable List<String> searches) {

            this.searches = searches;
            return this;
        }
        public Builder searches(String... searches) {
            return searches(List.of(searches));
        }
        public ServiceTaskSpecContainerSpecDnsConfig build() {
            final var _resultValue = new ServiceTaskSpecContainerSpecDnsConfig();
            _resultValue.nameservers = nameservers;
            _resultValue.options = options;
            _resultValue.searches = searches;
            return _resultValue;
        }
    }
}
