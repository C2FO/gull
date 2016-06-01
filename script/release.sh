#! /bin/bash


set -x

WORK=./gull-release
rm -rf $WORK
set -e
mkdir -p $WORK

SOURCE="github.com/c2fo/gull/source/bin/gull"

export GOARCH=amd64

export GOOS=darwin
go build -o $WORK/mac/gull -v $SOURCE

export GOOS=windows
go build -o $WORK/win/gull.exe -v $SOURCE

export GOOS=linux
go build -o $WORK/lin/gull -v $SOURCE