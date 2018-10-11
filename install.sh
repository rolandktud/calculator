#!/bin/bash
go get -u github.com/golang/protobuf/protoc-gen-go

cd $GOPATH/src/github.com/rolandktud/calculator/;dep ensure

protoc -I calculatorService/  calculatorService/calculator.proto --go_out=plugins=grpc:calculatorService

go install github.com/rolandktud/calculator/client
go install github.com/rolandktud/calculator/server

