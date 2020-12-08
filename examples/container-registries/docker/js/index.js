var docker = require("@pulumi/docker");
var pulumi = require("@pulumi/pulumi");

// Fetch the Docker Hub auth info from config.
var config = new pulumi.Config();
var username = config.require("dockerUsername");
var password = config.requireSecret("dockerPassword");

// Populate the registry info (creds and endpoint).
var imageName = `${username}/myapp`;
var registryInfo = {
    server: "docker.io",
    username: username,
    password: password,
};

// Build and publish the container image.
var image = new docker.Image("my-image", {
    build: "app",
    imageName: imageName,
    registry: registryInfo,
});

// Export the resuling base name in addition to the specific version pushed.
module.exports = {
    baseImageName: image.baseImageName,
    fullImageName = image.imageName,
};
