import * as azure from "@pulumi/azure";
import * as azuread from "@pulumi/azuread";
import * as docker from "@pulumi/docker";
import * as pulumi from "@pulumi/pulumi";
import * as random from "@pulumi/random";

// Conditionalize the auth mechanism.
const config = new pulumi.Config();
const useServicePrincipalAuth = !!config.getBoolean("useServicePrincipalAuth");

// Create a private ACR registry.
const rg = new azure.core.ResourceGroup("myrg");
const registry = new azure.containerservice.Registry("myregistry", {
    resourceGroupName: rg.name,
    adminEnabled: !useServicePrincipalAuth,
    sku: "Basic",
});

// Get registry info (creds and endpoint).
const imageName = registry.loginServer.apply(s => `${s}/myapp`);
let registryInfo: docker.ImageRegistry;
if (useServicePrincipalAuth) {
    const sp = new azuread.ServicePrincipal("mysp", {
        applicationId: new azuread.Application("myspapp").applicationId,
    });
    const spPassword = new azuread.ServicePrincipalPassword("mysp-pass", {
        servicePrincipalId: sp.id,
        value: new random.RandomPassword("mypass", {
            length: 32,
        }, { additionalSecretOutputs: [ "result" ] }).result,
        endDateRelative: "8760h",
    });
    const spAuth = new azure.authorization.Assignment("myauth", {
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
const image = new docker.Image("my-image", {
    build: "app",
    imageName: imageName,
    registry: registryInfo,
});

// Export the resuling base name in addition to the specific version pushed.
export const baseImageName = image.baseImageName;
export const fullImageName = image.imageName;
