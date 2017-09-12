#docker run -it --rm -v /home/mandeep/goprogs/src/laatoo/designer/dist/config/designer:/config/environments/designer -v /home/mandeep/goprogs/src/laatoo/dist/bin:/bin  -v /home/mandeep/goprogs/src/laatoo/designer/dist/publicdir:/bin/publicdir  -v /home/mandeep/goprogs/src/laatoo/designer/dist/appfiles:/bin/appfiles --net="host" designertester:latest
docker-compose -f Dockercomposedev.yml up
