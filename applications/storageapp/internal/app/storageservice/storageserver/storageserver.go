package storageserver

import (
	cldstrg "theterriblechild/CloudApp/applications/storageapp/internal/model"
	//contextutil "theterriblechild/CloudApp/tools/utils/contextutil"
	accesstoken "theterriblechild/CloudApp/applications/storageapp/internal/tools/auth/accesstoken"
	auth "theterriblechild/CloudApp/tools/auth/accesstoken"
	grpcutil "theterriblechild/CloudApp/tools/utils/grpc"
	"github.com/spf13/viper"

	//"github.com/golang/protobuf/proto"
	// "golang.org/x/net/netutil"
	//"github.com/go-stomp/stomp"
	"google.golang.org/grpc"
	//"google.golang.org/grpc/codes"
	"log"
	"net"
	"fmt"
	"os"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type StorageServer struct{}

var (
	storageLocation string
)

func (instance *StorageServer) InitializeServer() {
	storageLocation = viper.GetString("storagePath")
	grpcURL := fmt.Sprintf("%s:%d", viper.GetString("storageServer.accept"), viper.GetInt("storageServer.port"))
	maxRecvMsgSize := viper.GetInt("storageServer.maxMessageSize")
	if len(storageLocation) == 0 {
		log.Fatal("Storage Path missing")
	}
	stat, err := os.Stat(storageLocation)
	if err != nil && os.IsNotExist(err) {
		os.Mkdir(storageLocation, os.ModePerm)
	} else if err != nil {
		log.Fatalf("%s is not a valid storage location.", storageLocation)
	} else if stat != nil {
		if !stat.IsDir() {
			log.Fatalf("%s is not a directory", storageLocation)
		}
	}
	
	lis, err := net.Listen("tcp", grpcURL)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	chainstream := grpcutil.GetChainStreamInterceptorBuilder().AddLogInterceptor().AddAuthInterceptor(instance.authenticateRequest).Build()
	s := grpc.NewServer(grpc.MaxRecvMsgSize(maxRecvMsgSize), grpc.StreamInterceptor(chainstream))
	cldstrg.RegisterStorageServiceServer(s, instance)
	log.Printf("Initializing Storage Server. Listening on '%s'", grpcURL)
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (instance *StorageServer) authenticateRequest(method string, jwtStr string) error {
	log.Println("agent auth: " + method)
	log.Println(jwtStr)
	var tokenAuthenticator auth.TokenAuthenticator
	switch method {
	case "/cloudstorage.StorageService/DownloadFile":
		tokenAuthenticator = accesstoken.BuildDownloadTokenAuthentiactor("abc")
		break
	case "/cloudstorage.StorageService/UploadFile":
		tokenAuthenticator = accesstoken.BuildUploadTokenAuthentiactor("abc")
		break
	default:
		return status.Error(codes.InvalidArgument, "Invalid request.")
	}
	return tokenAuthenticator.AuthenticateJWTStringWithPermission(jwtStr)
}