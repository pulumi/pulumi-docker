import * as azure from "@pulumi/azure";
import * as docker from "@pulumi/docker";
import * as pulumi from "@pulumi/pulumi";

// Create a private ACR registry.
const rg = new azure.core.ResourceGroup("myrg")
const registry = new azure.containerservice.Registry("myotherregistry", { // <-- renamed the registry to cause a replacement
    resourceGroupName: rg.name,
    adminEnabled: true,
    sku: "Basic",
});

const imageName = pulumi.interpolate`${registry.loginServer}/myapp`;

// Build and publish the image using the auth information.
const image = new docker.Image("my-image", {
    build: {
        context: "app",
        cacheFrom: {
            images: [imageName]
        },
    },
    imageName: imageName,
    registry: {
        server: registry.loginServer,
        username: registry.adminUsername,
        password: registry.adminPassword,
    }
});

// Export the resulting image name
export const fullImageName = image.imageName;
