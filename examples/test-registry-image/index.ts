import * as aws from "@pulumi/aws";
import * as docker from "@pulumi/docker";

// Create a private ECR registry.
const repo = new aws.ecr.Repository("my-repo",{
    forceDelete: true,
});

// Get registry info (creds and endpoint) so we can build/publish to it.
const registryInfo = repo.registryId.apply(async id => {
    const credentials = await aws.ecr.getCredentials({ registryId: id });
    const decodedCredentials = Buffer.from(credentials.authorizationToken, "base64").toString();
    const [username, password] = decodedCredentials.split(":");
    if (!password || !username) {
        throw new Error("Invalid credentials");
    }
    return {
        address: credentials.proxyEndpoint,
        username: username,
        password: password,
    };
});


// Build and publish the image.
const image = new docker.Image("my-image", {
    build: {
        context: "app",
    },
    imageName: repo.repositoryUrl,
    skipPush: true
});

const ecrProvider = new docker.Provider("ecr-provider", {
    registryAuth: [registryInfo],
},
);

// Publish the image to the registry
const registryImage = new docker.RegistryImage("my-registry-image",
    {
        name: repo.repositoryUrl,
    },
    { provider: ecrProvider, dependsOn: [image] },
);

// Export the resulting image name
export const imageName = registryImage.name;
export const repoDigest = registryImage.sha256Digest;