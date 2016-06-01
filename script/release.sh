#! /bin/bash

set -e

WORK=/tmp/gull-release
mkdir -p $WORK

SORUCE=github.com/c2fo/gull/source/bin/gull

GOOS=darwin
GOARCH=amd64

go build -v $SOURCE -o $WORK/mac/gull

GOOS=windows

go build -v $SOURCE -o $WORK/win/gull.exe

GOOS=linux

go build -v $SOURCE -o $WORK/lin/gull