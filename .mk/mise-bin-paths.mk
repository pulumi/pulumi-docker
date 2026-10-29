# Example conversion runs `pulumi convert` from a temporary directory outside
# the repository. When mise shims come first on PATH (as they do in CI after
# jdx/mise-action), Pulumi resolves ambient plugins such as
# pulumi-converter-terraform to the shim, and the shim cannot resolve a version
# outside the project, so conversion fails. tfgen swallows that error and
# silently drops the converted examples (e.g. from docs/_index.md).
#
# Put the real tool directories for this project ahead of the shims for the
# targets that convert examples.
MISE_BIN_PATHS := $(shell mise bin-paths 2>/dev/null | paste -sd: -)
ifneq ($(MISE_BIN_PATHS),)
.make/schema .make/build_registry_docs: export PATH := $(MISE_BIN_PATHS):$(PATH)
endif
