#!/bin/sh

tag=5.7.21

docker stop sintor-db
docker rm -v sintor-db
docker run -i -t -d -p 3308:3306 -e MYSQL_ROOT_PASSWORD=xxxx --name sintor-db sintor-db:${tag}