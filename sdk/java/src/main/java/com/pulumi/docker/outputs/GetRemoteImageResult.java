// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetRemoteImageResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private final String id;
    /**
     * @return The name of the Docker image, including any tags or SHA256 repo digests.
     * 
     */
    private final String name;
    /**
     * @return The image sha256 digest in the form of `repo[:tag]@sha256:&lt;hash&gt;`. It may be empty in the edge case where the local image was pulled from a repo, tagged locally, and then referred to in the data source by that local name/tag.
     * 
     */
    private final String repoDigest;

    @CustomType.Constructor
    private GetRemoteImageResult(
        @CustomType.Parameter("id") String id,
        @CustomType.Parameter("name") String name,
        @CustomType.Parameter("repoDigest") String repoDigest) {
        this.id = id;
        this.name = name;
        this.repoDigest = repoDigest;
    }

    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The name of the Docker image, including any tags or SHA256 repo digests.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return The image sha256 digest in the form of `repo[:tag]@sha256:&lt;hash&gt;`. It may be empty in the edge case where the local image was pulled from a repo, tagged locally, and then referred to in the data source by that local name/tag.
     * 
     */
    public String repoDigest() {
        return this.repoDigest;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRemoteImageResult defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private String id;
        private String name;
        private String repoDigest;

        public Builder() {
    	      // Empty
        }

        public Builder(GetRemoteImageResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.repoDigest = defaults.repoDigest;
        }

        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        public Builder repoDigest(String repoDigest) {
            this.repoDigest = Objects.requireNonNull(repoDigest);
            return this;
        }        public GetRemoteImageResult build() {
            return new GetRemoteImageResult(id, name, repoDigest);
        }
    }
}