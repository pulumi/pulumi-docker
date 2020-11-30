var gcp = require("@pulumi/gcp");
var docker = require("@pulumi/docker");

// Create a private GCR registry.
var registry = new gcp.container.Registry("my-registry");
var registryUrl = registry.id.apply(_ =>
    gcp.container.getRegistryRepository().then(reg => reg.repositoryUrl));

// Get registry info (creds and endpoint).
var imageName = registryUrl.apply(url => `${url}/myapp`);
var registryInfo = undefined; // use gcloud for authentication.

// Build and publish the image.
var image = new docker.Image("my-image", {
    build: "./app",
    imageName: imageName,
    registry: registryInfo,
});

// Export the resuling base name in addition to the specific version pushed.
module.exports = {
    baseImageName: image.baseImageName,
    fullImageName: image.imageName,
};
