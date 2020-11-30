import * as docker from "@pulumi/docker";
import * as pulumi from "@pulumi/pulumi";

// Fetch the Docker Hub auth info from config.
const config = new pulumi.Config();
const username = config.require("dockerUsername");
const password = config.requireSecret("dockerPassword");

// Build and publish the image.
const image = new docker.Image("my-image", {
    build: "app",
    imageName: `${username}/myapp`,
    registry: {
        server: "docker.io",
        username: username,
        password: password,
    },
});
(image.imageName as any).isSecret = false;

// Export the resuling base name in addition to the specific version pushed.
export const imageName = image.imageName;
