## 0.17.5 (Unreleased)
* Regenerate SDK against tf2pulumi 0.6.0
* Add ability to skip push on image build
* Upgrade to v2.6.0 of the Docker Terraform Provider
* Upgrade to go1.13

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
