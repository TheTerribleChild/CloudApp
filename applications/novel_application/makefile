GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

all: test build

build:
	@echo Generating proto file
	protoc -I internal/model -I internal/app/connectorservice/proto internal/app/connectorservice/proto/*.proto --go_out=plugins=grpc:${GOPATH}/src
	protoc -I internal/model -I internal/app/novelapplication/proto internal/app/novelapplication/proto/*.proto --go_out=plugins=grpc:${GOPATH}/src
	protoc -I internal/model internal/model/*.proto --go_out=${GOPATH}/src

	@echo Building connector
	go build -ldflags=-w -buildmode=plugin -o bin/sourceconnectors/web69connector.so internal/app/connectorservice/sourceconnectors/web69connector/connector.go
	go build -o bin/connectorservice cmd/connectorservice/*.go
	
	@echo Building novelapplication
	go build -o bin/novelapplication cmd/novelapplication/*.go
test: 
		$(GOTEST) -v ./...
clean: 
		$(GOCLEAN)
		rm -f $(BINARY_NAME)
		rm -f $(BINARY_UNIX)
run:
		$(GOBUILD) -o $(BINARY_NAME) -v ./...
		./$(BINARY_NAME)
deps:
		$(GOGET) github.com/markbates/goth
		$(GOGET) github.com/markbates/pop