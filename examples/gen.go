//go:generate /bin/sh -c "export PATH=\"$(realpath ../bin):$PATH\"; pulumi convert -C buildx/yaml --from yaml --language dotnet --out ../csharp --generate-only"
//go:generate /bin/sh -c "export PATH=\"$(realpath ../bin):$PATH\"; pulumi convert -C buildx/yaml --from yaml --language go     --out ../go     --generate-only"
//go:generate /bin/sh -c "export PATH=\"$(realpath ../bin):$PATH\"; pulumi convert -C buildx/yaml --from yaml --language nodejs --out ../ts     --generate-only"
//go:generate /bin/sh -c "export PATH=\"$(realpath ../bin):$PATH\"; pulumi convert -C buildx/yaml --from yaml --language python --out ../py     --generate-only"
//go:generate /bin/sh -c "export PATH=\"$(realpath ../bin):$PATH\"; pulumi convert -C buildx/yaml --from yaml --language java   --out ../java   --generate-only"
//go:generate rm -rf buildx/*/app
//go:generate cp -r buildx/app buildx/yaml/.
//go:generate cp -r buildx/app buildx/csharp/.
//go:generate cp -r buildx/app buildx/go/.
//go:generate cp -r buildx/app buildx/ts/.
//go:generate cp -r buildx/app buildx/py/.
//go:generate cp -r buildx/app buildx/java/.

package examples
