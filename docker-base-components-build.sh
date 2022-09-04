#! /bin/bash

export COMPOSE_HTTP_TIMEOUT=500
export DOCKER_CLIENT_TIMEOUT=500
export COMPOSE_PARALLEL_LIMIT=1024
echo y | docker network prune
docker network create mall
docker-compose -f docker-compose.yml up -d
docker ps