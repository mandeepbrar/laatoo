docker ps --filter "status=exited" | grep 'weeks ago' | awk '{print $1}' | xargs --no-run-if-empty docker rm
docker rmi $(docker images --filter "dangling=true" -q --no-trunc) 
