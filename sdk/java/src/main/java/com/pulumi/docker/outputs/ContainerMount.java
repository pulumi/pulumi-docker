// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.docker.outputs.ContainerMountBindOptions;
import com.pulumi.docker.outputs.ContainerMountTmpfsOptions;
import com.pulumi.docker.outputs.ContainerMountVolumeOptions;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ContainerMount {
    /**
     * @return Optional configuration for the bind type.
     * 
     */
    private @Nullable ContainerMountBindOptions bindOptions;
    /**
     * @return Whether the mount should be read-only.
     * 
     */
    private @Nullable Boolean readOnly;
    /**
     * @return Mount source (e.g. a volume name, a host path).
     * 
     */
    private @Nullable String source;
    /**
     * @return Container path
     * 
     */
    private String target;
    /**
     * @return Optional configuration for the tmpfs type.
     * 
     */
    private @Nullable ContainerMountTmpfsOptions tmpfsOptions;
    /**
     * @return The mount type
     * 
     */
    private String type;
    /**
     * @return Optional configuration for the volume type.
     * 
     */
    private @Nullable ContainerMountVolumeOptions volumeOptions;

    private ContainerMount() {}
    /**
     * @return Optional configuration for the bind type.
     * 
     */
    public Optional<ContainerMountBindOptions> bindOptions() {
        return Optional.ofNullable(this.bindOptions);
    }
    /**
     * @return Whether the mount should be read-only.
     * 
     */
    public Optional<Boolean> readOnly() {
        return Optional.ofNullable(this.readOnly);
    }
    /**
     * @return Mount source (e.g. a volume name, a host path).
     * 
     */
    public Optional<String> source() {
        return Optional.ofNullable(this.source);
    }
    /**
     * @return Container path
     * 
     */
    public String target() {
        return this.target;
    }
    /**
     * @return Optional configuration for the tmpfs type.
     * 
     */
    public Optional<ContainerMountTmpfsOptions> tmpfsOptions() {
        return Optional.ofNullable(this.tmpfsOptions);
    }
    /**
     * @return The mount type
     * 
     */
    public String type() {
        return this.type;
    }
    /**
     * @return Optional configuration for the volume type.
     * 
     */
    public Optional<ContainerMountVolumeOptions> volumeOptions() {
        return Optional.ofNullable(this.volumeOptions);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ContainerMount defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable ContainerMountBindOptions bindOptions;
        private @Nullable Boolean readOnly;
        private @Nullable String source;
        private String target;
        private @Nullable ContainerMountTmpfsOptions tmpfsOptions;
        private String type;
        private @Nullable ContainerMountVolumeOptions volumeOptions;
        public Builder() {}
        public Builder(ContainerMount defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.bindOptions = defaults.bindOptions;
    	      this.readOnly = defaults.readOnly;
    	      this.source = defaults.source;
    	      this.target = defaults.target;
    	      this.tmpfsOptions = defaults.tmpfsOptions;
    	      this.type = defaults.type;
    	      this.volumeOptions = defaults.volumeOptions;
        }

        @CustomType.Setter
        public Builder bindOptions(@Nullable ContainerMountBindOptions bindOptions) {
            this.bindOptions = bindOptions;
            return this;
        }
        @CustomType.Setter
        public Builder readOnly(@Nullable Boolean readOnly) {
            this.readOnly = readOnly;
            return this;
        }
        @CustomType.Setter
        public Builder source(@Nullable String source) {
            this.source = source;
            return this;
        }
        @CustomType.Setter
        public Builder target(String target) {
            this.target = Objects.requireNonNull(target);
            return this;
        }
        @CustomType.Setter
        public Builder tmpfsOptions(@Nullable ContainerMountTmpfsOptions tmpfsOptions) {
            this.tmpfsOptions = tmpfsOptions;
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            this.type = Objects.requireNonNull(type);
            return this;
        }
        @CustomType.Setter
        public Builder volumeOptions(@Nullable ContainerMountVolumeOptions volumeOptions) {
            this.volumeOptions = volumeOptions;
            return this;
        }
        public ContainerMount build() {
            final var _resultValue = new ContainerMount();
            _resultValue.bindOptions = bindOptions;
            _resultValue.readOnly = readOnly;
            _resultValue.source = source;
            _resultValue.target = target;
            _resultValue.tmpfsOptions = tmpfsOptions;
            _resultValue.type = type;
            _resultValue.volumeOptions = volumeOptions;
            return _resultValue;
        }
    }
}
