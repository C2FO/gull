#! /bin/bash

set -e
set -x

if [[ $1 ]]; then
    go get -u github.com/alecthomas/gometalinter
    gometalinter --install --update
else
    gometalinter ./source/... --deadline=120s --disable=golint --disable=dupl --disable=goconst --disable gocyclo
fi