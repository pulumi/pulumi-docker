`Image` builds a Docker image and pushes it Docker and OCI compatible registries.
This resource enables running Docker builds as part of a Pulumi deployment.

Note: This resource does not delete tags, locally or remotely, when destroyed.

## Cross-platform builds

The Image resource supports cross-platform builds when the [Docker engine has cross-platform support enabled via emulators](https://docs.docker.com/build/building/multi-platform/#building-multi-platform-images).
The Image resource currently supports providing only a single operating system and architecture in the `platform` field, e.g.: `linux/amd64`.
To enable this support, you may need to install the emulators in the environment running your Pulumi program.

If you are using Linux, you may be using Docker Engine or Docker Desktop for Linux, depending on how you have installed Docker. The [FAQ for Docker Desktop for Linux](https://docs.docker.com/desktop/faqs/linuxfaqs/#context) describes the differences and how to select which Docker context is in use.

* For local development using Docker Desktop, this is enabled by default.
* For systems using Docker Engine, install the QEMU binaries and register them with using the docker image from [github.com/tonistiigi/binfmt](https://github.com/tonistiigi/binfmt):

  ```shell
  docker run --privileged --rm tonistiigi/binfmt --install all
  ```
* In a GitHub Actions workflow, the [docker/setup-qemu-action](https://github.com/docker/setup-qemu-action) can be used instead by adding this step to your workflow file. Example workflow usage:

  ```yaml
  name: Pulumi
  on:
    push:
      branches:
        - master
  jobs:
    up:
      name: Preview
      runs-on: ubuntu-latest
      steps:
          # This is the step added:
        - name: Set up QEMU
          uses: docker/setup-qemu-action@v2
          # The ordinary pulumi/actions workflow:
        - uses: actions/checkout@v3
        - uses: pulumi/actions@v4
          with:
            command: preview
            stack-name: org-name/stack-name
          env:
            PULUMI_ACCESS_TOKEN: ${{ secrets.PULUMI_ACCESS_TOKEN }}
  ```
