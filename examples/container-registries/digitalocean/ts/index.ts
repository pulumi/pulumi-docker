import * as digitalocean from "@pulumi/digitalocean";
import * as docker from "@pulumi/docker";
import * as pulumi from "@pulumi/pulumi";

// Get the Container Registry
const registry = digitalocean.getContainerRegistry({
    name: "development-pulumi-provider"
})

// Get registry info (creds and endpoint) so we can build/publish to it.
const imageName = pulumi.interpolate`${registry.then( x => x.endpoint)}/myapp`
const creds = new digitalocean.ContainerRegistryDockerCredentials("my-reg-creds", {
    registryName: registry.then(x => x.name),
    write: true,
});
const registryInfo = pulumi.all(
    [creds.dockerCredentials, registry.then( x => x.serverUrl)]
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
    build: "app",
    imageName: imageName,
    registry: registryInfo,
});

// Export the resuling base name in addition to the specific version pushed.
export const baseImageName = image.baseImageName;
export const fullImageName = image.imageName;
