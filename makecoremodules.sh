#!/bin/sh

docker run -it --rm -v /home/mandeep/goprogs:/plugin -e package=laatoo/services/cache/memory -e name=memorycache laatoomodulebuilder:latest
docker run -it --rm -v /home/mandeep/goprogs:/plugin -e package=laatoo/services/cache/appengine -e name=appenginecache laatoomodulebuilder:latest
docker run -it --rm -v /home/mandeep/goprogs:/plugin -e package=laatoo/services/cache/redis -e name=rediscache laatoomodulebuilder:latest


docker run -it --rm -v /home/mandeep/goprogs:/plugin -e package=laatoo/services/pubsub/redis -e name=redispubsub laatoomodulebuilder:latest

docker run -it --rm -v /home/mandeep/goprogs:/plugin -e package=laatoo/services/search/bleve -e name=blevesearch laatoomodulebuilder:latest
docker run -it --rm -v /home/mandeep/goprogs:/plugin -e package=laatoo/services/search/google -e name=googlesearch laatoomodulebuilder:latest
