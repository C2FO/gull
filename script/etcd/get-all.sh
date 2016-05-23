#! /bin/bash

# Reference: https://github.com/coreos/etcd/tree/master/etcdctl

export ETCDCTL_ENDPOINT=${ETCDCTL_ENDPOINT:-"http://localhost:2379"}

echo "Reading all values from etcd"

etcdctl --no-sync ls --recursive -p | while read -r line ; do
    if [[ "${line:${#line}-1}" != "/" ]]; then
    	stored=$(etcdctl --no-sync get $line)
    	echo "$line -> $stored"
    fi
done