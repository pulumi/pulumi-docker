import * as digitalocean from "@pulumi/digitalocean";
import * as docker from "@pulumi/docker";
import * as pulumi from "@pulumi/pulumi";

// Create a private DigitalOcean Container Registry.
const registry = new digitalocean.ContainerRegistry("my-reg", {
    subscriptionTierSlug: "starter",
}, { ignoreChanges: ["storageUsageBytes"] });

// Get registry info (creds and endpoint) so we can build/publish to it.
const creds = new digitalocean.ContainerRegistryDockerCredentials("my-reg-creds", {
    registryName: registry.name,
    write: true,
});
const registryInfo = pulumi.all(
    [creds.dockerCredentials, registry.serverUrl]
).apply(([authJson, serverUrl]) => {
    // We are given a Docker creds file; parse it to find the temp username/password.
    const auths = JSON.parse(authJson);
    const authToken = auths["auths"][serverUrl]["auth"];
    const decoded = Buffer.from(authToken, "base64").toString();
    const [username, password] = decoded.split(":");
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
const image = new docker.Image("my-image", {
    build: {
        context: "app"
    },
    imageName: registry.endpoint.apply(s => `${s}/myapp`),
    registry: registryInfo,
});

// Export the resuling base name in addition to the specific version pushed.
export const baseImageName = image.baseImageName;
export const fullImageName = image.imageName;
