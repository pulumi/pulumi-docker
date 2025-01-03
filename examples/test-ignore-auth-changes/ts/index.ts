import * as pulumi from "@pulumi/pulumi";
import * as docker from "@pulumi/docker";

const config = new pulumi.Config();

// Changes to username and password should not trigger a diff, but changes to address should.
const provider = new docker.Provider("docker", {
  registryAuth: [
    {
      address: config.require("address"),
      username: config.require("username"),
      password: config.require("password"),
    },
  ],
});
