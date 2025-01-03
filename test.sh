#!/usr/bin/env bash

set -e

go test -coverprofile=coverage.out -covermode=atomic -coverpkg=./src/... ./tests/...

if [ -f coverage.out ]; then
    go tool cover --html=coverage.out
fi
