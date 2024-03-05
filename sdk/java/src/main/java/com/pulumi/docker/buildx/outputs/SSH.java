// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker.buildx.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class SSH {
    /**
     * @return Useful for distinguishing different servers that are part of the same
     * build.
     * 
     * A value of `default` is appropriate if only dealing with a single host.
     * 
     */
    private String id;
    /**
     * @return SSH agent socket or private keys to expose to the build under the given
     * identifier.
     * 
     * Defaults to `[$SSH_AUTH_SOCK]`.
     * 
     * Note that your keys are **not** automatically added when using an
     * agent. Run `ssh-add -l` locally to confirm which public keys are
     * visible to the agent; these will be exposed to your build.
     * 
     */
    private @Nullable List<String> paths;

    private SSH() {}
    /**
     * @return Useful for distinguishing different servers that are part of the same
     * build.
     * 
     * A value of `default` is appropriate if only dealing with a single host.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return SSH agent socket or private keys to expose to the build under the given
     * identifier.
     * 
     * Defaults to `[$SSH_AUTH_SOCK]`.
     * 
     * Note that your keys are **not** automatically added when using an
     * agent. Run `ssh-add -l` locally to confirm which public keys are
     * visible to the agent; these will be exposed to your build.
     * 
     */
    public List<String> paths() {
        return this.paths == null ? List.of() : this.paths;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(SSH defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private @Nullable List<String> paths;
        public Builder() {}
        public Builder(SSH defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.paths = defaults.paths;
        }

        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("SSH", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder paths(@Nullable List<String> paths) {

            this.paths = paths;
            return this;
        }
        public Builder paths(String... paths) {
            return paths(List.of(paths));
        }
        public SSH build() {
            final var _resultValue = new SSH();
            _resultValue.id = id;
            _resultValue.paths = paths;
            return _resultValue;
        }
    }
}
