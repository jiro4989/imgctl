#!/bin/bash

set -eu

go run main.go version.go commands.go generate |
	go run main.go version.go commands.go scale -s 50 |
	go run main.go version.go commands.go trim -x 100 -y 290 |
  sort |
	go run main.go version.go commands.go paste
