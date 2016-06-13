#! /bin/bash

COMPILE_VARS="-X github.com/c2fo/gull/source/lib/common.ApplicationVersion=`cat ./VERSION.txt`"

WORK=./gull-release
rm -rf $WORK
set -e
mkdir -p $WORK

SOURCE="github.com/c2fo/gull/source/bin/gull"

build(){
    OS=$1
    SUFFIX=$2

    export GOOS=$OS
    export GOARCH=amd64

    echo "=== Building $GOARCH/$GOOS"

    OUTPUT="$WORK/gull-$SUFFIX"
    go build -ldflags "$COMPILE_VARS" -o $OUTPUT $SOURCE
}

build darwin mac64
build windows win64.exe
build linux lin64