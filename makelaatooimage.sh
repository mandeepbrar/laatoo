#!/bin/sh
docker build --rm -t="laatoo:latest" -f Dockerfile.laatoo .
docker build --rm -t="laatootester:latest" -f Dockerfile.laatootest .
