GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

all: test build

proto: 
	@echo Generating proto file
	protoc -I internal/model/ internal/model/proto/*.proto --go_out=plugins=grpc:${GOPATH}/src

build: proto

	@echo Building Management Service
	go build -o bin/ManagementService/ManagementService cmd/managementservice/*.go
	mkdir -p bin/ManagementService/logs

	@echo Building Storage Service
	go build -o bin/StorageService/StorageService cmd/storageservice/*.go
	mkdir -p bin/StorageService/logs
	
	@echo Building Agent
	go build -o bin/Agent/Agent cmd/agent/*.go
	mkdir -p bin/Agent/logs

test: 
		$(GOTEST) -v ./...
clean: 
		$(GOCLEAN)
		rm -rf bin/*
run:
		$(GOBUILD) -o $(BINARY_NAME) -v ./...
		./$(BINARY_NAME)
deps:
		$(GOGET) github.com/markbates/goth
		$(GOGET) github.com/markbates/pop