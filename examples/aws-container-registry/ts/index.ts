import * as aws from "@pulumi/aws";
import * as docker from "@pulumi/docker";
import * as random from "@pulumi/random";

// Create a private ECR registry.
const repo = new aws.ecr.Repository("my-repo",{
    forceDelete: true,
});

// Get registry info (creds and endpoint) so we can build/publish to it.
const registryInfo = repo.registryId.apply(async id => {
    const credentials = await aws.ecr.getAuthorizationToken({ registryId: id });
    const decodedCredentials = Buffer.from(credentials.authorizationToken, "base64").toString();
    const [username, password] = decodedCredentials.split(":");
    if (!password || !username) {
        throw new Error("Invalid credentials");
    }
    return {
        server: credentials.proxyEndpoint,
        username: credentials.userName,
        password: credentials.password,
    };
});


// Build and publish the image.
const image = new docker.Image("my-image", {
    build: {
        context: "app",
        args: {
            // Test for preview with dymanic build property.
            // See https://github.com/pulumi/pulumi-docker/issues/620
            KEY: new random.RandomUuid('guid', {}).result,
        },
    },
    imageName: repo.repositoryUrl,
    registry: registryInfo,
});

// Export the resulting image name
export const imageName = image.imageName;
export const repoDigest = image.repoDigest;
