#!/usr/bin/env bash

echo  "GOOS=linux go build"
 GOOS=linux go build -o grpc-server

docker build -t clarechu/grpc-server:v1.0 .

docker push clarechu/grpc-server:v1.0

rm -rf grpc-server