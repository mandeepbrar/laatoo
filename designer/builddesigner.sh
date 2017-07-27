#!/bin/sh
rm -fr dist/lib/designer
mkdir -p dist/lib/designer
docker run -it --rm -v /home/mandeep/goprogs:/plugin -e package=laatoo/designer/core -e name=designer laatoobuilder:latest
cp /home/mandeep/goprogs/bin/designer.so dist/lib/designer/
./makedockerimage.sh
