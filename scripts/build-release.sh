#!/bin/bash
set -e -x
# Cross-compile for all platforms with gox

rm -r pkg/*

PROJECT_NAME=${PWD##*/}
RELEASE_VERSION="$(git describe --abbrev=0)"

gox -output "pkg/{{.OS}}_{{.Arch}}/{{.Dir}}_${RELEASE_VERSION}"

cd pkg;
find * -type d -exec zip -r "${PROJECT_NAME}-{}.zip" "{}" \;
