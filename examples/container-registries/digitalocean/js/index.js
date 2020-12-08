var digitalocean = require("@pulumi/digitalocean");
var docker = require("@pulumi/docker");
var pulumi = require("@pulumi/pulumi");

// Create a private DigitalOcean Container Registry.
var registry = new digitalocean.ContainerRegistry("my-reg", {
    subscriptionTierSlug: "starter",
});

// Get registry info (creds and endpoint) so we can build/publish to it.
var imageName = registry.endpoint.apply(s => `${s}/myapp`);
var creds = new digitalocean.ContainerRegistryDockerCredentials("my-reg-creds", {
    registryName: registry.name,
    write: true,
});
var registryInfo = pulumi.all(
    [creds.dockerCredentials, registry.serverUrl]
).apply(([authJson, serverUrl]) => {
    // We are given a Docker creds file; parse it to find the temp username/password.
    var auths = JSON.parse(authJson);
    var authToken = auths["auths"][serverUrl]["auth"];
    var decoded = Buffer.from(authToken, "base64").toString();
    var [username, password] = decoded.split(":");
    if (!password || !username) {
        throw new Error("Invalid credentials");
    }
    return {
        server: serverUrl,
        username: username,
        password: password,
    };
});

// Build and publish the image.
var image = new docker.Image("my-image", {
    build: "app",
    imageName: imageName,
    registry: registryInfo,
});

// Export the resuling base name in addition to the specific version pushed.
module.exports = {
    baseImageName: image.baseImageName,
    fullImageName: image.imageName,
};
