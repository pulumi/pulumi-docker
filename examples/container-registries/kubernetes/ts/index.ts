import * as docker from "@pulumi/docker";
import * as k8s from "@pulumi/kubernetes";
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

// Ensure we can pull from the Docker Hub.
const pullSecret = new k8s.core.v1.Secret("my-regcred", {
    type: "kubernetes.io/dockerconfigjson",
    stringData: {
        ".dockerconfigjson": password.apply(password => JSON.stringify({
            auths: {
                "https://index.docker.io/v1/": {
                    username,
                    password,
                    auth: Buffer.from(`${username}:${password}`).toString("base64"),
                },
            },
        })),
    },
});

// Deploy a load-balanced service that uses this image.
const labels = { app: "my-app" };
const dep = new k8s.apps.v1.Deployment("my-app-dep", {
    spec: {
        selector: { matchLabels: labels },
        replicas: 1,
        template: {
            metadata: { labels: labels },
            spec: {
                containers: [{
                    name: labels.app,
                    image: image.imageName,
                }],
                imagePullSecrets: [{
                    name: pullSecret.metadata.name,
                }],
            },
        },
    },
});
const svc = new k8s.core.v1.Service("my-app-svc", {
    spec: {
        selector: labels,
        type: "LoadBalancer",
        ports: [{ port: 80 }],
    },
});

// Export the resuling base name in addition to the specific version pushed.
export const imageName = image.imageName;
// Export the k8s ingress IP to access the canary deployment
export const serviceIp = svc.status.loadBalancer.ingress[0].ip;
