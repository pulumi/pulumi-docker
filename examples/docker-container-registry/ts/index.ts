import * as docker from "@pulumi/docker";
import * as pulumi from "@pulumi/pulumi";

// Fetch the Docker Hub auth info from config.
const config = new pulumi.Config();
const username = config.require("dockerUsername");
const password = config.requireSecret("dockerPassword");

// Populate the registry info (creds and endpoint).
const imageName = `${username}/myapp`;
const registryInfo = {
  server: "docker.io",
  username: username,
  password: password,
};

// Build and publish the container image.
const image = new docker.Image("my-image", {
  build: {
    context: "app",
  },
  imageName: imageName,
  registry: registryInfo,
});

// Export the resulting image name
export const fullImageName = image.imageName;
export const repoDigest = image.repoDigest;

// buildx

const provider = new docker.Provider("docker", {
  registryAuth: [
    {
      address: registryInfo.server,
      username: registryInfo.username,
      password: registryInfo.password,
    },
  ],
});

const buildxImage = new docker.buildx.Image(
  "my-buildx-image",
  {
    tags: [`${imageName}:buildx`],
    exports: ["type=registry"],
    platforms: ["linux/arm64", "linux/amd64"],
    context: "app",
    file: "app/Dockerfile",
  },
  { provider: provider }
);

export const manifests = buildxImage.manifests;
