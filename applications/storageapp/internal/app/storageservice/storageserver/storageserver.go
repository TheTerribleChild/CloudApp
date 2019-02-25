package storageserver

import (
	cldstrg "theterriblechild/CloudApp/applications/storageapp/internal/model"
	//contextutil "theterriblechild/CloudApp/tools/utils/contextutil"
	accesstoken "theterriblechild/CloudApp/applications/storageapp/internal/tools/auth/accesstoken"
	auth "theterriblechild/CloudApp/tools/auth/accesstoken"
	grpcutil "theterriblechild/CloudApp/tools/utils/grpc"

	//"github.com/golang/protobuf/proto"
	// "golang.org/x/net/netutil"
	//"github.com/go-stomp/stomp"
	"google.golang.org/grpc"
	//"google.golang.org/grpc/codes"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type StorageServer struct {
	StorageLocation           string
	CacheLocation             string
	SecretKey                 string
	StorageServerTrustedIP    string
	StorageServerPort         int
	MaxRecvMsgSize            int
	CompressStorage           bool
	EncryptStorage            bool
	tokenAuthenticatorBuilder accesstoken.TokenAuthenticatorBuilder
}

func (instance *StorageServer) InitializeServer() {
	grpcURL := fmt.Sprintf("%s:%d", instance.StorageServerTrustedIP, instance.StorageServerPort)
	instance.tokenAuthenticatorBuilder = accesstoken.TokenAuthenticatorBuilder{Secret: instance.SecretKey}
	if len(instance.StorageLocation) == 0 {
		log.Fatal("Storage Path missing")
	}

	stat, err := os.Stat(instance.StorageLocation)
	if err != nil && os.IsNotExist(err) {
		os.Mkdir(instance.StorageLocation, os.ModePerm)
	} else if err != nil {
		log.Fatalf("%s is not a valid storage location.", instance.StorageLocation)
	} else if stat != nil {
		if !stat.IsDir() {
			log.Fatalf("%s is not a directory", instance.StorageLocation)
		}
	}
	if len(instance.CacheLocation) == 0 {
		instance.CacheLocation = filepath.Join(instance.StorageLocation, "StorageCache")
	}
	stat, err = os.Stat(instance.CacheLocation)
	if err != nil && os.IsNotExist(err) {
		os.Mkdir(instance.StorageLocation, os.ModePerm)
	} else if err != nil {
		log.Fatalf("%s is not a valid cache location.", instance.StorageLocation)
	} else if stat != nil {
		if !stat.IsDir() {
			log.Fatalf("%s is not a directory", instance.StorageLocation)
		}
	}

	lis, err := net.Listen("tcp", grpcURL)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	chainstream := grpcutil.GetChainStreamInterceptorBuilder().AddLogInterceptor().AddAuthInterceptor(instance.authenticateRequest).AddContextInjector(instance.injectContext).Build()
	s := grpc.NewServer(grpc.MaxRecvMsgSize(instance.MaxRecvMsgSize), grpc.StreamInterceptor(chainstream))
	cldstrg.RegisterStorageServiceServer(s, instance)
	log.Printf("Initializing Storage Server. Listening on '%s'", grpcURL)
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (instance *StorageServer) Ping(ctx context.Context, msg *cldstrg.Empty) (*cldstrg.Empty, error) {
	log.Println("Request to ping")
	return &cldstrg.Empty{}, nil
}

func (instance *StorageServer) authenticateRequest(method string, jwtStr string) error {
	log.Println("agent auth: " + method)
	var tokenAuthenticator auth.TokenAuthenticator
	switch method {
	case "/cloudstorage.StorageService/DownloadFile":
		tokenAuthenticator = instance.tokenAuthenticatorBuilder.BuildFileReadTokenAuthenticator()
		break
	case "/cloudstorage.StorageService/UploadFile":
		tokenAuthenticator = instance.tokenAuthenticatorBuilder.BuildFileWriteTokenAuthenticator()
		break
	case "/cloudstorage.StorageService/Ping":
		return nil
	default:
		return status.Error(codes.InvalidArgument, "Invalid request.")
	}
	return tokenAuthenticator.AuthenticateJWTStringWithPermission(jwtStr)
}

func (instance *StorageServer) injectContext(method string, streamContext *context.Context) error {

	switch method {
	case "/cloudstorage.StorageService/DownloadFile":
		break
	case "/cloudstorage.StorageService/UploadFile":
		break
	case "/cloudstorage.StorageService/Ping":
		return nil
	default:
		return status.Error(codes.InvalidArgument, "Invalid request.")
	}
	return nil
}