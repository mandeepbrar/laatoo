#!/bin/sh

docker run -it --rm -v /home/mandeep/goprogs:/plugin -e package=laatoo/services/security/role -e name=role laatoomodulebuilder:latest
docker run -it --rm -v /home/mandeep/goprogs:/plugin -e package=laatoo/services/security/user -e name=user laatoomodulebuilder:latest
docker run -it --rm -v /home/mandeep/goprogs:/plugin -e package=laatoo/services/security/db -e name=dblogin laatoomodulebuilder:latest


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

docker run -it --rm -v /home/mandeep/goprogs:/plugin -e package=laatoo/services/publicdir -e name=publicdirserver laatoomodulebuilder:latest
docker run -it --rm -v /home/mandeep/goprogs:/plugin -e package=laatoo/services/publicfiles -e name=publicfilesserver laatoomodulebuilder:latest

docker run -it --rm -v /home/mandeep/goprogs:/plugin -e package=laatoo/services/localauth -e name=localauth laatoomodulebuilder:latest

docker run -it --rm -v /home/mandeep/goprogs:/plugin -e package=laatoo/services/data/dataadapter -e name=dataadapter laatoomodulebuilder:latest

docker run -it --rm -v /home/mandeep/goprogs:/plugin -e package=laatoo/services/data/mongo -e name=mongodatabase laatoomodulebuilder:latest
docker run -it --rm -v /home/mandeep/goprogs:/plugin -e package=laatoo/services/data/gae -e name=gaedatastore laatoomodulebuilder:latest
docker run -it --rm -v /home/mandeep/goprogs:/plugin -e package=laatoo/services/data/sql -e name=sqldatabase laatoomodulebuilder:latest
