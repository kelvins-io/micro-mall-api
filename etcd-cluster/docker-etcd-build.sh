#! /bin/bash
docker network create etcd-cluster
docker-compose -f docker-compose-etcd.yml up -d

export  ETCDV3_SERVER_URLS=http://127.0.0.1:12379,http://127.0.0.1:22379,http://127.0.0.1:32379
export  ETCDV3_SERVER_URL=http://127.0.0.1:12379,http://127.0.0.1:22379,http://127.0.0.1:32379
export  ETCDCTL_API=3

etcdctl --endpoints=127.0.0.1:12379,127.0.0.1:22379,127.0.0.1:32379 member list