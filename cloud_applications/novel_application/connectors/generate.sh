#!/bin/bash


protoc -I service/ service/connectorservice.proto --go_out=plugins=grpc:service