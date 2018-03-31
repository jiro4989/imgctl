#!/bin/bash

git tag $1
export GITHUB_TOKEN=`cat ./res/token`
goreleaser --rm-dist
go install
