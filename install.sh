#!/bin/bash
if [[ -z "${GOPATH}" ]]; then
	echo "Please set GOPATH"
	exit 1
fi
if [[ -z "${GOBIN}" ]]; then
	echo "warning: GOBIN not set. Writing to default"
fi
go get -u github.com/golang/protobuf/protoc-gen-go

cd $GOPATH/src/github.com/rolandktud/calculator/;dep ensure

protoc -I calculatorService/  calculatorService/calculator.proto --go_out=plugins=grpc:calculatorService

go install github.com/rolandktud/calculator/client
go install github.com/rolandktud/calculator/server
