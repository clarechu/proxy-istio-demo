#!/usr/bin/env bash

echo  "GOOS=linux go build"
 GOOS=linux go build -o proxy

docker build -t clarechu/proxy:v1.0 .

docker push clarechu/proxy:v1.0

rm -rf proxy