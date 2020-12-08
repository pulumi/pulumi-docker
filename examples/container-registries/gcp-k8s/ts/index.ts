import * as gcp from "@pulumi/gcp";
import * as docker from "@pulumi/docker";
import * as k8s from "@pulumi/kubernetes";

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

// Create a load balanced Kubernetes service using this image, and export its IP.
const appLabels = { app: "myapp" };
const appDep = new k8s.apps.v1.Deployment("app-dep", {
    spec: {
        selector: { matchLabels: appLabels },
        replicas: 3,
        template: {
            metadata: { labels: appLabels },
            spec: {
                containers: [{
                    name: "myapp",
                    image: image.imageName,
                }],
            },
        },
    },
});
const appSvc = new k8s.core.v1.Service("app-svc", {
    metadata: { labels: appLabels },
    spec: {
        type: "LoadBalancer",
        ports: [{ port: 80, targetPort: 80, protocol: "TCP" }],
        selector: appLabels,
    },
});
export const appIp = appSvc.status.loadBalancer.ingress[0].ip;
