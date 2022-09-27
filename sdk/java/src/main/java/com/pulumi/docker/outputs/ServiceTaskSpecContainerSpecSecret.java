// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ServiceTaskSpecContainerSpecSecret {
    private @Nullable String fileGid;
    private @Nullable Integer fileMode;
    private String fileName;
    private @Nullable String fileUid;
    private String secretId;
    private @Nullable String secretName;

    private ServiceTaskSpecContainerSpecSecret() {}
    public Optional<String> fileGid() {
        return Optional.ofNullable(this.fileGid);
    }
    public Optional<Integer> fileMode() {
        return Optional.ofNullable(this.fileMode);
    }
    public String fileName() {
        return this.fileName;
    }
    public Optional<String> fileUid() {
        return Optional.ofNullable(this.fileUid);
    }
    public String secretId() {
        return this.secretId;
    }
    public Optional<String> secretName() {
        return Optional.ofNullable(this.secretName);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ServiceTaskSpecContainerSpecSecret defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String fileGid;
        private @Nullable Integer fileMode;
        private String fileName;
        private @Nullable String fileUid;
        private String secretId;
        private @Nullable String secretName;
        public Builder() {}
        public Builder(ServiceTaskSpecContainerSpecSecret defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.fileGid = defaults.fileGid;
    	      this.fileMode = defaults.fileMode;
    	      this.fileName = defaults.fileName;
    	      this.fileUid = defaults.fileUid;
    	      this.secretId = defaults.secretId;
    	      this.secretName = defaults.secretName;
        }

        @CustomType.Setter
        public Builder fileGid(@Nullable String fileGid) {
            this.fileGid = fileGid;
            return this;
        }
        @CustomType.Setter
        public Builder fileMode(@Nullable Integer fileMode) {
            this.fileMode = fileMode;
            return this;
        }
        @CustomType.Setter
        public Builder fileName(String fileName) {
            this.fileName = Objects.requireNonNull(fileName);
            return this;
        }
        @CustomType.Setter
        public Builder fileUid(@Nullable String fileUid) {
            this.fileUid = fileUid;
            return this;
        }
        @CustomType.Setter
        public Builder secretId(String secretId) {
            this.secretId = Objects.requireNonNull(secretId);
            return this;
        }
        @CustomType.Setter
        public Builder secretName(@Nullable String secretName) {
            this.secretName = secretName;
            return this;
        }
        public ServiceTaskSpecContainerSpecSecret build() {
            final var o = new ServiceTaskSpecContainerSpecSecret();
            o.fileGid = fileGid;
            o.fileMode = fileMode;
            o.fileName = fileName;
            o.fileUid = fileUid;
            o.secretId = secretId;
            o.secretName = secretName;
            return o;
        }
    }
}
