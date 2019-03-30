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
	$(GOGET) github.com/jmoiron/sqlx
	$(GOGET) github.com/favadi/protoc-go-inject-tag 
	$(GOGET) github.com/patrickmn/go-cache

proto: 
	@echo Generating proto file
	@echo Building common proto file
	protoc -I${GOPATH}/src/theterriblechild/CloudApp/common ${GOPATH}/src/theterriblechild/CloudApp/common/model/*.proto  --go_out=plugins=grpc:${GOPATH}/src
	protoc-go-inject-tag -input=${GOPATH}/src/theterriblechild/CloudApp/common/model/model.pb.go > /dev/null
	protoc -I${GOPATH}/src/theterriblechild/CloudApp/common ${GOPATH}/src/theterriblechild/CloudApp/common/*.proto  --go_out=plugins=grpc:${GOPATH}/src

	@echo Building Admin Service
	protoc -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I model/ -I ${GOPATH}/src/theterriblechild/CloudApp/common model/proto/adminservice.proto --go_out=plugins=grpc:${GOPATH}/src
	protoc -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I model/ -I ${GOPATH}/src/theterriblechild/CloudApp/common model/proto/adminservice.proto --grpc-gateway_out ${GOPATH}/src

config:
	cp configs/* bin/AdminService

compile: proto
	go build -o bin/AdminService/AdminService cmd/*.go