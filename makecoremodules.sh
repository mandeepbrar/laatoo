#!/bin/sh

nodeModulesFolder=/home/mandeep/goprogs/src/laatoo/designer/ui/nodemodules
pluginsRoot=/home/mandeep/goprogs
deploy=/home/mandeep/goprogs/src/laatoo/modules

docker run --rm -it -v $nodeModulesFolder:/nodemodules -v $pluginsRoot:/plugins -v $deploy:/deploy -e name=shell  -e packageFolder=laatoo/services/shell  laatoomodulebuilder

docker run --rm -it -v $nodeModulesFolder:/nodemodules -v $pluginsRoot:/plugins -v $deploy:/deploy -e name=redispubsub  -e packageFolder=laatoo/services/pubsub/redis  laatoomodulebuilder

docker run --rm -it -v $nodeModulesFolder:/nodemodules -v $pluginsRoot:/plugins -v $deploy:/deploy -e name=role  -e packageFolder=laatoo/services/security/objects/role  laatoomodulebuilder
docker run --rm -it -v $nodeModulesFolder:/nodemodules -v $pluginsRoot:/plugins -v $deploy:/deploy -e name=user  -e packageFolder=laatoo/services/security/objects/user  laatoomodulebuilder
docker run --rm -it -v $nodeModulesFolder:/nodemodules -v $pluginsRoot:/plugins -v $deploy:/deploy -e name=dblogin  -e packageFolder=laatoo/services/security/db  laatoomodulebuilder

docker run --rm -it -v $nodeModulesFolder:/nodemodules -v $pluginsRoot:/plugins -v $deploy:/deploy -e name=memorycache  -e packageFolder==laatoo/services/cache/memory  laatoomodulebuilder
docker run --rm -it -v $nodeModulesFolder:/nodemodules -v $pluginsRoot:/plugins -v $deploy:/deploy -e name=appenginecache  -e packageFolder=laatoo/services/cache/appengine  laatoomodulebuilder
docker run --rm -it -v $nodeModulesFolder:/nodemodules -v $pluginsRoot:/plugins -v $deploy:/deploy -e name=rediscache  -e packageFolder=laatoo/services/cache/redis  laatoomodulebuilder

docker run --rm -it -v $nodeModulesFolder:/nodemodules -v $pluginsRoot:/plugins -v $deploy:/deploy -e name=blevesearch  -e packageFolder=laatoo/services/search/bleve  laatoomodulebuilder
docker run --rm -it -v $nodeModulesFolder:/nodemodules -v $pluginsRoot:/plugins -v $deploy:/deploy -e name=googlesearch  -e packageFolder=laatoo/services/search/google  laatoomodulebuilder

docker run --rm -it -v $nodeModulesFolder:/nodemodules -v $pluginsRoot:/plugins -v $deploy:/deploy -e name=beanstalktasks  -e packageFolder=laatoo/services/tasks/beanstalk  laatoomodulebuilder
docker run --rm -it -v $nodeModulesFolder:/nodemodules -v $pluginsRoot:/plugins -v $deploy:/deploy -e name=gaetasks  -e packageFolder=laatoo/services/tasks/gae  laatoomodulebuilder

docker run --rm -it -v $nodeModulesFolder:/nodemodules -v $pluginsRoot:/plugins -v $deploy:/deploy -e name=filesystemstorage  -e packageFolder=laatoo/services/storage/filesystem  laatoomodulebuilder
docker run --rm -it -v $nodeModulesFolder:/nodemodules -v $pluginsRoot:/plugins -v $deploy:/deploy -e name=googlestorage  -e packageFolder=laatoo/services/storage/googlestorage  laatoomodulebuilder

docker run --rm -it -v $nodeModulesFolder:/nodemodules -v $pluginsRoot:/plugins -v $deploy:/deploy -e name=staticfileserver  -e packageFolder=laatoo/services/static  laatoomodulebuilder

docker run --rm -it -v $nodeModulesFolder:/nodemodules -v $pluginsRoot:/plugins -v $deploy:/deploy -e name=publicdirserver  -e packageFolder=laatoo/services/publicdir  laatoomodulebuilder
docker run --rm -it -v $nodeModulesFolder:/nodemodules -v $pluginsRoot:/plugins -v $deploy:/deploy -e name=publicfilesserver  -e packageFolder=laatoo/services/publicfiles  laatoomodulebuilder

docker run --rm -it -v $nodeModulesFolder:/nodemodules -v $pluginsRoot:/plugins -v $deploy:/deploy -e name=localauth  -e packageFolder=laatoo/services/localauth  laatoomodulebuilder

docker run --rm -it -v $nodeModulesFolder:/nodemodules -v $pluginsRoot:/plugins -v $deploy:/deploy -e name=dataadapter  -e packageFolder=laatoo/services/data/dataadapter  laatoomodulebuilder

docker run --rm -it -v $nodeModulesFolder:/nodemodules -v $pluginsRoot:/plugins -v $deploy:/deploy -e name=mongodatabase  -e packageFolder=laatoo/services/data/mongo  laatoomodulebuilder
docker run --rm -it -v $nodeModulesFolder:/nodemodules -v $pluginsRoot:/plugins -v $deploy:/deploy -e name=gaedatastore  -e packageFolder=laatoo/services/data/gae  laatoomodulebuilder
docker run --rm -it -v $nodeModulesFolder:/nodemodules -v $pluginsRoot:/plugins -v $deploy:/deploy -e name=sqldatabase  -e packageFolder=laatoo/services/data/sql  laatoomodulebuilder

docker run --rm -it -v $nodeModulesFolder:/nodemodules -v $pluginsRoot:/plugins -v $deploy:/deploy -e name=ui  -e packageFolder=laatoo/services/ui  laatoomodulebuilder
