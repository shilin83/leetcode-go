#!/usr/bin/env bash

set -e

go test -coverprofile=coverage.out -coverpkg=./src/... ./tests/...

go tool cover -html=coverage.out
