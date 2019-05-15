#!/bin/sh
docker run -it --rm -v /home/mandeep/goprogs/src/laatoo/server:/laatooserver -v /home/mandeep/goprogs/src/laatoo/sdk:/laatoosdk laatoocompiler:latest
rm -fr dist/bin/
mkdir -p dist/bin
cp --parents server/laatoo dist/bin/
