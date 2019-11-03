#!/bin/sh

read -p "Enter your secret : " secret
docker build --rm -t="laatootester:latest" -f Dockerfile.laatootest --build-arg seedpass=$secret .
