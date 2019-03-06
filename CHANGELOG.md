## 0.17.1 (Unreleased)

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
