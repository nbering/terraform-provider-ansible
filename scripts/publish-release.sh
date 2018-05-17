#!/bin/bash
set -e -x

if 
    ! which github-release
then
    go get github.com/aktau/github-release
fi

VERSION="$(git tag --points-at $(git rev-parse HEAD) | grep "v\(.*\)" 2> /dev/null)"

if 
    [ "${VERSION}" == "" ]
then
    echo "current commit is not tagged, please make sure to tag before releasing"
    exit 1
fi

GITHUB_USER="${1:-nbering}"
GITHUB_REPO="${2:-terraform-provider-ansible}"
RELEASE_ARGS="--user ${GITHUB_USER} --repo ${GITHUB_REPO} --tag ${VERSION}"

github-release release ${RELEASE_ARGS} \
		--name ${VERSION} \
		--description ${VERSION}

cd pkg
for file in \
    $(find . -maxdepth 1 -type f | cut -b 3-)
do
    github-release upload ${RELEASE_ARGS} \
		--name "${file}" \
		--file "${file}"
done
cd ..