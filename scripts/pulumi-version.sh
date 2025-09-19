#!/usr/bin/env bash

set -euo pipefail

module_path="github.com/pulumi/pulumi/pkg/v3"
gomod="provider/go.mod"

if [[ ! -f "$gomod" ]]; then
    echo "missing $gomod" >&2
    exit 1
fi

raw_version=$(awk -v module="$module_path" '
    $1 == module || $2 == module {
        for (i = 1; i <= NF; i++) {
            if ($i ~ /^v[0-9]/) {
                sub(/^v/, "", $i)
                print $i
                exit
            }
        }
    }
' "$gomod")

if [[ -z "${raw_version:-}" ]]; then
    echo "failed to determine Pulumi version from $gomod" >&2
    exit 1
fi

echo "PULUMI_VERSION=$raw_version"
# export PULUMI_VERSION=$raw_version
