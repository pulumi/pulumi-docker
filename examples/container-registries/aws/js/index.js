var aws = require("@pulumi/aws");
var docker = require("@pulumi/docker");

// Create a private ECR registry.
var repo = new aws.ecr.Repository("my-repo");

// Get registry info (creds and endpoint) so we can build/publish to it.
var imageName = repo.repositoryUrl;
var registryInfo = repo.registryId.apply(id => {
    return aws.ecr.getCredentials({ registryId: id }).then(credentials => {
        var decodedCredentials = Buffer.from(credentials.authorizationToken, "base64").toString();
        var [username, password] = decodedCredentials.split(":");
        if (!password || !username) {
            throw new Error("Invalid credentials");
        }
        return {
            server: credentials.proxyEndpoint,
            username: username,
            password: password,
        };
    });
});

// Build and publish the image.
var image = new docker.Image("my-image", {
    build: "app",
    imageName: imageName,
    registry: registryInfo,
});

// Export the resuling base name in addition to the specific version pushed.
module.exports = {
    baseImageName: image.baseImageName,
    fullImageName: image.imageName,
};
