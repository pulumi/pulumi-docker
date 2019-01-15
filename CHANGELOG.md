## 0.16.3 (unreleased)

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
