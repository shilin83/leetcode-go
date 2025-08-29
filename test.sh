#!/bin/bash

set -e

file="coverage.out"
cover="coverage"

if [ -n "$1" ]; then
    if [ "$1" == "$cover" ]; then
        go test -coverprofile="$file" -coverpkg=./src/... ./...
        go tool cover -html="$file"
    else
        printf "Invalid argument: %s\nUsage: /bin/bash test.sh %s\n" "$1" "$cover"
    fi
else
    go test ./...
fi
