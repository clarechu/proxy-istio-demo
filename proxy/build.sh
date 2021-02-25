#!/usr/bin/env bash

echo  "GOOS=linux go build"
 GOOS=linux go build -o proxy

scp proxy root@10.10.13.110:~/demo

rm -rf proxy