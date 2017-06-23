#!/bin/sh
docker run -it --rm -v /home/mandeep/goprogs/src/laatoo:/go/src/laatoo laatoobuilder:latest
docker run -it --rm -v /home/mandeep/goprogs/src/laatoo:/go/src/laatoo laatooservicesbuilder:latest
