## HEAD (Unreleased)
_(none)_

---

## 2.9.0 (2021-03-16)
* Upgrade to pulumi-terraform-bridge v2.21.0
* Release macOS arm64 binary

## 2.8.1 (2021-02-16)
* Upgrade to pulumi-terraform-bridge v2.19.0  
  **Please Note:** This includes a bug fix that stops mutating resources options in the nodejs provider

## 2.8.0 (2021-02-01)
* Upgrade to pulumi-terraform-bridge v2.18.1

## 2.7.0 (2021-01-27)
* Upgrade to v2.11.0 of the Docker Terraform Provider

## 2.6.1 (2021-01-13)
* Upgrade to pulumi-terraform-bridge v2.17.0
* Upgrade to Pulumi v2.17.0

## 2.6.0 (2021-01-08)
* Upgrade to v2.10.0 of the Docker Terraform Provider
* Improve build and push output [#251](https://github.com/pulumi/pulumi-docker/pull/251)

## 2.5.0 (2020-12-07)
* Upgrade to v2.8.0 of the Docker Terraform Provider

## 2.4.1 (2020-11-24)
* Upgrade to pulumi-terraform-bridge v2.13.2  
  * This adds support for import specific examples in documentation

## 2.4.0 (2020-10-26)
* Upgrade to Pulumi v2.12.0 and pulumi-terraform-bridge v2.11.0
* Improving the accuracy of previews leading to a more accurate understanding of what will actually change rather than assuming all output properties will change.  
  ** PLEASE NOTE:**  
  This new preview functionality can be disabled by setting `PULUMI_DISABLE_PROVIDER_PREVIEW` to `1` or `false`.

## 2.3.1 (2020-10-07)

* Always append specified environment variables to the current OS environment variable set [#212](https://github.com/pulumi/pulumi-docker/pull/212)
* Fix Python type hints for lists [#225](https://github.com/pulumi/pulumi-docker/pull/225)
* Upgrade to pulumi-terraform-bridge v2.8.0
* Upgrade to Pulumi v2.10.0
* Fix .NET concurrency issue leading to "No digest available for image" while building an image
  [#229](https://github.com/pulumi/pulumi-docker/pull/229)

## 2.3.0 (2020-08-31)
* Fix bug in python Image implementation [#190](https://github.com/pulumi/pulumi-docker/issues/190)
* Upgrade to v2.7.2 of the Docker Terraform Provider
* Upgrade to pulumi-terraform-bridge v2.7.3
* Upgrade to Pulumi v2.9.0, which adds type annotations and input/output classes to Python
* Avoid storing transient config from environment into the statefile

## 2.2.3 (2020-06-17)
* Switch to GitHub actions for build

## 2.2.2 (2020-06-05)
* Upgrade to v2.7.1 of the Docker Terraform Provider

## 2.2.1 (2020-05-28)
* Upgrade to Pulumi v2.3.0
* Upgrade to pulumi-terraform-bridge v2.4.0

## 2.2.0 (2020-05-21)
* Fix Python Docker Image build support for registries
* Fix Go Docker Image implementation [#175](https://github.com/pulumi/pulumi-docker/pull/175)

## 2.1.1 (2020-05-12)
* Upgrade to pulumi-terraform-bridge v2.3.1

## 2.1.0 (2020-04-28)
* Regenerate datasource examples to be async
* Upgrade to pulumi-terraform-bridge v2.1.0

## 2.0.0 (2020-04-18)
* Upgrade to Pulumi v2.0.0
* Upgrade to pulumi-terraform-bridge v2.0.0
* Refactor layout to support Go modules
* Upgrade to Pulumi v1.13.1
* Upgrade to pulumi-terraform-bridge v1.8.4

## 1.5.0 (2020-03-25)
* Add Docker Image build suport for Python

## 1.4.0 (2020-03-16)
* Add support for Go (https://github.com/pulumi/pulumi-docker/pull/147)
* Upgrade to Pulumi v1.12.1
* Upgrade to pulumi-terraform-bridge v1.8.2

## 1.3.0 (2020-02-27)
* Upgrade to v2.7.0 of the Docker Terraform Provider
* Rename `docker.Config` to `docker.ServiceConfig` to avoid collisions with `Config` package.

## 1.2.0 (2020-01-29)
* Upgrade to pulumi-terraform-bridge v1.6.4

## 1.1.0 (2019-12-18)
* Add Terraform resources to .NET SDK ([#121](https://github.com/pulumi/pulumi-docker/pull/121)).

## 1.0.0 (2019-12-06)
* Regenerate SDK against tf2pulumi 0.6.0
* Add ability to skip push on image build
* Upgrade to v2.6.0 of the Docker Terraform Provider
* Upgrade to go1.13
* Add support for .NET (https://github.com/pulumi/pulumi-docker/pull/115)
* Allow users to pass `target` in image.Image when using multi-target Dockerfiles

## 0.17.4 (Released September 5, 2019)

- Update to commit 8a5b696b491c of the Docker Terraform Provider.
- Upgrade pulumi-terraform to 3f206601e7
- Upgrade to Pulumi v1.0.0

## 0.17.3 (Released August 20, 2019)

- Fix image tag name collision during build time (https://github.com/pulumi/pulumi-docker/pull/90)
- Update dependency to latest version of `pulumi`

## 0.17.2 (Released July 19, 2019)

- Add ability to specify arbitrary extra `docker build` CLI options for `buildAndPush...()` functions.

## 0.17.1 (Released March 7, 2019)

## Improvements

- Fix an issue where the Python `pulumi_docker` package was depending on an older `pulumi` package.

## 0.17.0 (Released March 5, 2019)

### Important

Updating to v0.17.0 version of `@pulumi/pulumi`.  This is an update that will not play nicely
in side-by-side applications that pull in prior versions of this package.

See https://github.com/pulumi/pulumi/commit/7f5e089f043a70c02f7e03600d6404ff0e27cc9d for more details.

As such, we are rev'ing the minor version of the package from 0.16 to 0.17.  Recent version of `pulumi` will now detect, and warn, if different versions of `@pulumi/pulumi` are loaded into the same application.  If you encounter this warning, it is recommended you move to versions of the `@pulumi/...` packages that are compatible.  i.e. keep everything on 0.16.x until you are ready to move everything to 0.17.x.

## 0.16.4 (Released January 25th, 2019)

- docker.Image and docker.buildAndPushImage allow a wider set of inputs (i.e. promises and outputs), making it easier to pass in values produced by other resources.

### Improvements

## 0.16.3 (Released January 15th, 2019)

### Improvements

- Updated package constraints such that we do not depend on unreleased versions of `@pulumi/pulumi`.

## 0.16.2 (Released December 5th, 2018)

### Improvements

- Expose resources from Terraform's `docker` provider.

## 0.16.1 (Released Novemeber 13th, 2018)

### Improvements

- Fix an issue where image caching would not work as expected for multi-stage builds.

- Use a unique name per image when tagging it in the registry.

- Fix an issue which could cause iamges to be pushed when there were no relevent updates.

- Add a `registryServer` property to an `Image` to provide information about what registry the image belongs to.

- Don't run `docker login` for the same registry multiple times.

- Don't show output from `docker` invocations unless it fails or issues warnings.
