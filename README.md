# docker-web-server
## Overview
docker-we-server provide a web server with a simple unique endpoint

## Key feature
* Define an endpoint and the response for each endpoint
* Able to define from JSON file


## Usage
* Dockerhub is setting now... Sorry.
```
  git clone `https://github.com/oshiro3/docker-web-server.git`
  cd docker-web-serer
  docker build . -t docker-web-server
  docker run -d -p -v ${CONFIG_JSON_FILE}:/config/root.json ${HOST_PORT}:8888
```
