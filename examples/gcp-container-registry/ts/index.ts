import * as gcp from "@pulumi/gcp";
import * as docker from "@pulumi/docker";

// Create a private GCR registry.
const registry = new gcp.container.Registry("my-registry");
const registryUrl = registry.id.apply(_ =>
    gcp.container.getRegistryRepository().then(reg => reg.repositoryUrl));

// Get registry info (creds and endpoint).
const imageName = registryUrl.apply(url => `${url}/myapp`);

// Build and publish the image.
const image = new docker.Image("my-image", {
    build: {
        context: "app"
    },
    imageName: imageName,
});

// Export the resulting image name
export const fullImageName = image.imageName;
export const repoDigest = image.repoDigest;
