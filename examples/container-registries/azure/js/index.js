var azure = require("@pulumi/azure");
var azuread = require("@pulumi/azuread");
var docker = require("@pulumi/docker");
var pulumi = require("@pulumi/pulumi");
var random = require("@pulumi/random");

// Conditionalize the auth mechanism.
var config = new pulumi.Config();
var useServicePrincipalAuth = !!config.getBoolean("useServicePrincipalAuth");

// Create a private ACR registry.
var rg = new azure.core.ResourceGroup("myrg");
var registry = new azure.containerservice.Registry("myregistry", {
    resourceGroupName: rg.name,
    adminEnabled: !useServicePrincipalAuth,
    sku: "Basic",
});

// Get registry info (creds and endpoint) so we can build/publish to it.
var imageName = registry.loginServer.apply(s => `${s}/myapp`);
var registryInfo;
if (useServicePrincipalAuth) {
    var sp = new azuread.ServicePrincipal("mysp", {
        applicationId: new azuread.Application("myspapp").applicationId,
    });
    var spPassword = new azuread.ServicePrincipalPassword("mysp-pass", {
        servicePrincipalId: sp.id,
        value: new random.RandomPassword("mypass", {
            length: 32,
        }, { additionalSecretOutputs: [ "result" ] }).result,
        endDateRelative: "8760h",
    });
    var spAuth = new azure.authorization.Assignment("myauth", {
        scope: registry.id,
        roleDefinitionName: "acrpush",
        principalId: sp.id,
    });
    registryInfo = {
        server: registry.loginServer,
        username: sp.applicationId,
        password:  spAuth.id.apply(_ => spPassword.value),
    };
} else {
    registryInfo = {
        server: registry.loginServer,
        username: registry.adminUsername,
        password: registry.adminPassword,
    };
}

// Build and publish the image using the auth information.
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
