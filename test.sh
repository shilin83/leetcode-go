#!/usr/bin/env bash

set -e

go test -v -coverprofile=coverage.out -coverpkg=./src/... ./...

go tool cover -html=coverage.out
