#!/bin/sh
set -x

basedir=`pwd`/gopath/src/github.com/burdzwastaken/concourse-spinnaker-resource
build_dir=`pwd`/build-output

mkdir -p ${build_dir} > /dev/null

set -e

export GOPATH=`pwd`/gopath:${basedir}/cmd/out/vendor
export CGO_ENABLED=0

origbase=`pwd`
cd ${basedir}
go build -o ${build_dir}/assets/check ./cmd/check
go build -o ${build_dir}/assets/in ./cmd/in
go build -o ${build_dir}/assets/out ./cmd/out
cd ${origbase}

cp -a ${basedir}/Dockerfile ${build_dir}
