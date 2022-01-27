# go-httpecho

[![Docker Image Version (latest semver)](https://img.shields.io/docker/v/isayme/httpecho?sort=semver&style=flat-square)](https://hub.docker.com/r/isayme/httpecho)
![Docker Image Size (latest semver)](https://img.shields.io/docker/image-size/isayme/httpecho?sort=semver&style=flat-square)
![Docker Pulls](https://img.shields.io/docker/pulls/isayme/httpecho?style=flat-square)

http echo server with Golang

# Run with Dokcer

## Docker Cli
```
docker run -p 3000:3000 isayme/httpecho:lastest
```

## Docker Compose
```
version: '3'

services:
  httpecho:
    container_name: httpecho
    image: isayme/httpecho:latest
    ports:
      # http echo serve with 3000 port
      - '3000:3000'
    restart: unless-stopped
```
