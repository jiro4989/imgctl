#!/bin/bash

git tag $1
export GITHUB_TOKEN=`cat ./build/token`
goreleaser --rm-dist
go install
