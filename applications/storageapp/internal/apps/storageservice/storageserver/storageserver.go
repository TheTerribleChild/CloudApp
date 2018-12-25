package storageserver

import (
	"time"
	// "encoding/json"
	//"fmt"

	cldstrg "github.com/TheTerribleChild/CloudApp/applications/storageapp/internal/model"
	contextutil "github.com/TheTerribleChild/CloudApp/commons/utils/contextutil"
	accesstoken "github.com/TheTerribleChild/CloudApp/applications/storageapp/internal/common/auth/accesstoken"

	//"github.com/golang/protobuf/proto"
	// "golang.org/x/net/netutil"
	//"github.com/go-stomp/stomp"
	"google.golang.org/grpc"
	//"google.golang.org/grpc/codes"
	"log"
	"net"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type StorageServer struct{}

func (instance *StorageServer) InitializeServer() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	chainstream := grpc_middleware.ChainStreamServer(instance.StorageServerStreamLogInterceptor, instance.StorageServerStreamAuthInterceptor)
	s := grpc.NewServer(grpc.MaxRecvMsgSize(11*1024*1024), grpc.StreamInterceptor(chainstream))
	cldstrg.RegisterStorageServiceServer(s, instance)
	log.Println("Initializing Storage Server.")
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (instance *StorageServer) StorageServerStreamAuthInterceptor(srv interface{}, stream grpc.ServerStream,
	info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	requestContext := stream.Context()
	jwtStr, err := contextutil.GetAuth(requestContext)
	if len(jwtStr) == 0 || err != nil {
		log.Println("Missing authorization header.")
		return status.Error(codes.PermissionDenied, "Missing authorization header")
	}
	switch info.FullMethod{
	case "/cloudstorage.StorageServer/DownloadFile":
		downloadToken := accesstoken.UploadDownloadToken{}
		err = accesstoken.BuildAgentPollTokenAuthentiactor("abc").TokenAuthenticator.AuthenticateAndDecodeJWTString(jwtStr, &downloadToken)
		requestContext = contextutil.SetUserId(requestContext, downloadToken.UserId)
		requestContext = contextutil.SetAgentId(requestContext, downloadToken.AgentId)
		break
	case "/cloudstorage.StorageServer/UploadFile":
		uploadToken := accesstoken.UploadDownloadToken{}
		err = accesstoken.BuildAgentPollTokenAuthentiactor("abc").TokenAuthenticator.AuthenticateAndDecodeJWTString(jwtStr, &uploadToken)
		requestContext = contextutil.SetUserId(requestContext, uploadToken.UserId)
		requestContext = contextutil.SetAgentId(requestContext, uploadToken.AgentId)
		break
	default:
		return status.Error(codes.InvalidArgument, "Invalid request.")
	}
	if err != nil {
		log.Println("Unauthorized request." + err.Error())
		return status.Error(codes.PermissionDenied, "Unauthorized request.")
	}
	newStream := grpc_middleware.WrapServerStream(stream)
   	newStream.WrappedContext = requestContext
	handler(srv, newStream)
	return nil
}

func (instance *StorageServer) StorageServerStreamLogInterceptor(srv interface{}, stream grpc.ServerStream,
	info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	start := time.Now()
	toe, _ := contextutil.GetToe(stream.Context())
	log.Printf("[toe=%s]Request to: %s", toe, info.FullMethod)
	handler(srv, stream)
	log.Printf("[toe=%s]Request completed. Took: %dms", toe, time.Since(start)/time.Millisecond)
	return nil
}
