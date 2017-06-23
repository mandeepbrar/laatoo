#!/bin/sh
docker build -t="laatoobase:latest" -f Dockerfile.laatoobase .
docker build -t="laatoobuilder:latest" -f Dockerfile.laatoobuilder .
docker build -t="laatooservicesbuilder:latest" -f Dockerfile.laatooservicesbuilder .
