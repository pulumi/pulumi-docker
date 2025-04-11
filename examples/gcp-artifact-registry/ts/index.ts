import * as gcp from "@pulumi/gcp";
import * as pulumi from "@pulumi/pulumi";
import * as random from "@pulumi/random";
import * as docker from "@pulumi/docker";

// Create a random suffix for the repository name.
let randomSuffix = new random.RandomString("random-suffix", {
    length: 6,
    special: false,
    upper: false
});
let repoName = pulumi.concat(`docker-test-repo-`, randomSuffix.result);

// Create a private GCP artifact registry.
const registry = new gcp.artifactregistry.Repository("my-registry", {
    format: "DOCKER",
    repositoryId: repoName,
    location: "us-central1", // change to your desired region
});

const registryUrl = pulumi.interpolate `${registry.location}-docker.pkg.dev/${registry.project}/${registry.repositoryId}`;

// Get registry info (creds and endpoint).
const imageName = pulumi.interpolate `${registryUrl}/myapp`;

// Build and publish the image.
const image = new docker.Image("my-image", {
    build: {
        context: "app"
    },
    imageName: imageName,
});

// Export the resulting image name
export const fullImageName = image.imageName;
export const repoDigest = image.repoDigest;
