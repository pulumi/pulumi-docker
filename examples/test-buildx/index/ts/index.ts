import * as docker from "@pulumi/docker";
import * as pulumi from "@pulumi/pulumi";

const config = new pulumi.Config();

const provider = new docker.Provider("provider", {
  registryAuth: [
    {
      address: config.getSecret("address"),
      username: config.getSecret("username"),
      password: config.getSecret("password"),
    },
  ],
});

const dockerfile = {
  inline: `FROM scratch`,
};
const context = {
  location: ".",
};

const image1 = new docker.buildx.Image(
  `image1`,
  {
    tags: [`${config.require("tag")}-image1`],
    dockerfile: dockerfile,
    context: context,
    push: true,
  },
  { provider: provider }
);

const image2 = new docker.buildx.Image(
  `image2`,
  {
    tags: [`${config.require("tag")}-image2`],
    dockerfile: dockerfile,
    context: context,
    push: true,
  },
  { provider: provider }
);

new docker.buildx.Index(
  "index",
  {
    tag: config.require("tag"),
    sources: [image1.ref, image2.ref],
    push: true,
  },
  { provider: provider }
);

new docker.buildx.Index(
  "index-not-pushed",
  {
    tag: config.require("tag") + "-not-pushed",
    sources: [image1.ref, image2.ref],
    push: false,
  },
  { provider: provider }
);
