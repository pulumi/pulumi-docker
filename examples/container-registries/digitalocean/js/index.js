var digitalocean = require("@pulumi/digitalocean");
var docker = require("@pulumi/docker");
var pulumi = require("@pulumi/pulumi");

// Get the Container Registry
var registry = digitalocean.getContainerRegistry({
    name: "development-pulumi-provider"
})

// Get registry info (creds and endpoint) so we can build/publish to it.
var imageName = pulumi.interpolate`${registry.then( x => x.endpoint)}/myapp`
var creds = new digitalocean.ContainerRegistryDockerCredentials("my-reg-creds", {
    registryName: registry.then(x => x.name),
    write: true,
});
var registryInfo = pulumi.all(
    [creds.dockerCredentials, registry.then(x => x.serverUrl)]
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
