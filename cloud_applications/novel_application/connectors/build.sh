#!/bin/bash

#Build go files

echo "Building connector service"

bin_sourceconnectors=$bin_service/sourceconnectors
src_sourceconnectors=$src_service/sourceconnectors

mkdir -p $bin_sourceconnectors

echo "Building web69connector plugin"
go build -ldflags=-w -buildmode=plugin -o $bin_sourceconnectors/web69connector.so $src_sourceconnectors/web69connector/connector.go
echo "Building connectorservice"
go build -ldflags=-w -o $bin_service/connectorservice $src_service/main.go