#!/bin/bash

docker build . -t demo_api_build
docker create --name demo_api_build demo_api_build
docker cp demo_api_build:/go/src/github.com/bturbes/demo_api/demo_api demo_api
docker rm demo_api_build
docker build --file Dockerfile.prod --tag bturbes/demo_api .
