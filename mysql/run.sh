#!/bin/bash

docker run -d -p 3307:3306 -e MYSQL_ROOT_PASSWORD=supersecret -e MYSQL_DATABASE=RussianDoll -e MYSQL_USER=AdminUser -e MYSQL_PASSWORD=AdminPassword russian-doll-db