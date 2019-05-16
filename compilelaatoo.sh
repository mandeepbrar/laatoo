#!/bin/sh
docker run -it --rm -v /home/mandeep/goprogs/src/laatoo/server:/laatooserver laatoocompiler:latest
rm -fr dist/bin/
mkdir -p dist/bin
cp server/laatoo dist/bin/
