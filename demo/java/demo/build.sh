#!/usr/bin/env bash

mvn clean install

docker build -t clarechu/java-demo:v1.0 .

docker push clarechu/java-demo:v1.0

rm -rf target