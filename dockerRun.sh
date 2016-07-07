#!/usr/bin/env bash

docker rm --force hours-ui-mock-api > /dev/null 2>&1
docker run -it -p 3000:3000 -e PORT=3000 -d --name hours-ui-mock-api