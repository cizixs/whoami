#!/bin/bash

set -o errexit

if [[ $# != 1 ]]; then
    echo "usage"
    echo "$0 VERSION"
    exit 1
fi

VERSION=$1

CGO_ENABLED=0 go build -a -ldflags '-s' .
docker build -t whoami:$VERSION .
