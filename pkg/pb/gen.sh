#!/bin/bash
#
# Generate all gateway protobuf bindings.
# Run from repository root.
#
set -e

# directories containing protos to be built
DIRS="./meta ./tpl"

GOGOPROTO_ROOT="${GOPATH}/src/github.com/gogo/protobuf"
GOGOPROTO_PATH="${GOGOPROTO_ROOT}:${GOGOPROTO_ROOT}/protobuf"

PROJECT_PB_PATH="${GOPATH}/src/uuabc.com/sendmsg/pkg/pb"

for dir in ${DIRS}; do
	pushd ${dir}
		protoc --gofast_out=plugins=grpc,import_prefix=uuabc.com/sendmsg/pkg/pb/:. -I=.:"${GOGOPROTO_PATH}":"${PROJECT_PB_PATH}":"${GOPATH}/src" *.proto
		sed -i.bak -E "s/uuabc.com\/sendmsg\/pkg\/pb\/(gogoproto|github\.com|golang\.org|google\.golang\.org)/\1/g" *.pb.go
		sed -i.bak -E 's/uuabc.com\/sendmsg\/pkg\/pb\/(errors|fmt|io)/\1/g' *.pb.go
		sed -i.bak -E 's/import _ \"gogoproto\"//g' *.pb.go
		sed -i.bak -E 's/import fmt \"fmt\"//g' *.pb.go
		sed -i.bak -E 's/import math \"uuabc.com\/sendmsg\/pkg\/pb\/math\"//g' *.pb.go
		sed -i.bak -E 's/import encoding_binary \"uuabc.com\/sendmsg\/pkg\/pb\/encoding\/binary\"/ import encoding_binary \"encoding\/binary\"/g' *.pb.go
		rm -f *.bak
		goimports -w *.pb.go

		files=`ls *.pb.go`
		for f in ${files}; do
	        protoc-go-inject-tag -input=${f}
		done
	popd
done
