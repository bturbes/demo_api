#!/bin/bash -e

TAG=$1
IMAGE_NAME="bturbes/demo_api:${TAG:-latest}"

mkdir -p bin
docker run --rm -v "$PWD":/go/src/github.com/bturbes/demo_api -w /go/src/github.com/bturbes/demo_api golang:1.9 go build -o bin/demo_api.linux

docker build -t "$IMAGE_NAME" .
