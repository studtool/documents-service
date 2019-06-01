#!/usr/bin/env bash

command="$1"
version="v1.16.0"

if [[ "${command}" = "install" ]]; then
    curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh \
        | sh -s -- -b $(go env GOPATH)/bin "${version}"
elif [[ "${command}" = "run" ]]; then
    golangci-lint run
fi
