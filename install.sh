#!/bin/bash
set -ex
go get -u github.com/tools/godep
go get -u github.com/jteeuwen/go-bindata/...
go get -u github.com/watercraft/go-raml
cd $GOPATH/src/github.com/watercraft/go-raml
sh build.sh

