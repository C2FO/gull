#! /bin/bash

# Reference: https://github.com/coreos/etcd/tree/master/etcdctl

export ETCDCTL_PEERS=${ETCDCRL_PEERS:-"http://localhost:4002"}
export ETCDCTL_ENDPOINT=${ETCDCTL_ENDPOINT:-"http://localhost:4002"}

echo "Reading all values from etcd"

etcdctl --no-sync ls --recursive -p | while read -r line ; do
    if [[ "${line:${#line}-1}" != "/" ]]; then
    	stored=$(etcdctl --no-sync get $line)
    	echo "$line -> $stored"
    fi
done