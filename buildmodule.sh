#!/bin/sh
rm -fr dist/lib/punjabimehfil
mkdir -p dist/lib/punjabimehfil
docker run -it --rm -v /home/mandeep/goprogs:/plugin -e package=punjabimehfil/core -e name=punjabimehfil laatoobuilder:latest
cp /home/mandeep/goprogs/bin/punjabimehfil.so dist/lib/punjabimehfil/
./makedockerimage.sh
