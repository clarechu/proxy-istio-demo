#!/usr/bin/env bash

set -e

echo  "GOOS=linux go build"
 GOOS=linux go build -o grpc-client

docker build -t clarechu/grpc-client:v2.0 .

docker push clarechu/grpc-client:v2.0

rm -rf grpc-client