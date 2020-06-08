#!/bin/bash

docker run --volume=$HOME/mysql/data:/data -d -p 3307:3306 -e MYSQL_ROOT_PASSWORD=supersecret -e MYSQL_DATABASE=Graph -e MYSQL_USER=AdminUser -e MYSQL_PASSWORD=AdminPassword graph-db