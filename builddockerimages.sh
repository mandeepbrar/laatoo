#!/bin/sh
#docker build --rm -t="gonightly:latest" -f Dockerfile.gonightly .
docker build --rm -t="laatoobase:latest" -f Dockerfile.laatoobase .
#docker build --rm -t="laatooservicescompiler:latest" -f Dockerfile.laatooservicescompiler .
#docker build --rm -t="laatoobuilder:latest" -f Dockerfile.pluginbuilder .
docker build --rm -t="laatoouibase:latest" -f Dockerfile.nodebuilder .
#docker build --rm -t="laatoouiwithwasm:latest" -f Dockerfile.rustwasmsupport .
docker build --rm -t="laatoomodulebuilder:latest" -f Dockerfile.modulebuilder .
docker build --rm -t="laatoocompiler:latest" -f Dockerfile.laatoocompiler .
docker build --rm -t="laatootester:latest" -f Dockerfile.laatootest .
