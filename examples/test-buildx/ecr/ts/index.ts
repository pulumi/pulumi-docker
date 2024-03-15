import * as aws from "@pulumi/aws";
import * as docker from "@pulumi/docker";
import * as pulumi from "@pulumi/pulumi";

// Create a private ECR registry WITHOUT force delete -- our image should be
// deleted during teardown.
const repo = new aws.ecr.Repository("my-repo", {});

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
const image = new docker.buildx.Image("buildx-pushed-multi-plat", {
  tags: [pulumi.interpolate`${repo.repositoryUrl}:buildx`],
  push: true,
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

const notPushed = new docker.buildx.Image("buildx-not-pushed", {
  tags: [pulumi.interpolate`${repo.repositoryUrl}:buildx`],
  platforms: ["linux/arm64", "linux/amd64"],
  push: false,
  context: {
    location: "app",
  },
});

// Export the resulting image name
export const ref = image.ref;
