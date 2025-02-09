#!/usr/bin/env bash

set -e

go test -v -coverprofile=coverage.out ./...

go tool cover -html=coverage.out
