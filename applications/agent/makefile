GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get -u

get:
	$(GOGET) google.golang.org/grpc
	$(GOGET) github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	$(GOGET) github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	$(GOGET) github.com/golang/protobuf/protoc-gen-go
	$(GOGET) github.com/lib/pq
	$(GOGET) github.com/Masterminds/squirrel
	$(GOGET) github.com/jmoiron/sqlx
	$(GOGET) github.com/favadi/protoc-go-inject-tag 
	$(GOGET) github.com/patrickmn/go-cache

config:
	cp configs/* bin/Agent

compile:
	gofmt -s -w .
	go build -o bin/Agent/Agent cmd/*.go
