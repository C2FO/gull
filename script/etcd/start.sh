#! /bin/bash

docker run \
  -d \
  -p 2379:2379 \
  elcolio/etcd \
  -advertise-client-urls http://localhost:2379