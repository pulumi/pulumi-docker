import * as aws from "@pulumi/aws";
import * as docker from "@pulumi/docker";
import * as pulumi from "@pulumi/pulumi";
import * as random from "@pulumi/random";

// Create a private ECR registry.
const repo = new aws.ecr.Repository("my-repo", {
  forceDelete: true,
});

// Get registry info (creds and endpoint) so we can build/publish to it.
const registryInfo = repo.registryId.apply(async (id) => {
  const credentials = await aws.ecr.getCredentials({ registryId: id });
  const decodedCredentials = Buffer.from(
    credentials.authorizationToken,
    "base64"
  ).toString();
  const [username, password] = decodedCredentials.split(":");
  if (!password || !username) {
    throw new Error("Invalid credentials");
  }
  return {
    server: credentials.proxyEndpoint,
    username: username,
    password: password,
  };
});

// Build and publish the image.
const image = new docker.Image("my-image", {
  build: {
    context: "app",
    args: {
      // Test for preview with dymanic build property.
      // See https://github.com/pulumi/pulumi-docker/issues/620
      KEY: new random.RandomUuid("guid", {}).result,
    },
  },
  imageName: repo.repositoryUrl,
  registry: registryInfo,
});

// Export the resulting image name
export const imageName = image.baseImageName;
export const repoDigest = image.repoDigest;

// buildx

const buildxImage = new docker.buildx.Image("buildx", {
  tags: [pulumi.interpolate`${repo.repositoryUrl}:buildx`],
  exports: [{ registry: {} }],
  dockerfile: {
    location: "app/Dockerfile",
  },
  platforms: ["linux/arm64", "linux/amd64"],
  cacheTo: [
    {
      registry: {
        mode: "max",
        imageManifest: true,
        ociMediaTypes: true,
        ref: pulumi.interpolate`${repo.repositoryUrl}:cache`,
      },
    },
  ],
  cacheFrom: [
    {
      registry: {
        ref: pulumi.interpolate`${repo.repositoryUrl}:cache`,
      },
    },
  ],
  context: {
    location: "app",
  },
  registries: [
    {
      address: registryInfo.server,
      username: registryInfo.username,
      password: registryInfo.password,
    },
  ],
});
