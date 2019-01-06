GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

all: test build

get:
	go get github.com/spf13/viper

proto: 
	@echo Generating proto file
	protoc -I internal/model/ internal/model/proto/*.proto --go_out=plugins=grpc:${GOPATH}/src

config:
	cp configs/agent/* bin/Agent/
	cp configs/agentservice/* bin/AgentService/
	cp configs/storageservice/* bin/StorageService/

build: proto config

	@echo Building Management Service
	go build -o bin/AgentService/AgentService cmd/agentservice/*.go
	mkdir -p bin/AgentService/logs

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