#!/bin/sh
docker build -t="laatoobase:latest" -f Dockerfile.laatoobase .
docker build -t="laatoocompiler:latest" -f Dockerfile.laatoocompiler .
docker build -t="laatooservicescompiler:latest" -f Dockerfile.laatooservicescompiler .

docker build -t="laatoobuilder:latest" -f Dockerfile.pluginbuilder .
