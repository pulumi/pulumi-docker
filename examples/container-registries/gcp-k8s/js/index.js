var gcp = require("@pulumi/gcp");
var docker = require("@pulumi/docker");
var k8s = require("@pulumi/kubernetes");

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

// Create a load balanced Kubernetes service using this image, and export its IP.
var appLabels = { app: "myapp" };
var appDep = new k8s.apps.v1.Deployment("app-dep", {
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
var appSvc = new k8s.core.v1.Service("app-svc", {
    metadata: { labels: appLabels },
    spec: {
        type: "LoadBalancer",
        ports: [{ port: 80, targetPort: 80, protocol: "TCP" }],
        selector: appLabels,
    },
});

// Export the resuling base name in addition to the specific version pushed.
module.exports = {
    appIp: appSvc.status.loadbalancer.ingress[0].ip,
    baseImageName: image.baseImageName,
    fullImageName: image.imageName,
};
