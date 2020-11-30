import * as gcp from "@pulumi/gcp";
import * as docker from "@pulumi/docker";

// Create a private GCR registry.
const registry = new gcp.container.Registry("my-registry");
const registryUrl = registry.id.apply(_ =>
    gcp.container.getRegistryRepository().then(reg => reg.repositoryUrl));

// Get registry info (creds and endpoint).
const imageName = registryUrl.apply(url => `${url}/myapp`);
const registryInfo = undefined; // use gcloud for authentication.

// Build and publish the image.
const image = new docker.Image("my-image", {
    build: "./app",
    imageName: imageName,
    registry: registryInfo,
});

// Export the resuling base name in addition to the specific version pushed.
export const baseImageName = image.baseImageName;
export const fullImageName = image.imageName;
