#!/bin/bash

set -eu

go run main.go version.go commands.go generate --config ./res/config.toml

#go run main.go version.go commands.go generate --config ./res/config.toml |
#	go run main.go version.go commands.go scale --config ./res/config.toml |
#	go run main.go version.go commands.go trim --config ./res/config.toml |
#  sort |
#	go run main.go version.go commands.go paste --config ./res/config.toml
