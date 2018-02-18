#!/bin/bash
set -e -x
# Cross-compile for all platforms with gox

# Check if gox is installed. If not, get it.
which gox
if [[ $? -ne 0 ]]; then
    go get github.com/mitchellh/gox
fi

# Clean old builds if present.
if [ -d pkg ]; then
    rm -r pkg/*
fi

PROJECT_NAME=${PWD##*/}
RELEASE_VERSION="$(git describe --abbrev=0)"

gox -output "pkg/{{.OS}}_{{.Arch}}/{{.Dir}}_${RELEASE_VERSION}"

cd pkg;
find * -type d -exec zip -r "${PROJECT_NAME}-{}.zip" "{}" \;
