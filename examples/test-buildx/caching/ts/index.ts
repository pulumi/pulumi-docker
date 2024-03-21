import * as docker from "@pulumi/docker";
import * as pulumi from "@pulumi/pulumi";

const config = new pulumi.Config();

const start = new Date().getTime();

// docker buildx build \
//  -f Dockerfile \
//  --cache-to type=local,dest=tmp-arm64,mode=max,oci-mediatypes=true \
//  --cache-from type=local,src=tmp \
//  --build-arg SLEEP-MS=$SLEEP_MS \
//  -t cached \
//  --platform linux/arm64 \

const image = new docker.buildx.Image(`buildx-${config.require("name")}`, {
  tags: [config.require("name")],
  context: {
    location: ".",
  },
  platforms: ["linux/arm64"],
  buildArgs: {
    SLEEP_SECONDS: config.require("SLEEP_SECONDS"),
  },
  push: false,
  cacheTo: [{ raw: config.require("cacheTo") }],
  cacheFrom: [{ raw: config.require("cacheFrom") }],
  // Set registry auth if it was provided.
  registries: config.getSecret("username").apply((a) =>
    a
      ? [
          {
            address: config.getSecret("address"),
            username: config.getSecret("username"),
            password: config.getSecret("password"),
          },
        ]
      : undefined
  ),
});

export const durationSeconds = image.ref.apply(
  (_) => (new Date().getTime() - start) / 1000.0
);
