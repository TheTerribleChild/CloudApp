GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get -u

all: test build

get:
	$(GOGET) github.com/dgrijalva/jwt-go
	$(GOGET) github.com/gomodule/redigo/redis
	$(GOGET) github.com/google/uuid
	$(GOGET) github.com/grpc-ecosystem/go-grpc-middleware
	$(GOGET) github.com/streadway/amqp
	$(GOGET) github.com/spf13/viper

proto: 
	@echo Generating proto file
	protoc -I internal/model/ internal/model/proto/*.proto --go_out=plugins=grpc:${GOPATH}/src

config:
	cp configs/agent/* bin/Agent/
	cp configs/agentservice/* bin/AgentService/
	cp configs/storageservice/* bin/StorageService/

compile: proto

	@echo Building Management Service
	go build -o bin/AgentService/AgentService cmd/agentservice/*.go
	mkdir -p bin/AgentService/logs

	@echo Building Storage Service
	go build -o bin/StorageService/StorageService cmd/storageservice/*.go
	mkdir -p bin/StorageService/logs
	
	@echo Building Agent
	go build -o bin/Agent/Agent cmd/agent/*.go
	mkdir -p bin/Agent/logs

build: proto compile config

test: 
		$(GOTEST) -v ./...
clean: 
		$(GOCLEAN)
		rm -rf bin/*
run:
		$(GOBUILD) -o $(BINARY_NAME) -v ./...
		./$(BINARY_NAME)