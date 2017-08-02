#!/bin/sh

docker run -it --rm -v /home/mandeep/goprogs:/plugin -e package=laatoo/services/cache/memory -e name=memorycache laatoomodulebuilder:latest
docker run -it --rm -v /home/mandeep/goprogs:/plugin -e package=laatoo/services/cache/appengine -e name=appenginecache laatoomodulebuilder:latest
docker run -it --rm -v /home/mandeep/goprogs:/plugin -e package=laatoo/services/cache/redis -e name=rediscache laatoomodulebuilder:latest


docker run -it --rm -v /home/mandeep/goprogs:/plugin -e package=laatoo/services/pubsub/redis -e name=redispubsub laatoomodulebuilder:latest

docker run -it --rm -v /home/mandeep/goprogs:/plugin -e package=laatoo/services/search/bleve -e name=blevesearch laatoomodulebuilder:latest
docker run -it --rm -v /home/mandeep/goprogs:/plugin -e package=laatoo/services/search/google -e name=googlesearch laatoomodulebuilder:latest

docker run -it --rm -v /home/mandeep/goprogs:/plugin -e package=laatoo/services/tasks/beanstalk -e name=beanstalktasks laatoomodulebuilder:latest
docker run -it --rm -v /home/mandeep/goprogs:/plugin -e package=laatoo/services/tasks/gae -e name=gaetasks laatoomodulebuilder:latest

docker run -it --rm -v /home/mandeep/goprogs:/plugin -e package=laatoo/services/storage/filesystem -e name=filesystemstorage laatoomodulebuilder:latest
docker run -it --rm -v /home/mandeep/goprogs:/plugin -e package=laatoo/services/storage/googlestorage -e name=googlestorage laatoomodulebuilder:latest

docker run -it --rm -v /home/mandeep/goprogs:/plugin -e package=laatoo/services/static -e name=staticfileserver laatoomodulebuilder:latest
docker run -it --rm -v /home/mandeep/goprogs:/plugin -e package=laatoo/services/publicfiles -e name=publicfileserver laatoomodulebuilder:latest
