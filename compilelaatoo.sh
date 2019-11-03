#!/bin/sh
docker run -it --rm -v /home/mandeep/goprogs/src/laatoo/server:/laatooserver -v /home/mandeep/goprogs/src/laatoo/dist/bin:/laatoobin laatoocompiler:latest
#rm -fr dist/bin/
#mkdir -p dist/bin
#cp server/laatooserver dist/bin/
