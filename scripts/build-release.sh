#!/bin/bash
set -e -x
# Cross-compile for all platforms with gox

# Check if gox is installed. If not, get it.
set +e
which gox
if [[ $? -ne 0 ]]; then
    go get github.com/mitchellh/gox
fi

which zip
if [[ $? -ne 0 ]]; then
    which dpkg
    if [[ $? -ne 0 ]]; then
        echo "Missing zip utility, and no package install strategy available." >&2
        exit 1
    else
        set -e
        pushd "$(mktemp -d)"
            wget https://deb.debian.org/debian/pool/main/z/zip/zip_3.0-11+b1_amd64.deb
            dpkg -i zip_3.0-11+b1_amd64.deb
            rm zip_3.0-11+b1_amd64.deb
        popd
    fi
fi
set -e

# Clean old builds if present.
if [ -d pkg ]; then
    rm -r pkg/*
fi

PROJECT_NAME=${PWD##*/}
RELEASE_VERSION="$(git describe --abbrev=0)"

gox -output "pkg/{{.OS}}_{{.Arch}}/{{.Dir}}_${RELEASE_VERSION}"

cd pkg;
find * -type d -exec zip -r "${PROJECT_NAME}-{}.zip" "{}" \;
