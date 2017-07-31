#!/bin/sh

docker run -it --rm -v /home/mandeep/goprogs:/plugin -e package=laatoo/services/cache/memory -e name=memorycache laatoomodulebuilder:latest
