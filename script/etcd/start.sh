#! /bin/bash

docker run \
  -d \
  -p 4002:4001 \
  elcolio/etcd:latest \
  -advertise-client-urls http://localhost:4002