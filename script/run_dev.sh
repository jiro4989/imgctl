#!/bin/bash

set -eu

go run ./cmd/main.go ./cmd/version.go ./cmd/commands.go generate |
	go run ./cmd/main.go ./cmd/version.go ./cmd/commands.go scale -s 50 |
	go run ./cmd/main.go ./cmd/version.go ./cmd/commands.go trim -x 100 -y 290 |
  sort |
	go run ./cmd/main.go ./cmd/version.go ./cmd/commands.go paste
