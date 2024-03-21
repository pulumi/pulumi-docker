## Migrating v3 and v4 Image resources

The `buildx.Image` resource provides a superset of functionality over the `Image` resources available in versions 3 and 4 of the Pulumi Docker provider.
Existing `Image` resources can be converted to `build.Image` resources with minor modifications.

### Behavioral differences

There are several key behavioral differences to keep in mind when transitioning images to the new `buildx.Image` resource.

#### Previews

Version `3.x` of the Pulumi Docker provider always builds images during preview operations.
This is helpful as a safeguard to prevent "broken" images from merging, but users found the behavior unnecessarily redundant when running previews and updates locally.

Version `4.x` changed build-on-preview behavior to be opt-in.
By default, `v4.x` `Image` resources do _not_ build during previews, but this behavior can be toggled with the `buildOnPreview` option.
Some users felt this made previews in CI less helpful because they no longer detected bad images by default.

The default behavior of the `buildx.Image` resource has been changed to strike a better balance between CI use cases and manual updates.
By default, Pulumi will now only build `buildx.Image` resources during previews when it detects a CI environment like GitHub Actions.
Previews run in non-CI environments will not build images.
This behavior is still configurable with `buildOnPreview`.

#### Push behavior

Versions `3.x` and `4.x` of the Pulumi Docker provider attempt to push images to remote registries by default.
They expose a `skipPush: true` option to disable pushing.

The `buildx.Image` resource matches the Docker CLI's behavior and does not push images anywhere by default.

To push images to a registry you can include `push: true` (equivalent to Docker's `--push` flag) or configure an `export` of type `registry` (equivalent to Docker's `--output type=registry`).
Like Docker, if an image is configured without exports you will see a warning with instructions for how to enable pushing, but the build will still proceed normally.

#### Secrets

Version `3.x` of the Pulumi Docker provider supports secrets by way of the `extraOptions` field.

Version `4.x` of the Pulumi Docker provider does not support secrets.

The `buildx.Image` resource supports secrets but does not require those secrets to exist on-disk or in environment variables.
Instead, they should be passed directly as values.
(Please be sure to familiarize yourself with Pulumi's [native secret handling](https://www.pulumi.com/docs/concepts/secrets/).)
Pulumi also provides [ESC](https://www.pulumi.com/product/esc/) to make it easier to share secrets across stacks and environments.

#### Caching

Version `3.x` of the Pulumi Docker provider exposes `cacheFrom: bool | { stages: [...] }`.
It builds targets individually and pushes them to separate images for caching.

Version `4.x` exposes a similar parameter `cacheFrom: { images: [...] }` which pushes and pulls inline caches.

Both versions 3 and 4 require specific environment variables to be set and deviate from Docker's native caching behavior.
This can result in inefficient builds due to unnecessary image pulls, repeated file transfers, etc.

The `buildx.Image` resource delegates all caching behavior to Docker.
`cacheFrom` and `cacheTo` options (equivalent to Docker's `--cache-to` and `--cache-from`) are exposed and provide additional cache targets, such as local disk, S3 storage, etc.

#### Outputs

Versions `3.x` and `4.x` of the provider exposed a `repoDigest` output which was a fully qualified tag with digest.
In `4.x` this could also be a single sha256 hash if the image wasn't pushed.

Unlike earlier providers the `buildx.Image` resource can push multiple tags.
As a convenience, it exposes a `ref` output consisting of a tag with digest as long as the image was pushed.
If multiple tags were pushed this uses one at random.

If you need more control over tag references you can use the `digest` output, which is always a single sha256 hash as long as the image was exported somewhere.

#### Tag deletion and refreshes

Versions 3 and 4 of Pulumi Docker provider do not delete tags when the `Image` resource is deleted, nor do they confirm expected tags exist during `refresh` operations.

The `buidx.Image` will query your registries during `refresh` to ensure the expected tags exist.
If any are missing a subsequent `update` will push them.

When a `buildx.Image` is deleted, it will _attempt_ to also delete any pushed tags.
Deletion of remote tags is not guaranteed because not all registries support the manifest `DELETE` API (`docker.io` in particular).
Manifests are _not_ deleted in the same way during updates -- to do so safely would require a full build to determine whether a Pulumi operation should be an update or update-replace.

Use the [`retainOnDelete: true`](https://www.pulumi.com/docs/concepts/options/retainondelete/) option if you do not want tags deleted.

### Example migration

Examples of "fully-featured" `v3` and `v4` `Image` resources are shown below, along with an example `buildx.Image` resource showing how they would look after migration.

The `v3` resource leverages `buildx` via a `DOCKER_BUILDKIT` environment variable and CLI flags passed in with `extraOption`.
After migration, the environment variable is no longer needed and CLI flags are now properties on the `buildx.Image`.
In almost all cases, properties of `buildx.Image` are named after the Docker CLI flag they correspond to.

The `v4` resource is less functional than its `v3` counterpart because it lacks the flexibility of `extraOptions`.
It it is shown with parameters similar to the `v3` example for completeness.

{{% examples %}}
## Example Usage
{{% example %}}
### v3/v4 migration

```typescript

// v3 Image
const v3 = new docker.Image("v3-image", {
  imageName: "myregistry.com/user/repo:latest",
  localImageName: "local-tag",
  skipPush: false,
  build: {
    dockerfile: "./Dockerfile",
    context: "../app",
    target: "mytarget",
    args: {
      MY_BUILD_ARG: "foo",
    },
    env: {
      DOCKER_BUILDKIT: "1",
    },
    extraOptions: [
      "--cache-from",
      "type=registry,myregistry.com/user/repo:cache",
      "--cache-to",
      "type=registry,myregistry.com/user/repo:cache",
      "--add-host",
      "metadata.google.internal:169.254.169.254",
      "--secret",
      "id=mysecret,src=/local/secret",
      "--ssh",
      "default=/home/runner/.ssh/id_ed25519",
      "--network",
      "host",
      "--platform",
      "linux/amd64",
    ],
  },
  registry: {
    server: "myregistry.com",
    username: "username",
    password: pulumi.secret("password"),
  },
});

// v3 Image after migrating to buildx.Image
const v3Migrated = new docker.buildx.Image("v3-to-buildx", {
    tags: ["myregistry.com/user/repo:latest", "local-tag"],
    push: true,
    dockerfile: {
        location: "./Dockerfile",
    },
    context: {
        location: "../app",
    },
    targets: ["mytarget"],
    buildArgs: {
        MY_BUILD_ARG: "foo",
    },
    cacheFrom: [{ registry: { ref: "myregistry.com/user/repo:cache" } }],
    cacheTo: [{ registry: { ref: "myregistry.com/user/repo:cache" } }],
    secrets: {
        mysecret: "value",
    },
    addHosts: ["metadata.google.internal:169.254.169.254"],
    ssh: {
        default: ["/home/runner/.ssh/id_ed25519"],
    },
    network: "host",
    platforms: ["linux/amd64"],
    registries: [{
        address: "myregistry.com",
        username: "username",
        password: pulumi.secret("password"),
    }],
});


// v4 Image
const v4 = new docker.Image("v4-image", {
    imageName: "myregistry.com/user/repo:latest",
    skipPush: false,
    build: {
        dockerfile: "./Dockerfile",
        context: "../app",
        target: "mytarget",
        args: {
            MY_BUILD_ARG: "foo",
        },
        cacheFrom: {
            images: ["myregistry.com/user/repo:cache"],
        },
        addHosts: ["metadata.google.internal:169.254.169.254"],
        network: "host",
        platform: "linux/amd64",
    },
    buildOnPreview: true,
    registry: {
        server: "myregistry.com",
        username: "username",
        password: pulumi.secret("password"),
    },
});

// v4 Image after migrating to buildx.Image
const v4Migrated = new docker.buildx.Image("v4-to-buildx", {
    tags: ["myregistry.com/user/repo:latest"],
    push: true,
    dockerfile: {
        location: "./Dockerfile",
    },
    context: {
        location: "../app",
    },
    targets: ["mytarget"],
    buildArgs: {
        MY_BUILD_ARG: "foo",
    },
    cacheFrom: [{ registry: { ref: "myregistry.com/user/repo:cache" } }],
    cacheTo: [{ registry: { ref: "myregistry.com/user/repo:cache" } }],
    addHosts: ["metadata.google.internal:169.254.169.254"],
    network: "host",
    platforms: ["linux/amd64"],
    registries: [{
        address: "myregistry.com",
        username: "username",
        password: pulumi.secret("password"),
    }],
});

```

{{% /example %}}
