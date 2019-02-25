GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get -u

get:
	$(GOGET) github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	$(GOGET) github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	$(GOGET) github.com/golang/protobuf/protoc-gen-go
	$(GOGET) github.com/lib/pq
	$(GOGET) github.com/Masterminds/squirrel

proto: 
	@echo Generating proto file
	protoc -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I model/ model/proto/*.proto --go_out=plugins=grpc:${GOPATH}/src
	protoc -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I model/ model/proto/*.proto --grpc-gateway_out ${GOPATH}/src

compile: proto
	go build -o bin/AdminService/AdminService cmd/*.go