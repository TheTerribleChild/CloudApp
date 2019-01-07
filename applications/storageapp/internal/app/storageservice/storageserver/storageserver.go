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
	"log"
	"net"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type StorageServer struct{}

func (instance *StorageServer) InitializeServer() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	chainstream := grpcutil.GetChainStreamInterceptorBuilder().AddLogInterceptor().AddAuthInterceptor(instance.authenticateRequest).Build()
	s := grpc.NewServer(grpc.MaxRecvMsgSize(11*1024*1024), grpc.StreamInterceptor(chainstream))
	cldstrg.RegisterStorageServiceServer(s, instance)
	log.Println("Initializing Storage Server.")
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
