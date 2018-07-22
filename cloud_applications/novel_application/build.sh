#!/bin/bash

echo "Building novel applications"

services=("connectors")

for services in ${services[@]}
do
    bin_service=$bin_app/$services
    src_service=$src_app/$services
    mkdir -p $bin_service
    . ./$src_service/build.sh
done

go build -ldflags=-w -o $bin_app/tester $src_app/tester/main.go
