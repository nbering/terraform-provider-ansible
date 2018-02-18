#!/usr/bin/env bash

set -e -x

mkdir -p /go/src/github.com/nbering

ln -s /go/src/github.com/nbering/terraform-provider-ansible ./terraform-provider-ansible

pushd /go/src/github.com/nbering/terraform-provider-ansible
    make test
    make testacc
    make vet
popd
