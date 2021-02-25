#!/usr/bin/env bash

echo  "GOOS=linux go build"
 GOOS=linux go build -o demo

docker build -t clarechu/demo:v1.0 .

docker push clarechu/demo:v1.0

rm -rf demo