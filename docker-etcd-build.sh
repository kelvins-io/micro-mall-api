#! /bin/bash
docker network create etcd
docker-compose -f ./etcd-cluster/docker-compose-etcd.yml up -d